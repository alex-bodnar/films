package statistics

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"

	"films-api/internal/api/domain/statistics"
	repo "films-api/internal/api/repository/statistics"
	"films-api/pkg/log"
)

func TestRepository_GetByRequest(t *testing.T) {
	logger := log.New()

	DB, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	if err != nil {
		logger.Fatalf("test init db failed: %s", err)
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    statistics.FilmStatistic
		wantErr bool
	}{

		{
			name: "success get statistic by request",
			args: args{
				ctx: context.Background(),
			},
			want: statistics.FilmStatistic{
				Request:    "/film/" + strings.ToLower(gofakeit.FirstName()),
				TimeDB:     time.Duration(int64(gofakeit.Number(int(time.Millisecond*80), int(time.Second)))),
				TimeRedis:  time.Duration(int64(gofakeit.Number(int(time.Millisecond*40), int(time.Millisecond*80)))),
				TimeMemory: time.Duration(int64(gofakeit.Number(int(time.Millisecond), int(time.Millisecond*40)))),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepository(time.Second*100, DB)

			if err := r.Create(tt.args.ctx, tt.want); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := r.GetByRequest(tt.args.ctx, tt.want.Request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got.ID = 0

			assert.DeepEqual(t, got, tt.want)
		})
	}
}

func TestRepository_Create(t *testing.T) {
	logger := log.New()

	DB, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	if err != nil {
		logger.Fatalf("test init db failed: %s", err)
	}
	type args struct {
		ctx  context.Context
		stat statistics.FilmStatistic
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success create statistic",
			args: args{
				ctx: context.Background(),
				stat: statistics.FilmStatistic{
					Request:    "/film/" + strings.ToLower(gofakeit.FirstName()),
					TimeDB:     time.Duration(int64(gofakeit.Number(int(time.Millisecond*80), int(time.Second)))),
					TimeRedis:  time.Duration(int64(gofakeit.Number(int(time.Millisecond*40), int(time.Millisecond*80)))),
					TimeMemory: time.Duration(int64(gofakeit.Number(int(time.Millisecond), int(time.Millisecond*40)))),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepository(time.Second*100, DB)

			if err := r.Create(tt.args.ctx, tt.args.stat); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := r.GetByRequest(tt.args.ctx, tt.args.stat.Request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got.ID = 0

			assert.DeepEqual(t, got, tt.args.stat)
		})
	}
}

func TestRepository_Update(t *testing.T) {
	logger := log.New()

	DB, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	if err != nil {
		logger.Fatalf("test init db failed: %s", err)
	}
	type args struct {
		ctx  context.Context
		stat statistics.FilmStatistic
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success update statistic",
			args: args{
				ctx: context.Background(),
				stat: statistics.FilmStatistic{
					Request:    "/film/" + strings.ToLower(gofakeit.FirstName()),
					TimeDB:     time.Duration(int64(gofakeit.Number(int(time.Millisecond*80), int(time.Second)))),
					TimeRedis:  time.Duration(int64(gofakeit.Number(int(time.Millisecond), int(time.Millisecond*40)))),
					TimeMemory: time.Duration(int64(gofakeit.Number(int(time.Millisecond*40), int(time.Millisecond*80)))),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepository(time.Second*100, DB)

			if err := r.Create(tt.args.ctx, tt.args.stat); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := r.GetByRequest(tt.args.ctx, tt.args.stat.Request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.args.stat.ID = got.ID

			tt.args.stat.TimeMemory = got.TimeRedis
			tt.args.stat.TimeRedis = got.TimeMemory

			if err := r.Update(tt.args.ctx, tt.args.stat); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetAll(t *testing.T) {
	logger := log.New()

	DB, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	if err != nil {
		logger.Fatalf("test init db failed: %s", err)
	}
	type args struct {
		ctx    context.Context
		limit  uint64
		offset uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success get all statistic",
			args: args{
				ctx:    context.Background(),
				limit:  10,
				offset: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repo.NewRepository(time.Second*100, DB)

			_, err := r.GetAll(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
