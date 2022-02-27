package models

import "fmt"

// EventBasedGatewayRepository ...
type EventBasedGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// EventBasedGateway ...
type EventBasedGateway struct {
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
func (eventbasedgate *EventBasedGateway) SetID(suffix string) {
	eventbasedgate.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (eventbasedgate *EventBasedGateway) SetName(name string) {
	eventbasedgate.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (eventbasedgate *EventBasedGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	eventbasedgate.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (eventbasedgate *EventBasedGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	eventbasedgate.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (eventbasedgate *EventBasedGateway) SetCamundaJobPriority(priority int) {
	eventbasedgate.CamundaJobPriority = priority
}

/* Elements */

// SetExtensionElements ...
func (eventbasedgate *EventBasedGateway) SetExtensionElements() {
	eventbasedgate.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (eventbasedgate *EventBasedGateway) SetIncoming(num int) {
	eventbasedgate.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (eventbasedgate *EventBasedGateway) SetOutgoing(num int) {
	eventbasedgate.Outgoing = make([]Outgoing, num)
}
