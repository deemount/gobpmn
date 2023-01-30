package definitions

import (
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewLinkEventDefinition ...
func NewLinkEventDefinition() LinkEventDefinitionRepository {
	return &LinkEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (led *LinkEventDefinition) SetID(typ string, suffix interface{}) {
	led.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (led LinkEventDefinition) GetID() impl.STR_PTR {
	return &led.ID
}
