package definitions

import (
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/conditional"
	"github.com/deemount/gobpmn/models/time"
)

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
	compulsion.IFBaseID
	compulsion.IFBaseName
}

// CancelEventDefinitionRepository ...
type CancelEventDefinitionRepository interface {
	compulsion.IFBaseID
}

// CompensateEventDefinitionRepository ...
type CompensateEventDefinitionRepository interface{}

// ConditionalEventDefinitionRepository ...
type ConditionalEventDefinitionRepository interface {
	compulsion.IFBaseID

	SetCamundaVariableName(variableName string)
	GetCamundaVariableName() *string

	SetCondition()
	GetCondition() *conditional.Condition
}

// ErrorEventDefinitionRepository ...
type ErrorEventDefinitionRepository interface {
	compulsion.IFBaseID
}

// EscalationEventDefinitionRepository ...
type EscalationEventDefinitionRepository interface {
	compulsion.IFBaseID
}

// LinkEventDefinitionRepository ...
type LinkEventDefinitionRepository interface {
	compulsion.IFBaseID
}

// MessageEventDefinitionRepository ...
type MessageEventDefinitionRepository interface {
	compulsion.IFBaseID

	SetMsgRef(suffix string)
	GetMsgRef() *string
}

// SignalEventDefinitionRepository ...
type SignalEventDefinitionRepository interface {
	compulsion.IFBaseID

	SetSignalRef(suffix string)
	GetSignalRef() *string
}

// TerminateEventDefinitionRepository ...
type TerminateEventDefinitionRepository interface {
	compulsion.IFBaseID
}

// TimerEventDefinitionRepository ...
type TimerEventDefinitionRepository interface {
	compulsion.IFBaseID

	SetTimeDate()
	GetTimeDate() *time.TimeDate
	SetTimeDuration()
	GetTimeDuration() *time.TimeDuration
}
