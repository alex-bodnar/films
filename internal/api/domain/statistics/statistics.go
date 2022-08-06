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

		startTime time.Time
	}
)

// NewFilmStatistic - constructor.
func NewFilmStatistic(request string) FilmStatistic {
	return FilmStatistic{
		Request:   request,
		startTime: time.Now(),
	}
}

// ResetTime - reset start time.
func (f *FilmStatistic) ResetTime() {
	f.startTime = time.Now()
}

// FinishDB - set database request time.
func (f *FilmStatistic) FinishDB() {
	f.TimeDB = time.Since(f.startTime)
	f.ResetTime()
}

// FinishRedis - set redis request time.
func (f *FilmStatistic) FinishRedis() {
	f.TimeRedis = time.Since(f.startTime)
	f.ResetTime()
}

// FinishMemory - set memory request time.
func (f *FilmStatistic) FinishMemory() {
	f.TimeMemory = time.Since(f.startTime)
	f.ResetTime()
}

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
