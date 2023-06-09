// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	services "server/services"

	mock "github.com/stretchr/testify/mock"
)

// PermissionServices is an autogenerated mock type for the PermissionServices type
type PermissionServices struct {
	mock.Mock
}

// List provides a mock function with given fields:
func (_m *PermissionServices) List() ([]services.PermissionResponse, error) {
	ret := _m.Called()

	var r0 []services.PermissionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]services.PermissionResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []services.PermissionResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.PermissionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPermissionServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewPermissionServices creates a new instance of PermissionServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPermissionServices(t mockConstructorTestingTNewPermissionServices) *PermissionServices {
	mock := &PermissionServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
