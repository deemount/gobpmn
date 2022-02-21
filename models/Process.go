package models

// Process ...
type Process struct {
	ID                       string       `xml:"id,attr"`
	Name                     string       `xml:"name,attr,omitempty"`
	IsExecutable             bool         `xml:"isExecutable,attr"`
	CamundaVersionTag        string       `xml:"camunda:versionTag,attr,omitempty"`
	CamundaJobPriority       int          `xml:"camunda:jobPriority,attr,omitempty"`
	CamundaTaskPriority      int          `xml:"camunda:taskPriority,attr,omitempty"`
	CamundaCandidStartGroups string       `xml:"camunda:candidateStarterGroups,attr,omitempty"`
	CamundaCandidStartUsers  string       `xml:"camunda:candidateStarterUsers,attr,omitempty"`
	StartEvent               []StartEvent `xml:"bpmn:startEvent,omitemnpty"`
	//BoundEvent               []BoundaryEvent          `xml:"bpmn:boundaryEvent,omitemnpty"`
	EndEvent               []EndEvent               `xml:"bpmn:endEvent,omitempty"`
	IntermediateCatchEvent []IntermediateCatchEvent `xml:"bpmn:intermediateCatchEvent,omitempty"`
	IntermediateThrowEvent []IntermediateThrowEvent `xml:"bpmn:intermediateThrowEvent,omitempty"`
	Task                   []Task                   `xml:"bpmn:task,omitempty"`
	UserTask               []UserTask               `xml:"bpmn:userTask,omitempty"`
	ManualTask             []ManualTask             `xml:"bpmn:manualTask,omitempty"`
	ReceiveTask            []ReceiveTask            `xml:"bpmn:receiveTask,omitempty"`
	ScriptTask             []ScriptTask             `xml:"bpmn:scriptTask,omitempty"`
	SendTask               []SendTask               `xml:"bpmn:sendTask,omitempty"`
	ServiceTask            []ServiceTask            `xml:"bpmn:serviceTask,omitempty"`
	ExclusiveGate          []ExclusiveGateway       `xml:"bpmn:exclusiveGateway,omitempty"`
	InclusiveGate          []InclusiveGateway       `xml:"bpmn:inclusiveGateway,omitempty"`
	ParallelGate           []ParallelGateway        `xml:"bpmn:parallelGateway,omitempty"`
	ComplexGate            []ComplexGateway         `xml:"bpmn:complexGateway,omitempty"`
	EventBasedGate         []EventBasedGateway      `xml:"bpmn:eventBasedGateway,omitempty"`
	SubProcess             []SubProcess             `xml:"bpmn:subProcess,omitempty"`
	SequenceFlow           []SequenceFlow           `xml:"bpmn:sequenceFlow,omitempty"`
	//CallActivity             []CallActivity           `xml:"bpmn:callActivity,omitempty"`
	//Group                    []Group                  `xml:"bpmn:group,omitempty"`
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

// SetStartEvent ...
func (proc *Process) SetStartEvent(num int) {
	proc.StartEvent = make([]StartEvent, num)
}

// SetIntermedCatchEvent ...
func (proc *Process) SetIntermedCatchEvent(num int) {
	proc.IntermediateCatchEvent = make([]IntermediateCatchEvent, num)
}

// SetIntermedThrowEvent ...
func (proc *Process) SetIntermedThrowEvent(num int) {
	proc.IntermediateThrowEvent = make([]IntermediateThrowEvent, num)
}

// SetTask ...
func (proc *Process) SetTask(num int) {
	proc.Task = make([]Task, num)
}

// SetUserTask ...
func (proc *Process) SetUserTask(num int) {
	proc.UserTask = make([]UserTask, num)
}

/*
// SetBoundEvent ...
func (proc *Process) SetBoundEvent(num int) {
	proc.BoundEvent = make([]BoundaryEvent, num)
}
*/

// SetEndEvent ...
func (proc *Process) SetEndEvent(num int) {
	proc.EndEvent = make([]EndEvent, num)
}

/*
// SetExclGate
func (proc *Process) SetExclGate(num int) {
	proc.ExclGate = make([]ExclusiveGateway, num)
}
*/

// SetSubProcess ...
func (proc *Process) SetSubProcess(num int) {
	proc.SubProcess = make([]SubProcess, num)
}

// SetSequenceFlow ...
func (proc *Process) SetSequenceFlow(num int) {
	proc.SequenceFlow = make([]SequenceFlow, num)
}

/*

// SetCallActivity ...
func (proc *Process) SetCallActivity(num int) {
	proc.CallActivity = make([]CallActivity, num)
}

// SetGroup ...
func (proc *Process) SetGroup(num int) {
	proc.Group = make([]Group, num)
}
*/
