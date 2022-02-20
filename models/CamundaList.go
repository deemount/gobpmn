package models

// CamundaList ...
type CamundaList struct {
	CamundaValue []CamundaValue `xml:"camunda:value,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaValue ...
func (clist *CamundaList) SetCamundaValue(num int) {
	clist.CamundaValue = make([]CamundaValue, num)
}
