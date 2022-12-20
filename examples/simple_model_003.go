package examples

import (
	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

// SimpleModel003Repository ...
type SimpleModel003Repository interface {
	Create()
	SetElements()
	SetDefinitionsAttributes()
	SetProcess()
	SetStartEvent()
	SetStartEventState()
	SetTask()
	SetTaskState()
	SetEndEvent()
	SetFromStartEventSequenceFlow()
	SetFromStartEventStateSequenceFlow()
	SetFromTaskSequenceFlow()
	SetFromTaskStateSequenceFlow()
	SetDiagram()
}

// SimpleModel003 ...
type SimpleModel003 struct {
	def                 *models.Definitions
	isExecutable        bool
	ProcessHash         string
	StartEventCounter   int64
	FromStartEvent      string
	StartEventStateHash string
	FromStartEventState string
	TaskHash            string
	FromTask            string
	TaskStateHash       string
	FromTaskState       string
	EndEventHash        string
}

// NewSimpleModel003 ...
func NewSimpleModel003(def *models.Definitions) SimpleModel003 {
	return SimpleModel003{
		def:                 def,
		isExecutable:        true,
		ProcessHash:         utils.GenerateHash(),
		StartEventCounter:   1,
		FromStartEvent:      utils.GenerateHash(),
		StartEventStateHash: utils.GenerateHash(),
		FromStartEventState: utils.GenerateHash(),
		TaskHash:            utils.GenerateHash(),
		FromTask:            utils.GenerateHash(),
		TaskStateHash:       utils.GenerateHash(),
		FromTaskState:       utils.GenerateHash(),
		EndEventHash:        utils.GenerateHash(),
	}
}

// Create ...
func (simpleModel003 *SimpleModel003) Create() {
	simpleModel003.SetElements()
	simpleModel003.SetDefinitionsAttributes()
	simpleModel003.SetProcess()
	simpleModel003.SetStartEvent()
	simpleModel003.SetStartEventState()
	simpleModel003.SetTask()
	simpleModel003.SetTaskState()
	simpleModel003.SetEndEvent()
	simpleModel003.SetFromStartEventSequenceFlow()
	simpleModel003.SetFromStartEventStateSequenceFlow()
	simpleModel003.SetFromTaskSequenceFlow()
	simpleModel003.SetFromTaskStateSequenceFlow()
	simpleModel003.SetDiagram()
}

// SetElements ...
func (simpleModel003 *SimpleModel003) SetElements() {
	simpleModel003.def.SetProcess(1)
	simpleModel003.def.Process[0].SetStartEvent(1)
	simpleModel003.def.Process[0].SetTask(1)
	simpleModel003.def.Process[0].SetEndEvent(1)
	simpleModel003.def.Process[0].SetIntermediateThrowEvent(2)
	simpleModel003.def.Process[0].SetSequenceFlow(4)
	simpleModel003.def.SetDiagram(1)
	simpleModel003.def.Diagram[0].SetPlane()
	simpleModel003.def.Diagram[0].Plane[0].SetShape(5)
	simpleModel003.def.Diagram[0].Plane[0].SetEdge(4)
}

// SetDefinitionsAttributes ...
func (simpleModel003 *SimpleModel003) SetDefinitionsAttributes() {
	simpleModel003.def.SetDefaultAttributes()
}

// SetProcess ...
func (simpleModel003 *SimpleModel003) SetProcess() {
	// generics
	simpleModel003.def.Process[0].SetID("hash", simpleModel003.ProcessHash)
	simpleModel003.def.Process[0].SetIsExecutable(simpleModel003.isExecutable)
}

// SetStartEvent ...
func (simpleModel003 *SimpleModel003) SetStartEvent() {
	// generics
	simpleModel003.def.Process[0].StartEvent[0].SetID("counter", simpleModel003.StartEventCounter)
	simpleModel003.def.Process[0].StartEvent[0].SetName("Begin of process")
	// outgoing
	simpleModel003.def.Process[0].StartEvent[0].SetOutgoing(1)
	simpleModel003.def.Process[0].StartEvent[0].Outgoing[0].SetFlow(simpleModel003.FromStartEvent)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Shape[0].SetID("startevent", simpleModel003.StartEventCounter+1)
	p.Shape[0].SetElement("startevent", simpleModel003.StartEventCounter)
	p.Shape[0].SetBounds()
	p.Shape[0].Bounds[0].SetCoordinates(179, 99)
	p.Shape[0].Bounds[0].SetSize(36, 36)
}

// SetStartEventState ...
func (simpleModel003 *SimpleModel003) SetStartEventState() {
	// generics
	simpleModel003.def.Process[0].IntermediateThrowEvent[0].SetID("event", simpleModel003.StartEventStateHash)
	simpleModel003.def.Process[0].IntermediateThrowEvent[0].SetName("State after event")
	// incoming
	simpleModel003.def.Process[0].IntermediateThrowEvent[0].SetIncoming(1)
	simpleModel003.def.Process[0].IntermediateThrowEvent[0].Incoming[0].SetFlow(simpleModel003.FromStartEvent)
	// outgoing
	simpleModel003.def.Process[0].IntermediateThrowEvent[0].SetOutgoing(1)
	simpleModel003.def.Process[0].IntermediateThrowEvent[0].Outgoing[0].SetFlow(simpleModel003.FromStartEventState)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Shape[1].SetID("event", simpleModel003.StartEventStateHash)
	p.Shape[1].SetElement("event", simpleModel003.StartEventStateHash)
	p.Shape[1].SetBounds()
	p.Shape[1].Bounds[0].SetCoordinates(272, 99)
	p.Shape[1].Bounds[0].SetSize(36, 36)
}

// SetTask ...
func (simpleModel003 *SimpleModel003) SetTask() {
	// generics
	simpleModel003.def.Process[0].Task[0].SetID("activity", simpleModel003.TaskHash)
	simpleModel003.def.Process[0].Task[0].SetName("Do something")
	// incoming
	simpleModel003.def.Process[0].Task[0].SetIncoming(1)
	simpleModel003.def.Process[0].Task[0].Incoming[0].SetFlow(simpleModel003.FromStartEventState)
	// outgoing
	simpleModel003.def.Process[0].Task[0].SetOutgoing(1)
	simpleModel003.def.Process[0].Task[0].Outgoing[0].SetFlow(simpleModel003.FromTask)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Shape[2].SetID("activity", simpleModel003.TaskHash)
	p.Shape[2].SetElement("activity", simpleModel003.TaskHash)
	p.Shape[2].SetBounds()
	p.Shape[2].Bounds[0].SetCoordinates(370, 77)
	p.Shape[2].Bounds[0].SetSize(100, 80)
}

// SetTaskState ...
func (simpleModel003 *SimpleModel003) SetTaskState() {
	// generics
	simpleModel003.def.Process[0].IntermediateThrowEvent[1].SetID("event", simpleModel003.TaskStateHash)
	simpleModel003.def.Process[0].IntermediateThrowEvent[1].SetName("State after task")
	// incoming
	simpleModel003.def.Process[0].IntermediateThrowEvent[1].SetIncoming(1)
	simpleModel003.def.Process[0].IntermediateThrowEvent[1].Incoming[0].SetFlow(simpleModel003.FromTask)
	// outgoing
	simpleModel003.def.Process[0].IntermediateThrowEvent[1].SetOutgoing(1)
	simpleModel003.def.Process[0].IntermediateThrowEvent[1].Outgoing[0].SetFlow(simpleModel003.FromTaskState)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Shape[3].SetID("event", simpleModel003.TaskStateHash)
	p.Shape[3].SetElement("event", simpleModel003.TaskStateHash)
	p.Shape[3].SetBounds()
	p.Shape[3].Bounds[0].SetCoordinates(532, 99)
	p.Shape[3].Bounds[0].SetSize(36, 36)
}

// SetEndEvent ...
func (simpleModel003 *SimpleModel003) SetEndEvent() {
	// generics
	simpleModel003.def.Process[0].EndEvent[0].SetID("event", simpleModel003.EndEventHash)
	simpleModel003.def.Process[0].EndEvent[0].SetName("End of process")
	// incoming
	simpleModel003.def.Process[0].EndEvent[0].SetIncoming(1)
	simpleModel003.def.Process[0].EndEvent[0].Incoming[0].SetFlow(simpleModel003.FromTaskState)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Shape[4].SetID("event", simpleModel003.EndEventHash)
	p.Shape[4].SetElement("event", simpleModel003.EndEventHash)
	p.Shape[4].SetBounds()
	p.Shape[4].Bounds[0].SetCoordinates(632, 99)
	p.Shape[4].Bounds[0].SetSize(36, 36)
}

// SetFromStartEventSequenceFlow ...
func (simpleModel003 *SimpleModel003) SetFromStartEventSequenceFlow() {
	// generics
	simpleModel003.def.Process[0].SequenceFlow[0].SetID("flow", simpleModel003.FromStartEvent)
	simpleModel003.def.Process[0].SequenceFlow[0].SetSourceRef("startevent", simpleModel003.StartEventCounter)
	simpleModel003.def.Process[0].SequenceFlow[0].SetTargetRef("event", simpleModel003.StartEventStateHash)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Edge[0].SetID("flow", simpleModel003.FromStartEvent)
	p.Edge[0].SetElement("flow", simpleModel003.FromStartEvent)
	p.Edge[0].SetWaypoint()
	p.Edge[0].Waypoint[0].SetCoordinates(215, 117)
	p.Edge[0].Waypoint[1].SetCoordinates(272, 117)
}

// SetFromStartEventStateSequenceFlow ...
func (simpleModel003 *SimpleModel003) SetFromStartEventStateSequenceFlow() {
	// generics
	simpleModel003.def.Process[0].SequenceFlow[1].SetID("flow", simpleModel003.FromStartEventState)
	simpleModel003.def.Process[0].SequenceFlow[1].SetSourceRef("event", simpleModel003.StartEventStateHash)
	simpleModel003.def.Process[0].SequenceFlow[1].SetTargetRef("activity", simpleModel003.TaskHash)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Edge[1].SetID("flow", simpleModel003.FromStartEventState)
	p.Edge[1].SetElement("flow", simpleModel003.FromStartEventState)
	p.Edge[1].SetWaypoint()
	p.Edge[1].Waypoint[0].SetCoordinates(308, 117)
	p.Edge[1].Waypoint[1].SetCoordinates(370, 117)
}

// SetFromTaskSequenceFlow ...
func (simpleModel003 *SimpleModel003) SetFromTaskSequenceFlow() {
	// generics
	simpleModel003.def.Process[0].SequenceFlow[2].SetID("flow", simpleModel003.FromTask)
	simpleModel003.def.Process[0].SequenceFlow[2].SetSourceRef("activity", simpleModel003.TaskHash)
	simpleModel003.def.Process[0].SequenceFlow[2].SetTargetRef("event", simpleModel003.TaskStateHash)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Edge[2].SetID("flow", simpleModel003.FromTask)
	p.Edge[2].SetElement("flow", simpleModel003.FromTask)
	p.Edge[2].SetWaypoint()
	p.Edge[2].Waypoint[0].SetCoordinates(470, 117)
	p.Edge[2].Waypoint[1].SetCoordinates(532, 117)
}

// SetFromTaskStateSequenceFlow ...
func (simpleModel003 *SimpleModel003) SetFromTaskStateSequenceFlow() {
	// generics
	simpleModel003.def.Process[0].SequenceFlow[3].SetID("flow", simpleModel003.FromTaskState)
	simpleModel003.def.Process[0].SequenceFlow[3].SetSourceRef("event", simpleModel003.TaskStateHash)
	simpleModel003.def.Process[0].SequenceFlow[3].SetTargetRef("event", simpleModel003.EndEventHash)
	// shape attributes
	p := simpleModel003.getPlane()
	p.Edge[3].SetID("flow", simpleModel003.FromTaskState)
	p.Edge[3].SetElement("flow", simpleModel003.FromTaskState)
	p.Edge[3].SetWaypoint()
	p.Edge[3].Waypoint[0].SetCoordinates(568, 117)
	p.Edge[3].Waypoint[1].SetCoordinates(632, 117)
}

// SetDiagram ...
func (simpleModel003 *SimpleModel003) SetDiagram() {
	// diagram attributes
	var n int64 = 1
	simpleModel003.def.Diagram[0].SetID(n)
	// plane attributes
	p := simpleModel003.getPlane()
	p.SetID("plane", n)
	p.SetElement("process", simpleModel003.ProcessHash)
}

// getPlane ...
func (simpleModel003 SimpleModel003) getPlane() *models.Plane {
	return &simpleModel003.def.Diagram[0].Plane[0]
}
