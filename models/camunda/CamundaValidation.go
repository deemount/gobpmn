package camunda

// CamundaValidation ...
type CamundaValidation struct {
	CamundaConstraint []CamundaConstraint `xml:"camunda:constraint,omitempty" json:"constraint,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaConstraint ...
func (validation *CamundaValidation) SetCamundaConstraint(num int) {
	validation.CamundaConstraint = make([]CamundaConstraint, num)
}
