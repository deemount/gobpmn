package pool

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/loop"
	"github.com/deemount/gobpmn/models/marker"
)

type PoolID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type PoolName interface {
	SetName(name string)
	GetName() STR_PTR
}

type PoolBaseDocumentation interface {
	SetDocumentation()
	GetDocumentation() *attributes.Documentation
}

type PoolBaseExtensionElements interface {
	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

type PoolBaseCoreElements interface {
	PoolBaseDocumentation
	PoolBaseExtensionElements
}

type PoolBase interface {
	PoolID
	PoolName
}

// CollaborationRepository ...
type CollaborationRepository interface {
	PoolID
	PoolBaseCoreElements

	SetParticipant(num int)
	GetParticipant(num int) *Participant

	SetMessageFlow(num int)
	GetMessageFlow(num int) *marker.MessageFlow
}

// FlowNodeRefRepository ...
type FlowNodeRefRepository interface {
	PoolID
}

// LaneRepository ...
type LaneRepository interface {
	PoolID

	SetFlowNodeRef(num int)
	GetFlowNodeRef(num int) *FlowNodeRef
}

// LaneSetRepository ...
type LaneSetRepository interface {
	PoolID
	SetLane(num int)
	GetLane(num int) *Lane
}

// ParticipantRepository ...
type ParticipantRepository interface {
	PoolBase

	SetProcessRef(typ string, suffix string)
	GetProcessRef() *string

	SetDocumentation()
	GetDocumentation() *attributes.Documentation

	SetParticipantMultiplicity()
	GetParticipantMultiplicity() *loop.ParticipantMultiplicity
}
