package statistics

import (
	"films-api/internal/api/domain/statistics"
	"time"
)

type (
	// filmStatistic database model with total field.
	filmStatisticList struct {
		FilmStatistics []filmStatistic
		Total          uint64
	}

	// filmStatistic database model.
	filmStatistic struct {
		ID         int64  `db:"id"`
		Request    string `db:"request"`
		TimeDB     int64  `db:"time_db"`
		TimeRedis  int64  `db:"time_redis"`
		TimeMemory int64  `db:"time_memory"`
	}
)

// toDatabase converts domain model to database model.
func toDatabase(stat statistics.FilmStatistic) filmStatistic {
	return filmStatistic{
		ID:         stat.ID,
		Request:    stat.Request,
		TimeDB:     stat.TimeDB.Microseconds(),
		TimeRedis:  stat.TimeRedis.Microseconds(),
		TimeMemory: stat.TimeMemory.Microseconds(),
	}
}

// toDomain converts database model to domain model.
func (f filmStatistic) toDomain() statistics.FilmStatistic {
	return statistics.FilmStatistic{
		ID:         f.ID,
		Request:    f.Request,
		TimeDB:     time.Duration(int64(time.Microsecond) * f.TimeDB),
		TimeRedis:  time.Duration(int64(time.Microsecond) * f.TimeRedis),
		TimeMemory: time.Duration(int64(time.Microsecond) * f.TimeMemory),
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
