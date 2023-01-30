package camunda

import "github.com/deemount/gobpmn/models/bpmn/impl"

// NewCamundaOutputParameter ...
func NewCamundaOutputParameter() CamundaOutputParameterRepository {
	return &CamundaOutputParameter{}
}

/*
 * Default Setters
 */

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

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetName ...
func (outputParameter CamundaOutputParameter) GetLocalVariableName() impl.STR_PTR {
	return &outputParameter.LocalVariableName
}

// GetVariableAssignmentValue ...
func (outputParameter CamundaOutputParameter) GetVariableAssignmentValue() impl.STR_PTR {
	return &outputParameter.VariableAssignmentValue
}

/* Elements */

/** Camunda **/

// GetCamundaScript ...
func (outputParameter CamundaOutputParameter) GetCamundaScript() *CamundaScript {
	return &outputParameter.CamundaScript[0]
}

// GetCamundaList ...
func (outputParameter CamundaOutputParameter) GetCamundaList() *CamundaList {
	return &outputParameter.CamundaList[0]
}

// GetCamundaMap ...
func (outputParameter CamundaOutputParameter) GetCamundaMap() *CamundaMap {
	return &outputParameter.CamundaMap[0]
}
