package examples

import (
	"fmt"

	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

// SimpleModel001Repository ...
type SimpleModel001Repository interface {
	Create()
	SetElements()
	SetDefinitionsAttributes()
	SetProcess()
	SetStartEvent()
	SetTask()
	SetEndEvent()
	SetFromStartEventSequenceFlow()
	SetFromTaskSequenceFlow()
	SetDiagram()
}

// SimpleModel001 ...
type SimpleModel001 struct {
	def               *models.Definitions
	isExecutable      bool
	ProcessHash       string
	StartEventCounter int64
	FromStartEvent    string
	TaskHash          string
	FromTask          string
	EndEventHash      string
}

// NewSimpleModel001 ...
func NewSimpleModel001(def *models.Definitions) SimpleModel001 {
	return SimpleModel001{
		def:               def,
		isExecutable:      true,
		ProcessHash:       utils.GenerateHash(),
		StartEventCounter: 1,
		FromStartEvent:    utils.GenerateHash(),
		TaskHash:          utils.GenerateHash(),
		FromTask:          utils.GenerateHash(),
		EndEventHash:      utils.GenerateHash(),
	}
}

// Create ...
func (simpleModel001 *SimpleModel001) Create() {
	simpleModel001.SetElements()
	simpleModel001.SetDefinitionsAttributes()
	simpleModel001.SetProcess()
	simpleModel001.SetStartEvent()
	simpleModel001.SetTask()
	simpleModel001.SetEndEvent()
	simpleModel001.SetFromStartEventSequenceFlow()
	simpleModel001.SetFromTaskSequenceFlow()
	simpleModel001.SetDiagram()
}

// SetElements ...
func (simpleModel001 *SimpleModel001) SetElements() {
	simpleModel001.def.SetProcess(1)
	simpleModel001.def.Process[0].SetStartEvent(1)
	simpleModel001.def.Process[0].SetTask(1)
	simpleModel001.def.Process[0].SetEndEvent(1)
	simpleModel001.def.Process[0].SetSequenceFlow(2)
	simpleModel001.def.SetDiagram()
	simpleModel001.def.Diagram[0].SetPlane()
	simpleModel001.def.Diagram[0].Plane[0].SetShape(3)
	simpleModel001.def.Diagram[0].Plane[0].SetEdge(2)
}

// SetDefinitionsAttributes ...
func (simpleModel001 *SimpleModel001) SetDefinitionsAttributes() {
	simpleModel001.def.SetDefaultAttributes()
}

// SetProcess ...
func (simpleModel001 *SimpleModel001) SetProcess() {
	// generics
	simpleModel001.def.Process[0].SetID(simpleModel001.ProcessHash)
	simpleModel001.def.Process[0].SetIsExecutable(simpleModel001.isExecutable)
}

// SetStartEvent ...
func (simpleModel001 *SimpleModel001) SetStartEvent() {
	// generics
	simpleModel001.def.Process[0].StartEvent[0].SetID(simpleModel001.StartEventCounter)
	// outgoing
	simpleModel001.def.Process[0].StartEvent[0].SetOutgoing(1)
	simpleModel001.def.Process[0].StartEvent[0].Outgoing[0].SetFlow(simpleModel001.FromStartEvent)
	// shape attributes
	simpleModel001.def.Diagram[0].Plane[0].Shape[0].SetID("startevent", simpleModel001.StartEventCounter+1)
	simpleModel001.def.Diagram[0].Plane[0].Shape[0].SetElement("startevent", simpleModel001.StartEventCounter)
	simpleModel001.def.Diagram[0].Plane[0].Shape[0].SetBounds()
	simpleModel001.def.Diagram[0].Plane[0].Shape[0].Bounds[0].SetCoordinates(179, 159)
	simpleModel001.def.Diagram[0].Plane[0].Shape[0].Bounds[0].SetSize(36, 36)
}

// SetTask ...
func (simpleModel001 *SimpleModel001) SetTask() {
	// generics
	simpleModel001.def.Process[0].Task[0].SetID(simpleModel001.TaskHash)
	// incoming
	simpleModel001.def.Process[0].Task[0].SetIncoming(1)
	simpleModel001.def.Process[0].Task[0].Incoming[0].SetFlow(simpleModel001.FromStartEvent)
	// outgoing
	simpleModel001.def.Process[0].Task[0].SetOutgoing(1)
	simpleModel001.def.Process[0].Task[0].Outgoing[0].SetFlow(simpleModel001.FromTask)
	// shape attributes
	simpleModel001.def.Diagram[0].Plane[0].Shape[1].SetID("activity", simpleModel001.TaskHash)
	simpleModel001.def.Diagram[0].Plane[0].Shape[1].SetElement("activity", simpleModel001.TaskHash)
	simpleModel001.def.Diagram[0].Plane[0].Shape[1].SetBounds()
	simpleModel001.def.Diagram[0].Plane[0].Shape[1].Bounds[0].SetCoordinates(270, 137)
	simpleModel001.def.Diagram[0].Plane[0].Shape[1].Bounds[0].SetSize(100, 80)
}

// SetEndEvent ...
func (simpleModel001 *SimpleModel001) SetEndEvent() {
	// generics
	simpleModel001.def.Process[0].EndEvent[0].SetID(simpleModel001.EndEventHash)
	// incoming
	simpleModel001.def.Process[0].EndEvent[0].SetIncoming(1)
	simpleModel001.def.Process[0].EndEvent[0].Incoming[0].SetFlow(simpleModel001.FromTask)
	// shape attributes
	simpleModel001.def.Diagram[0].Plane[0].Shape[2].SetID("event", simpleModel001.EndEventHash)
	simpleModel001.def.Diagram[0].Plane[0].Shape[2].SetElement("event", simpleModel001.EndEventHash)
	simpleModel001.def.Diagram[0].Plane[0].Shape[2].SetBounds()
	simpleModel001.def.Diagram[0].Plane[0].Shape[2].Bounds[0].SetCoordinates(432, 159)
	simpleModel001.def.Diagram[0].Plane[0].Shape[2].Bounds[0].SetSize(36, 36)
}

// SetFromStartEventSequenceFlow ...
func (simpleModel001 *SimpleModel001) SetFromStartEventSequenceFlow() {
	// generics
	simpleModel001.def.Process[0].SequenceFlow[0].SetID(simpleModel001.FromStartEvent)
	simpleModel001.def.Process[0].SequenceFlow[0].SetSourceRef(fmt.Sprintf("StartEvent_%d", simpleModel001.StartEventCounter))
	simpleModel001.def.Process[0].SequenceFlow[0].SetTargetRef(fmt.Sprintf("Activity_%s", simpleModel001.TaskHash))
	// shape attributes
	simpleModel001.def.Diagram[0].Plane[0].Edge[0].SetID("flow", simpleModel001.FromStartEvent)
	simpleModel001.def.Diagram[0].Plane[0].Edge[0].SetElement("flow", simpleModel001.FromStartEvent)
	simpleModel001.def.Diagram[0].Plane[0].Edge[0].SetWaypoint()
	simpleModel001.def.Diagram[0].Plane[0].Edge[0].Waypoint[0].SetCoordinates(215, 177)
	simpleModel001.def.Diagram[0].Plane[0].Edge[0].Waypoint[1].SetCoordinates(270, 177)
}

// SetFromTaskSequenceFlow ...
func (simpleModel001 *SimpleModel001) SetFromTaskSequenceFlow() {
	// generics
	simpleModel001.def.Process[0].SequenceFlow[1].SetID(simpleModel001.FromTask)
	simpleModel001.def.Process[0].SequenceFlow[1].SetSourceRef(fmt.Sprintf("Activity_%s", simpleModel001.TaskHash))
	simpleModel001.def.Process[0].SequenceFlow[1].SetTargetRef(fmt.Sprintf("Event_%s", simpleModel001.EndEventHash))
	// shape attributes
	simpleModel001.def.Diagram[0].Plane[0].Edge[1].SetID("flow", simpleModel001.FromTask)
	simpleModel001.def.Diagram[0].Plane[0].Edge[1].SetElement("flow", simpleModel001.FromTask)
	simpleModel001.def.Diagram[0].Plane[0].Edge[1].SetWaypoint()
	simpleModel001.def.Diagram[0].Plane[0].Edge[1].Waypoint[0].SetCoordinates(370, 177)
	simpleModel001.def.Diagram[0].Plane[0].Edge[1].Waypoint[1].SetCoordinates(432, 177)
}

//
func (simpleModel001 *SimpleModel001) SetDiagram() {
	// diagram attributes
	var n int64 = 1
	simpleModel001.def.Diagram[0].SetID(n)
	// plane attributes
	simpleModel001.def.Diagram[0].Plane[0].SetID(n)
	simpleModel001.def.Diagram[0].Plane[0].SetElement("process", simpleModel001.ProcessHash)
}
