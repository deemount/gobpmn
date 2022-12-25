package attributes

import "github.com/deemount/gobpmn/models/camunda"

// ExtensionElementsRepository ...
type ExtensionElementsRepository interface {
	SetCamundaProperties()
	SetCamundaFailedJobRetryCycle()
	SetCamundaFormData()
	SetCamundaInputOutput()
	SetCamundaTaskListener(num int)
	SetCamundaExecutionListener(num int)
	SetCamundaIn(num int)
	SetCamundaOut(num int)

	GetCamundaProperties() *camunda.CamundaProperties
	GetCamundaFailedJobRetryCycle() *camunda.CamundaFailedJobRetryCycle
	GetCamundaFormData() *camunda.CamundaFormData
	GetCamundaInputOutput() *camunda.CamundaInputOutput
	GetCamundaTaskListener(num int) *camunda.CamundaTaskListener
	GetCamundaExecutionListener(num int) *camunda.CamundaExecutionListener
	GetCamundaIn(num int) *camunda.CamundaIn
	GetCamundaOut(num int) *camunda.CamundaOut
}

// ExtensionElements ...
type ExtensionElements struct {
	CamundaProperties          []camunda.CamundaProperties          `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
	CamundaFailedJobRetryCycle []camunda.CamundaFailedJobRetryCycle `xml:"camunda:failedJobRetryCycle,omitempty" json:"failedJobRetryCycle,omitempty"`
	CamundaFormData            []camunda.CamundaFormData            `xml:"camunda:formData,omitempty" json:"formData,omitempty"`
	CamundaInputOutput         []camunda.CamundaInputOutput         `xml:"camunda:inputOutput,omitempty" json:"inputOutput,omitempty"`
	CamundaTaskListener        []camunda.CamundaTaskListener        `xml:"camunda:taskListener,omitempty" json:"taskListener,omitempty"`
	CamundaExecutionListener   []camunda.CamundaExecutionListener   `xml:"camunda:executionListener,omitempty" json:"executionListener,omitempty"`
	CamundaIn                  []camunda.CamundaIn                  `xml:"camunda:in,omitempty" json:"in,omitempty"`
	CamundaOut                 []camunda.CamundaOut                 `xml:"camunda:out,omitempty" json:"out,omitempty"`
}

// TExtensionElements ...
type TExtensionElements struct {
	CamundaProperties          []camunda.CamundaProperties          `xml:"properties,omitempty" json:"properties,omitempty"`
	CamundaFailedJobRetryCycle []camunda.CamundaFailedJobRetryCycle `xml:"failedJobRetryCycle,omitempty" json:"failedJobRetryCycle,omitempty"`
	CamundaFormData            []camunda.CamundaFormData            `xml:"formData,omitempty" json:"formData,omitempty"`
	CamundaInputOutput         []camunda.CamundaInputOutput         `xml:"inputOutput,omitempty" json:"inputOutput,omitempty"`
	CamundaTaskListener        []camunda.CamundaTaskListener        `xml:"taskListener,omitempty" json:"taskListener,omitempty"`
	CamundaExecutionListener   []camunda.CamundaExecutionListener   `xml:"executionListener,omitempty" json:"executionListener,omitempty"`
	CamundaIn                  []camunda.CamundaIn                  `xml:"in,omitempty" json:"in,omitempty"`
	CamundaOut                 []camunda.CamundaOut                 `xml:"out,omitempty" json:"out,omitempty"`
}

func NewExtensionElements() ExtensionElementsRepository {
	return &ExtensionElements{}
}

/**
 * Default Setters
 */

/*** Make Elements ***/

/** Camunda **/

// SetCamundaProperties ...
func (extel *ExtensionElements) SetCamundaProperties() {
	extel.CamundaProperties = make([]camunda.CamundaProperties, 1)
}

// SetCamundaFailedJobRetryCycle ...
func (extel *ExtensionElements) SetCamundaFailedJobRetryCycle() {
	extel.CamundaFailedJobRetryCycle = make([]camunda.CamundaFailedJobRetryCycle, 1)
}

// SetCamundaFormData ...
func (extel *ExtensionElements) SetCamundaFormData() {
	extel.CamundaFormData = make([]camunda.CamundaFormData, 1)
}

// SetCamundaInputOutput ...
func (extel *ExtensionElements) SetCamundaInputOutput() {
	extel.CamundaInputOutput = make([]camunda.CamundaInputOutput, 1)
}

// SetCamundaTaskListener ...
func (extel *ExtensionElements) SetCamundaTaskListener(num int) {
	extel.CamundaTaskListener = make([]camunda.CamundaTaskListener, num)
}

// SetCamundaExecutionListener ...
func (extel *ExtensionElements) SetCamundaExecutionListener(num int) {
	extel.CamundaExecutionListener = make([]camunda.CamundaExecutionListener, num)
}

// SetCamundaIn ...
func (extel *ExtensionElements) SetCamundaIn(num int) {
	extel.CamundaIn = make([]camunda.CamundaIn, num)
}

// SetCamundaOut ...
func (extel *ExtensionElements) SetCamundaOut(num int) {
	extel.CamundaOut = make([]camunda.CamundaOut, num)
}

/**
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaProperties ...
func (extel ExtensionElements) GetCamundaProperties() *camunda.CamundaProperties {
	return &extel.CamundaProperties[0]
}

// GetCamundaFailedJobRetryCycle ...
func (extel ExtensionElements) GetCamundaFailedJobRetryCycle() *camunda.CamundaFailedJobRetryCycle {
	return &extel.CamundaFailedJobRetryCycle[0]
}

// GetCamundaFormData ...
func (extel ExtensionElements) GetCamundaFormData() *camunda.CamundaFormData {
	return &extel.CamundaFormData[0]
}

// GetCamundaInputOutput ...
func (extel ExtensionElements) GetCamundaInputOutput() *camunda.CamundaInputOutput {
	return &extel.CamundaInputOutput[0]
}

// GetCamundaTaskListener ...
func (extel ExtensionElements) GetCamundaTaskListener(num int) *camunda.CamundaTaskListener {
	return &extel.CamundaTaskListener[num]
}

// GetCamundaExecutionListener ...
func (extel ExtensionElements) GetCamundaExecutionListener(num int) *camunda.CamundaExecutionListener {
	return &extel.CamundaExecutionListener[num]
}

// GetCamundaIn ...
func (extel ExtensionElements) GetCamundaIn(num int) *camunda.CamundaIn {
	return &extel.CamundaIn[num]
}

// GetCamundaOut ...
func (extel ExtensionElements) GetCamundaOut(num int) *camunda.CamundaOut {
	return &extel.CamundaOut[num]
}
