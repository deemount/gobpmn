package events

import (
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/elements"
)

// boundary events constant
const BOUNDARY_TIMER string = "boundaryTimer"
const BOUNDARY_MESSAGE string = "boundaryMessage"
const BOUNDARY_SIGNAL string = "boundarySignal"
const BOUNDARY_COMPENSATION string = "compensationBoundaryCatch"
const BOUNDARY_ERROR string = "boundaryError"
const BOUNDARY_ESCALATION string = "boundaryEscalation"
const BOUNDARY_CANCEL string = "cancelBoundaryCatch"
const BOUNDARY_CONDITIONAL string = "boundaryConditional"

// start events constants
const START_EVENT string = "startEvent"
const START_EVENT_TIMER string = "startTimerEvent"
const START_EVENT_MESSAGE string = "messageStartEvent"
const START_EVENT_SIGNAL string = "signalStartEvent"
const START_EVENT_ESCALATION string = "escalationStartEvent"
const START_EVENT_COMPENSATION string = "compensationStartEvent"
const START_EVENT_ERROR string = "errorStartEvent"
const START_EVENT_CONDITIONAL string = "conditionalStartEvent"

// intermediate catch events
const INTERMEDIATE_EVENT_CATCH string = "intermediateCatchEvent"
const INTERMEDIATE_EVENT_MESSAGE string = "intermediateMessageCatch"
const INTERMEDIATE_EVENT_TIMER string = "intermediateTimer"
const INTERMEDIATE_EVENT_LINK string = "intermediateLinkCatch"
const INTERMEDIATE_EVENT_SIGNAL string = "intermediateSignalCatch"
const INTERMEDIATE_EVENT_CONDITIONAL string = "intermediateConditional"

// intermediate throw events
const INTERMEDIATE_EVENT_THROW string = "intermediateThrowEvent"
const INTERMEDIATE_EVENT_SIGNAL_THROW string = "intermediateSignalThrow"
const INTERMEDIATE_EVENT_COMPENSATION_THROW string = "intermediateCompensationThrowEvent"
const INTERMEDIATE_EVENT_MESSAGE_THROW string = "intermediateMessageThrowEvent"
const INTERMEDIATE_EVENT_NONE_THROW string = "intermediateNoneThrowEvent"
const INTERMEDIATE_EVENT_ESCALATION_THROW string = "intermediateEscalationThrowEvent"

// typed interfaces
// @definitions
type Type_Cancel_Event_Definition definitions.CancelEventDefinitionRepository
type Type_Compensate_Event_Definition definitions.CompensateEventDefinitionRepository
type Type_Conditional_Event_Definition definitions.ConditionalEventDefinitionRepository
type Type_Error_Event_Definition definitions.ErrorEventDefinitionRepository
type Type_Escalation_Event_Definition definitions.EscalationEventDefinitionRepository
type Type_Link_Event_Definition definitions.LinkEventDefinitionRepository
type Type_Signal_Event_Definition definitions.SignalEventDefinitionRepository
type Type_Timer_Event_Definition definitions.TimerEventDefinitionRepository
type Type_Terminate_Event_Definition definitions.TerminateEventDefinitionRepository

// @elements
type Type_Start_Event elements.StartEventRepository
type Type_End_Event elements.EndEventRepository
type Type_Boundary_Event elements.BoundaryEventRepository
type Type_Intermediate_Catch_Event elements.IntermediateCatchEventRepository
type Type_Intermediate_Throw_Event elements.IntermediateThrowEventRepository

// EventsRepository ...
type EventsRepository interface {
	elements.StartEventRepository
	elements.EndEventRepository
	elements.BoundaryEventRepository
	elements.IntermediateCatchEventRepository
	elements.IntermediateThrowEventRepository
	definitions.CancelEventDefinitionRepository
	definitions.CompensateEventDefinitionRepository
	definitions.ConditionalEventDefinitionRepository
	definitions.ErrorEventDefinitionRepository
	definitions.EscalationEventDefinitionRepository
	definitions.LinkEventDefinitionRepository
	definitions.SignalEventDefinitionRepository
	definitions.TimerEventDefinitionRepository
	definitions.TerminateEventDefinitionRepository
}

type Events struct {
	EventsRepository
}

func NewEvents() EventsRepository {
	return &Events{}
}

func (events *Events) SetID(typ string, suffix interface{}) {}
