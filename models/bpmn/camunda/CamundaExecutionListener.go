package camunda

import "github.com/deemount/gobpmn/models/bpmn/impl"

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
func (executionListener CamundaExecutionListener) GetClass() impl.STR_PTR {
	return &executionListener.Class
}

// GetEvent ...
func (executionListener CamundaExecutionListener) GetEvent() impl.STR_PTR {
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
