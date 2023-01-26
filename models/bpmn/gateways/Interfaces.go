package gateways

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/camunda"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

type GatewayID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type GatewayName interface {
	SetName(name string)
	GetName() STR_PTR
}

type GatewayBase interface {
	GatewayID
	GatewayName
	marker.MarkerIncomingOutgoing
	camunda.CamundaDefaultAttributes
	attributes.AttributesBaseElements
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
