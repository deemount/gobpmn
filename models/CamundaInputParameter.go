package models

// CamundaInputParameter ...
type CamundaInputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	CamundaScript           []CamundaScript `xml:"camunda:script,omitempty" json:"script,omitempty"`
	CamundaList             []CamundaList   `xml:"camunda:list,omitempty" json:"list,omitempty"`
	CamundaMap              []CamundaMap    `xml:"camunda:map,omitempty" json:"map,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetName ...
func (inputParameter *CamundaInputParameter) SetLocalVariableName(variable string) {
	inputParameter.LocalVariableName = variable
}

// SetVariableAssignmentValue ...
func (inputParameter *CamundaInputParameter) SetVariableAssignmentValue(value string) {
	inputParameter.VariableAssignmentValue = value
}

/* Elements */

/** Camunda **/

// SetCamundaScript ...
func (inputParameter *CamundaInputParameter) SetCamundaScript() {
	inputParameter.CamundaScript = make([]CamundaScript, 1)
}

// SetCamundaList ...
func (inputParameter *CamundaInputParameter) SetCamundaList() {
	inputParameter.CamundaList = make([]CamundaList, 1)
}

// SetCamundaMap ...
func (inputParameter *CamundaInputParameter) SetCamundaMap() {
	inputParameter.CamundaMap = make([]CamundaMap, 1)
}
