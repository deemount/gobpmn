package camunda

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
