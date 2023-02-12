package examples

/**************************************************************************************/
/**
 * @Import
 *
 * Import models and utils package to build the model
 **/

import (
	"github.com/deemount/gobpmn/models/bpmn/canvas"
<<<<<<< HEAD
	"github.com/deemount/gobpmn/models/bpmn/collaboration"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
=======
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/pool"
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
	"github.com/deemount/gobpmn/utils"
)

/**************************************************************************************/
/**
 * @BASE CLASS
 *
 *
 **/

type Model interface {
	Create() model
}

/**************************************************************************************/
/**
 * @OBJECTS
 * - pool
 * - event
 * - task
 * - sequence
 * - model
 **/

type Pool2 struct {
	IsExecutable    bool
	CollaborationID string
	ParticipantID   string
	ProcessID       string
}

type event struct {
	StartEventHash string
	EndEventHash   string
}

type task struct {
	TaskHash string
}

type sequence struct {
	FromStartEventHash string
	FromTaskHash       string
}

type model struct {
	def core.DefinitionsRepository
	Pool2
	event
	task
	sequence
}

/**************************************************************************************/
/**
 * @BUILDER
 *
 *
 **/

func NewModel(def *core.Definitions) Model {
	return &model{
		// Definitions
		def: def,
		// Pool
		Pool2: Pool2{
			IsExecutable:    true,
			CollaborationID: "uniqueCollaborationId",
			ParticipantID:   "uniqueParticipantId",
			ProcessID:       "uniqueProcessId",
		},
		// Event
		event: event{
			StartEventHash: utils.GenerateHash(),
			EndEventHash:   utils.GenerateHash(),
		},
		// Task
		task: task{
			TaskHash: utils.GenerateHash(),
		},
		// Sequence
		sequence: sequence{
			FromStartEventHash: utils.GenerateHash(),
			FromTaskHash:       utils.GenerateHash(),
		},
	}
}

/**************************************************************************************/
/**
 * @Create
 * @@setMainElements
 * @@setInnerElements
 * @@setDefinitionsAttributes
 * @@setCollaboration
 * @@setProcess
 * @@setStartEvent
 * @@setEndEvent
 * @@setTask
 * @@fromStartEvent
 * @@fromTask
 * @@setDiagram
 **/

func (m model) Create() model {
	m.setMainElements()  // Initialize number of main elements
	m.setInnerElements() // Initialize number of inner elements
	m.setDefinitionsAttributes()
	m.setCollaboration()
	// Process
	m.setProcess()
	// Events
	m.setStartEvent()
	m.setEndEvent()
	// Tasks
	m.setTask()
	// Sequences
	m.fromStartEvent()
	m.fromTask()
	// Diagram
	m.setDiagram()
	return m
}

// Def ...
func (m model) Def() *core.DefinitionsRepository {
	return &m.def
}

/**************************************************************************************/

/**
 * @Setters I
 * @setMainElements
 * @@models.Definitions: SetCollaboration
 * @@models.Definitions: SetProcess
 * @@models.Definitions: SetDiagram
 *
 * @setInnerElements
 * @@Model: GetCollaboration
 * @@Model: GetDiagram
 *
 * @setDefinitionsAttributes
 * @@models.Definitions: SetDefaultAttributes
 *
 * @setCollaboration
 * @@Model: GetParticipant
 * @@Model: GetCollaboration
 **/

func (m *model) setMainElements() {
	m.def.SetCollaboration()
	m.def.SetProcess(1)
	m.def.SetDiagram(1)
}

func (m *model) setInnerElements() {
	collaboration := m.GetCollaboration()
	collaboration.SetParticipant(1)
	process := m.GetProcess()
	process.SetStartEvent(1)
	process.SetEndEvent(1)
	process.SetTask(1)
	process.SetSequenceFlow(2)
	diagram := m.GetDiagram()
	diagram.SetPlane()
	plane := m.GetPlane()
	plane.SetShape(4)
	plane.SetEdge(2)
}

func (m *model) setDefinitionsAttributes() {
	m.def.SetDefaultAttributes()
}

func (m *model) setCollaboration() {
	participant := m.GetParticipant(m.GetCollaboration())
	m.GetCollaboration().SetID("id", m.CollaborationID)
	participant.SetID("id", m.ParticipantID)
	participant.SetName("Participant")
	participant.SetProcessRef("id", m.ProcessID)
	m.setPool()
}

/**
 * @Setters II
 * @Process
 * @private setPool
 * @private setProcess
 * @private setDiagram
**/
func (m *model) setPool() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: m.GetShapePool(m.GetPlane()),
			T: "id",
			I: true,
			H: m.ParticipantID,
			B: canvas.Bounds{X: 150, Y: 80, Width: 420, Height: 160}})

}

func (m *model) setProcess() {
	m.GetProcess().SetID("id", m.ProcessID)
	m.GetProcess().SetIsExecutable(m.IsExecutable)
}

func (m *model) setDiagram() {
	// diagram attributes
	var n int64 = 1
	diagram := m.GetDiagram()
	diagram.SetID("diagram", n)
	// plane attributes
	p := m.GetPlane()
	p.SetID("plane", n)
	p.SetElement("id", m.CollaborationID)
}

/**
 * @Setters III
 * @Event
 * @private setStartEvent
 * @private setEndEvent
**/
func (m *model) setStartEvent() {
	e, d := m.GetStartEvent()
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: e,
			N:  "Begin of Process",
			H:  []string{m.StartEventHash, m.FromStartEventHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: m.StartEventHash,
			B: canvas.Bounds{X: 225, Y: 142, Width: 36, Height: 36}})
}

func (m *model) setEndEvent() {
	e, d := m.GetEndEvent()
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: e,
			N:  "End of Process",
			H:  []string{m.EndEventHash, m.FromTaskHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: m.EndEventHash,
			B: canvas.Bounds{X: 492, Y: 142, Width: 36, Height: 36}})
}

/**
 * @Setters IV
 * @Task
 * @private setTask
**/
func (m *model) setTask() {
	e, d := m.GetTask()
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: e,
			T:  "activity",
			N:  "Task",
			H:  []string{m.TaskHash, m.FromStartEventHash, m.FromTaskHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: m.TaskHash,
			B: canvas.Bounds{X: 320, Y: 120, Width: 100, Height: 80}})
}

/**
 * @Setters V
 * @Sequence
 * @private fromStartEvent
 * @private fromTask
**/
func (m *model) fromStartEvent() {
	e, d := m.GetFromStartEvent()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{m.FromStartEventHash, m.StartEventHash, m.TaskHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: m.FromStartEventHash,
			W: []canvas.Waypoint{{X: 261, Y: 160}, {X: 320, Y: 160}}})
}

func (m *model) fromTask() {
	e, d := m.GetFromTask()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{m.FromTaskHash, m.TaskHash, m.EndEventHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: m.FromTaskHash,
			W: []canvas.Waypoint{{X: 420, Y: 160}, {X: 492, Y: 160}}})
}

/**************************************************************************************/

/**
 * @Getters I
 * @GetCollaboration -> models.Collaboration
 * @GetParticipant -> models.Participant
 * @GetProcess -> models.Process
 * @GetDiagram -> models.Diagram
**/
func (m model) GetCollaboration() *collaboration.Collaboration {
	return m.def.GetCollaboration()
}

func (m model) GetParticipant(e *collaboration.Collaboration) *collaboration.Participant {
	return e.GetParticipant(0)
}

func (m model) GetProcess() *process.Process {
	return m.def.GetProcess(0)
}

func (m model) GetDiagram() *canvas.Diagram {
	return m.def.GetDiagram(0)
}

func (m model) GetPlane() *canvas.Plane {
	return m.GetDiagram().GetPlane()
}

/**
 * @Getters II
 * @GetStartEvent -> models.StartEvent, models.Shape
 * @GetEndEvent -> models.EndEvent, models.Shape
 * @GetTask -> models.Task, models.Shape
 * @GetFromStartEvent -> models.SequenceFlow, models.Shape
 * @GetFromTask -> models.SequenceFlow
**/
func (m model) GetStartEvent() (*elements.StartEvent, *canvas.Shape) {
	start := m.GetProcess().GetStartEvent(0)
	plane := m.GetPlane()
	shape := m.GetShapeStartEvent(plane)
	return start, shape
}

func (m model) GetEndEvent() (*elements.EndEvent, *canvas.Shape) {
	end := m.GetProcess().GetEndEvent(0)
	plane := m.GetPlane()
	shape := m.GetShapeEndEvent(plane)
	return end, shape
}

func (m model) GetTask() (*tasks.Task, *canvas.Shape) {
	task := m.GetProcess().GetTask(0)
	plane := m.GetPlane()
	shape := m.GetShapeTask(plane)
	return task, shape
}

func (m model) GetFromStartEvent() (*flow.SequenceFlow, *canvas.Edge) {
	flow := m.GetProcess().GetSequenceFlow(0)
	plane := m.GetPlane()
	edge := m.GetEdgeFromStartEvent(plane)
	return flow, edge
}

func (m model) GetFromTask() (*flow.SequenceFlow, *canvas.Edge) {
	flow := m.GetProcess().GetSequenceFlow(1)
	plane := m.GetPlane()
	edge := m.GetEdgeFromTask(plane)
	return flow, edge
}

/**
 * @Getters III
 * @Shapes
 * @GetShapePool -> models.Shape
 * @GetShapeStartEvent -> models.Shape
 * @GetShapeEndEvent -> models.Shape
 * @GetShapeTask -> models.Shape
 * @Edges
 * @GetEdgeFromStartEvent -> models.Edge
 * @GetEdgeFromTask -> models.Edge
**/
func (m model) GetShapePool(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(0)
}

func (m model) GetShapeStartEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(1)
}

func (m model) GetShapeEndEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(2)
}

func (m model) GetShapeTask(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(3)
}

func (m model) GetEdgeFromStartEvent(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(0)
}

func (m model) GetEdgeFromTask(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(1)
}
