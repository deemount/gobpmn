package models

import "fmt"

// IntermediateThrowEventRepository ...
type IntermediateThrowEventRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	ID                        string                      `xml:"id,attr,omitempty" json:"id"`
	Name                      string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements         []ExtensionElements         `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []Incoming                  `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                  []Outgoing                  `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	CompensateEventDefinition []CompensateEventDefinition `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	LinkEventDefinition       []LinkEventDefinition       `xml:"bpmn:linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	EscalationEventDefinition []EscalationEventDefinition `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []MessageEventDefinition    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetID(suffix string) {
	intermediateThrowEvent.ID = fmt.Sprintf("Event_%s", suffix)
}

// SetName ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetName(name string) {
	intermediateThrowEvent.Name = name
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetExtensionElements() {
	intermediateThrowEvent.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetIncoming(num int) {
	intermediateThrowEvent.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetOutgoing(num int) {
	intermediateThrowEvent.Outgoing = make([]Outgoing, num)
}

// SetLinkEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetLinkEventDefinition() {
	intermediateThrowEvent.LinkEventDefinition = make([]LinkEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetEscalationEventDefinition() {
	intermediateThrowEvent.EscalationEventDefinition = make([]EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetMessageEventDefinition() {
	intermediateThrowEvent.MessageEventDefinition = make([]MessageEventDefinition, 1)
}
