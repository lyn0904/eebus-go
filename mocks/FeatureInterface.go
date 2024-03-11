// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/spine-go/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"
)

// FeatureInterface is an autogenerated mock type for the FeatureInterface type
type FeatureInterface struct {
	mock.Mock
}

type FeatureInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *FeatureInterface) EXPECT() *FeatureInterface_Expecter {
	return &FeatureInterface_Expecter{mock: &_m.Mock}
}

// AddResponseCallback provides a mock function with given fields: msgCounterReference, function
func (_m *FeatureInterface) AddResponseCallback(msgCounterReference model.MsgCounterType, function func(api.ResponseMessage)) error {
	ret := _m.Called(msgCounterReference, function)

	if len(ret) == 0 {
		panic("no return value specified for AddResponseCallback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.MsgCounterType, func(api.ResponseMessage)) error); ok {
		r0 = rf(msgCounterReference, function)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FeatureInterface_AddResponseCallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddResponseCallback'
type FeatureInterface_AddResponseCallback_Call struct {
	*mock.Call
}

// AddResponseCallback is a helper method to define mock.On call
//   - msgCounterReference model.MsgCounterType
//   - function func(api.ResponseMessage)
func (_e *FeatureInterface_Expecter) AddResponseCallback(msgCounterReference interface{}, function interface{}) *FeatureInterface_AddResponseCallback_Call {
	return &FeatureInterface_AddResponseCallback_Call{Call: _e.mock.On("AddResponseCallback", msgCounterReference, function)}
}

func (_c *FeatureInterface_AddResponseCallback_Call) Run(run func(msgCounterReference model.MsgCounterType, function func(api.ResponseMessage))) *FeatureInterface_AddResponseCallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MsgCounterType), args[1].(func(api.ResponseMessage)))
	})
	return _c
}

func (_c *FeatureInterface_AddResponseCallback_Call) Return(_a0 error) *FeatureInterface_AddResponseCallback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_AddResponseCallback_Call) RunAndReturn(run func(model.MsgCounterType, func(api.ResponseMessage)) error) *FeatureInterface_AddResponseCallback_Call {
	_c.Call.Return(run)
	return _c
}

// AddResultCallback provides a mock function with given fields: function
func (_m *FeatureInterface) AddResultCallback(function func(api.ResponseMessage)) {
	_m.Called(function)
}

// FeatureInterface_AddResultCallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddResultCallback'
type FeatureInterface_AddResultCallback_Call struct {
	*mock.Call
}

// AddResultCallback is a helper method to define mock.On call
//   - function func(api.ResponseMessage)
func (_e *FeatureInterface_Expecter) AddResultCallback(function interface{}) *FeatureInterface_AddResultCallback_Call {
	return &FeatureInterface_AddResultCallback_Call{Call: _e.mock.On("AddResultCallback", function)}
}

func (_c *FeatureInterface_AddResultCallback_Call) Run(run func(function func(api.ResponseMessage))) *FeatureInterface_AddResultCallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(api.ResponseMessage)))
	})
	return _c
}

func (_c *FeatureInterface_AddResultCallback_Call) Return() *FeatureInterface_AddResultCallback_Call {
	_c.Call.Return()
	return _c
}

func (_c *FeatureInterface_AddResultCallback_Call) RunAndReturn(run func(func(api.ResponseMessage))) *FeatureInterface_AddResultCallback_Call {
	_c.Call.Return(run)
	return _c
}

// Bind provides a mock function with given fields:
func (_m *FeatureInterface) Bind() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Bind")
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

// FeatureInterface_Bind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bind'
type FeatureInterface_Bind_Call struct {
	*mock.Call
}

// Bind is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) Bind() *FeatureInterface_Bind_Call {
	return &FeatureInterface_Bind_Call{Call: _e.mock.On("Bind")}
}

func (_c *FeatureInterface_Bind_Call) Run(run func()) *FeatureInterface_Bind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_Bind_Call) Return(_a0 *model.MsgCounterType, _a1 error) *FeatureInterface_Bind_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FeatureInterface_Bind_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *FeatureInterface_Bind_Call {
	_c.Call.Return(run)
	return _c
}

// HasBinding provides a mock function with given fields:
func (_m *FeatureInterface) HasBinding() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HasBinding")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeatureInterface_HasBinding_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasBinding'
type FeatureInterface_HasBinding_Call struct {
	*mock.Call
}

// HasBinding is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) HasBinding() *FeatureInterface_HasBinding_Call {
	return &FeatureInterface_HasBinding_Call{Call: _e.mock.On("HasBinding")}
}

func (_c *FeatureInterface_HasBinding_Call) Run(run func()) *FeatureInterface_HasBinding_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_HasBinding_Call) Return(_a0 bool) *FeatureInterface_HasBinding_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_HasBinding_Call) RunAndReturn(run func() bool) *FeatureInterface_HasBinding_Call {
	_c.Call.Return(run)
	return _c
}

// HasSubscription provides a mock function with given fields:
func (_m *FeatureInterface) HasSubscription() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HasSubscription")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeatureInterface_HasSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasSubscription'
type FeatureInterface_HasSubscription_Call struct {
	*mock.Call
}

// HasSubscription is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) HasSubscription() *FeatureInterface_HasSubscription_Call {
	return &FeatureInterface_HasSubscription_Call{Call: _e.mock.On("HasSubscription")}
}

func (_c *FeatureInterface_HasSubscription_Call) Run(run func()) *FeatureInterface_HasSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_HasSubscription_Call) Return(_a0 bool) *FeatureInterface_HasSubscription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_HasSubscription_Call) RunAndReturn(run func() bool) *FeatureInterface_HasSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// Subscribe provides a mock function with given fields:
func (_m *FeatureInterface) Subscribe() (*model.MsgCounterType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Subscribe")
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

// FeatureInterface_Subscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscribe'
type FeatureInterface_Subscribe_Call struct {
	*mock.Call
}

// Subscribe is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) Subscribe() *FeatureInterface_Subscribe_Call {
	return &FeatureInterface_Subscribe_Call{Call: _e.mock.On("Subscribe")}
}

func (_c *FeatureInterface_Subscribe_Call) Run(run func()) *FeatureInterface_Subscribe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_Subscribe_Call) Return(_a0 *model.MsgCounterType, _a1 error) *FeatureInterface_Subscribe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FeatureInterface_Subscribe_Call) RunAndReturn(run func() (*model.MsgCounterType, error)) *FeatureInterface_Subscribe_Call {
	_c.Call.Return(run)
	return _c
}

// NewFeatureInterface creates a new instance of FeatureInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeatureInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeatureInterface {
	mock := &FeatureInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
