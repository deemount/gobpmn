package repository

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

// BPMNF ...
type BPMNF struct {

	// models
	Def models.Definitions

	// string pointer
	DefHash     *string
	ProcHash    *string
	FlowHash    *string
	EventHash   *string
	CVersionTag *string
	Name        *string
	FormKey     *string

	// int64 pointer
	N       *int64
	Counter *int64
}

// NewBPMNF ...
func NewBPMNF() BPMNF {

	files, _ := ioutil.ReadDir("files/")
	var n int64 = 1
	var counter int64 = int64(len(files))

	var defHash string = utils.GenerateHash()
	var procHash string = utils.GenerateHash()
	var flowHash string = utils.GenerateHash()
	var eventHash string = utils.GenerateHash()
	var cVersionTag string = ""
	var name string = ""
	var formKey string = "create"

	if counter == 0 {
		counter++
	}
	return BPMNF{
		DefHash:     &defHash,
		ProcHash:    &procHash,
		FlowHash:    &flowHash,
		EventHash:   &eventHash,
		CVersionTag: &cVersionTag,
		Name:        &name,
		N:           &n,
		FormKey:     &formKey,
		Counter:     &counter,
	}
}

/*
	Definitions Management
*/

// setDefinitions ...
func (bpm *BPMNF) setDefinitions() {

	// set namespaces
	bpm.Def.SetBpmn()
	//d.SetXSD()
	bpm.Def.SetBpmndi()

	// set specification
	bpm.Def.SetDC()
	bpm.Def.SetBioc()

	//
	bpm.Def.SetDefinitionsID(*bpm.DefHash)
	bpm.Def.SetTargetNamespace()

	// set exporter
	bpm.Def.SetExporter()
	bpm.Def.SetExporterVersion()

}

// setProcess
func (bpm *BPMNF) setProcess() {

	// set process
	bpm.Def.Proc.SetID(*bpm.ProcHash)
	bpm.Def.Proc.SetName(*bpm.Name)
	bpm.Def.Proc.SetIsExecutable()
	bpm.Def.Proc.SetCamundaVersionTag(*bpm.CVersionTag)

	// set diagram
	bpm.Def.Diagram.SetID(*bpm.N)

	// set start event
	bpm.Def.Proc.StartEvent.SetID(*bpm.N)
	bpm.Def.Proc.StartEvent.SetCamundaFormKey(*bpm.FormKey)
	bpm.Def.Proc.StartEvent.Outgoing.SetFlow(*bpm.FlowHash)

	// set intermediate throw event

	bpm.Def.Proc.IntermediateThrowEvent.SetID(*bpm.EventHash)
	bpm.Def.Proc.IntermediateThrowEvent.Incoming.SetFlow(*bpm.FlowHash)

	// set sequence flow (start event)
	bpm.Def.Proc.SequenceFlow.SetID(*bpm.FlowHash)
	bpm.Def.Proc.SequenceFlow.SetStartEvent(*bpm.N)
	bpm.Def.Proc.SequenceFlow.SetEventRef(*bpm.EventHash)

}

// setDiagram
func (bpm *BPMNF) setDiagram() {

	// set plane
	bpm.Def.Diagram.Plane.SetID(*bpm.N)
	bpm.Def.Diagram.Plane.SetElement(*bpm.ProcHash)

	// set edge
	bpm.Def.Diagram.Plane.Edge.SetID(*bpm.FlowHash)
	bpm.Def.Diagram.Plane.Edge.SetElement(*bpm.FlowHash)
	bpm.Def.Diagram.Plane.Edge.Waypoint = append(bpm.Def.Diagram.Plane.Edge.Waypoint,
		models.Waypoint{X: "190", Y: "100"},
		models.Waypoint{X: "250", Y: "100"})

	// set shape
	bpm.Def.Diagram.Plane.Shape =
		[]models.Shape{
			{ // start event
				ID:      fmt.Sprintf("_BPMNShape_StartEvent_%d", *bpm.N+1),
				Element: fmt.Sprintf("StartEvent_%d", bpm.N),
				Stroke:  "rgb(67, 160, 71)",
				Fill:    "rgb(200, 230, 201)",
				Bounds:  models.Bounds{X: "150", Y: "80", Width: "40", Height: "40"},
			},
			{ // first state
				ID:      fmt.Sprintf("Event_%s_di", *bpm.EventHash),
				Element: fmt.Sprintf("Event_%s", *bpm.EventHash),
				Stroke:  "black",
				Fill:    "white",
				Bounds:  models.Bounds{X: "250", Y: "80", Width: "40", Height: "40"},
			}}

}

// Set sets the definitions, process and diagram elements in the instance
func (bpm *BPMNF) Set() {

	bpm.setDefinitions()
	bpm.setProcess()
	bpm.setDiagram()

}

// Create ...
func (bpm *BPMNF) Create() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(bpm.Def, " ", "  ")

	// create file
	f, err := os.Create("files/diagram_" + fmt.Sprintf("%d", *bpm.Counter) + ".bpmn")
	if err != nil {
		return err
	}
	defer f.Close()

	// add xml header
	w := []byte(fmt.Sprintf("%v", xml.Header+string(b)))

	// write bytes to file
	_, err = f.Write(w)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}

	return nil

}
