package repository

/**************************************************************************************/
/**
 * @Import
 *
 * Import models and utils package to build the model
 **/

import (
	"github.com/deemount/gobpmn/models"
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

type pool struct {
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
	def models.DefinitionsRepository
	pool
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

func NewModel(def *models.Definitions) Model {
	return &model{
		// Definitions
		def: new(models.Definitions),
		// Pool
		pool: pool{
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
func (m model) Def() *models.DefinitionsRepository {
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
	m.GetDiagram().SetPlane()
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
	e := m.GetShapePool(m.GetPlane())
	e.SetID("id", m.ParticipantID)
	e.SetElement("id", m.ParticipantID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.Bounds[0].SetCoordinates(150, 80)
	e.Bounds[0].SetSize(420, 160)
}

func (m *model) setProcess() {
	m.GetProcess().SetID("id", m.ProcessID)
	m.GetProcess().SetIsExecutable(m.IsExecutable)
}

func (m *model) setDiagram() {
	// diagram attributes
	var n int64 = 1
	m.GetDiagram().SetID(n)
	// plane attributes
	p := m.GetPlane()
	p.SetID(n)
	p.SetElement("id", m.CollaborationID)
}

/**
 * @Setters III
 * @Event
 * @private setStartEvent
 * @private setEndEvent
**/
func (m *model) setStartEvent() {
	// assign
	e, d := m.GetStartEvent()
	// element
	e.SetID("event", m.StartEventHash)
	e.SetName("Begin of Process")
	// outgoing
	e.SetOutgoing(1)
	e.Outgoing[0].SetFlow(m.FromStartEventHash)
	// shape
	d.SetID("event", m.StartEventHash)
	d.SetElement("event", m.StartEventHash)
	d.SetBounds()
	d.Bounds[0].SetCoordinates(225, 142)
	d.Bounds[0].SetSize(36, 36)
}

func (m *model) setEndEvent() {
	// assign
	e, d := m.GetEndEvent()
	// element
	e.SetID("event", m.EndEventHash)
	e.SetName("End of Process")
	// incoming
	e.SetIncoming(1)
	e.Incoming[0].SetFlow(m.FromTaskHash)
	// shape
	d.SetID("event", m.EndEventHash)
	d.SetElement("event", m.EndEventHash)
	d.SetBounds()
	d.Bounds[0].SetCoordinates(492, 142)
	d.Bounds[0].SetSize(36, 36)
}

/**
 * @Setters IV
 * @Task
 * @private setTask
**/
func (m *model) setTask() {
	// assign
	e, d := m.GetTask()
	// element
	e.SetID(m.TaskHash)
	e.SetName("Task")
	// incoming
	e.SetIncoming(1)
	e.Incoming[0].SetFlow(m.FromStartEventHash)
	// outgoing
	e.SetOutgoing(1)
	e.Outgoing[0].SetFlow(m.FromTaskHash)
	// shape
	d.SetID("activity", m.TaskHash)
	d.SetElement("activity", m.TaskHash)
	d.SetBounds()
	d.Bounds[0].SetCoordinates(320, 120)
	d.Bounds[0].SetSize(100, 80)
}

/**
 * @Setters V
 * @Sequence
 * @private fromStartEvent
 * @private fromTask
**/
func (m *model) fromStartEvent() {
	// assign
	e, d := m.GetFromStartEvent()
	// element
	e.SetID(m.FromStartEventHash)
	e.SetSourceRef("event", m.StartEventHash)
	e.SetTargetRef("activity", m.TaskHash)
	// edge
	d.SetID("flow", m.FromStartEventHash)
	d.SetElement("flow", m.FromStartEventHash)
	d.SetWaypoint()
	d.Waypoint[0].SetCoordinates(261, 160)
	d.Waypoint[1].SetCoordinates(320, 160)
}

func (m *model) fromTask() {
	// assign
	e, d := m.GetFromTask()
	// element
	e.SetID(m.FromTaskHash)
	e.SetSourceRef("activity", m.TaskHash)
	e.SetTargetRef("event", m.EndEventHash)
	// edge
	d.SetID("flow", m.FromTaskHash)
	d.SetElement("flow", m.FromTaskHash)
	d.SetWaypoint()
	d.Waypoint[0].SetCoordinates(420, 160)
	d.Waypoint[1].SetCoordinates(492, 160)
}

/**************************************************************************************/

/**
 * @Getters I
 * @GetCollaboration -> models.Collaboration
 * @GetParticipant -> models.Participant
 * @GetProcess -> models.Process
 * @GetDiagram -> models.Diagram
**/
func (m model) GetCollaboration() *models.Collaboration {
	return m.def.GetCollaboration()
}

func (m model) GetParticipant(e *models.Collaboration) *models.Participant {
	return e.GetParticipant(0)
}

func (m model) GetProcess() *models.Process {
	return m.def.GetProcess(0)
}

func (m model) GetDiagram() *models.Diagram {
	return m.def.GetDiagram(0)
}

func (m model) GetPlane() *models.Plane {
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
func (m model) GetStartEvent() (*models.StartEvent, *models.Shape) {
	start := m.GetProcess().GetStartEvent(0)
	plane := m.GetPlane()
	shape := m.GetShapeStartEvent(plane)
	return start, shape
}

func (m model) GetEndEvent() (*models.EndEvent, *models.Shape) {
	end := m.GetProcess().GetEndEvent(0)
	plane := m.GetPlane()
	shape := m.GetShapeEndEvent(plane)
	return end, shape
}

func (m model) GetTask() (*models.Task, *models.Shape) {
	task := m.GetProcess().GetTask(0)
	plane := m.GetPlane()
	shape := m.GetShapeTask(plane)
	return task, shape
}

func (m model) GetFromStartEvent() (*models.SequenceFlow, *models.Edge) {
	flow := m.GetProcess().GetSequenceFlow(0)
	plane := m.GetPlane()
	edge := m.GetEdgeFromStartEvent(plane)
	return flow, edge
}

func (m model) GetFromTask() (*models.SequenceFlow, *models.Edge) {
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
func (m model) GetShapePool(e *models.Plane) *models.Shape {
	return e.GetShape(0)
}

func (m model) GetShapeStartEvent(e *models.Plane) *models.Shape {
	return e.GetShape(1)
}

func (m model) GetShapeEndEvent(e *models.Plane) *models.Shape {
	return e.GetShape(2)
}

func (m model) GetShapeTask(e *models.Plane) *models.Shape {
	return e.GetShape(3)
}

func (m model) GetEdgeFromStartEvent(e *models.Plane) *models.Edge {
	return e.GetEdge(0)
}

func (m model) GetEdgeFromTask(e *models.Plane) *models.Edge {
	return e.GetEdge(1)
}
