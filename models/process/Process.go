package process

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events/elements"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/models/subprocesses"
	"github.com/deemount/gobpmn/models/tasks"
)

// ProcessRepository ...
type ProcessRepository interface {

	/*
	 * Base
	 */
	// Setter
	SetID(typ string, suffix interface{})
	SetName(name string)
	SetIsExecutable(isExec bool)
	// Getter
	GetID() *string
	GetName() *string
	GetIsExecutable() *bool

	/*
	 * Events
	 */
	// Setter
	SetStartEvent(num int)
	SetBoundaryEvent(num int)
	SetEndEvent(num int)
	SetIntermediateCatchEvent(num int)
	SetIntermediateThrowEvent(num int)
	// Getter
	GetStartEvent(num int) *elements.StartEvent
	GetBoundaryEvent(num int) *elements.BoundaryEvent
	GetEndEvent(num int) *elements.EndEvent
	GetIntermediateCatchEvent(num int) *elements.IntermediateCatchEvent
	GetIntermediateThrowEvent(num int) *elements.IntermediateThrowEvent

	/*
	 * Gateways
	 */
	// Setter
	SetExclusiveGateway(num int)
	SetInclusiveGateway(num int)
	SetParallelGateway(num int)
	SetComplexGateway(num int)
	SetEventBasedGateway(num int)
	// Getter
	GetExclusiveGateway(num int) *gateways.ExclusiveGateway
	GetInclusiveGateway(num int) *gateways.InclusiveGateway
	GetParallelGateway(num int) *gateways.ParallelGateway
	GetComplexGateway(num int) *gateways.ComplexGateway
	GetEventBasedGateway(num int) *gateways.EventBasedGateway

	/*
	 * Marker
	 */
	SetSequenceFlow(num int)
	GetSequenceFlow(num int) *marker.SequenceFlow

	/*
	 * Camunda
	 */
	// Setter
	SetCamundaVersionTag(tag string)
	SetCamundaJobPriority(priority int)
	SetCamundaTaskPriority(priority int)
	SetCamundaCandidateStarterGroups(groups string)
	SetCamundaCandiddateStarterUsers(users string)
	// Getter
	GetCamundaVersionTag() *string
	GetCamundaJobPriority() *int
	GetCamundaTaskPriority() *int
	GetCamundaCandidateStarterGroups() *string
	GetCamundaCandiddateStarterUsers() *string

	/*
	 *
	 */
	SetDocumentation()
	GetDocumentation() *attributes.Documentation

	/*
	 * Pool
	 */
	SetLaneSet()
	GetLaneSet() *pool.LaneSet

	/*
	 * Tasks
	 */
	// Setter
	SetTask(num int)
	SetUserTask(num int)
	SetManualTask(num int)
	SetReceiveTask(num int)
	SetScriptTask(num int)
	SetSendTask(num int)
	SetServiceTask(num int)
	// Getter
	GetTask(num int) *tasks.Task
	GetUserTask(num int) *tasks.UserTask
	GetManualTask(num int) *tasks.ManualTask
	GetReceiveTask(num int) *tasks.ReceiveTask
	GetScriptTask(num int) *tasks.ScriptTask
	GetSendTask(num int) *tasks.SendTask
	GetServiceTask(num int) *tasks.ServiceTask

	/*
	 * Subprocesses
	 */
	// Setter
	SetCallActivity(num int)
	SetSubProcess(num int)
	SetTransaction(num int)
	SetAdHocSubprocess(num int)
	// Getter
	GetCallActivity(num int) *subprocesses.CallActivity
	GetSubProcess(num int) *subprocesses.SubProcess
	GetTransaction(num int) *subprocesses.Transaction
	GetAdHocSubProcess(num int) *subprocesses.AdHocSubProcess
}

// Process ...
type Process struct {
	// Base
	ID           string `xml:"id,attr" json:"id" csv:"ID"`
	Name         string `xml:"name,attr,omitempty" json:"name,omitempty" csv:"NAME"`
	IsExecutable bool   `xml:"isExecutable,attr" json:"isExecutable,omitempty" csv:"IS_EXECUTABLE"`
	// Camunda
	CamundaVersionTag             string `xml:"camunda:versionTag,attr,omitempty" json:"versionTag,omitempty" csv:"VERSION_TAG"`
	CamundaJobPriority            int    `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty" csv:"JOB_PRIORITY"`
	CamundaTaskPriority           int    `xml:"camunda:taskPriority,attr,omitempty" json:"taskPriority,omitempty" csv:"TASK_PRIORITY"`
	CamundaCandidateStarterGroups string `xml:"camunda:candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty" csv:"CANDIDATE_STARTER_GROUPS"`
	CamundaCandidateStarterUsers  string `xml:"camunda:candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty" csv:"CANDIDATE_STARTER_USERS"`
	//
	Documentation []attributes.Documentation `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"-"`
	// Pool
	LaneSet []pool.LaneSet `xml:"bpmn:laneSet,omitempty" json:"laneSet,omitempty" csv:"-"`
	// Events
	StartEvent             []elements.StartEvent             `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty" csv:"-"`
	BoundaryEvent          []elements.BoundaryEvent          `xml:"bpmn:boundaryEvent,omitemnpty" json:"boundaryEvent,omitempty" csv:"-"`
	EndEvent               []elements.EndEvent               `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent []elements.IntermediateCatchEvent `xml:"bpmn:intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty" csv:"-"`
	IntermediateThrowEvent []elements.IntermediateThrowEvent `xml:"bpmn:intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty" csv:"-"`
	// Tasks
	Task        []tasks.Task        `xml:"bpmn:task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask    []tasks.UserTask    `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask  []tasks.ManualTask  `xml:"bpmn:manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask []tasks.ReceiveTask `xml:"bpmn:receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask  []tasks.ScriptTask  `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask    []tasks.SendTask    `xml:"bpmn:sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask []tasks.ServiceTask `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
	// Subprocesses
	CallActivity    []subprocesses.CallActivity    `xml:"bpmn:callActivity,omitempty" json:"callActivity,omitempty" csv:"-"`
	SubProcess      []subprocesses.SubProcess      `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty" csv:"-"`
	Transaction     []subprocesses.Transaction     `xml:"bpmn:transaction,omitempty" json:"transaction,omitempty" csv:"-"`
	AdHocSubprocess []subprocesses.AdHocSubProcess `xml:"bpmn:adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty" csv:"-"`
	// Gateways
	ExclusiveGateway  []gateways.ExclusiveGateway  `xml:"bpmn:exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty" csv:"-"`
	InclusiveGateway  []gateways.InclusiveGateway  `xml:"bpmn:inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty" csv:"-"`
	ParallelGateway   []gateways.ParallelGateway   `xml:"bpmn:parallelGateway,omitempty" json:"parallelGateway,omitempty" csv:"-"`
	ComplexGateway    []gateways.ComplexGateway    `xml:"bpmn:complexGateway,omitempty" json:"complexGateway,omitempty" csv:"-"`
	EventBasedGateway []gateways.EventBasedGateway `xml:"bpmn:eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty" csv:"-"`
	// Marker
	Association  []marker.Association  `xml:"bpmn:association,omitempty" json:"association,omitempty"`
	SequenceFlow []marker.SequenceFlow `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty" csv:"-"`
}

// TProcess ...
type TProcess struct {
	ID                     string                             `xml:"id,attr" json:"id"`
	Name                   string                             `xml:"name,attr,omitempty" json:"name,omitempty"`
	IsExecutable           bool                               `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	VersionTag             string                             `xml:"versionTag,attr,omitempty" json:"versionTag,omitempty"`
	JobPriority            int                                `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	TaskPriority           int                                `xml:"taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CandidateStarterGroups string                             `xml:"candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CandidateStarterUsers  string                             `xml:"candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	Documentation          []attributes.Documentation         `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LaneSet                []pool.LaneSet                     `xml:"laneSet,omitempty" json:"laneSet,omitempty"`
	StartEvent             []elements.TStartEvent             `xml:"startEvent,omitemnpty" json:"startEvent,omitempty"`
	BoundaryEvent          []elements.TBoundaryEvent          `xml:"boundaryEvent,omitemnpty" json:"boundaryEvent,omitempty"`
	EndEvent               []elements.TEndEvent               `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent []elements.TIntermediateCatchEvent `xml:"intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty"`
	IntermediateThrowEvent []elements.TIntermediateThrowEvent `xml:"intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty"`
	Task                   []tasks.TTask                      `xml:"task,omitempty" json:"task,omitempty"`
	UserTask               []tasks.TUserTask                  `xml:"userTask,omitempty" json:"userTask,omitempty"`
	ManualTask             []tasks.TManualTask                `xml:"manualTask,omitempty" json:"manualTask,omitempty"`
	ReceiveTask            []tasks.TReceiveTask               `xml:"receiveTask,omitempty" json:"receiveTask,omitempty"`
	ScriptTask             []tasks.TScriptTask                `xml:"scriptTask,omitempty" json:"scriptTask,omitempty"`
	SendTask               []tasks.TSendTask                  `xml:"sendTask,omitempty" json:"sendTask,omitempty"`
	ServiceTask            []tasks.TServiceTask               `xml:"serviceTask,omitempty" json:"serviceTask,omitempty"`
	CallActivity           []subprocesses.TCallActivity       `xml:"callActivity,omitempty" json:"callActivity,omitempty"`
	SubProcess             []subprocesses.TSubProcess         `xml:"subProcess,omitempty" json:"subProcess,omitempty"`
	Transaction            []subprocesses.TTransaction        `xml:"transaction,omitempty" json:"transaction,omitempty"`
	AdHocSubprocess        []subprocesses.TAdHocSubProcess    `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty" csv:"-"`
	ExclusiveGateway       []gateways.TExclusiveGateway       `xml:"exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
	InclusiveGateway       []gateways.TInclusiveGateway       `xml:"inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
	ParallelGateway        []gateways.TParallelGateway        `xml:"parallelGateway,omitempty" json:"parallelGateway,omitempty"`
	ComplexGateway         []gateways.TComplexGateway         `xml:"complexGateway,omitempty" json:"complexGateway,omitempty"`
	EventBasedGateway      []gateways.TEventBasedGateway      `xml:"eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty"`
	Association            []marker.TAssociation              `xml:"association,omitempty" json:"association,omitempty"`
	SequenceFlow           []marker.TSequenceFlow             `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

func NewProcess() ProcessRepository {
	return &Process{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (process *Process) SetID(typ string, suffix interface{}) {
	switch typ {
	case "hash":
		process.ID = fmt.Sprintf("Process_%v", suffix)
		break
	case "id":
		process.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (process *Process) SetName(name string) {
	process.Name = name
}

// SetIsExecutable ...
func (process *Process) SetIsExecutable(isExec bool) {
	process.IsExecutable = isExec
}

/** Camunda **/

// SetCamundaVersionTag ...
func (process *Process) SetCamundaVersionTag(tag string) {
	process.CamundaVersionTag = tag
}

// SetCamundaJobpriority ...
func (process *Process) SetCamundaJobPriority(priority int) {
	process.CamundaJobPriority = priority
}

// SetCamundaTaskPriority ...
func (process *Process) SetCamundaTaskPriority(priority int) {
	process.CamundaTaskPriority = priority
}

// SetCamundaCandidateStarterGroups ...
func (process *Process) SetCamundaCandidateStarterGroups(groups string) {
	process.CamundaCandidateStarterGroups = groups
}

// SetCamundaCandidateStarterUsers
func (process *Process) SetCamundaCandiddateStarterUsers(users string) {
	process.CamundaCandidateStarterUsers = users
}

/*** Make Elements ***/

/** BPMN **/

/** Documentation **/

// SetDocumentation ...
func (process *Process) SetDocumentation() {
	process.Documentation = make([]attributes.Documentation, 1)
}

/** LaneSet **/

// SetLaneSet ...
func (process *Process) SetLaneSet() {
	process.LaneSet = make([]pool.LaneSet, 1)
}

/*** Event ***/

// SetStartEvent ...
func (process *Process) SetStartEvent(num int) {
	process.StartEvent = make([]elements.StartEvent, num)
}

// SetBoundEvent ...
func (process *Process) SetBoundaryEvent(num int) {
	process.BoundaryEvent = make([]elements.BoundaryEvent, num)
}

// SetEndEvent ...
func (process *Process) SetEndEvent(num int) {
	process.EndEvent = make([]elements.EndEvent, num)
}

// SetIntermedCatchEvent ...
func (process *Process) SetIntermediateCatchEvent(num int) {
	process.IntermediateCatchEvent = make([]elements.IntermediateCatchEvent, num)
}

// SetIntermedThrowEvent ...
func (process *Process) SetIntermediateThrowEvent(num int) {
	process.IntermediateThrowEvent = make([]elements.IntermediateThrowEvent, num)
}

/*** Task ***/

// SetTask ...
func (process *Process) SetTask(num int) {
	process.Task = make([]tasks.Task, num)
}

// SetUserTask ...
func (process *Process) SetUserTask(num int) {
	process.UserTask = make([]tasks.UserTask, num)
}

// SetManualTask ...
func (process *Process) SetManualTask(num int) {
	process.ManualTask = make([]tasks.ManualTask, num)
}

// SetReceiveTask ...
func (process *Process) SetReceiveTask(num int) {
	process.ReceiveTask = make([]tasks.ReceiveTask, num)
}

// SetScriptTask ...
func (process *Process) SetScriptTask(num int) {
	process.ScriptTask = make([]tasks.ScriptTask, num)
}

// SetSendTask ...
func (process *Process) SetSendTask(num int) {
	process.SendTask = make([]tasks.SendTask, num)
}

// SetServiceTask ...
func (process *Process) SetServiceTask(num int) {
	process.ServiceTask = make([]tasks.ServiceTask, num)
}

// SetCallActivity ...
func (process *Process) SetCallActivity(num int) {
	process.CallActivity = make([]subprocesses.CallActivity, num)
}

// SetSubProcess ...
func (process *Process) SetSubProcess(num int) {
	process.SubProcess = make([]subprocesses.SubProcess, num)
}

// SetTransaction ...
func (process *Process) SetTransaction(num int) {
	process.Transaction = make([]subprocesses.Transaction, num)
}

// SetAdHocSubProcess ...
func (process *Process) SetAdHocSubprocess(num int) {
	process.AdHocSubprocess = make([]subprocesses.AdHocSubProcess, num)
}

/*** Gateway ***/

// SetExclusiveGateway
func (process *Process) SetExclusiveGateway(num int) {
	process.ExclusiveGateway = make([]gateways.ExclusiveGateway, num)
}

// SetInclsuiveGateway
func (process *Process) SetInclusiveGateway(num int) {
	process.InclusiveGateway = make([]gateways.InclusiveGateway, num)
}

// SetParallelGateway
func (process *Process) SetParallelGateway(num int) {
	process.ParallelGateway = make([]gateways.ParallelGateway, num)
}

// SetComplexGateway
func (process *Process) SetComplexGateway(num int) {
	process.ComplexGateway = make([]gateways.ComplexGateway, num)
}

// SetEventBasedGateway
func (process *Process) SetEventBasedGateway(num int) {
	process.EventBasedGateway = make([]gateways.EventBasedGateway, num)
}

/*** Marker ***/

// SetAssociation ...
func (process *Process) SetAssociation(num int) {
	process.Association = make([]marker.Association, num)
}

// SetSequenceFlow ...
func (process *Process) SetSequenceFlow(num int) {
	process.SequenceFlow = make([]marker.SequenceFlow, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (process Process) GetID() *string {
	return &process.ID
}

// GetName ...
func (process Process) GetName() *string {
	return &process.Name
}

// GetIsExecutable ...
func (process Process) GetIsExecutable() *bool {
	return &process.IsExecutable
}

/** Camunda **/

// GetCamundaVersionTag ...
func (process *Process) GetCamundaVersionTag() *string {
	return &process.CamundaVersionTag
}

// GetCamundaJobpriority ...
func (process *Process) GetCamundaJobPriority() *int {
	return &process.CamundaJobPriority
}

// GetCamundaTaskPriority ...
func (process *Process) GetCamundaTaskPriority() *int {
	return &process.CamundaTaskPriority
}

// GetCamundaCandidateStarterGroups ...
func (process *Process) GetCamundaCandidateStarterGroups() *string {
	return &process.CamundaCandidateStarterGroups
}

// GetCamundaCandidateStarterUsers
func (process *Process) GetCamundaCandiddateStarterUsers() *string {
	return &process.CamundaCandidateStarterUsers
}

/* Elements */

/** BPMN **/

/** Documentation **/

// GetDocumentation ...
func (process Process) GetDocumentation() *attributes.Documentation {
	return &process.Documentation[0]
}

/** LaneSet **/

// GetLaneSet ...
func (process Process) GetLaneSet() *pool.LaneSet {
	return &process.LaneSet[0]
}

/*** Event ***/

// GetStartEvent ...
func (process Process) GetStartEvent(num int) *elements.StartEvent {
	return &process.StartEvent[num]
}

// GetBoundaryEvent ...
func (process Process) GetBoundaryEvent(num int) *elements.BoundaryEvent {
	return &process.BoundaryEvent[num]
}

// GetEndEvent ...
func (process Process) GetEndEvent(num int) *elements.EndEvent {
	return &process.EndEvent[num]
}

// GetIntermedCatchEvent ...
func (process Process) GetIntermediateCatchEvent(num int) *elements.IntermediateCatchEvent {
	return &process.IntermediateCatchEvent[num]
}

// GetIntermedThrowEvent ...
func (process Process) GetIntermediateThrowEvent(num int) *elements.IntermediateThrowEvent {
	return &process.IntermediateThrowEvent[num]
}

/*** Task ***/

// GetTask ...
func (process Process) GetTask(num int) *tasks.Task {
	return &process.Task[num]
}

// GetUserTask ...
func (process Process) GetUserTask(num int) *tasks.UserTask {
	return &process.UserTask[num]
}

// GetManualTask ...
func (process Process) GetManualTask(num int) *tasks.ManualTask {
	return &process.ManualTask[num]
}

// GetReceiveTask ...
func (process Process) GetReceiveTask(num int) *tasks.ReceiveTask {
	return &process.ReceiveTask[num]
}

// GetScriptTask ...
func (process Process) GetScriptTask(num int) *tasks.ScriptTask {
	return &process.ScriptTask[num]
}

// GetSendTask ...
func (process Process) GetSendTask(num int) *tasks.SendTask {
	return &process.SendTask[num]
}

// GetServiceTask ...
func (process Process) GetServiceTask(num int) *tasks.ServiceTask {
	return &process.ServiceTask[num]
}

// GetCallActivity ...
func (process Process) GetCallActivity(num int) *subprocesses.CallActivity {
	return &process.CallActivity[num]
}

// GetSubProcess ...
func (process Process) GetSubProcess(num int) *subprocesses.SubProcess {
	return &process.SubProcess[num]
}

// GetTransaction ...
func (process Process) GetTransaction(num int) *subprocesses.Transaction {
	return &process.Transaction[num]
}

// GetAdHocSubProcess ...
func (process Process) GetAdHocSubProcess(num int) *subprocesses.AdHocSubProcess {
	return &process.AdHocSubprocess[num]
}

/*** Gateway ***/

// GetExclusiveGateway
func (process Process) GetExclusiveGateway(num int) *gateways.ExclusiveGateway {
	return &process.ExclusiveGateway[num]
}

// GetInclsuiveGateway
func (process Process) GetInclusiveGateway(num int) *gateways.InclusiveGateway {
	return &process.InclusiveGateway[num]
}

// GetParallelGateway
func (process Process) GetParallelGateway(num int) *gateways.ParallelGateway {
	return &process.ParallelGateway[num]
}

// GetComplexGateway
func (process Process) GetComplexGateway(num int) *gateways.ComplexGateway {
	return &process.ComplexGateway[num]
}

// GetEventBasedGateway
func (process Process) GetEventBasedGateway(num int) *gateways.EventBasedGateway {
	return &process.EventBasedGateway[num]
}

/*** Marker ***/

// GetAssociation ...
func (process Process) GetAssociation(num int) *marker.Association {
	return &process.Association[num]
}

// GetSequenceFlow ...
func (process Process) GetSequenceFlow(num int) *marker.SequenceFlow {
	return &process.SequenceFlow[num]
}
