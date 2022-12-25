package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/marker"
)

// IntermediateCatchEventRepository ...
type IntermediateCatchEventRepository interface {
	EventElementsBase

	SetIncoming(num int)
	SetOutgoing(num int)
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing

	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements

	SetConditionalEventDefinition()
	SetTimerEventDefinition()
	SetMessageEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	GetTimerEventDefinition() *definitions.TimerEventDefinition
	GetMessageEventDefinition() *definitions.MessageEventDefinition
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	ID                         string                                   `xml:"id,attr,omitempty" json:"id"`
	Name                       string                                   `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements          []attributes.ExtensionElements           `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                   []marker.Incoming                        `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                   []marker.Outgoing                        `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	ConditionalEventDefinition []definitions.ConditionalEventDefinition `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	TimerEventDefinition       []definitions.TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	MessageEventDefinition     []definitions.MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	ID                         string                                    `xml:"id,attr,omitempty" json:"id"`
	Name                       string                                    `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements          []attributes.TExtensionElements           `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                   []marker.Incoming                         `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                   []marker.Outgoing                         `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	ConditionalEventDefinition []definitions.TConditionalEventDefinition `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	TimerEventDefinition       []definitions.TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	MessageEventDefinition     []definitions.MessageEventDefinition      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

func NewIntermediateCatchEvent() IntermediateCatchEventRepository {
	return &IntermediateCatchEvent{}
}

/* Attributes */

/** BPMN **/

// SetID ...
func (ice *IntermediateCatchEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	case "event":
		ice.ID = fmt.Sprintf("Event_%v", suffix)
		break
	case "id":
		ice.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (ice *IntermediateCatchEvent) SetName(name string) {
	ice.Name = name
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (ice *IntermediateCatchEvent) SetExtensionElements() {
	ice.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (ice *IntermediateCatchEvent) SetIncoming(num int) {
	ice.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (ice *IntermediateCatchEvent) SetOutgoing(num int) {
	ice.Outgoing = make([]marker.Outgoing, num)
}

// SetConditionalEventDefinition ...
func (ice *IntermediateCatchEvent) SetConditionalEventDefinition() {
	ice.ConditionalEventDefinition = make([]definitions.ConditionalEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (ice *IntermediateCatchEvent) SetTimerEventDefinition() {
	ice.TimerEventDefinition = make([]definitions.TimerEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (ice *IntermediateCatchEvent) SetMessageEventDefinition() {
	ice.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ice IntermediateCatchEvent) GetID() eventsbase.STR_PTR {
	return &ice.ID
}

// GetName ...
func (ice IntermediateCatchEvent) GetName() eventsbase.STR_PTR {
	return &ice.Name
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (ice IntermediateCatchEvent) GetExtensionElements() *attributes.ExtensionElements {
	return &ice.ExtensionElements[0]
}

// GetIncoming ...
func (ice IntermediateCatchEvent) GetIncoming(num int) *marker.Incoming {
	return &ice.Incoming[num]
}

// GetOutgoing ...
func (ice IntermediateCatchEvent) GetOutgoing(num int) *marker.Outgoing {
	return &ice.Outgoing[num]
}

// GetConditionalEventDefinition ...
func (ice IntermediateCatchEvent) GetConditionalEventDefinition() *definitions.ConditionalEventDefinition {
	return &ice.ConditionalEventDefinition[0]
}

// GetTimerEventDefinition ...
func (ice IntermediateCatchEvent) GetTimerEventDefinition() *definitions.TimerEventDefinition {
	return &ice.TimerEventDefinition[0]
}

// GetMessageEventDefinition ...
func (ice IntermediateCatchEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &ice.MessageEventDefinition[0]
}
