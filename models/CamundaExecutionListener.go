package models

// CamundaExecutionListener ...
type CamundaExecutionListener struct {
	Class         string         `xml:"class,attr,omitempty"`
	Event         string         `xml:"event,attr,omitempty"`
	DelegateExpr  string         `xml:"delegateExpression,attr,omitempty"`
	CamundaScript CamundaScript  `xml:"camunda:script,innerxml,omitempty"`
	CamundaField  []CamundaField `xml:"camunda:field,omitempty"`
}

// SetClass ..
func (cel *CamundaExecutionListener) SetClass(cls string) {
	cel.Class = cls
}

// SetEvent ...
func (cel *CamundaExecutionListener) SetEvent(event string) {
	cel.Event = event
}

// SetDelegateExpression ...
func (cel *CamundaExecutionListener) SetDelegateExpression(expr string) {
	cel.DelegateExpr = expr
}
