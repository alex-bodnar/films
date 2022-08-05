package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type (
	// StatusHTTP – describes an interface for work with service status over HTTP.
	StatusHTTP interface {
		CheckStatus(ctx *fiber.Ctx) error
	}

	// FilmHTTP – describes an interface for work with film over HTTP.
	FilmHTTP interface {
		GetByName(ctx *fiber.Ctx) error
	}
)
