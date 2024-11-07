// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/lyn0904/eebus-go/api"
	mock "github.com/stretchr/testify/mock"

	spine_goapi "github.com/lyn0904/spine-go/api"
)

// CemVAPDInterface is an autogenerated mock type for the CemVAPDInterface type
type CemVAPDInterface struct {
	mock.Mock
}

type CemVAPDInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CemVAPDInterface) EXPECT() *CemVAPDInterface_Expecter {
	return &CemVAPDInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *CemVAPDInterface) AddFeatures() {
	_m.Called()
}

// CemVAPDInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type CemVAPDInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *CemVAPDInterface_Expecter) AddFeatures() *CemVAPDInterface_AddFeatures_Call {
	return &CemVAPDInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *CemVAPDInterface_AddFeatures_Call) Run(run func()) *CemVAPDInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemVAPDInterface_AddFeatures_Call) Return() *CemVAPDInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemVAPDInterface_AddFeatures_Call) RunAndReturn(run func()) *CemVAPDInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *CemVAPDInterface) AddUseCase() {
	_m.Called()
}

// CemVAPDInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type CemVAPDInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *CemVAPDInterface_Expecter) AddUseCase() *CemVAPDInterface_AddUseCase_Call {
	return &CemVAPDInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *CemVAPDInterface_AddUseCase_Call) Run(run func()) *CemVAPDInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemVAPDInterface_AddUseCase_Call) Return() *CemVAPDInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemVAPDInterface_AddUseCase_Call) RunAndReturn(run func()) *CemVAPDInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *CemVAPDInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for AvailableScenariosForEntity")
	}

	var r0 []uint
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []uint); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint)
		}
	}

	return r0
}

// CemVAPDInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type CemVAPDInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVAPDInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *CemVAPDInterface_AvailableScenariosForEntity_Call {
	return &CemVAPDInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *CemVAPDInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVAPDInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVAPDInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *CemVAPDInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemVAPDInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *CemVAPDInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *CemVAPDInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for IsCompatibleEntityType")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CemVAPDInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type CemVAPDInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVAPDInterface_Expecter) IsCompatibleEntityType(entity interface{}) *CemVAPDInterface_IsCompatibleEntityType_Call {
	return &CemVAPDInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *CemVAPDInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVAPDInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVAPDInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *CemVAPDInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemVAPDInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CemVAPDInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *CemVAPDInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
	ret := _m.Called(entity, scenario)

	if len(ret) == 0 {
		panic("no return value specified for IsScenarioAvailableAtEntity")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, uint) bool); ok {
		r0 = rf(entity, scenario)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CemVAPDInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type CemVAPDInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *CemVAPDInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *CemVAPDInterface_IsScenarioAvailableAtEntity_Call {
	return &CemVAPDInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *CemVAPDInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *CemVAPDInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *CemVAPDInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *CemVAPDInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemVAPDInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *CemVAPDInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// PVYieldTotal provides a mock function with given fields: entity
func (_m *CemVAPDInterface) PVYieldTotal(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for PVYieldTotal")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemVAPDInterface_PVYieldTotal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PVYieldTotal'
type CemVAPDInterface_PVYieldTotal_Call struct {
	*mock.Call
}

// PVYieldTotal is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVAPDInterface_Expecter) PVYieldTotal(entity interface{}) *CemVAPDInterface_PVYieldTotal_Call {
	return &CemVAPDInterface_PVYieldTotal_Call{Call: _e.mock.On("PVYieldTotal", entity)}
}

func (_c *CemVAPDInterface_PVYieldTotal_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVAPDInterface_PVYieldTotal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVAPDInterface_PVYieldTotal_Call) Return(_a0 float64, _a1 error) *CemVAPDInterface_PVYieldTotal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemVAPDInterface_PVYieldTotal_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *CemVAPDInterface_PVYieldTotal_Call {
	_c.Call.Return(run)
	return _c
}

// Power provides a mock function with given fields: entity
func (_m *CemVAPDInterface) Power(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for Power")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemVAPDInterface_Power_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Power'
type CemVAPDInterface_Power_Call struct {
	*mock.Call
}

// Power is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVAPDInterface_Expecter) Power(entity interface{}) *CemVAPDInterface_Power_Call {
	return &CemVAPDInterface_Power_Call{Call: _e.mock.On("Power", entity)}
}

func (_c *CemVAPDInterface_Power_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVAPDInterface_Power_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVAPDInterface_Power_Call) Return(_a0 float64, _a1 error) *CemVAPDInterface_Power_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemVAPDInterface_Power_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *CemVAPDInterface_Power_Call {
	_c.Call.Return(run)
	return _c
}

// PowerNominalPeak provides a mock function with given fields: entity
func (_m *CemVAPDInterface) PowerNominalPeak(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for PowerNominalPeak")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemVAPDInterface_PowerNominalPeak_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PowerNominalPeak'
type CemVAPDInterface_PowerNominalPeak_Call struct {
	*mock.Call
}

// PowerNominalPeak is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVAPDInterface_Expecter) PowerNominalPeak(entity interface{}) *CemVAPDInterface_PowerNominalPeak_Call {
	return &CemVAPDInterface_PowerNominalPeak_Call{Call: _e.mock.On("PowerNominalPeak", entity)}
}

func (_c *CemVAPDInterface_PowerNominalPeak_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVAPDInterface_PowerNominalPeak_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVAPDInterface_PowerNominalPeak_Call) Return(_a0 float64, _a1 error) *CemVAPDInterface_PowerNominalPeak_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemVAPDInterface_PowerNominalPeak_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *CemVAPDInterface_PowerNominalPeak_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *CemVAPDInterface) RemoteEntitiesScenarios() []eebus_goapi.RemoteEntityScenarios {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoteEntitiesScenarios")
	}

	var r0 []eebus_goapi.RemoteEntityScenarios
	if rf, ok := ret.Get(0).(func() []eebus_goapi.RemoteEntityScenarios); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]eebus_goapi.RemoteEntityScenarios)
		}
	}

	return r0
}

// CemVAPDInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type CemVAPDInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *CemVAPDInterface_Expecter) RemoteEntitiesScenarios() *CemVAPDInterface_RemoteEntitiesScenarios_Call {
	return &CemVAPDInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *CemVAPDInterface_RemoteEntitiesScenarios_Call) Run(run func()) *CemVAPDInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemVAPDInterface_RemoteEntitiesScenarios_Call) Return(_a0 []eebus_goapi.RemoteEntityScenarios) *CemVAPDInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemVAPDInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []eebus_goapi.RemoteEntityScenarios) *CemVAPDInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *CemVAPDInterface) RemoveUseCase() {
	_m.Called()
}

// CemVAPDInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type CemVAPDInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *CemVAPDInterface_Expecter) RemoveUseCase() *CemVAPDInterface_RemoveUseCase_Call {
	return &CemVAPDInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *CemVAPDInterface_RemoveUseCase_Call) Run(run func()) *CemVAPDInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemVAPDInterface_RemoveUseCase_Call) Return() *CemVAPDInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemVAPDInterface_RemoveUseCase_Call) RunAndReturn(run func()) *CemVAPDInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *CemVAPDInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// CemVAPDInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type CemVAPDInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *CemVAPDInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *CemVAPDInterface_UpdateUseCaseAvailability_Call {
	return &CemVAPDInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *CemVAPDInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *CemVAPDInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CemVAPDInterface_UpdateUseCaseAvailability_Call) Return() *CemVAPDInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemVAPDInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *CemVAPDInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// NewCemVAPDInterface creates a new instance of CemVAPDInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCemVAPDInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CemVAPDInterface {
	mock := &CemVAPDInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
