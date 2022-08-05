// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	film "films-api/internal/api/domain/film"
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