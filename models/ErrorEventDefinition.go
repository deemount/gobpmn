package models

import "fmt"

// ErrorEventDefinition ...
type ErrorEventDefinition struct {
	ID string `xml:"id,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (eed *ErrorEventDefinition) SetID(suffix string) {
	eed.ID = fmt.Sprintf("ErrorEventDefinition_%s", suffix)
}