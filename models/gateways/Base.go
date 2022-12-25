package gateways

import "github.com/deemount/gobpmn/models/marker"

type STR_PTR *string

// gateways
const GATEWAY_EXCLUSIVE string = "exclusiveGateway"
const GATEWAY_INCLUSIVE string = "inclusiveGateway"
const GATEWAY_PARALLEL string = "parallelGateway"
const GATEWAY_COMPLEX string = "complexGateway"
const GATEWAY_EVENT_BASED string = "eventBasedGateway"

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

type GatewayBase interface {
	GatewayID
	GatewayName
	GatewayMarker
	GatewayCamundaBase
}
