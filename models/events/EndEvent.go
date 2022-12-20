package events

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// EndEventRepository ...
type EndEventRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)

	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)

	SetDocumentation()
	SetExtensionElements()
	SetIncoming(num int)
	SetCompensateEventDefinition()
	SetEscalationEventDefinition()
	SetMessageEventDefinition()
	SetErrorEventDefinition()
	SetSignalEventDefinition()
	SetTerminateEventDefinition()

	GetID() *string
	GetName() *string

	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool

	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *camunda.ExtensionElements
	GetIncoming(num int) *marker.Incoming
	GetCompensateEventDefinition() *CompensateEventDefinition
	GetEscalationEventDefinition() *EscalationEventDefinition
	GetMessageEventDefinition() *MessageEventDefinition
	GetErrorEventDefinition() *ErrorEventDefinition
	GetSignalEventDefinition() *SignalEventDefinition
	GetTerminateEventDefinition() *TerminateEventDefinition
}

// EndEvent ...
type EndEvent struct {
	ID                        string                      `xml:"id,attr" json:"id"`
	Name                      string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore        bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter         bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority        int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation             []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements         []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	CompensateEventDefinition []CompensateEventDefinition `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition []EscalationEventDefinition `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []MessageEventDefinition    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      []ErrorEventDefinition      `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     []SignalEventDefinition     `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  []TerminateEventDefinition  `xml:"bpmn:terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
}

// TEndEvent ...
type TEndEvent struct {
	ID                        string                      `xml:"id,attr" json:"id"`
	Name                      string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore               bool                        `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter                bool                        `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority               int                         `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation             []attributes.Documentation  `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements         []camunda.ExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []marker.Incoming           `xml:"incoming,omitempty" json:"incoming,omitempty"`
	CompensateEventDefinition []CompensateEventDefinition `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition []EscalationEventDefinition `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []MessageEventDefinition    `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      []ErrorEventDefinition      `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     []SignalEventDefinition     `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  []TerminateEventDefinition  `xml:"terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
}

func NewEndEvent() EndEventRepository {
	return &EndEvent{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (endEvent *EndEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	case "event":
		endEvent.ID = fmt.Sprintf("Event_%v", suffix)
		break
	case "id":
		endEvent.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (endEvent *EndEvent) SetName(name string) {
	endEvent.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (endEvent *EndEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	endEvent.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncBefore ...
func (endEvent *EndEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	endEvent.CamundaAsyncAfter = asyncAfter
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (endEvent *EndEvent) SetDocumentation() {
	endEvent.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (endEvent *EndEvent) SetExtensionElements() {
	endEvent.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (endEvent *EndEvent) SetIncoming(num int) {
	endEvent.Incoming = make([]marker.Incoming, num)
}

/*** Event Definitions ***/

// SetCompensateEventDefinition ...
func (endEvent *EndEvent) SetCompensateEventDefinition() {
	endEvent.CompensateEventDefinition = make([]CompensateEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (endEvent *EndEvent) SetEscalationEventDefinition() {
	endEvent.EscalationEventDefinition = make([]EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (endEvent *EndEvent) SetMessageEventDefinition() {
	endEvent.MessageEventDefinition = make([]MessageEventDefinition, 1)
}

// SetErrorEventDefinition ...
func (endEvent *EndEvent) SetErrorEventDefinition() {
	endEvent.ErrorEventDefinition = make([]ErrorEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (endEvent *EndEvent) SetSignalEventDefinition() {
	endEvent.SignalEventDefinition = make([]SignalEventDefinition, 1)
}

// SetTerminateEventDefinition ...
func (endEvent *EndEvent) SetTerminateEventDefinition() {
	endEvent.TerminateEventDefinition = make([]TerminateEventDefinition, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (endEvent EndEvent) GetID() *string {
	return &endEvent.ID
}

// GetName ...
func (endEvent EndEvent) GetName() *string {
	return &endEvent.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (endEvent *EndEvent) GetCamundaAsyncBefore() *bool {
	return &endEvent.CamundaAsyncBefore
}

// GetCamundaAsyncBefore ...
func (endEvent *EndEvent) GetCamundaAsyncAfter() *bool {
	return &endEvent.CamundaAsyncAfter
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (endEvent EndEvent) GetDocumentation() *attributes.Documentation {
	return &endEvent.Documentation[0]
}

// GetExtensionElements ...
func (endEvent EndEvent) GetExtensionElements() *camunda.ExtensionElements {
	return &endEvent.ExtensionElements[0]
}

// GetIncoming ...
func (endEvent EndEvent) GetIncoming(num int) *marker.Incoming {
	return &endEvent.Incoming[num]
}

/*** Event Definitions ***/

// GetCompensateEventDefinition ...
func (endEvent EndEvent) GetCompensateEventDefinition() *CompensateEventDefinition {
	return &endEvent.CompensateEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (endEvent EndEvent) GetEscalationEventDefinition() *EscalationEventDefinition {
	return &endEvent.EscalationEventDefinition[0]
}

// GetMessageEventDefinition ...
func (endEvent EndEvent) GetMessageEventDefinition() *MessageEventDefinition {
	return &endEvent.MessageEventDefinition[0]
}

// GetErrorEventDefinition ...
func (endEvent EndEvent) GetErrorEventDefinition() *ErrorEventDefinition {
	return &endEvent.ErrorEventDefinition[0]
}

// GetSignalEventDefinition ...
func (endEvent EndEvent) GetSignalEventDefinition() *SignalEventDefinition {
	return &endEvent.SignalEventDefinition[0]
}

// GetTerminateEventDefinition ...
func (endEvent EndEvent) GetTerminateEventDefinition() *TerminateEventDefinition {
	return &endEvent.TerminateEventDefinition[0]
}
