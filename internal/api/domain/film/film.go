package film

import "time"

type (
	// FilmList type alias for a slice of film.
	FilmList []Film

	// Film domain model.
	Film struct {
		ID              int64
		Name            string
		Description     string
		ReleaseYear     int64
		Language        string
		RentalDuration  int64
		RentalRate      float64
		Length          int64
		ReplacementCost float64
		Rating          string
		SpecialFeatures []string
		LastUpdate      time.Time

		Actors     []Actor
		Categories []Category
	}

	// Actor domain model.
	Actor struct {
		ID       int64
		Name     string
		LastName string
	}

	// Category domain model.
	Category struct {
		ID   int64
		Name string
	}
)
