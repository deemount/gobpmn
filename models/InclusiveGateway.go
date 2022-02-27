package models

import "fmt"

// InclusiveGatewayRepository ...
type InclusiveGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// InclusiveGateway ...
type InclusiveGateway struct {
	ID                 string              `xml:"id,attr"`
	Name               string              `xml:"name,attr,omitempty"`
	CamundaAsyncBefore bool                `xml:"camunda:asyncBefore,attr,omitempty"`
	CamundaAsyncAfter  bool                `xml:"camunda:asyncAfter,attr,omitempty"`
	CamundaJobPriority int                 `xml:"camunda:jobPriority,attr,omitempty"`
	Documentation      []Documentation     `xml:"bpmn:documentation,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
	Incoming           []Incoming          `xml:"bpmn:incoming,omitempty"`
	Outgoing           []Outgoing          `xml:"bpmn:outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusivegate *InclusiveGateway) SetID(suffix string) {
	inclusivegate.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (inclusivegate *InclusiveGateway) SetName(name string) {
	inclusivegate.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (inclusivegate *InclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	inclusivegate.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (inclusivegate *InclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	inclusivegate.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (inclusivegate *InclusiveGateway) SetCamundaJobPriority(priority int) {
	inclusivegate.CamundaJobPriority = priority
}

/* Elements */

// SetExtensionElements ...
func (inclusivegate *InclusiveGateway) SetExtensionElements() {
	inclusivegate.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (inclusivegate *InclusiveGateway) SetIncoming(num int) {
	inclusivegate.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (inclusivegate *InclusiveGateway) SetOutgoing(num int) {
	inclusivegate.Outgoing = make([]Outgoing, num)
}
