package statistics

import (
	"context"
	"errors"

	"films-api/internal/api/domain/statistics"
	"films-api/internal/api/repository"
	"films-api/internal/api/services"
	"films-api/pkg/errs"
	"films-api/pkg/log"
)

var _ services.Statistics = &Service{}

// Service - defines statistics service struct.
type Service struct {
	ctx            context.Context
	logger         log.Logger
	statisticsRepo repository.Statistics
}

// NewService - constructor.
func NewService(ctx context.Context, statisticsRepo repository.Statistics, logger log.Logger) *Service {
	return &Service{
		ctx:            ctx,
		statisticsRepo: statisticsRepo,
		logger:         logger,
	}
}

// Update - update statistics in other goroutine.
func (s Service) Update(stat statistics.FilmStatistic) {
	go s.update(s.ctx, stat)
}

// update - update statistics.
func (s Service) update(ctx context.Context, stat statistics.FilmStatistic) {
	repoStatistic, err := s.statisticsRepo.GetByRequest(ctx, stat.Request)
	if err != nil {
		if !errors.As(err, &errs.NotFound{}) {
			s.logger.Error(err)
			return
		}

		if err := s.statisticsRepo.Create(ctx, stat); err != nil {
			s.logger.Error(err)
		}

		return
	}

	repoStatistic.SetFilmStatisticValue(stat)

	if err := s.statisticsRepo.Update(ctx, repoStatistic); err != nil {
		s.logger.Error(err)
	}
}

// GetAll - get all statistics.
func (s Service) GetAll(ctx context.Context, limit, offset uint64) (statistics.FilmStatisticList, error) {
	result, err := s.statisticsRepo.GetAll(ctx, limit, offset)
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			s.logger.Debug(err)
			return statistics.FilmStatisticList{}, errs.Empty{What: "statistics"}
		}

		s.logger.Error(err)
		return statistics.FilmStatisticList{}, errs.Internal{}
	}

	return result, nil
}
