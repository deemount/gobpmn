package models

import "fmt"

// ServiceTaskRepository ...
type ServiceTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ServiceTask ...
type ServiceTask struct {
	ID                string              `xml:"id,attr,omitempty" json:"id"`
	Name              string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
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

/* Elements */

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
