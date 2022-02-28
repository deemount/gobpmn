package models

import "fmt"

// BoundaryEvent ...
type BoundaryEvent struct {
	ID                         string                       `xml:"id,attr" json:"id"`
	AttachedToRef              string                       `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity             bool                         `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"`
	MessageEventDefinition     []MessageEventDefinition     `xml:"bpmn:messageEventDefinition,omitempty" json:"messageEventDefinition,omitempty"`
	TimerEventDefinition       []TimerEventDefinition       `xml:"bpmn:timerEventDefinition,omitempty" json:"timerEventDefinition,omitempty"`
	EscalationEventDefinition  []EscalationEventDefinition  `xml:"bpmn:escalationEventDefinition,omitempty" json:"escalationEventDefinition,omitempty"`
	ConditionalEventDefinition []ConditionalEventDefinition `xml:"bpmn:conditionalEventDefinition,omitempty" json:"conditionalEventDefinition,omitempty"`
	ErrorEventDefinition       []ErrorEventDefinition       `xml:"bpmn:errorEventDefinition,omitempty" json:"errorEventDefinition,omitempty"`
	SignalEventDefinition      []SignalEventDefinition      `xml:"bpmn:signalEventDefinition,omitempty" json:"signalEventDefinition,omitempty"`
	CompensateEventDefinition  []CompensateEventDefinition  `xml:"bpmn:compensateEventDefinition,omitempty" json:"compensateEventDefinition,omitempty"`
	Outgoing                   []Outgoing                   `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (be *BoundaryEvent) SetID(suffix string) {
	be.ID = fmt.Sprintf("Event_%s", suffix)
}

// SetAttachedToRef ...
func (be *BoundaryEvent) SetAttachedToRef(ref string) {
	be.AttachedToRef = ref
}

// SetCancelActivity ...
func (be *BoundaryEvent) SetCancelActivity(cancel bool) {
	be.CancelActivity = cancel
}

/* Elements */

/** BPMN **/

/*** Event Definition ***/

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
func (be *BoundaryEvent) SetOutgoing() {
	be.Outgoing = make([]Outgoing, 1)
}
