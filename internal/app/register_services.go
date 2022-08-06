package app

import (
	"context"

	"films-api/internal/api/services/film"
	"films-api/internal/api/services/statistics"
)

// registerServices register services in app struct.
func (a *App) registerServices(ctx context.Context) {
	a.statisticsService = statistics.NewService(ctx, a.statisticsRepo, a.logger)
	a.filmService = film.NewService(
		a.filmPostgres,
		a.filmRedisCache,
		a.filmLocalCache,
		a.statisticsService,
		a.logger,
	)
}
