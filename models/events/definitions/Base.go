package definitions

import "github.com/deemount/gobpmn/models/events/eventsbase"

type DefinitionsID interface{ eventsbase.EventsID }
type DefinitionsName interface{ eventsbase.EventsName }

type DefinitionsGetTerminateBase interface {
	// @EndEvent only
	GetTerminateEventDefinition() *TerminateEventDefinition
}

type DefinitionsGetElementsBase interface {
	GetMessageEventDefinition() *MessageEventDefinition
	GetEscalationEventDefinition() *EscalationEventDefinition
	GetErrorEventDefinition() *ErrorEventDefinition
	GetSignalEventDefinition() *SignalEventDefinition
	GetCompensateEventDefinition() *CompensateEventDefinition
}

type DefinitionsGetElements interface {
	DefinitionsGetElementsBase
	// @BoundaryEvent
	GetCancelEventDefinition() *CancelEventDefinition
	GetTimerEventDefinition() *TimerEventDefinition
	GetConditionalEventDefinition() *ConditionalEventDefinition
}

type DefinitionsBase interface {
	eventsbase.EventsBase
}
