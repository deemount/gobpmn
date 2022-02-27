package repository

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

type Options struct {
	// integer values
	Counter int
}

// BPMNFRepository ...
type BPMNFRepository interface {
	Set()
	Create() error
}

// BPMNF ...
type BPMNF struct {
	// options
	Options Options
	Def     models.Definitions
}

type BPMNFOption func(o Options) Options

// NewBPMNF ...
func NewBPMNF(opt ...BPMNFOption) BPMNF {
	options := Options{}
	for _, o := range opt {
		options = o(options)
	}
	files, _ := ioutil.ReadDir("files/")
	options.Counter = len(files)
	// count up
	if options.Counter == 0 {
		options.Counter++
	}
	return BPMNF{Options: options}
}

// Set ...
func (bpm *BPMNF) Set() {

	//++

	/** set definitions **/
	// set namespaces
	def := &bpm.Def
	def.SetBpmn()
	def.SetBpmndi()
	def.SetDC()
	// set definitions id & target namespace
	defHash := utils.GenerateHash()
	def.SetID(defHash)
	def.SetTargetNamespace()
	// set exporter & version
	def.SetExporter()
	def.SetExporterVersion()

	//++

	/* set definition elements */
	def.SetCollaboration()
	def.SetProcess(1)
	def.Process[0].SetStartEvent(1)
	def.Process[0].SetTask(1)
	def.Process[0].SetEndEvent(1)
	def.Process[0].SetSequenceFlow(2)
	def.SetDiagram()
	def.Diagram[0].SetPlane()
	def.Diagram[0].Plane[0].SetShape(4)
	def.Diagram[0].Plane[0].SetEdge(2)

	/* set elements attributes */

	/** set collaboration attributes **/
	collabHash := utils.GenerateHash()
	def.Collaboration[0].SetID(collabHash)
	// set participant
	def.Collaboration[0].SetParticipant(1)
	participHash := utils.GenerateHash()
	def.Collaboration[0].Participant[0].SetID(participHash)
	// set process reference
	procHash := utils.GenerateHash()
	def.Collaboration[0].Participant[0].SetProcessRef(procHash)

	/*** set shape attributes (collaboration) ***/
	def.Diagram[0].Plane[0].Shape[0].SetID("participant", participHash)
	def.Diagram[0].Plane[0].Shape[0].SetElement("participant", participHash)
	def.Diagram[0].Plane[0].Shape[0].SetIsHorizontal(true)
	def.Diagram[0].Plane[0].Shape[0].SetBounds()
	def.Diagram[0].Plane[0].Shape[0].Bounds[0].SetCoordinates(129, 60)
	def.Diagram[0].Plane[0].Shape[0].Bounds[0].SetSize(600, 250)

	/** set process attributes **/
	// element
	def.Process[0].SetID(procHash)
	name := "Test"
	def.Process[0].SetName(name)
	isExecutable := true
	def.Process[0].SetIsExecutable(isExecutable)
	cVersionTag := "v0.1.0"
	def.Process[0].SetCamundaVersionTag(cVersionTag)

	/** set start event attributes **/
	// generics
	var stevN int64 = 1
	outFromStartEvent := utils.GenerateHash()
	def.Process[0].StartEvent[0].SetID(stevN)
	// Outgoing
	def.Process[0].StartEvent[0].SetOutgoing(1)
	def.Process[0].StartEvent[0].Outgoing[0].SetFlow(outFromStartEvent)

	/*** set shape attributes (startevent) ***/
	def.Diagram[0].Plane[0].Shape[1].SetID("startevent", stevN+1)
	def.Diagram[0].Plane[0].Shape[1].SetElement("startevent", stevN)
	def.Diagram[0].Plane[0].Shape[1].SetBounds()
	def.Diagram[0].Plane[0].Shape[1].Bounds[0].SetCoordinates(179, 150)
	def.Diagram[0].Plane[0].Shape[1].Bounds[0].SetSize(36, 36)

	/** set task attributes **/
	// generics
	taskHash := utils.GenerateHash()
	outFromTask := utils.GenerateHash()
	def.Process[0].Task[0].SetID(taskHash)
	// Incoming
	def.Process[0].Task[0].SetIncoming(1)
	def.Process[0].Task[0].Incoming[0].SetFlow(outFromStartEvent)
	// Outgoing
	def.Process[0].Task[0].SetOutgoing(1)
	def.Process[0].Task[0].Outgoing[0].SetFlow(outFromTask)

	/*** set shape attributes (task) ***/
	def.Diagram[0].Plane[0].Shape[2].SetID("activity", taskHash)
	def.Diagram[0].Plane[0].Shape[2].SetElement("activity", taskHash)
	def.Diagram[0].Plane[0].Shape[2].SetBounds()
	def.Diagram[0].Plane[0].Shape[2].Bounds[0].SetCoordinates(270, 137)
	def.Diagram[0].Plane[0].Shape[2].Bounds[0].SetSize(100, 80)

	/** set end event attributes **/
	// generics
	endEventHash := utils.GenerateHash()
	def.Process[0].EndEvent[0].SetID(endEventHash)
	// Incoming
	def.Process[0].EndEvent[0].SetIncoming(1)
	def.Process[0].EndEvent[0].Incoming[0].SetFlow(outFromTask)

	/*** set shape attributes (endevent) ***/
	def.Diagram[0].Plane[0].Shape[3].SetID("event", endEventHash)
	def.Diagram[0].Plane[0].Shape[3].SetElement("event", endEventHash)
	def.Diagram[0].Plane[0].Shape[3].SetBounds()
	def.Diagram[0].Plane[0].Shape[3].Bounds[0].SetCoordinates(432, 159)
	def.Diagram[0].Plane[0].Shape[3].Bounds[0].SetSize(36, 36)

	/** set sequence flow attributes **/

	// #1
	def.Process[0].SequenceFlow[0].SetID(outFromStartEvent)
	def.Process[0].SequenceFlow[0].SetSourceRef(fmt.Sprintf("StartEvent_%d", stevN))
	def.Process[0].SequenceFlow[0].SetTargetRef(fmt.Sprintf("Activity_%s", taskHash))
	/*** set shape attributes (sequence flow #1) ***/
	def.Diagram[0].Plane[0].Edge[0].SetID("flow", outFromStartEvent)
	def.Diagram[0].Plane[0].Edge[0].SetElement("flow", outFromStartEvent)
	def.Diagram[0].Plane[0].Edge[0].SetWaypoint()
	def.Diagram[0].Plane[0].Edge[0].Waypoint[0].SetCoordinates(215, 177)
	def.Diagram[0].Plane[0].Edge[0].Waypoint[1].SetCoordinates(270, 177)

	// #2
	def.Process[0].SequenceFlow[1].SetID(outFromTask)
	def.Process[0].SequenceFlow[1].SetSourceRef(fmt.Sprintf("Activity_%s", taskHash))
	def.Process[0].SequenceFlow[1].SetTargetRef(fmt.Sprintf("Event_%s", endEventHash))
	/*** set shape attributes (sequence flow #2) ***/
	def.Diagram[0].Plane[0].Edge[1].SetID("flow", outFromStartEvent)
	def.Diagram[0].Plane[0].Edge[1].SetElement("flow", outFromStartEvent)
	def.Diagram[0].Plane[0].Edge[1].SetWaypoint()
	def.Diagram[0].Plane[0].Edge[1].Waypoint[0].SetCoordinates(370, 177)
	def.Diagram[0].Plane[0].Edge[1].Waypoint[1].SetCoordinates(432, 177)
	//++

	/** set diagram attributes **/
	var n int64 = 1
	def.Diagram[0].SetID(n)

	/** set plane attributes **/
	def.Diagram[0].Plane[0].SetID(n)
	def.Diagram[0].Plane[0].SetElement("collaboration", procHash)

}

// Create ...
func (bpm *BPMNF) Create() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&bpm.Def, " ", "  ")

	// create file
	f, err := os.Create("files/diagram_" + fmt.Sprintf("%d", bpm.Options.Counter) + ".bpmn")
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
