// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "github.com/lyn0904/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// SmartEnergyManagementPsClientInterface is an autogenerated mock type for the SmartEnergyManagementPsClientInterface type
type SmartEnergyManagementPsClientInterface struct {
	mock.Mock
}

type SmartEnergyManagementPsClientInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *SmartEnergyManagementPsClientInterface) EXPECT() *SmartEnergyManagementPsClientInterface_Expecter {
	return &SmartEnergyManagementPsClientInterface_Expecter{mock: &_m.Mock}
}

// RequestData provides a mock function with given fields:
func (_m *SmartEnergyManagementPsClientInterface) RequestData() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RequestData")
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

// SmartEnergyManagementPsClientInterface_RequestData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestData'
type SmartEnergyManagementPsClientInterface_RequestData_Call struct {
	*mock.Call
}

// RequestData is a helper method to define mock.On call
func (_e *SmartEnergyManagementPsClientInterface_Expecter) RequestData() *SmartEnergyManagementPsClientInterface_RequestData_Call {
	return &SmartEnergyManagementPsClientInterface_RequestData_Call{Call: _e.mock.On("RequestData")}
}

func (_c *SmartEnergyManagementPsClientInterface_RequestData_Call) Run(run func()) *SmartEnergyManagementPsClientInterface_RequestData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SmartEnergyManagementPsClientInterface_RequestData_Call) Return(_a0 *model.MsgCounterType, _a1 error) *SmartEnergyManagementPsClientInterface_RequestData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SmartEnergyManagementPsClientInterface_RequestData_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *SmartEnergyManagementPsClientInterface_RequestData_Call {
	_c.Call.Return(run)
	return _c
}

// WriteData provides a mock function with given fields: data
func (_m *SmartEnergyManagementPsClientInterface) WriteData(data *model.SmartEnergyManagementPsDataType) (*model.MsgCounterType, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for WriteData")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.SmartEnergyManagementPsDataType) (*model.MsgCounterType, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*model.SmartEnergyManagementPsDataType) *model.MsgCounterType); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.SmartEnergyManagementPsDataType) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SmartEnergyManagementPsClientInterface_WriteData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteData'
type SmartEnergyManagementPsClientInterface_WriteData_Call struct {
	*mock.Call
}

// WriteData is a helper method to define mock.On call
//   - data *model.SmartEnergyManagementPsDataType
func (_e *SmartEnergyManagementPsClientInterface_Expecter) WriteData(data interface{}) *SmartEnergyManagementPsClientInterface_WriteData_Call {
	return &SmartEnergyManagementPsClientInterface_WriteData_Call{Call: _e.mock.On("WriteData", data)}
}

func (_c *SmartEnergyManagementPsClientInterface_WriteData_Call) Run(run func(data *model.SmartEnergyManagementPsDataType)) *SmartEnergyManagementPsClientInterface_WriteData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.SmartEnergyManagementPsDataType))
	})
	return _c
}

func (_c *SmartEnergyManagementPsClientInterface_WriteData_Call) Return(_a0 *model.MsgCounterType, _a1 error) *SmartEnergyManagementPsClientInterface_WriteData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SmartEnergyManagementPsClientInterface_WriteData_Call) RunAndReturn(run func(*model.SmartEnergyManagementPsDataType) (*model.MsgCounterType, error)) *SmartEnergyManagementPsClientInterface_WriteData_Call {
	_c.Call.Return(run)
	return _c
}

// NewSmartEnergyManagementPsClientInterface creates a new instance of SmartEnergyManagementPsClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSmartEnergyManagementPsClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SmartEnergyManagementPsClientInterface {
	mock := &SmartEnergyManagementPsClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
