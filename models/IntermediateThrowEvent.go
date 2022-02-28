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
func (ite *IntermediateThrowEvent) SetID(suffix string) {
	ite.ID = fmt.Sprintf("Event_%s", suffix)
}

// SetName ...
func (ite *IntermediateThrowEvent) SetName(name string) {
	ite.Name = name
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (ite *IntermediateThrowEvent) SetExtensionElements() {
	ite.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (ite *IntermediateThrowEvent) SetIncoming(num int) {
	ite.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (ite *IntermediateThrowEvent) SetOutgoing(num int) {
	ite.Outgoing = make([]Outgoing, num)
}

// SetLinkEventDefinition ...
func (ite *IntermediateThrowEvent) SetLinkEventDefinition() {
	ite.LinkEventDefinition = make([]LinkEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (ite *IntermediateThrowEvent) SetEscalationEventDefinition() {
	ite.EscalationEventDefinition = make([]EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (ite *IntermediateThrowEvent) SetMessageEventDefinition() {
	ite.MessageEventDefinition = make([]MessageEventDefinition, 1)
}
