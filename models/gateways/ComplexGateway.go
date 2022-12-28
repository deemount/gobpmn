package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
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
func (complexgate *ComplexGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		complexgate.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		complexgate.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (complexgate *ComplexGateway) SetName(name string) {
	complexgate.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexgate *ComplexGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	complexgate.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (complexgate *ComplexGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	complexgate.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (complexgate *ComplexGateway) SetCamundaJobPriority(priority int) {
	complexgate.CamundaJobPriority = priority
}

/*** Make Elements ***/

/*** Attributes ***/

// SetDocumentation ...
func (complexgate *ComplexGateway) SetDocumentation() {
	complexgate.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (complexgate *ComplexGateway) SetExtensionElements() {
	complexgate.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (complexgate *ComplexGateway) SetIncoming(num int) {
	complexgate.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (complexgate *ComplexGateway) SetOutgoing(num int) {
	complexgate.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (complexgate ComplexGateway) GetID() STR_PTR {
	return &complexgate.ID
}

// SetName ...
func (complexgate ComplexGateway) GetName() STR_PTR {
	return &complexgate.Name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexgate ComplexGateway) GetCamundaAsyncBefore() *bool {
	return &complexgate.CamundaAsyncBefore
}

// SetCamundaAsyncAfter ...
func (complexgate ComplexGateway) GetCamundaAsyncAfter() *bool {
	return &complexgate.CamundaAsyncAfter
}

// SetCamundaJobPriority ...
func (complexgate ComplexGateway) GetCamundaJobPriority() *int {
	return &complexgate.CamundaJobPriority
}

/* Elements */

/*** Attributes ***/

// GetDocumentation ...
func (complexgate ComplexGateway) GetDocumentation() *attributes.Documentation {
	return &complexgate.Documentation[0]
}

// GetExtensionElements ...
func (complexgate ComplexGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &complexgate.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (complexgate ComplexGateway) GetIncoming(num int) *marker.Incoming {
	return &complexgate.Incoming[num]
}

// GetOutgoing ...
func (complexgate ComplexGateway) GetOutgoing(num int) *marker.Outgoing {
	return &complexgate.Outgoing[num]
}
