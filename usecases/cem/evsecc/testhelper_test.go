package evsecc

import (
	"fmt"
	"testing"
	"time"

	"github.com/lyn0904/eebus-go/api"
	"github.com/lyn0904/eebus-go/mocks"
	"github.com/lyn0904/eebus-go/service"
	shipapi "github.com/lyn0904/ship-go/api"
	"github.com/lyn0904/ship-go/cert"
	shipmocks "github.com/lyn0904/ship-go/mocks"
	spineapi "github.com/lyn0904/spine-go/api"
	spinemocks "github.com/lyn0904/spine-go/mocks"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/spine"
	"github.com/lyn0904/spine-go/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestCemEVSECCSuite(t *testing.T) {
	suite.Run(t, new(CemEVSECCSuite))
}

type CemEVSECCSuite struct {
	suite.Suite

	sut *EVSECC

	service api.ServiceInterface

	remoteDevice     spineapi.DeviceRemoteInterface
	mockRemoteEntity *spinemocks.EntityRemoteInterface
	evseEntity       spineapi.EntityRemoteInterface

	eventCalled bool
}

func (s *CemEVSECCSuite) Event(ski string, device spineapi.DeviceRemoteInterface, entity spineapi.EntityRemoteInterface, event api.EventType) {
	s.eventCalled = true
}

func (s *CemEVSECCSuite) BeforeTest(suiteName, testName string) {
	s.eventCalled = false
	cert, _ := cert.CreateCertificate("test", "test", "DE", "test")
	configuration, _ := api.NewConfiguration(
		"test", "test", "test", "test",
		[]shipapi.DeviceCategoryType{shipapi.DeviceCategoryTypeEnergyManagementSystem},
		model.DeviceTypeTypeEnergyManagementSystem,
		[]model.EntityTypeType{model.EntityTypeTypeCEM},
		9999, cert, time.Second*4)

	serviceHandler := mocks.NewServiceReaderInterface(s.T())
	serviceHandler.EXPECT().ServicePairingDetailUpdate(mock.Anything, mock.Anything).Return().Maybe()

	s.service = service.NewService(configuration, serviceHandler)
	_ = s.service.Setup()

	mockRemoteDevice := spinemocks.NewDeviceRemoteInterface(s.T())
	s.mockRemoteEntity = spinemocks.NewEntityRemoteInterface(s.T())
	mockRemoteFeature := spinemocks.NewFeatureRemoteInterface(s.T())
	mockRemoteDevice.EXPECT().FeatureByEntityTypeAndRole(mock.Anything, mock.Anything, mock.Anything).Return(mockRemoteFeature).Maybe()
	mockRemoteDevice.EXPECT().Ski().Return(remoteSki).Maybe()
	s.mockRemoteEntity.EXPECT().Device().Return(mockRemoteDevice).Maybe()
	s.mockRemoteEntity.EXPECT().EntityType().Return(mock.Anything).Maybe()
	entityAddress := &model.EntityAddressType{}
	s.mockRemoteEntity.EXPECT().Address().Return(entityAddress).Maybe()
	mockRemoteFeature.EXPECT().DataCopy(mock.Anything).Return(mock.Anything).Maybe()
	mockRemoteFeature.EXPECT().Address().Return(&model.FeatureAddressType{}).Maybe()
	mockRemoteFeature.EXPECT().Operations().Return(nil).Maybe()

	localEntity := s.service.LocalDevice().EntityForType(model.EntityTypeTypeCEM)
	s.sut = NewEVSECC(localEntity, s.Event)
	s.sut.AddFeatures()
	s.sut.AddUseCase()

	var entities []spineapi.EntityRemoteInterface
	s.remoteDevice, entities = setupDevices(s.service, s.T())
	s.evseEntity = entities[0]
}

const remoteSki string = "testremoteski"

func setupDevices(
	eebusService api.ServiceInterface, t *testing.T) (
	spineapi.DeviceRemoteInterface,
	[]spineapi.EntityRemoteInterface) {
	localDevice := eebusService.LocalDevice()

	writeHandler := shipmocks.NewShipConnectionDataWriterInterface(t)
	writeHandler.EXPECT().WriteShipMessageWithPayload(mock.Anything).Return().Maybe()
	sender := spine.NewSender(writeHandler)
	remoteDevice := spine.NewDeviceRemote(localDevice, remoteSki, sender)

	remoteDeviceName := "remote"

	var remoteFeatures = []struct {
		featureType   model.FeatureTypeType
		supportedFcts []model.FunctionType
	}{
		{model.FeatureTypeTypeDeviceClassification,
			[]model.FunctionType{
				model.FunctionTypeDeviceClassificationManufacturerData,
			},
		},
		{model.FeatureTypeTypeDeviceDiagnosis,
			[]model.FunctionType{
				model.FunctionTypeDeviceDiagnosisStateData,
			},
		},
	}

	var featureInformations []model.NodeManagementDetailedDiscoveryFeatureInformationType
	for index, feature := range remoteFeatures {
		supportedFcts := []model.FunctionPropertyType{}
		for _, fct := range feature.supportedFcts {
			supportedFct := model.FunctionPropertyType{
				Function: util.Ptr(fct),
				PossibleOperations: &model.PossibleOperationsType{
					Read: &model.PossibleOperationsReadType{},
				},
			}
			supportedFcts = append(supportedFcts, supportedFct)
		}

		featureInformation := model.NodeManagementDetailedDiscoveryFeatureInformationType{
			Description: &model.NetworkManagementFeatureDescriptionDataType{
				FeatureAddress: &model.FeatureAddressType{
					Device:  util.Ptr(model.AddressDeviceType(remoteDeviceName)),
					Entity:  []model.AddressEntityType{1},
					Feature: util.Ptr(model.AddressFeatureType(index)),
				},
				FeatureType:       util.Ptr(feature.featureType),
				Role:              util.Ptr(model.RoleTypeServer),
				SupportedFunction: supportedFcts,
			},
		}
		featureInformations = append(featureInformations, featureInformation)
	}

	detailedData := &model.NodeManagementDetailedDiscoveryDataType{
		DeviceInformation: &model.NodeManagementDetailedDiscoveryDeviceInformationType{
			Description: &model.NetworkManagementDeviceDescriptionDataType{
				DeviceAddress: &model.DeviceAddressType{
					Device: util.Ptr(model.AddressDeviceType(remoteDeviceName)),
				},
			},
		},
		EntityInformation: []model.NodeManagementDetailedDiscoveryEntityInformationType{
			{
				Description: &model.NetworkManagementEntityDescriptionDataType{
					EntityAddress: &model.EntityAddressType{
						Device: util.Ptr(model.AddressDeviceType(remoteDeviceName)),
						Entity: []model.AddressEntityType{1},
					},
					EntityType: util.Ptr(model.EntityTypeTypeEVSE),
				},
			},
			{
				Description: &model.NetworkManagementEntityDescriptionDataType{
					EntityAddress: &model.EntityAddressType{
						Device: util.Ptr(model.AddressDeviceType(remoteDeviceName)),
						Entity: []model.AddressEntityType{1, 1},
					},
					EntityType: util.Ptr(model.EntityTypeTypeEV),
				},
			},
		},
		FeatureInformation: featureInformations,
	}

	entities, err := remoteDevice.AddEntityAndFeatures(true, detailedData)
	if err != nil {
		fmt.Println(err)
	}
	remoteDevice.UpdateDevice(detailedData.DeviceInformation.Description)

	for _, entity := range entities {
		entity.UpdateDeviceAddress(*remoteDevice.Address())
	}

	localDevice.AddRemoteDeviceForSki(remoteSki, remoteDevice)

	return remoteDevice, entities
}
