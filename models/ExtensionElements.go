package models

// ExtensionElements ...
type ExtensionElements struct {
	Properties                 []Properties                 `xml:"bpmn:properties,omitempty"`
	CamundaFailedJobRetryCycle []CamundaFailedJobRetryCycle `xml:"camunda:failedJobRetryCycle,omitempty"`
	CamundaInputOutput         []CamundaInputOutput         `xml:"camunda:inputOutput,attr,omitempty"`
	CamundaExecListener        []CamundaExecutionListener   `xml:"camunda:executionListener,omitempty"`
}

/* Elements */

// SetProperties ...
func (extel *ExtensionElements) SetProperties() {
	extel.Properties = make([]Properties, 1)
}

// SetCamundaFailedJobRetryCycle ...
func (extel *ExtensionElements) SetCamundaFailedJobRetryCycle() {
	extel.CamundaFailedJobRetryCycle = make([]CamundaFailedJobRetryCycle, 1)
}

// SetCamundaInputOutput ...
func (extel *ExtensionElements) SetCamundaInputOutput() {
	extel.CamundaInputOutput = make([]CamundaInputOutput, 1)
}

// SetCamundaExecutionListener ...
func (extel *ExtensionElements) SetCmaundaExecutionListener(num int) {
	extel.CamundaExecListener = make([]CamundaExecutionListener, num)
}
