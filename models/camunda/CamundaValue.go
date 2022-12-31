package camunda

import "github.com/deemount/gobpmn/models/compulsion"

// NewCamundaValue ...
func NewCamundaValue() CamundaValueRepository {
	return &CamundaValue{}
}

/*
 * Default Setters
 */

/* Content */

// SetValue ...
func (value *CamundaValue) SetValue(val string) {
	value.Value = val
}

/*
 * Default Getters
 */

/* Content */

// GetValue ...
func (value CamundaValue) GetValue() compulsion.STR_PTR {
	return &value.Value
}
