package gateways

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

type GatewayID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type GatewayName interface {
	SetName(name string)
	GetName() STR_PTR
}

type GatewayMarker interface {
	SetIncoming(num int)
	GetIncoming(num int) *marker.Incoming
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
}

type GatewayCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
}

type GatewayBaseCoreElements interface {
	SetDocumentation()
	GetDocumentation() *attributes.Documentation
	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

type GatewayBase interface {
	GatewayID
	GatewayName
	GatewayMarker
	GatewayCamundaBase
	GatewayBaseCoreElements
}

type GatewaysElementsRepository interface {
	SetExclusiveGateway(num int)
	GetExclusiveGateway(num int) *ExclusiveGateway
	SetInclusiveGateway(num int)
	GetInclusiveGateway(num int) *InclusiveGateway
	SetParallelGateway(num int)
	GetParallelGateway(num int) *ParallelGateway
	SetComplexGateway(num int)
	GetComplexGateway(num int) *ComplexGateway
	SetEventBasedGateway(num int)
	GetEventBasedGateway(num int) *EventBasedGateway
}

// ComplexGatewayRepository ...
type ComplexGatewayRepository interface {
	GatewayBase
}

// EventBasedGatewayRepository ...
type EventBasedGatewayRepository interface {
	GatewayBase
}

// ExclusiveGatewayRepository ...
type ExclusiveGatewayRepository interface {
	GatewayBase
}

// InclusiveGatewayRepository ...
type InclusiveGatewayRepository interface {
	GatewayBase
}

// ParallelGatewayRepository ...
type ParallelGatewayRepository interface {
	GatewayBase
}
