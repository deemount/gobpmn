package models

// CamundaProperties ...
type CamundaProperties struct {
	CamundaProperty []CamundaProperty `xml:"camunda:properties,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaProperties ...
func (cproperties *CamundaProperties) SetName(num int) {
	cproperties.CamundaProperty = make([]CamundaProperty, num)
}
