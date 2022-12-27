package camunda

// CamundaInputOutputRepository ...
type CamundaInputOutputRepository interface {
	SetCamundaInputParameter(num int)
	GetCamundaInputParameter(num int) *CamundaInputParameter
	SetCamundaOutputParameter(num int)
	GetCamundaOutputParameter(num int) *CamundaOutputParameter
}

// CamundaInputOutput ...
type CamundaInputOutput struct {
	CamundaInputParameter  []CamundaInputParameter  `xml:"camunda:inputParameter,omitempty" json:"inputParameter,omitempty"`
	CamundaOutputParameter []CamundaOutputParameter `xml:"camunda:outputParameter,omitempty" json:"outputParameter,omitempty"`
}

// TCamundaInputOutput ...
type TCamundaInputOutput struct {
	InputParameter  []CamundaInputParameter  `xml:"inputParameter,omitempty" json:"inputParameter,omitempty"`
	OutputParameter []CamundaOutputParameter `xml:"outputParameter,omitempty" json:"outputParameter,omitempty"`
}

// NewCamundaInputOutput ...
func NewCamundaInputOutput() CamundaInputOutputRepository {
	return &CamundaInputOutput{}
}

/*
 * Default Setters
 */

/*** Make Elements ***/

/** Camunda **/

// SetCamundaInputParameter ...
func (cio *CamundaInputOutput) SetCamundaInputParameter(num int) {
	cio.CamundaInputParameter = make([]CamundaInputParameter, num)
}

// SetCamundaOutputParameter ...
func (cio *CamundaInputOutput) SetCamundaOutputParameter(num int) {
	cio.CamundaOutputParameter = make([]CamundaOutputParameter, num)
}

/*
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaInputParameter ...
func (cio CamundaInputOutput) GetCamundaInputParameter(num int) *CamundaInputParameter {
	return &cio.CamundaInputParameter[num]
}

// GetCamundaOutputParameter ...
func (cio CamundaInputOutput) GetCamundaOutputParameter(num int) *CamundaOutputParameter {
	return &cio.CamundaOutputParameter[num]
}
