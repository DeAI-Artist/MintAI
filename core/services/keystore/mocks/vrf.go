// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	vrfkey "github.com/DeAI-Artist/MintAI/core/services/keystore/keys/vrfkey"
)

// VRF is an autogenerated mock type for the VRF type
type VRF struct {
	mock.Mock
}

// Add provides a mock function with given fields: key
func (_m *VRF) Add(key vrfkey.KeyV2) error {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(vrfkey.KeyV2) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields:
func (_m *VRF) Create() (vrfkey.KeyV2, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func() (vrfkey.KeyV2, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() vrfkey.KeyV2); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(vrfkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *VRF) Delete(id string) (vrfkey.KeyV2, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (vrfkey.KeyV2, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) vrfkey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(vrfkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Export provides a mock function with given fields: id, password
func (_m *VRF) Export(id string, password string) ([]byte, error) {
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

// GenerateProof provides a mock function with given fields: id, seed
func (_m *VRF) GenerateProof(id string, seed *big.Int) (vrfkey.Proof, error) {
	ret := _m.Called(id, seed)

	if len(ret) == 0 {
		panic("no return value specified for GenerateProof")
	}

	var r0 vrfkey.Proof
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *big.Int) (vrfkey.Proof, error)); ok {
		return rf(id, seed)
	}
	if rf, ok := ret.Get(0).(func(string, *big.Int) vrfkey.Proof); ok {
		r0 = rf(id, seed)
	} else {
		r0 = ret.Get(0).(vrfkey.Proof)
	}

	if rf, ok := ret.Get(1).(func(string, *big.Int) error); ok {
		r1 = rf(id, seed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *VRF) Get(id string) (vrfkey.KeyV2, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (vrfkey.KeyV2, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) vrfkey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(vrfkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *VRF) GetAll() ([]vrfkey.KeyV2, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]vrfkey.KeyV2, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []vrfkey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vrfkey.KeyV2)
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
func (_m *VRF) Import(keyJSON []byte, password string) (vrfkey.KeyV2, error) {
	ret := _m.Called(keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string) (vrfkey.KeyV2, error)); ok {
		return rf(keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func([]byte, string) vrfkey.KeyV2); ok {
		r0 = rf(keyJSON, password)
	} else {
		r0 = ret.Get(0).(vrfkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewVRF creates a new instance of VRF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVRF(t interface {
	mock.TestingT
	Cleanup(func())
}) *VRF {
	mock := &VRF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
