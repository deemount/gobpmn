package models

// CamundaProperty ...
type CamundaProperty struct {
	CamundaProperties []CamundaProperties `xml:"camunda:properties,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaProperties ...
func (cproperty *CamundaProperty) SetName(num int) {
	cproperty.CamundaProperties = make([]CamundaProperties, num)
}
