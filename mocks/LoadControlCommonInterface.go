// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "github.com/lyn0904/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// LoadControlCommonInterface is an autogenerated mock type for the LoadControlCommonInterface type
type LoadControlCommonInterface struct {
	mock.Mock
}

type LoadControlCommonInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *LoadControlCommonInterface) EXPECT() *LoadControlCommonInterface_Expecter {
	return &LoadControlCommonInterface_Expecter{mock: &_m.Mock}
}

// CheckEventPayloadDataForFilter provides a mock function with given fields: payloadData, filter
func (_m *LoadControlCommonInterface) CheckEventPayloadDataForFilter(payloadData interface{}, filter interface{}) bool {
	ret := _m.Called(payloadData, filter)

	if len(ret) == 0 {
		panic("no return value specified for CheckEventPayloadDataForFilter")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) bool); ok {
		r0 = rf(payloadData, filter)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckEventPayloadDataForFilter'
type LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call struct {
	*mock.Call
}

// CheckEventPayloadDataForFilter is a helper method to define mock.On call
//   - payloadData interface{}
//   - filter interface{}
func (_e *LoadControlCommonInterface_Expecter) CheckEventPayloadDataForFilter(payloadData interface{}, filter interface{}) *LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call {
	return &LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call{Call: _e.mock.On("CheckEventPayloadDataForFilter", payloadData, filter)}
}

func (_c *LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call) Run(run func(payloadData interface{}, filter interface{})) *LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(interface{}))
	})
	return _c
}

func (_c *LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call) Return(_a0 bool) *LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call) RunAndReturn(run func(interface{}, interface{}) bool) *LoadControlCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// GetLimitDataForFilter provides a mock function with given fields: filter
func (_m *LoadControlCommonInterface) GetLimitDataForFilter(filter model.LoadControlLimitDescriptionDataType) ([]model.LoadControlLimitDataType, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetLimitDataForFilter")
	}

	var r0 []model.LoadControlLimitDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.LoadControlLimitDescriptionDataType) ([]model.LoadControlLimitDataType, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.LoadControlLimitDescriptionDataType) []model.LoadControlLimitDataType); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.LoadControlLimitDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.LoadControlLimitDescriptionDataType) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadControlCommonInterface_GetLimitDataForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLimitDataForFilter'
type LoadControlCommonInterface_GetLimitDataForFilter_Call struct {
	*mock.Call
}

// GetLimitDataForFilter is a helper method to define mock.On call
//   - filter model.LoadControlLimitDescriptionDataType
func (_e *LoadControlCommonInterface_Expecter) GetLimitDataForFilter(filter interface{}) *LoadControlCommonInterface_GetLimitDataForFilter_Call {
	return &LoadControlCommonInterface_GetLimitDataForFilter_Call{Call: _e.mock.On("GetLimitDataForFilter", filter)}
}

func (_c *LoadControlCommonInterface_GetLimitDataForFilter_Call) Run(run func(filter model.LoadControlLimitDescriptionDataType)) *LoadControlCommonInterface_GetLimitDataForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.LoadControlLimitDescriptionDataType))
	})
	return _c
}

func (_c *LoadControlCommonInterface_GetLimitDataForFilter_Call) Return(_a0 []model.LoadControlLimitDataType, _a1 error) *LoadControlCommonInterface_GetLimitDataForFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LoadControlCommonInterface_GetLimitDataForFilter_Call) RunAndReturn(run func(model.LoadControlLimitDescriptionDataType) ([]model.LoadControlLimitDataType, error)) *LoadControlCommonInterface_GetLimitDataForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// GetLimitDataForId provides a mock function with given fields: limitId
func (_m *LoadControlCommonInterface) GetLimitDataForId(limitId model.LoadControlLimitIdType) (*model.LoadControlLimitDataType, error) {
	ret := _m.Called(limitId)

	if len(ret) == 0 {
		panic("no return value specified for GetLimitDataForId")
	}

	var r0 *model.LoadControlLimitDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.LoadControlLimitIdType) (*model.LoadControlLimitDataType, error)); ok {
		return rf(limitId)
	}
	if rf, ok := ret.Get(0).(func(model.LoadControlLimitIdType) *model.LoadControlLimitDataType); ok {
		r0 = rf(limitId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoadControlLimitDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.LoadControlLimitIdType) error); ok {
		r1 = rf(limitId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadControlCommonInterface_GetLimitDataForId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLimitDataForId'
type LoadControlCommonInterface_GetLimitDataForId_Call struct {
	*mock.Call
}

// GetLimitDataForId is a helper method to define mock.On call
//   - limitId model.LoadControlLimitIdType
func (_e *LoadControlCommonInterface_Expecter) GetLimitDataForId(limitId interface{}) *LoadControlCommonInterface_GetLimitDataForId_Call {
	return &LoadControlCommonInterface_GetLimitDataForId_Call{Call: _e.mock.On("GetLimitDataForId", limitId)}
}

func (_c *LoadControlCommonInterface_GetLimitDataForId_Call) Run(run func(limitId model.LoadControlLimitIdType)) *LoadControlCommonInterface_GetLimitDataForId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.LoadControlLimitIdType))
	})
	return _c
}

func (_c *LoadControlCommonInterface_GetLimitDataForId_Call) Return(_a0 *model.LoadControlLimitDataType, _a1 error) *LoadControlCommonInterface_GetLimitDataForId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LoadControlCommonInterface_GetLimitDataForId_Call) RunAndReturn(run func(model.LoadControlLimitIdType) (*model.LoadControlLimitDataType, error)) *LoadControlCommonInterface_GetLimitDataForId_Call {
	_c.Call.Return(run)
	return _c
}

// GetLimitDescriptionForId provides a mock function with given fields: limitId
func (_m *LoadControlCommonInterface) GetLimitDescriptionForId(limitId model.LoadControlLimitIdType) (*model.LoadControlLimitDescriptionDataType, error) {
	ret := _m.Called(limitId)

	if len(ret) == 0 {
		panic("no return value specified for GetLimitDescriptionForId")
	}

	var r0 *model.LoadControlLimitDescriptionDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.LoadControlLimitIdType) (*model.LoadControlLimitDescriptionDataType, error)); ok {
		return rf(limitId)
	}
	if rf, ok := ret.Get(0).(func(model.LoadControlLimitIdType) *model.LoadControlLimitDescriptionDataType); ok {
		r0 = rf(limitId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.LoadControlLimitDescriptionDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.LoadControlLimitIdType) error); ok {
		r1 = rf(limitId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadControlCommonInterface_GetLimitDescriptionForId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLimitDescriptionForId'
type LoadControlCommonInterface_GetLimitDescriptionForId_Call struct {
	*mock.Call
}

// GetLimitDescriptionForId is a helper method to define mock.On call
//   - limitId model.LoadControlLimitIdType
func (_e *LoadControlCommonInterface_Expecter) GetLimitDescriptionForId(limitId interface{}) *LoadControlCommonInterface_GetLimitDescriptionForId_Call {
	return &LoadControlCommonInterface_GetLimitDescriptionForId_Call{Call: _e.mock.On("GetLimitDescriptionForId", limitId)}
}

func (_c *LoadControlCommonInterface_GetLimitDescriptionForId_Call) Run(run func(limitId model.LoadControlLimitIdType)) *LoadControlCommonInterface_GetLimitDescriptionForId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.LoadControlLimitIdType))
	})
	return _c
}

func (_c *LoadControlCommonInterface_GetLimitDescriptionForId_Call) Return(_a0 *model.LoadControlLimitDescriptionDataType, _a1 error) *LoadControlCommonInterface_GetLimitDescriptionForId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LoadControlCommonInterface_GetLimitDescriptionForId_Call) RunAndReturn(run func(model.LoadControlLimitIdType) (*model.LoadControlLimitDescriptionDataType, error)) *LoadControlCommonInterface_GetLimitDescriptionForId_Call {
	_c.Call.Return(run)
	return _c
}

// GetLimitDescriptionsForFilter provides a mock function with given fields: filter
func (_m *LoadControlCommonInterface) GetLimitDescriptionsForFilter(filter model.LoadControlLimitDescriptionDataType) ([]model.LoadControlLimitDescriptionDataType, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetLimitDescriptionsForFilter")
	}

	var r0 []model.LoadControlLimitDescriptionDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.LoadControlLimitDescriptionDataType) ([]model.LoadControlLimitDescriptionDataType, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.LoadControlLimitDescriptionDataType) []model.LoadControlLimitDescriptionDataType); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.LoadControlLimitDescriptionDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.LoadControlLimitDescriptionDataType) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLimitDescriptionsForFilter'
type LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call struct {
	*mock.Call
}

// GetLimitDescriptionsForFilter is a helper method to define mock.On call
//   - filter model.LoadControlLimitDescriptionDataType
func (_e *LoadControlCommonInterface_Expecter) GetLimitDescriptionsForFilter(filter interface{}) *LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call {
	return &LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call{Call: _e.mock.On("GetLimitDescriptionsForFilter", filter)}
}

func (_c *LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call) Run(run func(filter model.LoadControlLimitDescriptionDataType)) *LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.LoadControlLimitDescriptionDataType))
	})
	return _c
}

func (_c *LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call) Return(_a0 []model.LoadControlLimitDescriptionDataType, _a1 error) *LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call) RunAndReturn(run func(model.LoadControlLimitDescriptionDataType) ([]model.LoadControlLimitDescriptionDataType, error)) *LoadControlCommonInterface_GetLimitDescriptionsForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// NewLoadControlCommonInterface creates a new instance of LoadControlCommonInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoadControlCommonInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoadControlCommonInterface {
	mock := &LoadControlCommonInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
