package models

type CamundaField struct {
	Name          string            `xml:"name,attr,omitempty"`
	CamundaExpr   CamundaExpression `xml:"camunda:expression,innerxml,omitempty"`
	CamundaString CamundaString     `xml:"camunda:string,innerxml,omitempty"`
}

// SetName ...
func (cf *CamundaField) SetName(name string) {
	cf.Name = name
}
