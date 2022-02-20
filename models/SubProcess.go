package models

// SubProcess ...
type SubProcess struct {
	ID               string         `xml:"id,attr"`
	TriggeredByEvent bool           `xml:"triggeredByEvent,attr,omitempty"`
	StartEvent       []StartEvent   `xml:"bpmn:startEvent,omitemnpty"`
	Task             []Task         `xml:"bpmn:task,omitempty"`
	UserTask         []UserTask     `xml:"bpm:userTask,omitempty"`
	EndEvent         []EndEvent     `xml:"bpmn:endEvent,omitempty"`
	SequenceFlow     []SequenceFlow `xml:"bpmn:sequenceFlow,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (subprocess *SubProcess) SetID(suffix string) {
	subprocess.ID = "Activity_" + suffix
}

// SetTriggeredByEvent ...
func (subprocess *SubProcess) SetTriggeredByEvent(triggered bool) {
	subprocess.TriggeredByEvent = triggered
}

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
