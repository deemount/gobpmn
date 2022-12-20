package examples

import (
	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

// SimpleModel002Repository ...
type SimpleModel002Repository interface {
	Create()
	SetElements()
	SetDefinitionsAttributes()
	SetCollaboration()
	SetProcess()
	SetStartEvent()
	SetTask()
	SetEndEvent()
	SetFromStartEventSequenceFlow()
	SetFromTaskSequenceFlow()
	SetDiagram()
}

// SimpleModel002 ...
type SimpleModel002 struct {
	def               *models.Definitions
	isExecutable      bool
	CollaborationHash string
	ParticipantHash   string
	ProcessHash       string
	StartEventCounter int64
	FromStartEvent    string
	TaskHash          string
	FromTask          string
	EndEventHash      string
}

// NewSimpleModel002 ...
func NewSimpleModel002(def *models.Definitions) SimpleModel002 {
	return SimpleModel002{
		def:               def,
		isExecutable:      true,
		CollaborationHash: utils.GenerateHash(),
		ParticipantHash:   utils.GenerateHash(),
		ProcessHash:       utils.GenerateHash(),
		StartEventCounter: 1,
		FromStartEvent:    utils.GenerateHash(),
		TaskHash:          utils.GenerateHash(),
		FromTask:          utils.GenerateHash(),
		EndEventHash:      utils.GenerateHash(),
	}
}

// Create ...
func (simpleModel002 *SimpleModel002) Create() {
	simpleModel002.SetElements()
	simpleModel002.SetDefinitionsAttributes()
	simpleModel002.SetCollaboration()
	simpleModel002.SetProcess()
	simpleModel002.SetStartEvent()
	simpleModel002.SetTask()
	simpleModel002.SetEndEvent()
	simpleModel002.SetFromStartEventSequenceFlow()
	simpleModel002.SetFromTaskSequenceFlow()
	simpleModel002.SetDiagram()
}

// SetElements ...
func (simpleModel002 *SimpleModel002) SetElements() {
	simpleModel002.def.SetCollaboration()
	simpleModel002.def.Collaboration[0].SetParticipant(1)
	simpleModel002.def.SetProcess(1)
	simpleModel002.def.Process[0].SetStartEvent(1)
	simpleModel002.def.Process[0].SetTask(1)
	simpleModel002.def.Process[0].SetEndEvent(1)
	simpleModel002.def.Process[0].SetSequenceFlow(2)
	simpleModel002.def.SetDiagram(1)
	simpleModel002.def.Diagram[0].SetPlane()
	simpleModel002.def.Diagram[0].Plane[0].SetShape(4)
	simpleModel002.def.Diagram[0].Plane[0].SetEdge(2)
}

// SetDefinitionsAttributes ...
func (simpleModel002 *SimpleModel002) SetDefinitionsAttributes() {
	simpleModel002.def.SetDefaultAttributes()
}

// SetCollaboration ...
func (simpleModel002 *SimpleModel002) SetCollaboration() {
	// generics
	simpleModel002.def.Collaboration[0].SetID("collaboration", simpleModel002.CollaborationHash)
	// participant attributes
	// generics
	simpleModel002.def.Collaboration[0].Participant[0].SetID("participant", simpleModel002.ParticipantHash)
	simpleModel002.def.Collaboration[0].Participant[0].SetProcessRef("process", simpleModel002.ProcessHash)
	// shape attributes
	p := simpleModel002.getPlane()
	p.Shape[0].SetID("participant", simpleModel002.ParticipantHash)
	p.Shape[0].SetElement("participant", simpleModel002.ParticipantHash)
	p.Shape[0].SetIsHorizontal(true)
	p.Shape[0].SetBounds()
	p.Shape[0].Bounds[0].SetCoordinates(129, 52)
	p.Shape[0].Bounds[0].SetSize(600, 250)
}

// SetProcess ...
func (simpleModel002 *SimpleModel002) SetProcess() {
	// generics
	simpleModel002.def.Process[0].SetID("hash", simpleModel002.ProcessHash)
	simpleModel002.def.Process[0].SetIsExecutable(simpleModel002.isExecutable)
}

// SetStartEvent ...
func (simpleModel002 *SimpleModel002) SetStartEvent() {
	// generics
	simpleModel002.def.Process[0].StartEvent[0].SetID("counter", simpleModel002.StartEventCounter)
	simpleModel002.def.Process[0].StartEvent[0].SetName("Begin of Process")
	// outgoing
	simpleModel002.def.Process[0].StartEvent[0].SetOutgoing(1)
	simpleModel002.def.Process[0].StartEvent[0].Outgoing[0].SetFlow(simpleModel002.FromStartEvent)
	// shape attributes
	p := simpleModel002.getPlane()
	p.Shape[1].SetID("startevent", simpleModel002.StartEventCounter+1)
	p.Shape[1].SetElement("startevent", simpleModel002.StartEventCounter)
	p.Shape[1].SetBounds()
	p.Shape[1].Bounds[0].SetCoordinates(179, 159)
	p.Shape[1].Bounds[0].SetSize(36, 36)
}

// SetTask ...
func (simpleModel002 *SimpleModel002) SetTask() {
	// generics
	simpleModel002.def.Process[0].Task[0].SetID("activity", simpleModel002.TaskHash)
	// incoming
	simpleModel002.def.Process[0].Task[0].SetIncoming(1)
	simpleModel002.def.Process[0].Task[0].Incoming[0].SetFlow(simpleModel002.FromStartEvent)
	// outgoing
	simpleModel002.def.Process[0].Task[0].SetOutgoing(1)
	simpleModel002.def.Process[0].Task[0].Outgoing[0].SetFlow(simpleModel002.FromTask)
	// shape attributes
	p := simpleModel002.getPlane()
	p.Shape[2].SetID("activity", simpleModel002.TaskHash)
	p.Shape[2].SetElement("activity", simpleModel002.TaskHash)
	p.Shape[2].SetBounds()
	p.Shape[2].Bounds[0].SetCoordinates(270, 137)
	p.Shape[2].Bounds[0].SetSize(100, 80)
}

// SetEndEvent ...
func (simpleModel002 *SimpleModel002) SetEndEvent() {
	// generics
	simpleModel002.def.Process[0].EndEvent[0].SetID("event", simpleModel002.EndEventHash)
	simpleModel002.def.Process[0].StartEvent[0].SetName("End of Process")
	// incoming
	simpleModel002.def.Process[0].EndEvent[0].SetIncoming(1)
	simpleModel002.def.Process[0].EndEvent[0].Incoming[0].SetFlow(simpleModel002.FromTask)
	// shape attributes
	p := simpleModel002.getPlane()
	p.Shape[3].SetID("event", simpleModel002.EndEventHash)
	p.Shape[3].SetElement("event", simpleModel002.EndEventHash)
	p.Shape[3].SetBounds()
	p.Shape[3].Bounds[0].SetCoordinates(432, 159)
	p.Shape[3].Bounds[0].SetSize(36, 36)
}

// SetFromStartEventSequenceFlow ...
func (simpleModel002 *SimpleModel002) SetFromStartEventSequenceFlow() {
	// generics
	simpleModel002.def.Process[0].SequenceFlow[0].SetID("flow", simpleModel002.FromStartEvent)
	simpleModel002.def.Process[0].SequenceFlow[0].SetSourceRef("startevent", simpleModel002.StartEventCounter)
	simpleModel002.def.Process[0].SequenceFlow[0].SetTargetRef("activity", simpleModel002.TaskHash)
	// shape attributes
	p := simpleModel002.getPlane()
	p.Edge[0].SetID("flow", simpleModel002.FromStartEvent)
	p.Edge[0].SetElement("flow", simpleModel002.FromStartEvent)
	p.Edge[0].SetWaypoint()
	p.Edge[0].Waypoint[0].SetCoordinates(215, 177)
	p.Edge[0].Waypoint[1].SetCoordinates(270, 177)
}

// SetFromTaskSequenceFlow ...
func (simpleModel002 *SimpleModel002) SetFromTaskSequenceFlow() {
	// generics
	simpleModel002.def.Process[0].SequenceFlow[1].SetID("flow", simpleModel002.FromTask)
	simpleModel002.def.Process[0].SequenceFlow[1].SetSourceRef("activity", simpleModel002.TaskHash)
	simpleModel002.def.Process[0].SequenceFlow[1].SetTargetRef("evemt", simpleModel002.EndEventHash)
	// shape attributes
	p := simpleModel002.getPlane()
	p.Edge[1].SetID("flow", simpleModel002.FromTask)
	p.Edge[1].SetElement("flow", simpleModel002.FromTask)
	p.Edge[1].SetWaypoint()
	p.Edge[1].Waypoint[0].SetCoordinates(370, 177)
	p.Edge[1].Waypoint[1].SetCoordinates(432, 177)
}

// SetDiagram ...
func (simpleModel002 *SimpleModel002) SetDiagram() {
	// diagram attributes
	var n int64 = 1
	simpleModel002.def.Diagram[0].SetID(n)
	// plane attributes
	p := simpleModel002.getPlane()
	p.SetID("plane", n)
	p.SetElement("collaboration", simpleModel002.CollaborationHash)
}

// getPlane ...
func (simpleModel002 SimpleModel002) getPlane() *models.Plane {
	return &simpleModel002.def.Diagram[0].Plane[0]
}
