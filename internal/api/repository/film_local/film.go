package film_redis

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"

	"films-api/internal/api/domain/film"
	"films-api/internal/config"
	"films-api/pkg/errs"
)

// Repository implements repository.FilmCache
type Repository struct {
	redisCache *cache.Cache
	timeLive   time.Duration
}

// NewRepository constructor.
func NewRepository(conf config.LocalCache) *Repository {
	return &Repository{
		redisCache: cache.New(&cache.Options{LocalCache: cache.NewTinyLFU(conf.NumberOfRecords, conf.TimeLive)}),
		timeLive:   conf.TimeLive,
	}
}

// SetByName add new film to cache.
func (r Repository) SetByName(ctx context.Context, name string, data film.FilmList) error {
	item := cache.Item{
		Ctx:   ctx,
		Key:   name,
		Value: data,
		TTL:   r.timeLive,
	}

	if err := r.redisCache.Set(&item); err != nil {
		return errs.Internal{Cause: err.Error()}
	}

	return nil
}

// GetByName returns film by name.
func (r Repository) GetByName(ctx context.Context, name string) (film.FilmList, error) {
	result := film.FilmList{}
	if err := r.redisCache.Get(ctx, name, &result); err != nil {
		if err == cache.ErrCacheMiss {
			return film.FilmList{}, errs.NotFound{What: "film with key " + name}
		}

		return film.FilmList{}, errs.Internal{Cause: err.Error()}
	}

	return result, nil
}
