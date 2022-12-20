package events

import (
	"fmt"

	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// IntermediateThrowEventRepository ...
type IntermediateThrowEventRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)

	SetExtensionElements()
	SetIncoming(num int)
	SetOutgoing(num int)
	SetLinkEventDefinition()
	SetEscalationEventDefinition()
	SetMessageEventDefinition()

	GetID() *string
	GetName() *string

	GetExtensionElements() *camunda.ExtensionElements
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing
	GetLinkEventDefinition() *LinkEventDefinition
	GetEscalationEventDefinition() *EscalationEventDefinition
	GetMessageEventDefinition() *MessageEventDefinition
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	ID                        string                      `xml:"id,attr,omitempty" json:"id"`
	Name                      string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements         []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                  []marker.Outgoing           `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	CompensateEventDefinition []CompensateEventDefinition `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	LinkEventDefinition       []LinkEventDefinition       `xml:"bpmn:linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	EscalationEventDefinition []EscalationEventDefinition `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []MessageEventDefinition    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

// TIntermediateThrowEvent ...
type TIntermediateThrowEvent struct {
	ID                        string                       `xml:"id,attr,omitempty" json:"id"`
	Name                      string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements         []camunda.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []marker.Incoming            `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                  []marker.Outgoing            `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	CompensateEventDefinition []CompensateEventDefinition  `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	LinkEventDefinition       []LinkEventDefinition        `xml:"linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	EscalationEventDefinition []EscalationEventDefinition  `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []MessageEventDefinition     `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

func NewIntermediateThrowEvent() IntermediateThrowEventRepository {
	return &IntermediateThrowEvent{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	case "event":
		intermediateThrowEvent.ID = fmt.Sprintf("Event_%v", suffix)
		break
	case "id":
		intermediateThrowEvent.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetName(name string) {
	intermediateThrowEvent.Name = name
}

/*** Make Elements +++*/

/** BPMN **/

// SetExtensionElements ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetExtensionElements() {
	intermediateThrowEvent.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetIncoming(num int) {
	intermediateThrowEvent.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetOutgoing(num int) {
	intermediateThrowEvent.Outgoing = make([]marker.Outgoing, num)
}

// SetLinkEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetLinkEventDefinition() {
	intermediateThrowEvent.LinkEventDefinition = make([]LinkEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetEscalationEventDefinition() {
	intermediateThrowEvent.EscalationEventDefinition = make([]EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetMessageEventDefinition() {
	intermediateThrowEvent.MessageEventDefinition = make([]MessageEventDefinition, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (intermediateThrowEvent *IntermediateThrowEvent) GetID() *string {
	return &intermediateThrowEvent.ID
}

// GetName ...
func (intermediateThrowEvent *IntermediateThrowEvent) GetName() *string {
	return &intermediateThrowEvent.Name
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (intermediateThrowEvent *IntermediateThrowEvent) GetExtensionElements() *camunda.ExtensionElements {
	return &intermediateThrowEvent.ExtensionElements[0]
}

// SetIncoming ...
func (intermediateThrowEvent *IntermediateThrowEvent) GetIncoming(num int) *marker.Incoming {
	return &intermediateThrowEvent.Incoming[num]
}

// SetOutgoing ...
func (intermediateThrowEvent *IntermediateThrowEvent) GetOutgoing(num int) *marker.Outgoing {
	return &intermediateThrowEvent.Outgoing[num]
}

// SetLinkEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) GetLinkEventDefinition() *LinkEventDefinition {
	return &intermediateThrowEvent.LinkEventDefinition[0]
}

// SetEscalationEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) GetEscalationEventDefinition() *EscalationEventDefinition {
	return &intermediateThrowEvent.EscalationEventDefinition[0]
}

// SetMessageEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) GetMessageEventDefinition() *MessageEventDefinition {
	return &intermediateThrowEvent.MessageEventDefinition[0]
}
