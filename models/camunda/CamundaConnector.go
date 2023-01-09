package camunda

// NewCamundaConnector ...
func NewCamundaConnector() CamundaConnectorRepository {
	return &CamundaConnector{}
}

/*
 * Default Setters
 */

/*** Make Elements ***/

/** Camunda **/

// SetCamundaInputOutput ...
func (cconnector *CamundaConnector) SetCamundaInputOutput() {
	cconnector.CamundaInputOutput = make([]CamundaInputOutput, 1)
}

// SetCamundaInputOutput ...
func (cconnector *CamundaConnector) SetCamundaConnectorID() {
	cconnector.CamundaConnectorID = make([]CamundaConnectorID, 1)
}

/*
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaInputOutput ...
func (cconnector CamundaConnector) GetCamundaInputOutput() *CamundaInputOutput {
	return &cconnector.CamundaInputOutput[0]
}

// GetCamundaConnectorID ...
func (cconnector CamundaConnector) GetCamundaConnectorID() *CamundaConnectorID {
	return &cconnector.CamundaConnectorID[0]
}
