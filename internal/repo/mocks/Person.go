// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/punkestu/ecommerce-go/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Person is an autogenerated mock type for the Person type
type Person struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Person) Create(_a0 entity.Person) (int32, error) {
	ret := _m.Called(_a0)

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Person) (int32, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(entity.Person) int32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(entity.Person) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: _a0
func (_m *Person) GetByEmail(_a0 string) (*entity.Person, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Person, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Person); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *Person) GetByID(_a0 int32) (*entity.Person, error) {
	ret := _m.Called(_a0)

	var r0 *entity.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(int32) (*entity.Person, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int32) *entity.Person); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPerson creates a new instance of Person. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPerson(t interface {
	mock.TestingT
	Cleanup(func())
}) *Person {
	mock := &Person{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
