package events

import (
	"fmt"

	"github.com/deemount/gobpmn/models/marker"
)

// BoundaryEventRepository ...
type BoundaryEventRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)
	SetAttachedToRef(ref string)
	SetCancelActivity(cancel bool)

	SetCancelEventDefinition()
	SetMessageEventDefinition()
	SetTimerEventDefinition()
	SetEscalationEventDefinition()
	SetConditionalEventDefinition()
	SetErrorEventDefinition()
	SetSignalEventDefinition()
	SetCompensateEventDefinition()
	SetOutgoing(num int)

	GetID() *string
	GetName() *string
	GetAttachedToRef() *string
	GetCancelActivity() *bool

	GetCancelEventDefinition() *CancelEventDefinition
	GetMessageEventDefinition() *MessageEventDefinition
	GetTimerEventDefinition() *TimerEventDefinition
	GetEscalationEventDefinition() *EscalationEventDefinition
	GetErrorEventDefinition() *ErrorEventDefinition
	GetSignalEventDefinition() *SignalEventDefinition
	GetCompensateEventDefinition() *CompensateEventDefinition
	GetOutgoing(num int) *marker.Outgoing
}

// BoundaryEvent ...
type BoundaryEvent struct {
	ID                         string                       `xml:"id,attr" json:"id"`
	Name                       string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	AttachedToRef              string                       `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity             bool                         `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"`
	CancelEventDefinition      []CancelEventDefinition      `xml:"bpmn:cancelEventDefinition,omitempty" json:"cancelEventDefinition,omitempty"`
	MessageEventDefinition     []MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       []TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  []EscalationEventDefinition  `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition []ConditionalEventDefinition `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       []ErrorEventDefinition       `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      []SignalEventDefinition      `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  []CompensateEventDefinition  `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	Outgoing                   []marker.Outgoing            `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TBoundaryEvent ...
type TBoundaryEvent struct {
	ID                         string                        `xml:"id,attr" json:"id"`
	Name                       string                        `xml:"name,attr,omitempty" json:"name,omitempty"`
	AttachedToRef              string                        `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity             bool                          `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"`
	CancelEventDefinition      []CancelEventDefinition       `xml:"cancelEventDefinition,omitempty" json:"cancelEventDefinition,omitempty"`
	MessageEventDefinition     []MessageEventDefinition      `xml:"messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       []TTimerEventDefinition       `xml:"timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  []EscalationEventDefinition   `xml:"escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition []TConditionalEventDefinition `xml:"conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       []ErrorEventDefinition        `xml:"errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      []SignalEventDefinition       `xml:"signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  []CompensateEventDefinition   `xml:"compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	Outgoing                   []marker.Outgoing             `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
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
	be.CancelEventDefinition = make([]CancelEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (be *BoundaryEvent) SetMessageEventDefinition() {
	be.MessageEventDefinition = make([]MessageEventDefinition, 1)
}

// SetTimerEventDefinition ...
func (be *BoundaryEvent) SetTimerEventDefinition() {
	be.TimerEventDefinition = make([]TimerEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (be *BoundaryEvent) SetEscalationEventDefinition() {
	be.EscalationEventDefinition = make([]EscalationEventDefinition, 1)
}

// SetConditionalEventDefinition ...
func (be *BoundaryEvent) SetConditionalEventDefinition() {
	be.ConditionalEventDefinition = make([]ConditionalEventDefinition, 1)
}

// SetErrorEventDefinition ...
func (be *BoundaryEvent) SetErrorEventDefinition() {
	be.ErrorEventDefinition = make([]ErrorEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (be *BoundaryEvent) SetSignalEventDefinition() {
	be.SignalEventDefinition = make([]SignalEventDefinition, 1)
}

// SetCompensateEventDefinition ...
func (be *BoundaryEvent) SetCompensateEventDefinition() {
	be.CompensateEventDefinition = make([]CompensateEventDefinition, 1)
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
func (be BoundaryEvent) GetID() *string {
	return &be.ID
}

// GetName ...
func (be BoundaryEvent) GetName() *string {
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
func (be BoundaryEvent) GetCancelEventDefinition() *CancelEventDefinition {
	return &be.CancelEventDefinition[0]
}

// GetMessageEventDefinition ...
func (be BoundaryEvent) GetMessageEventDefinition() *MessageEventDefinition {
	return &be.MessageEventDefinition[0]
}

// GetTimerEventDefinition ...
func (be BoundaryEvent) GetTimerEventDefinition() *TimerEventDefinition {
	return &be.TimerEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (be BoundaryEvent) GetEscalationEventDefinition() *EscalationEventDefinition {
	return &be.EscalationEventDefinition[0]
}

// GetConditionalEventDefinition ...
func (be BoundaryEvent) GetConditionalEventDefinition() *ConditionalEventDefinition {
	return &be.ConditionalEventDefinition[0]
}

// GetErrorEventDefinition ...
func (be BoundaryEvent) GetErrorEventDefinition() *ErrorEventDefinition {
	return &be.ErrorEventDefinition[0]
}

// GetSignalEventDefinition ...
func (be BoundaryEvent) GetSignalEventDefinition() *SignalEventDefinition {
	return &be.SignalEventDefinition[0]
}

// GetCompensateEventDefinition ...
func (be BoundaryEvent) GetCompensateEventDefinition() *CompensateEventDefinition {
	return &be.CompensateEventDefinition[0]
}

/*** Marker ***/

// SetOutgoing ...
func (be BoundaryEvent) GetOutgoing(num int) *marker.Outgoing {
	return &be.Outgoing[num]
}
