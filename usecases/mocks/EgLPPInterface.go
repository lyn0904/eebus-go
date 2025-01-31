// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/lyn0904/eebus-go/api"
	api "github.com/lyn0904/eebus-go/usecases/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/lyn0904/spine-go/model"

	spine_goapi "github.com/lyn0904/spine-go/api"

	time "time"
)

// EgLPPInterface is an autogenerated mock type for the EgLPPInterface type
type EgLPPInterface struct {
	mock.Mock
}

type EgLPPInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *EgLPPInterface) EXPECT() *EgLPPInterface_Expecter {
	return &EgLPPInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *EgLPPInterface) AddFeatures() {
	_m.Called()
}

// EgLPPInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type EgLPPInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *EgLPPInterface_Expecter) AddFeatures() *EgLPPInterface_AddFeatures_Call {
	return &EgLPPInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *EgLPPInterface_AddFeatures_Call) Run(run func()) *EgLPPInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EgLPPInterface_AddFeatures_Call) Return() *EgLPPInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPPInterface_AddFeatures_Call) RunAndReturn(run func()) *EgLPPInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *EgLPPInterface) AddUseCase() {
	_m.Called()
}

// EgLPPInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type EgLPPInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *EgLPPInterface_Expecter) AddUseCase() *EgLPPInterface_AddUseCase_Call {
	return &EgLPPInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *EgLPPInterface_AddUseCase_Call) Run(run func()) *EgLPPInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EgLPPInterface_AddUseCase_Call) Return() *EgLPPInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPPInterface_AddUseCase_Call) RunAndReturn(run func()) *EgLPPInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *EgLPPInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
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

// EgLPPInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type EgLPPInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPPInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *EgLPPInterface_AvailableScenariosForEntity_Call {
	return &EgLPPInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *EgLPPInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPPInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPPInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *EgLPPInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EgLPPInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *EgLPPInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeDurationMinimum provides a mock function with given fields: entity
func (_m *EgLPPInterface) FailsafeDurationMinimum(entity spine_goapi.EntityRemoteInterface) (time.Duration, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for FailsafeDurationMinimum")
	}

	var r0 time.Duration
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (time.Duration, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) time.Duration); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPPInterface_FailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeDurationMinimum'
type EgLPPInterface_FailsafeDurationMinimum_Call struct {
	*mock.Call
}

// FailsafeDurationMinimum is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPPInterface_Expecter) FailsafeDurationMinimum(entity interface{}) *EgLPPInterface_FailsafeDurationMinimum_Call {
	return &EgLPPInterface_FailsafeDurationMinimum_Call{Call: _e.mock.On("FailsafeDurationMinimum", entity)}
}

func (_c *EgLPPInterface_FailsafeDurationMinimum_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPPInterface_FailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPPInterface_FailsafeDurationMinimum_Call) Return(_a0 time.Duration, _a1 error) *EgLPPInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPPInterface_FailsafeDurationMinimum_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (time.Duration, error)) *EgLPPInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeProductionActivePowerLimit provides a mock function with given fields: entity
func (_m *EgLPPInterface) FailsafeProductionActivePowerLimit(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for FailsafeProductionActivePowerLimit")
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

// EgLPPInterface_FailsafeProductionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeProductionActivePowerLimit'
type EgLPPInterface_FailsafeProductionActivePowerLimit_Call struct {
	*mock.Call
}

// FailsafeProductionActivePowerLimit is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPPInterface_Expecter) FailsafeProductionActivePowerLimit(entity interface{}) *EgLPPInterface_FailsafeProductionActivePowerLimit_Call {
	return &EgLPPInterface_FailsafeProductionActivePowerLimit_Call{Call: _e.mock.On("FailsafeProductionActivePowerLimit", entity)}
}

func (_c *EgLPPInterface_FailsafeProductionActivePowerLimit_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPPInterface_FailsafeProductionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPPInterface_FailsafeProductionActivePowerLimit_Call) Return(_a0 float64, _a1 error) *EgLPPInterface_FailsafeProductionActivePowerLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPPInterface_FailsafeProductionActivePowerLimit_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *EgLPPInterface_FailsafeProductionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *EgLPPInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
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

// EgLPPInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type EgLPPInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPPInterface_Expecter) IsCompatibleEntityType(entity interface{}) *EgLPPInterface_IsCompatibleEntityType_Call {
	return &EgLPPInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *EgLPPInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPPInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPPInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *EgLPPInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EgLPPInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *EgLPPInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsHeartbeatWithinDuration provides a mock function with given fields: entity
func (_m *EgLPPInterface) IsHeartbeatWithinDuration(entity spine_goapi.EntityRemoteInterface) bool {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for IsHeartbeatWithinDuration")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EgLPPInterface_IsHeartbeatWithinDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsHeartbeatWithinDuration'
type EgLPPInterface_IsHeartbeatWithinDuration_Call struct {
	*mock.Call
}

// IsHeartbeatWithinDuration is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPPInterface_Expecter) IsHeartbeatWithinDuration(entity interface{}) *EgLPPInterface_IsHeartbeatWithinDuration_Call {
	return &EgLPPInterface_IsHeartbeatWithinDuration_Call{Call: _e.mock.On("IsHeartbeatWithinDuration", entity)}
}

func (_c *EgLPPInterface_IsHeartbeatWithinDuration_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPPInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPPInterface_IsHeartbeatWithinDuration_Call) Return(_a0 bool) *EgLPPInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EgLPPInterface_IsHeartbeatWithinDuration_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *EgLPPInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *EgLPPInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
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

// EgLPPInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type EgLPPInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *EgLPPInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *EgLPPInterface_IsScenarioAvailableAtEntity_Call {
	return &EgLPPInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *EgLPPInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *EgLPPInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *EgLPPInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *EgLPPInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EgLPPInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *EgLPPInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// ProductionLimit provides a mock function with given fields: entity
func (_m *EgLPPInterface) ProductionLimit(entity spine_goapi.EntityRemoteInterface) (api.LoadLimit, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ProductionLimit")
	}

	var r0 api.LoadLimit
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (api.LoadLimit, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) api.LoadLimit); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(api.LoadLimit)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPPInterface_ProductionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProductionLimit'
type EgLPPInterface_ProductionLimit_Call struct {
	*mock.Call
}

// ProductionLimit is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPPInterface_Expecter) ProductionLimit(entity interface{}) *EgLPPInterface_ProductionLimit_Call {
	return &EgLPPInterface_ProductionLimit_Call{Call: _e.mock.On("ProductionLimit", entity)}
}

func (_c *EgLPPInterface_ProductionLimit_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPPInterface_ProductionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPPInterface_ProductionLimit_Call) Return(limit api.LoadLimit, resultErr error) *EgLPPInterface_ProductionLimit_Call {
	_c.Call.Return(limit, resultErr)
	return _c
}

func (_c *EgLPPInterface_ProductionLimit_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (api.LoadLimit, error)) *EgLPPInterface_ProductionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// ProductionNominalMax provides a mock function with given fields: entity
func (_m *EgLPPInterface) ProductionNominalMax(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ProductionNominalMax")
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

// EgLPPInterface_ProductionNominalMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProductionNominalMax'
type EgLPPInterface_ProductionNominalMax_Call struct {
	*mock.Call
}

// ProductionNominalMax is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPPInterface_Expecter) ProductionNominalMax(entity interface{}) *EgLPPInterface_ProductionNominalMax_Call {
	return &EgLPPInterface_ProductionNominalMax_Call{Call: _e.mock.On("ProductionNominalMax", entity)}
}

func (_c *EgLPPInterface_ProductionNominalMax_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPPInterface_ProductionNominalMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPPInterface_ProductionNominalMax_Call) Return(_a0 float64, _a1 error) *EgLPPInterface_ProductionNominalMax_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPPInterface_ProductionNominalMax_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *EgLPPInterface_ProductionNominalMax_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *EgLPPInterface) RemoteEntitiesScenarios() []eebus_goapi.RemoteEntityScenarios {
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

// EgLPPInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type EgLPPInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *EgLPPInterface_Expecter) RemoteEntitiesScenarios() *EgLPPInterface_RemoteEntitiesScenarios_Call {
	return &EgLPPInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *EgLPPInterface_RemoteEntitiesScenarios_Call) Run(run func()) *EgLPPInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EgLPPInterface_RemoteEntitiesScenarios_Call) Return(_a0 []eebus_goapi.RemoteEntityScenarios) *EgLPPInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EgLPPInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []eebus_goapi.RemoteEntityScenarios) *EgLPPInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *EgLPPInterface) RemoveUseCase() {
	_m.Called()
}

// EgLPPInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type EgLPPInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *EgLPPInterface_Expecter) RemoveUseCase() *EgLPPInterface_RemoveUseCase_Call {
	return &EgLPPInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *EgLPPInterface_RemoveUseCase_Call) Run(run func()) *EgLPPInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EgLPPInterface_RemoveUseCase_Call) Return() *EgLPPInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPPInterface_RemoveUseCase_Call) RunAndReturn(run func()) *EgLPPInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// StartHeartbeat provides a mock function with given fields:
func (_m *EgLPPInterface) StartHeartbeat() {
	_m.Called()
}

// EgLPPInterface_StartHeartbeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartHeartbeat'
type EgLPPInterface_StartHeartbeat_Call struct {
	*mock.Call
}

// StartHeartbeat is a helper method to define mock.On call
func (_e *EgLPPInterface_Expecter) StartHeartbeat() *EgLPPInterface_StartHeartbeat_Call {
	return &EgLPPInterface_StartHeartbeat_Call{Call: _e.mock.On("StartHeartbeat")}
}

func (_c *EgLPPInterface_StartHeartbeat_Call) Run(run func()) *EgLPPInterface_StartHeartbeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EgLPPInterface_StartHeartbeat_Call) Return() *EgLPPInterface_StartHeartbeat_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPPInterface_StartHeartbeat_Call) RunAndReturn(run func()) *EgLPPInterface_StartHeartbeat_Call {
	_c.Call.Return(run)
	return _c
}

// StopHeartbeat provides a mock function with given fields:
func (_m *EgLPPInterface) StopHeartbeat() {
	_m.Called()
}

// EgLPPInterface_StopHeartbeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopHeartbeat'
type EgLPPInterface_StopHeartbeat_Call struct {
	*mock.Call
}

// StopHeartbeat is a helper method to define mock.On call
func (_e *EgLPPInterface_Expecter) StopHeartbeat() *EgLPPInterface_StopHeartbeat_Call {
	return &EgLPPInterface_StopHeartbeat_Call{Call: _e.mock.On("StopHeartbeat")}
}

func (_c *EgLPPInterface_StopHeartbeat_Call) Run(run func()) *EgLPPInterface_StopHeartbeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EgLPPInterface_StopHeartbeat_Call) Return() *EgLPPInterface_StopHeartbeat_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPPInterface_StopHeartbeat_Call) RunAndReturn(run func()) *EgLPPInterface_StopHeartbeat_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *EgLPPInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// EgLPPInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type EgLPPInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *EgLPPInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *EgLPPInterface_UpdateUseCaseAvailability_Call {
	return &EgLPPInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *EgLPPInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *EgLPPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *EgLPPInterface_UpdateUseCaseAvailability_Call) Return() *EgLPPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPPInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *EgLPPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFailsafeDurationMinimum provides a mock function with given fields: entity, duration
func (_m *EgLPPInterface) WriteFailsafeDurationMinimum(entity spine_goapi.EntityRemoteInterface, duration time.Duration) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, duration)

	if len(ret) == 0 {
		panic("no return value specified for WriteFailsafeDurationMinimum")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, time.Duration) (*model.MsgCounterType, error)); ok {
		return rf(entity, duration)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, time.Duration) *model.MsgCounterType); ok {
		r0 = rf(entity, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface, time.Duration) error); ok {
		r1 = rf(entity, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPPInterface_WriteFailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFailsafeDurationMinimum'
type EgLPPInterface_WriteFailsafeDurationMinimum_Call struct {
	*mock.Call
}

// WriteFailsafeDurationMinimum is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - duration time.Duration
func (_e *EgLPPInterface_Expecter) WriteFailsafeDurationMinimum(entity interface{}, duration interface{}) *EgLPPInterface_WriteFailsafeDurationMinimum_Call {
	return &EgLPPInterface_WriteFailsafeDurationMinimum_Call{Call: _e.mock.On("WriteFailsafeDurationMinimum", entity, duration)}
}

func (_c *EgLPPInterface_WriteFailsafeDurationMinimum_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, duration time.Duration)) *EgLPPInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(time.Duration))
	})
	return _c
}

func (_c *EgLPPInterface_WriteFailsafeDurationMinimum_Call) Return(_a0 *model.MsgCounterType, _a1 error) *EgLPPInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPPInterface_WriteFailsafeDurationMinimum_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, time.Duration) (*model.MsgCounterType, error)) *EgLPPInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFailsafeProductionActivePowerLimit provides a mock function with given fields: entity, value
func (_m *EgLPPInterface) WriteFailsafeProductionActivePowerLimit(entity spine_goapi.EntityRemoteInterface, value float64) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, value)

	if len(ret) == 0 {
		panic("no return value specified for WriteFailsafeProductionActivePowerLimit")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, float64) (*model.MsgCounterType, error)); ok {
		return rf(entity, value)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, float64) *model.MsgCounterType); ok {
		r0 = rf(entity, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface, float64) error); ok {
		r1 = rf(entity, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFailsafeProductionActivePowerLimit'
type EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call struct {
	*mock.Call
}

// WriteFailsafeProductionActivePowerLimit is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - value float64
func (_e *EgLPPInterface_Expecter) WriteFailsafeProductionActivePowerLimit(entity interface{}, value interface{}) *EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call {
	return &EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call{Call: _e.mock.On("WriteFailsafeProductionActivePowerLimit", entity, value)}
}

func (_c *EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, value float64)) *EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(float64))
	})
	return _c
}

func (_c *EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call) Return(_a0 *model.MsgCounterType, _a1 error) *EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, float64) (*model.MsgCounterType, error)) *EgLPPInterface_WriteFailsafeProductionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// WriteProductionLimit provides a mock function with given fields: entity, limit, resultCB
func (_m *EgLPPInterface) WriteProductionLimit(entity spine_goapi.EntityRemoteInterface, limit api.LoadLimit, resultCB func(model.ResultDataType)) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, limit, resultCB)

	if len(ret) == 0 {
		panic("no return value specified for WriteProductionLimit")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, api.LoadLimit, func(model.ResultDataType)) (*model.MsgCounterType, error)); ok {
		return rf(entity, limit, resultCB)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, api.LoadLimit, func(model.ResultDataType)) *model.MsgCounterType); ok {
		r0 = rf(entity, limit, resultCB)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface, api.LoadLimit, func(model.ResultDataType)) error); ok {
		r1 = rf(entity, limit, resultCB)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPPInterface_WriteProductionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteProductionLimit'
type EgLPPInterface_WriteProductionLimit_Call struct {
	*mock.Call
}

// WriteProductionLimit is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - limit api.LoadLimit
//   - resultCB func(model.ResultDataType)
func (_e *EgLPPInterface_Expecter) WriteProductionLimit(entity interface{}, limit interface{}, resultCB interface{}) *EgLPPInterface_WriteProductionLimit_Call {
	return &EgLPPInterface_WriteProductionLimit_Call{Call: _e.mock.On("WriteProductionLimit", entity, limit, resultCB)}
}

func (_c *EgLPPInterface_WriteProductionLimit_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, limit api.LoadLimit, resultCB func(model.ResultDataType))) *EgLPPInterface_WriteProductionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(api.LoadLimit), args[2].(func(model.ResultDataType)))
	})
	return _c
}

func (_c *EgLPPInterface_WriteProductionLimit_Call) Return(_a0 *model.MsgCounterType, _a1 error) *EgLPPInterface_WriteProductionLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPPInterface_WriteProductionLimit_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, api.LoadLimit, func(model.ResultDataType)) (*model.MsgCounterType, error)) *EgLPPInterface_WriteProductionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// NewEgLPPInterface creates a new instance of EgLPPInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEgLPPInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *EgLPPInterface {
	mock := &EgLPPInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
