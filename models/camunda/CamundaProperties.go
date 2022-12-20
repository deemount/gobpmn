package camunda

// CamundaProperties ...
type CamundaProperties struct {
	CamundaProperty []CamundaProperty `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaProperties ...
func (properties *CamundaProperties) SetName(num int) {
	properties.CamundaProperty = make([]CamundaProperty, num)
}
