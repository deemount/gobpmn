package examples

import (
	"github.com/deemount/gobpmn/factory"
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
)

type SimpleModel001 interface{ Create() SM001 }

// SimpleModel001 ...
type SM001 struct {
	Def            core.DefinitionsRepository
	IsExecutable   bool
	Process        factory.Builder
	StartEvent     factory.Builder
	FromStartEvent factory.Builder
	Task           factory.Builder
	FromTask       factory.Builder
	EndEvent       factory.Builder
}

// NewSimpleModel001 ...
func NewSimpleModel001() SimpleModel001 {
	sm001 := Builder.Inject(SM001{}).(SM001)
	sm001.Def = core.NewDefinitions()
	Builder.Build(sm001.Def)
	return &sm001
}

// Create ...
func (sm001 SM001) Create() SM001 {
	sm001.setInnerElements()
	sm001.attributes()
	sm001.setProcess()
	sm001.setStartEvent()
	sm001.setTask()
	sm001.setEndEvent()
	sm001.setDiagram()
	return sm001
}

// Build ...
func (sm001 SM001) Build() core.DefinitionsRepository { return sm001.Def }

// setInnerElements ...
func (sm001 *SM001) setInnerElements() {
	//Processes
	sm001.Def.SetProcess(Builder.NumProc)
	sm001.process().SetStartEvent(Builder.NumStartEvent)
	sm001.process().SetTask(Builder.NumTask)
	sm001.process().SetEndEvent(Builder.NumEndEvent)
	sm001.process().SetSequenceFlow(2)
	// Canvas
	diagram := sm001.diagram()
	diagram.SetPlane()
	plane := sm001.plane()
	plane.SetShape(3)
	plane.SetEdge(2)
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
	// diagram attributes
	var n int64 = 1
	sm001.diagram().SetID("diagram", n)
	// plane attributes
	p := sm001.plane()
	p.SetID("plane", n)
	p.SetElement("process", sm001.Process.Suffix)
}

/**** Default Setter/Getter ****/

func (sm001 *SM001) attributes()              { sm001.Def.SetDefaultAttributes() }
func (sm001 SM001) process() *process.Process { return sm001.Def.GetProcess(0) }
func (sm001 SM001) diagram() *canvas.Diagram  { return sm001.Def.GetDiagram(0) }
func (sm001 SM001) plane() *canvas.Plane      { return sm001.diagram().GetPlane() }
