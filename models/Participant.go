package models

// Participant ...
type Participant struct {
	ID                      string                    `xml:"id,attr"`
	Name                    string                    `xml:"name,attr,omitempty"`
	ProcessRef              string                    `xml:"processRef,attr"`
	Documentation           []Documentation           `xml:"bpmn:documentation,omitempty"`
	ParticipantMultiplicity []ParticipantMultiplicity `xml:"bpmn:participantMultiplicity,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (pp *Participant) SetID(suffix string) {
	pp.ID = "Participant_" + suffix
}

// SetName ...
func (pp *Participant) SetName(name string) {
	pp.Name = name
}

// SetProcessRef ...
func (pp *Participant) SetProcessRef(suffix string) {
	pp.ProcessRef = "Process_" + suffix
}

/* Elements */

/** BPMN **/

// SetParticipantMultiplicity ...
func (participant *Participant) SetParticipantMultiplicity() {
	participant.ParticipantMultiplicity = make([]ParticipantMultiplicity, 1)
}
