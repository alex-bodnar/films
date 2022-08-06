// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	film "films-api/internal/api/domain/film"
	statistics "films-api/internal/api/domain/statistics"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFilmPostgres is a mock of FilmPostgres interface.
type MockFilmPostgres struct {
	ctrl     *gomock.Controller
	recorder *MockFilmPostgresMockRecorder
}

// MockFilmPostgresMockRecorder is the mock recorder for MockFilmPostgres.
type MockFilmPostgresMockRecorder struct {
	mock *MockFilmPostgres
}

// NewMockFilmPostgres creates a new mock instance.
func NewMockFilmPostgres(ctrl *gomock.Controller) *MockFilmPostgres {
	mock := &MockFilmPostgres{ctrl: ctrl}
	mock.recorder = &MockFilmPostgresMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilmPostgres) EXPECT() *MockFilmPostgresMockRecorder {
	return m.recorder
}

// GetByName mocks base method.
func (m *MockFilmPostgres) GetByName(ctx context.Context, name string) (film.FilmList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(film.FilmList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockFilmPostgresMockRecorder) GetByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockFilmPostgres)(nil).GetByName), ctx, name)
}

// MockStatistics is a mock of Statistics interface.
type MockStatistics struct {
	ctrl     *gomock.Controller
	recorder *MockStatisticsMockRecorder
}

// MockStatisticsMockRecorder is the mock recorder for MockStatistics.
type MockStatisticsMockRecorder struct {
	mock *MockStatistics
}

// NewMockStatistics creates a new mock instance.
func NewMockStatistics(ctrl *gomock.Controller) *MockStatistics {
	mock := &MockStatistics{ctrl: ctrl}
	mock.recorder = &MockStatisticsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatistics) EXPECT() *MockStatisticsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStatistics) Create(ctx context.Context, stat statistics.FilmStatistic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, stat)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStatisticsMockRecorder) Create(ctx, stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStatistics)(nil).Create), ctx, stat)
}

// GetAll mocks base method.
func (m *MockStatistics) GetAll(ctx context.Context, limit, offset uint64) (statistics.FilmStatisticList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, limit, offset)
	ret0, _ := ret[0].(statistics.FilmStatisticList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockStatisticsMockRecorder) GetAll(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStatistics)(nil).GetAll), ctx, limit, offset)
}

// GetByRequest mocks base method.
func (m *MockStatistics) GetByRequest(ctx context.Context, req string) (statistics.FilmStatistic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRequest", ctx, req)
	ret0, _ := ret[0].(statistics.FilmStatistic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByRequest indicates an expected call of GetByRequest.
func (mr *MockStatisticsMockRecorder) GetByRequest(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRequest", reflect.TypeOf((*MockStatistics)(nil).GetByRequest), ctx, req)
}

// Update mocks base method.
func (m *MockStatistics) Update(ctx context.Context, stat statistics.FilmStatistic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, stat)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStatisticsMockRecorder) Update(ctx, stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStatistics)(nil).Update), ctx, stat)
}

// MockFilmCache is a mock of FilmCache interface.
type MockFilmCache struct {
	ctrl     *gomock.Controller
	recorder *MockFilmCacheMockRecorder
}

// MockFilmCacheMockRecorder is the mock recorder for MockFilmCache.
type MockFilmCacheMockRecorder struct {
	mock *MockFilmCache
}

// NewMockFilmCache creates a new mock instance.
func NewMockFilmCache(ctrl *gomock.Controller) *MockFilmCache {
	mock := &MockFilmCache{ctrl: ctrl}
	mock.recorder = &MockFilmCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilmCache) EXPECT() *MockFilmCacheMockRecorder {
	return m.recorder
}

// GetByName mocks base method.
func (m *MockFilmCache) GetByName(ctx context.Context, name string) (film.FilmList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(film.FilmList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockFilmCacheMockRecorder) GetByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockFilmCache)(nil).GetByName), ctx, name)
}

// SetByName mocks base method.
func (m *MockFilmCache) SetByName(ctx context.Context, name string, data film.FilmList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetByName", ctx, name, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetByName indicates an expected call of SetByName.
func (mr *MockFilmCacheMockRecorder) SetByName(ctx, name, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetByName", reflect.TypeOf((*MockFilmCache)(nil).SetByName), ctx, name, data)
}
