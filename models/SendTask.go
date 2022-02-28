package models

import "fmt"

// SendTaskRepository ...
type SendTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// SendTask ...
type SendTask struct {
	ID                string              `xml:"id,attr,omitempty"`
	Name              string              `xml:"name,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
	Incoming          []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
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

/* Elements */

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
