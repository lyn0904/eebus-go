package evsoc

import (
	"github.com/lyn0904/eebus-go/api"
	ucapi "github.com/lyn0904/eebus-go/usecases/api"
	usecase "github.com/lyn0904/eebus-go/usecases/usecase"
	spineapi "github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/spine"
)

type EVSOC struct {
	*usecase.UseCaseBase
}

var _ ucapi.CemEVSOCInterface = (*EVSOC)(nil)

func NewEVSOC(localEntity spineapi.EntityLocalInterface, eventCB api.EntityEventCallback) *EVSOC {
	validActorTypes := []model.UseCaseActorType{
		model.UseCaseActorTypeEV,
	}
	validEntityTypes := []model.EntityTypeType{
		model.EntityTypeTypeEV,
	}
	useCaseScenarios := []api.UseCaseScenario{
		{
			Scenario:       model.UseCaseScenarioSupportType(1),
			Mandatory:      true,
			ServerFeatures: []model.FeatureTypeType{model.FeatureTypeTypeMeasurement},
		},
	}

	usecase := usecase.NewUseCaseBase(
		localEntity,
		model.UseCaseActorTypeCEM,
		model.UseCaseNameTypeEVStateOfCharge,
		"1.0.0",
		"RC1",
		useCaseScenarios,
		eventCB,
		UseCaseSupportUpdate,
		validActorTypes,
		validEntityTypes,
	)

	uc := &EVSOC{
		UseCaseBase: usecase,
	}

	_ = spine.Events.Subscribe(uc)

	return uc
}

func (e *EVSOC) AddFeatures() {
	// client features
	var clientFeatures = []model.FeatureTypeType{
		model.FeatureTypeTypeElectricalConnection,
		model.FeatureTypeTypeMeasurement,
	}
	for _, feature := range clientFeatures {
		_ = e.LocalEntity.GetOrAddFeature(feature, model.RoleTypeClient)
	}
}

func (e *EVSOC) UpdateUseCaseAvailability(available bool) {
	e.LocalEntity.SetUseCaseAvailability(model.UseCaseActorTypeCEM, e.UseCaseName, available)
}
