package models

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

/* Attributes */

/** BPMN **/

// SetID ...
func (ee *EndEvent) SetID(suffix string) {
	ee.ID = "Event_" + suffix
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (ee *EndEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	ee.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncBefore ...
func (ee *EndEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	ee.CamundaAsyncAfter = asyncAfter
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (ee *EndEvent) SetDocumentation() {
	ee.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (ee *EndEvent) SetExtensionElements() {
	ee.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (ee *EndEvent) SetIncoming(num int) {
	ee.Incoming = make([]Incoming, num)
}

/*** Event Definitions ***/

// SetCompensateEventDefinition ...
func (ee *EndEvent) SetCompensateEventDefinition() {
	ee.CompensateEventDefinition = make([]CompensateEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (ee *EndEvent) SetEscalationEventDefinition() {
	ee.EscalationEventDefinition = make([]EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (ee *EndEvent) SetMessageEventDefinition() {
	ee.MessageEventDefinition = make([]MessageEventDefinition, 1)
}

// SetErrorEventDefinition ...
func (ee *EndEvent) SetErrorEventDefinition() {
	ee.ErrorEventDefinition = make([]ErrorEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (ee *EndEvent) SetSignalEventDefinition() {
	ee.SignalEventDefinition = make([]SignalEventDefinition, 1)
}

// SetTerminateEventDefinition ...
func (ee *EndEvent) SetTerminateEventDefinition() {
	ee.TerminateEventDefinition = make([]TerminateEventDefinition, 1)
}
