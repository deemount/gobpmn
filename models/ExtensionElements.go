package models

// ExtensionElements ...
type ExtensionElements struct {
	Properties                 Properties                 `xml:"bpmn:properties,omitempty"`
	CamundaFailedJobRetryCycle CamundaFailedJobRetryCycle `xml:"camunda:failedJobRetryCycle,innerxml,omitempty"`
	CamundaExecListener        []CamundaExecutionListener `xml:"camunda:executionListener,omitempty"`
}
