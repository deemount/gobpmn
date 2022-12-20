package camunda

// CamundaList ...
type CamundaList struct {
	CamundaValue []CamundaValue `xml:"camunda:value,omitempty" json:"value,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaValue ...
func (list *CamundaList) SetCamundaValue(num int) {
	list.CamundaValue = make([]CamundaValue, num)
}
