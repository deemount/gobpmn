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
	p.elements()
	p.attributes()
	p.setProcess()
	p.setStartEvent()
	p.setTask()
	p.setEndEvent()
	p.setDiagram()
	return p
}

// Call ...
func (p Process) Call() core.DefinitionsRepository {
	return p.Def
}

// setInnerElements ...
func (p *Process) elements() {
	// Process Elements
	p.process().SetStartEvent(Builder.NumStartEvent)
	p.process().SetTask(Builder.NumTask)
	p.process().SetEndEvent(Builder.NumEndEvent)
	p.process().SetSequenceFlow(2) // TODO: calculate num of flows by the elements above?
	// Canvas
	p.diagram().SetPlane()
	p.plane().SetShape(Builder.NumShape) // TODO: calculate num of flows by the elements above?
	p.plane().SetEdge(2)                 // TODO: calculate num of flows by the elements above?
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
			SF:    p.process().GetSequenceFlow(0),
			ED:    p.plane().GetEdge(0),
			BSPTR: p.plane().GetShape(0).GetBounds(),
			T:     "flow",
			ST:    "startevent",
			TT:    "activity",
			H:     []string{p.FromStartEvent.Suffix, p.StartEvent.Suffix, p.Task.Suffix}})
}

// setTask ...
func (p *Process) setTask() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA:     p.process().GetTask(0),
			SH:     p.plane().GetShape(1),
			WPPREV: p.plane().GetEdge(0).GetWaypoint(1),
			T:      "activity",
			N:      "Task",
			H:      []string{p.Task.Suffix, p.FromStartEvent.Suffix, p.FromTask.Suffix}})
	p.fromTask()
}

// fromTask ...
func (p *Process) fromTask() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF:    p.process().GetSequenceFlow(1),
			ED:    p.plane().GetEdge(1),
			BSPTR: p.plane().GetShape(1).GetBounds(),
			T:     "flow",
			ST:    "activity",
			TT:    "event",
			H:     []string{p.FromTask.Suffix, p.Task.Suffix, p.EndEvent.Suffix}})
}

// setEndEvent ...
func (p *Process) setEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE:     p.process().GetEndEvent(0),
			SH:     p.plane().GetShape(2),
			WPPREV: p.plane().GetEdge(1).GetWaypoint(1),
			T:      "event",
			N:      "End of Process",
			H:      []string{p.EndEvent.Suffix, p.FromTask.Suffix}})
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
