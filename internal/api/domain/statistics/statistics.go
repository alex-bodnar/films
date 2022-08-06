package statistics

import "time"

const (
	// FilmRout - film route.
	FilmRout = "/film/"
)

type (
	// FilmStatisticList describes list of statistic with their total number.
	FilmStatisticList struct {
		FilmStatistics []FilmStatistic
		Total          uint64
	}

	// FilmStatistic domain model.
	FilmStatistic struct {
		ID         int64
		Request    string
		TimeDB     time.Duration
		TimeRedis  time.Duration
		TimeMemory time.Duration
	}
)

// SetFilmStatisticValue - update non-zero value of statistic.
func (f *FilmStatistic) SetFilmStatisticValue(stat FilmStatistic) {
	if stat.TimeDB != 0 {
		f.TimeDB = stat.TimeDB
	}

	if stat.TimeRedis != 0 {
		f.TimeRedis = stat.TimeRedis
	}

	if stat.TimeMemory != 0 {
		f.TimeMemory = stat.TimeMemory
	}
}
