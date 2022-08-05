package services

import (
	"context"

	"films-api/internal/api/domain/film"
)

//go:generate mockgen -source services.go -destination ./services_mock.go -package services
type (
	// FilmService - describe an interface for working with film.
	FilmService interface {
		GetByName(ctx context.Context, name string) (film.FilmList, error)
	}
)
