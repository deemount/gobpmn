package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

func NewExclusiveGateway() ExclusiveGatewayRepository {
	return &InclusiveGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (exclusiveGateway *ExclusiveGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		exclusiveGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		exclusiveGateway.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (exclusiveGateway *ExclusiveGateway) SetName(name string) {
	exclusiveGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	exclusiveGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	exclusiveGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaJobPriority(priority int) {
	exclusiveGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

/*** Attributes ***/

// SetDocumentation ...
func (exclusiveGateway *ExclusiveGateway) SetDocumentation() {
	exclusiveGateway.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (exclusiveGateway *ExclusiveGateway) SetExtensionElements() {
	exclusiveGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (exclusiveGateway *ExclusiveGateway) SetIncoming(num int) {
	exclusiveGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (exclusiveGateway *ExclusiveGateway) SetOutgoing(num int) {
	exclusiveGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (exclusiveGateway ExclusiveGateway) GetID() *string {
	return &exclusiveGateway.ID
}

// GetName ...
func (exclusiveGateway ExclusiveGateway) GetName() *string {
	return &exclusiveGateway.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (exclusiveGateway ExclusiveGateway) GetCamundaAsyncBefore() *bool {
	return &exclusiveGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (exclusiveGateway ExclusiveGateway) GetCamundaAsyncAfter() *bool {
	return &exclusiveGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (exclusiveGateway ExclusiveGateway) GetCamundaJobPriority() *int {
	return &exclusiveGateway.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (exclusiveGateway ExclusiveGateway) GetDocumentation() *attributes.Documentation {
	return &exclusiveGateway.Documentation[0]
}

// GetExtensionElements ...
func (exclusiveGateway ExclusiveGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &exclusiveGateway.ExtensionElements[0]
}

// GetIncoming ...
func (exclusiveGateway ExclusiveGateway) GetIncoming(num int) *marker.Incoming {
	return &exclusiveGateway.Incoming[num]
}

// GetOutgoing ...
func (exclusiveGateway ExclusiveGateway) GetOutgoing(num int) *marker.Outgoing {
	return &exclusiveGateway.Outgoing[num]
}
