// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	buyer "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/buyer"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Service) GetAll() ([]buyer.Buyer, error) {
	ret := _m.Called()

	var r0 []buyer.Buyer
	if rf, ok := ret.Get(0).(func() []buyer.Buyer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]buyer.Buyer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetId provides a mock function with given fields: id
func (_m *Service) GetId(id int64) (*buyer.Buyer, error) {
	ret := _m.Called(id)

	var r0 *buyer.Buyer
	if rf, ok := ret.Get(0).(func(int64) *buyer.Buyer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*buyer.Buyer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: cardNumberId, firstName, lastName
func (_m *Service) Store(cardNumberId int64, firstName string, lastName string) (buyer.Buyer, error) {
	ret := _m.Called(cardNumberId, firstName, lastName)

	var r0 buyer.Buyer
	if rf, ok := ret.Get(0).(func(int64, string, string) buyer.Buyer); ok {
		r0 = rf(cardNumberId, firstName, lastName)
	} else {
		r0 = ret.Get(0).(buyer.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string, string) error); ok {
		r1 = rf(cardNumberId, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, cardNumberId, lastName
func (_m *Service) Update(id int64, cardNumberId int64, lastName string) (buyer.Buyer, error) {
	ret := _m.Called(id, cardNumberId, lastName)

	var r0 buyer.Buyer
	if rf, ok := ret.Get(0).(func(int64, int64, string) buyer.Buyer); ok {
		r0 = rf(id, cardNumberId, lastName)
	} else {
		r0 = ret.Get(0).(buyer.Buyer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, string) error); ok {
		r1 = rf(id, cardNumberId, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
