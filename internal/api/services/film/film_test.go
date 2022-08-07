package film

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"

	"films-api/internal/api/domain/film"
	"films-api/internal/api/repository"
	"films-api/internal/api/services"
	"films-api/pkg/errs"
	"films-api/pkg/log"
)

func TestService_GetByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := log.NewTest(t, zaptest.Level(zapcore.FatalLevel))

	filmPostgres := repository.NewMockFilmPostgres(ctrl)
	filmRedisCache := repository.NewMockFilmCache(ctrl)
	filmLocalCache := repository.NewMockFilmCache(ctrl)
	statisticsService := services.NewMockStatistics(ctrl)

	service := NewService(
		filmPostgres,
		filmRedisCache,
		filmLocalCache,
		statisticsService,
		logger,
	)

	type mocker struct {
		filmPostgres      *repository.MockFilmPostgres
		filmRedisCache    *repository.MockFilmCache
		filmLocalCache    *repository.MockFilmCache
		statisticsService *services.MockStatistics
	}

	m := mocker{
		filmPostgres:      filmPostgres,
		filmRedisCache:    filmRedisCache,
		filmLocalCache:    filmLocalCache,
		statisticsService: statisticsService,
	}

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name         string
		args         args
		mockFn       func(ctx context.Context, m mocker)
		wantFilmList film.FilmList
		wantErr      bool
	}{
		{
			name: "Success get film from postgres",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmRedisCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmPostgres.EXPECT().GetByName(ctx, "rocky").Return([]film.Film{{Name: "rocky"}}, nil),
					m.filmLocalCache.EXPECT().SetByName(ctx, "rocky", []film.Film{{Name: "rocky"}}).Return(nil),
					m.filmRedisCache.EXPECT().SetByName(ctx, "rocky", []film.Film{{Name: "rocky"}}).Return(nil),
					m.statisticsService.EXPECT().Update(gomock.Any()),
				)

			},
			wantFilmList: []film.Film{{Name: "rocky"}},
			wantErr:      false,
		},
		{
			name: "Fail update redis cache",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmRedisCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmPostgres.EXPECT().GetByName(ctx, "rocky").Return([]film.Film{{Name: "rocky"}}, nil),
					m.filmLocalCache.EXPECT().SetByName(ctx, "rocky", []film.Film{{Name: "rocky"}}).Return(nil),
					m.filmRedisCache.EXPECT().SetByName(ctx, "rocky", []film.Film{{Name: "rocky"}}).Return(errs.Internal{}),
				)

			},
			wantFilmList: nil,
			wantErr:      true,
		},
		{
			name: "Fail update local cache",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmRedisCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmPostgres.EXPECT().GetByName(ctx, "rocky").Return([]film.Film{{Name: "rocky"}}, nil),
					m.filmLocalCache.EXPECT().SetByName(ctx, "rocky", []film.Film{{Name: "rocky"}}).Return(errs.Internal{}),
				)

			},
			wantFilmList: nil,
			wantErr:      true,
		},
		{
			name: "Fail get by name from postgres",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmRedisCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmPostgres.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.Internal{}),
				)

			},
			wantFilmList: nil,
			wantErr:      true,
		},
		{
			name: "Fail not found film in postgres",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmRedisCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmPostgres.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
				)

			},
			wantFilmList: nil,
			wantErr:      true,
		},
		{
			name: "Fail get by name from redis",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmRedisCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.Internal{}),
				)

			},
			wantFilmList: nil,
			wantErr:      true,
		},
		{
			name: "Success get by name from redis",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmRedisCache.EXPECT().GetByName(ctx, "rocky").Return([]film.Film{{Name: "rocky"}}, nil),
					m.filmLocalCache.EXPECT().SetByName(ctx, "rocky", []film.Film{{Name: "rocky"}}).Return(nil),
					m.statisticsService.EXPECT().Update(gomock.Any()),
				)

			},
			wantFilmList: []film.Film{{Name: "rocky"}},
			wantErr:      false,
		},
		{
			name: "Fail update local cache",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.NotFound{}),
					m.filmRedisCache.EXPECT().GetByName(ctx, "rocky").Return([]film.Film{{Name: "rocky"}}, nil),
					m.filmLocalCache.EXPECT().SetByName(ctx, "rocky", []film.Film{{Name: "rocky"}}).Return(errs.Internal{}),
				)

			},
			wantFilmList: nil,
			wantErr:      true,
		},
		{
			name: "Fail get by name from local cache",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return(nil, errs.Internal{}),
				)

			},
			wantFilmList: nil,
			wantErr:      true,
		},
		{
			name: "Success get by name from local cache",
			args: args{
				ctx:  context.Background(),
				name: "rocky",
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.filmLocalCache.EXPECT().GetByName(ctx, "rocky").Return([]film.Film{{Name: "rocky"}}, nil),
					m.statisticsService.EXPECT().Update(gomock.Any()),
				)

			},
			wantFilmList: []film.Film{{Name: "rocky"}},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args.ctx, m)

			gotFilmList, err := service.GetByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotFilmList, tt.wantFilmList) {
				t.Errorf("Service.GetByName() = %v, want %v", gotFilmList, tt.wantFilmList)
			}
		})
	}
}
