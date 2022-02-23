package models

import "fmt"

// SetIntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	ID                        string                      `xml:"id,attr,omitempty"`
	Name                      string                      `xml:"name,attr,omitempty"`
	ExtensionElements         []ExtensionElements         `xml:"bpmn:extensionElements,omitempty"`
	Incoming                  []Incoming                  `xml:"bpmn:incoming,omitempty"`
	Outgoing                  []Outgoing                  `xml:"bpmn:outgoing,omitempty"`
	EscalationEventDefinition []EscalationEventDefinition `xml:"bpmn:escalationEventDefinition,omitempty"`
	MessageEventDefinition    []MessageEventDefinition    `xml:"bpmn:messageEventDefinition,omitempty"`
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

// SetEscalationEventDefinition ...
func (ite *IntermediateThrowEvent) SetEscalationEventDefinition() {
	ite.EscalationEventDefinition = make([]EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (ite *IntermediateThrowEvent) SetMessageEventDefinition() {
	ite.MessageEventDefinition = make([]MessageEventDefinition, 1)
}
