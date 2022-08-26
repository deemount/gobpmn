package models

import "fmt"

// EndEvent ...
type EndEvent struct {
	ID                        string                      `xml:"id,attr" json:"id"`
	Name                      string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore        bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter         bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority        int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation             []Documentation             `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements         []ExtensionElements         `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []Incoming                  `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
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
	Documentation             []Documentation             `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements         []ExtensionElements         `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []Incoming                  `xml:"incoming,omitempty" json:"incoming,omitempty"`
	CompensateEventDefinition []CompensateEventDefinition `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	EscalationEventDefinition []EscalationEventDefinition `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []MessageEventDefinition    `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	ErrorEventDefinition      []ErrorEventDefinition      `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition     []SignalEventDefinition     `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	TerminateEventDefinition  []TerminateEventDefinition  `xml:"terminateEventDefinition,omitempty" json:"terminateEventDefinition,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (endEvent *EndEvent) SetID(typ string, suffix string) {
	switch typ {
	case "event":
		endEvent.ID = fmt.Sprintf("Event_%s", suffix)
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
	endEvent.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (endEvent *EndEvent) SetExtensionElements() {
	endEvent.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (endEvent *EndEvent) SetIncoming(num int) {
	endEvent.Incoming = make([]Incoming, num)
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
