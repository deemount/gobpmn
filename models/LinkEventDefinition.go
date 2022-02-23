package models

import "fmt"

// LinkEventDefinition ...
type LinkEventDefinition struct {
	ID string `xml:"id,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (led *LinkEventDefinition) SetID(suffix string) {
	led.ID = fmt.Sprintf("LinkEventDefinition_%s", suffix)
}
