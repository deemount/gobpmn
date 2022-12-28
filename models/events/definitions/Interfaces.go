package definitions

import (
	"github.com/deemount/gobpmn/models/conditional"
	"github.com/deemount/gobpmn/models/time"
)

type DefinitionsID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type DefinitionsName interface {
	SetName(name string)
	GetName() STR_PTR
}

// @EndEvent only
type DefinitionsGetTerminateBase interface {
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
	DefinitionsID
	DefinitionsName
}

// CancelEventDefinitionRepository ...
type CancelEventDefinitionRepository interface {
	DefinitionsID
}

// CompensateEventDefinitionRepository ...
type CompensateEventDefinitionRepository interface{}

// ConditionalEventDefinitionRepository ...
type ConditionalEventDefinitionRepository interface {
	DefinitionsID

	SetCamundaVariableName(variableName string)
	GetCamundaVariableName() *string

	SetCondition()
	GetCondition() *conditional.Condition
}

// ErrorEventDefinitionRepository ...
type ErrorEventDefinitionRepository interface {
	DefinitionsID
}

// EscalationEventDefinitionRepository ...
type EscalationEventDefinitionRepository interface {
	DefinitionsID
}

// LinkEventDefinitionRepository ...
type LinkEventDefinitionRepository interface {
	DefinitionsID
}

// MessageEventDefinitionRepository ...
type MessageEventDefinitionRepository interface {
	DefinitionsID

	SetMsgRef(suffix string)
	GetMsgRef() *string
}

// SignalEventDefinitionRepository ...
type SignalEventDefinitionRepository interface {
	DefinitionsID

	SetSignalRef(suffix string)
	GetSignalRef() *string
}

// TerminateEventDefinitionRepository ...
type TerminateEventDefinitionRepository interface {
	DefinitionsID
}

// TimerEventDefinitionRepository ...
type TimerEventDefinitionRepository interface {
	DefinitionsID

	SetTimeDate()
	GetTimeDate() *time.TimeDate
	SetTimeDuration()
	GetTimeDuration() *time.TimeDuration
}
