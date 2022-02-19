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
	BPMNFCounter int
	Def          models.Definitions
}

// NewBPMNF ...
func NewBPMNF() BPMNF {
	files, _ := ioutil.ReadDir("files/")
	counter := len(files)
	if counter == 0 {
		counter++
	}
	return BPMNF{BPMNFCounter: counter}
}

// Set ...
func (bpm *BPMNF) Set() {

	// set namespaces
	def := &bpm.Def
	def.SetBpmn()
	def.SetBpmndi()
	def.SetDC()

	// set definitions id & target namespace
	defHash := utils.GenerateHash()
	def.SetDefinitionsID(defHash)

	def.SetTargetNamespace()

	// set exporter & version
	def.SetExporter()
	def.SetExporterVersion()

	//++

	// set collaboration
	collabHash := utils.GenerateHash()
	collab := &def.Collab
	collab.SetID(collabHash)

	// set participant
	collab.Participant = collab.SetParticipant(1)

	participHash := utils.GenerateHash()
	collab.Participant[0].SetID(participHash)

	// set process reference
	procHash := utils.GenerateHash()
	collab.Participant[0].SetProcessRef(procHash)

	// set process
	def.Proc = def.SetProcess(1)
	def.Proc[0].SetID(procHash)
	name := "Test"
	def.Proc[0].SetName(name)
	isExecutable := true
	def.Proc[0].SetIsExecutable(isExecutable)
	cVersionTag := "v0.1.0"
	def.Proc[0].SetCamundaVersionTag(cVersionTag)

	// set start event
	var stevN int64 = 1
	outFromStartEvent := utils.GenerateHash()
	def.Proc[0].SetStartEvent(1)
	def.Proc[0].StartEvent[0].SetID(stevN)
	def.Proc[0].StartEvent[0].SetOutgoing(1)
	def.Proc[0].StartEvent[0].Outgoing[0].SetFlow(outFromStartEvent)

	// set task
	taskHash := utils.GenerateHash()
	outFromTask := utils.GenerateHash()
	def.Proc[0].SetTask(1)
	def.Proc[0].Task[0].SetID(taskHash)
	def.Proc[0].Task[0].SetIncoming(1)
	def.Proc[0].Task[0].Incoming[0].SetFlow(outFromStartEvent)
	def.Proc[0].Task[0].SetOutgoing(1)
	def.Proc[0].Task[0].Outgoing[0].SetFlow(outFromTask)

	/*
		// set end event
		endEventHash := utils.GenerateHash()
		def.Proc[0].EndEvent = []models.EndEvent{
			{
				ID: fmt.Sprintf("Event_%s", endEventHash),
				Incoming: []models.Incoming{
					{
						Flow: fmt.Sprintf("Flow_%s", outFromTask),
					},
				},
			},
		}

		// set sequence flow
		def.Proc[0].SequenceFlow = []models.SequenceFlow{
			{
				ID:        fmt.Sprintf("Flow_%s", outFromStartEvent),
				SourceRef: fmt.Sprintf("StartEvent_%d", stevN),
				TargetRef: fmt.Sprintf("Activity_%s", taskHash),
			},
			{
				ID:        fmt.Sprintf("Flow_%s", outFromTask),
				SourceRef: fmt.Sprintf("Activity_%s", taskHash),
				TargetRef: fmt.Sprintf("Event_%s", endEventHash),
			},
		}

		//++

		// set diagram
		var n int64 = 1
		diagram := def.Diagram
		diagram.SetID(n)

		// set plane
		plane := diagram.Plane
		plane.SetID(n)
		plane.SetElement("collaboration", procHash)

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
	f, err := os.Create("files/diagram_" + fmt.Sprintf("%d", bpm.BPMNFCounter) + ".bpmn")
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
