package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/marker"
)

// StartEventRepository ...
type StartEventRepository interface {
	EventElementsBase
	EventElementsCamundaBase
	EventElementsMarkerOutgoing
	EventElementsCoreElements

	SetIsInterrupting(isInterrupt bool)
	GetIsInterrupting() *bool

	SetCamundaFormKey(key string)
	GetCamundaFormKey() *string
	SetCamundaFormRef(ref string)
	GetCamundaFormRef() *string
	SetCamundaFormRefBinding(bind string)
	GetCamundaFormRefBinding() *string
	SetCamundaFormRefVersion(version string)
	GetCamundaFormRefVersion() *string
	SetCamundaInitiator(initiator string)
	GetCamundaInitiator() *string

	SetConditionalEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	SetTimerEventDefinition()
	GetTimerEventDefinition() *definitions.TimerEventDefinition

	SetMessageEventDefinition()
	GetMessageEventDefinition() *definitions.MessageEventDefinition

	String() string
}

// StartEvent ...
type StartEvent struct {
	compulsion.CompulsionCoreAttributes
	compulsion.CompulsionCamundaCoreAttributes
	compulsion.CompulsionCoreElements
	IsInterrupting         bool                                     `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	CamundaFormKey         string                                   `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaFormRef         string                                   `xml:"camunda:formRef,attr,omitempty" json:"formRef,omitempty"`
	CamundaFormRefBind     string                                   `xml:"camunda:formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	CamundaFormRefVersion  string                                   `xml:"camunda:formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	CamundaInitiator       string                                   `xml:"camunda:initiator,attr,omitempty" json:"init,omitempty"`
	ConditionalEventDef    []definitions.ConditionalEventDefinition `xml:"bpmn:conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MessageEventDefinition []definitions.MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef          []definitions.TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	Outgoing               []marker.Outgoing                        `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	compulsion.CompulsionCoreAttributes
	compulsion.TCompulsionCamundaCoreAttributes
	compulsion.TCompulsionCoreElements
	IsInterrupting         bool                                      `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	FormKey                string                                    `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	FormRef                string                                    `xml:"formRef,attr,omitempty" json:"formRef,omitempty"`
	FormRefBind            string                                    `xml:"formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	FormRefVersion         string                                    `xml:"formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	Initiator              string                                    `xml:"initiator,attr,omitempty" json:"init,omitempty"`
	ConditionalEventDef    []definitions.TConditionalEventDefinition `xml:"conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MessageEventDefinition []definitions.MessageEventDefinition      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef          []definitions.TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	Outgoing               []marker.Outgoing                         `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewStartEvent() StartEventRepository {
	return &StartEvent{}
}

/*
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
	startEvent.CamundaInitiator = initiator
}

/*** Make Elements ***/

// SetDocumentation ...
func (startEvent *StartEvent) SetDocumentation() {
	startEvent.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (startEvent *StartEvent) SetExtensionElements() {
	startEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetConditionalEventDefinition ...
func (startEvent *StartEvent) SetConditionalEventDefinition() {
	startEvent.ConditionalEventDef = make([]definitions.ConditionalEventDefinition, 1)
}

// SetMessagEventDefinition ...
func (startEvent *StartEvent) SetMessageEventDefinition() {
	startEvent.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (startEvent *StartEvent) SetTimerEventDefinition() {
	startEvent.TimerEventDef = make([]definitions.TimerEventDefinition, 1)
}

// SetOutgoing ...
func (startEvent *StartEvent) SetOutgoing(num int) {
	startEvent.Outgoing = make([]marker.Outgoing, num)
}

/*
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
	return &startEvent.CamundaInitiator
}

/* Elements */

// GetDocumentation ...
func (startEvent StartEvent) GetDocumentation() *attributes.Documentation {
	return &startEvent.Documentation[0]
}

// GetExtensionElements ...
func (startEvent StartEvent) GetExtensionElements() *attributes.ExtensionElements {
	return &startEvent.ExtensionElements[0]
}

// GetConditionalEventDefinition ...
func (startEvent StartEvent) GetConditionalEventDefinition() *definitions.ConditionalEventDefinition {
	return &startEvent.ConditionalEventDef[0]
}

// GetMessageEventDefinition ...
func (startEvent StartEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &startEvent.MessageEventDefinition[0]
}

// GetTimerEventDefinition ...
func (startEvent StartEvent) GetTimerEventDefinition() *definitions.TimerEventDefinition {
	return &startEvent.TimerEventDef[0]
}

// GetOutgoing ...
func (startEvent StartEvent) GetOutgoing(num int) *marker.Outgoing {
	return &startEvent.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (startEvent StartEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", startEvent.ID, startEvent.Name)
}
