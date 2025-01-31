// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "github.com/lyn0904/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// MeasurementCommonInterface is an autogenerated mock type for the MeasurementCommonInterface type
type MeasurementCommonInterface struct {
	mock.Mock
}

type MeasurementCommonInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MeasurementCommonInterface) EXPECT() *MeasurementCommonInterface_Expecter {
	return &MeasurementCommonInterface_Expecter{mock: &_m.Mock}
}

// CheckEventPayloadDataForFilter provides a mock function with given fields: payloadData, filter
func (_m *MeasurementCommonInterface) CheckEventPayloadDataForFilter(payloadData interface{}, filter interface{}) bool {
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

// MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckEventPayloadDataForFilter'
type MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call struct {
	*mock.Call
}

// CheckEventPayloadDataForFilter is a helper method to define mock.On call
//   - payloadData interface{}
//   - filter interface{}
func (_e *MeasurementCommonInterface_Expecter) CheckEventPayloadDataForFilter(payloadData interface{}, filter interface{}) *MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call {
	return &MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call{Call: _e.mock.On("CheckEventPayloadDataForFilter", payloadData, filter)}
}

func (_c *MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call) Run(run func(payloadData interface{}, filter interface{})) *MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(interface{}))
	})
	return _c
}

func (_c *MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call) Return(_a0 bool) *MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call) RunAndReturn(run func(interface{}, interface{}) bool) *MeasurementCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// GetConstraintsForFilter provides a mock function with given fields: filter
func (_m *MeasurementCommonInterface) GetConstraintsForFilter(filter model.MeasurementConstraintsDataType) ([]model.MeasurementConstraintsDataType, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetConstraintsForFilter")
	}

	var r0 []model.MeasurementConstraintsDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.MeasurementConstraintsDataType) ([]model.MeasurementConstraintsDataType, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.MeasurementConstraintsDataType) []model.MeasurementConstraintsDataType); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.MeasurementConstraintsDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.MeasurementConstraintsDataType) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeasurementCommonInterface_GetConstraintsForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConstraintsForFilter'
type MeasurementCommonInterface_GetConstraintsForFilter_Call struct {
	*mock.Call
}

// GetConstraintsForFilter is a helper method to define mock.On call
//   - filter model.MeasurementConstraintsDataType
func (_e *MeasurementCommonInterface_Expecter) GetConstraintsForFilter(filter interface{}) *MeasurementCommonInterface_GetConstraintsForFilter_Call {
	return &MeasurementCommonInterface_GetConstraintsForFilter_Call{Call: _e.mock.On("GetConstraintsForFilter", filter)}
}

func (_c *MeasurementCommonInterface_GetConstraintsForFilter_Call) Run(run func(filter model.MeasurementConstraintsDataType)) *MeasurementCommonInterface_GetConstraintsForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MeasurementConstraintsDataType))
	})
	return _c
}

func (_c *MeasurementCommonInterface_GetConstraintsForFilter_Call) Return(_a0 []model.MeasurementConstraintsDataType, _a1 error) *MeasurementCommonInterface_GetConstraintsForFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MeasurementCommonInterface_GetConstraintsForFilter_Call) RunAndReturn(run func(model.MeasurementConstraintsDataType) ([]model.MeasurementConstraintsDataType, error)) *MeasurementCommonInterface_GetConstraintsForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataForFilter provides a mock function with given fields: filter
func (_m *MeasurementCommonInterface) GetDataForFilter(filter model.MeasurementDescriptionDataType) ([]model.MeasurementDataType, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetDataForFilter")
	}

	var r0 []model.MeasurementDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.MeasurementDescriptionDataType) ([]model.MeasurementDataType, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.MeasurementDescriptionDataType) []model.MeasurementDataType); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.MeasurementDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.MeasurementDescriptionDataType) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeasurementCommonInterface_GetDataForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataForFilter'
type MeasurementCommonInterface_GetDataForFilter_Call struct {
	*mock.Call
}

// GetDataForFilter is a helper method to define mock.On call
//   - filter model.MeasurementDescriptionDataType
func (_e *MeasurementCommonInterface_Expecter) GetDataForFilter(filter interface{}) *MeasurementCommonInterface_GetDataForFilter_Call {
	return &MeasurementCommonInterface_GetDataForFilter_Call{Call: _e.mock.On("GetDataForFilter", filter)}
}

func (_c *MeasurementCommonInterface_GetDataForFilter_Call) Run(run func(filter model.MeasurementDescriptionDataType)) *MeasurementCommonInterface_GetDataForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MeasurementDescriptionDataType))
	})
	return _c
}

func (_c *MeasurementCommonInterface_GetDataForFilter_Call) Return(_a0 []model.MeasurementDataType, _a1 error) *MeasurementCommonInterface_GetDataForFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MeasurementCommonInterface_GetDataForFilter_Call) RunAndReturn(run func(model.MeasurementDescriptionDataType) ([]model.MeasurementDataType, error)) *MeasurementCommonInterface_GetDataForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataForId provides a mock function with given fields: measurementId
func (_m *MeasurementCommonInterface) GetDataForId(measurementId model.MeasurementIdType) (*model.MeasurementDataType, error) {
	ret := _m.Called(measurementId)

	if len(ret) == 0 {
		panic("no return value specified for GetDataForId")
	}

	var r0 *model.MeasurementDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.MeasurementIdType) (*model.MeasurementDataType, error)); ok {
		return rf(measurementId)
	}
	if rf, ok := ret.Get(0).(func(model.MeasurementIdType) *model.MeasurementDataType); ok {
		r0 = rf(measurementId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MeasurementDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.MeasurementIdType) error); ok {
		r1 = rf(measurementId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeasurementCommonInterface_GetDataForId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataForId'
type MeasurementCommonInterface_GetDataForId_Call struct {
	*mock.Call
}

// GetDataForId is a helper method to define mock.On call
//   - measurementId model.MeasurementIdType
func (_e *MeasurementCommonInterface_Expecter) GetDataForId(measurementId interface{}) *MeasurementCommonInterface_GetDataForId_Call {
	return &MeasurementCommonInterface_GetDataForId_Call{Call: _e.mock.On("GetDataForId", measurementId)}
}

func (_c *MeasurementCommonInterface_GetDataForId_Call) Run(run func(measurementId model.MeasurementIdType)) *MeasurementCommonInterface_GetDataForId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MeasurementIdType))
	})
	return _c
}

func (_c *MeasurementCommonInterface_GetDataForId_Call) Return(_a0 *model.MeasurementDataType, _a1 error) *MeasurementCommonInterface_GetDataForId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MeasurementCommonInterface_GetDataForId_Call) RunAndReturn(run func(model.MeasurementIdType) (*model.MeasurementDataType, error)) *MeasurementCommonInterface_GetDataForId_Call {
	_c.Call.Return(run)
	return _c
}

// GetDescriptionForId provides a mock function with given fields: measurementId
func (_m *MeasurementCommonInterface) GetDescriptionForId(measurementId model.MeasurementIdType) (*model.MeasurementDescriptionDataType, error) {
	ret := _m.Called(measurementId)

	if len(ret) == 0 {
		panic("no return value specified for GetDescriptionForId")
	}

	var r0 *model.MeasurementDescriptionDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.MeasurementIdType) (*model.MeasurementDescriptionDataType, error)); ok {
		return rf(measurementId)
	}
	if rf, ok := ret.Get(0).(func(model.MeasurementIdType) *model.MeasurementDescriptionDataType); ok {
		r0 = rf(measurementId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MeasurementDescriptionDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.MeasurementIdType) error); ok {
		r1 = rf(measurementId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeasurementCommonInterface_GetDescriptionForId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDescriptionForId'
type MeasurementCommonInterface_GetDescriptionForId_Call struct {
	*mock.Call
}

// GetDescriptionForId is a helper method to define mock.On call
//   - measurementId model.MeasurementIdType
func (_e *MeasurementCommonInterface_Expecter) GetDescriptionForId(measurementId interface{}) *MeasurementCommonInterface_GetDescriptionForId_Call {
	return &MeasurementCommonInterface_GetDescriptionForId_Call{Call: _e.mock.On("GetDescriptionForId", measurementId)}
}

func (_c *MeasurementCommonInterface_GetDescriptionForId_Call) Run(run func(measurementId model.MeasurementIdType)) *MeasurementCommonInterface_GetDescriptionForId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MeasurementIdType))
	})
	return _c
}

func (_c *MeasurementCommonInterface_GetDescriptionForId_Call) Return(_a0 *model.MeasurementDescriptionDataType, _a1 error) *MeasurementCommonInterface_GetDescriptionForId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MeasurementCommonInterface_GetDescriptionForId_Call) RunAndReturn(run func(model.MeasurementIdType) (*model.MeasurementDescriptionDataType, error)) *MeasurementCommonInterface_GetDescriptionForId_Call {
	_c.Call.Return(run)
	return _c
}

// GetDescriptionsForFilter provides a mock function with given fields: filter
func (_m *MeasurementCommonInterface) GetDescriptionsForFilter(filter model.MeasurementDescriptionDataType) ([]model.MeasurementDescriptionDataType, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetDescriptionsForFilter")
	}

	var r0 []model.MeasurementDescriptionDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.MeasurementDescriptionDataType) ([]model.MeasurementDescriptionDataType, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.MeasurementDescriptionDataType) []model.MeasurementDescriptionDataType); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.MeasurementDescriptionDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.MeasurementDescriptionDataType) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeasurementCommonInterface_GetDescriptionsForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDescriptionsForFilter'
type MeasurementCommonInterface_GetDescriptionsForFilter_Call struct {
	*mock.Call
}

// GetDescriptionsForFilter is a helper method to define mock.On call
//   - filter model.MeasurementDescriptionDataType
func (_e *MeasurementCommonInterface_Expecter) GetDescriptionsForFilter(filter interface{}) *MeasurementCommonInterface_GetDescriptionsForFilter_Call {
	return &MeasurementCommonInterface_GetDescriptionsForFilter_Call{Call: _e.mock.On("GetDescriptionsForFilter", filter)}
}

func (_c *MeasurementCommonInterface_GetDescriptionsForFilter_Call) Run(run func(filter model.MeasurementDescriptionDataType)) *MeasurementCommonInterface_GetDescriptionsForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MeasurementDescriptionDataType))
	})
	return _c
}

func (_c *MeasurementCommonInterface_GetDescriptionsForFilter_Call) Return(_a0 []model.MeasurementDescriptionDataType, _a1 error) *MeasurementCommonInterface_GetDescriptionsForFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MeasurementCommonInterface_GetDescriptionsForFilter_Call) RunAndReturn(run func(model.MeasurementDescriptionDataType) ([]model.MeasurementDescriptionDataType, error)) *MeasurementCommonInterface_GetDescriptionsForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// NewMeasurementCommonInterface creates a new instance of MeasurementCommonInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMeasurementCommonInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MeasurementCommonInterface {
	mock := &MeasurementCommonInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
