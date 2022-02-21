package models

// CamundaOutputParameter ...
type CamundaOutputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty"`
	CamundaScript           []CamundaScript `xml:"camunda:script,omitempty"`
	CamundaList             []CamundaList   `xml:"camunda:list,omitempty"`
	CamundaMap              []CamundaMap    `xml:"camunda:map,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetName ...
func (cop *CamundaOutputParameter) SetLocalVariableName(variable string) {
	cop.LocalVariableName = variable
}

// SetVariableAssignmentValue ...
func (cop *CamundaOutputParameter) SetVariableAssignmentValue(value string) {
	cop.VariableAssignmentValue = value
}

/* Elements */

/** Camunda **/

// SetCamundaScript ...
func (cop *CamundaOutputParameter) SetCamundaScript() {
	cop.CamundaScript = make([]CamundaScript, 1)
}

// SetCamundaList ...
func (cop *CamundaOutputParameter) SetCamundaList() {
	cop.CamundaList = make([]CamundaList, 1)
}

// SetCamundaMap ...
func (cop *CamundaOutputParameter) SetCamundaMap() {
	cop.CamundaMap = make([]CamundaMap, 1)
}
