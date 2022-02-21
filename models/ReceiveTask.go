package models

import "fmt"

// ReceiveTask ...
type ReceiveTask struct {
	ID                string              `xml:"id,attr,omitempty"`
	Name              string              `xml:"name,attr,omitempty"`
	MessageRef        string              `xml:"messageRef,attr,omitempty"`
	CamundaAsyncBef   bool                `xml:"camunda:asyncBefore,attr,omitempty"`
	CamundaAsyncAft   bool                `xml:"camunda:asyncAfter,attr,omitempty"`
	CamundaJobPrio    int                 `xml:"camunda:jobPriority,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
	Incoming          []Incoming          `xml:"bpmn:incoming,omitempty"`
	Outgoing          []Outgoing          `xml:"bpmn:outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (rtask *ReceiveTask) SetID(suffix string) {
	rtask.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (rtask *ReceiveTask) SetName(name string) {
	rtask.Name = name
}

// SetMessageRef ...
func (rtask *ReceiveTask) SetMessageRef(suffix string) {
	rtask.MessageRef = fmt.Sprintf("Message_%s", suffix)
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (rtask *ReceiveTask) SetCamundaAsyncBefore(asyncBefore bool) {
	rtask.CamundaAsyncBef = asyncBefore
}

// SetCamundaAsyncAfter ...
func (rtask *ReceiveTask) SetCamundaAsyncAfter(asyncAfter bool) {
	rtask.CamundaAsyncAft = asyncAfter
}

// SetCamundaJobPriority ...
func (rtask *ReceiveTask) SetCamundaJobPriority(priority int) {
	rtask.CamundaJobPrio = priority
}

/* Elements */

// SetExtensionElements ...
func (rtask *ReceiveTask) SetExtensionElements() {
	rtask.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (rtask *ReceiveTask) SetIncoming(num int) {
	rtask.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (rtask *ReceiveTask) SetOutgoing(num int) {
	rtask.Outgoing = make([]Outgoing, num)
}
