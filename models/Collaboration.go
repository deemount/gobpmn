package models

import "fmt"

// Collaboration ...
type Collaboration struct {
	ID                string              `xml:"id,attr" json:"id"`
	Documentation     []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Participant       []Participant       `xml:"bpmn:participant" json:"participant,omitempty"`
	MessageFlow       []MessageFlow       `xml:"bpmn:messageFlow,omitempty" json:"messageFlow,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (collaboration *Collaboration) SetID(suffix string) {
	collaboration.ID = fmt.Sprintf("Collaboration_%s", suffix)
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (collaboration *Collaboration) SetDocumentation() {
	collaboration.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (collaboration *Collaboration) SetExtensionElements() {
	collaboration.ExtensionElements = make([]ExtensionElements, 1)
}

// SetParticipant ...
func (collaboration *Collaboration) SetParticipant(num int) {
	collaboration.Participant = make([]Participant, num)
}

// SetMessageFlow ...
func (collaboration *Collaboration) SetMessageFlow(num int) {
	collaboration.MessageFlow = make([]MessageFlow, num)
}
