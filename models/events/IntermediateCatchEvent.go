package events

import (
	"fmt"

	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// IntermediateCatchEventRepository ...
type IntermediateCatchEventRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)

	SetExtensionElements()
	SetIncoming(num int)
	SetOutgoing(num int)
	SetConditionalEventDefinition()
	SetTimerEventDefinition()
	SetMessageEventDefinition()

	GetID(suffix string) *string
	GetName(name string) *string

	GetExtensionElements() *camunda.ExtensionElements
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing
	GetConditionalEventDefinition() *ConditionalEventDefinition
	GetTimerEventDefinition() *TimerEventDefinition
	GetMessageEventDefinition() *MessageEventDefinition
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	ID                         string                       `xml:"id,attr,omitempty" json:"id"`
	Name                       string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements          []camunda.ExtensionElements  `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                   []marker.Incoming            `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                   []marker.Outgoing            `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	ConditionalEventDefinition []ConditionalEventDefinition `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	TimerEventDefinition       []TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	MessageEventDefinition     []MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	ID                         string                        `xml:"id,attr,omitempty" json:"id"`
	Name                       string                        `xml:"name,attr,omitempty" json:"name,omitempty"`
	ExtensionElements          []camunda.TExtensionElements  `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                   []marker.Incoming             `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                   []marker.Outgoing             `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	ConditionalEventDefinition []TConditionalEventDefinition `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	TimerEventDefinition       []TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	MessageEventDefinition     []MessageEventDefinition      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
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
	ice.ExtensionElements = make([]camunda.ExtensionElements, 1)
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
	ice.ConditionalEventDefinition = make([]ConditionalEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (ice *IntermediateCatchEvent) SetTimerEventDefinition() {
	ice.TimerEventDefinition = make([]TimerEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (ice *IntermediateCatchEvent) SetMessageEventDefinition() {
	ice.MessageEventDefinition = make([]MessageEventDefinition, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ice IntermediateCatchEvent) GetID(suffix string) *string {
	return &ice.ID
}

// GetName ...
func (ice IntermediateCatchEvent) GetName(name string) *string {
	return &ice.Name
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (ice IntermediateCatchEvent) GetExtensionElements() *camunda.ExtensionElements {
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
func (ice IntermediateCatchEvent) GetConditionalEventDefinition() *ConditionalEventDefinition {
	return &ice.ConditionalEventDefinition[0]
}

// GetTimerEventDefinition ...
func (ice IntermediateCatchEvent) GetTimerEventDefinition() *TimerEventDefinition {
	return &ice.TimerEventDefinition[0]
}

// GetMessageEventDefinition ...
func (ice IntermediateCatchEvent) GetMessageEventDefinition() *MessageEventDefinition {
	return &ice.MessageEventDefinition[0]
}
