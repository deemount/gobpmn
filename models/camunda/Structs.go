package camunda

// ExtensionElements ...
type ExtensionElements struct {
	CamundaConnector           CAMUNDA_CONNECTOR_SLC              `xml:"camunda:connector,omitempty" json:"connector,omitempty"`
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
	CamundaConnector    TCAMUNDA_CONNECTOR_SLC              `xml:"connector,omitempty" json:"connector,omitempty"`
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

// CamundaConnector ...
type CamundaConnector struct {
	CamundaInputOutput []CamundaInputOutput `xml:"camunda:inputOutput,omitempty" json:"inputOutput,omitempty"`
	CamundaConnectorID []CamundaConnectorID `xml:"camunda:connectorId,omitempty" json:"connectorId,omitempty"`
}

// TCamundaConnector ...
type TCamundaConnector struct {
	InputOutput []TCamundaInputOutput `xml:"inputOutput,omitempty" json:"inputOutput,omitempty"`
	ConnectorID []TCamundaConnectorID `xml:"connectorId,omitempty" json:"connectorId,omitempty"`
}

// CamundaConnectorID ...
type CamundaConnectorID struct {
	ID string `xml:",innerxml,omitempty" json:"id"`
}

// TCamundaConnectorID ...
type TCamundaConnectorID struct {
	ID string `xml:",innerxml,omitempty" json:"id"`
}

// CamundaConstraint ...
type CamundaConstraint struct {
	Name   string `xml:"name,attr,omitempty" json:"name"`
	Config string `xml:"config,attr,omitempty" json:"config,omitempty"`
}

// CamundaEntry ...
type CamundaEntry struct {
	Key   string `xml:"key,attr,omitempty" json:"key"`
	Value string `xml:",innerxml,omitempty" json:"value,omitempty"`
}

// CamundaExecutionListener ...
type CamundaExecutionListener struct {
	Class         string          `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event         string          `xml:"event,attr,omitempty" json:"event,omitempty"`
	DelegateExpr  string          `xml:"delegateExpression,attr,omitempty" json:"delegateExpression,omitempty"`
	CamundaScript []CamundaScript `xml:"camunda:script,innerxml,omitempty" json:"script,omitempty"`
	CamundaField  []CamundaField  `xml:"camunda:field,omitempty" json:"field,omitempty"`
}

// TCamundaExecutionListener ...
type TCamundaExecutionListener struct {
	Class        string          `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event        string          `xml:"event,attr,omitempty" json:"event,omitempty"`
	DelegateExpr string          `xml:"delegateExpression,attr,omitempty" json:"delegateExpression,omitempty"`
	Script       []CamundaScript `xml:"script,innerxml,omitempty" json:"script,omitempty"`
	Field        []CamundaField  `xml:"field,omitempty" json:"field,omitempty"`
}

// CamundaExpression ...
type CamundaExpression struct{}

// CamundaFailedJobyRetryCycle ...
type CamundaFailedJobRetryCycle struct{}

// CamundaField ...
type CamundaField struct {
	Name              string              `xml:"name,attr,omitempty" json:"name"`
	CamundaExpression []CamundaExpression `xml:"camunda:expression,omitempty" json:"expression,omitempty"`
	CamundaString     []CamundaString     `xml:"camunda:string,omitempty" json:"string,omitempty"`
}

// TCamundaField ...
type TCamundaField struct {
	Name       string              `xml:"name,attr,omitempty" json:"name"`
	Expression []CamundaExpression `xml:"expression,omitempty" json:"expression,omitempty"`
	String     []CamundaString     `xml:"string,omitempty" json:"string,omitempty"`
}

// CamundaFormData ...
type CamundaFormData struct {
	CamundaFormField []CamundaFormField `xml:"camunda:formField,omitempty" json:"formData,omitempty"`
}

// TCamundaFormData ...
type TCamundaFormData struct {
	FormField []CamundaFormField `xml:"formField,omitempty" json:"formData,omitempty"`
}

// CamundaFormField ...
type CamundaFormField struct {
	ID                string              `xml:"id,attr,omitempty" json:"id"`
	Label             string              `xml:"label,attr,omitempty" json:"label,omitempty"`
	Typ               string              `xml:"type,attr,omitempty" json:"type,omitempty"`
	DefaultValue      string              `xml:"defaultValue,attr,omitempty" json:"defaultValue,omitempty"`
	CamundaProperties []CamundaProperties `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
	CamundaValidation []CamundaValidation `xml:"camunda:validation,omitempty" json:"validation,omitempty"`
}

// TCamundaFormField ...
type TCamundaFormField struct {
	ID           string              `xml:"id,attr,omitempty" json:"id"`
	Label        string              `xml:"label,attr,omitempty" json:"label,omitempty"`
	Typ          string              `xml:"type,attr,omitempty" json:"type,omitempty"`
	DefaultValue string              `xml:"defaultValue,attr,omitempty" json:"defaultValue,omitempty"`
	Properties   []CamundaProperties `xml:"properties,omitempty" json:"properties,omitempty"`
	Validation   []CamundaValidation `xml:"validation,omitempty" json:"validation,omitempty"`
}

// CamundaIn ...
type CamundaIn struct{}

// CamundaInputOutput ...
type CamundaInputOutput struct {
	CamundaInputParameter  []CamundaInputParameter  `xml:"camunda:inputParameter,omitempty" json:"inputParameter,omitempty"`
	CamundaOutputParameter []CamundaOutputParameter `xml:"camunda:outputParameter,omitempty" json:"outputParameter,omitempty"`
}

// TCamundaInputOutput ...
type TCamundaInputOutput struct {
	InputParameter  []CamundaInputParameter  `xml:"inputParameter,omitempty" json:"inputParameter,omitempty"`
	OutputParameter []CamundaOutputParameter `xml:"outputParameter,omitempty" json:"outputParameter,omitempty"`
}

// CamundaInputParameter ...
type CamundaInputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	CamundaScript           []CamundaScript `xml:"camunda:script,omitempty" json:"script,omitempty"`
	CamundaList             []CamundaList   `xml:"camunda:list,omitempty" json:"list,omitempty"`
	CamundaMap              []CamundaMap    `xml:"camunda:map,omitempty" json:"map,omitempty"`
}

// TCamundaInputParameter ...
type TCamundaInputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	Script                  []CamundaScript `xml:"script,omitempty" json:"script,omitempty"`
	List                    []CamundaList   `xml:"list,omitempty" json:"list,omitempty"`
	Map                     []CamundaMap    `xml:"map,omitempty" json:"map,omitempty"`
}

// CamundaList ...
type CamundaList struct {
	CamundaValue []CamundaValue `xml:"camunda:value,omitempty" json:"value,omitempty"`
}

// TCamundaList ...
type TCamundaList struct {
	Value []CamundaValue `xml:"value,omitempty" json:"value,omitempty"`
}

// CamundaMap ...
type CamundaMap struct {
	CamundaEntry []CamundaEntry `xml:"camunda:entry,omitempty" json:"entry,omitempty"`
}

// TCamundaMap ...
type TCamundaMap struct {
	Entry []CamundaEntry `xml:"entry,omitempty" json:"entry,omitempty"`
}

// CamundaOut ...
type CamundaOut struct{}

// CamundaOutputParameter ...
type CamundaOutputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	CamundaScript           []CamundaScript `xml:"camunda:script,omitempty" json:"script,omitempty"`
	CamundaList             []CamundaList   `xml:"camunda:list,omitempty" json:"list,omitempty"`
	CamundaMap              []CamundaMap    `xml:"camunda:map,omitempty" json:"map,omitempty"`
}

// TCamundaOutputParameter ...
type TCamundaOutputParameter struct {
	LocalVariableName       string          `xml:"name,attr,omitempty" json:"localVariableName,omitempty"`
	VariableAssignmentValue string          `xml:",innerxml,omitempty" json:"variableAssignmentValue,omitempty"`
	Script                  []CamundaScript `xml:"script,omitempty" json:"script,omitempty"`
	List                    []CamundaList   `xml:"list,omitempty" json:"list,omitempty"`
	Map                     []CamundaMap    `xml:"map,omitempty" json:"map,omitempty"`
}

// CamundaProperties ...
type CamundaProperties struct {
	CamundaProperty []CamundaProperty `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
}

// TCamundaProperties ...
type TCamundaProperties struct {
	Property []CamundaProperty `xml:"properties,omitempty" json:"properties,omitempty"`
}

// CamundaProperty ...
type CamundaProperty struct {
	Name  string `xml:"name,attr,omitempty" json:"name"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

// CamundaScript ...
type CamundaScript struct {
	ScriptFormat string `xml:"scriptFormat,attr,omitempty"`
}

// CamundaString ...
type CamundaString struct{}

// CamundaTaskListener ...
type CamundaTaskListener struct {
	Class        string         `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event        string         `xml:"event,attr,omitempty" json:"event,omitempty"`
	ListenerID   string         `xml:"id,attr,omitempty" json:"listenerId,omitempty"`
	CamundaField []CamundaField `xml:"camunda:field,omitempty" json:"field,omitempty"`
}

// TCamundaTaskListener ...
type TCamundaTaskListener struct {
	Class      string         `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event      string         `xml:"event,attr,omitempty" json:"event,omitempty"`
	ListenerID string         `xml:"id,attr,omitempty" json:"listenerId,omitempty"`
	Field      []CamundaField `xml:"field,omitempty" json:"field,omitempty"`
}

// CamundaValidation ...
type CamundaValidation struct {
	CamundaConstraint []CamundaConstraint `xml:"camunda:constraint,omitempty" json:"constraint,omitempty"`
}

// TCamundaValidation ...
type TCamundaValidation struct {
	Constraint []CamundaConstraint `xml:"constraint,omitempty" json:"constraint,omitempty"`
}

// CamundaValue ...
type CamundaValue struct {
	Value string `xml:",innerxml,omitempty" json:"value,omitempty"`
}
