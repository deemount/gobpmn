package models

import "fmt"

// ParticipantRepository ...
type ParticipantRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// Participant ...
type Participant struct {
	ID                      string                    `xml:"id,attr" json:"id"`
	Name                    string                    `xml:"name,attr,omitempty" json:"name,omitempty"`
	ProcessRef              string                    `xml:"processRef,attr" json:"processRef,omitempty"`
	Documentation           []Documentation           `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ParticipantMultiplicity []ParticipantMultiplicity `xml:"bpmn:participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty"`
}

// TParticipant ...
type TParticipant struct {
	ID                      string                    `xml:"id,attr" json:"id"`
	Name                    string                    `xml:"name,attr,omitempty" json:"name,omitempty"`
	ProcessRef              string                    `xml:"processRef,attr" json:"processRef,omitempty"`
	Documentation           []Documentation           `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ParticipantMultiplicity []ParticipantMultiplicity `xml:"participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (participant *Participant) SetID(suffix string) {
	participant.ID = fmt.Sprintf("Participant_%s", suffix)
}

// SetName ...
func (participant *Participant) SetName(name string) {
	participant.Name = name
}

// SetProcessRef ...
func (participant *Participant) SetProcessRef(suffix string) {
	participant.ProcessRef = fmt.Sprintf("Process_%s", suffix)
}

/* Elements */

/** BPMN **/

// SetParticipantMultiplicity ...
func (participant *Participant) SetParticipantMultiplicity() {
	participant.ParticipantMultiplicity = make([]ParticipantMultiplicity, 1)
}
