package models

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	ID       string   `xml:"id,attr"`
	Incoming Incoming `xml:"bpmn:incoming"`
}

// SetID ...
func (ite *IntermediateThrowEvent) SetID(suffix string) {
	ite.ID = "Event_" + suffix
}
