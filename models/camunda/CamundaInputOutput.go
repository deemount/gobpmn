package camunda

// CamundaInputOutput ...
type CamundaInputOutput struct {
	CamundaInputParameter  []CamundaInputParameter  `xml:"camunda:inputParameter,omitempty" json:"inputParameter,omitempty"`
	CamundaOutputParameter []CamundaOutputParameter `xml:"camunda:outputParameter,omitempty" json:"outputParameter,omitempty"`
}

/* Elements */

// SetCamundaInputParameter ...
func (cio *CamundaInputOutput) SetCamundaInputParameter(num int) {
	cio.CamundaInputParameter = make([]CamundaInputParameter, num)
}

// SetCamundaOutputParameter ...
func (cio *CamundaInputOutput) SetCamundaOutputParameter(num int) {
	cio.CamundaOutputParameter = make([]CamundaOutputParameter, num)
}
