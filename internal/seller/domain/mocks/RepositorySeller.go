// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/seller/domain"
)

// RepositorySeller is an autogenerated mock type for the RepositorySeller type
type RepositorySeller struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, seller
func (_m *RepositorySeller) Create(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	ret := _m.Called(ctx, seller)

	var r0 *domain.Seller
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Seller) *domain.Seller); ok {
		r0 = rf(ctx, seller)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Seller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Seller) error); ok {
		r1 = rf(ctx, seller)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *RepositorySeller) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *RepositorySeller) GetAll(ctx context.Context) (*[]domain.Seller, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.Seller
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.Seller); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Seller)
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

// GetById provides a mock function with given fields: ctx, id
func (_m *RepositorySeller) GetById(ctx context.Context, id int64) (*domain.Seller, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Seller
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Seller); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Seller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, seller
func (_m *RepositorySeller) Update(ctx context.Context, seller *domain.Seller) (*domain.Seller, error) {
	ret := _m.Called(ctx, seller)

	var r0 *domain.Seller
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Seller) *domain.Seller); ok {
		r0 = rf(ctx, seller)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Seller)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Seller) error); ok {
		r1 = rf(ctx, seller)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositorySeller interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositorySeller creates a new instance of RepositorySeller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositorySeller(t mockConstructorTestingTNewRepositorySeller) *RepositorySeller {
	mock := &RepositorySeller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
