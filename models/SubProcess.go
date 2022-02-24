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

// SetCamundaAsyncAfter ...

// SetCamundaJobPriority ...

/* Elements */

// SetStartEvent ...
func (subprocess *SubProcess) SetStartEvent(num int) {
	subprocess.StartEvent = make([]StartEvent, num)
}

// SetTask ...
func (subprocess *SubProcess) SetTask(num int) {
	subprocess.Task = make([]Task, num)
}

// SetUserTask ...
func (subprocess *SubProcess) SetUserTask(num int) {
	subprocess.UserTask = make([]UserTask, num)
}

// SetEndEvent ...
func (subprocess *SubProcess) SetEndEvent(num int) {
	subprocess.EndEvent = make([]EndEvent, num)
}

// SetSequenceFlow ...
func (subprocess *SubProcess) SetSequenceFlow(num int) {
	subprocess.SequenceFlow = make([]SequenceFlow, num)
}
