package models

import "fmt"

// SignalRepository ...
type SignalRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// Signal ...
type Signal struct {
	ID   string `xml:"id,attr,omitempty" json:"id"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
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
