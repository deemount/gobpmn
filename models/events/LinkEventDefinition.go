package events

import "fmt"

// LinkEventDefinitionRepository ...
type LinkEventDefinitionRepository interface {
	SetID(typ string, suffix interface{})

	GetID() *string
}

// LinkEventDefinition ...
type LinkEventDefinition struct {
	ID string `xml:"id,attr,omitempty" json:"id"`
}

func NewLinkEventDefinition() LinkEventDefinitionRepository {
	return &LinkEventDefinition{}
}

/**
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

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (led LinkEventDefinition) GetID() *string {
	return &led.ID
}
