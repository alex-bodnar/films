package film

import (
	"context"
	"errors"

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
	filmPostgres   repository.FilmPostgres
	filmRedisCache repository.FilmCache
	filmLocalCache repository.FilmCache

	statisticsService services.Statistics

	logger log.Logger
}

// NewService - constructor.
func NewService(
	filmPostgres repository.FilmPostgres,
	filmRedisCache repository.FilmCache,
	filmLocalCache repository.FilmCache,

	statistics services.Statistics,

	logger log.Logger,
) *Service {
	return &Service{
		filmPostgres:   filmPostgres,
		filmRedisCache: filmRedisCache,
		filmLocalCache: filmLocalCache,

		statisticsService: statistics,

		logger: logger,
	}
}

// GetByName - get film by name.
func (s Service) GetByName(ctx context.Context, name string) (filmList film.FilmList, err error) {
	filmStatistic := statistics.NewFilmStatistic(statistics.FilmRout + name)

	defer func() {
		if filmList != nil {
			s.statisticsService.Update(filmStatistic)
		}
	}()

	filmList, err = s.filmLocalCache.GetByName(ctx, name)
	switch {
	case err == nil:
		filmStatistic.FinishMemory()
		return filmList, nil
	case !errors.As(err, &errs.NotFound{}):
		s.logger.Error(err)
		return nil, errs.Internal{}
	}

	filmList, err = s.filmRedisCache.GetByName(ctx, name)
	switch {
	case err == nil:
		filmStatistic.FinishRedis()

		if err = s.filmLocalCache.SetByName(ctx, name, filmList); err != nil {
			s.logger.Error(err)
			return nil, errs.Internal{}
		}

		return filmList, nil
	case !errors.As(err, &errs.NotFound{}):
		s.logger.Error(err)
		return nil, errs.Internal{}
	}

	filmList, err = s.filmPostgres.GetByName(ctx, name)
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			s.logger.Debug(err)
			return nil, err
		}

		s.logger.Error(err)

		return nil, errs.Internal{}
	}

	filmStatistic.FinishDB()

	if err = s.filmLocalCache.SetByName(ctx, name, filmList); err != nil {
		s.logger.Error(err)
		return nil, errs.Internal{}
	}

	if err = s.filmRedisCache.SetByName(ctx, name, filmList); err != nil {
		s.logger.Error(err)
		return nil, errs.Internal{}
	}

	return filmList, nil
}
