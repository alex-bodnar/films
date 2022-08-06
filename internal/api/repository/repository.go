package repository

import (
	"context"

	"films-api/internal/api/domain/film"
	"films-api/internal/api/domain/statistics"
)

//go:generate mockgen -source repository.go -destination ./repository_mock.go -package repository

type (
	// FilmPostgres - describe an interface for working with film database models.
	FilmPostgres interface {
		GetByName(ctx context.Context, name string) (film.FilmList, error)
	}

	// Statistics - describe an interface for working with statistics database models.
	Statistics interface {
		GetByRequest(ctx context.Context, req string) (statistics.FilmStatistic, error)
		GetAll(ctx context.Context, limit, offset uint64) (statistics.FilmStatisticList, error)
		Create(ctx context.Context, stat statistics.FilmStatistic) error
		Update(ctx context.Context, stat statistics.FilmStatistic) error
	}
)
