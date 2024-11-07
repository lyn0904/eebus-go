package api

import "github.com/lyn0904/spine-go/model"

//go:generate mockery

var PhaseNameMapping = []model.ElectricalConnectionPhaseNameType{model.ElectricalConnectionPhaseNameTypeA, model.ElectricalConnectionPhaseNameTypeB, model.ElectricalConnectionPhaseNameTypeC}
