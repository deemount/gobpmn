package models

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	Incoming []Incoming `xml:"bpmn:incoming,omitempty"`
	Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty"`
}
