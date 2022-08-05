package film

import (
	"context"
	"errors"

	"films-api/internal/api/domain/film"
	"films-api/internal/api/repository"
	"films-api/internal/api/services"
	"films-api/pkg/errs"
	"films-api/pkg/log"
)

var _ services.FilmService = &Service{}

// Service - defines film service struct.
type Service struct {
	logger       log.Logger
	filmPostgres repository.FilmPostgres
}

// NewService - constructor.
func NewService(filmPostgres repository.FilmPostgres, logger log.Logger) *Service {
	return &Service{
		filmPostgres: filmPostgres,
		logger:       logger,
	}
}

// GetByName - get film by name.
func (s Service) GetByName(ctx context.Context, name string) (film.FilmList, error) {
	filmList, err := s.filmPostgres.GetByName(ctx, name)
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			s.logger.Debug(err)
			return nil, err
		}

		s.logger.Error(err)

		return nil, err
	}

	return filmList, nil
}
