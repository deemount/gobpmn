package models

import "fmt"

// InclusiveGatewayRepository ...
type InclusiveGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// InclusiveGateway ...
type InclusiveGateway struct {
	ID                 string              `xml:"id,attr" json:"id"`
	Name               string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore"`
	CamundaAsyncAfter  bool                `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter"`
	CamundaJobPriority int                 `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusiveGateway *InclusiveGateway) SetID(suffix string) {
	inclusiveGateway.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (inclusiveGateway *InclusiveGateway) SetName(name string) {
	inclusiveGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (inclusiveGateway *InclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	inclusiveGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (inclusiveGateway *InclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	inclusiveGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (inclusiveGateway *InclusiveGateway) SetCamundaJobPriority(priority int) {
	inclusiveGateway.CamundaJobPriority = priority
}

/* Elements */

// SetExtensionElements ...
func (inclusiveGateway *InclusiveGateway) SetExtensionElements() {
	inclusiveGateway.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (inclusiveGateway *InclusiveGateway) SetIncoming(num int) {
	inclusiveGateway.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (inclusiveGateway *InclusiveGateway) SetOutgoing(num int) {
	inclusiveGateway.Outgoing = make([]Outgoing, num)
}
