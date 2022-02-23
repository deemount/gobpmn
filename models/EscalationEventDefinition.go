package models

import "fmt"

// EscalationEventDefinition ...
type EscalationEventDefinition struct {
	ID string `xml:"id,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (esced *EscalationEventDefinition) SetID(suffix string) {
	esced.ID = fmt.Sprintf("EscalationEventDefinition_%s", suffix)
}
