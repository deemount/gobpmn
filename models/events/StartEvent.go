package events

import (
	"fmt"

	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// StartEventRepository ...
type StartEventRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)
	SetIsInterrupting(isInterrupt bool)

	SetCamundaFormKey(key string)
	SetCamundaFormRef(ref string)
	SetCamundaFormRefBinding(bind string)
	SetCamundaFormRefVersion(version string)
	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaJobPriority(priority int)
	SetCamundaInitiator(initiator string)

	SetExtensionElements()
	SetConditionalEventDefinition()
	SetMsgEventDefinition()
	SetTimerEventDefinition()
	SetOutgoing(num int)

	String() string

	GetID() *string
	GetName() *string
	GetIsInterrupting() *bool

	GetCamundaFormKey() *string
	GetCamundaFormRef() *string
	GetCamundaFormRefBinding() *string
	GetCamundaFormRefVersion() *string
	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaJobPriority() *int
	GetCamundaInitiator() *string

	GetExtensionElements() *camunda.ExtensionElements
	GetConditionalEventDefinition() *ConditionalEventDefinition
	GetMsgEventDefinition() *MessageEventDefinition
	GetTimerEventDefinition() *TimerEventDefinition
	GetOutgoing(num int) *marker.Outgoing
}

// StartEvent ...
type StartEvent struct {
	ID                    string                       `xml:"id,attr" json:"id"`
	Name                  string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	IsInterrupting        bool                         `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	CamundaFormKey        string                       `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaFormRef        string                       `xml:"camunda:formRef,attr,omitempty" json:"formRef,omitempty"`
	CamundaFormRefBind    string                       `xml:"camunda:formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	CamundaFormRefVersion string                       `xml:"camunda:formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	CamundaAsyncBefore    bool                         `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter     bool                         `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority    int                          `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	CamundaInit           string                       `xml:"camunda:initiator,attr,omitempty" json:"init,omitempty"`
	ExtensionElements     []camunda.ExtensionElements  `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	ConditionalEventDef   []ConditionalEventDefinition `xml:"bpmn:conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MsgEventDef           []MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef         []TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	Outgoing              []marker.Outgoing            `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	ID                  string                        `xml:"id,attr" json:"id"`
	Name                string                        `xml:"name,attr,omitempty" json:"name,omitempty"`
	IsInterrupting      bool                          `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	FormKey             string                        `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	FormRef             string                        `xml:"formRef,attr,omitempty" json:"formRef,omitempty"`
	FormRefBind         string                        `xml:"formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	FormRefVersion      string                        `xml:"formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	AsyncBefore         bool                          `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter          bool                          `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority         int                           `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Init                string                        `xml:"initiator,attr,omitempty" json:"init,omitempty"`
	ExtensionElements   []camunda.TExtensionElements  `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	ConditionalEventDef []TConditionalEventDefinition `xml:"conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MsgEventDef         []MessageEventDefinition      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef       []TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	Outgoing            []marker.Outgoing             `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewStartEvent() StartEventRepository {
	return &StartEvent{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (startEvent *StartEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	case "counter":
		startEvent.ID = fmt.Sprintf("StartEvent_%d", suffix)
		break
	case "event":
		startEvent.ID = fmt.Sprintf("Event_%s", suffix)
		break
	case "id":
		startEvent.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (startEvent *StartEvent) SetName(name string) {
	startEvent.Name = name
}

// SetIsInterrupting ...
func (startEvent *StartEvent) SetIsInterrupting(isInterrupt bool) {
	startEvent.IsInterrupting = isInterrupt
}

/** Camunda **/

// SetCamundaFormKey ...
func (startEvent *StartEvent) SetCamundaFormKey(key string) {
	startEvent.CamundaFormKey = key
}

// SetCamundaFormRef ...
func (startEvent *StartEvent) SetCamundaFormRef(ref string) {
	startEvent.CamundaFormRef = ref
}

// SetCamundaFormRefBinding ...
// possible arguments: latest, deployment, version
// version is dependent to attribute camunda:formRefVersion
func (startEvent *StartEvent) SetCamundaFormRefBinding(bind string) {
	startEvent.CamundaFormRefBind = bind
}

// SetCamundaFormRefVersion ...
func (startEvent *StartEvent) SetCamundaFormRefVersion(version string) {
	startEvent.CamundaFormRefVersion = version
}

// SetCamundaAsyncBefore ...
func (startEvent *StartEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	startEvent.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (startEvent *StartEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	startEvent.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (startEvent *StartEvent) SetCamundaJobPriority(priority int) {
	startEvent.CamundaJobPriority = priority
}

// SetCamundaInitiator ...
func (startEvent *StartEvent) SetCamundaInitiator(initiator string) {
	startEvent.CamundaInit = initiator
}

/*** Make Elements ***/

// SetExtensionElements ...
func (startEvent *StartEvent) SetExtensionElements() {
	startEvent.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetConditionalEventDefinition ...
func (startEvent *StartEvent) SetConditionalEventDefinition() {
	startEvent.ConditionalEventDef = make([]ConditionalEventDefinition, 1)
}

// SetMsgEventDefinition ...
func (startEvent *StartEvent) SetMsgEventDefinition() {
	startEvent.MsgEventDef = make([]MessageEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (startEvent *StartEvent) SetTimerEventDefinition() {
	startEvent.TimerEventDef = make([]TimerEventDefinition, 1)
}

// SetOutgoing ...
func (startEvent *StartEvent) SetOutgoing(num int) {
	startEvent.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default String
 */

// String ...
func (startEvent StartEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", startEvent.ID, startEvent.Name)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (startEvent StartEvent) GetID() *string {
	return &startEvent.ID
}

// GetName ...
func (startEvent StartEvent) GetName() *string {
	return &startEvent.Name
}

// SetIsInterrupting ...
func (startEvent StartEvent) GetIsInterrupting() *bool {
	return &startEvent.IsInterrupting
}

/** Camunda **/

// SetCamundaFormKey ...
func (startEvent StartEvent) GetCamundaFormKey() *string {
	return &startEvent.CamundaFormKey
}

// GetCamundaFormRef ...
func (startEvent StartEvent) GetCamundaFormRef() *string {
	return &startEvent.CamundaFormRef
}

// GetCamundaFormRefBinding ...
func (startEvent StartEvent) GetCamundaFormRefBinding() *string {
	return &startEvent.CamundaFormRefBind
}

// GetCamundaFormRefVersion ...
func (startEvent StartEvent) GetCamundaFormRefVersion() *string {
	return &startEvent.CamundaFormRefVersion
}

// GetCamundaAsyncBefore ...
func (startEvent StartEvent) GetCamundaAsyncBefore() *bool {
	return &startEvent.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (startEvent StartEvent) GetCamundaAsyncAfter() *bool {
	return &startEvent.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (startEvent StartEvent) GetCamundaJobPriority() *int {
	return &startEvent.CamundaJobPriority
}

// GetCamundaInitiator ...
func (startEvent StartEvent) GetCamundaInitiator() *string {
	return &startEvent.CamundaInit
}

/* Elements */

// SetExtensionElements ...
func (startEvent StartEvent) GetExtensionElements() *camunda.ExtensionElements {
	return &startEvent.ExtensionElements[0]
}

// SetConditionalEventDefinition ...
func (startEvent StartEvent) GetConditionalEventDefinition() *ConditionalEventDefinition {
	return &startEvent.ConditionalEventDef[0]
}

// SetMsgEventDefinition ...
func (startEvent StartEvent) GetMsgEventDefinition() *MessageEventDefinition {
	return &startEvent.MsgEventDef[0]
}

// SetTimerEventDefinition ...
func (startEvent StartEvent) GetTimerEventDefinition() *TimerEventDefinition {
	return &startEvent.TimerEventDef[0]
}

// SetOutgoing ...
func (startEvent StartEvent) GetOutgoing(num int) *marker.Outgoing {
	return &startEvent.Outgoing[num]
}
