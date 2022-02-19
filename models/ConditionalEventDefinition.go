package models

// ConditionalEventDefinition ...
type ConditionalEventDefinition struct {
	ID   string    `xml:"id,attr"`
	Cond Condition `xml:"bpmn:condition,omitempty"`
}

// SetID ...
func (ced *ConditionalEventDefinition) SetID(suffix string) {
	ced.ID = "ConditionalEventDefinition_" + suffix
}
