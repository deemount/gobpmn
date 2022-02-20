package models

import "fmt"

// Signal ...
type Signal struct {
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (signal *Signal) SetID(suffix string) {
	signal.ID = fmt.Sprintf("Signal_%s", suffix)
}

// SetName ...
func (signal *Signal) SetName(suffix string) {
	signal.Name = fmt.Sprintf("Signal_%s", suffix)
}
