package events

import "fmt"

// EscalationEventDefinitionRepository ...
type EscalationEventDefinitionRepository interface {
	SetID(typ string, suffix interface{})

	GetID() *string
}

// EscalationEventDefinition ...
type EscalationEventDefinition struct {
	ID string `xml:"id,attr,omitempty" json:"id"`
}

func NewEscalationEventDefinition() EscalationEventDefinitionRepository {
	return &EscalationEventDefinition{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (esced *EscalationEventDefinition) SetID(typ string, suffix interface{}) {
	switch typ {
	case "eed":
		esced.ID = fmt.Sprintf("EscalationEventDefinition_%v", suffix)
		break
	case "id":
		esced.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (esced EscalationEventDefinition) GetID() *string {
	return &esced.ID
}
