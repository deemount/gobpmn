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
func (proc *Process) SetID(suffix string) {
	proc.ID = "Process_" + suffix
}

// SetName ...
func (proc *Process) SetName(name string) {
	proc.Name = name
}

// SetIsExecutable ...
func (proc *Process) SetIsExecutable(isExec bool) {
	proc.IsExecutable = isExec
}

/** Camunda **/

// SetCamundaVersionTag ...
func (proc *Process) SetCamundaVersionTag(tag string) {
	proc.CamundaVersionTag = tag
}

// SetCamundaJobpriority ...
func (proc *Process) SetCamundaJobPriority(priority int) {
	proc.CamundaJobPriority = priority
}

// SetCamundaTaskPriority ...
func (proc *Process) SetCamundaTaskPriority(priority int) {
	proc.CamundaTaskPriority = priority
}

// SetCamundaCandidStartGroups ...
func (proc *Process) SetCamundaCandidStartGroups(groups string) {
	proc.CamundaCandidStartGroups = groups
}

// SetCamundaCandidStartUsers
func (proc *Process) SetCamundaCandidStartUsers(users string) {
	proc.CamundaCandidStartUsers = users
}

/* Elements */

/** BPMN **/

/** LaneSet **/

// SetLaneSet ...
func (proc *Process) SetLaneSet() {
	proc.LaneSet = make([]LaneSet, 1)
}

/*** Event ***/

// SetStartEvent ...
func (proc *Process) SetStartEvent(num int) {
	proc.StartEvent = make([]StartEvent, num)
}

// SetBoundEvent ...
func (proc *Process) SetBoundaryEvent(num int) {
	proc.BoundaryEvent = make([]BoundaryEvent, num)
}

// SetEndEvent ...
func (proc *Process) SetEndEvent(num int) {
	proc.EndEvent = make([]EndEvent, num)
}

// SetIntermedCatchEvent ...
func (proc *Process) SetIntermedCatchEvent(num int) {
	proc.IntermediateCatchEvent = make([]IntermediateCatchEvent, num)
}

// SetIntermedThrowEvent ...
func (proc *Process) SetIntermedThrowEvent(num int) {
	proc.IntermediateThrowEvent = make([]IntermediateThrowEvent, num)
}

/*** Task ***/

// SetTask ...
func (proc *Process) SetTask(num int) {
	proc.Task = make([]Task, num)
}

// SetUserTask ...
func (proc *Process) SetUserTask(num int) {
	proc.UserTask = make([]UserTask, num)
}

// SetManualTask ...
func (proc *Process) SetManualTask(num int) {
	proc.ManualTask = make([]ManualTask, num)
}

// SetReceiveTask ...
func (proc *Process) SetReceiveTask(num int) {
	proc.ReceiveTask = make([]ReceiveTask, num)
}

// SetScriptTask ...
func (proc *Process) SetScriptTask(num int) {
	proc.ScriptTask = make([]ScriptTask, num)
}

// SetSendTask ...
func (proc *Process) SetSendTask(num int) {
	proc.SendTask = make([]SendTask, num)
}

// SetServiceTask ...
func (proc *Process) SetServiceTask(num int) {
	proc.ServiceTask = make([]ServiceTask, num)
}

// SetCallActivity ...
func (proc *Process) SetCallActivity(num int) {
	proc.CallActivity = make([]CallActivity, num)
}

// SetSubProcess ...
func (proc *Process) SetSubProcess(num int) {
	proc.SubProcess = make([]SubProcess, num)
}

// SetTransaction ...
func (proc *Process) SetTransaction(num int) {
	proc.Transaction = make([]Transaction, num)
}

/*** Gateway ***/

// SetExclusiveGateway
func (proc *Process) SetExclusiveGateway(num int) {
	proc.ExclusiveGateway = make([]ExclusiveGateway, num)
}

// SetInclsuiveGateway
func (proc *Process) SetInclusiveGateway(num int) {
	proc.InclusiveGateway = make([]InclusiveGateway, num)
}

// SetParallelGateway
func (proc *Process) SetParallelGateway(num int) {
	proc.ParallelGateway = make([]ParallelGateway, num)
}

// SetComplexGateway
func (proc *Process) SetComplexGateway(num int) {
	proc.ComplexGateway = make([]ComplexGateway, num)
}

// SetEventBasedGateway
func (proc *Process) SetEventBasedGateway(num int) {
	proc.EventBasedGateway = make([]EventBasedGateway, num)
}

/*** Marker ***/

// SetSequenceFlow ...
func (proc *Process) SetSequenceFlow(num int) {
	proc.SequenceFlow = make([]SequenceFlow, num)
}

// SetGroup ...
func (proc *Process) SetGroup(num int) {
	proc.Group = make([]Group, num)
}
