package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/marker"
)

// StartEventRepository ...
type StartEventRepository interface {
	EventElementsBase
	EventElementsCamundaBase
	EventElementsMarkerOutgoing

	SetIsInterrupting(isInterrupt bool)
	GetIsInterrupting() *bool

	SetCamundaFormKey(key string)
	SetCamundaFormRef(ref string)
	SetCamundaFormRefBinding(bind string)
	SetCamundaFormRefVersion(version string)
	GetCamundaFormKey() *string
	GetCamundaFormRef() *string
	GetCamundaFormRefBinding() *string
	GetCamundaFormRefVersion() *string
	SetCamundaInitiator(initiator string)
	GetCamundaInitiator() *string

	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements

	SetConditionalEventDefinition()
	SetMsgEventDefinition()
	SetTimerEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	GetMsgEventDefinition() *definitions.MessageEventDefinition
	GetTimerEventDefinition() *definitions.TimerEventDefinition

	String() string
}

// StartEvent ...
type StartEvent struct {
	ID                    string                                   `xml:"id,attr" json:"id"`
	Name                  string                                   `xml:"name,attr,omitempty" json:"name,omitempty"`
	IsInterrupting        bool                                     `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	CamundaFormKey        string                                   `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaFormRef        string                                   `xml:"camunda:formRef,attr,omitempty" json:"formRef,omitempty"`
	CamundaFormRefBind    string                                   `xml:"camunda:formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	CamundaFormRefVersion string                                   `xml:"camunda:formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	CamundaAsyncBefore    bool                                     `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter     bool                                     `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority    int                                      `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	CamundaInit           string                                   `xml:"camunda:initiator,attr,omitempty" json:"init,omitempty"`
	ExtensionElements     []attributes.ExtensionElements           `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	ConditionalEventDef   []definitions.ConditionalEventDefinition `xml:"bpmn:conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MsgEventDef           []definitions.MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef         []definitions.TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	Outgoing              []marker.Outgoing                        `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	ID                  string                                    `xml:"id,attr" json:"id"`
	Name                string                                    `xml:"name,attr,omitempty" json:"name,omitempty"`
	IsInterrupting      bool                                      `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	FormKey             string                                    `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	FormRef             string                                    `xml:"formRef,attr,omitempty" json:"formRef,omitempty"`
	FormRefBind         string                                    `xml:"formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	FormRefVersion      string                                    `xml:"formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	AsyncBefore         bool                                      `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter          bool                                      `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority         int                                       `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Init                string                                    `xml:"initiator,attr,omitempty" json:"init,omitempty"`
	ExtensionElements   []attributes.TExtensionElements           `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	ConditionalEventDef []definitions.TConditionalEventDefinition `xml:"conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MsgEventDef         []definitions.MessageEventDefinition      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef       []definitions.TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	Outgoing            []marker.Outgoing                         `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
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
	startEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetConditionalEventDefinition ...
func (startEvent *StartEvent) SetConditionalEventDefinition() {
	startEvent.ConditionalEventDef = make([]definitions.ConditionalEventDefinition, 1)
}

// SetMsgEventDefinition ...
func (startEvent *StartEvent) SetMsgEventDefinition() {
	startEvent.MsgEventDef = make([]definitions.MessageEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (startEvent *StartEvent) SetTimerEventDefinition() {
	startEvent.TimerEventDef = make([]definitions.TimerEventDefinition, 1)
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
func (startEvent StartEvent) GetID() eventsbase.STR_PTR {
	return &startEvent.ID
}

// GetName ...
func (startEvent StartEvent) GetName() eventsbase.STR_PTR {
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
func (startEvent StartEvent) GetExtensionElements() *attributes.ExtensionElements {
	return &startEvent.ExtensionElements[0]
}

// SetConditionalEventDefinition ...
func (startEvent StartEvent) GetConditionalEventDefinition() *definitions.ConditionalEventDefinition {
	return &startEvent.ConditionalEventDef[0]
}

// SetMsgEventDefinition ...
func (startEvent StartEvent) GetMsgEventDefinition() *definitions.MessageEventDefinition {
	return &startEvent.MsgEventDef[0]
}

// SetTimerEventDefinition ...
func (startEvent StartEvent) GetTimerEventDefinition() *definitions.TimerEventDefinition {
	return &startEvent.TimerEventDef[0]
}

// SetOutgoing ...
func (startEvent StartEvent) GetOutgoing(num int) *marker.Outgoing {
	return &startEvent.Outgoing[num]
}
