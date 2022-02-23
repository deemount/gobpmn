package models

import "fmt"

// CallActivity ...
type CallActivity struct {
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (ca *CallActivity) SetID(suffix string) {
	ca.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (ca *CallActivity) SetName(name string) {
	ca.Name = name
}
