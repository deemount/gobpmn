package gateways

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/camunda"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

type GatewayID interface {
	SetID(typ string, suffix interface{})
	GetID() impl.STR_PTR
}

type GatewayName interface {
	SetName(name string)
	GetName() impl.STR_PTR
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
	GetExclusiveGateway(num int) EXCLUSIVE_GATEWAY_PTR
	SetInclusiveGateway(num int)
	GetInclusiveGateway(num int) INCLUSIVE_GATEWAY_PTR
	SetParallelGateway(num int)
	GetParallelGateway(num int) PARALLEL_GATEWAY_PTR
	SetComplexGateway(num int)
	GetComplexGateway(num int) COMPLEX_GATEWAY_PTR
	SetEventBasedGateway(num int)
	GetEventBasedGateway(num int) EVENT_BASED_GATEWAYS_PTR
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
