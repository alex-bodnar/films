package app

import (
	filmPostgres "films-api/internal/api/repository/film_postgres"
	"films-api/internal/api/repository/statistics"
)

func (a *App) registerRepositories() {
	a.filmPostgres = filmPostgres.NewRepository(a.config.Storage.Postgres.QueryTimeout, a.db)
	a.statisticsRepo = statistics.NewRepository(a.config.Storage.Redis.QueryTimeout, a.db)
}
