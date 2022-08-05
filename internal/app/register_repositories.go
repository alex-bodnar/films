package app

import (
	_ "github.com/jackc/pgx/stdlib"

	filmPostgres "films-api/internal/api/repository/film_postgres"
)

func (a *App) registerRepositories() {
	a.filmPostgres = filmPostgres.NewRepository(a.config.Storage.Postgres.QueryTimeout, a.db)
}
