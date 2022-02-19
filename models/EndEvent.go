package models

// EndEvent ...
type EndEvent struct {
	ID       string     `xml:"id,attr"`
	Incoming []Incoming `xml:"bpmn:incoming,omitempty"`
}

// SetID ...
func (ee *EndEvent) SetID(suffix string) {
	ee.ID = "Event_" + suffix
}
