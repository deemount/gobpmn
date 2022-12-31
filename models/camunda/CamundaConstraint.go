package camunda

import "github.com/deemount/gobpmn/models/compulsion"

// NewCamundaConstraint ...
func NewCamunadConstraint() CamundaConstraintRepository {
	return &CamundaConstraint{}
}

/*
 * Default Getters
 */

/* Attributes */

// SetName ...
func (cconstraint *CamundaConstraint) SetName(name string) {
	cconstraint.Name = name
}

// SetConfig ...
func (cconstraint *CamundaConstraint) SetConfig(config string) {
	cconstraint.Config = config
}

/*
 * Default Getters
 */

/* Attributes */

// GetName ...
func (cconstraint CamundaConstraint) GetName() compulsion.STR_PTR {
	return &cconstraint.Name
}

// GetConfig ...
func (cconstraint CamundaConstraint) GetConfig() *string {
	return &cconstraint.Config
}
