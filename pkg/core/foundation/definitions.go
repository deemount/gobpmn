package foundation

import (
	"crypto/rand"
	"encoding/xml"
	"fmt"
	"hash/fnv"

	"github.com/deemount/gobpmn/pkg/core/infrastructure"
)

var (
	schemaOMGBpmnModel = "http://www.omg.org/spec/BPMN/20100524/MODEL"
	schemaBpmnDI       = "http://www.omg.org/spec/BPMN/20100524/DI"
	schemaOMGDC        = "http://www.omg.org/spec/DD/20100524/DC"
	schemaBpmnIOSchema = "http://bpmn.io/schema/bpmn"
)

type (

	// Interface ...
	DefinitionsRepository interface {
		SetID(typ string, suffix any)
		GetID() *string
		SetBpmn()
		SetBpmnDI()
		SetDC()
		SetTargetNamespace()
		SetProcess(num int)
		GetProcess(num int) *infrastructure.Process
		SetCollaboration()
		GetCollaboration() *infrastructure.Collaboration
		SetMainElements(num int)
		SetDefaultAttributes()
	}

	// Definitions ...
	Definitions struct {
		XMLName         xml.Name                       `xml:"bpmn:definitions" json:"-"`
		Bpmn            string                         `xml:"xmlns:bpmn,attr" json:"-"`
		BpmnDI          string                         `xml:"xmlns:bpmndi,attr" json:"-"`
		DC              string                         `xml:"xmlns:dc,attr,omitempty" json:"-"`
		ID              string                         `xml:"id,attr" json:"id"`
		TargetNamespace string                         `xml:"targetNamespace,attr" json:"-"`
		Collaboration   []infrastructure.Collaboration `xml:"bpmn:collaboration,omitempty" json:"collaboration,omitempty"`
		Process         []infrastructure.Process       `xml:"bpmn:process,omitempty" json:"process"`
		Diagram         []infrastructure.Diagram       `xml:"bpmndi:BPMNDiagram,omitempty" json:"diagram,omitempty"`
	}

	// TDefinitions ...
	TDefinitions struct {
		XMLName       xml.Name                        `xml:"definitions" json:"-"`
		ID            string                          `xml:"id,attr" json:"id"`
		Collaboration []infrastructure.TCollaboration `xml:"collaboration,omitempty" json:"collaboration,omitempty"`
		Process       []infrastructure.TProcess       `xml:"process,omitempty" json:"process"`
	}
)

func NewDefinitions() DefinitionsRepository {
	return &Definitions{}
}

// SetID ...
func (d *Definitions) SetID(typ string, suffix any) {
	switch typ {
	case "definitions":
		d.ID = fmt.Sprintf("Definitions_%v", suffix)
	case "id":
		d.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (d Definitions) GetID() *string {
	return &d.ID
}

// SetBpmn ...
func (d *Definitions) SetBpmn() {
	d.Bpmn = schemaOMGBpmnModel
}

// SetBpmnDI ...
func (d *Definitions) SetBpmnDI() {
	d.BpmnDI = schemaBpmnDI
}

// SetDC ...
func (d *Definitions) SetDC() {
	d.DC = schemaOMGDC
}

// SetTargetNamespace ...
func (d *Definitions) SetTargetNamespace() {
	d.TargetNamespace = schemaBpmnIOSchema
}

// SetCollaboration ...
func (d *Definitions) SetCollaboration() {
	d.Collaboration = make([]infrastructure.Collaboration, 1)
}

// GetCollaboration ...
func (d Definitions) GetCollaboration() *infrastructure.Collaboration {
	return &d.Collaboration[0]
}

// SetProcess ...
func (d *Definitions) SetProcess(num int) {
	if num == 0 {
		num = 1
	}
	d.Process = make([]infrastructure.Process, num)
}

// GetProcess ...
func (d Definitions) GetProcess(num int) *infrastructure.Process {
	return &d.Process[num]
}

// SetDiagram ...
func (d *Definitions) SetDiagram(num int) {
	d.Diagram = make([]infrastructure.Diagram, num)
}

// GetDiagram ...
func (d Definitions) GetDiagram(num int) *infrastructure.Diagram {
	return &d.Diagram[num]
}

// SetDefaultAttributes ...
func (d *Definitions) SetDefaultAttributes() {
	h := GenerateHash()
	d.SetBpmn()
	d.SetBpmnDI()
	d.SetDC()
	d.SetID("definitions", h)
	d.SetTargetNamespace()
}

// SetMainElements ...
func (d *Definitions) SetMainElements(num int) {
	d.SetProcess(num)
	d.SetDiagram(1)
}

// GenerateHash ...
func GenerateHash() string {

	n := 8
	b := make([]byte, n)
	h := fnv.New32a()
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)

	if _, err := h.Write([]byte(s)); err != nil {
		panic(err)
	}
	defer h.Reset()

	return fmt.Sprintf("%x", h.Sum(nil))

}
