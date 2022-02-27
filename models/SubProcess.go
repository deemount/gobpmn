package models

// SubProcess ...
type SubProcess struct {
	ID                 string              `xml:"id,attr"`
	Name               string              `xml:"name,attr,omitempty"`
	TriggeredByEvent   bool                `xml:"triggeredByEvent,attr,omitempty"`
	CamundaAsyncBefore bool                `xml:"camunda:asyncBefore,attr,omitempty"`
	CamundaAsyncAfter  bool                `xml:"camunda:asyncAfter,attr,omitempty"`
	CamundaJobPriority int                 `xml:"camunda:jobPriority,attr,omitempty"`
	Documentation      []Documentation     `xml:"bpmn:documentation,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
	StartEvent         []StartEvent        `xml:"bpmn:startEvent,omitemnpty"`
	EndEvent           []EndEvent          `xml:"bpmn:endEvent,omitempty"`
	Task               []Task              `xml:"bpmn:task,omitempty"`
	UserTask           []UserTask          `xml:"bpmn:userTask,omitempty"`
	ManualTask         []ManualTask        `xml:"bpmn:manualTask,omitempty"`
	ReceiveTask        []ReceiveTask       `xml:"bpmn:receiveTask,omitempty"`
	ScriptTask         []ScriptTask        `xml:"bpmn:scriptTask,omitempty"`
	SendTask           []SendTask          `xml:"bpmn:sendTask,omitempty"`
	ServiceTask        []ServiceTask       `xml:"bpmn:serviceTask,omitempty"`
	SubProcess         []SubProcess        `xml:"bpmn:subProcess,omitempty"` // is that possible ?
	ExclusiveGateway   []ExclusiveGateway  `xml:"bpmn:exclusiveGateway,omitempty"`
	InclusiveGateway   []InclusiveGateway  `xml:"bpmn:inclusiveGateway,omitempty"`
	ParallelGateway    []ParallelGateway   `xml:"bpmn:parallelGateway,omitempty"`
	ComplexGateway     []ComplexGateway    `xml:"bpmn:complexGateway,omitempty"`
	EventBasedGateway  []EventBasedGateway `xml:"bpmn:eventBasedGateway,omitempty"`
	SequenceFlow       []SequenceFlow      `xml:"bpmn:sequenceFlow,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (subprocess *SubProcess) SetID(suffix string) {
	subprocess.ID = "Activity_" + suffix
}

// SetName ...
func (subprocess *SubProcess) SetName(name string) {
	subprocess.Name = name
}

// SetTriggeredByEvent ...
func (subprocess *SubProcess) SetTriggeredByEvent(triggered bool) {
	subprocess.TriggeredByEvent = triggered
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (subprocess *SubProcess) SetCamundaAsyncBefore(asyncBefore bool) {
	subprocess.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (subprocess *SubProcess) SetCamundaAsyncAfter(asyncAfter bool) {
	subprocess.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (subprocess *SubProcess) SetCamundaJobPriority(priority int) {
	subprocess.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (subprocess *SubProcess) SetDocumentation() {
	subprocess.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (subprocess *SubProcess) SetExtensionElements() {
	subprocess.ExtensionElements = make([]ExtensionElements, 1)
}

/*** Event ***/

// SetStartEvent ...
func (subprocess *SubProcess) SetStartEvent(num int) {
	subprocess.StartEvent = make([]StartEvent, num)
}

// SetEndEvent ...
func (subprocess *SubProcess) SetEndEvent(num int) {
	subprocess.EndEvent = make([]EndEvent, num)
}

/*** Task ***/

// SetTask ...
func (subprocess *SubProcess) SetTask(num int) {
	subprocess.Task = make([]Task, num)
}

// SetUserTask ...
func (subprocess *SubProcess) SetUserTask(num int) {
	subprocess.UserTask = make([]UserTask, num)
}

// SetManualTask ...
func (subprocess *SubProcess) SetManualTask(num int) {
	subprocess.ManualTask = make([]ManualTask, num)
}

// SetReceiveTask ...
func (subprocess *SubProcess) SetReceiveTask(num int) {
	subprocess.ReceiveTask = make([]ReceiveTask, num)
}

// SetScriptTask ...
func (subprocess *SubProcess) SetScriptTask(num int) {
	subprocess.ScriptTask = make([]ScriptTask, num)
}

// SetSendTask ...
func (subprocess *SubProcess) SetSendTask(num int) {
	subprocess.SendTask = make([]SendTask, num)
}

// SetServiceTask ...
func (subprocess *SubProcess) SetServiceTask(num int) {
	subprocess.ServiceTask = make([]ServiceTask, num)
}

/*** Gateway ***/

// SetExclusiveGateway
func (subprocess *SubProcess) SetExclusiveGateway(num int) {
	subprocess.ExclusiveGateway = make([]ExclusiveGateway, num)
}

// SetInclsuiveGateway
func (subprocess *SubProcess) SetInclusiveGateway(num int) {
	subprocess.InclusiveGateway = make([]InclusiveGateway, num)
}

// SetParallelGateway
func (subprocess *SubProcess) SetParallelGateway(num int) {
	subprocess.ParallelGateway = make([]ParallelGateway, num)
}

// SetComplexGateway
func (subprocess *SubProcess) SetComplexGateway(num int) {
	subprocess.ComplexGateway = make([]ComplexGateway, num)
}

// SetEventBasedGateway
func (subprocess *SubProcess) SetEventBasedGateway(num int) {
	subprocess.EventBasedGateway = make([]EventBasedGateway, num)
}

/*** Marker ***/

// SetSequenceFlow ...
func (subprocess *SubProcess) SetSequenceFlow(num int) {
	subprocess.SequenceFlow = make([]SequenceFlow, num)
}
