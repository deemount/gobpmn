package models

import "fmt"

// BusinessRuleTaskRepository ...
type BusinessRuleTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// BusinessRuleTask ...
type BusinessRuleTask struct {
	ID                 string              `xml:"id,attr"`
	Name               string              `xml:"name,attr,omitempty"`
	CamundaAsyncBefore bool                `xml:"camunda:asyncBefore,attr,omitempty"`
	CamundaAsyncAfter  bool                `xml:"camunda:asyncAfter,attr,omitempty"`
	CamundaJobPriority int                 `xml:"camunda:jobPriority,attr,omitempty"`
	CamundaClass       string              `xml:"camunda:class,attr,omitempty"`
	Documentation      []Documentation     `xml:"bpmn:documentation,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
	Incoming           []Incoming          `xml:"bpmn:incoming,omitempty"`
	Outgoing           []Outgoing          `xml:"bpmn:outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (brt *BusinessRuleTask) SetID(suffix string) {
	brt.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (brt *BusinessRuleTask) SetName(name string) {
	brt.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (brt *BusinessRuleTask) SetCamundaAsyncBefore(asyncBefore bool) {
	brt.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (brt *BusinessRuleTask) SetCamundaAsyncAfter(asyncAfter bool) {
	brt.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (brt *BusinessRuleTask) SetCamundaJobPriority(priority int) {
	brt.CamundaJobPriority = priority
}

// SetCamundaClass ...
func (brt *BusinessRuleTask) SetCamundaClass(class string) {
	brt.CamundaClass = class
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (brt *BusinessRuleTask) SetDocumentation() {
	brt.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (brt *BusinessRuleTask) SetExtensionElements() {
	brt.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (brt *BusinessRuleTask) SetIncoming(num int) {
	brt.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (brt *BusinessRuleTask) SetOutgoing(num int) {
	brt.Outgoing = make([]Outgoing, num)
}
