package models

import "strconv"

// StartEvent ...
type StartEvent struct {
	ID                    string                       `xml:"id,attr"`
	IsInterrupting        bool                         `xml:"isInterrupting,attr,omitempty"`
	CamundaFormKey        string                       `xml:"camunda:formKey,attr,omitempty"`
	CamundaFormRef        string                       `xml:"camunda:formRef,attr,omitempty"`
	CamundaFormRefBind    string                       `xml:"camunda:formRefBinding,attr,omitempty"`
	CamundaFormRefVersion string                       `xml:"camunda:formRefVersion,attr,omitempty"`
	CamundaAsyncBef       bool                         `xml:"camunda:asyncBefore,attr,omitempty"`
	CamundaAsyncAft       bool                         `xml:"camunda:asyncAfter,attr,omitempty"`
	CamundaJobPrio        int                          `xml:"camunda:jobPriority,attr,omitempty"`
	CamundaInit           string                       `xml:"camunda:initiator,attr,omitempty"`
	ExtensionElements     []ExtensionElements          `xml:"bpmn:extensionElements,omitempty"`
	ConditionalEventDef   []ConditionalEventDefinition `xml:"bpmn:conditionalEventDefintion,omitempty"`
	MsgEventDef           []MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty"`
	TimerEventDef         []TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty"`
	Outgoing              []Outgoing                   `xml:"bpmn:outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (stev *StartEvent) SetID(num int64) {
	stev.ID = "StartEvent_" + strconv.FormatInt(num, 16)
}

// SetIsInterrupting ...
func (stev *StartEvent) SetIsInterrupting(isInterrupt bool) {
	stev.IsInterrupting = isInterrupt
}

/** Camunda **/

// SetCamundaFormKey ...
func (stev *StartEvent) SetCamundaFormKey(key string) {
	stev.CamundaFormKey = key
}

// SetCamundaFormRef ...
func (stev *StartEvent) SetCamundaFormRef(ref string) {
	stev.CamundaFormRef = ref
}

// SetCamundaFormRefBinding ...
// possible arguments: latest, deployment, version
// version is dependent to attribute camunda:formRefVersion
func (stev *StartEvent) SetCamundaFormRefBinding(bind string) {
	stev.CamundaFormRefBind = bind
}

// SetCamundaFormRefVersion ...
func (stev *StartEvent) SetCamundaFormRefVersion(version string) {
	stev.CamundaFormRefVersion = version
}

// SetCamundaAsyncBefore ...
func (stev *StartEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	stev.CamundaAsyncBef = asyncBefore
}

// SetCamundaAsyncAfter ...
func (stev *StartEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	stev.CamundaAsyncAft = asyncAfter
}

// SetCamundaJobPriority ...
func (stev *StartEvent) SetCamundaJobPriority(priority int) {
	stev.CamundaJobPrio = priority
}

// SetCamundaInitiator ...
func (stev *StartEvent) SetCamundaInitiator(initiator string) {
	stev.CamundaInit = initiator
}

/*** Make Elements ***/

// SetExtensionElements ...
func (stev *StartEvent) SetExtensionElements() {
	stev.ExtensionElements = make([]ExtensionElements, 1)
}

// SetConditionalEventDefinition ...
func (stev *StartEvent) SetConditionalEventDefinition() {
	stev.ConditionalEventDef = make([]ConditionalEventDefinition, 1)
}

// SetMsgEventDefinition ...
func (stev *StartEvent) SetMsgEventDefinition() {
	stev.MsgEventDef = make([]MessageEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (stev *StartEvent) SetTimerEventDefinition() {
	stev.TimerEventDef = make([]TimerEventDefinition, 1)
}

// SetOutgoing ...
func (stev *StartEvent) SetOutgoing(num int) {
	stev.Outgoing = make([]Outgoing, num)
}
