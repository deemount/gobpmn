package elements

import (
	"context"
	"fmt"
	"log"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/events/definitions"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

// NewEndEvent ...
func NewEndEvent() EndEventRepository {
	return &EndEvent{}
}

func (endEvent TEndEvent) Handle(ctx context.Context) error {
	log.Printf("after: %+v\n", endEvent)
	return nil
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (endEvent *EndEvent) SetID(typ string, suffix interface{}) {
	endEvent.ID = SetID(typ, suffix)
}

// SetName ...
func (endEvent *EndEvent) SetName(name string) {
	endEvent.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (endEvent *EndEvent) SetCamundaAsyncBefore(asyncBefore bool) {
	endEvent.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncBefore ...
func (endEvent *EndEvent) SetCamundaAsyncAfter(asyncAfter bool) {
	endEvent.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (endEvent *EndEvent) SetCamundaJobPriority(priority int) {
	endEvent.CamundaJobPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (endEvent *EndEvent) SetDocumentation() {
	endEvent.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (endEvent *EndEvent) SetExtensionElements() {
	endEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (endEvent *EndEvent) SetIncoming(num int) {
	endEvent.Incoming = make([]marker.Incoming, num)
}

/*** Event Definitions ***/

// SetCompensateEventDefinition ...
func (endEvent *EndEvent) SetCompensateEventDefinition() {
	endEvent.CompensateEventDefinition = make([]definitions.CompensateEventDefinition, 1)
}

// SetEscalationEventDefinition ...
func (endEvent *EndEvent) SetEscalationEventDefinition() {
	endEvent.EscalationEventDefinition = make([]definitions.EscalationEventDefinition, 1)
}

// SetMessageEventDefinition ...
func (endEvent *EndEvent) SetMessageEventDefinition() {
	endEvent.MessageEventDefinition = make([]definitions.MessageEventDefinition, 1)
}

// SetErrorEventDefinition ...
func (endEvent *EndEvent) SetErrorEventDefinition() {
	endEvent.ErrorEventDefinition = make([]definitions.ErrorEventDefinition, 1)
}

// SetSignalEventDefinition ...
func (endEvent *EndEvent) SetSignalEventDefinition() {
	endEvent.SignalEventDefinition = make([]definitions.SignalEventDefinition, 1)
}

// SetTerminateEventDefinition ...
func (endEvent *EndEvent) SetTerminateEventDefinition() {
	endEvent.TerminateEventDefinition = make([]definitions.TerminateEventDefinition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (endEvent EndEvent) GetID() impl.STR_PTR {
	return &endEvent.ID
}

// GetName ...
func (endEvent EndEvent) GetName() impl.STR_PTR {
	return &endEvent.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (endEvent EndEvent) GetCamundaAsyncBefore() *bool {
	return &endEvent.CamundaAsyncBefore
}

// GetCamundaAsyncBefore ...
func (endEvent EndEvent) GetCamundaAsyncAfter() *bool {
	return &endEvent.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (endEvent EndEvent) GetCamundaJobPriority() *int {
	return &endEvent.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (endEvent EndEvent) GetDocumentation() *attributes.Documentation {
	return &endEvent.Documentation[0]
}

// GetExtensionElements ...
func (endEvent EndEvent) GetExtensionElements() *attributes.ExtensionElements {
	return &endEvent.ExtensionElements[0]
}

// GetIncoming ...
func (endEvent EndEvent) GetIncoming(num int) *marker.Incoming {
	return &endEvent.Incoming[num]
}

/*** Event Definitions ***/

// GetCompensateEventDefinition ...
func (endEvent EndEvent) GetCompensateEventDefinition() *definitions.CompensateEventDefinition {
	return &endEvent.CompensateEventDefinition[0]
}

// GetEscalationEventDefinition ...
func (endEvent EndEvent) GetEscalationEventDefinition() *definitions.EscalationEventDefinition {
	return &endEvent.EscalationEventDefinition[0]
}

// GetMessageEventDefinition ...
func (endEvent EndEvent) GetMessageEventDefinition() *definitions.MessageEventDefinition {
	return &endEvent.MessageEventDefinition[0]
}

// GetErrorEventDefinition ...
func (endEvent EndEvent) GetErrorEventDefinition() *definitions.ErrorEventDefinition {
	return &endEvent.ErrorEventDefinition[0]
}

// GetSignalEventDefinition ...
func (endEvent EndEvent) GetSignalEventDefinition() *definitions.SignalEventDefinition {
	return &endEvent.SignalEventDefinition[0]
}

// GetTerminateEventDefinition ...
func (endEvent EndEvent) GetTerminateEventDefinition() *definitions.TerminateEventDefinition {
	return &endEvent.TerminateEventDefinition[0]
}

/*
 * Default String
 */

// String ...
func (endEvent EndEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", endEvent.ID, endEvent.Name)
}

// String ...
func (endEvent TEndEvent) String() string {
	return fmt.Sprintf("id=%v, name=%v", endEvent.ID, endEvent.Name)
}
