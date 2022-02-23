package models

import "fmt"

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	ID                         string                       `xml:"id,attr,omitempty"`
	Name                       string                       `xml:"name,attr,omitempty"`
	ExtensionElements          []ExtensionElements          `xml:"bpmn:extensionElements,omitempty"`
	Incoming                   []Incoming                   `xml:"bpmn:incoming,omitempty"`
	Outgoing                   []Outgoing                   `xml:"bpmn:outgoing,omitempty"`
	ConditionalEventDefinition []ConditionalEventDefinition `xml:"bpmn:conditionalEventDefinition,omitempty"`
	TimerEventDefinition       []TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty"`
	MessageEventDefinition     []MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (ice *IntermediateCatchEvent) SetID(suffix string) {
	ice.ID = fmt.Sprintf("Event_%s", suffix)
}

// SetName ...
func (ice *IntermediateCatchEvent) SetName(name string) {
	ice.Name = name
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (ice *IntermediateCatchEvent) SetExtensionElements() {
	ice.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (ice *IntermediateCatchEvent) SetIncoming(num int) {
	ice.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (ice *IntermediateCatchEvent) SetOutgoing(num int) {
	ice.Outgoing = make([]Outgoing, num)
}

// SetConditionalEventDefinition ...
func (ice *IntermediateCatchEvent) SetConditionalEventDefinition() {
	ice.ConditionalEventDefinition = make([]ConditionalEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (ice *IntermediateCatchEvent) SetTimerEventDefinition() {
	ice.TimerEventDefinition = make([]TimerEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (ice *IntermediateCatchEvent) SetMessageEventDefinition() {
	ice.MessageEventDefinition = make([]MessageEventDefinition, 1)
}
