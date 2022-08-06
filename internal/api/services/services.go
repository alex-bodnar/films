package services

import (
	"context"

	"films-api/internal/api/domain/film"
	"films-api/internal/api/domain/statistics"
)

//go:generate mockgen -source services.go -destination ./services_mock.go -package services
type (
	// FilmService - describe an interface for working with film.
	FilmService interface {
		GetByName(ctx context.Context, name string) (film.FilmList, error)
	}

	// Statistics - describe an interface for working with statistics.
	Statistics interface {
		Update(stat statistics.FilmStatistic)
		GetAll(ctx context.Context, limit, offset uint64) (statistics.FilmStatisticList, error)
	}
)
