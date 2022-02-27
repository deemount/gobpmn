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

	/** set end event attributes **/
	// generics
	endEventHash := utils.GenerateHash()
	def.Process[0].EndEvent[0].SetID(endEventHash)
	// Incoming
	def.Process[0].EndEvent[0].SetIncoming(1)
	def.Process[0].EndEvent[0].Incoming[0].SetFlow(outFromTask)

	/** set sequence flow attributes **/
	def.Process[0].SequenceFlow[0].SetID(outFromStartEvent)
	def.Process[0].SequenceFlow[0].SetSourceRef(fmt.Sprintf("StartEvent_%d", stevN))
	def.Process[0].SequenceFlow[0].SetTargetRef(fmt.Sprintf("Activity_%s", taskHash))

	def.Process[0].SequenceFlow[1].SetID(outFromTask)
	def.Process[0].SequenceFlow[1].SetSourceRef(fmt.Sprintf("Activity_%s", taskHash))
	def.Process[0].SequenceFlow[1].SetTargetRef(fmt.Sprintf("Event_%s", endEventHash))

	//++

	// set diagram
	var n int64 = 1
	def.SetDiagram()
	def.Diagram[0].SetID(n)

	// set plane
	def.Diagram[0].SetPlane()
	def.Diagram[0].Plane[0].SetID(n)
	def.Diagram[0].Plane[0].SetElement("collaboration", procHash)

	/*
		// set shape
		plane.Shape = []models.Shape{
			{
				ID:           fmt.Sprintf("Participant_%s_di", participHash),
				Element:      fmt.Sprintf("Participant_%s", participHash),
				IsHorizontal: "true",
				Bounds: models.Bounds{
					X:      129,
					Y:      60,
					Width:  600,
					Height: 250,
				},
			},
			{
				ID:      fmt.Sprintf("_BPMNShape_StartEvent_%d", stevN+1),
				Element: fmt.Sprintf("StartEvent_%d", stevN),
				Bounds: models.Bounds{
					X:      179,
					Y:      159,
					Width:  36,
					Height: 36,
				},
			},
			{
				ID:      fmt.Sprintf("Activity_%s_di", taskHash),
				Element: fmt.Sprintf("Activity_%s", taskHash),
				Bounds: models.Bounds{
					X:      270,
					Y:      137,
					Width:  100,
					Height: 80,
				},
			},
			{
				ID:      fmt.Sprintf("Event_%s_di", endEventHash),
				Element: fmt.Sprintf("Event_%s", endEventHash),
				Bounds: models.Bounds{
					X:      432,
					Y:      159,
					Width:  36,
					Height: 36,
				},
			},
		}

		plane.Edge = []models.Edge{
			{
				ID:      fmt.Sprintf("Flow_%s_di", outFromStartEvent),
				Element: fmt.Sprintf("Flow_%s", outFromStartEvent),
				Waypoint: []models.Waypoint{
					{
						X: 215,
						Y: 177,
					},
					{
						X: 270,
						Y: 177,
					},
				},
			},
			{
				ID:      fmt.Sprintf("Flow_%s_di", outFromStartEvent),
				Element: fmt.Sprintf("Flow_%s", outFromStartEvent),
				Waypoint: []models.Waypoint{
					{
						X: 215,
						Y: 177,
					},
					{
						X: 270,
						Y: 177,
					},
				},
			},
		}
	*/
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
