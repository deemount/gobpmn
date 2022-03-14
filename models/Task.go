package models

import "fmt"

// TaskRepository ...
type TaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// Task ...
type Task struct {
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

// TTask ...
type TTask struct {
	ID                 string              `xml:"id,attr,omitempty" json:"id"`
	Name               string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                 `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []Documentation     `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []Incoming          `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []Outgoing          `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
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

/** Camunda **/

// SetCamundaAsyncBefore ...
func (task *Task) SetCamundaAsyncBefore(asyncBefore bool) {
	task.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (task *Task) SetCamundaAsyncAfter(asyncAfter bool) {
	task.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (task *Task) SetCamundaJobPriority(priority int) {
	task.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (task *Task) SetDocumentation() {
	task.Documentation = make([]Documentation, 1)
}

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
