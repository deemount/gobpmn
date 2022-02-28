package models

import "fmt"

// Collaboration ...
type Collaboration struct {
	ID                string              `xml:"id,attr" json:"id,omitempty"`
	Documentation     []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Participant       []Participant       `xml:"bpmn:participant" json:"participant,omitempty"`
	MessageFlow       []MessageFlow       `xml:"bpmn:messageFlow,omitempty" json:"messageFlow,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (cb *Collaboration) SetID(suffix string) {
	cb.ID = fmt.Sprintf("Collaboration_%s", suffix)
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (cb *Collaboration) SetDocumentation() {
	cb.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (cb *Collaboration) SetExtensionElements() {
	cb.ExtensionElements = make([]ExtensionElements, 1)
}

// SetParticipant ...
func (cb *Collaboration) SetParticipant(num int) {
	cb.Participant = make([]Participant, num)
}

// SetMessageFlow ...
func (cb *Collaboration) SetMessageFlow(num int) {
	cb.MessageFlow = make([]MessageFlow, num)
}
