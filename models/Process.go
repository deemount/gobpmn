package models

import (
	"fmt"

	"github.com/deemount/gobpmn/models/activities"
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/subprocesses"
)

// ProcessRepository ...
type ProcessRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)
	SetIsExecutable(isExec bool)

	SetCamundaVersionTag(tag string)
	SetCamundaJobPriority(priority int)
	SetCamundaTaskPriority(priority int)
	SetCamundaCandidateStarterGroups(groups string)
	SetCamundaCandiddateStarterUsers(users string)

	SetDocumentation()
	SetLaneSet()

	SetStartEvent(num int)
	SetBoundaryEvent(num int)
	SetEndEvent(num int)
	SetIntermediateCatchEvent(num int)
	SetIntermediateThrowEvent(num int)

	SetTask(num int)
	SetUserTask(num int)
	SetManualTask(num int)
	SetReceiveTask(num int)
	SetScriptTask(num int)
	SetSendTask(num int)
	SetServiceTask(num int)

	SetCallActivity(num int)
	SetSubProcess(num int)
	SetTransaction(num int)
	SetAdHocSubprocess(num int)

	SetExclusiveGateway(num int)
	SetInclusiveGateway(num int)
	SetParallelGateway(num int)
	SetComplexGateway(num int)
	SetEventBasedGateway(num int)

	SetSequenceFlow(num int)
	SetGroup(num int)

	GetID() *string
	GetName() *string
	GetIsExecutable() *bool

	GetCamundaVersionTag() *string
	GetCamundaJobPriority() *int
	GetCamundaTaskPriority() *int
	GetCamundaCandidateStarterGroups() *string
	GetCamundaCandiddateStarterUsers() *string

	GetDocumentation() *attributes.Documentation
	GetLaneSet() *LaneSet

	GetStartEvent(num int) *events.StartEvent
	GetBoundaryEvent(num int) *events.BoundaryEvent
	GetEndEvent(num int) *events.EndEvent
	GetIntermediateCatchEvent(num int) *events.IntermediateCatchEvent
	GetIntermediateThrowEvent(num int) *events.IntermediateThrowEvent

	GetTask(num int) *activities.Task
	GetUserTask(num int) *activities.UserTask
	GetManualTask(num int) *activities.ManualTask
	GetReceiveTask(num int) *activities.ReceiveTask
	GetScriptTask(num int) *activities.ScriptTask
	GetSendTask(num int) *activities.SendTask
	GetServiceTask(num int) *activities.ServiceTask

	GetCallActivity(num int) *subprocesses.CallActivity
	GetSubProcess(num int) *subprocesses.SubProcess
	GetTransaction(num int) *subprocesses.Transaction
	GetAdHocSubProcess(num int) *subprocesses.AdHocSubProcess

	GetExclusiveGateway(num int) *gateways.ExclusiveGateway
	GetInclusiveGateway(num int) *gateways.InclusiveGateway
	GetParallelGateway(num int) *gateways.ParallelGateway
	GetComplexGateway(num int) *gateways.ComplexGateway
	GetEventBasedGateway(num int) *gateways.EventBasedGateway

	GetSequenceFlow(num int) *marker.SequenceFlow
	GetGroup(num int) *Group
}

// Process ...
type Process struct {
	ID                            string                          `xml:"id,attr" json:"id" csv:"ID"`
	Name                          string                          `xml:"name,attr,omitempty" json:"name,omitempty" csv:"NAME"`
	IsExecutable                  bool                            `xml:"isExecutable,attr" json:"isExecutable,omitempty" csv:"IS_EXECUTABLE"`
	CamundaVersionTag             string                          `xml:"camunda:versionTag,attr,omitempty" json:"versionTag,omitempty" csv:"VERSION_TAG"`
	CamundaJobPriority            int                             `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty" csv:"JOB_PRIORITY"`
	CamundaTaskPriority           int                             `xml:"camunda:taskPriority,attr,omitempty" json:"taskPriority,omitempty" csv:"TASK_PRIORITY"`
	CamundaCandidateStarterGroups string                          `xml:"camunda:candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty" csv:"CANDIDATE_STARTER_GROUPS"`
	CamundaCandidateStarterUsers  string                          `xml:"camunda:candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty" csv:"CANDIDATE_STARTER_USERS"`
	Documentation                 []attributes.Documentation      `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"-"`
	LaneSet                       []LaneSet                       `xml:"bpmn:laneSet,omitempty" json:"laneSet,omitempty" csv:"-"`
	StartEvent                    []events.StartEvent             `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty" csv:"-"`
	BoundaryEvent                 []events.BoundaryEvent          `xml:"bpmn:boundaryEvent,omitemnpty" json:"boundaryEvent,omitempty" csv:"-"`
	EndEvent                      []events.EndEvent               `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent        []events.IntermediateCatchEvent `xml:"bpmn:intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty" csv:"-"`
	IntermediateThrowEvent        []events.IntermediateThrowEvent `xml:"bpmn:intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty" csv:"-"`
	Task                          []activities.Task               `xml:"bpmn:task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask                      []activities.UserTask           `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask                    []activities.ManualTask         `xml:"bpmn:manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask                   []activities.ReceiveTask        `xml:"bpmn:receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask                    []activities.ScriptTask         `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask                      []activities.SendTask           `xml:"bpmn:sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask                   []activities.ServiceTask        `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
	CallActivity                  []subprocesses.CallActivity     `xml:"bpmn:callActivity,omitempty" json:"callActivity,omitempty" csv:"-"`
	SubProcess                    []subprocesses.SubProcess       `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty" csv:"-"`
	Transaction                   []subprocesses.Transaction      `xml:"bpmn:transaction,omitempty" json:"transaction,omitempty" csv:"-"`
	AdHocSubprocess               []subprocesses.AdHocSubProcess  `xml:"bpmn:adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty" csv:"-"`
	ExclusiveGateway              []gateways.ExclusiveGateway     `xml:"bpmn:exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty" csv:"-"`
	InclusiveGateway              []gateways.InclusiveGateway     `xml:"bpmn:inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty" csv:"-"`
	ParallelGateway               []gateways.ParallelGateway      `xml:"bpmn:parallelGateway,omitempty" json:"parallelGateway,omitempty" csv:"-"`
	ComplexGateway                []gateways.ComplexGateway       `xml:"bpmn:complexGateway,omitempty" json:"complexGateway,omitempty" csv:"-"`
	EventBasedGateway             []gateways.EventBasedGateway    `xml:"bpmn:eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty" csv:"-"`
	SequenceFlow                  []marker.SequenceFlow           `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty" csv:"-"`
	Group                         []Group                         `xml:"bpmn:group,omitempty" json:"group,omitempty" csv:"-"`
}

// TProcess ...
type TProcess struct {
	ID                     string                           `xml:"id,attr" json:"id"`
	Name                   string                           `xml:"name,attr,omitempty" json:"name,omitempty"`
	IsExecutable           bool                             `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	VersionTag             string                           `xml:"versionTag,attr,omitempty" json:"versionTag,omitempty"`
	JobPriority            int                              `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	TaskPriority           int                              `xml:"taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CandidateStarterGroups string                           `xml:"candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CandidateStarterUsers  string                           `xml:"candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	Documentation          []attributes.Documentation       `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LaneSet                []LaneSet                        `xml:"laneSet,omitempty" json:"laneSet,omitempty"`
	StartEvent             []events.TStartEvent             `xml:"startEvent,omitemnpty" json:"startEvent,omitempty"`
	BoundaryEvent          []events.TBoundaryEvent          `xml:"boundaryEvent,omitemnpty" json:"boundaryEvent,omitempty"`
	EndEvent               []events.TEndEvent               `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent []events.TIntermediateCatchEvent `xml:"intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty"`
	IntermediateThrowEvent []events.TIntermediateThrowEvent `xml:"intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty"`
	Task                   []activities.TTask               `xml:"task,omitempty" json:"task,omitempty"`
	UserTask               []activities.TUserTask           `xml:"userTask,omitempty" json:"userTask,omitempty"`
	ManualTask             []activities.TManualTask         `xml:"manualTask,omitempty" json:"manualTask,omitempty"`
	ReceiveTask            []activities.TReceiveTask        `xml:"receiveTask,omitempty" json:"receiveTask,omitempty"`
	ScriptTask             []activities.TScriptTask         `xml:"scriptTask,omitempty" json:"scriptTask,omitempty"`
	SendTask               []activities.TSendTask           `xml:"sendTask,omitempty" json:"sendTask,omitempty"`
	ServiceTask            []activities.TServiceTask        `xml:"serviceTask,omitempty" json:"serviceTask,omitempty"`
	CallActivity           []subprocesses.TCallActivity     `xml:"callActivity,omitempty" json:"callActivity,omitempty"`
	SubProcess             []subprocesses.TSubProcess       `xml:"subProcess,omitempty" json:"subProcess,omitempty"`
	Transaction            []subprocesses.TTransaction      `xml:"transaction,omitempty" json:"transaction,omitempty"`
	AdHocSubprocess        []subprocesses.TAdHocSubProcess  `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty" csv:"-"`
	ExclusiveGateway       []gateways.TExclusiveGateway     `xml:"exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
	InclusiveGateway       []gateways.TInclusiveGateway     `xml:"inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
	ParallelGateway        []gateways.TParallelGateway      `xml:"parallelGateway,omitempty" json:"parallelGateway,omitempty"`
	ComplexGateway         []gateways.TComplexGateway       `xml:"complexGateway,omitempty" json:"complexGateway,omitempty"`
	EventBasedGateway      []gateways.TEventBasedGateway    `xml:"eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty"`
	SequenceFlow           []marker.TSequenceFlow           `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
	Group                  []Group                          `xml:"group,omitempty" json:"group,omitempty"`
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
	process.LaneSet = make([]LaneSet, 1)
}

/*** Event ***/

// SetStartEvent ...
func (process *Process) SetStartEvent(num int) {
	process.StartEvent = make([]events.StartEvent, num)
}

// SetBoundEvent ...
func (process *Process) SetBoundaryEvent(num int) {
	process.BoundaryEvent = make([]events.BoundaryEvent, num)
}

// SetEndEvent ...
func (process *Process) SetEndEvent(num int) {
	process.EndEvent = make([]events.EndEvent, num)
}

// SetIntermedCatchEvent ...
func (process *Process) SetIntermediateCatchEvent(num int) {
	process.IntermediateCatchEvent = make([]events.IntermediateCatchEvent, num)
}

// SetIntermedThrowEvent ...
func (process *Process) SetIntermediateThrowEvent(num int) {
	process.IntermediateThrowEvent = make([]events.IntermediateThrowEvent, num)
}

/*** Task ***/

// SetTask ...
func (process *Process) SetTask(num int) {
	process.Task = make([]activities.Task, num)
}

// SetUserTask ...
func (process *Process) SetUserTask(num int) {
	process.UserTask = make([]activities.UserTask, num)
}

// SetManualTask ...
func (process *Process) SetManualTask(num int) {
	process.ManualTask = make([]activities.ManualTask, num)
}

// SetReceiveTask ...
func (process *Process) SetReceiveTask(num int) {
	process.ReceiveTask = make([]activities.ReceiveTask, num)
}

// SetScriptTask ...
func (process *Process) SetScriptTask(num int) {
	process.ScriptTask = make([]activities.ScriptTask, num)
}

// SetSendTask ...
func (process *Process) SetSendTask(num int) {
	process.SendTask = make([]activities.SendTask, num)
}

// SetServiceTask ...
func (process *Process) SetServiceTask(num int) {
	process.ServiceTask = make([]activities.ServiceTask, num)
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

// SetSequenceFlow ...
func (process *Process) SetSequenceFlow(num int) {
	process.SequenceFlow = make([]marker.SequenceFlow, num)
}

// SetGroup ...
func (process *Process) SetGroup(num int) {
	process.Group = make([]Group, num)
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
func (process Process) GetLaneSet() *LaneSet {
	return &process.LaneSet[0]
}

/*** Event ***/

// GetStartEvent ...
func (process Process) GetStartEvent(num int) *events.StartEvent {
	return &process.StartEvent[num]
}

// GetBoundaryEvent ...
func (process Process) GetBoundaryEvent(num int) *events.BoundaryEvent {
	return &process.BoundaryEvent[num]
}

// GetEndEvent ...
func (process Process) GetEndEvent(num int) *events.EndEvent {
	return &process.EndEvent[num]
}

// GetIntermedCatchEvent ...
func (process Process) GetIntermediateCatchEvent(num int) *events.IntermediateCatchEvent {
	return &process.IntermediateCatchEvent[num]
}

// GetIntermedThrowEvent ...
func (process Process) GetIntermediateThrowEvent(num int) *events.IntermediateThrowEvent {
	return &process.IntermediateThrowEvent[num]
}

/*** Task ***/

// GetTask ...
func (process Process) GetTask(num int) *activities.Task {
	return &process.Task[num]
}

// GetUserTask ...
func (process Process) GetUserTask(num int) *activities.UserTask {
	return &process.UserTask[num]
}

// GetManualTask ...
func (process Process) GetManualTask(num int) *activities.ManualTask {
	return &process.ManualTask[num]
}

// GetReceiveTask ...
func (process Process) GetReceiveTask(num int) *activities.ReceiveTask {
	return &process.ReceiveTask[num]
}

// GetScriptTask ...
func (process Process) GetScriptTask(num int) *activities.ScriptTask {
	return &process.ScriptTask[num]
}

// GetSendTask ...
func (process Process) GetSendTask(num int) *activities.SendTask {
	return &process.SendTask[num]
}

// GetServiceTask ...
func (process Process) GetServiceTask(num int) *activities.ServiceTask {
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

// GetSequenceFlow ...
func (process Process) GetSequenceFlow(num int) *marker.SequenceFlow {
	return &process.SequenceFlow[num]
}

// GetGroup ...
func (process Process) GetGroup(num int) *Group {
	return &process.Group[num]
}
