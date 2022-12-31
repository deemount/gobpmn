package definitions

import (
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/conditional"
)

// NewConditionalEventDefinition ...
func NewConditionalEventDefinition() ConditionalEventDefinitionRepository {
	return &ConditionalEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetID(typ string, suffix interface{}) {
	conditionalEventDefinition.ID = SetID(typ, suffix)
}

/** Camunda **/

// SetCamudnaVariableName ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCamundaVariableName(variableName string) {
	conditionalEventDefinition.CamundaVariableName = variableName
}

/*** Make Elements ***/

// SetCondition ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCondition() {
	conditionalEventDefinition.Condition = make([]conditional.Condition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (conditionalEventDefinition ConditionalEventDefinition) GetID() compulsion.STR_PTR {
	return &conditionalEventDefinition.ID
}

/** Camunda **/

// GetCamundaVariableName ...
func (conditionalEventDefinition ConditionalEventDefinition) GetCamundaVariableName() *string {
	return &conditionalEventDefinition.CamundaVariableName
}

/* Elements */

// GetCondition ...
func (conditionalEventDefinition ConditionalEventDefinition) GetCondition() *conditional.Condition {
	return &conditionalEventDefinition.Condition[0]
}
