package events

import "fmt"

// CancelEventDefinitionRepository ...
type CancelEventDefinitionRepository interface {
	SetID(typ string, suffix interface{})
	GetID() *string
}

// CancelEventDefinition ...
type CancelEventDefinition struct {
	ID string `xml:"id,attr,omitempty" json:"id"`
}

func NewCancelEventDefinition() CancelEventDefinitionRepository {
	return &CancelEventDefinition{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ced *CancelEventDefinition) SetID(typ string, suffix interface{}) {
	switch typ {
	case "ced":
		ced.ID = fmt.Sprintf("CancelEventDefinition_%v", suffix)
		break
	case "id":
		ced.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ced CancelEventDefinition) GetID() *string {
	return &ced.ID
}
