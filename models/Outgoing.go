package models

import "fmt"

// Outgoing ...
type Outgoing struct {
	Flow string `xml:",innerxml" json:"flow,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetFlow ...
func (og *Outgoing) SetFlow(suffix string) {
	og.Flow = fmt.Sprintf("Flow_%s", suffix)
}
