package elements

import (
	"context"
	"fmt"
	"log"

	"github.com/deemount/gobpmn/models/bpmn/events/definitions"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

// NewBoundaryEvent ...
func NewBoundaryEvent() BoundaryEventRepository {
	return &BoundaryEvent{}
}

func (be TBoundaryEvent) Handle(ctx context.Context) error {
	log.Printf("after: %+v\n", be)
	return nil
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (be *BoundaryEvent) SetID(typ string, suffix interface{}) {
	be.ID = SetID(typ, suffix)
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

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (be BoundaryEvent) GetID() impl.STR_PTR {
	return &be.ID
}

// GetName ...
func (be BoundaryEvent) GetName() impl.STR_PTR {
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

/*
 * Default String
 */

// String ...
func (be BoundaryEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", be.ID, be.Name)
}

// String ...
func (be TBoundaryEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", be.ID, be.Name)
}
