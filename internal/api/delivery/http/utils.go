package http

import (
	"github.com/gofiber/fiber/v2"

	"films-api/pkg/errs"
)

const (
	// DefaultLimit describes default quantity of entities to return in response.
	DefaultLimit = 20
)

type (
	// Pagination is a model to store Pagination parameters.
	Pagination struct {
		Offset uint64 `query:"offset" json:"offset"`
		Limit  uint64 `query:"limit" json:"limit"`
		Total  uint64 `query:"total" json:"total"`
	}
)

// SetPagination setups Pagination model.
// 	Keys are: "limit", "offset"
func (r *Pagination) SetPagination(ctx *fiber.Ctx) error {
	if err := ctx.QueryParser(r); err != nil {
		return errs.BadRequest{Cause: "[pagination]::invalid"}
	}

	if r.Limit == 0 {
		r.Limit = DefaultLimit
	}

	return nil
}
