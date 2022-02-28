package models

import "fmt"

// BusinessRuleTaskRepository ...
type BusinessRuleTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// BusinessRuleTask ...
type BusinessRuleTask struct {
	ID                 string              `xml:"id,attr" json:"id"`
	Name               string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                 `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	CamundaClass       string              `xml:"camunda:class,attr,omitempty" json:"class,omitempty"`
	Documentation      []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (businessRuleTask *BusinessRuleTask) SetID(suffix string) {
	businessRuleTask.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (businessRuleTask *BusinessRuleTask) SetName(name string) {
	businessRuleTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (businessRuleTask *BusinessRuleTask) SetCamundaAsyncBefore(asyncBefore bool) {
	businessRuleTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (businessRuleTask *BusinessRuleTask) SetCamundaAsyncAfter(asyncAfter bool) {
	businessRuleTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (businessRuleTask *BusinessRuleTask) SetCamundaJobPriority(priority int) {
	businessRuleTask.CamundaJobPriority = priority
}

// SetCamundaClass ...
func (businessRuleTask *BusinessRuleTask) SetCamundaClass(class string) {
	businessRuleTask.CamundaClass = class
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (businessRuleTask *BusinessRuleTask) SetDocumentation() {
	businessRuleTask.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (businessRuleTask *BusinessRuleTask) SetExtensionElements() {
	businessRuleTask.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (businessRuleTask *BusinessRuleTask) SetIncoming(num int) {
	businessRuleTask.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (businessRuleTask *BusinessRuleTask) SetOutgoing(num int) {
	businessRuleTask.Outgoing = make([]Outgoing, num)
}
