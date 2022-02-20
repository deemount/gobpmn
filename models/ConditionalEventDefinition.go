package models

// ConditionalEventDefinition ...
type ConditionalEventDefinition struct {
	ID                  string      `xml:"id,attr"`
	CamundaVariableName string      `xml:"camunda:variableName"`
	Condition           []Condition `xml:"bpmn:condition,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (ced *ConditionalEventDefinition) SetID(suffix string) {
	ced.ID = "ConditionalEventDefinition_" + suffix
}

/** Camunda **/

// SetCamudnaVariableName ...
func (ced *ConditionalEventDefinition) SetCamundaVariableName(variableName string) {
	ced.CamundaVariableName = variableName
}

/* Elements */

// SetCondition ...
func (ced *ConditionalEventDefinition) SetCondition() {
	ced.Condition = make([]Condition, 1)
}
