package camunda

// ExtensionElements ...
type ExtensionElements struct {
	CamundaProperties          CAMUNDA_PROPERTIES_SLC             `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
	CamundaFailedJobRetryCycle CAMUNDA_FAILED_JOB_RETRY_CYCLE_SLC `xml:"camunda:failedJobRetryCycle,omitempty" json:"failedJobRetryCycle,omitempty"`
	CamundaFormData            CAMUNDA_FORM_DATA_SLC              `xml:"camunda:formData,omitempty" json:"formData,omitempty"`
	CamundaInputOutput         CAMUNDA_IO_SLC                     `xml:"camunda:inputOutput,omitempty" json:"inputOutput,omitempty"`
	CamundaTaskListener        CAMUNDA_TASK_LISTENER_SLC          `xml:"camunda:taskListener,omitempty" json:"taskListener,omitempty"`
	CamundaExecutionListener   CAMUNDA_EXECUTION_LISTENER_SLC     `xml:"camunda:executionListener,omitempty" json:"executionListener,omitempty"`
	CamundaIn                  CAMUNDA_IN_SLC                     `xml:"camunda:in,omitempty" json:"in,omitempty"`
	CamundaOut                 CAMUNDA_OUT_SLC                    `xml:"camunda:out,omitempty" json:"out,omitempty"`
}

// TExtensionElements ...
type TExtensionElements struct {
	Properties          TCAMUNDA_PROPERTIES_SLC             `xml:"properties,omitempty" json:"properties,omitempty"`
	FailedJobRetryCycle TCAMUNDA_FAILED_JOB_RETRY_CYCLE_SLC `xml:"failedJobRetryCycle,omitempty" json:"failedJobRetryCycle,omitempty"`
	FormData            TCAMUNDA_FORM_DATA_SLC              `xml:"formData,omitempty" json:"formData,omitempty"`
	InputOutput         TCAMUNDA_IO_SLC                     `xml:"inputOutput,omitempty" json:"inputOutput,omitempty"`
	TaskListener        TCAMUNDA_TASK_LISTENER_SLC          `xml:"taskListener,omitempty" json:"taskListener,omitempty"`
	ExecutionListener   TCAMUNDA_EXECUTION_LISTENER_SLC     `xml:"executionListener,omitempty" json:"executionListener,omitempty"`
	In                  TCAMUNDA_IN_SLC                     `xml:"in,omitempty" json:"in,omitempty"`
	Out                 TCAMUNDA_OUT_SLC                    `xml:"out,omitempty" json:"out,omitempty"`
}

type CoreAttributes struct {
	CamundaAsyncBefore bool `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int  `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
}

type TCoreAttributes struct {
	AsyncBefore bool `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter  bool `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority int  `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
}
