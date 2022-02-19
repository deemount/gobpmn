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

// SetID ...
func (sproc *SubProcess) SetID(suffix string) {
	sproc.ID = "Activity_" + suffix
}

// SetTriggeredByEvent ...
func (sproc *SubProcess) SetTriggeredByEvent(triggered bool) {
	sproc.TriggeredByEvent = triggered
}
