package mpc

import (
	"github.com/lyn0904/eebus-go/api"
	ucapi "github.com/lyn0904/eebus-go/usecases/api"
	usecase "github.com/lyn0904/eebus-go/usecases/usecase"
	spineapi "github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/spine"
)

type MPC struct {
	*usecase.UseCaseBase
}

var _ ucapi.MaMPCInterface = (*MPC)(nil)

func NewMPC(localEntity spineapi.EntityLocalInterface, eventCB api.EntityEventCallback) *MPC {
	validActorTypes := []model.UseCaseActorType{model.UseCaseActorTypeMonitoredUnit}
	validEntityTypes := []model.EntityTypeType{
		model.EntityTypeTypeCompressor,
		model.EntityTypeTypeElectricalImmersionHeater,
		model.EntityTypeTypeEVSE,
		model.EntityTypeTypeHeatPumpAppliance,
		model.EntityTypeTypeInverter,
		model.EntityTypeTypeSmartEnergyAppliance,
		model.EntityTypeTypeSubMeterElectricity,
	}
	useCaseScenarios := []api.UseCaseScenario{
		{
			Scenario:  model.UseCaseScenarioSupportType(1),
			Mandatory: true,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeElectricalConnection,
				model.FeatureTypeTypeMeasurement,
			},
		},
		{
			Scenario:  model.UseCaseScenarioSupportType(2),
			Mandatory: false,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeElectricalConnection,
				model.FeatureTypeTypeMeasurement,
			},
		},
		{
			Scenario:  model.UseCaseScenarioSupportType(3),
			Mandatory: false,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeElectricalConnection,
				model.FeatureTypeTypeMeasurement,
			},
		},
		{
			Scenario:  model.UseCaseScenarioSupportType(4),
			Mandatory: false,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeElectricalConnection,
				model.FeatureTypeTypeMeasurement,
			},
		},
		{
			Scenario:  model.UseCaseScenarioSupportType(5),
			Mandatory: false,
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeElectricalConnection,
				model.FeatureTypeTypeMeasurement,
			},
		},
	}

	usecase := usecase.NewUseCaseBase(
		localEntity,
		model.UseCaseActorTypeMonitoringAppliance,
		model.UseCaseNameTypeMonitoringOfPowerConsumption,
		"1.0.0",
		"release",
		useCaseScenarios,
		eventCB,
		UseCaseSupportUpdate,
		validActorTypes,
		validEntityTypes)

	uc := &MPC{
		UseCaseBase: usecase,
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

func (e *MPC) AddFeatures() {
	// client features
	var clientFeatures = []model.FeatureTypeType{
		model.FeatureTypeTypeElectricalConnection,
		model.FeatureTypeTypeMeasurement,
	}
	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}
}
