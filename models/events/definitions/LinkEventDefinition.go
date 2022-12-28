package definitions

import (
	"fmt"
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
	switch typ {
	case "led":
		led.ID = fmt.Sprintf("LinkEventDefinition_%v", suffix)
		break
	case "id":
		led.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (led LinkEventDefinition) GetID() STR_PTR {
	return &led.ID
}
