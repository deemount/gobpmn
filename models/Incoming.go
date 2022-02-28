package models

// Incoming ...
type Incoming struct {
	Flow string `xml:",innerxml" json:"flow"`
}

/* Content */

/** BPMN **/

// SetFlow ...
func (incoming *Incoming) SetFlow(suffix string) {
	incoming.Flow = "Flow_" + suffix
}
