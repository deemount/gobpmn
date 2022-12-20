package events

import "fmt"

// TerminateEventDefinitionRepository ...
type TerminateEventDefinitionRepository interface {
	SetID(typ string, suffix interface{})
	GetID() *string
}

// TerminateEventDefinition ...
type TerminateEventDefinition struct {
	ID string `xml:"id,attr,omitempty" json:"id"`
}

func NewTerminateEventDefinition() TerminateEventDefinitionRepository {
	return &TerminateEventDefinition{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ted *TerminateEventDefinition) SetID(typ string, suffix interface{}) {
	switch typ {
	case "ted":
		ted.ID = fmt.Sprintf("TerminateEventDefinition_%v", suffix)
		break
	case "id":
		ted.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ted TerminateEventDefinition) GetID() *string {
	return &ted.ID
}
