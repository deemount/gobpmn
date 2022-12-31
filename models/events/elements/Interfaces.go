package elements

import (
	"context"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/marker"
)

// Notice: belongs to definitions package
type EventsSetTerminateBase interface {
	SetTerminateEventDefinition()
}

// Notice: belongs to definitions package
type EventsSetDefinitionsBase interface {
	SetMessageEventDefinition()
	SetEscalationEventDefinition()
	SetErrorEventDefinition()
	SetSignalEventDefinition()
	SetCompensateEventDefinition()
}

// Notice: belongs to definitions package
type EventsSetDefinitions interface {
	EventsSetDefinitionsBase

	SetTimerEventDefinition()
	SetCancelEventDefinition()
	SetConditionalEventDefinition()
}

type EventElementsDefinitions interface {
	definitions.DefinitionsGetElements
	EventsSetDefinitions
}

type EventElementsCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
}

type EventElementsCoreThrowCatchElements interface {
	SetLinkEventDefinition()
	GetLinkEventDefinition() *definitions.LinkEventDefinition
	SetMessageEventDefinition()
	GetMessageEventDefinition() *definitions.MessageEventDefinition
}

// BoundaryEventRepository ...
type BoundaryEventRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
	EventElementsDefinitions
	marker.MarkerOutgoing

	SetAttachedToRef(ref string)
	GetAttachedToRef() *string

	// maybe @deprecated @7.17 execution platform
	// Notice: maybe token out of a older modeler version.
	// Needs a checkout
	SetCancelActivity(cancel bool)
	GetCancelActivity() *bool

	String() string
}

// TBoundaryEventRepository ...
type TBoundaryEventRepository interface {
	String() string
	Handle(ctx context.Context) error
}

// EndEventRepository ...
type EndEventRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
	EventElementsCamundaBase
	marker.MarkerIncoming
	attributes.AttributesBaseElements
	EventsSetDefinitionsBase
	EventsSetTerminateBase
	definitions.DefinitionsGetElementsBase
	definitions.DefinitionsGetTerminateBase

	String() string
}

// TEndEventRepository ...
type TEndEventRepository interface {
	String() string
	Handle(ctx context.Context) error
}

// IntermediateCatchEventRepository ...
type IntermediateCatchEventRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
	EventElementsCamundaBase
	attributes.AttributesBaseElements
	marker.MarkerIncomingOutgoing
	EventElementsCoreThrowCatchElements

	SetConditionalEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	SetTimerEventDefinition()
	GetTimerEventDefinition() *definitions.TimerEventDefinition

	String() string
}

// TIntermediateCatchEventRepository ...
type TIntermediateCatchEventRepository interface {
	String() string
	Handle(ctx context.Context) error
}

// IntermediateThrowEventRepository ...
type IntermediateThrowEventRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
	marker.MarkerIncomingOutgoing
	attributes.AttributesBaseElements
	EventElementsCoreThrowCatchElements

	SetCompensateEventDefinition()
	GetCompensateEventDefinition() *definitions.CompensateEventDefinition
	SetEscalationEventDefinition()
	GetEscalationEventDefinition() *definitions.EscalationEventDefinition

	String() string
}

// TIntermediateThrowEventRepository ...
type TIntermediateThrowEventRepository interface {
	String() string
	Handle(ctx context.Context) error
}

// StartEventRepository ...
type StartEventRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
	EventElementsCamundaBase
	marker.MarkerOutgoing
	attributes.AttributesBaseElements

	SetIsInterrupting(isInterrupt bool)
	GetIsInterrupting() *bool

	SetCamundaFormKey(key string)
	GetCamundaFormKey() *string
	SetCamundaFormRef(ref string)
	GetCamundaFormRef() *string
	SetCamundaFormRefBinding(bind string)
	GetCamundaFormRefBinding() *string
	SetCamundaFormRefVersion(version string)
	GetCamundaFormRefVersion() *string
	SetCamundaInitiator(initiator string)
	GetCamundaInitiator() *string

	SetConditionalEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	SetTimerEventDefinition()
	GetTimerEventDefinition() *definitions.TimerEventDefinition

	SetMessageEventDefinition()
	GetMessageEventDefinition() *definitions.MessageEventDefinition

	String() string
}

// TStartEventRepository ...
type TStartEventRepository interface {
	String() string
	Handle(ctx context.Context) error
}
