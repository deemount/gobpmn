package models

import "fmt"

// EventBasedGatewayRepository ...
type EventBasedGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// EventBasedGateway ...
type EventBasedGateway struct {
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
func (eventBasedGateway *EventBasedGateway) SetID(suffix string) {
	eventBasedGateway.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (eventBasedGateway *EventBasedGateway) SetName(name string) {
	eventBasedGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (eventBasedGateway *EventBasedGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	eventBasedGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (eventBasedGateway *EventBasedGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	eventBasedGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (eventBasedGateway *EventBasedGateway) SetCamundaJobPriority(priority int) {
	eventBasedGateway.CamundaJobPriority = priority
}

/* Elements */

// SetExtensionElements ...
func (eventBasedGateway *EventBasedGateway) SetExtensionElements() {
	eventBasedGateway.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (eventBasedGateway *EventBasedGateway) SetIncoming(num int) {
	eventBasedGateway.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (eventBasedGateway *EventBasedGateway) SetOutgoing(num int) {
	eventBasedGateway.Outgoing = make([]Outgoing, num)
}
