package models

// CamundaExecutionListener ...
type CamundaExecutionListener struct {
	Class         string          `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event         string          `xml:"event,attr,omitempty" json:"event,omitempty"`
	DelegateExpr  string          `xml:"delegateExpression,attr,omitempty" json:"delegateExpression,omitempty"`
	CamundaScript []CamundaScript `xml:"camunda:script,innerxml,omitempty" json:"script,omitempty"`
	CamundaField  []CamundaField  `xml:"camunda:field,omitempty" json:"field,omitempty"`
}

/* Attributes */

// SetClass ..
func (executionListener *CamundaExecutionListener) SetClass(cls string) {
	executionListener.Class = cls
}

// SetEvent ...
func (executionListener *CamundaExecutionListener) SetEvent(event string) {
	executionListener.Event = event
}

// SetDelegateExpression ...
func (executionListener *CamundaExecutionListener) SetDelegateExpression(expr string) {
	executionListener.DelegateExpr = expr
}

/* Elements */

// SetCamundaScript ...
func (executionListener *CamundaExecutionListener) SetCamundaScript() {
	executionListener.CamundaScript = make([]CamundaScript, 1)
}

// SetCamundaField ...
func (executionListener *CamundaExecutionListener) SetCamundaField(num int) {
	executionListener.CamundaField = make([]CamundaField, num)
}
