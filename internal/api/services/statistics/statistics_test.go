package statistics

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"

	"films-api/internal/api/domain/statistics"
	"films-api/internal/api/repository"
	"films-api/pkg/errs"
	"films-api/pkg/log"
)

func TestService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := log.NewTest(t, zaptest.Level(zapcore.FatalLevel))

	statisticsRepo := repository.NewMockStatistics(ctrl)

	service := NewService(
		context.Background(),
		statisticsRepo,
		logger,
	)

	type mocker struct {
		statisticsRepo *repository.MockStatistics
	}

	m := mocker{
		statisticsRepo: statisticsRepo,
	}

	type args struct {
		ctx    context.Context
		limit  uint64
		offset uint64
	}
	tests := []struct {
		name    string
		args    args
		mockFn  func(ctx context.Context, m mocker)
		want    statistics.FilmStatisticList
		wantErr bool
	}{
		{
			name: "success get statistics",
			args: args{
				ctx:    context.Background(),
				limit:  10,
				offset: 0,
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.statisticsRepo.EXPECT().GetAll(ctx, uint64(10), uint64(0)).Return(statistics.FilmStatisticList{}, nil),
				)
			},
			want:    statistics.FilmStatisticList{},
			wantErr: false,
		},
		{
			name: "Fail get statistics",
			args: args{
				ctx:    context.Background(),
				limit:  10,
				offset: 0,
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.statisticsRepo.EXPECT().GetAll(ctx, uint64(10), uint64(0)).Return(statistics.FilmStatisticList{}, errs.Internal{}),
				)
			},
			want:    statistics.FilmStatisticList{},
			wantErr: true,
		},
		{
			name: "Fail not found statistics",
			args: args{
				ctx:    context.Background(),
				limit:  10,
				offset: 0,
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.statisticsRepo.EXPECT().GetAll(ctx, uint64(10), uint64(0)).Return(statistics.FilmStatisticList{}, errs.NotFound{}),
				)
			},
			want:    statistics.FilmStatisticList{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args.ctx, m)

			got, err := service.GetAll(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := log.NewTest(t, zaptest.Level(zapcore.FatalLevel))

	statisticsRepo := repository.NewMockStatistics(ctrl)

	ctx := context.Background()

	service := NewService(
		ctx,
		statisticsRepo,
		logger,
	)

	type mocker struct {
		statisticsRepo *repository.MockStatistics
	}

	m := mocker{
		statisticsRepo: statisticsRepo,
	}

	type args struct {
		stat statistics.FilmStatistic
	}
	tests := []struct {
		name   string
		args   args
		mockFn func(ctx context.Context, m mocker)
	}{
		{
			name: "success update statistics",
			args: args{
				stat: statistics.FilmStatistic{
					Request: "request",
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.statisticsRepo.EXPECT().GetByRequest(ctx, "request").Return(statistics.FilmStatistic{}, nil),
					m.statisticsRepo.EXPECT().Update(ctx, statistics.FilmStatistic{}).Return(nil),
				)
			},
		},
		{
			name: "fail update statistics",
			args: args{
				stat: statistics.FilmStatistic{
					Request: "request",
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.statisticsRepo.EXPECT().GetByRequest(ctx, "request").Return(statistics.FilmStatistic{}, nil),
					m.statisticsRepo.EXPECT().Update(ctx, statistics.FilmStatistic{}).Return(errs.Internal{}),
				)
			},
		},
		{
			name: "fail get statistics",
			args: args{
				stat: statistics.FilmStatistic{
					Request: "request",
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.statisticsRepo.EXPECT().GetByRequest(ctx, "request").Return(statistics.FilmStatistic{}, errs.Internal{}),
				)
			},
		},
		{
			name: "success create new row",
			args: args{
				stat: statistics.FilmStatistic{
					Request: "request",
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.statisticsRepo.EXPECT().GetByRequest(ctx, "request").Return(statistics.FilmStatistic{}, errs.NotFound{}),
					m.statisticsRepo.EXPECT().Create(ctx, statistics.FilmStatistic{Request: "request"}).Return(nil),
				)
			},
		},
		{
			name: "fail create new row",
			args: args{
				stat: statistics.FilmStatistic{
					Request: "request",
				},
			},
			mockFn: func(ctx context.Context, m mocker) {
				gomock.InOrder(
					m.statisticsRepo.EXPECT().GetByRequest(ctx, "request").Return(statistics.FilmStatistic{}, errs.NotFound{}),
					m.statisticsRepo.EXPECT().Create(ctx, statistics.FilmStatistic{Request: "request"}).Return(errs.Internal{}),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(ctx, m)

			service.Update(tt.args.stat)
			time.Sleep(time.Millisecond)
		})
	}
}
