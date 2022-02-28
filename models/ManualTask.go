package models

import "fmt"

// ManualTaskRepository ...
type ManualTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ManualTask ...
type ManualTask struct {
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
func (manualTask *ManualTask) SetID(suffix string) {
	manualTask.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (manualTask *ManualTask) SetName(name string) {
	manualTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (manualTask *ManualTask) SetCamundaAsyncBefore(asyncBefore bool) {
	manualTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (manualTask *ManualTask) SetCamundaAsyncAfter(asyncAfter bool) {
	manualTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (manualTask *ManualTask) SetCamundaJobPriority(priority int) {
	manualTask.CamundaJobPriority = priority
}

/* Elements */

// SetDocumentation ...
func (manualTask *ManualTask) SetDocumentation() {
	manualTask.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (manualTask *ManualTask) SetExtensionElements() {
	manualTask.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (manualTask *ManualTask) SetIncoming(num int) {
	manualTask.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (manualTask *ManualTask) SetOutgoing(num int) {
	manualTask.Outgoing = make([]Outgoing, num)
}
