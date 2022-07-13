// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/locality/domain"
)

// LocalityService is an autogenerated mock type for the LocalityService type
type LocalityService struct {
	mock.Mock
}

// CreateLocality provides a mock function with given fields: ctx, locality
func (_m *LocalityService) CreateLocality(ctx context.Context, locality *domain.LocalityModel) (*domain.LocalityModel, error) {
	ret := _m.Called(ctx, locality)

	var r0 *domain.LocalityModel
	if rf, ok := ret.Get(0).(func(context.Context, *domain.LocalityModel) *domain.LocalityModel); ok {
		r0 = rf(ctx, locality)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.LocalityModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.LocalityModel) error); ok {
		r1 = rf(ctx, locality)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReportSeller provides a mock function with given fields: ctx
func (_m *LocalityService) GetAllReportSeller(ctx context.Context) (*[]domain.ReportSeller, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.ReportSeller
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.ReportSeller); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.ReportSeller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIdReportSeller provides a mock function with given fields: ctx, locality_id
func (_m *LocalityService) GetByIdReportSeller(ctx context.Context, locality_id int64) (*[]domain.ReportSeller, error) {
	ret := _m.Called(ctx, locality_id)

	var r0 *[]domain.ReportSeller
	if rf, ok := ret.Get(0).(func(context.Context, int64) *[]domain.ReportSeller); ok {
		r0 = rf(ctx, locality_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.ReportSeller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, locality_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportCarrie provides a mock function with given fields: ctx, locality_id
func (_m *LocalityService) ReportCarrie(ctx context.Context, locality_id int64) (*[]domain.ReportCarrie, error) {
	ret := _m.Called(ctx, locality_id)

	var r0 *[]domain.ReportCarrie
	if rf, ok := ret.Get(0).(func(context.Context, int64) *[]domain.ReportCarrie); ok {
		r0 = rf(ctx, locality_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.ReportCarrie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, locality_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLocalityService interface {
	mock.TestingT
	Cleanup(func())
}

// NewLocalityService creates a new instance of LocalityService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocalityService(t mockConstructorTestingTNewLocalityService) *LocalityService {
	mock := &LocalityService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}