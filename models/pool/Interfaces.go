package pool

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/loop"
	"github.com/deemount/gobpmn/models/marker"
)

// CollaborationRepository ...
type CollaborationRepository interface {
	compulsion.IFBaseID
	attributes.AttributesBaseElements

	SetParticipant(num int)
	GetParticipant(num int) *Participant

	SetMessageFlow(num int)
	GetMessageFlow(num int) *marker.MessageFlow
}

// FlowNodeRefRepository ...
type FlowNodeRefRepository interface {
	compulsion.IFBaseID
}

// LaneRepository ...
type LaneRepository interface {
	compulsion.IFBaseID

	SetFlowNodeRef(num int)
	GetFlowNodeRef(num int) *FlowNodeRef
}

// LaneSetRepository ...
type LaneSetRepository interface {
	compulsion.IFBaseID
	SetLane(num int)
	GetLane(num int) *Lane
}

// ParticipantRepository ...
type ParticipantRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName

	SetProcessRef(typ string, suffix string)
	GetProcessRef() *string

	SetDocumentation()
	GetDocumentation() *attributes.Documentation

	SetParticipantMultiplicity()
	GetParticipantMultiplicity() *loop.ParticipantMultiplicity
}
