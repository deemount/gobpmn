package models

// ExtensionElements ...
type ExtensionElements struct {
	Properties                 Properties                 `xml:"bpmn:properties,omitempty"`
	CamundaFailedJobRetryCycle CamundaFailedJobRetryCycle `xml:"camunda:failedJobRetryCycle,omitempty"`
	CamundaInputOutput         []CamundaInputOutput       `xml:"camunda:inputOutput,attr,omitempty"`
	CamundaExecListener        []CamundaExecutionListener `xml:"camunda:executionListener,omitempty"`
}

// SetProperties ...
func (extel *ExtensionElements) SetProperties() {
	extel.Properties = Properties{}
}

// SetCamundaFailedJobRetryCycle ...
func (extel *ExtensionElements) SetCamundaFailedJobRetryCycle() {
	extel.CamundaFailedJobRetryCycle = CamundaFailedJobRetryCycle{}
}

// SetCamundaExecutionListener ...
func (extel *ExtensionElements) SetCmaundaExecutionListener(num int) {
	extel.CamundaExecListener = make([]CamundaExecutionListener, num)
}
