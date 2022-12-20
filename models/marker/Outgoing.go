package marker

import "fmt"

// OutgoingRepository ...
type OutgoingRepository interface {
	SetFlow(suffix string)

	GetFlow() *string
}

// Outgoing ...
type Outgoing struct {
	Flow string `xml:",innerxml" json:"flow"`
}

func NewOutgoing() OutgoingRepository {
	return &Outgoing{}
}

/**
 * Default Setters
 */

/* Content */

/** BPMN **/

// SetFlow ...
func (outgoing *Outgoing) SetFlow(suffix string) {
	outgoing.Flow = fmt.Sprintf("Flow_%s", suffix)
}

/**
 * Default Getters
 */

/* Content */

/** BPMN **/

// GetFlow ...
func (outgoing Outgoing) GetFlow() *string {
	return &outgoing.Flow
}
