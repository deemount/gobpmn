package examples

import (
	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

// SimpleModel004Repository ...
type SimpleModel004Repository interface {
	Create() simpleModel004
}

// SimpleModel004 ...
type simpleModel004 struct {
	def                 *models.Definitions
	isExecutable        bool
	CollaborationHash   string
	ParticipantHash     string
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

// NewSimpleModel004 ...
func NewSimpleModel004(def *models.Definitions) SimpleModel004Repository {
	return &simpleModel004{
		def:                 def,
		isExecutable:        true,
		CollaborationHash:   utils.GenerateHash(),
		ParticipantHash:     utils.GenerateHash(),
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
func (m004 simpleModel004) Create() simpleModel004 {
	m004.setElements()
	m004.setDefinitionsAttributes()
	m004.setCollaboration()
	m004.setProcess()
	m004.setStartEvent()
	m004.setStartEventState()
	m004.setTask()
	m004.setTaskState()
	m004.setEndEvent()
	m004.setFromStartEventSequenceFlow()
	m004.setFromStartEventStateSequenceFlow()
	m004.setFromTaskSequenceFlow()
	m004.setFromTaskStateSequenceFlow()
	m004.setDiagram()
	return m004
}

// SetElements ...
func (m004 *simpleModel004) setElements() {
	m004.def.SetCollaboration()
	m004.def.Collaboration[0].SetParticipant(1)
	m004.def.SetProcess(1)
	m004.def.Process[0].SetStartEvent(1)
	m004.def.Process[0].SetTask(1)
	m004.def.Process[0].SetEndEvent(1)
	m004.def.Process[0].SetIntermediateThrowEvent(2)
	m004.def.Process[0].SetSequenceFlow(4)
	m004.def.SetDiagram(1)
	m004.def.Diagram[0].SetPlane()
	m004.def.Diagram[0].Plane[0].SetShape(6)
	m004.def.Diagram[0].Plane[0].SetEdge(4)
}

// SetDefinitionsAttributes ...
func (m004 *simpleModel004) setDefinitionsAttributes() {
	m004.def.SetDefaultAttributes()
}

// SetCollaboration ...
func (simpleModel004 *simpleModel004) setCollaboration() {
	// generics
	simpleModel004.def.Collaboration[0].SetID("collaboration", simpleModel004.CollaborationHash)
	// participant attributes
	// generics
	simpleModel004.def.Collaboration[0].Participant[0].SetID("participant", simpleModel004.ParticipantHash)
	simpleModel004.def.Collaboration[0].Participant[0].SetProcessRef("process", simpleModel004.ProcessHash)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Shape[0].SetID("participant", simpleModel004.ParticipantHash)
	p.Shape[0].SetElement("participant", simpleModel004.ParticipantHash)
	p.Shape[0].SetIsHorizontal(true)
	p.Shape[0].SetBounds()
	p.Shape[0].Bounds[0].SetCoordinates(110, 10)
	p.Shape[0].Bounds[0].SetSize(650, 250)
}

// SetProcess ...
func (simpleModel004 *simpleModel004) setProcess() {
	// generics
	simpleModel004.def.Process[0].SetID("hash", simpleModel004.ProcessHash)
	simpleModel004.def.Process[0].SetIsExecutable(simpleModel004.isExecutable)
}

// SetStartEvent ...
func (simpleModel004 *simpleModel004) setStartEvent() {
	// generics
	simpleModel004.def.Process[0].StartEvent[0].SetID("counter", simpleModel004.StartEventCounter)
	simpleModel004.def.Process[0].StartEvent[0].SetName("Begin of process")
	// outgoing
	simpleModel004.def.Process[0].StartEvent[0].SetOutgoing(1)
	simpleModel004.def.Process[0].StartEvent[0].Outgoing[0].SetFlow(simpleModel004.FromStartEvent)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Shape[1].SetID("startevent", simpleModel004.StartEventCounter+1)
	p.Shape[1].SetElement("startevent", simpleModel004.StartEventCounter)
	p.Shape[1].SetBounds()
	p.Shape[1].Bounds[0].SetCoordinates(179, 99)
	p.Shape[1].Bounds[0].SetSize(36, 36)
}

// SetStartEventState ...
func (simpleModel004 *simpleModel004) setStartEventState() {
	// generics
	simpleModel004.def.Process[0].IntermediateThrowEvent[0].SetID(simpleModel004.StartEventStateHash)
	simpleModel004.def.Process[0].IntermediateThrowEvent[0].SetName("State after event")
	// incoming
	simpleModel004.def.Process[0].IntermediateThrowEvent[0].SetIncoming(1)
	simpleModel004.def.Process[0].IntermediateThrowEvent[0].Incoming[0].SetFlow(simpleModel004.FromStartEvent)
	// outgoing
	simpleModel004.def.Process[0].IntermediateThrowEvent[0].SetOutgoing(1)
	simpleModel004.def.Process[0].IntermediateThrowEvent[0].Outgoing[0].SetFlow(simpleModel004.FromStartEventState)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Shape[2].SetID("event", simpleModel004.StartEventStateHash)
	p.Shape[2].SetElement("event", simpleModel004.StartEventStateHash)
	p.Shape[2].SetBounds()
	p.Shape[2].Bounds[0].SetCoordinates(272, 99)
	p.Shape[2].Bounds[0].SetSize(36, 36)
}

// SetTask ...
func (simpleModel004 *simpleModel004) setTask() {
	// generics
	simpleModel004.def.Process[0].Task[0].SetID(simpleModel004.TaskHash)
	simpleModel004.def.Process[0].Task[0].SetName("Do something")
	// incoming
	simpleModel004.def.Process[0].Task[0].SetIncoming(1)
	simpleModel004.def.Process[0].Task[0].Incoming[0].SetFlow(simpleModel004.FromStartEventState)
	// outgoing
	simpleModel004.def.Process[0].Task[0].SetOutgoing(1)
	simpleModel004.def.Process[0].Task[0].Outgoing[0].SetFlow(simpleModel004.FromTask)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Shape[3].SetID("activity", simpleModel004.TaskHash)
	p.Shape[3].SetElement("activity", simpleModel004.TaskHash)
	p.Shape[3].SetBounds()
	p.Shape[3].Bounds[0].SetCoordinates(370, 77)
	p.Shape[3].Bounds[0].SetSize(100, 80)
}

// SetTaskState ...
func (simpleModel004 *simpleModel004) setTaskState() {
	// generics
	simpleModel004.def.Process[0].IntermediateThrowEvent[1].SetID(simpleModel004.TaskStateHash)
	simpleModel004.def.Process[0].IntermediateThrowEvent[1].SetName("State after task")
	// incoming
	simpleModel004.def.Process[0].IntermediateThrowEvent[1].SetIncoming(1)
	simpleModel004.def.Process[0].IntermediateThrowEvent[1].Incoming[0].SetFlow(simpleModel004.FromTask)
	// outgoing
	simpleModel004.def.Process[0].IntermediateThrowEvent[1].SetOutgoing(1)
	simpleModel004.def.Process[0].IntermediateThrowEvent[1].Outgoing[0].SetFlow(simpleModel004.FromTaskState)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Shape[4].SetID("event", simpleModel004.TaskStateHash)
	p.Shape[4].SetElement("event", simpleModel004.TaskStateHash)
	p.Shape[4].SetBounds()
	p.Shape[4].Bounds[0].SetCoordinates(532, 99)
	p.Shape[4].Bounds[0].SetSize(36, 36)
}

// SetEndEvent ...
func (simpleModel004 *simpleModel004) setEndEvent() {
	// generics
	simpleModel004.def.Process[0].EndEvent[0].SetID("event", simpleModel004.EndEventHash)
	simpleModel004.def.Process[0].EndEvent[0].SetName("End of process")
	// incoming
	simpleModel004.def.Process[0].EndEvent[0].SetIncoming(1)
	simpleModel004.def.Process[0].EndEvent[0].Incoming[0].SetFlow(simpleModel004.FromTaskState)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Shape[5].SetID("event", simpleModel004.EndEventHash)
	p.Shape[5].SetElement("event", simpleModel004.EndEventHash)
	p.Shape[5].SetBounds()
	p.Shape[5].Bounds[0].SetCoordinates(632, 99)
	p.Shape[5].Bounds[0].SetSize(36, 36)
}

// SetFromStartEventSequenceFlow ...
func (simpleModel004 *simpleModel004) setFromStartEventSequenceFlow() {
	// generics
	simpleModel004.def.Process[0].SequenceFlow[0].SetID(simpleModel004.FromStartEvent)
	simpleModel004.def.Process[0].SequenceFlow[0].SetSourceRef("startevent", simpleModel004.StartEventCounter)
	simpleModel004.def.Process[0].SequenceFlow[0].SetTargetRef("event", simpleModel004.StartEventStateHash)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Edge[0].SetID("flow", simpleModel004.FromStartEvent)
	p.Edge[0].SetElement("flow", simpleModel004.FromStartEvent)
	p.Edge[0].SetWaypoint()
	p.Edge[0].Waypoint[0].SetCoordinates(215, 117)
	p.Edge[0].Waypoint[1].SetCoordinates(272, 117)
}

// SetFromStartEventStateSequenceFlow ...
func (simpleModel004 *simpleModel004) setFromStartEventStateSequenceFlow() {
	// generics
	simpleModel004.def.Process[0].SequenceFlow[1].SetID(simpleModel004.FromStartEventState)
	simpleModel004.def.Process[0].SequenceFlow[1].SetSourceRef("event", simpleModel004.StartEventStateHash)
	simpleModel004.def.Process[0].SequenceFlow[1].SetTargetRef("activity", simpleModel004.TaskHash)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Edge[1].SetID("flow", simpleModel004.FromStartEventState)
	p.Edge[1].SetElement("flow", simpleModel004.FromStartEventState)
	p.Edge[1].SetWaypoint()
	p.Edge[1].Waypoint[0].SetCoordinates(308, 117)
	p.Edge[1].Waypoint[1].SetCoordinates(370, 117)
}

// SetFromTaskSequenceFlow ...
func (simpleModel004 *simpleModel004) setFromTaskSequenceFlow() {
	// generics
	simpleModel004.def.Process[0].SequenceFlow[2].SetID(simpleModel004.FromTask)
	simpleModel004.def.Process[0].SequenceFlow[2].SetSourceRef("activity", simpleModel004.TaskHash)
	simpleModel004.def.Process[0].SequenceFlow[2].SetTargetRef("event", simpleModel004.TaskStateHash)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Edge[2].SetID("flow", simpleModel004.FromTask)
	p.Edge[2].SetElement("flow", simpleModel004.FromTask)
	p.Edge[2].SetWaypoint()
	p.Edge[2].Waypoint[0].SetCoordinates(470, 117)
	p.Edge[2].Waypoint[1].SetCoordinates(532, 117)
}

// SetFromTaskStateSequenceFlow ...
func (simpleModel004 *simpleModel004) setFromTaskStateSequenceFlow() {
	// generics
	simpleModel004.def.Process[0].SequenceFlow[3].SetID(simpleModel004.FromTaskState)
	simpleModel004.def.Process[0].SequenceFlow[3].SetSourceRef("event", simpleModel004.TaskStateHash)
	simpleModel004.def.Process[0].SequenceFlow[3].SetTargetRef("event", simpleModel004.EndEventHash)
	// shape attributes
	p := simpleModel004.getPlane()
	p.Edge[3].SetID("flow", simpleModel004.FromTaskState)
	p.Edge[3].SetElement("flow", simpleModel004.FromTaskState)
	p.Edge[3].SetWaypoint()
	p.Edge[3].Waypoint[0].SetCoordinates(568, 117)
	p.Edge[3].Waypoint[1].SetCoordinates(632, 117)
}

// SetDiagram ...
func (simpleModel004 *simpleModel004) setDiagram() {
	// diagram attributes
	var n int64 = 1
	simpleModel004.def.Diagram[0].SetID(n)
	// plane attributes
	p := simpleModel004.getPlane()
	p.SetID(n)
	p.SetElement("collaboration", simpleModel004.CollaborationHash)
}

// getPlane ...
func (simpleModel004 simpleModel004) getPlane() *models.Plane {
	return &simpleModel004.def.Diagram[0].Plane[0]
}
