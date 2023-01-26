package definitions

import (
	"github.com/deemount/gobpmn/models/bpmn/conditional"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/time"
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
	impl.IFBaseID
	impl.IFBaseName
}

// CancelEventDefinitionRepository ...
type CancelEventDefinitionRepository interface {
	impl.IFBaseID
}

// CompensateEventDefinitionRepository ...
type CompensateEventDefinitionRepository interface{}

// ConditionalEventDefinitionRepository ...
type ConditionalEventDefinitionRepository interface {
	impl.IFBaseID

	SetCamundaVariableName(variableName string)
	GetCamundaVariableName() *string

	SetCondition()
	GetCondition() *conditional.Condition
}

// ErrorEventDefinitionRepository ...
type ErrorEventDefinitionRepository interface {
	impl.IFBaseID
}

// EscalationEventDefinitionRepository ...
type EscalationEventDefinitionRepository interface {
	impl.IFBaseID
}

// LinkEventDefinitionRepository ...
type LinkEventDefinitionRepository interface {
	impl.IFBaseID
}

// MessageEventDefinitionRepository ...
type MessageEventDefinitionRepository interface {
	impl.IFBaseID

	SetMsgRef(suffix string)
	GetMsgRef() *string
}

// SignalEventDefinitionRepository ...
type SignalEventDefinitionRepository interface {
	impl.IFBaseID

	SetSignalRef(suffix string)
	GetSignalRef() *string
}

// TerminateEventDefinitionRepository ...
type TerminateEventDefinitionRepository interface {
	impl.IFBaseID
}

// TimerEventDefinitionRepository ...
type TimerEventDefinitionRepository interface {
	impl.IFBaseID

	SetTimeDate()
	GetTimeDate() *time.TimeDate
	SetTimeDuration()
	GetTimeDuration() *time.TimeDuration
}
