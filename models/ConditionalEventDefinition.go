package models

// ConditionalEventDefinition ...
type ConditionalEventDefinition struct {
	ID                  string      `xml:"id,attr" json:"id"`
	CamundaVariableName string      `xml:"camunda:variableName" json:"variableName,omitempty"`
	Condition           []Condition `xml:"bpmn:condition,omitempty" json:"condition,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetID(suffix string) {
	conditionalEventDefinition.ID = "ConditionalEventDefinition_" + suffix
}

/** Camunda **/

// SetCamudnaVariableName ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCamundaVariableName(variableName string) {
	conditionalEventDefinition.CamundaVariableName = variableName
}

/* Elements */

// SetCondition ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCondition() {
	conditionalEventDefinition.Condition = make([]Condition, 1)
}
