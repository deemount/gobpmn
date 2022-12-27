package camunda

// CamundaExecutionListener ...
type CamundaExecutionListenerRepository interface {
	CamundaBaseEvent
	CamundaBaseClass
	CamundaBaseScriptElements
	SetDelegateExpression(expr string)
	GetDelegateExpression() *string
	SetCamundaField(num int)
	GetCamundaField(num int) *CamundaField
}

// CamundaExecutionListener ...
type CamundaExecutionListener struct {
	Class         string          `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event         string          `xml:"event,attr,omitempty" json:"event,omitempty"`
	DelegateExpr  string          `xml:"delegateExpression,attr,omitempty" json:"delegateExpression,omitempty"`
	CamundaScript []CamundaScript `xml:"camunda:script,innerxml,omitempty" json:"script,omitempty"`
	CamundaField  []CamundaField  `xml:"camunda:field,omitempty" json:"field,omitempty"`
}

// TCamundaExecutionListener ...
type TCamundaExecutionListener struct {
	Class        string          `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event        string          `xml:"event,attr,omitempty" json:"event,omitempty"`
	DelegateExpr string          `xml:"delegateExpression,attr,omitempty" json:"delegateExpression,omitempty"`
	Script       []CamundaScript `xml:"script,innerxml,omitempty" json:"script,omitempty"`
	Field        []CamundaField  `xml:"field,omitempty" json:"field,omitempty"`
}

// NewCamundaExecutionListener ...
func NewCamundaExecutionListener() CamundaExecutionListenerRepository {
	return &CamundaExecutionListener{}
}

/*
 * Default Setters
 */

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

/*** Make Elements ***/

// SetCamundaScript ...
func (executionListener *CamundaExecutionListener) SetCamundaScript() {
	executionListener.CamundaScript = make([]CamundaScript, 1)
}

// SetCamundaField ...
func (executionListener *CamundaExecutionListener) SetCamundaField(num int) {
	executionListener.CamundaField = make([]CamundaField, num)
}

/*
 * Default Getters
 */

/* Attributes */

// GetClass ..
func (executionListener CamundaExecutionListener) GetClass() STR_PTR {
	return &executionListener.Class
}

// GetEvent ...
func (executionListener CamundaExecutionListener) GetEvent() STR_PTR {
	return &executionListener.Event
}

// GetDelegateExpression ...
func (executionListener CamundaExecutionListener) GetDelegateExpression() *string {
	return &executionListener.DelegateExpr
}

/* Elements */

// GetCamundaScript ...
func (executionListener CamundaExecutionListener) GetCamundaScript() *CamundaScript {
	return &executionListener.CamundaScript[0]
}

// GetCamundaField ...
func (executionListener CamundaExecutionListener) GetCamundaField(num int) *CamundaField {
	return &executionListener.CamundaField[num]
}
