package models

import "fmt"

// ExclusiveGatewayRepository ...
type ExclusiveGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ExclusiveGateway ...
type ExclusiveGateway struct {
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
func (exclusiveGateway *ExclusiveGateway) SetID(suffix string) {
	exclusiveGateway.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (exclusiveGateway *ExclusiveGateway) SetName(name string) {
	exclusiveGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	exclusiveGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	exclusiveGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaJobPriority(priority int) {
	exclusiveGateway.CamundaJobPriority = priority
}

/* Elements */

// SetExtensionElements ...
func (exclusiveGateway *ExclusiveGateway) SetExtensionElements() {
	exclusiveGateway.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (exclusiveGateway *ExclusiveGateway) SetIncoming(num int) {
	exclusiveGateway.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (exclusiveGateway *ExclusiveGateway) SetOutgoing(num int) {
	exclusiveGateway.Outgoing = make([]Outgoing, num)
}
