// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/employees/domain"
)

// EmployeeService is an autogenerated mock type for the EmployeeService type
type EmployeeService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, cardNumberId, firstName, lastName, warehouseId
func (_m *EmployeeService) Create(ctx context.Context, cardNumberId string, firstName string, lastName string, warehouseId int64) (domain.Employee, error) {
	ret := _m.Called(ctx, cardNumberId, firstName, lastName, warehouseId)

	var r0 domain.Employee
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, int64) domain.Employee); ok {
		r0 = rf(ctx, cardNumberId, firstName, lastName, warehouseId)
	} else {
		r0 = ret.Get(0).(domain.Employee)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, int64) error); ok {
		r1 = rf(ctx, cardNumberId, firstName, lastName, warehouseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *EmployeeService) Delete(ctx context.Context, id int64) error {
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
func (_m *EmployeeService) GetAll(ctx context.Context) ([]domain.Employee, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Employee
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Employee); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Employee)
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

// GetAllReportInboundOrders provides a mock function with given fields: ctx
func (_m *EmployeeService) GetAllReportInboundOrders(ctx context.Context) ([]domain.EmployeeInboundOrdersReport, error) {
	ret := _m.Called(ctx)

	var r0 []domain.EmployeeInboundOrdersReport
	if rf, ok := ret.Get(0).(func(context.Context) []domain.EmployeeInboundOrdersReport); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.EmployeeInboundOrdersReport)
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
func (_m *EmployeeService) GetById(ctx context.Context, id int64) (*domain.Employee, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Employee
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Employee); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Employee)
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

// GetReportInboundOrdersById provides a mock function with given fields: ctx, employeeID
func (_m *EmployeeService) GetReportInboundOrdersById(ctx context.Context, employeeID int64) (domain.EmployeeInboundOrdersReport, error) {
	ret := _m.Called(ctx, employeeID)

	var r0 domain.EmployeeInboundOrdersReport
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.EmployeeInboundOrdersReport); ok {
		r0 = rf(ctx, employeeID)
	} else {
		r0 = ret.Get(0).(domain.EmployeeInboundOrdersReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, employeeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFullname provides a mock function with given fields: ctx, id, firstName, lastName
func (_m *EmployeeService) UpdateFullname(ctx context.Context, id int64, firstName string, lastName string) (*domain.Employee, error) {
	ret := _m.Called(ctx, id, firstName, lastName)

	var r0 *domain.Employee
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) *domain.Employee); ok {
		r0 = rf(ctx, id, firstName, lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Employee)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string) error); ok {
		r1 = rf(ctx, id, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewEmployeeServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmployeeService creates a new instance of EmployeeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmployeeService(t NewEmployeeServiceT) *EmployeeService {
	mock := &EmployeeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
