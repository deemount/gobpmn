package models

import "fmt"

// TerminateEventDefinition ...
type TerminateEventDefinition struct {
	ID string `xml:"id,attr,omitempty" json:"" id`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (ted *TerminateEventDefinition) SetID(suffix string) {
	ted.ID = fmt.Sprintf("TerminateEventDefinition_%s", suffix)
}
