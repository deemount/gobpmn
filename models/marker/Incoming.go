package marker

import "fmt"

// IncomingRepository ...
type IncomingRepository interface {
	SetFlow(suffix string)
	GetFlow() *string
}

// Incoming ...
type Incoming struct {
	Flow string `xml:",innerxml" json:"flow"`
}

func NewIncoming() IncomingRepository {
	return &Incoming{}
}

/*
 * Default Setters
 */

/* Content */

/** BPMN **/

// SetFlow ...
func (incoming *Incoming) SetFlow(suffix string) {
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
