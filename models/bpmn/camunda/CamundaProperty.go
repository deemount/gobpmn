package camunda

import "github.com/deemount/gobpmn/models/bpmn/impl"

// NewCamundaProperty ...
func NewCamundaProperty() CamundaPropertyRepository {
	return &CamundaProperty{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetName ...
func (property *CamundaProperty) SetName(name string) {
	property.Name = name
}

// SetValue ...
func (property *CamundaProperty) SetValue(value string) {
	property.Value = value
}

/*
 * Default Getters
 */

/* Attributes */

// GetName ...
func (property CamundaProperty) GetName() impl.STR_PTR {
	return &property.Name
}

// GetValue ...
func (property CamundaProperty) GetValue() impl.STR_PTR {
	return &property.Value
}
