// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "github.com/lyn0904/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// DeviceDiagnosisClientInterface is an autogenerated mock type for the DeviceDiagnosisClientInterface type
type DeviceDiagnosisClientInterface struct {
	mock.Mock
}

type DeviceDiagnosisClientInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *DeviceDiagnosisClientInterface) EXPECT() *DeviceDiagnosisClientInterface_Expecter {
	return &DeviceDiagnosisClientInterface_Expecter{mock: &_m.Mock}
}

// RequestHeartbeat provides a mock function with given fields:
func (_m *DeviceDiagnosisClientInterface) RequestHeartbeat() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestHeartbeat")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.MsgCounterType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.MsgCounterType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceDiagnosisClientInterface_RequestHeartbeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestHeartbeat'
type DeviceDiagnosisClientInterface_RequestHeartbeat_Call struct {
	*mock.Call
}

// RequestHeartbeat is a helper method to define mock.On call
func (_e *DeviceDiagnosisClientInterface_Expecter) RequestHeartbeat() *DeviceDiagnosisClientInterface_RequestHeartbeat_Call {
	return &DeviceDiagnosisClientInterface_RequestHeartbeat_Call{Call: _e.mock.On("RequestHeartbeat")}
}

func (_c *DeviceDiagnosisClientInterface_RequestHeartbeat_Call) Run(run func()) *DeviceDiagnosisClientInterface_RequestHeartbeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceDiagnosisClientInterface_RequestHeartbeat_Call) Return(_a0 *model.MsgCounterType, _a1 error) *DeviceDiagnosisClientInterface_RequestHeartbeat_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceDiagnosisClientInterface_RequestHeartbeat_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *DeviceDiagnosisClientInterface_RequestHeartbeat_Call {
	_c.Call.Return(run)
	return _c
}

// RequestState provides a mock function with given fields:
func (_m *DeviceDiagnosisClientInterface) RequestState() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestState")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.MsgCounterType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.MsgCounterType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceDiagnosisClientInterface_RequestState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestState'
type DeviceDiagnosisClientInterface_RequestState_Call struct {
	*mock.Call
}

// RequestState is a helper method to define mock.On call
func (_e *DeviceDiagnosisClientInterface_Expecter) RequestState() *DeviceDiagnosisClientInterface_RequestState_Call {
	return &DeviceDiagnosisClientInterface_RequestState_Call{Call: _e.mock.On("RequestState")}
}

func (_c *DeviceDiagnosisClientInterface_RequestState_Call) Run(run func()) *DeviceDiagnosisClientInterface_RequestState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceDiagnosisClientInterface_RequestState_Call) Return(_a0 *model.MsgCounterType, _a1 error) *DeviceDiagnosisClientInterface_RequestState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceDiagnosisClientInterface_RequestState_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *DeviceDiagnosisClientInterface_RequestState_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeviceDiagnosisClientInterface creates a new instance of DeviceDiagnosisClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceDiagnosisClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceDiagnosisClientInterface {
	mock := &DeviceDiagnosisClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
