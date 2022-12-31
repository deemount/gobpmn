package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

func NewInclusiveGateway() InclusiveGatewayRepository {
	return &InclusiveGateway{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusiveGateway *InclusiveGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		inclusiveGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		inclusiveGateway.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (inclusiveGateway *InclusiveGateway) SetName(name string) {
	inclusiveGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (inclusiveGateway *InclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	inclusiveGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (inclusiveGateway *InclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	inclusiveGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (inclusiveGateway *InclusiveGateway) SetCamundaJobPriority(priority int) {
	inclusiveGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

/*** Attributes ***/

// SetDocumentation ...
func (inclusiveGateway *InclusiveGateway) SetDocumentation() {
	inclusiveGateway.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (inclusiveGateway *InclusiveGateway) SetExtensionElements() {
	inclusiveGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (inclusiveGateway *InclusiveGateway) SetIncoming(num int) {
	inclusiveGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (inclusiveGateway *InclusiveGateway) SetOutgoing(num int) {
	inclusiveGateway.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusiveGateway InclusiveGateway) GetID() STR_PTR {
	return &inclusiveGateway.ID

}

// SetName ...
func (inclusiveGateway InclusiveGateway) GetName() STR_PTR {
	return &inclusiveGateway.Name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (inclusiveGateway InclusiveGateway) GetCamundaAsyncBefore() *bool {
	return &inclusiveGateway.CamundaAsyncBefore
}

// SetCamundaAsyncAfter ...
func (inclusiveGateway InclusiveGateway) GetCamundaAsyncAfter() *bool {
	return &inclusiveGateway.CamundaAsyncAfter
}

// SetCamundaJobPriority ...
func (inclusiveGateway InclusiveGateway) GetCamundaJobPriority() *int {
	return &inclusiveGateway.CamundaJobPriority
}

/* Elements */

/*** Attributes ***/

// GetDocumentation ...
func (inclusiveGateway InclusiveGateway) GetDocumentation() *attributes.Documentation {
	return &inclusiveGateway.Documentation[0]
}

// GetExtensionElements ...
func (inclusiveGateway InclusiveGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &inclusiveGateway.ExtensionElements[0]
}

// GetIncoming ...
func (inclusiveGateway InclusiveGateway) GetIncoming(num int) *marker.Incoming {
	return &inclusiveGateway.Incoming[num]
}

// GetOutgoing ...
func (inclusiveGateway InclusiveGateway) GetOutgoing(num int) *marker.Outgoing {
	return &inclusiveGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (inclusiveGateway InclusiveGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", inclusiveGateway.ID, inclusiveGateway.Name)
}
