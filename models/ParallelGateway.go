package models

import "fmt"

// ParallelGatewayRepository ...
type ParallelGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ParallelGateway ...
type ParallelGateway struct {
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
func (parallelGateway *ParallelGateway) SetID(suffix string) {
	parallelGateway.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (parallelGateway *ParallelGateway) SetName(name string) {
	parallelGateway.Name = name
}

/* Elements */

// SetExtensionElements ...
func (parallelGateway *ParallelGateway) SetExtensionElements() {
	parallelGateway.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (parallelGateway *ParallelGateway) SetIncoming(num int) {
	parallelGateway.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (parallelGateway *ParallelGateway) SetOutgoing(num int) {
	parallelGateway.Outgoing = make([]Outgoing, num)
}
