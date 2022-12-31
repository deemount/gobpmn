package examples

import (
	"testing"

	"github.com/deemount/gobpmn/models/core"
)

type Processes struct {
	model []model
}

func (p *Processes) AddModel(m model) {
	p.model = append(p.model, m)
}

func TestNewModel(t *testing.T) {

	def := new(core.Definitions)

	m := NewModel(def)
	t.Logf("NewModel(tDef)=%+v", &m)

}

func TestCreate(t *testing.T) {

	def := new(core.Definitions)

	m := NewModel(def)
	c := m.Create()
	t.Logf("Create=%+v", c)

}

func TestAddModel(t *testing.T) {

	var p Processes

	def := new(core.Definitions)

	m := NewModel(def)
	c := m.Create()
	p.AddModel(c)

	m = NewModel(def)
	c = m.Create()
	p.AddModel(c)

	t.Logf("Processes=%+v", p)

}
