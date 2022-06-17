// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	seller "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/seller"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateSeller provides a mock function with given fields: cid, companyName, address, telephone
func (_m *Service) CreateSeller(cid int64, companyName string, address string, telephone string) (seller.Seller, error) {
	ret := _m.Called(cid, companyName, address, telephone)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(int64, string, string, string) seller.Seller); ok {
		r0 = rf(cid, companyName, address, telephone)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string, string, string) error); ok {
		r1 = rf(cid, companyName, address, telephone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSeller provides a mock function with given fields: id
func (_m *Service) DeleteSeller(id int64) error {
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
func (_m *Service) GetAll() ([]seller.Seller, error) {
	ret := _m.Called()

	var r0 []seller.Seller
	if rf, ok := ret.Get(0).(func() []seller.Seller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]seller.Seller)
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

// GetById provides a mock function with given fields: id
func (_m *Service) GetById(id int64) (seller.Seller, error) {
	ret := _m.Called(id)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(int64) seller.Seller); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSellerAddresAndTel provides a mock function with given fields: id, address, telephone
func (_m *Service) UpdateSellerAddresAndTel(id int64, address string, telephone string) (seller.Seller, error) {
	ret := _m.Called(id, address, telephone)

	var r0 seller.Seller
	if rf, ok := ret.Get(0).(func(int64, string, string) seller.Seller); ok {
		r0 = rf(id, address, telephone)
	} else {
		r0 = ret.Get(0).(seller.Seller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string, string) error); ok {
		r1 = rf(id, address, telephone)
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
