package examples

/**************************************************************************************/
/**
 * @Import
 *
 * Import models and utils package to build the model
 **/
import (
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/process"
)

/**************************************************************************************/
/**
 * @BASE CLASS
 *
 *
 **/

type PrivateBusiness interface {
	Create() privateBusinessModel
}

type privateBusinessEvent struct{}

type privateBusinessTask struct{}

type privateBusinessSequence struct{}

type privateBusinessModel struct {
	def *core.Definitions
	privateBusinessEvent
	privateBusinessTask
	privateBusinessSequence
}

func NewPrivateBusiness() PrivateBusiness {
	return &privateBusinessModel{
		def:                     new(core.Definitions),
		privateBusinessEvent:    privateBusinessEvent{},
		privateBusinessTask:     privateBusinessTask{},
		privateBusinessSequence: privateBusinessSequence{},
	}
}

func (pb privateBusinessModel) Create() privateBusinessModel {
	pb.setMainElements()  // Initialize number of main elements
	pb.setInnerElements() // Initialize number of inner elements
	pb.setDefinitionsAttributes()
	return pb
}

func (pb *privateBusinessModel) setMainElements() {
	pb.def.SetCollaboration()
	pb.def.SetProcess(1)
	pb.def.SetDiagram(1)
}

func (pb *privateBusinessModel) setInnerElements() {
	diagram := pb.GetDiagram()
	diagram.SetPlane()
	plane := pb.GetPlane()
	plane.SetShape(5)
	plane.SetEdge(4)
}

func (pb *privateBusinessModel) setDefinitionsAttributes() {
	pb.def.SetDefaultAttributes()
}

func (pb privateBusinessModel) GetProcess() *process.Process {
	return pb.def.GetProcess(0)
}

func (pb privateBusinessModel) GetDiagram() *canvas.Diagram {
	return pb.def.GetDiagram(0)
}

func (pb privateBusinessModel) GetPlane() *canvas.Plane {
	return pb.GetDiagram().GetPlane()
}
