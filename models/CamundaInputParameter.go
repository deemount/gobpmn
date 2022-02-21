package models

// CamundaInputParameter ...
type CamundaInputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty"`
	CamundaScript           []CamundaScript `xml:"camunda:script,omitempty"`
	CamundaList             []CamundaList   `xml:"camunda:list,omitempty"`
	CamundaMap              []CamundaMap    `xml:"camunda:map,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetName ...
func (cip *CamundaInputParameter) SetLocalVariableName(variable string) {
	cip.LocalVariableName = variable
}

// SetVariableAssignmentValue ...
func (cip *CamundaInputParameter) SetVariableAssignmentValue(value string) {
	cip.VariableAssignmentValue = value
}

/* Elements */

/** Camunda **/

// SetCamundaScript ...
func (cip *CamundaInputParameter) SetCamundaScript() {
	cip.CamundaScript = make([]CamundaScript, 1)
}

// SetCamundaList ...
func (cip *CamundaInputParameter) SetCamundaList() {
	cip.CamundaList = make([]CamundaList, 1)
}

// SetCamundaMap ...
func (cip *CamundaInputParameter) SetCamundaMap() {
	cip.CamundaMap = make([]CamundaMap, 1)
}
