package camunda

// CamundaOutputParameter ...
type CamundaOutputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	CamundaScript           []CamundaScript `xml:"camunda:script,omitempty" json:"script,omitempty"`
	CamundaList             []CamundaList   `xml:"camunda:list,omitempty" json:"list,omitempty"`
	CamundaMap              []CamundaMap    `xml:"camunda:map,omitempty" json:"map,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetName ...
func (outputParameter *CamundaOutputParameter) SetLocalVariableName(variable string) {
	outputParameter.LocalVariableName = variable
}

// SetVariableAssignmentValue ...
func (outputParameter *CamundaOutputParameter) SetVariableAssignmentValue(value string) {
	outputParameter.VariableAssignmentValue = value
}

/* Elements */

/** Camunda **/

// SetCamundaScript ...
func (outputParameter *CamundaOutputParameter) SetCamundaScript() {
	outputParameter.CamundaScript = make([]CamundaScript, 1)
}

// SetCamundaList ...
func (outputParameter *CamundaOutputParameter) SetCamundaList() {
	outputParameter.CamundaList = make([]CamundaList, 1)
}

// SetCamundaMap ...
func (outputParameter *CamundaOutputParameter) SetCamundaMap() {
	outputParameter.CamundaMap = make([]CamundaMap, 1)
}
