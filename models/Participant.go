package models

import "fmt"

// ParticipantRepository ...
type ParticipantRepository interface {
	SetID(typ string, suffix string)
	SetName(name string)
}

// Participant ...
type Participant struct {
	ID                      string                    `xml:"id,attr" json:"id" csv:"ID"`
	Name                    string                    `xml:"name,attr,omitempty" json:"name,omitempty" csv:"NAME"`
	ProcessRef              string                    `xml:"processRef,attr,omitempty" json:"processRef,omitempty" csv:"PROCESS_REF"`
	Documentation           []Documentation           `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"DOCUMENTATION"`
	ParticipantMultiplicity []ParticipantMultiplicity `xml:"bpmn:participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty" csv:"PARTICIPANT_MULTIPLICITY"`
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
func (participant *Participant) SetID(typ string, suffix string) {
	switch typ {
	case "participant":
		participant.ID = fmt.Sprintf("Participant_%s", suffix)
		break
	case "id":
		participant.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (participant *Participant) SetName(name string) {
	participant.Name = name
}

// SetProcessRef ...
func (participant *Participant) SetProcessRef(typ string, suffix string) {
	switch typ {
	case "process":
		participant.ProcessRef = fmt.Sprintf("Process_%s", suffix)
		break
	case "id":
		participant.ProcessRef = fmt.Sprintf("%s", suffix)
		break
	}
}

/* Elements */

/** BPMN **/

// SetParticipantMultiplicity ...
func (participant *Participant) SetParticipantMultiplicity() {
	participant.ParticipantMultiplicity = make([]ParticipantMultiplicity, 1)
}
