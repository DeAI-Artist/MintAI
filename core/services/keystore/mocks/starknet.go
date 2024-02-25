// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	starkkey "github.com/DeAI-Artist/MintAI/core/services/keystore/keys/starkkey"
	mock "github.com/stretchr/testify/mock"
)

// StarkNet is an autogenerated mock type for the StarkNet type
type StarkNet struct {
	mock.Mock
}

// Add provides a mock function with given fields: key
func (_m *StarkNet) Add(key starkkey.Key) error {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(starkkey.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields:
func (_m *StarkNet) Create() (starkkey.Key, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 starkkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func() (starkkey.Key, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() starkkey.Key); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(starkkey.Key)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *StarkNet) Delete(id string) (starkkey.Key, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 starkkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (starkkey.Key, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) starkkey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(starkkey.Key)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields:
func (_m *StarkNet) EnsureKey() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EnsureKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *StarkNet) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *StarkNet) Get(id string) (starkkey.Key, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 starkkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (starkkey.Key, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) starkkey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(starkkey.Key)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *StarkNet) GetAll() ([]starkkey.Key, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []starkkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]starkkey.Key, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []starkkey.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]starkkey.Key)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: keyJSON, password
func (_m *StarkNet) Import(keyJSON []byte, password string) (starkkey.Key, error) {
	ret := _m.Called(keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 starkkey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string) (starkkey.Key, error)); ok {
		return rf(keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func([]byte, string) starkkey.Key); ok {
		r0 = rf(keyJSON, password)
	} else {
		r0 = ret.Get(0).(starkkey.Key)
	}

	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStarkNet creates a new instance of StarkNet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStarkNet(t interface {
	mock.TestingT
	Cleanup(func())
}) *StarkNet {
	mock := &StarkNet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
