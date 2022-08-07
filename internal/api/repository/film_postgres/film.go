package film_postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"films-api/internal/api/domain/film"
	"films-api/internal/api/repository"
	"films-api/pkg/errs"
)

var _ repository.FilmPostgres = &Repository{}

// Repository implements repository.FilmPostgres
type Repository struct {
	queryTimeout time.Duration
	db           *sqlx.DB
}

// NewRepository constructor.
func NewRepository(qt time.Duration, db *sqlx.DB) *Repository {
	return &Repository{
		queryTimeout: qt,
		db:           db,
	}
}

// GetByName returns film by name.
func (r Repository) GetByName(ctx context.Context, name string) (film.FilmList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query := `SELECT film_id, title, description, release_year, language.name AS language,
                  rental_duration, rental_rate, length, replacement_cost, rating, film.last_update, 
                  special_features
              FROM film 
              INNER JOIN language
                  ON film.language_id = language.language_id
              WHERE fulltext @@ plainto_tsquery('english', $1)`

	var result filmList
	if err := r.db.SelectContext(ctx, &result, query, name); err != nil {
		return film.FilmList{}, errs.Internal{Cause: err.Error()}
	}

	if len(result) == 0 {
		return film.FilmList{}, errs.NotFound{What: "film"}
	}

	actor := `SELECT film_actor.actor_id, actor.first_name, actor.last_name
              FROM film_actor
              INNER JOIN actor
                  ON actor.actor_id = film_actor.actor_id
              WHERE film_id = $1`

	category := `SELECT film_category.category_id, category.name
                 FROM film_category
                 INNER JOIN category
                     ON category.category_id = film_category.category_id
                 WHERE film_id = $1`

	for index := range result {
		if err := r.db.SelectContext(ctx, &result[index].Actors, actor, result[index].ID); err != nil {
			return film.FilmList{}, errs.Internal{Cause: err.Error()}
		}

		if err := r.db.SelectContext(ctx, &result[index].Categories, category, result[index].ID); err != nil {
			return film.FilmList{}, errs.Internal{Cause: err.Error()}
		}
	}

	return result.toDomain(), nil
}
