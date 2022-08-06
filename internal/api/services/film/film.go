package film

import (
	"context"
	"errors"
	"time"

	"films-api/internal/api/domain/film"
	"films-api/internal/api/domain/statistics"
	"films-api/internal/api/repository"
	"films-api/internal/api/services"
	"films-api/pkg/errs"
	"films-api/pkg/log"
)

var _ services.FilmService = &Service{}

// Service - defines film service struct.
type Service struct {
	logger            log.Logger
	filmPostgres      repository.FilmPostgres
	statisticsService services.Statistics
}

// NewService - constructor.
func NewService(filmPostgres repository.FilmPostgres, statistics services.Statistics, logger log.Logger) *Service {
	return &Service{
		filmPostgres:      filmPostgres,
		statisticsService: statistics,
		logger:            logger,
	}
}

// GetByName - get film by name.
func (s Service) GetByName(ctx context.Context, name string) (film.FilmList, error) {
	startTime := time.Now()

	filmList, err := s.filmPostgres.GetByName(ctx, name)
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			s.logger.Debug(err)
			return nil, err
		}

		s.logger.Error(err)

		return nil, errs.Internal{}
	}

	filmStatistic := statistics.FilmStatistic{
		Request: statistics.FilmRout + name,
		TimeDB:  time.Since(startTime),
	}

	s.statisticsService.Update(filmStatistic)

	return filmList, nil
}
