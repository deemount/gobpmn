package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

func NewParallelGateway() ParallelGatewayRepository {
	return &ParallelGateway{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (parallelGateway *ParallelGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		parallelGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		parallelGateway.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (parallelGateway *ParallelGateway) SetName(name string) {
	parallelGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (parallelGateway *ParallelGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	parallelGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (parallelGateway *ParallelGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	parallelGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (parallelGateway *ParallelGateway) SetCamundaJobPriority(priority int) {
	parallelGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

/*** Attributes ***/

// SetDocumentation ...
func (parallelGateway *ParallelGateway) SetDocumentation() {
	parallelGateway.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (parallelGateway *ParallelGateway) SetExtensionElements() {
	parallelGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (parallelGateway *ParallelGateway) SetIncoming(num int) {
	parallelGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (parallelGateway *ParallelGateway) SetOutgoing(num int) {
	parallelGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (parallelGateway ParallelGateway) GetID() STR_PTR {
	return &parallelGateway.ID
}

// GetName ...
func (parallelGateway ParallelGateway) GetName() STR_PTR {
	return &parallelGateway.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (parallelGateway ParallelGateway) GetCamundaAsyncBefore() *bool {
	return &parallelGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (parallelGateway ParallelGateway) GetCamundaAsyncAfter() *bool {
	return &parallelGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (parallelGateway ParallelGateway) GetCamundaJobPriority() *int {
	return &parallelGateway.CamundaJobPriority
}

/* Elements */

/*** Attributes ***/

// GetDocumentation ...
func (parallelGateway ParallelGateway) GetDocumentation() *attributes.Documentation {
	return &parallelGateway.Documentation[0]
}

// GetExtensionElements ...
func (parallelGateway ParallelGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &parallelGateway.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (parallelGateway ParallelGateway) GetIncoming(num int) *marker.Incoming {
	return &parallelGateway.Incoming[num]
}

// GetOutgoing ...
func (parallelGateway ParallelGateway) GetOutgoing(num int) *marker.Outgoing {
	return &parallelGateway.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (parallelGateway ParallelGateway) String() string {
	return fmt.Sprintf("id=%v, name=%v", parallelGateway.ID, parallelGateway.Name)
}
