package repository

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"

	"github.com/deemount/gobpmn/models"
)

// Options ...
type Options struct {

	// integer values
	N       int64
	Counter int64

	// hash values
	DefHash         string
	ProcHash        string
	FlowHash        string
	EventHash       string
	ParticipantHash string
	CollaboHash     string

	// other values
	CVersionTag string
	Name        string
	FormKey     string

	// bool
	HasCollab     bool
	HasStartEvent bool
	HasFirstState bool
}

// BPMNF ...
type BPMNF struct {
	// options
	Opts Options
	// models
	Def models.Definitions
}

type BPMNFOption func(o Options) Options

// NewBPMNF ...
func NewBPMNF(opt ...BPMNFOption) BPMNF {
	opts := Options{}
	for _, o := range opt {
		opts = o(opts)
	}
	// count up
	if opts.Counter == 0 {
		opts.Counter++
	}
	return BPMNF{Opts: opts}
}

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
	bpm.Def.SetDefinitionsID(bpm.Opts.DefHash)
	bpm.Def.SetTargetNamespace()

	// set exporter
	bpm.Def.SetExporter()
	bpm.Def.SetExporterVersion()

}

// setCollaboration ...
func (bpm *BPMNF) setCollaboration() {

	bpm.Def.Collab.SetID(bpm.Opts.CollaboHash)
	bpm.Def.Collab.Participant.SetID(bpm.Opts.ParticipantHash)
	bpm.Def.Collab.Participant.SetName("Person 1")
	bpm.Def.Collab.Participant.SetProcessRef(bpm.Opts.ProcHash)

}

// setProcess
func (bpm *BPMNF) setProcess() {

	// set process
	bpm.Def.Proc.SetID(bpm.Opts.ProcHash)
	bpm.Def.Proc.SetName(bpm.Opts.Name)
	bpm.Def.Proc.SetIsExecutable()
	bpm.Def.Proc.SetCamundaVersionTag(bpm.Opts.CVersionTag)

	// set diagram
	bpm.Def.Diagram.SetID(bpm.Opts.N)

	// set start event
	if bpm.Opts.HasStartEvent {
		bpm.Def.Proc.StartEvent.SetID(bpm.Opts.N)
		bpm.Def.Proc.StartEvent.SetCamundaFormKey(bpm.Opts.FormKey)
		bpm.Def.Proc.StartEvent.Outgoing.SetFlow(bpm.Opts.FlowHash)

		// set intermediate throw event
		bpm.Def.Proc.IntermediateThrowEvent.SetID(bpm.Opts.EventHash)
		bpm.Def.Proc.IntermediateThrowEvent.Incoming.SetFlow(bpm.Opts.FlowHash)

		// set sequence flow (start event)
		bpm.Def.Proc.SequenceFlow.SetID(bpm.Opts.FlowHash)
		bpm.Def.Proc.SequenceFlow.SetStartEvent(bpm.Opts.N)
		bpm.Def.Proc.SequenceFlow.SetEventRef(bpm.Opts.EventHash)
	}

}

// setDiagram
func (bpm *BPMNF) setDiagram() {

	// assign
	var collab models.Shape
	var startEvent models.Shape
	var firstState models.Shape

	// set plane
	bpm.Def.Diagram.Plane.SetID(bpm.Opts.N)
	if bpm.Opts.HasCollab {
		bpm.Def.Diagram.Plane.SetCollaborationElement(bpm.Opts.CollaboHash)
	} else {
		bpm.Def.Diagram.Plane.SetProcessElement(bpm.Opts.ProcHash)
	}

	// set edge
	bpm.Def.Diagram.Plane.Edge.SetID(bpm.Opts.FlowHash)
	bpm.Def.Diagram.Plane.Edge.SetElement(bpm.Opts.FlowHash)
	bpm.Def.Diagram.Plane.Edge.Waypoint = append(bpm.Def.Diagram.Plane.Edge.Waypoint,
		models.Waypoint{X: "190", Y: "100"},
		models.Waypoint{X: "250", Y: "100"})

	if bpm.Opts.HasCollab {
		collab = models.Shape{
			ID:           fmt.Sprintf("Participant_%s_di", bpm.Opts.ParticipantHash),
			Element:      fmt.Sprintf("Participant_%s", bpm.Opts.ParticipantHash),
			IsHorizontal: strconv.FormatBool(true),
			Bounds:       models.Bounds{X: "100", Y: "10", Width: "600", Height: "300"},
		}
	}

	if bpm.Opts.HasStartEvent {
		startEvent = models.Shape{
			ID:      fmt.Sprintf("_BPMNShape_StartEvent_%d", bpm.Opts.N+1),
			Element: fmt.Sprintf("StartEvent_%d", bpm.Opts.N),
			Stroke:  "rgb(67, 160, 71)",
			Fill:    "rgb(200, 230, 201)",
			Bounds:  models.Bounds{X: "150", Y: "80", Width: "40", Height: "40"},
		}
		if bpm.Opts.HasFirstState {
			firstState = models.Shape{
				ID:      fmt.Sprintf("Event_%s_di", bpm.Opts.EventHash),
				Element: fmt.Sprintf("Event_%s", bpm.Opts.EventHash),
				Stroke:  "black",
				Fill:    "white",
				Bounds:  models.Bounds{X: "250", Y: "80", Width: "40", Height: "40"},
			}
		}
	}

	// set shape
	bpm.Def.Diagram.Plane.Shape = []models.Shape{collab, startEvent, firstState}

}

// Set sets the definitions, process and diagram elements in the instance
func (bpm *BPMNF) Set() {

	if bpm.Opts.HasCollab {
		bpm.setCollaboration()
	}
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
	f, err := os.Create("files/diagram_" + fmt.Sprintf("%d", bpm.Opts.Counter) + ".bpmn")
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
