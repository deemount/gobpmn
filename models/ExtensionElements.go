package models

// ExtensionElements ...
type ExtensionElements struct {
	CamundaProperties          []CamundaProperties          `xml:"camunda:properties,omitempty"`
	CamundaFailedJobRetryCycle []CamundaFailedJobRetryCycle `xml:"camunda:failedJobRetryCycle,omitempty"`
	CamundaFormData            []CamundaFormData            `xml:"camunda:formData,omitempty"`
	CamundaInputOutput         []CamundaInputOutput         `xml:"camunda:inputOutput,omitempty"`
	CamundaTaskListener        []CamundaTaskListener        `xml:"camunda:taskListener,omitempty"`
	CamundaExecutionListener   []CamundaExecutionListener   `xml:"camunda:executionListener,omitempty"`
	CamundaIn                  []CamundaIn                  `xml:"camunda:in,omitempty"`
	CamundaOut                 []CamundaOut                 `xml:"camunda:out,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaProperties ...
func (extel *ExtensionElements) SetCamundaProperties() {
	extel.CamundaProperties = make([]CamundaProperties, 1)
}

// SetCamundaFailedJobRetryCycle ...
func (extel *ExtensionElements) SetCamundaFailedJobRetryCycle() {
	extel.CamundaFailedJobRetryCycle = make([]CamundaFailedJobRetryCycle, 1)
}

// SetCamundaFormData ...
func (extel *ExtensionElements) SetCamundaFormData() {
	extel.CamundaFormData = make([]CamundaFormData, 1)
}

// SetCamundaInputOutput ...
func (extel *ExtensionElements) SetCamundaInputOutput() {
	extel.CamundaInputOutput = make([]CamundaInputOutput, 1)
}

// SetCamundaTaskListener ...
func (extel *ExtensionElements) SetCamundaTaskListener(num int) {
	extel.CamundaTaskListener = make([]CamundaTaskListener, num)
}

// SetCamundaExecutionListener ...
func (extel *ExtensionElements) SetCamundaExecutionListener(num int) {
	extel.CamundaExecutionListener = make([]CamundaExecutionListener, num)
}

// SetCamundaIn ...
func (extel *ExtensionElements) SetCamundaIn(num int) {
	extel.CamundaIn = make([]CamundaIn, num)
}

// SetCamundaOut ...
func (extel *ExtensionElements) SetCamundaOut(num int) {
	extel.CamundaOut = make([]CamundaOut, num)
}
