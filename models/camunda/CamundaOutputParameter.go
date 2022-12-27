package camunda

// CamundaOutputParameter ...
type CamundaOutputParameterRepository interface {
	CamundaBaseScriptElements
	SetLocalVariableName(variable string)
	GetLocalVariableName() *string
	SetVariableAssignmentValue(value string)
	GetVariableAssignmentValue() *string
	SetCamundaList()
	GetCamundaList() *CamundaList
	SetCamundaMap()
	GetCamundaMap() *CamundaMap
}

// CamundaOutputParameter ...
type CamundaOutputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	CamundaScript           []CamundaScript `xml:"camunda:script,omitempty" json:"script,omitempty"`
	CamundaList             []CamundaList   `xml:"camunda:list,omitempty" json:"list,omitempty"`
	CamundaMap              []CamundaMap    `xml:"camunda:map,omitempty" json:"map,omitempty"`
}

// TCamundaOutputParameter ...
type TCamundaOutputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	Script                  []CamundaScript `xml:"script,omitempty" json:"script,omitempty"`
	List                    []CamundaList   `xml:"list,omitempty" json:"list,omitempty"`
	Map                     []CamundaMap    `xml:"map,omitempty" json:"map,omitempty"`
}

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
func (outputParameter CamundaOutputParameter) GetLocalVariableName() *string {
	return &outputParameter.LocalVariableName
}

// GetVariableAssignmentValue ...
func (outputParameter CamundaOutputParameter) GetVariableAssignmentValue() *string {
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
