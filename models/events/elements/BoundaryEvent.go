package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/marker"
)

// BoundaryEventRepository ...
type BoundaryEventRepository interface {
	EventElementsBase
	EventsElementsDefinitions
	EventElementsMarkerOutgoing

	SetAttachedToRef(ref string)
	GetAttachedToRef() *string

	// maybe @deprecated @7.17 execution platform
	// Notice: maybe token out of a older modeler version.
	// Needs a checkout
	SetCancelActivity(cancel bool)
	GetCancelActivity() *bool
}

// BoundaryEvent ...
type BoundaryEvent struct {
	ID                         string                                   `xml:"id,attr" json:"id"`
	Name                       string                                   `xml:"name,attr,omitempty" json:"name,omitempty"`
	AttachedToRef              string                                   `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity             bool                                     `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	CancelEventDefinition      []definitions.CancelEventDefinition      `xml:"bpmn:cancelEventDefinition,omitempty" json:"cancelEventDefinition,omitempty"`
	MessageEventDefinition     []definitions.MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       []definitions.TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  []definitions.EscalationEventDefinition  `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition []definitions.ConditionalEventDefinition `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       []definitions.ErrorEventDefinition       `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      []definitions.SignalEventDefinition      `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  []definitions.CompensateEventDefinition  `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	Outgoing                   []marker.Outgoing                        `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TBoundaryEvent ...
type TBoundaryEvent struct {
	ID                         string                                    `xml:"id,attr" json:"id"`
	Name                       string                                    `xml:"name,attr,omitempty" json:"name,omitempty"`
	AttachedToRef              string                                    `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity             bool                                      `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	CancelEventDefinition      []definitions.CancelEventDefinition       `xml:"cancelEventDefinition,omitempty" json:"cancelEventDefinition,omitempty"`
	MessageEventDefinition     []definitions.MessageEventDefinition      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       []definitions.TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  []definitions.EscalationEventDefinition   `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition []definitions.TConditionalEventDefinition `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       []definitions.ErrorEventDefinition        `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      []definitions.SignalEventDefinition       `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  []definitions.CompensateEventDefinition   `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	Outgoing                   []marker.Outgoing                         `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewBoundaryEvent() BoundaryEventRepository {
	return &BoundaryEvent{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (be *BoundaryEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	case "event":
		be.ID = fmt.Sprintf("Event_%v", suffix)
		break
	case "id":
		be.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (be *BoundaryEvent) SetName(name string) {
	be.ID = name
}

// SetAttachedToRef ...
func (be *BoundaryEvent) SetAttachedToRef(ref string) {
	be.AttachedToRef = ref
}

// SetCancelActivity ...
func (be *BoundaryEvent) SetCancelActivity(cancel bool) {
	be.CancelActivity = cancel
}

/*** Make Elements ***/

/** BPMN **/

/*** Event Definition ***/

// SetMessageEventDefinition ...
func (be *BoundaryEvent) SetCancelEventDefinition() {
	be.CancelEventDefinition = make([]definitions.CancelEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (be *BoundaryEvent) SetMessageEventDefinition() {
	be.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (be *BoundaryEvent) SetTimerEventDefinition() {
	be.TimerEventDefinition = make([]definitions.TimerEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (be *BoundaryEvent) SetEscalationEventDefinition() {
	be.EscalationEventDefinition = make([]definitions.EscalationEventDefinition, 1)
}

// SetErrorEventDefinition ...
func (be *BoundaryEvent) SetErrorEventDefinition() {
	be.ErrorEventDefinition = make([]definitions.ErrorEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (be *BoundaryEvent) SetSignalEventDefinition() {
	be.SignalEventDefinition = make([]definitions.SignalEventDefinition, 1)
}

// SetCompensateEventDefinition ...
func (be *BoundaryEvent) SetCompensateEventDefinition() {
	be.CompensateEventDefinition = make([]definitions.CompensateEventDefinition, 1)
}

// SetConditionalEventDefinition ...
func (be *BoundaryEvent) SetConditionalEventDefinition() {
	be.ConditionalEventDefinition = make([]definitions.ConditionalEventDefinition, 1)
}

/*** Marker ***/

// SetOutgoing ...
func (be *BoundaryEvent) SetOutgoing(num int) {
	be.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (be BoundaryEvent) GetID() eventsbase.STR_PTR {
	return &be.ID
}

// GetName ...
func (be BoundaryEvent) GetName() eventsbase.STR_PTR {
	return &be.Name
}

// GetAttachedToRef ...
func (be BoundaryEvent) GetAttachedToRef() *string {
	return &be.AttachedToRef
}

// GetCancelActivity ...
func (be BoundaryEvent) GetCancelActivity() *bool {
	return &be.CancelActivity
}

/* Elements */

/** BPMN **/

/*** Event Definition ***/

// GetMessageEventDefinition ...
func (be BoundaryEvent) GetCancelEventDefinition() *definitions.CancelEventDefinition {
	return &be.CancelEventDefinition[0]
}

// GetMessageEventDefinition ...
func (be BoundaryEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &be.MessageEventDefinition[0]
}

// GetTimerEventDefinition ...
func (be BoundaryEvent) GetTimerEventDefinition() *definitions.TimerEventDefinition {
	return &be.TimerEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (be BoundaryEvent) GetEscalationEventDefinition() *definitions.EscalationEventDefinition {
	return &be.EscalationEventDefinition[0]
}

// GetErrorEventDefinition ...
func (be BoundaryEvent) GetErrorEventDefinition() *definitions.ErrorEventDefinition {
	return &be.ErrorEventDefinition[0]
}

// GetSignalEventDefinition ...
func (be BoundaryEvent) GetSignalEventDefinition() *definitions.SignalEventDefinition {
	return &be.SignalEventDefinition[0]
}

// GetCompensateEventDefinition ...
func (be BoundaryEvent) GetCompensateEventDefinition() *definitions.CompensateEventDefinition {
	return &be.CompensateEventDefinition[0]
}

// GetConditionalEventDefinition ...
func (be BoundaryEvent) GetConditionalEventDefinition() *definitions.ConditionalEventDefinition {
	return &be.ConditionalEventDefinition[0]
}

/*** Marker ***/

// SetOutgoing ...
func (be BoundaryEvent) GetOutgoing(num int) *marker.Outgoing {
	return &be.Outgoing[num]
}