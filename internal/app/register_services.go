package app

import (
	"films-api/internal/api/services/film"
)

// registerServices register services in app struct.
func (a *App) registerServices() {
	a.filmService = film.NewService(a.filmPostgres, a.logger)
}
