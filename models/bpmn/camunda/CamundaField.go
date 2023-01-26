package camunda

import "github.com/deemount/gobpmn/models/bpmn/impl"

// NewCamundaField ...
func NewCamundaField() CamundaFieldRepository {
	return &CamundaField{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetName ...
func (cf *CamundaField) SetName(name string) {
	cf.Name = name
}

/*** Make Elements ***/

/** Camunda **/

// SetCamundaExpression ...
func (cf *CamundaField) SetCamundaExpression() {
	cf.CamundaExpression = make([]CamundaExpression, 1)
}

// SetCamundaString ...
func (cf *CamundaField) SetCamundaString() {
	cf.CamundaString = make([]CamundaString, 1)
}

/*
 * Default Getters
 */

/* Attributes */

// GetName ...
func (cf CamundaField) GetName() impl.STR_PTR {
	return &cf.Name
}

/* Elements */

/** Camunda **/

// GetCamundaExpression ...
func (cf CamundaField) GetCamundaExpression() *CamundaExpression {
	return &cf.CamundaExpression[0]
}

// GetCamundaString ...
func (cf CamundaField) GetCamundaString() *CamundaString {
	return &cf.CamundaString[0]
}
