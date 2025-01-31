// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/lyn0904/eebus-go/api"
	api "github.com/lyn0904/eebus-go/usecases/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/lyn0904/spine-go/model"

	spine_goapi "github.com/lyn0904/spine-go/api"
)

// CemOSCEVInterface is an autogenerated mock type for the CemOSCEVInterface type
type CemOSCEVInterface struct {
	mock.Mock
}

type CemOSCEVInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CemOSCEVInterface) EXPECT() *CemOSCEVInterface_Expecter {
	return &CemOSCEVInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *CemOSCEVInterface) AddFeatures() {
	_m.Called()
}

// CemOSCEVInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type CemOSCEVInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *CemOSCEVInterface_Expecter) AddFeatures() *CemOSCEVInterface_AddFeatures_Call {
	return &CemOSCEVInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *CemOSCEVInterface_AddFeatures_Call) Run(run func()) *CemOSCEVInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemOSCEVInterface_AddFeatures_Call) Return() *CemOSCEVInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemOSCEVInterface_AddFeatures_Call) RunAndReturn(run func()) *CemOSCEVInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *CemOSCEVInterface) AddUseCase() {
	_m.Called()
}

// CemOSCEVInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type CemOSCEVInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *CemOSCEVInterface_Expecter) AddUseCase() *CemOSCEVInterface_AddUseCase_Call {
	return &CemOSCEVInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *CemOSCEVInterface_AddUseCase_Call) Run(run func()) *CemOSCEVInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemOSCEVInterface_AddUseCase_Call) Return() *CemOSCEVInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemOSCEVInterface_AddUseCase_Call) RunAndReturn(run func()) *CemOSCEVInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *CemOSCEVInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
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

// CemOSCEVInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type CemOSCEVInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemOSCEVInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *CemOSCEVInterface_AvailableScenariosForEntity_Call {
	return &CemOSCEVInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *CemOSCEVInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemOSCEVInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemOSCEVInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *CemOSCEVInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemOSCEVInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *CemOSCEVInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// CurrentLimits provides a mock function with given fields: entity
func (_m *CemOSCEVInterface) CurrentLimits(entity spine_goapi.EntityRemoteInterface) ([]float64, []float64, []float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for CurrentLimits")
	}

	var r0 []float64
	var r1 []float64
	var r2 []float64
	var r3 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) ([]float64, []float64, []float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []float64); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) []float64); ok {
		r1 = rf(entity)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]float64)
		}
	}

	if rf, ok := ret.Get(2).(func(spine_goapi.EntityRemoteInterface) []float64); ok {
		r2 = rf(entity)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]float64)
		}
	}

	if rf, ok := ret.Get(3).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r3 = rf(entity)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// CemOSCEVInterface_CurrentLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CurrentLimits'
type CemOSCEVInterface_CurrentLimits_Call struct {
	*mock.Call
}

// CurrentLimits is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemOSCEVInterface_Expecter) CurrentLimits(entity interface{}) *CemOSCEVInterface_CurrentLimits_Call {
	return &CemOSCEVInterface_CurrentLimits_Call{Call: _e.mock.On("CurrentLimits", entity)}
}

func (_c *CemOSCEVInterface_CurrentLimits_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemOSCEVInterface_CurrentLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemOSCEVInterface_CurrentLimits_Call) Return(_a0 []float64, _a1 []float64, _a2 []float64, _a3 error) *CemOSCEVInterface_CurrentLimits_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

func (_c *CemOSCEVInterface_CurrentLimits_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) ([]float64, []float64, []float64, error)) *CemOSCEVInterface_CurrentLimits_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *CemOSCEVInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
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

// CemOSCEVInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type CemOSCEVInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemOSCEVInterface_Expecter) IsCompatibleEntityType(entity interface{}) *CemOSCEVInterface_IsCompatibleEntityType_Call {
	return &CemOSCEVInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *CemOSCEVInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemOSCEVInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemOSCEVInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *CemOSCEVInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemOSCEVInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CemOSCEVInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *CemOSCEVInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
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

// CemOSCEVInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type CemOSCEVInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *CemOSCEVInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *CemOSCEVInterface_IsScenarioAvailableAtEntity_Call {
	return &CemOSCEVInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *CemOSCEVInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *CemOSCEVInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *CemOSCEVInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *CemOSCEVInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemOSCEVInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *CemOSCEVInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// LoadControlLimits provides a mock function with given fields: entity
func (_m *CemOSCEVInterface) LoadControlLimits(entity spine_goapi.EntityRemoteInterface) ([]api.LoadLimitsPhase, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for LoadControlLimits")
	}

	var r0 []api.LoadLimitsPhase
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) ([]api.LoadLimitsPhase, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []api.LoadLimitsPhase); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.LoadLimitsPhase)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemOSCEVInterface_LoadControlLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadControlLimits'
type CemOSCEVInterface_LoadControlLimits_Call struct {
	*mock.Call
}

// LoadControlLimits is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CemOSCEVInterface_Expecter) LoadControlLimits(entity interface{}) *CemOSCEVInterface_LoadControlLimits_Call {
	return &CemOSCEVInterface_LoadControlLimits_Call{Call: _e.mock.On("LoadControlLimits", entity)}
}

func (_c *CemOSCEVInterface_LoadControlLimits_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CemOSCEVInterface_LoadControlLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CemOSCEVInterface_LoadControlLimits_Call) Return(limits []api.LoadLimitsPhase, resultErr error) *CemOSCEVInterface_LoadControlLimits_Call {
	_c.Call.Return(limits, resultErr)
	return _c
}

func (_c *CemOSCEVInterface_LoadControlLimits_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) ([]api.LoadLimitsPhase, error)) *CemOSCEVInterface_LoadControlLimits_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *CemOSCEVInterface) RemoteEntitiesScenarios() []eebus_goapi.RemoteEntityScenarios {
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

// CemOSCEVInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type CemOSCEVInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *CemOSCEVInterface_Expecter) RemoteEntitiesScenarios() *CemOSCEVInterface_RemoteEntitiesScenarios_Call {
	return &CemOSCEVInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *CemOSCEVInterface_RemoteEntitiesScenarios_Call) Run(run func()) *CemOSCEVInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemOSCEVInterface_RemoteEntitiesScenarios_Call) Return(_a0 []eebus_goapi.RemoteEntityScenarios) *CemOSCEVInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemOSCEVInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []eebus_goapi.RemoteEntityScenarios) *CemOSCEVInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *CemOSCEVInterface) RemoveUseCase() {
	_m.Called()
}

// CemOSCEVInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type CemOSCEVInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *CemOSCEVInterface_Expecter) RemoveUseCase() *CemOSCEVInterface_RemoveUseCase_Call {
	return &CemOSCEVInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *CemOSCEVInterface_RemoveUseCase_Call) Run(run func()) *CemOSCEVInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemOSCEVInterface_RemoveUseCase_Call) Return() *CemOSCEVInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemOSCEVInterface_RemoveUseCase_Call) RunAndReturn(run func()) *CemOSCEVInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// SetOperatingState provides a mock function with given fields: failureState
func (_m *CemOSCEVInterface) SetOperatingState(failureState bool) error {
	ret := _m.Called(failureState)

	if len(ret) == 0 {
		panic("no return value specified for SetOperatingState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(failureState)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CemOSCEVInterface_SetOperatingState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetOperatingState'
type CemOSCEVInterface_SetOperatingState_Call struct {
	*mock.Call
}

// SetOperatingState is a helper method to define mock.On call
//   - failureState bool
func (_e *CemOSCEVInterface_Expecter) SetOperatingState(failureState interface{}) *CemOSCEVInterface_SetOperatingState_Call {
	return &CemOSCEVInterface_SetOperatingState_Call{Call: _e.mock.On("SetOperatingState", failureState)}
}

func (_c *CemOSCEVInterface_SetOperatingState_Call) Run(run func(failureState bool)) *CemOSCEVInterface_SetOperatingState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CemOSCEVInterface_SetOperatingState_Call) Return(_a0 error) *CemOSCEVInterface_SetOperatingState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CemOSCEVInterface_SetOperatingState_Call) RunAndReturn(run func(bool) error) *CemOSCEVInterface_SetOperatingState_Call {
	_c.Call.Return(run)
	return _c
}

// StartHeartbeat provides a mock function with given fields:
func (_m *CemOSCEVInterface) StartHeartbeat() {
	_m.Called()
}

// CemOSCEVInterface_StartHeartbeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartHeartbeat'
type CemOSCEVInterface_StartHeartbeat_Call struct {
	*mock.Call
}

// StartHeartbeat is a helper method to define mock.On call
func (_e *CemOSCEVInterface_Expecter) StartHeartbeat() *CemOSCEVInterface_StartHeartbeat_Call {
	return &CemOSCEVInterface_StartHeartbeat_Call{Call: _e.mock.On("StartHeartbeat")}
}

func (_c *CemOSCEVInterface_StartHeartbeat_Call) Run(run func()) *CemOSCEVInterface_StartHeartbeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemOSCEVInterface_StartHeartbeat_Call) Return() *CemOSCEVInterface_StartHeartbeat_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemOSCEVInterface_StartHeartbeat_Call) RunAndReturn(run func()) *CemOSCEVInterface_StartHeartbeat_Call {
	_c.Call.Return(run)
	return _c
}

// StopHeartbeat provides a mock function with given fields:
func (_m *CemOSCEVInterface) StopHeartbeat() {
	_m.Called()
}

// CemOSCEVInterface_StopHeartbeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopHeartbeat'
type CemOSCEVInterface_StopHeartbeat_Call struct {
	*mock.Call
}

// StopHeartbeat is a helper method to define mock.On call
func (_e *CemOSCEVInterface_Expecter) StopHeartbeat() *CemOSCEVInterface_StopHeartbeat_Call {
	return &CemOSCEVInterface_StopHeartbeat_Call{Call: _e.mock.On("StopHeartbeat")}
}

func (_c *CemOSCEVInterface_StopHeartbeat_Call) Run(run func()) *CemOSCEVInterface_StopHeartbeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CemOSCEVInterface_StopHeartbeat_Call) Return() *CemOSCEVInterface_StopHeartbeat_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemOSCEVInterface_StopHeartbeat_Call) RunAndReturn(run func()) *CemOSCEVInterface_StopHeartbeat_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *CemOSCEVInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// CemOSCEVInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type CemOSCEVInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *CemOSCEVInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *CemOSCEVInterface_UpdateUseCaseAvailability_Call {
	return &CemOSCEVInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *CemOSCEVInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *CemOSCEVInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CemOSCEVInterface_UpdateUseCaseAvailability_Call) Return() *CemOSCEVInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *CemOSCEVInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *CemOSCEVInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// WriteLoadControlLimits provides a mock function with given fields: entity, limits, resultCB
func (_m *CemOSCEVInterface) WriteLoadControlLimits(entity spine_goapi.EntityRemoteInterface, limits []api.LoadLimitsPhase, resultCB func(model.ResultDataType)) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, limits, resultCB)

	if len(ret) == 0 {
		panic("no return value specified for WriteLoadControlLimits")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, []api.LoadLimitsPhase, func(model.ResultDataType)) (*model.MsgCounterType, error)); ok {
		return rf(entity, limits, resultCB)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, []api.LoadLimitsPhase, func(model.ResultDataType)) *model.MsgCounterType); ok {
		r0 = rf(entity, limits, resultCB)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface, []api.LoadLimitsPhase, func(model.ResultDataType)) error); ok {
		r1 = rf(entity, limits, resultCB)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CemOSCEVInterface_WriteLoadControlLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteLoadControlLimits'
type CemOSCEVInterface_WriteLoadControlLimits_Call struct {
	*mock.Call
}

// WriteLoadControlLimits is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - limits []api.LoadLimitsPhase
//   - resultCB func(model.ResultDataType)
func (_e *CemOSCEVInterface_Expecter) WriteLoadControlLimits(entity interface{}, limits interface{}, resultCB interface{}) *CemOSCEVInterface_WriteLoadControlLimits_Call {
	return &CemOSCEVInterface_WriteLoadControlLimits_Call{Call: _e.mock.On("WriteLoadControlLimits", entity, limits, resultCB)}
}

func (_c *CemOSCEVInterface_WriteLoadControlLimits_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, limits []api.LoadLimitsPhase, resultCB func(model.ResultDataType))) *CemOSCEVInterface_WriteLoadControlLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].([]api.LoadLimitsPhase), args[2].(func(model.ResultDataType)))
	})
	return _c
}

func (_c *CemOSCEVInterface_WriteLoadControlLimits_Call) Return(_a0 *model.MsgCounterType, _a1 error) *CemOSCEVInterface_WriteLoadControlLimits_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CemOSCEVInterface_WriteLoadControlLimits_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, []api.LoadLimitsPhase, func(model.ResultDataType)) (*model.MsgCounterType, error)) *CemOSCEVInterface_WriteLoadControlLimits_Call {
	_c.Call.Return(run)
	return _c
}

// NewCemOSCEVInterface creates a new instance of CemOSCEVInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCemOSCEVInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CemOSCEVInterface {
	mock := &CemOSCEVInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
