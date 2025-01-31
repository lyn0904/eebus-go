package evcem

import (
	"github.com/lyn0904/eebus-go/api"
	ucapi "github.com/lyn0904/eebus-go/usecases/api"
	usecase "github.com/lyn0904/eebus-go/usecases/usecase"
	spineapi "github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/spine"
)

type EVCEM struct {
	*usecase.UseCaseBase

	service api.ServiceInterface
}

var _ ucapi.CemEVCEMInterface = (*EVCEM)(nil)

func NewEVCEM(service api.ServiceInterface, localEntity spineapi.EntityLocalInterface, eventCB api.EntityEventCallback) *EVCEM {
	validActorTypes := []model.UseCaseActorType{
		model.UseCaseActorTypeEV,
	}
	validEntityTypes := []model.EntityTypeType{
		model.EntityTypeTypeEV,
	}
	useCaseScenarios := []api.UseCaseScenario{
		{
			Scenario: model.UseCaseScenarioSupportType(1),
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeElectricalConnection,
				model.FeatureTypeTypeMeasurement,
			},
		},
		{
			Scenario: model.UseCaseScenarioSupportType(2),
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeElectricalConnection,
				model.FeatureTypeTypeMeasurement,
			},
		},
		{
			Scenario: model.UseCaseScenarioSupportType(3),
			ServerFeatures: []model.FeatureTypeType{
				model.FeatureTypeTypeElectricalConnection,
				model.FeatureTypeTypeMeasurement,
			},
		},
	}

	usecase := usecase.NewUseCaseBase(
		localEntity,
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeMeasurementOfElectricityDuringEVCharging,
		"1.0.1",
		"release",
		useCaseScenarios,
		eventCB,
		UseCaseSupportUpdate,
		validActorTypes,
		validEntityTypes)

	uc := &EVCEM{
		UseCaseBase: usecase,
		service:     service,
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

func (e *EVCEM) AddFeatures() {
	// client features
	var clientFeatures = []model.FeatureTypeType{
		model.FeatureTypeTypeElectricalConnection,
		model.FeatureTypeTypeMeasurement,
	}
	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}
}
