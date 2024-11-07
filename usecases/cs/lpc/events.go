package lpc

import (
	"github.com/enbility/eebus-go/features/client"
	"github.com/enbility/eebus-go/features/server"
	internal "github.com/enbility/eebus-go/usecases/internal"
	"github.com/lyn0904/ship-go/logging"
	spineapi "github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/util"
)

// handle SPINE events
func (e *LPC) HandleEvent(payload spineapi.EventPayload) {
	if internal.IsDeviceConnected(payload) {
		e.deviceConnected(payload)
		return
	}

	if !e.IsCompatibleEntityType(payload.Entity) {
		return
	}

	// did we receive a binding to the loadControl server and the
	// heartbeatWorkaround is required?
	if payload.EventType == spineapi.EventTypeBindingChange &&
		payload.ChangeType == spineapi.ElementChangeAdd &&
		payload.LocalFeature != nil &&
		payload.LocalFeature.Type() == model.FeatureTypeTypeLoadControl &&
		payload.LocalFeature.Role() == model.RoleTypeServer {
		e.subscribeHeartbeatWorkaround(payload)
		return
	}

	if internal.IsHeartbeat(payload) && e.EventCB != nil {
		e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateHeartbeat)
		return
	}

	if payload.EventType != spineapi.EventTypeDataChange ||
		payload.ChangeType != spineapi.ElementChangeUpdate ||
		payload.CmdClassifier == nil ||
		*payload.CmdClassifier != model.CmdClassifierTypeWrite {
		return
	}

	switch payload.Data.(type) {
	case *model.LoadControlLimitListDataType:
		serverF := e.LocalEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl, model.RoleTypeServer)

		if payload.Function != model.FunctionTypeLoadControlLimitListData ||
			payload.LocalFeature != serverF {
			return
		}

		e.loadControlLimitDataUpdate(payload)
	case *model.DeviceConfigurationKeyValueListDataType:
		serverF := e.LocalEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeDeviceConfiguration, model.RoleTypeServer)

		if payload.Function != model.FunctionTypeDeviceConfigurationKeyValueListData ||
			payload.LocalFeature != serverF {
			return
		}

		e.configurationDataUpdate(payload)
	}
}

// a remote device was connected and we know its entities
func (e *LPC) deviceConnected(payload spineapi.EventPayload) {
	if payload.Device == nil {
		return
	}

	// check if there is a DeviceDiagnosis server on one or more entities
	remoteDevice := payload.Device

	var deviceDiagEntities []spineapi.EntityRemoteInterface

	entities := remoteDevice.Entities()
	for _, entity := range entities {
		if !e.IsCompatibleEntityType(entity) {
			continue
		}

		deviceDiagF := entity.FeatureOfTypeAndRole(model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
		if deviceDiagF == nil {
			continue
		}

		deviceDiagEntities = append(deviceDiagEntities, entity)
	}

	logging.Log().Debug("cs-lpc:", len(deviceDiagEntities), "DeviceDiagnosis Server found")

	// the remote device does not have a DeviceDiagnosis Server, which it should
	if len(deviceDiagEntities) == 0 {
		return
	}

	// we only found one matching entity, as it should be, subscribe
	if len(deviceDiagEntities) == 1 {
		if localDeviceDiag, err := client.NewDeviceDiagnosis(e.LocalEntity, deviceDiagEntities[0]); err == nil {
			e.heartbeatDiag = localDeviceDiag
			if !localDeviceDiag.HasSubscription() {
				if _, err := localDeviceDiag.Subscribe(); err != nil {
					logging.Log().Debug(err)
				}
			}

			if _, err := localDeviceDiag.RequestHeartbeat(); err != nil {
				logging.Log().Debug(err)
			}
		}

		return
	}

	// we found more than one matching entity, this is not good
	// according to KEO the subscription should be done on the entity that requests a binding to
	// the local loadControlLimit server feature
	e.heartbeatKeoWorkaround = true
}

// subscribe to the DeviceDiagnosis Server of the entity that created a binding
func (e *LPC) subscribeHeartbeatWorkaround(payload spineapi.EventPayload) {
	// is the workaround is needed?
	if e.heartbeatKeoWorkaround {
		if localDeviceDiag, err := client.NewDeviceDiagnosis(e.LocalEntity, payload.Entity); err == nil {
			e.heartbeatDiag = localDeviceDiag
			if !localDeviceDiag.HasSubscription() {
				if _, err := localDeviceDiag.Subscribe(); err != nil {
					logging.Log().Debug(err)
				}
			}

			if _, err := localDeviceDiag.RequestHeartbeat(); err != nil {
				logging.Log().Debug(err)
			}
		}
	}
}

// the load control limit data was updated
func (e *LPC) loadControlLimitDataUpdate(payload spineapi.EventPayload) {
	if lc, err := server.NewLoadControl(e.LocalEntity); err == nil {
		filter := model.LoadControlLimitDescriptionDataType{
			LimitType:      util.Ptr(model.LoadControlLimitTypeTypeSignDependentAbsValueLimit),
			LimitCategory:  util.Ptr(model.LoadControlCategoryTypeObligation),
			LimitDirection: util.Ptr(model.EnergyDirectionTypeConsume),
			ScopeType:      util.Ptr(model.ScopeTypeTypeActivePowerLimit),
		}
		if lc.CheckEventPayloadDataForFilter(payload.Data, filter) && e.EventCB != nil {
			e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateLimit)
		}
	}
}

// the configuration key data of an SMGW was updated
func (e *LPC) configurationDataUpdate(payload spineapi.EventPayload) {
	if dc, err := server.NewDeviceConfiguration(e.LocalEntity); err == nil {
		filter := model.DeviceConfigurationKeyValueDescriptionDataType{
			KeyName: util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeConsumptionActivePowerLimit),
		}
		if dc.CheckEventPayloadDataForFilter(payload.Data, filter) && e.EventCB != nil {
			e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateFailsafeConsumptionActivePowerLimit)
		}
		filter = model.DeviceConfigurationKeyValueDescriptionDataType{
			KeyName: util.Ptr(model.DeviceConfigurationKeyNameTypeFailsafeDurationMinimum),
		}
		if dc.CheckEventPayloadDataForFilter(payload.Data, filter) && e.EventCB != nil {
			e.EventCB(payload.Ski, payload.Device, payload.Entity, DataUpdateFailsafeDurationMinimum)
		}
	}
}
