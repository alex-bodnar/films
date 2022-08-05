package film

import (
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"

	"films-api/internal/api/services"
	"films-api/pkg/errs"
	"films-api/pkg/http/responder"
)

// Handler - define http handler struct for handling film requests.
type Handler struct {
	responder.Responder
	filmService services.FilmService
}

// NewHandler - constructor.
func NewHandler(filmService services.FilmService) *Handler {
	return &Handler{
		filmService: filmService,
	}
}

// GetByName - handle http request for getting film by name.
func (h *Handler) GetByName(ctx *fiber.Ctx) error {
	filmName := ctx.Params(ParameterTitle)
	if filmName == "" {
		return errs.BadRequest{Cause: ParameterTitle + "::is_required"}
	}

	decodedFilmName, err := url.QueryUnescape(filmName)
	if err != nil {
		return errs.BadRequest{Cause: ParameterTitle + "::is_invalid"}
	}

	films, err := h.filmService.GetByName(ctx.Context(), strings.ToLower(decodedFilmName))
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusOK, toFilmListResponse(films))
}
