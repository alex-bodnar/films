package repository

import (
	"context"

	"films-api/internal/api/domain/film"
)

//go:generate mockgen -source repository.go -destination ./repository_mock.go -package repository

type (
	// FilmPostgres - describe an interface for working with film database models.
	FilmPostgres interface {
		GetByName(ctx context.Context, name string) (film.FilmList, error)
	}
)
