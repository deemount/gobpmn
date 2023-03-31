package collaboration

import (
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
	GetParticipant(num int) PARTICIPANT_PTR

	SetMessageFlow(num int)
	GetMessageFlow(num int) *flow.MessageFlow
}

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
