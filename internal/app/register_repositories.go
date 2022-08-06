package app

import (
	filmLocal "films-api/internal/api/repository/film_local"
	filmPostgres "films-api/internal/api/repository/film_postgres"
	filmRedis "films-api/internal/api/repository/film_redis"
	"films-api/internal/api/repository/statistics"
)

func (a *App) registerRepositories() {
	a.filmPostgres = filmPostgres.NewRepository(a.config.Storage.Postgres.QueryTimeout, a.db)
	a.statisticsRepo = statistics.NewRepository(a.config.Storage.Redis.QueryTimeout, a.db)
	a.filmRedisCache = filmRedis.NewRepository(a.config.Extra.RedisCache.TimeLive, a.rdb)
	a.filmLocalCache = filmLocal.NewRepository(a.config.Extra.LocalCache)
}
