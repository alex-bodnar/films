package app

import (
	"films-api/internal/api/delivery/http/film"
	"films-api/internal/api/delivery/http/status"
)

func (a *App) registerHTTPHandlers() {
	a.statusHTTPHandler = status.NewHandler(
		a.meta.Info.AppName,
		a.meta.Info.Tag,
		a.meta.Info.Version,
		a.meta.Info.Commit,
		a.meta.Info.Date,
		a.meta.Info.FortuneCookie,
	)

	a.filmHTTPHandler = film.NewHandler(a.filmService)
}
