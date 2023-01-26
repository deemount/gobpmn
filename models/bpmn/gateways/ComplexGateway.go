package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

func NewComplexGateway() ComplexGatewayRepository {
	return &ComplexGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (complexGateway *ComplexGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		complexGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		complexGateway.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (complexGateway *ComplexGateway) SetName(name string) {
	complexGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexGateway *ComplexGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	complexGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (complexGateway *ComplexGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	complexGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (complexGateway *ComplexGateway) SetCamundaJobPriority(priority int) {
	complexGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

/*** Attributes ***/

// SetDocumentation ...
func (complexGateway *ComplexGateway) SetDocumentation() {
	complexGateway.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (complexGateway *ComplexGateway) SetExtensionElements() {
	complexGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (complexGateway *ComplexGateway) SetIncoming(num int) {
	complexGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (complexGateway *ComplexGateway) SetOutgoing(num int) {
	complexGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (complexGateway ComplexGateway) GetID() STR_PTR {
	return &complexGateway.ID
}

// SetName ...
func (complexGateway ComplexGateway) GetName() STR_PTR {
	return &complexGateway.Name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexGateway ComplexGateway) GetCamundaAsyncBefore() *bool {
	return &complexGateway.CamundaAsyncBefore
}

// SetCamundaAsyncAfter ...
func (complexGateway ComplexGateway) GetCamundaAsyncAfter() *bool {
	return &complexGateway.CamundaAsyncAfter
}

// SetCamundaJobPriority ...
func (complexGateway ComplexGateway) GetCamundaJobPriority() *int {
	return &complexGateway.CamundaJobPriority
}

/* Elements */

/*** Attributes ***/

// GetDocumentation ...
func (complexGateway ComplexGateway) GetDocumentation() *attributes.Documentation {
	return &complexGateway.Documentation[0]
}

// GetExtensionElements ...
func (complexGateway ComplexGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &complexGateway.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (complexGateway ComplexGateway) GetIncoming(num int) *marker.Incoming {
	return &complexGateway.Incoming[num]
}

// GetOutgoing ...
func (complexGateway ComplexGateway) GetOutgoing(num int) *marker.Outgoing {
	return &complexGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (complexGateway ComplexGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", complexGateway.ID, complexGateway.Name)
}
