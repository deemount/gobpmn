package pool

import (
<<<<<<< HEAD
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

=======
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/loop"
)

// CollaborationRepository ...
type CollaborationRepository interface {
	impl.IFBaseID
	attributes.AttributesBaseElements

	SetParticipant(num int)
	GetParticipant(num int) *Participant

	SetMessageFlow(num int)
	GetMessageFlow(num int) *flow.MessageFlow
}

>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
// FlowNodeRefRepository ...
type FlowNodeRefRepository interface {
	impl.IFBaseID
}

// LaneRepository ...
type LaneRepository interface {
	impl.IFBaseID

	SetFlowNodeRef(num int)
	GetFlowNodeRef(num int) *FlowNodeRef
}

// LaneSetRepository ...
type LaneSetRepository interface {
	impl.IFBaseID
	SetLane(num int)
	GetLane(num int) *Lane
}
<<<<<<< HEAD
=======

// ParticipantRepository ...
type ParticipantRepository interface {
	impl.IFBaseID
	impl.IFBaseName

	SetProcessRef(typ string, suffix string)
	GetProcessRef() *string

	SetDocumentation()
	GetDocumentation() *attributes.Documentation

	SetParticipantMultiplicity()
	GetParticipantMultiplicity() *loop.ParticipantMultiplicity
}
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
