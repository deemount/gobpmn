package camunda

// NewCamundaValidation ...
func NewCamundaValidation() CamundaValidationRepository {
	return &CamundaValidation{}
}

/*
 * Default Setters
 */

/*** Make Elements ***/

/** Camunda **/

// SetCamundaConstraint ...
func (validation *CamundaValidation) SetCamundaConstraint(num int) {
	validation.CamundaConstraint = make([]CamundaConstraint, num)
}

/*
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaConstraint ...
func (validation CamundaValidation) GetCamundaConstraint(num int) *CamundaConstraint {
	return &validation.CamundaConstraint[num]
}
