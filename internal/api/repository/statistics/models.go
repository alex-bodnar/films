package statistics

import (
	"time"

	"films-api/internal/api/domain/statistics"
)

type (
	// filmStatistic database model with total field.
	filmStatisticList struct {
		FilmStatistics []filmStatistic
		Total          uint64
	}

	// filmStatistic database model.
	filmStatistic struct {
		ID         int64         `db:"id"`
		Request    string        `db:"request"`
		TimeDB     time.Duration `db:"time_db"`
		TimeRedis  time.Duration `db:"time_redis"`
		TimeMemory time.Duration `db:"time_memory"`
	}
)

// toDatabase converts domain model to database model.
func toDatabase(stat statistics.FilmStatistic) filmStatistic {
	return filmStatistic{
		ID:         stat.ID,
		Request:    stat.Request,
		TimeDB:     stat.TimeDB,
		TimeRedis:  stat.TimeRedis,
		TimeMemory: stat.TimeMemory,
	}
}

// toDomain converts database model to domain model.
func (f filmStatistic) toDomain() statistics.FilmStatistic {
	return statistics.FilmStatistic{
		ID:         f.ID,
		Request:    f.Request,
		TimeDB:     f.TimeDB,
		TimeRedis:  f.TimeRedis,
		TimeMemory: f.TimeMemory,
	}
}

// toDomain converts database model to domain model.
func (f filmStatisticList) toDomain() statistics.FilmStatisticList {
	result := make([]statistics.FilmStatistic, 0, len(f.FilmStatistics))
	for _, filmStat := range f.FilmStatistics {
		result = append(result, filmStat.toDomain())
	}

	return statistics.FilmStatisticList{
		FilmStatistics: result,
		Total:          f.Total,
	}
}
