package models

import "fmt"

// Collaboration ...
type Collaboration struct {
	ID          string        `xml:"id,attr"`
	Participant []Participant `xml:"bpmn:participant"`
	MessageFlow []MessageFlow `xml:"bpmn:messageFlow,omitempty"`
}

// SetID ...
func (cb *Collaboration) SetID(suffix string) {
	cb.ID = fmt.Sprintf("Collaboration_%s", suffix)
}

/* Make Elements */

// SetParticipant ...
func (cb *Collaboration) SetParticipant(num int) []Participant {
	cb.Participant = make([]Participant, num, 1)
	return cb.Participant
}
