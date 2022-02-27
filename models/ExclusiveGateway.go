package models

import "fmt"

// ExclusiveGatewayRepository ...
type ExclusiveGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ExclusiveGateway ...
type ExclusiveGateway struct {
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
func (exclusivegate *ExclusiveGateway) SetID(suffix string) {
	exclusivegate.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (exclusivegate *ExclusiveGateway) SetName(name string) {
	exclusivegate.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (exclusivegate *ExclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	exclusivegate.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (exclusivegate *ExclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	exclusivegate.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (exclusivegate *ExclusiveGateway) SetCamundaJobPriority(priority int) {
	exclusivegate.CamundaJobPriority = priority
}

/* Elements */

// SetExtensionElements ...
func (exclusivegate *ExclusiveGateway) SetExtensionElements() {
	exclusivegate.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (exclusivegate *ExclusiveGateway) SetIncoming(num int) {
	exclusivegate.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (exclusivegate *ExclusiveGateway) SetOutgoing(num int) {
	exclusivegate.Outgoing = make([]Outgoing, num)
}
