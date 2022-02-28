package models

import "fmt"

// TaskRepository ...
type TaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// Task ...
type Task struct {
	ID                string              `xml:"id,attr" json:"id"`
	Name              string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (task *Task) SetID(suffix string) {
	task.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (task *Task) SetName(name string) {
	task.Name = name
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (task *Task) SetExtensionElements() {
	task.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (task *Task) SetIncoming(num int) {
	task.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (task *Task) SetOutgoing(num int) {
	task.Outgoing = make([]Outgoing, num)
}
