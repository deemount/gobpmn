package models

import "fmt"

// SendTaskRepository ...
type SendTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// SendTask ...
type SendTask struct {
	ID                 string              `xml:"id,attr,omitempty" json:"id"`
	Name               string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                 `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (sendTask *SendTask) SetID(suffix string) {
	sendTask.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (sendTask *SendTask) SetName(name string) {
	sendTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (sendTask *SendTask) SetCamundaAsyncBefore(asyncBefore bool) {
	sendTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (sendTask *SendTask) SetCamundaAsyncAfter(asyncAfter bool) {
	sendTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (sendTask *SendTask) SetCamundaJobPriority(priority int) {
	sendTask.CamundaJobPriority = priority
}

/* Elements */

// SetDocumentation ...
func (sendTask *SendTask) SetDocumentation() {
	sendTask.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (sendTask *SendTask) SetExtensionElements() {
	sendTask.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (sendTask *SendTask) SetIncoming(num int) {
	sendTask.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (sendTask *SendTask) SetOutgoing(num int) {
	sendTask.Outgoing = make([]Outgoing, num)
}
