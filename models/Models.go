package models

import (
	"log"

	"github.com/deemount/gobpmn/models/bpmn/core"
)

type IFModels interface {
	core.DefinitionsRepository
}

type Models struct {
	IFModels
}

func NewModels() IFModels {
	return &Models{
<<<<<<< HEAD
		core.NewDefinitions(&Models{})}
=======
		core.NewDefinitions()}
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

func (m *Models) Test() {

	log.Printf("%+v", m)

}
