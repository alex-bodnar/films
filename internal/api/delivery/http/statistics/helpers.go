package statistics

import (
	"films-api/internal/api/delivery/http"
	"films-api/internal/api/domain/statistics"
)

type (
	// StatisticList response model with pagination and data.
	StatisticList struct {
		Pagination     http.Pagination     `json:"pagination"`
		FilmStatistics []statisticResponse `json:"film_statistics"`
	}

	// FilmStatistic response model.
	statisticResponse struct {
		ID         int64  `json:"id"`
		Request    string `json:"request"`
		TimeDB     int64  `json:"time_db,omitempty"`
		TimeRedis  int64  `json:"time_redis,omitempty"`
		TimeMemory int64  `json:"time_memory,omitempty"`
	}
)

// toStatisticResponse converts domain model to response model.
func toStatisticResponse(domain statistics.FilmStatistic) statisticResponse {
	return statisticResponse{
		ID:         domain.ID,
		Request:    domain.Request,
		TimeDB:     domain.TimeDB.Milliseconds(),
		TimeRedis:  domain.TimeRedis.Milliseconds(),
		TimeMemory: domain.TimeMemory.Milliseconds(),
	}
}

// toStatisticList converts domain model to response model.
func toStatisticList(domain statistics.FilmStatisticList, pagination http.Pagination) StatisticList {
	result := make([]statisticResponse, 0, len(domain.FilmStatistics))
	for _, filmStat := range domain.FilmStatistics {
		result = append(result, toStatisticResponse(filmStat))
	}

	pagination.Total = domain.Total

	return StatisticList{
		Pagination:     pagination,
		FilmStatistics: result,
	}
}
