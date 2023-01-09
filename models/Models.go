package models

import (
	"log"

	"github.com/deemount/gobpmn/models/core"
)

type IFModels interface {
	core.DefinitionsRepository
}

type Models struct {
	IFModels
}

func NewModels() IFModels {
	return &Models{
		core.NewDefinitions()}
}

func (m *Models) Test() {

	log.Printf("%+v", m)

}
