package marker

import "fmt"

// NewIncoming ...
func NewIncoming() IncomingRepository {
	return &Incoming{}
}

/*
 * Default Setters
 */

/* Content */

/** BPMN **/

// SetFlow ...
func (incoming *Incoming) SetFlow(suffix interface{}) {
	incoming.Flow = fmt.Sprintf("Flow_%s", suffix)
}

/*
 * Default Getters
 */

/* Content */

/** BPMN **/

// GetFlow ...
func (incoming Incoming) GetFlow() *string {
	return &incoming.Flow
}
