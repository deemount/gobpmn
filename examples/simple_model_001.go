package examples

import (
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/utils"
)

// SimpleModel001Repository ...
type SimpleModel001Repository interface {
	GetProcess()
	GetDiagram()
	GetPlane()
}

// SimpleModel001 ...
type SimpleModel001 struct {
	def               *core.Definitions
	isExecutable      bool
	ProcessHash       string
	StartEventCounter int64
	FromStartEvent    string
	TaskHash          string
	FromTask          string
	EndEventHash      string
}

// NewSimpleModel001 ...
func NewSimpleModel001(def *core.Definitions) SimpleModel001 {
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
	simpleModel001.GetProcess().SetStartEvent(1)
	simpleModel001.GetProcess().SetTask(1)
	simpleModel001.GetProcess().SetEndEvent(1)
	simpleModel001.GetProcess().SetSequenceFlow(2)
	simpleModel001.def.SetDiagram(1)
	simpleModel001.GetDiagram().SetPlane()
	simpleModel001.GetPlane().SetShape(3)
	simpleModel001.GetPlane().SetEdge(2)
}

// SetDefinitionsAttributes ...
func (simpleModel001 *SimpleModel001) SetDefinitionsAttributes() {
	simpleModel001.def.SetDefaultAttributes()
}

// SetProcess ...
func (simpleModel001 *SimpleModel001) SetProcess() {
	// generics
	simpleModel001.GetProcess().SetID("hash", simpleModel001.ProcessHash)
	simpleModel001.GetProcess().SetIsExecutable(simpleModel001.isExecutable)
}

// SetStartEvent ...
func (simpleModel001 *SimpleModel001) SetStartEvent() {
	// assign
	process := simpleModel001.GetProcess()
	p := simpleModel001.GetPlane()
	// generics
	process.GetStartEvent(0).SetID("counter", simpleModel001.StartEventCounter)
	// outgoing
	process.GetStartEvent(0).SetOutgoing(1)
	process.GetStartEvent(0).GetOutgoing(0).SetFlow(simpleModel001.FromStartEvent)
	// shape attributes
	p.GetShape(0).SetID("startevent", simpleModel001.StartEventCounter+1)
	p.GetShape(0).SetElement("startevent", simpleModel001.StartEventCounter)
	p.GetShape(0).SetBounds()
	p.GetShape(0).GetBounds().SetCoordinates(179, 159)
	p.GetShape(0).GetBounds().SetSize(36, 36)
}

// SetTask ...
func (simpleModel001 *SimpleModel001) SetTask() {
	// assign
	process := simpleModel001.GetProcess()
	p := simpleModel001.GetPlane()
	// generics
	process.GetTask(0).SetID("activity", simpleModel001.TaskHash)
	// incoming
	process.GetTask(0).SetIncoming(1)
	process.GetTask(0).Incoming[0].SetFlow(simpleModel001.FromStartEvent)
	// outgoing
	process.GetTask(0).SetOutgoing(1)
	process.GetTask(0).GetOutgoing(0).SetFlow(simpleModel001.FromTask)
	// shape attributes
	p.GetShape(1).SetID("activity", simpleModel001.TaskHash)
	p.GetShape(1).SetElement("activity", simpleModel001.TaskHash)
	p.GetShape(1).SetBounds()
	p.GetShape(1).GetBounds().SetCoordinates(270, 137)
	p.GetShape(1).GetBounds().SetSize(100, 80)
}

// SetEndEvent ...
func (simpleModel001 *SimpleModel001) SetEndEvent() {
	// assign
	process := simpleModel001.GetProcess()
	p := simpleModel001.GetPlane()
	// generics
	process.GetEndEvent(0).SetID("event", simpleModel001.EndEventHash)
	// incoming
	process.GetEndEvent(0).SetIncoming(1)
	process.GetEndEvent(0).Incoming[0].SetFlow(simpleModel001.FromTask)
	// shape attributes
	p.GetShape(2).SetID("event", simpleModel001.EndEventHash)
	p.GetShape(2).SetElement("event", simpleModel001.EndEventHash)
	p.GetShape(2).SetBounds()
	p.GetShape(2).GetBounds().SetCoordinates(432, 159)
	p.GetShape(2).GetBounds().SetSize(36, 36)
}

// SetFromStartEventSequenceFlow ...
func (simpleModel001 *SimpleModel001) SetFromStartEventSequenceFlow() {
	// assign
	process := simpleModel001.GetProcess()
	p := simpleModel001.GetPlane()
	// generics
	process.GetSequenceFlow(0).SetID("flow", simpleModel001.FromStartEvent)
	process.GetSequenceFlow(0).SetSourceRef("startevent", simpleModel001.StartEventCounter)
	process.GetSequenceFlow(0).SetTargetRef("activity", simpleModel001.TaskHash)
	// shape attributes
	p.GetEdge(0).SetID("flow", simpleModel001.FromStartEvent)
	p.GetEdge(0).SetElement("flow", simpleModel001.FromStartEvent)
	p.GetEdge(0).SetWaypoint()
	p.GetEdge(0).GetWaypoint(0).SetCoordinates(215, 177)
	p.GetEdge(0).GetWaypoint(1).SetCoordinates(270, 177)
}

// SetFromTaskSequenceFlow ...
func (simpleModel001 *SimpleModel001) SetFromTaskSequenceFlow() {
	// assign
	process := simpleModel001.GetProcess()
	p := simpleModel001.GetPlane()
	// generics
	process.GetSequenceFlow(1).SetID("flow", simpleModel001.FromTask)
	process.GetSequenceFlow(1).SetSourceRef("activity", simpleModel001.TaskHash)
	process.GetSequenceFlow(1).SetTargetRef("event", simpleModel001.EndEventHash)
	// shape attributes
	p.GetEdge(1).SetID("flow", simpleModel001.FromTask)
	p.GetEdge(1).SetElement("flow", simpleModel001.FromTask)
	p.GetEdge(1).SetWaypoint()
	p.GetEdge(1).GetWaypoint(0).SetCoordinates(370, 177)
	p.GetEdge(1).GetWaypoint(1).SetCoordinates(432, 177)
}

func (simpleModel001 *SimpleModel001) SetDiagram() {
	// diagram attributes
	var n int64 = 1
	simpleModel001.def.Diagram[0].SetID("diagram", n)
	// plane attributes
	p := simpleModel001.GetPlane()
	p.SetID("plane", n)
	p.SetElement("process", simpleModel001.ProcessHash)
}

/*
 * Getter
 */

// GetProcess ...
func (simpleModel001 SimpleModel001) GetProcess() *process.Process {
	return &simpleModel001.def.Process[0]
}

// GetDiagram ...
func (simpleModel001 SimpleModel001) GetDiagram() *canvas.Diagram {
	return &simpleModel001.def.Diagram[0]
}

// GetPlane ...
func (simpleModel001 SimpleModel001) GetPlane() *canvas.Plane {
	return &simpleModel001.def.Diagram[0].Plane[0]
}
