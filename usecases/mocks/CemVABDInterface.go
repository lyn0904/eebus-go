// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/enbility/eebus-go/api"
	mock "github.com/stretchr/testify/mock"

	spine_goapi "github.com/lyn0904/spine-go/api"
)

// CemVABDInterface is an autogenerated mock type for the CemVABDInterface type
type CemVABDInterface struct {
	mock.Mock
}

type CemVABDInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CemVABDInterface) EXPECT() *CemVABDInterface_Expecter {
	return &CemVABDInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *CemVABDInterface) AddFeatures() {
	_m.Called()
}

// CemVABDInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type CemVABDInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *CemVABDInterface_Expecter) AddFeatures() *CemVABDInterface_AddFeatures_Call {
	return &CemVABDInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *CemVABDInterface_AddFeatures_Call) Run(run func()) *CemVABDInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemVABDInterface_AddFeatures_Call) Return() *CemVABDInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemVABDInterface_AddFeatures_Call) RunAndReturn(run func()) *CemVABDInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *CemVABDInterface) AddUseCase() {
	_m.Called()
}

// CemVABDInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type CemVABDInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *CemVABDInterface_Expecter) AddUseCase() *CemVABDInterface_AddUseCase_Call {
	return &CemVABDInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *CemVABDInterface_AddUseCase_Call) Run(run func()) *CemVABDInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemVABDInterface_AddUseCase_Call) Return() *CemVABDInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemVABDInterface_AddUseCase_Call) RunAndReturn(run func()) *CemVABDInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *CemVABDInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
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

// CemVABDInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type CemVABDInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVABDInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *CemVABDInterface_AvailableScenariosForEntity_Call {
	return &CemVABDInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *CemVABDInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVABDInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVABDInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *CemVABDInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemVABDInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *CemVABDInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// EnergyCharged provides a mock function with given fields: entity
func (_m *CemVABDInterface) EnergyCharged(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for EnergyCharged")
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

// CemVABDInterface_EnergyCharged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnergyCharged'
type CemVABDInterface_EnergyCharged_Call struct {
	*mock.Call
}

// EnergyCharged is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVABDInterface_Expecter) EnergyCharged(entity interface{}) *CemVABDInterface_EnergyCharged_Call {
	return &CemVABDInterface_EnergyCharged_Call{Call: _e.mock.On("EnergyCharged", entity)}
}

func (_c *CemVABDInterface_EnergyCharged_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVABDInterface_EnergyCharged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVABDInterface_EnergyCharged_Call) Return(_a0 float64, _a1 error) *CemVABDInterface_EnergyCharged_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemVABDInterface_EnergyCharged_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *CemVABDInterface_EnergyCharged_Call {
	_c.Call.Return(run)
	return _c
}

// EnergyDischarged provides a mock function with given fields: entity
func (_m *CemVABDInterface) EnergyDischarged(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for EnergyDischarged")
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

// CemVABDInterface_EnergyDischarged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnergyDischarged'
type CemVABDInterface_EnergyDischarged_Call struct {
	*mock.Call
}

// EnergyDischarged is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVABDInterface_Expecter) EnergyDischarged(entity interface{}) *CemVABDInterface_EnergyDischarged_Call {
	return &CemVABDInterface_EnergyDischarged_Call{Call: _e.mock.On("EnergyDischarged", entity)}
}

func (_c *CemVABDInterface_EnergyDischarged_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVABDInterface_EnergyDischarged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVABDInterface_EnergyDischarged_Call) Return(_a0 float64, _a1 error) *CemVABDInterface_EnergyDischarged_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemVABDInterface_EnergyDischarged_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *CemVABDInterface_EnergyDischarged_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *CemVABDInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
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

// CemVABDInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type CemVABDInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVABDInterface_Expecter) IsCompatibleEntityType(entity interface{}) *CemVABDInterface_IsCompatibleEntityType_Call {
	return &CemVABDInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *CemVABDInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVABDInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVABDInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *CemVABDInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemVABDInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CemVABDInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *CemVABDInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
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

// CemVABDInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type CemVABDInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *CemVABDInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *CemVABDInterface_IsScenarioAvailableAtEntity_Call {
	return &CemVABDInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *CemVABDInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *CemVABDInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *CemVABDInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *CemVABDInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemVABDInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *CemVABDInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// Power provides a mock function with given fields: entity
func (_m *CemVABDInterface) Power(entity spine_goapi.EntityRemoteInterface) (float64, error) {
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

// CemVABDInterface_Power_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Power'
type CemVABDInterface_Power_Call struct {
	*mock.Call
}

// Power is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVABDInterface_Expecter) Power(entity interface{}) *CemVABDInterface_Power_Call {
	return &CemVABDInterface_Power_Call{Call: _e.mock.On("Power", entity)}
}

func (_c *CemVABDInterface_Power_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVABDInterface_Power_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVABDInterface_Power_Call) Return(_a0 float64, _a1 error) *CemVABDInterface_Power_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemVABDInterface_Power_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *CemVABDInterface_Power_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *CemVABDInterface) RemoteEntitiesScenarios() []eebus_goapi.RemoteEntityScenarios {
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

// CemVABDInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type CemVABDInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *CemVABDInterface_Expecter) RemoteEntitiesScenarios() *CemVABDInterface_RemoteEntitiesScenarios_Call {
	return &CemVABDInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *CemVABDInterface_RemoteEntitiesScenarios_Call) Run(run func()) *CemVABDInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemVABDInterface_RemoteEntitiesScenarios_Call) Return(_a0 []eebus_goapi.RemoteEntityScenarios) *CemVABDInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemVABDInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []eebus_goapi.RemoteEntityScenarios) *CemVABDInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *CemVABDInterface) RemoveUseCase() {
	_m.Called()
}

// CemVABDInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type CemVABDInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *CemVABDInterface_Expecter) RemoveUseCase() *CemVABDInterface_RemoveUseCase_Call {
	return &CemVABDInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *CemVABDInterface_RemoveUseCase_Call) Run(run func()) *CemVABDInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemVABDInterface_RemoveUseCase_Call) Return() *CemVABDInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemVABDInterface_RemoveUseCase_Call) RunAndReturn(run func()) *CemVABDInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// StateOfCharge provides a mock function with given fields: entity
func (_m *CemVABDInterface) StateOfCharge(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for StateOfCharge")
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

// CemVABDInterface_StateOfCharge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StateOfCharge'
type CemVABDInterface_StateOfCharge_Call struct {
	*mock.Call
}

// StateOfCharge is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemVABDInterface_Expecter) StateOfCharge(entity interface{}) *CemVABDInterface_StateOfCharge_Call {
	return &CemVABDInterface_StateOfCharge_Call{Call: _e.mock.On("StateOfCharge", entity)}
}

func (_c *CemVABDInterface_StateOfCharge_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemVABDInterface_StateOfCharge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemVABDInterface_StateOfCharge_Call) Return(_a0 float64, _a1 error) *CemVABDInterface_StateOfCharge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemVABDInterface_StateOfCharge_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *CemVABDInterface_StateOfCharge_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *CemVABDInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// CemVABDInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type CemVABDInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *CemVABDInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *CemVABDInterface_UpdateUseCaseAvailability_Call {
	return &CemVABDInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *CemVABDInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *CemVABDInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CemVABDInterface_UpdateUseCaseAvailability_Call) Return() *CemVABDInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemVABDInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *CemVABDInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// NewCemVABDInterface creates a new instance of CemVABDInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCemVABDInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CemVABDInterface {
	mock := &CemVABDInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
