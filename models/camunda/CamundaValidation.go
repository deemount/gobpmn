package camunda

// CamundaValidationRepository ...
type CamundaValidationRepository interface {
	SetCamundaConstraint(num int)
	GetCamundaConstraint(num int) *CamundaConstraint
}

// CamundaValidation ...
type CamundaValidation struct {
	CamundaConstraint []CamundaConstraint `xml:"camunda:constraint,omitempty" json:"constraint,omitempty"`
}

// TCamundaValidation ...
type TCamundaValidation struct {
	Constraint []CamundaConstraint `xml:"constraint,omitempty" json:"constraint,omitempty"`
}

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
