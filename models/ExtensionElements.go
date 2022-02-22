package models

// ExtensionElements ...
type ExtensionElements struct {
	CamundaProperties          []CamundaProperties          `xml:"camunda:properties,omitempty"`
	CamundaFailedJobRetryCycle []CamundaFailedJobRetryCycle `xml:"camunda:failedJobRetryCycle,omitempty"`
	CamundaFormData            []CamundaFormData            `xml:"camunda:formData,omitempty"`
	CamundaInputOutput         []CamundaInputOutput         `xml:"camunda:inputOutput,attr,omitempty"`
	CamundaExecListener        []CamundaExecutionListener   `xml:"camunda:executionListener,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaFailedJobRetryCycle ...
func (extel *ExtensionElements) SetCamundaFailedJobRetryCycle() {
	extel.CamundaFailedJobRetryCycle = make([]CamundaFailedJobRetryCycle, 1)
}

// SetCamundaProperties ...
func (extel *ExtensionElements) SetCamundaProperties() {
	extel.CamundaProperties = make([]CamundaProperties, 1)
}

// SetCamundaInputOutput ...
func (extel *ExtensionElements) SetCamundaInputOutput() {
	extel.CamundaInputOutput = make([]CamundaInputOutput, 1)
}

// SetCamundaExecutionListener ...
func (extel *ExtensionElements) SetCmaundaExecutionListener(num int) {
	extel.CamundaExecListener = make([]CamundaExecutionListener, num)
}
