package film

import (
	"time"

	"films-api/internal/api/domain/film"
)

const (
	// ParameterTitle it is film name parameter.
	ParameterTitle = "title"
)

type (
	// filmListResponse - response model for films.
	filmListResponse struct {
		Total uint64         `json:"total"`
		Films []filmResponse `json:"films,omitempty"`
	}

	// filmResponse - response model for film.
	filmResponse struct {
		ID              int64     `json:"id"`
		Name            string    `json:"name"`
		Description     string    `json:"description,omitempty"`
		ReleaseYear     int64     `json:"release_year,omitempty"`
		Language        string    `json:"language,omitempty"`
		RentalDuration  int64     `json:"rental_duration,omitempty"`
		RentalRate      float64   `json:"rental_rate,omitempty"`
		Length          int64     `json:"length,omitempty"`
		ReplacementCost float64   `json:"replacement_cost,omitempty"`
		Rating          string    `json:"rating,omitempty"`
		SpecialFeatures []string  `json:"special_features,omitempty"`
		LastUpdate      time.Time `json:"last_update,omitempty"`

		Actors     []actor    `json:"actors,omitempty"`
		Categories []category `json:"categories,omitempty"`
	}

	// actor - response model for actor.
	actor struct {
		ID       int64  `json:"id"`
		Name     string `json:"name,omitempty"`
		LastName string `json:"last_name,omitempty"`
	}

	// category - response model for category.
	category struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)

// toActorResponse converts domain film.Actor to response actor model.
func toActorResponse(domain film.Actor) actor {
	return actor{
		ID:       domain.ID,
		Name:     domain.Name,
		LastName: domain.LastName,
	}
}

// toCategoryResponse converts domain film.Category to response category model.
func toCategoryResponse(domain film.Category) category {
	return category{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

// toFilmResponse converts domain film.Film to response film model.
func toFilmResponse(domain film.Film) filmResponse {
	actors := make([]actor, 0, len(domain.Actors))
	for _, a := range domain.Actors {
		actors = append(actors, toActorResponse(a))
	}

	categories := make([]category, 0, len(domain.Categories))
	for _, c := range domain.Categories {
		categories = append(categories, toCategoryResponse(c))
	}

	return filmResponse{
		ID:              domain.ID,
		Name:            domain.Name,
		Description:     domain.Description,
		ReleaseYear:     domain.ReleaseYear,
		Language:        domain.Language,
		RentalDuration:  domain.RentalDuration,
		RentalRate:      domain.RentalRate,
		Length:          domain.Length,
		ReplacementCost: domain.ReplacementCost,
		Rating:          domain.Rating,
		SpecialFeatures: domain.SpecialFeatures,
		LastUpdate:      domain.LastUpdate,
		Actors:          actors,
		Categories:      categories,
	}
}

// toFilmListResponse converts domain film.FilmList to response film list model.
func toFilmListResponse(domain film.FilmList) filmListResponse {
	films := make([]filmResponse, 0, len(domain))
	for _, f := range domain {
		films = append(films, toFilmResponse(f))
	}

	return filmListResponse{
		Total: uint64(len(domain)),
		Films: films,
	}
}
