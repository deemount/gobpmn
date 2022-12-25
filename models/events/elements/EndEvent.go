package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/marker"
)

// EndEventRepository ...
type EndEventRepository interface {
	EventElementsBase
	EventElementsCamundaBase
	EventElementsMarkerIncoming
	eventsbase.EventsSetDefinitionsBase
	eventsbase.EventsSetTerminateBase
	definitions.DefinitionsGetElementsBase
	definitions.DefinitionsGetTerminateBase

	SetDocumentation()
	SetExtensionElements()
	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *attributes.ExtensionElements
}

// EndEvent ...
type EndEvent struct {
	ID                        string                                  `xml:"id,attr" json:"id"`
	Name                      string                                  `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore        bool                                    `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter         bool                                    `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority        int                                     `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation             []attributes.Documentation              `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements         []attributes.ExtensionElements          `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []marker.Incoming                       `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	CompensateEventDefinition []definitions.CompensateEventDefinition `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition []definitions.EscalationEventDefinition `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []definitions.MessageEventDefinition    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      []definitions.ErrorEventDefinition      `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     []definitions.SignalEventDefinition     `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  []definitions.TerminateEventDefinition  `xml:"bpmn:terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
}

// TEndEvent ...
type TEndEvent struct {
	ID                        string                                  `xml:"id,attr" json:"id"`
	Name                      string                                  `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore               bool                                    `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter                bool                                    `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority               int                                     `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation             []attributes.Documentation              `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements         []attributes.TExtensionElements         `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []marker.Incoming                       `xml:"incoming,omitempty" json:"incoming,omitempty"`
	CompensateEventDefinition []definitions.CompensateEventDefinition `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition []definitions.EscalationEventDefinition `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []definitions.MessageEventDefinition    `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      []definitions.ErrorEventDefinition      `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     []definitions.SignalEventDefinition     `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  []definitions.TerminateEventDefinition  `xml:"terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
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

// SetCamundaJobPriority ...
func (endEvent *EndEvent) SetCamundaJobPriority(priority int) {
	endEvent.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (endEvent *EndEvent) SetDocumentation() {
	endEvent.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (endEvent *EndEvent) SetExtensionElements() {
	endEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (endEvent *EndEvent) SetIncoming(num int) {
	endEvent.Incoming = make([]marker.Incoming, num)
}

/*** Event Definitions ***/

// SetCompensateEventDefinition ...
func (endEvent *EndEvent) SetCompensateEventDefinition() {
	endEvent.CompensateEventDefinition = make([]definitions.CompensateEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (endEvent *EndEvent) SetEscalationEventDefinition() {
	endEvent.EscalationEventDefinition = make([]definitions.EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (endEvent *EndEvent) SetMessageEventDefinition() {
	endEvent.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

// SetErrorEventDefinition ...
func (endEvent *EndEvent) SetErrorEventDefinition() {
	endEvent.ErrorEventDefinition = make([]definitions.ErrorEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (endEvent *EndEvent) SetSignalEventDefinition() {
	endEvent.SignalEventDefinition = make([]definitions.SignalEventDefinition, 1)
}

// SetTerminateEventDefinition ...
func (endEvent *EndEvent) SetTerminateEventDefinition() {
	endEvent.TerminateEventDefinition = make([]definitions.TerminateEventDefinition, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (endEvent EndEvent) GetID() eventsbase.STR_PTR {
	return &endEvent.ID
}

// GetName ...
func (endEvent EndEvent) GetName() eventsbase.STR_PTR {
	return &endEvent.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (endEvent EndEvent) GetCamundaAsyncBefore() *bool {
	return &endEvent.CamundaAsyncBefore
}

// GetCamundaAsyncBefore ...
func (endEvent EndEvent) GetCamundaAsyncAfter() *bool {
	return &endEvent.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (endEvent EndEvent) GetCamundaJobPriority() *int {
	return &endEvent.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (endEvent EndEvent) GetDocumentation() *attributes.Documentation {
	return &endEvent.Documentation[0]
}

// GetExtensionElements ...
func (endEvent EndEvent) GetExtensionElements() *attributes.ExtensionElements {
	return &endEvent.ExtensionElements[0]
}

// GetIncoming ...
func (endEvent EndEvent) GetIncoming(num int) *marker.Incoming {
	return &endEvent.Incoming[num]
}

/*** Event Definitions ***/

// GetCompensateEventDefinition ...
func (endEvent EndEvent) GetCompensateEventDefinition() *definitions.CompensateEventDefinition {
	return &endEvent.CompensateEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (endEvent EndEvent) GetEscalationEventDefinition() *definitions.EscalationEventDefinition {
	return &endEvent.EscalationEventDefinition[0]
}

// GetMessageEventDefinition ...
func (endEvent EndEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &endEvent.MessageEventDefinition[0]
}

// GetErrorEventDefinition ...
func (endEvent EndEvent) GetErrorEventDefinition() *definitions.ErrorEventDefinition {
	return &endEvent.ErrorEventDefinition[0]
}

// GetSignalEventDefinition ...
func (endEvent EndEvent) GetSignalEventDefinition() *definitions.SignalEventDefinition {
	return &endEvent.SignalEventDefinition[0]
}

// GetTerminateEventDefinition ...
func (endEvent EndEvent) GetTerminateEventDefinition() *definitions.TerminateEventDefinition {
	return &endEvent.TerminateEventDefinition[0]
}
