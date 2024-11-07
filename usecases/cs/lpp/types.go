package lpp

import "github.com/lyn0904/eebus-go/api"

const (
	// Update of the list of remote entities supporting the Use Case
	//
	// Use `RemoteEntities` to get the current data
	UseCaseSupportUpdate api.EventType = "cs-lpp-UseCaseSupportUpdate"

	// Load control obligation limit data update received
	//
	// Use `ProductionLimit` to get the current data
	//
	// Use Case LPC, Scenario 1
	DataUpdateLimit api.EventType = "cs-lpp-DataUpdateLimit"

	// An incoming load control obligation limit needs to be approved or denied
	//
	// Use `PendingProductionLimits` to get the currently pending write approval requests
	// and invoke `ApproveOrDenyProductionLimit` for each
	//
	// Use Case LPC, Scenario 1
	WriteApprovalRequired api.EventType = "cs-lpp-WriteApprovalRequired"

	// Failsafe limit for the produced active (real) power of the
	// Controllable System data update received
	//
	// Use `FailsafeProductionActivePowerLimit` to get the current data
	//
	// Use Case LPC, Scenario 2
	DataUpdateFailsafeProductionActivePowerLimit api.EventType = "cs-lpp-DataUpdateFailsafeProductionActivePowerLimit"

	// Minimum time the Controllable System remains in "failsafe state" unless conditions
	// specified in this Use Case permit leaving the "failsafe state" data update received
	//
	// Use `FailsafeDurationMinimum` to get the current data
	//
	// Use Case LPC, Scenario 2
	DataUpdateFailsafeDurationMinimum api.EventType = "cs-lpp-DataUpdateFailsafeDurationMinimum"

	// Indicates a notify heartbeat event the application should care of.
	// E.g. going into or out of the Failsafe state
	//
	// Use Case LPP, Scenario 3
	DataUpdateHeartbeat api.EventType = "uclppserver-DataUpdateHeartbeat"
)
