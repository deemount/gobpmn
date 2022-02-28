package models

import "fmt"

// ReceiveTaskRepository ...
type ReceiveTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ReceiveTask ...
type ReceiveTask struct {
	ID                 string              `xml:"id,attr" json:"id,omitempty"`
	Name               string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	MessageRef         string              `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
	CamundaAsyncBefore bool                `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,attr,omitempty"`
	CamundaAsyncAfter  bool                `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,attr,omitempty"`
	CamundaJobPriority int                 `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,attr,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (receiveTask *ReceiveTask) SetID(suffix string) {
	receiveTask.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (receiveTask *ReceiveTask) SetName(name string) {
	receiveTask.Name = name
}

// SetMessageRef ...
func (receiveTask *ReceiveTask) SetMessageRef(suffix string) {
	receiveTask.MessageRef = fmt.Sprintf("Message_%s", suffix)
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (receiveTask *ReceiveTask) SetCamundaAsyncBefore(asyncBefore bool) {
	receiveTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (receiveTask *ReceiveTask) SetCamundaAsyncAfter(asyncAfter bool) {
	receiveTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (receiveTask *ReceiveTask) SetCamundaJobPriority(priority int) {
	receiveTask.CamundaJobPriority = priority
}

/* Elements */

// SetExtensionElements ...
func (receiveTask *ReceiveTask) SetExtensionElements() {
	receiveTask.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (receiveTask *ReceiveTask) SetIncoming(num int) {
	receiveTask.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (receiveTask *ReceiveTask) SetOutgoing(num int) {
	receiveTask.Outgoing = make([]Outgoing, num)
}
