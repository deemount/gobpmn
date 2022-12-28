package definitions

import (
	"fmt"
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
	switch typ {
	case "eed":
		eed.ID = fmt.Sprintf("ErrorEventDefinition_%v", suffix)
		break
	case "id":
		eed.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (eed ErrorEventDefinition) GetID() STR_PTR {
	return &eed.ID
}
