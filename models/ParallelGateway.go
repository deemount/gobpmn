package models

import "fmt"

// ParallelGatewayRepository ...
type ParallelGatewayRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ParallelGateway ...
type ParallelGateway struct {
	ID                string              `xml:"id,attr"`
	Name              string              `xml:"name,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
	Incoming          []Incoming          `xml:"bpmn:incoming,omitempty"`
	Outgoing          []Outgoing          `xml:"bpmn:outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (parallelgate *ParallelGateway) SetID(suffix string) {
	parallelgate.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (parallelgate *ParallelGateway) SetName(name string) {
	parallelgate.Name = name
}

/* Elements */

// SetExtensionElements ...
func (parallelgate *ParallelGateway) SetExtensionElements() {
	parallelgate.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (parallelgate *ParallelGateway) SetIncoming(num int) {
	parallelgate.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (parallelgate *ParallelGateway) SetOutgoing(num int) {
	parallelgate.Outgoing = make([]Outgoing, num)
}
