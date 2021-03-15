package models

import "strconv"

// StartEvent ...
type StartEvent struct {
	ID             string   `xml:"id,attr"`
	CamundaFormKey string   `xml:"camunda:formKey,attr"`
	Outgoing       Outgoing `xml:"bpmn:outgoing"`
}

// SetID ...
func (stev *StartEvent) SetID(num int64) {
	stev.ID = "StartEvent_" + strconv.FormatInt(num, 16)
}

// SetCamundaFormKey ...
func (stev *StartEvent) SetCamundaFormKey(name string) {
	stev.CamundaFormKey = name
}
