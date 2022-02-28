package models

import "fmt"

// ServiceTaskRepository ...
type ServiceTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ServiceTask ...
type ServiceTask struct {
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
func (serviceTask *ServiceTask) SetID(suffix string) {
	serviceTask.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (serviceTask *ServiceTask) SetName(name string) {
	serviceTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (serviceTask *ServiceTask) SetCamundaAsyncBefore(asyncBefore bool) {
	serviceTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (serviceTask *ServiceTask) SetCamundaAsyncAfter(asyncAfter bool) {
	serviceTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (serviceTask *ServiceTask) SetCamundaJobPriority(priority int) {
	serviceTask.CamundaJobPriority = priority
}

/* Elements */

// SetDocumentation ...
func (serviceTask *ServiceTask) SetDocumentation() {
	serviceTask.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (serviceTask *ServiceTask) SetExtensionElements() {
	serviceTask.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (serviceTask *ServiceTask) SetIncoming(num int) {
	serviceTask.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (serviceTask *ServiceTask) SetOutgoing(num int) {
	serviceTask.Outgoing = make([]Outgoing, num)
}
