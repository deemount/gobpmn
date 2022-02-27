package models

// ProcessRepository ...
type ProcessRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// Process ...
type Process struct {
	ID                       string                   `xml:"id,attr"`
	Name                     string                   `xml:"name,attr,omitempty"`
	IsExecutable             bool                     `xml:"isExecutable,attr"`
	CamundaVersionTag        string                   `xml:"camunda:versionTag,attr,omitempty"`
	CamundaJobPriority       int                      `xml:"camunda:jobPriority,attr,omitempty"`
	CamundaTaskPriority      int                      `xml:"camunda:taskPriority,attr,omitempty"`
	CamundaCandidStartGroups string                   `xml:"camunda:candidateStarterGroups,attr,omitempty"`
	CamundaCandidStartUsers  string                   `xml:"camunda:candidateStarterUsers,attr,omitempty"`
	LaneSet                  []LaneSet                `xml:"bpmn:laneSet,omitempty"`
	StartEvent               []StartEvent             `xml:"bpmn:startEvent,omitemnpty"`
	BoundaryEvent            []BoundaryEvent          `xml:"bpmn:boundaryEvent,omitemnpty"`
	EndEvent                 []EndEvent               `xml:"bpmn:endEvent,omitempty"`
	IntermediateCatchEvent   []IntermediateCatchEvent `xml:"bpmn:intermediateCatchEvent,omitempty"`
	IntermediateThrowEvent   []IntermediateThrowEvent `xml:"bpmn:intermediateThrowEvent,omitempty"`
	Task                     []Task                   `xml:"bpmn:task,omitempty"`
	UserTask                 []UserTask               `xml:"bpmn:userTask,omitempty"`
	ManualTask               []ManualTask             `xml:"bpmn:manualTask,omitempty"`
	ReceiveTask              []ReceiveTask            `xml:"bpmn:receiveTask,omitempty"`
	ScriptTask               []ScriptTask             `xml:"bpmn:scriptTask,omitempty"`
	SendTask                 []SendTask               `xml:"bpmn:sendTask,omitempty"`
	ServiceTask              []ServiceTask            `xml:"bpmn:serviceTask,omitempty"`
	CallActivity             []CallActivity           `xml:"bpmn:callActivity,omitempty"`
	SubProcess               []SubProcess             `xml:"bpmn:subProcess,omitempty"`
	Transaction              []Transaction            `xml:"bpmn:transaction,omitempty"`
	ExclusiveGateway         []ExclusiveGateway       `xml:"bpmn:exclusiveGateway,omitempty"`
	InclusiveGateway         []InclusiveGateway       `xml:"bpmn:inclusiveGateway,omitempty"`
	ParallelGateway          []ParallelGateway        `xml:"bpmn:parallelGateway,omitempty"`
	ComplexGateway           []ComplexGateway         `xml:"bpmn:complexGateway,omitempty"`
	EventBasedGateway        []EventBasedGateway      `xml:"bpmn:eventBasedGateway,omitempty"`
	SequenceFlow             []SequenceFlow           `xml:"bpmn:sequenceFlow,omitempty"`
	Group                    []Group                  `xml:"bpmn:group,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (process *Process) SetID(suffix string) {
	process.ID = "Process_" + suffix
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
func (process *Process) SetIntermedCatchEvent(num int) {
	process.IntermediateCatchEvent = make([]IntermediateCatchEvent, num)
}

// SetIntermedThrowEvent ...
func (process *Process) SetIntermedThrowEvent(num int) {
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
