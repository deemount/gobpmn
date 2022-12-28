package definitions

import (
	"fmt"

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
	conditionalEventDefinition.Condition = make([]conditional.Condition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (conditionalEventDefinition ConditionalEventDefinition) GetID() STR_PTR {
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
