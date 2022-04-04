package models

import "fmt"

// ProcessRepository ...
type ProcessRepository interface {
	SetID(suffix string)
	SetName(name string)
	GetID() string
}

// Process ...
type Process struct {
	ID                       string                   `xml:"id,attr" json:"id" csv:"ID"`
	Name                     string                   `xml:"name,attr,omitempty" json:"name,omitempty" csv:"NAME"`
	IsExecutable             bool                     `xml:"isExecutable,attr" json:"isExecutable,omitempty" csv:"IS_EXECUTABLE"`
	CamundaVersionTag        string                   `xml:"camunda:versionTag,attr,omitempty" json:"versionTag,omitempty" csv:"VERSION_TAG"`
	CamundaJobPriority       int                      `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty" csv:"JOB_PRIORITY"`
	CamundaTaskPriority      int                      `xml:"camunda:taskPriority,attr,omitempty" json:"taskPriority,omitempty" csv:"TASK_PRIORITY"`
	CamundaCandidStartGroups string                   `xml:"camunda:candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty" csv:"CANDIDATE_STARTER_GROUPS"`
	CamundaCandidStartUsers  string                   `xml:"camunda:candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty" csv:"CANDIDATE_STARTER_USERS"`
	Documentation            []Documentation          `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"-"`
	LaneSet                  []LaneSet                `xml:"bpmn:laneSet,omitempty" json:"laneSet,omitempty" csv:"-"`
	StartEvent               []StartEvent             `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty" csv:"-"`
	BoundaryEvent            []BoundaryEvent          `xml:"bpmn:boundaryEvent,omitemnpty" json:"boundaryEvent,omitempty" csv:"-"`
	EndEvent                 []EndEvent               `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent   []IntermediateCatchEvent `xml:"bpmn:intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty" csv:"-"`
	IntermediateThrowEvent   []IntermediateThrowEvent `xml:"bpmn:intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty" csv:"-"`
	Task                     []Task                   `xml:"bpmn:task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask                 []UserTask               `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask               []ManualTask             `xml:"bpmn:manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask              []ReceiveTask            `xml:"bpmn:receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask               []ScriptTask             `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask                 []SendTask               `xml:"bpmn:sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask              []ServiceTask            `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
	CallActivity             []CallActivity           `xml:"bpmn:callActivity,omitempty" json:"callActivity,omitempty" csv:"-"`
	SubProcess               []SubProcess             `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty" csv:"-"`
	Transaction              []Transaction            `xml:"bpmn:transaction,omitempty" json:"transaction,omitempty" csv:"-"`
	ExclusiveGateway         []ExclusiveGateway       `xml:"bpmn:exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty" csv:"-"`
	InclusiveGateway         []InclusiveGateway       `xml:"bpmn:inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty" csv:"-"`
	ParallelGateway          []ParallelGateway        `xml:"bpmn:parallelGateway,omitempty" json:"parallelGateway,omitempty" csv:"-"`
	ComplexGateway           []ComplexGateway         `xml:"bpmn:complexGateway,omitempty" json:"complexGateway,omitempty" csv:"-"`
	EventBasedGateway        []EventBasedGateway      `xml:"bpmn:eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty" csv:"-"`
	SequenceFlow             []SequenceFlow           `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty" csv:"-"`
	Group                    []Group                  `xml:"bpmn:group,omitempty" json:"group,omitempty" csv:"-"`
}

// TProcess ...
type TProcess struct {
	ID                     string                    `xml:"id,attr" json:"id"`
	Name                   string                    `xml:"name,attr,omitempty" json:"name,omitempty"`
	IsExecutable           bool                      `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	VersionTag             string                    `xml:"versionTag,attr,omitempty" json:"versionTag,omitempty"`
	JobPriority            int                       `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	TaskPriority           int                       `xml:"taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CandidStartGroups      string                    `xml:"candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CandidStartUsers       string                    `xml:"candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	Documentation          []Documentation           `xml:"documentation,omitempty" json:"documentation,omitempty"`
	LaneSet                []LaneSet                 `xml:"laneSet,omitempty" json:"laneSet,omitempty"`
	StartEvent             []TStartEvent             `xml:"startEvent,omitemnpty" json:"startEvent,omitempty"`
	BoundaryEvent          []BoundaryEvent           `xml:"boundaryEvent,omitemnpty" json:"boundaryEvent,omitempty"`
	EndEvent               []TEndEvent               `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent []TIntermediateCatchEvent `xml:"intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty"`
	IntermediateThrowEvent []IntermediateThrowEvent  `xml:"intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty"`
	Task                   []TTask                   `xml:"task,omitempty" json:"task,omitempty"`
	UserTask               []UserTask                `xml:"userTask,omitempty" json:"userTask,omitempty"`
	ManualTask             []ManualTask              `xml:"manualTask,omitempty" json:"manualTask,omitempty"`
	ReceiveTask            []ReceiveTask             `xml:"receiveTask,omitempty" json:"receiveTask,omitempty"`
	ScriptTask             []ScriptTask              `xml:"scriptTask,omitempty" json:"scriptTask,omitempty"`
	SendTask               []SendTask                `xml:"sendTask,omitempty" json:"sendTask,omitempty"`
	ServiceTask            []ServiceTask             `xml:"serviceTask,omitempty" json:"serviceTask,omitempty"`
	CallActivity           []CallActivity            `xml:"callActivity,omitempty" json:"callActivity,omitempty"`
	SubProcess             []SubProcess              `xml:"subProcess,omitempty" json:"subProcess,omitempty"`
	Transaction            []Transaction             `xml:"transaction,omitempty" json:"transaction,omitempty"`
	ExclusiveGateway       []ExclusiveGateway        `xml:"exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
	InclusiveGateway       []InclusiveGateway        `xml:"inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
	ParallelGateway        []ParallelGateway         `xml:"parallelGateway,omitempty" json:"parallelGateway,omitempty"`
	ComplexGateway         []ComplexGateway          `xml:"complexGateway,omitempty" json:"complexGateway,omitempty"`
	EventBasedGateway      []EventBasedGateway       `xml:"eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty"`
	SequenceFlow           []TSequenceFlow           `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
	Group                  []Group                   `xml:"group,omitempty" json:"group,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (process *Process) SetID(typ string, suffix string) {
	switch typ {
	case "hash":
		process.ID = fmt.Sprintf("Process_%s", suffix)
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

// SetCamundaCandidStartGroups ...
func (process *Process) SetCamundaCandidStartGroups(groups string) {
	process.CamundaCandidStartGroups = groups
}

// SetCamundaCandidStartUsers
func (process *Process) SetCamundaCandidStartUsers(users string) {
	process.CamundaCandidStartUsers = users
}

/* Elements */

/** BPMN **/

/** Documentation **/

// SetDocumentation ...
func (process *Process) SetDocumentation() {
	process.Documentation = make([]Documentation, 1)
}

/** LaneSet **/

// SetLaneSet ...
func (process *Process) SetLaneSet() {
	process.LaneSet = make([]LaneSet, 1)
}

/*** Event ***/

// SetStartEvent ...
func (process *Process) SetStartEvent(num int) {
	process.StartEvent = make([]StartEvent, num)
}

// SetBoundEvent ...
func (process *Process) SetBoundaryEvent(num int) {
	process.BoundaryEvent = make([]BoundaryEvent, num)
}

// SetEndEvent ...
func (process *Process) SetEndEvent(num int) {
	process.EndEvent = make([]EndEvent, num)
}

// SetIntermedCatchEvent ...
func (process *Process) SetIntermediateCatchEvent(num int) {
	process.IntermediateCatchEvent = make([]IntermediateCatchEvent, num)
}

// SetIntermedThrowEvent ...
func (process *Process) SetIntermediateThrowEvent(num int) {
	process.IntermediateThrowEvent = make([]IntermediateThrowEvent, num)
}

/*** Task ***/

// SetTask ...
func (process *Process) SetTask(num int) {
	process.Task = make([]Task, num)
}

// SetUserTask ...
func (process *Process) SetUserTask(num int) {
	process.UserTask = make([]UserTask, num)
}

// SetManualTask ...
func (process *Process) SetManualTask(num int) {
	process.ManualTask = make([]ManualTask, num)
}

// SetReceiveTask ...
func (process *Process) SetReceiveTask(num int) {
	process.ReceiveTask = make([]ReceiveTask, num)
}

// SetScriptTask ...
func (process *Process) SetScriptTask(num int) {
	process.ScriptTask = make([]ScriptTask, num)
}

// SetSendTask ...
func (process *Process) SetSendTask(num int) {
	process.SendTask = make([]SendTask, num)
}

// SetServiceTask ...
func (process *Process) SetServiceTask(num int) {
	process.ServiceTask = make([]ServiceTask, num)
}

// SetCallActivity ...
func (process *Process) SetCallActivity(num int) {
	process.CallActivity = make([]CallActivity, num)
}

// SetSubProcess ...
func (process *Process) SetSubProcess(num int) {
	process.SubProcess = make([]SubProcess, num)
}

// SetTransaction ...
func (process *Process) SetTransaction(num int) {
	process.Transaction = make([]Transaction, num)
}

/*** Gateway ***/

// SetExclusiveGateway
func (process *Process) SetExclusiveGateway(num int) {
	process.ExclusiveGateway = make([]ExclusiveGateway, num)
}

// SetInclsuiveGateway
func (process *Process) SetInclusiveGateway(num int) {
	process.InclusiveGateway = make([]InclusiveGateway, num)
}

// SetParallelGateway
func (process *Process) SetParallelGateway(num int) {
	process.ParallelGateway = make([]ParallelGateway, num)
}

// SetComplexGateway
func (process *Process) SetComplexGateway(num int) {
	process.ComplexGateway = make([]ComplexGateway, num)
}

// SetEventBasedGateway
func (process *Process) SetEventBasedGateway(num int) {
	process.EventBasedGateway = make([]EventBasedGateway, num)
}

/*** Marker ***/

// SetSequenceFlow ...
func (process *Process) SetSequenceFlow(num int) {
	process.SequenceFlow = make([]SequenceFlow, num)
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
func (process Process) GetID() string {
	return process.ID
}
