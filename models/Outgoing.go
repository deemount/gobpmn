package models

import "fmt"

// Outgoing ...
type Outgoing struct {
	Flow string `xml:",innerxml" json:"flow"`
}

/* Content */

/** BPMN **/

// SetFlow ...
func (outgoing *Outgoing) SetFlow(suffix string) {
	outgoing.Flow = fmt.Sprintf("Flow_%s", suffix)
}
