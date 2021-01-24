package repository

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deemount/gobpmn/models"
)

const ()

// BPMNF ...
type BPMNF struct {
	Def models.Definitions
}

// NewBPMNF ...
func NewBPMNF() *BPMNF {
	return &BPMNF{}
}

// Create ...
func (bpm BPMNF) Create() string {

	d := bpm.Def

	// set namespaces
	d.SetBpmn()
	d.SetXSD()
	d.SetBpmndi()

	//
	d.SetDefinitionsID()
	d.SetTargetNamespace()

	// set exporter
	d.SetExporter()
	d.SetExporterVersion()

	// set process
	d.Proc.SetID()
	d.Proc.SetIsExecutable()
	d.Proc.SetCamundaVersionTag()

	// set diagram
	d.Diagram.SetID()

	// set plane
	d.Diagram.Plane.SetID()
	d.Diagram.Plane.SetElement()

	log.Printf("%#v", d)

	out, _ := xml.MarshalIndent(&d, " ", "  ")

	return fmt.Sprintf("%v", xml.Header+string(out))

}
