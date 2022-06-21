// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	employees "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/employees"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: cardNumberId, firstName, lastName, warehouseId
func (_m *Repository) Create(cardNumberId string, firstName string, lastName string, warehouseId int64) (employees.Employee, error) {
	ret := _m.Called(cardNumberId, firstName, lastName, warehouseId)

	var r0 employees.Employee
	if rf, ok := ret.Get(0).(func(string, string, string, int64) employees.Employee); ok {
		r0 = rf(cardNumberId, firstName, lastName, warehouseId)
	} else {
		r0 = ret.Get(0).(employees.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int64) error); ok {
		r1 = rf(cardNumberId, firstName, lastName, warehouseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int64) error {
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
func (_m *Repository) GetAll() ([]employees.Employee, error) {
	ret := _m.Called()

	var r0 []employees.Employee
	if rf, ok := ret.Get(0).(func() []employees.Employee); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]employees.Employee)
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
func (_m *Repository) GetById(id int64) (*employees.Employee, error) {
	ret := _m.Called(id)

	var r0 *employees.Employee
	if rf, ok := ret.Get(0).(func(int64) *employees.Employee); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*employees.Employee)
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

// UpdateFullname provides a mock function with given fields: id, firstName, lastName
func (_m *Repository) UpdateFullname(id int64, firstName string, lastName string) (*employees.Employee, error) {
	ret := _m.Called(id, firstName, lastName)

	var r0 *employees.Employee
	if rf, ok := ret.Get(0).(func(int64, string, string) *employees.Employee); ok {
		r0 = rf(id, firstName, lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*employees.Employee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string, string) error); ok {
		r1 = rf(id, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t NewRepositoryT) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}