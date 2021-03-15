package models

// Collaboration ...
/*
<bpmn:collaboration id="Collaboration_0w0i6yn">
*/
type Collaboration struct {
	ID          string      `xml:"id,attr"`
	Participant Participant `xml:"bpmn:participant"`
}

func (cb *Collaboration) SetID(suffix string) {
	cb.ID = "Collaboration_" + suffix
}
