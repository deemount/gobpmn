package definitions

import (
	"github.com/deemount/gobpmn/models/compulsion"
)

// NewErrorEventDefinition ...
func NewErrorEventDefinition() ErrorEventDefinitionRepository {
	return &ErrorEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (eed *ErrorEventDefinition) SetID(typ string, suffix interface{}) {
	eed.ID = SetID(typ, suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (eed ErrorEventDefinition) GetID() compulsion.STR_PTR {
	return &eed.ID
}
