package models

type CamundaField struct {
	Name              string              `xml:"name,attr,omitempty"`
	CamundaExpression []CamundaExpression `xml:"camunda:expression,innerxml,omitempty"`
	CamundaString     []CamundaString     `xml:"camunda:string,innerxml,omitempty"`
}

/* Attributes */

// SetName ...
func (cf *CamundaField) SetName(name string) {
	cf.Name = name
}

/* Elements */

/** Camunda **/

// SetCamundaExpression ...

// SetCamundaString ...
