package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewSignal ...
func NewSignal() SignalRepository {
	return &Signal{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (signal *Signal) SetID(typ string, suffix interface{}) {
	switch typ {
	case "signal":
		signal.ID = fmt.Sprintf("Signal_%v", suffix)
		break
	case "id":
		signal.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (signal *Signal) SetName(suffix string) {
	signal.Name = fmt.Sprintf("Signal_%s", suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (signal Signal) GetID() impl.STR_PTR {
	return &signal.ID
}

// GetName ...
func (signal Signal) GetName() impl.STR_PTR {
	return &signal.Name
}
