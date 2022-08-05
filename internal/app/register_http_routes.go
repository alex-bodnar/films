package app

import "github.com/gofiber/fiber/v2"

func (a *App) registerHTTPRoutes(app *fiber.App) {
	router := app.Group("/v1/films-api")
	router.Get("/status", a.statusHTTPHandler.CheckStatus)

	router.Get("/film/:title", a.filmHTTPHandler.GetByName)
}
