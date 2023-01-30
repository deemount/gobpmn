package marker

import "fmt"

// NewOutgoing ...
func NewOutgoing() OutgoingRepository {
	return &Outgoing{}
}

/*
 * Default Setters
 */

/* Content */

/** BPMN **/

// SetFlow ...
func (outgoing *Outgoing) SetFlow(suffix interface{}) {
	outgoing.Flow = fmt.Sprintf("Flow_%s", suffix)
}

/*
 * Default Getters
 */

/* Content */

/** BPMN **/

// GetFlow ...
func (outgoing Outgoing) GetFlow() *string {
	return &outgoing.Flow
}
