package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/marker"
)

// IntermediateCatchEventRepository ...
type IntermediateCatchEventRepository interface {
	EventElementsBase
	EventElementsCamundaBase
	EventElementsCoreElements
	EventElementsMarker
	EventElementsCoreThrowCatchElements

	SetConditionalEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	SetTimerEventDefinition()
	GetTimerEventDefinition() *definitions.TimerEventDefinition
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	compulsion.CompulsionCoreAttributes
	compulsion.CompulsionCamundaCoreAttributes
	compulsion.CompulsionCoreElements
	compulsion.CompulsionCoreIncomingOutgoing
	MessageEventDefinition     []definitions.MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	LinkEventDefinition        []definitions.LinkEventDefinition        `xml:"bpmn:linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	ConditionalEventDefinition []definitions.ConditionalEventDefinition `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	TimerEventDefinition       []definitions.TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	compulsion.CompulsionCoreAttributes
	compulsion.TCompulsionCamundaCoreAttributes
	compulsion.TCompulsionCoreElements
	compulsion.TCompulsionCoreIncomingOutgoing
	ConditionalEventDefinition []definitions.TConditionalEventDefinition `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	LinkEventDefinition        []definitions.LinkEventDefinition         `xml:"linkEventDefinition,omitempty" json:"linkEventDefinition,omitempty"`
	TimerEventDefinition       []definitions.TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	MessageEventDefinition     []definitions.MessageEventDefinition      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
}

// NewIntermediateCatchEvent ...
func NewIntermediateCatchEvent() IntermediateCatchEventRepository {
	return &IntermediateCatchEvent{}
}

/*
 * Default Setters
 */

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

/** Camunda **/

// SetCamundaAsyncBefore ...
func (ice *IntermediateCatchEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	ice.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncBefore ...
func (ice *IntermediateCatchEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	ice.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (ice *IntermediateCatchEvent) SetCamundaJobPriority(priority int) {
	ice.CamundaJobPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (ice *IntermediateCatchEvent) SetDocumentation() {
	ice.Documentation = make([]attributes.Documentation, 1)
}

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

// SetLinkEventDefinition ...
func (ice *IntermediateCatchEvent) SetLinkEventDefinition() {
	ice.LinkEventDefinition = make([]definitions.LinkEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (ice *IntermediateCatchEvent) SetTimerEventDefinition() {
	ice.TimerEventDefinition = make([]definitions.TimerEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (ice *IntermediateCatchEvent) SetMessageEventDefinition() {
	ice.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

/*
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

/** Camunda **/

// GetCamundaAsyncBefore ...
func (ice IntermediateCatchEvent) GetCamundaAsyncBefore() *bool {
	return &ice.CamundaAsyncBefore
}

// GetCamundaAsyncBefore ...
func (ice IntermediateCatchEvent) GetCamundaAsyncAfter() *bool {
	return &ice.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (ice IntermediateCatchEvent) GetCamundaJobPriority() *int {
	return &ice.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (ice IntermediateCatchEvent) GetDocumentation() *attributes.Documentation {
	return &ice.Documentation[0]
}

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

// GetLinkEventDefinition ...
func (ice IntermediateCatchEvent) GetLinkEventDefinition() *definitions.LinkEventDefinition {
	return &ice.LinkEventDefinition[0]
}

// GetTimerEventDefinition ...
func (ice IntermediateCatchEvent) GetTimerEventDefinition() *definitions.TimerEventDefinition {
	return &ice.TimerEventDefinition[0]
}

// GetMessageEventDefinition ...
func (ice IntermediateCatchEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &ice.MessageEventDefinition[0]
}
