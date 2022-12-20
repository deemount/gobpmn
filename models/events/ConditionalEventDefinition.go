package events

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
)

// ConditionalEventDefinitionRepository ...
type ConditionalEventDefinitionRepository interface {
	SetID(typ string, suffix interface{})

	SetCamundaVariableName(variableName string)

	SetCondition()

	GetID() *string

	GetCamundaVariableName() *string

	GetCondition() *attributes.Condition
}

// ConditionalEventDefinition ...
type ConditionalEventDefinition struct {
	ID                  string                 `xml:"id,attr" json:"id"`
	CamundaVariableName string                 `xml:"camunda:variableName" json:"variableName,omitempty"`
	Condition           []attributes.Condition `xml:"bpmn:condition,omitempty" json:"condition,omitempty"`
}

// TConditionalEventDefinition ...
type TConditionalEventDefinition struct {
	ID                  string                 `xml:"id,attr" json:"id"`
	CamundaVariableName string                 `xml:"variableName" json:"variableName,omitempty"`
	Condition           []attributes.Condition `xml:"condition,omitempty" json:"condition,omitempty"`
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetID(typ string, suffix interface{}) {
	switch typ {
	case "ced":
		conditionalEventDefinition.ID = fmt.Sprintf("ConditionalEventDefinition_%v", suffix)
		break
	case "id":
		conditionalEventDefinition.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/** Camunda **/

// SetCamudnaVariableName ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCamundaVariableName(variableName string) {
	conditionalEventDefinition.CamundaVariableName = variableName
}

/*** Make Elements ***/

// SetCondition ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCondition() {
	conditionalEventDefinition.Condition = make([]attributes.Condition, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (conditionalEventDefinition ConditionalEventDefinition) GetID() *string {
	return &conditionalEventDefinition.ID
}

/** Camunda **/

// GetCamundaVariableName ...
func (conditionalEventDefinition ConditionalEventDefinition) GetCamundaVariableName() *string {
	return &conditionalEventDefinition.CamundaVariableName
}

/* Elements */

// GetCondition ...
func (conditionalEventDefinition ConditionalEventDefinition) GetCondition() *attributes.Condition {
	return &conditionalEventDefinition.Condition[0]
}
