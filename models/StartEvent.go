package models

import "strconv"

// StartEventRepository ...
type StartEventRepository interface {
	SetID(suffix string)
	SetName(name string)
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
	ExtensionElements     []ExtensionElements          `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	ConditionalEventDef   []ConditionalEventDefinition `xml:"bpmn:conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MsgEventDef           []MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef         []TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	Outgoing              []Outgoing                   `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	ID                  string                       `xml:"id,attr" json:"id"`
	Name                string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	IsInterrupting      bool                         `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	FormKey             string                       `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	FormRef             string                       `xml:"formRef,attr,omitempty" json:"formRef,omitempty"`
	FormRefBind         string                       `xml:"formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	FormRefVersion      string                       `xml:"formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	AsyncBefore         bool                         `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter          bool                         `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority         int                          `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Init                string                       `xml:"initiator,attr,omitempty" json:"init,omitempty"`
	ExtensionElements   []ExtensionElements          `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	ConditionalEventDef []ConditionalEventDefinition `xml:"conditionalEventDefintion,omitempty" json:"conditionalEventDefinition,omitempty"`
	MsgEventDef         []MessageEventDefinition     `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDef       []TimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	Outgoing            []Outgoing                   `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (startEvent *StartEvent) SetID(num int64) {
	startEvent.ID = "StartEvent_" + strconv.FormatInt(num, 16)
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
	startEvent.ExtensionElements = make([]ExtensionElements, 1)
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
	startEvent.Outgoing = make([]Outgoing, num)
}
