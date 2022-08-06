package statistics

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"

	"films-api/internal/api/domain/statistics"
	"films-api/internal/api/repository"
	"films-api/pkg/errs"
)

var _ repository.Statistics = &Repository{}

// Repository implements repository.Statistics
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

// GetByRequest returns statistic by name.
func (r Repository) GetByRequest(ctx context.Context, req string) (statistics.FilmStatistic, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query := `SELECT id, request, time_db, time_redis, time_memory 
			  FROM responce_time_log WHERE request = $1`

	var result filmStatistic
	if err := r.db.GetContext(ctx, &result, query, req); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return statistics.FilmStatistic{}, errs.NotFound{What: "request in statistic table"}
		}

		return statistics.FilmStatistic{}, errs.Internal{Cause: err.Error()}
	}

	return result.toDomain(), nil
}

// GetAll returns all statistic.
func (r Repository) GetAll(ctx context.Context, limit, offset uint64) (statistics.FilmStatisticList, error) {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query := `SELECT count(*) FROM responce_time_log`

	var result filmStatisticList
	if err := r.db.GetContext(ctx, &result.Total, query); err != nil {
		return statistics.FilmStatisticList{}, errs.Internal{Cause: err.Error()}
	}

	if result.Total == 0 {
		return statistics.FilmStatisticList{}, errs.NotFound{What: "statistic in statistic table"}
	}

	query = `SELECT id, request, time_db, time_redis, time_memory 
			  FROM responce_time_log ORDER BY id LIMIT $1 OFFSET $2`

	if err := r.db.SelectContext(ctx, &result.FilmStatistics, query, limit, offset); err != nil {
		return statistics.FilmStatisticList{}, errs.Internal{Cause: err.Error()}
	}

	return result.toDomain(), nil
}

// Create insert new value to database.
func (r Repository) Create(ctx context.Context, stat statistics.FilmStatistic) error {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query := `INSERT INTO responce_time_log (
				request, time_db, time_redis, time_memory
			) VALUES (
				:request, :time_db, :time_redis, :time_memory)`

	if _, err := r.db.NamedExecContext(ctx, query, toDatabase(stat)); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}

// Update update value in database.
func (r Repository) Update(ctx context.Context, stat statistics.FilmStatistic) error {
	ctx, cancel := context.WithTimeout(ctx, r.queryTimeout)
	defer cancel()

	query := `UPDATE responce_time_log SET
				request = :request, time_db = :time_db, 
				time_redis = :time_redis, time_memory = :time_memory
			WHERE id = :id`

	if _, err := r.db.NamedExecContext(ctx, query, toDatabase(stat)); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}
