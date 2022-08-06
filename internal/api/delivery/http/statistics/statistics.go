package statistics

import (
	"github.com/gofiber/fiber/v2"

	"films-api/internal/api/delivery/http"
	"films-api/internal/api/services"
	"films-api/pkg/http/responder"
)

// Handler - define http handler struct for handling statistics requests.
type Handler struct {
	responder.Responder
	statisticsService services.Statistics
}

// NewHandler - constructor.
func NewHandler(statisticsService services.Statistics) *Handler {
	return &Handler{
		statisticsService: statisticsService,
	}
}

func (h Handler) GetAll(ctx *fiber.Ctx) error {
	var pagination http.Pagination
	if err := pagination.SetPagination(ctx); err != nil {
		return err
	}

	statistics, err := h.statisticsService.GetAll(ctx.Context(), pagination.Limit, pagination.Offset)
	if err != nil {
		return err
	}

	return h.Respond(ctx, fiber.StatusOK, toStatisticList(statistics, pagination))
}
