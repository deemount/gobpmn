package camunda

// CamundaInputParameterRepository ...
type CamundaInputParameterRepository interface {
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

// CamundaInputParameter ...
type CamundaInputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	CamundaScript           []CamundaScript `xml:"camunda:script,omitempty" json:"script,omitempty"`
	CamundaList             []CamundaList   `xml:"camunda:list,omitempty" json:"list,omitempty"`
	CamundaMap              []CamundaMap    `xml:"camunda:map,omitempty" json:"map,omitempty"`
}

// TCamundaInputParameter ...
type TCamundaInputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	Script                  []CamundaScript `xml:"script,omitempty" json:"script,omitempty"`
	List                    []CamundaList   `xml:"list,omitempty" json:"list,omitempty"`
	Map                     []CamundaMap    `xml:"map,omitempty" json:"map,omitempty"`
}

// NewCamundaInputParameter
func NewCamundaInputParameter() CamundaInputParameterRepository {
	return &CamundaInputParameter{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetLocalVariableName ...
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

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetLocalVariableName ...
func (inputParameter CamundaInputParameter) GetLocalVariableName() *string {
	return &inputParameter.LocalVariableName
}

// GetVariableAssignmentValue ...
func (inputParameter CamundaInputParameter) GetVariableAssignmentValue() *string {
	return &inputParameter.VariableAssignmentValue
}

/* Elements */

/** Camunda **/

// GetCamundaScript ...
func (inputParameter CamundaInputParameter) GetCamundaScript() *CamundaScript {
	return &inputParameter.CamundaScript[0]
}

// GetCamundaList ...
func (inputParameter CamundaInputParameter) GetCamundaList() *CamundaList {
	return &inputParameter.CamundaList[0]
}

// GetCamundaMap ...
func (inputParameter CamundaInputParameter) GetCamundaMap() *CamundaMap {
	return &inputParameter.CamundaMap[0]
}
