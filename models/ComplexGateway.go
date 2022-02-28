package models

import "fmt"

// ComplexGatewayRepository ...
type ComplexGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ComplexGateway ...
type ComplexGateway struct {
	ID                 string              `xml:"id,attr" json:"id"`
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
func (complexgate *ComplexGateway) SetID(suffix string) {
	complexgate.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (complexgate *ComplexGateway) SetName(name string) {
	complexgate.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexgate *ComplexGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	complexgate.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (complexgate *ComplexGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	complexgate.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (complexgate *ComplexGateway) SetCamundaJobPriority(priority int) {
	complexgate.CamundaJobPriority = priority
}

/* Elements */

// SetExtensionElements ...
func (complexgate *ComplexGateway) SetExtensionElements() {
	complexgate.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (complexgate *ComplexGateway) SetIncoming(num int) {
	complexgate.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (complexgate *ComplexGateway) SetOutgoing(num int) {
	complexgate.Outgoing = make([]Outgoing, num)
}
