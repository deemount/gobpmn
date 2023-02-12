package examples

import (
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/process"
)

type PrivateBusiness interface {
	Create() PrivateBusinessModel
}

type PrivateBusinessModel struct {
	Def core.DefinitionsRepository
	PrivateBusinessEvent
	PrivateBusinessTask
	PrivateBusinessSequence
}

type PrivateBusinessEvent struct{}
type PrivateBusinessTask struct{}
type PrivateBusinessSequence struct{}

func NewPrivateBusiness() PrivateBusiness {
	pb := Builder.Inject(PrivateBusinessModel{}).(PrivateBusinessModel)
	pb.Def = core.NewDefinitions()
	return &pb
}

func (pb PrivateBusinessModel) Create() PrivateBusinessModel {
	Builder.Build(pb.Def)
	pb.setInnerElements()
	pb.setDefinitionsAttributes()
	return pb
}

func (pb *PrivateBusinessModel) setMainElements() {
	pb.Def.SetCollaboration()
	pb.Def.SetProcess(1)
	pb.Def.SetDiagram(1)
}

func (pb PrivateBusinessModel) Build() core.DefinitionsRepository { return pb.Def }

func (pb *PrivateBusinessModel) setInnerElements() {
	diagram := pb.diagram()
	diagram.SetPlane()
	plane := pb.plane()
	plane.SetShape(5)
	plane.SetEdge(4)
}

func (pb *PrivateBusinessModel) setDefinitionsAttributes() { pb.Def.SetDefaultAttributes() }
func (pb PrivateBusinessModel) process() *process.Process  { return pb.Def.GetProcess(0) }
func (pb PrivateBusinessModel) diagram() *canvas.Diagram   { return pb.Def.GetDiagram(0) }
func (pb PrivateBusinessModel) plane() *canvas.Plane       { return pb.diagram().GetPlane() }
