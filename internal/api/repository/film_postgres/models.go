package film_postgres

import (
	"films-api/internal/api/domain/film"
	"time"

	"github.com/lib/pq"
)

type (
	// filmList type alias for a slice of film.
	filmList []filmDB

	// film database model.
	filmDB struct {
		ID              int64          `db:"film_id"`
		Name            string         `db:"title"`
		Description     string         `db:"description"`
		ReleaseYear     int64          `db:"release_year"`
		Language        string         `db:"language"`
		RentalDuration  int64          `db:"rental_duration"`
		RentalRate      float64        `db:"rental_rate"`
		Length          int64          `db:"length"`
		ReplacementCost float64        `db:"replacement_cost"`
		Rating          string         `db:"rating"`
		SpecialFeatures pq.StringArray `db:"special_features"`
		LastUpdate      time.Time      `db:"last_update"`

		Actors     []actor
		Categories []category
	}

	// actor database model.
	actor struct {
		ID       int64  `db:"actor_id"`
		Name     string `db:"first_name"`
		LastName string `db:"last_name"`
	}

	// category database model.
	category struct {
		ID   int64  `db:"category_id"`
		Name string `db:"name"`
	}
)

// toDomain converts database model to domain model.
func (a actor) toDomain() film.Actor {
	return film.Actor{
		ID:       a.ID,
		Name:     a.Name,
		LastName: a.LastName,
	}
}

// toDomain converts database model to domain model.
func (c category) toDomain() film.Category {
	return film.Category{
		ID:   c.ID,
		Name: c.Name,
	}
}

// toDomain converts database model to domain model.
func (f filmDB) toDomain() film.Film {
	actors := make([]film.Actor, 0, len(f.Actors))
	for _, a := range f.Actors {
		actors = append(actors, a.toDomain())
	}

	categories := make([]film.Category, 0, len(f.Categories))
	for _, c := range f.Categories {
		categories = append(categories, c.toDomain())
	}

	return film.Film{
		ID:              f.ID,
		Name:            f.Name,
		Description:     f.Description,
		ReleaseYear:     f.ReleaseYear,
		Language:        f.Language,
		RentalDuration:  f.RentalDuration,
		RentalRate:      f.RentalRate,
		Length:          f.Length,
		ReplacementCost: f.ReplacementCost,
		Rating:          f.Rating,
		SpecialFeatures: f.SpecialFeatures,
		LastUpdate:      f.LastUpdate,
		Actors:          actors,
		Categories:      categories,
	}
}

// toDomain converts database model to domain model.
func (f filmList) toDomain() []film.Film {
	films := make([]film.Film, 0, len(f))
	for _, film := range f {
		films = append(films, film.toDomain())
	}

	return films
}
