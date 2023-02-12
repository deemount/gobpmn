package examples

import (
<<<<<<< HEAD
	"github.com/deemount/gobpmn/factory"
=======
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
)

// SimpleModel001Repository ...
type SimpleModel001 interface {
<<<<<<< HEAD
	Create() SM001
}

// SimpleModel001 ...
type SM001 struct {
	Def            core.DefinitionsRepository
	IsExecutable   bool
	Process        factory.BpmnBuilder
	StartEvent     factory.BpmnBuilder
	FromStartEvent factory.BpmnBuilder
	Task           factory.BpmnBuilder
	FromTask       factory.BpmnBuilder
	EndEvent       factory.BpmnBuilder
=======
	Create() simpleModel001
}

// SimpleModel001 ...
type simpleModel001 struct {
	def            core.DefinitionsRepository
	isExecutable   bool
	Process        Hash
	StartEventHash string
	FromStartEvent Hash
	Task           Hash
	FromTask       Hash
	EndEvent       Hash
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

// NewSimpleModel001 ...
func NewSimpleModel001() SimpleModel001 {
<<<<<<< HEAD
	sm001 := bb.Inject(SM001{}).(SM001)
	sm001.Def = core.NewDefinitions(sm001)
	return &sm001
}

// Create ...
func (sm001 SM001) Create() SM001 {
	bb.Build(sm001.Def)
=======
	return &simpleModel001{
		def:            core.NewDefinitions(),
		isExecutable:   true,
		StartEventHash: "1",
	}
}

// Create ...
func (sm001 simpleModel001) Create() simpleModel001 {
	core.SetMainElements(sm001.def, 1) // TODO: Collaboration 'must-be' optional
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	sm001.setInnerElements()
	sm001.setDefinitionsAttributes()
	sm001.setProcess()
	sm001.setStartEvent()
	sm001.setTask()
	sm001.setEndEvent()
	sm001.setDiagram()
	return sm001
}

<<<<<<< HEAD
// Build ...
func (sm001 SM001) Build() core.DefinitionsRepository { return sm001.Def }

// setInnerElements ...
func (sm001 *SM001) setInnerElements() {
	//Processes
	sm001.Def.SetProcess(1)
=======
// Def ...
func (sm001 simpleModel001) Def() core.DefinitionsRepository {
	return sm001.def
}

// setInnerElements ...
func (sm001 *simpleModel001) setInnerElements() {
	//Processes
	sm001.def.SetProcess(1)
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	sm001.process().SetStartEvent(1)
	sm001.process().SetTask(1)
	sm001.process().SetEndEvent(1)
	sm001.process().SetSequenceFlow(2)
<<<<<<< HEAD
=======

>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	// Canvas
	diagram := sm001.diagram()
	diagram.SetPlane()
	plane := sm001.plane()
	plane.SetShape(3)
	plane.SetEdge(2)
}

// SetDefinitionsAttributes ...
<<<<<<< HEAD
func (sm001 *SM001) setDefinitionsAttributes() {
	sm001.Def.SetDefaultAttributes()
}

// setProcess ...
func (sm001 *SM001) setProcess() {
	sm001.process().SetID("process", sm001.Process.Suffix)
	sm001.process().SetIsExecutable(sm001.IsExecutable)
}

// setStartEvent ...
func (sm001 *SM001) setStartEvent() {
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: sm001.process().GetStartEvent(0),
			SH: sm001.plane().GetShape(0),
			T:  "startevent",
			N:  "Begin of Process",
			H:  []string{sm001.StartEvent.Suffix, sm001.FromStartEvent.Suffix}})
	sm001.fromStartEvent()
}

// fromStartEvent ...
func (sm001 *SM001) fromStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: sm001.process().GetSequenceFlow(0),
			T:  "flow",
			ST: "startevent",
			TT: "activity",
			H:  []string{sm001.FromStartEvent.Suffix, sm001.StartEvent.Suffix, sm001.Task.Suffix}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: sm001.plane().GetEdge(0),
			T: "flow",
			H: sm001.FromStartEvent.Suffix,
			W: []canvas.Waypoint{{X: 215, Y: 177}, {X: 270, Y: 177}}})
}

// setTask ...
func (sm001 *SM001) setTask() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: sm001.process().GetTask(0),
			T:  "activity",
			N:  "Task",
			H:  []string{sm001.Task.Suffix, sm001.FromStartEvent.Suffix, sm001.FromTask.Suffix}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: sm001.plane().GetShape(1),
			T: "activity",
			H: sm001.Task.Suffix,
			B: canvas.Bounds{X: 270, Y: 137}})
	sm001.fromTask()
}

// fromTask ...
func (sm001 *SM001) fromTask() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: sm001.process().GetSequenceFlow(1),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{sm001.FromTask.Suffix, sm001.Task.Suffix, sm001.EndEvent.Suffix}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: sm001.plane().GetEdge(1),
			T: "flow",
			H: sm001.FromTask.Suffix,
			W: []canvas.Waypoint{{X: 370, Y: 177}, {X: 432, Y: 177}}})
}

// setEndEvent ...
func (sm001 *SM001) setEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: sm001.process().GetEndEvent(0),
			SH: sm001.plane().GetShape(2),
			T:  "event",
			N:  "End of Process",
			H:  []string{sm001.EndEvent.Suffix, sm001.FromTask.Suffix},
			BS: canvas.Bounds{X: 432, Y: 159}})
}

// setDiagram ...
func (sm001 *SM001) setDiagram() {
=======
func (sm001 *simpleModel001) setDefinitionsAttributes() {
	sm001.def.SetDefaultAttributes()
}

// setProcess ...
func (sm001 *simpleModel001) setProcess() {
	sm001.process().SetID("hash", sm001.Process.Hash())
	sm001.process().SetIsExecutable(sm001.isExecutable)
}

// SetStartEvent ...
func (sm001 *simpleModel001) setStartEvent() {
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: sm001.process().GetStartEvent(0),
			T:  "startevent",
			N:  "Begin of Process",
			H:  []string{sm001.StartEventHash, sm001.FromStartEvent.Hash()}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: sm001.plane().GetShape(0),
			T: "startevent",
			H: sm001.StartEventHash})
	sm001.fromStartEvent()
}

// SetTask ...
func (sm001 *simpleModel001) setTask() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: sm001.process().GetTask(0),
			T:  "activity",
			N:  "Task",
			H:  []string{sm001.Task.string, sm001.FromStartEvent.string, sm001.FromTask.Hash()}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: sm001.plane().GetShape(1),
			T: "activity",
			H: sm001.Task.string,
			B: canvas.Bounds{X: 270, Y: 137}})
	sm001.fromTask()
}

// SetEndEvent ...
func (sm001 *simpleModel001) setEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: sm001.process().GetEndEvent(0),
			T:  "event",
			N:  "End of Process",
			H:  []string{sm001.EndEvent.string, sm001.FromTask.string}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: sm001.plane().GetShape(2),
			T: "event",
			H: sm001.EndEvent.string,
			B: canvas.Bounds{X: 432, Y: 159}})
}

// fromStartEvent ...
func (sm001 *simpleModel001) fromStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: sm001.process().GetSequenceFlow(0),
			T:  "flow",
			ST: "startevent",
			TT: "activity",
			H:  []string{sm001.FromStartEvent.string, sm001.StartEventHash, sm001.Task.Hash()}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: sm001.plane().GetEdge(0),
			T: "flow",
			H: sm001.FromStartEvent.string,
			W: []canvas.Waypoint{{X: 215, Y: 177}, {X: 270, Y: 177}}})
}

// fromTask ...
func (sm001 *simpleModel001) fromTask() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: sm001.process().GetSequenceFlow(1),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{sm001.FromTask.string, sm001.Task.string, sm001.EndEvent.Hash()}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: sm001.plane().GetEdge(1),
			T: "flow",
			H: sm001.FromTask.string,
			W: []canvas.Waypoint{{X: 370, Y: 177}, {X: 432, Y: 177}}})
}

// setDiagram ...
func (sm001 *simpleModel001) setDiagram() {
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	// diagram attributes
	var n int64 = 1
	sm001.diagram().SetID("diagram", n)
	// plane attributes
	p := sm001.plane()
	p.SetID("plane", n)
<<<<<<< HEAD
	p.SetElement("process", sm001.Process.Suffix)
=======
	p.SetElement("process", sm001.Process.string)
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

/*
 * Getter
 */

// process ...
<<<<<<< HEAD
func (sm001 SM001) process() *process.Process {
	return sm001.Def.GetProcess(0)
}

// diagram ...
func (sm001 SM001) diagram() *canvas.Diagram {
	return sm001.Def.GetDiagram(0)
}

// plane ...
func (sm001 SM001) plane() *canvas.Plane {
=======
func (sm001 simpleModel001) process() *process.Process {
	return sm001.def.GetProcess(0)
}

// diagram ...
func (sm001 simpleModel001) diagram() *canvas.Diagram {
	return sm001.def.GetDiagram(0)
}

// plane ...
func (sm001 simpleModel001) plane() *canvas.Plane {
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	return sm001.diagram().GetPlane()
}
