package models

// CamundaExecutionListener ...
type CamundaExecutionListener struct {
	Class         string          `xml:"class,attr,omitempty"`
	Event         string          `xml:"event,attr,omitempty"`
	DelegateExpr  string          `xml:"delegateExpression,attr,omitempty"`
	CamundaScript []CamundaScript `xml:"camunda:script,innerxml,omitempty"`
	CamundaField  []CamundaField  `xml:"camunda:field,omitempty"`
}

/* Attributes */

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

/* Elements */

// SetCamundaScript ...
func (cel *CamundaExecutionListener) SetCamundaScript() {
	cel.CamundaScript = make([]CamundaScript, 1)
}

// SetCamundaField ...
func (cel *CamundaExecutionListener) SetCamundaField(num int) {
	cel.CamundaField = make([]CamundaField, num)
}
