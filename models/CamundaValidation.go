package models

// CamundaValidation ...
type CamundaValidation struct {
	CamundaConstraint []CamundaConstraint `xml:"camunda:constraint,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaConstraint ...
func (cvalid *CamundaValidation) SetCamundaConstraint(num int) {
	cvalid.CamundaConstraint = make([]CamundaConstraint, num)
}
