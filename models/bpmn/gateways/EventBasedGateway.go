package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

func NewEventBasedGateway() EventBasedGatewayRepository {
	return &InclusiveGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (eventBasedGateway *EventBasedGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		eventBasedGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		eventBasedGateway.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (eventBasedGateway *EventBasedGateway) SetName(name string) {
	eventBasedGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (eventBasedGateway *EventBasedGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	eventBasedGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (eventBasedGateway *EventBasedGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	eventBasedGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (eventBasedGateway *EventBasedGateway) SetCamundaJobPriority(priority int) {
	eventBasedGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

/*** Attributes ***/

// SetDocumentation ...
func (eventBasedGateway *EventBasedGateway) SetDocumentation() {
	eventBasedGateway.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (eventBasedGateway *EventBasedGateway) SetExtensionElements() {
	eventBasedGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (eventBasedGateway *EventBasedGateway) SetIncoming(num int) {
	eventBasedGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (eventBasedGateway *EventBasedGateway) SetOutgoing(num int) {
	eventBasedGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (eventBasedGateway EventBasedGateway) GetID() STR_PTR {
	return &eventBasedGateway.ID
}

// GetName ...
func (eventBasedGateway EventBasedGateway) GetName() STR_PTR {
	return &eventBasedGateway.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (eventBasedGateway EventBasedGateway) GetCamundaAsyncBefore() *bool {
	return &eventBasedGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (eventBasedGateway EventBasedGateway) GetCamundaAsyncAfter() *bool {
	return &eventBasedGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (eventBasedGateway EventBasedGateway) GetCamundaJobPriority() *int {
	return &eventBasedGateway.CamundaJobPriority
}

/* Elements */

/*** Attributes ***/

// GetDocumentation ...
func (eventBasedGateway *EventBasedGateway) GetDocumentation() *attributes.Documentation {
	return &eventBasedGateway.Documentation[0]
}

// GetExtensionElements ...
func (eventBasedGateway EventBasedGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &eventBasedGateway.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (eventBasedGateway EventBasedGateway) GetIncoming(num int) *marker.Incoming {
	return &eventBasedGateway.Incoming[num]
}

// GetOutgoing ...
func (eventBasedGateway EventBasedGateway) GetOutgoing(num int) *marker.Outgoing {
	return &eventBasedGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (eventBasedGateway EventBasedGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", eventBasedGateway.ID, eventBasedGateway.Name)
}
