package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/marker"
)

// IntermediateThrowEventRepository ...
type IntermediateThrowEventRepository interface {
	EventElementsBase

	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements

	SetIncoming(num int)
	SetOutgoing(num int)
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing

	SetLinkEventDefinition()
	SetEscalationEventDefinition()
	SetMessageEventDefinition()
	GetLinkEventDefinition() *definitions.LinkEventDefinition
	GetEscalationEventDefinition() *definitions.EscalationEventDefinition
	GetMessageEventDefinition() *definitions.MessageEventDefinition
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	ID                        string                                  `xml:"id,attr,omitempty" json:"id"`
	Name                      string                                  `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements         []attributes.ExtensionElements          `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []marker.Incoming                       `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                  []marker.Outgoing                       `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	CompensateEventDefinition []definitions.CompensateEventDefinition `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	LinkEventDefinition       []definitions.LinkEventDefinition       `xml:"bpmn:linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	EscalationEventDefinition []definitions.EscalationEventDefinition `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []definitions.MessageEventDefinition    `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

// TIntermediateThrowEvent ...
type TIntermediateThrowEvent struct {
	ID                        string                                  `xml:"id,attr,omitempty" json:"id"`
	Name                      string                                  `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements         []attributes.TExtensionElements         `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                  []marker.Incoming                       `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                  []marker.Outgoing                       `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	CompensateEventDefinition []definitions.CompensateEventDefinition `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	LinkEventDefinition       []definitions.LinkEventDefinition       `xml:"linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	EscalationEventDefinition []definitions.EscalationEventDefinition `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	MessageEventDefinition    []definitions.MessageEventDefinition    `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
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
	intermediateThrowEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
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
	intermediateThrowEvent.LinkEventDefinition = make([]definitions.LinkEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetEscalationEventDefinition() {
	intermediateThrowEvent.EscalationEventDefinition = make([]definitions.EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetMessageEventDefinition() {
	intermediateThrowEvent.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (intermediateThrowEvent IntermediateThrowEvent) GetID() eventsbase.STR_PTR {
	return &intermediateThrowEvent.ID
}

// GetName ...
func (intermediateThrowEvent IntermediateThrowEvent) GetName() eventsbase.STR_PTR {
	return &intermediateThrowEvent.Name
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (intermediateThrowEvent IntermediateThrowEvent) GetExtensionElements() *attributes.ExtensionElements {
	return &intermediateThrowEvent.ExtensionElements[0]
}

// SetIncoming ...
func (intermediateThrowEvent IntermediateThrowEvent) GetIncoming(num int) *marker.Incoming {
	return &intermediateThrowEvent.Incoming[num]
}

// SetOutgoing ...
func (intermediateThrowEvent IntermediateThrowEvent) GetOutgoing(num int) *marker.Outgoing {
	return &intermediateThrowEvent.Outgoing[num]
}

// SetLinkEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetLinkEventDefinition() *definitions.LinkEventDefinition {
	return &intermediateThrowEvent.LinkEventDefinition[0]
}

// SetEscalationEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetEscalationEventDefinition() *definitions.EscalationEventDefinition {
	return &intermediateThrowEvent.EscalationEventDefinition[0]
}

// SetMessageEventDefinition ...
func (intermediateThrowEvent IntermediateThrowEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &intermediateThrowEvent.MessageEventDefinition[0]
}
