// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "github.com/lyn0904/spine-go/model"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// DeviceDiagnosisCommonInterface is an autogenerated mock type for the DeviceDiagnosisCommonInterface type
type DeviceDiagnosisCommonInterface struct {
	mock.Mock
}

type DeviceDiagnosisCommonInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *DeviceDiagnosisCommonInterface) EXPECT() *DeviceDiagnosisCommonInterface_Expecter {
	return &DeviceDiagnosisCommonInterface_Expecter{mock: &_m.Mock}
}

// GetState provides a mock function with given fields:
func (_m *DeviceDiagnosisCommonInterface) GetState() (*model.DeviceDiagnosisStateDataType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 *model.DeviceDiagnosisStateDataType
	var r1 error
	if rf, ok := ret.Get(0).(func() (*model.DeviceDiagnosisStateDataType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *model.DeviceDiagnosisStateDataType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceDiagnosisStateDataType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceDiagnosisCommonInterface_GetState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetState'
type DeviceDiagnosisCommonInterface_GetState_Call struct {
	*mock.Call
}

// GetState is a helper method to define mock.On call
func (_e *DeviceDiagnosisCommonInterface_Expecter) GetState() *DeviceDiagnosisCommonInterface_GetState_Call {
	return &DeviceDiagnosisCommonInterface_GetState_Call{Call: _e.mock.On("GetState")}
}

func (_c *DeviceDiagnosisCommonInterface_GetState_Call) Run(run func()) *DeviceDiagnosisCommonInterface_GetState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceDiagnosisCommonInterface_GetState_Call) Return(_a0 *model.DeviceDiagnosisStateDataType, _a1 error) *DeviceDiagnosisCommonInterface_GetState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceDiagnosisCommonInterface_GetState_Call) RunAndReturn(run func() (*model.DeviceDiagnosisStateDataType, error)) *DeviceDiagnosisCommonInterface_GetState_Call {
	_c.Call.Return(run)
	return _c
}

// IsHeartbeatWithinDuration provides a mock function with given fields: duration
func (_m *DeviceDiagnosisCommonInterface) IsHeartbeatWithinDuration(duration time.Duration) bool {
	ret := _m.Called(duration)

	if len(ret) == 0 {
		panic("no return value specified for IsHeartbeatWithinDuration")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(time.Duration) bool); ok {
		r0 = rf(duration)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsHeartbeatWithinDuration'
type DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call struct {
	*mock.Call
}

// IsHeartbeatWithinDuration is a helper method to define mock.On call
//   - duration time.Duration
func (_e *DeviceDiagnosisCommonInterface_Expecter) IsHeartbeatWithinDuration(duration interface{}) *DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call {
	return &DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call{Call: _e.mock.On("IsHeartbeatWithinDuration", duration)}
}

func (_c *DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call) Run(run func(duration time.Duration)) *DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call) Return(_a0 bool) *DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call) RunAndReturn(run func(time.Duration) bool) *DeviceDiagnosisCommonInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeviceDiagnosisCommonInterface creates a new instance of DeviceDiagnosisCommonInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceDiagnosisCommonInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceDiagnosisCommonInterface {
	mock := &DeviceDiagnosisCommonInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
