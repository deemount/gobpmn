package camunda

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

	GetCamundaProperties() *CamundaProperties
	GetCamundaFailedJobRetryCycle() *CamundaFailedJobRetryCycle
	GetCamundaFormData() *CamundaFormData
	GetCamundaInputOutput() *CamundaInputOutput
	GetCamundaTaskListener(num int) *CamundaTaskListener
	GetCamundaExecutionListener(num int) *CamundaExecutionListener
	GetCamundaIn(num int) *CamundaIn
	GetCamundaOut(num int) *CamundaOut
}

// ExtensionElements ...
type ExtensionElements struct {
	CamundaProperties          []CamundaProperties          `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
	CamundaFailedJobRetryCycle []CamundaFailedJobRetryCycle `xml:"camunda:failedJobRetryCycle,omitempty" json:"failedJobRetryCycle,omitempty"`
	CamundaFormData            []CamundaFormData            `xml:"camunda:formData,omitempty" json:"formData,omitempty"`
	CamundaInputOutput         []CamundaInputOutput         `xml:"camunda:inputOutput,omitempty" json:"inputOutput,omitempty"`
	CamundaTaskListener        []CamundaTaskListener        `xml:"camunda:taskListener,omitempty" json:"taskListener,omitempty"`
	CamundaExecutionListener   []CamundaExecutionListener   `xml:"camunda:executionListener,omitempty" json:"executionListener,omitempty"`
	CamundaIn                  []CamundaIn                  `xml:"camunda:in,omitempty" json:"in,omitempty"`
	CamundaOut                 []CamundaOut                 `xml:"camunda:out,omitempty" json:"out,omitempty"`
}

// TExtensionElements ...
type TExtensionElements struct {
	CamundaProperties          []CamundaProperties          `xml:"properties,omitempty" json:"properties,omitempty"`
	CamundaFailedJobRetryCycle []CamundaFailedJobRetryCycle `xml:"failedJobRetryCycle,omitempty" json:"failedJobRetryCycle,omitempty"`
	CamundaFormData            []CamundaFormData            `xml:"formData,omitempty" json:"formData,omitempty"`
	CamundaInputOutput         []CamundaInputOutput         `xml:"inputOutput,omitempty" json:"inputOutput,omitempty"`
	CamundaTaskListener        []CamundaTaskListener        `xml:"taskListener,omitempty" json:"taskListener,omitempty"`
	CamundaExecutionListener   []CamundaExecutionListener   `xml:"executionListener,omitempty" json:"executionListener,omitempty"`
	CamundaIn                  []CamundaIn                  `xml:"in,omitempty" json:"in,omitempty"`
	CamundaOut                 []CamundaOut                 `xml:"out,omitempty" json:"out,omitempty"`
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

/**
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaProperties ...
func (extel ExtensionElements) GetCamundaProperties() *CamundaProperties {
	return &extel.CamundaProperties[0]
}

// GetCamundaFailedJobRetryCycle ...
func (extel ExtensionElements) GetCamundaFailedJobRetryCycle() *CamundaFailedJobRetryCycle {
	return &extel.CamundaFailedJobRetryCycle[0]
}

// GetCamundaFormData ...
func (extel ExtensionElements) GetCamundaFormData() *CamundaFormData {
	return &extel.CamundaFormData[0]
}

// GetCamundaInputOutput ...
func (extel ExtensionElements) GetCamundaInputOutput() *CamundaInputOutput {
	return &extel.CamundaInputOutput[0]
}

// GetCamundaTaskListener ...
func (extel ExtensionElements) GetCamundaTaskListener(num int) *CamundaTaskListener {
	return &extel.CamundaTaskListener[num]
}

// GetCamundaExecutionListener ...
func (extel ExtensionElements) GetCamundaExecutionListener(num int) *CamundaExecutionListener {
	return &extel.CamundaExecutionListener[num]
}

// GetCamundaIn ...
func (extel ExtensionElements) GetCamundaIn(num int) *CamundaIn {
	return &extel.CamundaIn[num]
}

// GetCamundaOut ...
func (extel ExtensionElements) GetCamundaOut(num int) *CamundaOut {
	return &extel.CamundaOut[num]
}
