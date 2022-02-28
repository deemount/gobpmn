package models

type CamundaField struct {
	Name              string              `xml:"name,attr,omitempty" json:"name"`
	CamundaExpression []CamundaExpression `xml:"camunda:expression,innerxml,omitempty" json:"expression,omitempty"`
	CamundaString     []CamundaString     `xml:"camunda:string,innerxml,omitempty" json:"string,omitempty"`
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
