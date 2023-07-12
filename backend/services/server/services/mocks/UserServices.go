// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entities "server/entities"
	services "server/services"

	mock "github.com/stretchr/testify/mock"
)

// UserServices is an autogenerated mock type for the UserServices type
type UserServices struct {
	mock.Mock
}

// AccessAdminPage provides a mock function with given fields: role
func (_m *UserServices) AccessAdminPage(role string) error {
	ret := _m.Called(role)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeRole provides a mock function with given fields: role, id, newRole
func (_m *UserServices) ChangeRole(role string, id uint, newRole string) error {
	ret := _m.Called(role, id, newRole)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint, string) error); ok {
		r0 = rf(role, id, newRole)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields:
func (_m *UserServices) Count() (int, error) {
	ret := _m.Called()

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: registerUserInput
func (_m *UserServices) CreateUser(registerUserInput *services.RegisterUserInput) (*entities.User, error) {
	ret := _m.Called(registerUserInput)

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*services.RegisterUserInput) (*entities.User, error)); ok {
		return rf(registerUserInput)
	}
	if rf, ok := ret.Get(0).(func(*services.RegisterUserInput) *entities.User); ok {
		r0 = rf(registerUserInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*services.RegisterUserInput) error); ok {
		r1 = rf(registerUserInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: role, id
func (_m *UserServices) Delete(role string, id uint) error {
	ret := _m.Called(role, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(role, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: username
func (_m *UserServices) GetUser(username string) (*entities.User, error) {
	ret := _m.Called(username)

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GoogleOAuth provides a mock function with given fields: googleUser
func (_m *UserServices) GoogleOAuth(googleUser *services.GoogleUserResult) (string, string, error) {
	ret := _m.Called(googleUser)

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(*services.GoogleUserResult) (string, string, error)); ok {
		return rf(googleUser)
	}
	if rf, ok := ret.Get(0).(func(*services.GoogleUserResult) string); ok {
		r0 = rf(googleUser)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*services.GoogleUserResult) string); ok {
		r1 = rf(googleUser)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(*services.GoogleUserResult) error); ok {
		r2 = rf(googleUser)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// List provides a mock function with given fields: page, pageSize
func (_m *UserServices) List(page int, pageSize int) ([]services.UserResponse, error) {
	ret := _m.Called(page, pageSize)

	var r0 []services.UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]services.UserResponse, error)); ok {
		return rf(page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(int, int) []services.UserResponse); ok {
		r0 = rf(page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]services.UserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginWithEmail provides a mock function with given fields: inputUser
func (_m *UserServices) LoginWithEmail(inputUser *services.LoginUserInput) (string, string, error) {
	ret := _m.Called(inputUser)

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(*services.LoginUserInput) (string, string, error)); ok {
		return rf(inputUser)
	}
	if rf, ok := ret.Get(0).(func(*services.LoginUserInput) string); ok {
		r0 = rf(inputUser)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*services.LoginUserInput) string); ok {
		r1 = rf(inputUser)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(*services.LoginUserInput) error); ok {
		r2 = rf(inputUser)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LoginWithUsername provides a mock function with given fields: inputUser
func (_m *UserServices) LoginWithUsername(inputUser *services.LoginUserInput) (string, string, error) {
	ret := _m.Called(inputUser)

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(*services.LoginUserInput) (string, string, error)); ok {
		return rf(inputUser)
	}
	if rf, ok := ret.Get(0).(func(*services.LoginUserInput) string); ok {
		r0 = rf(inputUser)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*services.LoginUserInput) string); ok {
		r1 = rf(inputUser)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(*services.LoginUserInput) error); ok {
		r2 = rf(inputUser)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UserUpgrateRole provides a mock function with given fields: role, username
func (_m *UserServices) UserUpgrateRole(role string, username string) (string, string, error) {
	ret := _m.Called(role, username)

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (string, string, error)); ok {
		return rf(role, username)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(role, username)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(role, username)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(role, username)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewUserServices interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserServices creates a new instance of UserServices. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserServices(t mockConstructorTestingTNewUserServices) *UserServices {
	mock := &UserServices{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
