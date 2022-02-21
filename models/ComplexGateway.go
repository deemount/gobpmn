package models

import "fmt"

// ComplexGateway ...
type ComplexGateway struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (complexgate *ComplexGateway) SetID(suffix string) {
	complexgate.ID = fmt.Sprintf("Gateway_%s", suffix)
}
