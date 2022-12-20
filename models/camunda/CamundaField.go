package camunda

// CamundaField ...
type CamundaField struct {
	Name              string              `xml:"name,attr,omitempty" json:"name"`
	CamundaExpression []CamundaExpression `xml:"camunda:expression,omitempty" json:"expression,omitempty"`
	CamundaString     []CamundaString     `xml:"camunda:string,omitempty" json:"string,omitempty"`
}

/* Attributes */

// SetName ...
func (cf *CamundaField) SetName(name string) {
	cf.Name = name
}

/* Elements */

/** Camunda **/

// SetCamundaExpression ...
func (field *CamundaField) SetCamundaExpression(expression string) {
	field.CamundaExpression = make([]CamundaExpression, 1)
}

// SetCamundaString ...
func (field *CamundaField) SetCamundaString(str string) {
	field.CamundaString = make([]CamundaString, 1)
}
