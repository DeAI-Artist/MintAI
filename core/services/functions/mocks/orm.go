// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	functions "github.com/DeAI-Artist/MintAI/core/services/functions"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/DeAI-Artist/MintAI/core/services/pg"

	time "time"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CreateRequest provides a mock function with given fields: request, qopts
func (_m *ORM) CreateRequest(request *functions.Request, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, request)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*functions.Request, ...pg.QOpt) error); ok {
		r0 = rf(request, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: requestID, qopts
func (_m *ORM) FindById(requestID functions.RequestID, qopts ...pg.QOpt) (*functions.Request, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *functions.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, ...pg.QOpt) (*functions.Request, error)); ok {
		return rf(requestID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(functions.RequestID, ...pg.QOpt) *functions.Request); ok {
		r0 = rf(requestID, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*functions.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(functions.RequestID, ...pg.QOpt) error); ok {
		r1 = rf(requestID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOldestEntriesByState provides a mock function with given fields: state, limit, qopts
func (_m *ORM) FindOldestEntriesByState(state functions.RequestState, limit uint32, qopts ...pg.QOpt) ([]functions.Request, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, state, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindOldestEntriesByState")
	}

	var r0 []functions.Request
	var r1 error
	if rf, ok := ret.Get(0).(func(functions.RequestState, uint32, ...pg.QOpt) ([]functions.Request, error)); ok {
		return rf(state, limit, qopts...)
	}
	if rf, ok := ret.Get(0).(func(functions.RequestState, uint32, ...pg.QOpt) []functions.Request); ok {
		r0 = rf(state, limit, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]functions.Request)
		}
	}

	if rf, ok := ret.Get(1).(func(functions.RequestState, uint32, ...pg.QOpt) error); ok {
		r1 = rf(state, limit, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PruneOldestRequests provides a mock function with given fields: maxRequestsInDB, batchSize, qopts
func (_m *ORM) PruneOldestRequests(maxRequestsInDB uint32, batchSize uint32, qopts ...pg.QOpt) (uint32, uint32, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, maxRequestsInDB, batchSize)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PruneOldestRequests")
	}

	var r0 uint32
	var r1 uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(uint32, uint32, ...pg.QOpt) (uint32, uint32, error)); ok {
		return rf(maxRequestsInDB, batchSize, qopts...)
	}
	if rf, ok := ret.Get(0).(func(uint32, uint32, ...pg.QOpt) uint32); ok {
		r0 = rf(maxRequestsInDB, batchSize, qopts...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(uint32, uint32, ...pg.QOpt) uint32); ok {
		r1 = rf(maxRequestsInDB, batchSize, qopts...)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	if rf, ok := ret.Get(2).(func(uint32, uint32, ...pg.QOpt) error); ok {
		r2 = rf(maxRequestsInDB, batchSize, qopts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetConfirmed provides a mock function with given fields: requestID, qopts
func (_m *ORM) SetConfirmed(requestID functions.RequestID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetConfirmed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, ...pg.QOpt) error); ok {
		r0 = rf(requestID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetError provides a mock function with given fields: requestID, errorType, computationError, readyAt, readyForProcessing, qopts
func (_m *ORM) SetError(requestID functions.RequestID, errorType functions.ErrType, computationError []byte, readyAt time.Time, readyForProcessing bool, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID, errorType, computationError, readyAt, readyForProcessing)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, functions.ErrType, []byte, time.Time, bool, ...pg.QOpt) error); ok {
		r0 = rf(requestID, errorType, computationError, readyAt, readyForProcessing, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetFinalized provides a mock function with given fields: requestID, reportedResult, reportedError, qopts
func (_m *ORM) SetFinalized(requestID functions.RequestID, reportedResult []byte, reportedError []byte, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID, reportedResult, reportedError)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetFinalized")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, []byte, []byte, ...pg.QOpt) error); ok {
		r0 = rf(requestID, reportedResult, reportedError, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetResult provides a mock function with given fields: requestID, computationResult, readyAt, qopts
func (_m *ORM) SetResult(requestID functions.RequestID, computationResult []byte, readyAt time.Time, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, requestID, computationResult, readyAt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetResult")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(functions.RequestID, []byte, time.Time, ...pg.QOpt) error); ok {
		r0 = rf(requestID, computationResult, readyAt, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TimeoutExpiredResults provides a mock function with given fields: cutoff, limit, qopts
func (_m *ORM) TimeoutExpiredResults(cutoff time.Time, limit uint32, qopts ...pg.QOpt) ([]functions.RequestID, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, cutoff, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TimeoutExpiredResults")
	}

	var r0 []functions.RequestID
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, uint32, ...pg.QOpt) ([]functions.RequestID, error)); ok {
		return rf(cutoff, limit, qopts...)
	}
	if rf, ok := ret.Get(0).(func(time.Time, uint32, ...pg.QOpt) []functions.RequestID); ok {
		r0 = rf(cutoff, limit, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]functions.RequestID)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, uint32, ...pg.QOpt) error); ok {
		r1 = rf(cutoff, limit, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}