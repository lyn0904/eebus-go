package api

import (
	"time"

	"github.com/enbility/eebus-go/api"
	spineapi "github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
)

// Actor: Energy Guard
// UseCase: Limitation of Power Consumption
type EgLPCInterface interface {
	api.UseCaseInterface

	// Scenario 1

	// return the current consumption limit data
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//
	// return values:
	//   - limit: load limit data
	//
	// possible errors:
	//   - ErrDataNotAvailable if no such limit is (yet) available
	//   - and others
	ConsumptionLimit(entity spineapi.EntityRemoteInterface) (limit LoadLimit, resultErr error)

	// send new LoadControlLimits
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//   - limit: load limit data
	//   - resultCB: callback function for handling the result response
	WriteConsumptionLimit(
		entity spineapi.EntityRemoteInterface,
		limit LoadLimit,
		resultCB func(result model.ResultDataType),
	) (*model.MsgCounterType, error)

	// Scenario 2

	// return Failsafe limit for the consumed active (real) power of the
	// Controllable System. This limit becomes activated in "init" state or "failsafe state".
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//
	// return values:
	//   - positive values are used for consumption
	FailsafeConsumptionActivePowerLimit(entity spineapi.EntityRemoteInterface) (float64, error)

	// send new Failsafe Consumption Active Power Limit
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//   - value: the new limit in W
	WriteFailsafeConsumptionActivePowerLimit(entity spineapi.EntityRemoteInterface, value float64) (*model.MsgCounterType, error)

	// return minimum time the Controllable System remains in "failsafe state" unless conditions
	// specified in this Use Case permit leaving the "failsafe state"
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//
	// return values:
	//   - negative values are used for production
	FailsafeDurationMinimum(entity spineapi.EntityRemoteInterface) (time.Duration, error)

	// send new Failsafe Duration Minimum
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	//   - duration: the duration, between 2h and 24h
	WriteFailsafeDurationMinimum(entity spineapi.EntityRemoteInterface, duration time.Duration) (*model.MsgCounterType, error)

	// Scenario 3

	// start sending heartbeat from the local entity supporting this usecase
	//
	// the heartbeat is started by default when a non 0 timeout is set in the service configuration
	StartHeartbeat()

	// stop sending heartbeat from the local entity supporting this usecase
	StopHeartbeat()

	// check wether there was a heartbeat received within the last 2 minutes
	//
	// returns true, if the last heartbeat is within 2 minutes, otherwise false
	IsHeartbeatWithinDuration(entity spineapi.EntityRemoteInterface) bool

	// Scenario 4

	// return nominal maximum active (real) power the Controllable System is
	// able to consume according to the contract (EMS), device label or data sheet.
	//
	// parameters:
	//   - entity: the entity of the e.g. EVSE
	ConsumptionNominalMax(entity spineapi.EntityRemoteInterface) (float64, error)
}
