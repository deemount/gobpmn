package small_process

import (
	"github.com/deemount/gobpmn/factory"
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
)

// Process ...
type Process struct {
	Def            core.DefinitionsRepository
	IsExecutable   bool
	Process        factory.Builder
	StartEvent     factory.Builder
	FromStartEvent factory.Builder
	Task           factory.Builder
	FromTask       factory.Builder
	EndEvent       factory.Builder
}

// New ...
func New() Proxy {
	p := Builder.Inject(Process{}).(Process)
	p.Def = core.NewDefinitions()
	Builder.Build(p.Def)
	return &p
}

// Build ...
func (p Process) Build() Process {
	p.setInnerElements()
	p.attributes()
	p.setProcess()
	p.setStartEvent()
	p.setTask()
	p.setEndEvent()
	p.setDiagram()
	return p
}

// Build ...
func (p Process) Call() core.DefinitionsRepository { return p.Def }

// setInnerElements ...
func (p *Process) setInnerElements() {
	//Processes
	p.Def.SetProcess(Builder.NumProc)
	p.process().SetStartEvent(Builder.NumStartEvent)
	p.process().SetTask(Builder.NumTask)
	p.process().SetEndEvent(Builder.NumEndEvent)
	p.process().SetSequenceFlow(2)
	// Canvas
	diagram := p.diagram()
	diagram.SetPlane()
	plane := p.plane()
	plane.SetShape(3)
	plane.SetEdge(2)
}

// setProcess ...
func (p *Process) setProcess() {
	p.process().SetID("process", p.Process.Suffix)
	p.process().SetIsExecutable(p.IsExecutable)
}

// setStartEvent ...
func (p *Process) setStartEvent() {
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: p.process().GetStartEvent(0),
			SH: p.plane().GetShape(0),
			T:  "startevent",
			N:  "Begin of Process",
			H:  []string{p.StartEvent.Suffix, p.FromStartEvent.Suffix}})
	p.fromStartEvent()
}

// fromStartEvent ...
func (p *Process) fromStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.process().GetSequenceFlow(0),
			T:  "flow",
			ST: "startevent",
			TT: "activity",
			H:  []string{p.FromStartEvent.Suffix, p.StartEvent.Suffix, p.Task.Suffix}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: p.plane().GetEdge(0),
			T: "flow",
			H: p.FromStartEvent.Suffix,
			W: []canvas.Waypoint{{X: 215, Y: 177}, {X: 270, Y: 177}}})
}

// setTask ...
func (p *Process) setTask() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: p.process().GetTask(0),
			T:  "activity",
			N:  "Task",
			H:  []string{p.Task.Suffix, p.FromStartEvent.Suffix, p.FromTask.Suffix}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: p.plane().GetShape(1),
			T: "activity",
			H: p.Task.Suffix,
			B: canvas.Bounds{X: 270, Y: 137}})
	p.fromTask()
}

// fromTask ...
func (p *Process) fromTask() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.process().GetSequenceFlow(1),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{p.FromTask.Suffix, p.Task.Suffix, p.EndEvent.Suffix}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: p.plane().GetEdge(1),
			T: "flow",
			H: p.FromTask.Suffix,
			W: []canvas.Waypoint{{X: 370, Y: 177}, {X: 432, Y: 177}}})
}

// setEndEvent ...
func (p *Process) setEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: p.process().GetEndEvent(0),
			SH: p.plane().GetShape(2),
			T:  "event",
			N:  "End of Process",
			H:  []string{p.EndEvent.Suffix, p.FromTask.Suffix},
			BS: canvas.Bounds{X: 432, Y: 159}})
}

// setDiagram ...
func (p *Process) setDiagram() {
	// diagram attributes
	var n int64 = 1
	p.diagram().SetID("diagram", n)
	// plane attributes
	plane := p.plane()
	plane.SetID("plane", n)
	plane.SetElement("process", p.Process.Suffix)
}

/**** Default Setter/Getter ****/

func (p *Process) attributes()              { p.Def.SetDefaultAttributes() }
func (p Process) process() *process.Process { return p.Def.GetProcess(0) }
func (p Process) diagram() *canvas.Diagram  { return p.Def.GetDiagram(0) }
func (p Process) plane() *canvas.Plane      { return p.diagram().GetPlane() }
