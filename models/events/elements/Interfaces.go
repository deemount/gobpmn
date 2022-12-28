package elements

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/marker"
)

type EventElementsID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}
type EventElementsName interface {
	SetName(name string)
	GetName() STR_PTR
}

type EventElementsMarkerIncoming interface {
	SetIncoming(num int)
	GetIncoming(num int) *marker.Incoming
}

type EventElementsMarkerOutgoing interface {
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
}

type EventElementsMarker interface {
	EventElementsMarkerIncoming
	EventElementsMarkerOutgoing
}

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

type EventElementsCoreElements interface {
	SetDocumentation()
	GetDocumentation() *attributes.Documentation
	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

type EventElementsBase interface {
	EventElementsID
	EventElementsName
}

// BoundaryEventRepository ...
type BoundaryEventRepository interface {
	EventElementsBase
	EventElementsDefinitions
	EventElementsMarkerOutgoing

	SetAttachedToRef(ref string)
	GetAttachedToRef() *string

	// maybe @deprecated @7.17 execution platform
	// Notice: maybe token out of a older modeler version.
	// Needs a checkout
	SetCancelActivity(cancel bool)
	GetCancelActivity() *bool
}

// EndEventRepository ...
type EndEventRepository interface {
	EventElementsBase
	EventElementsCamundaBase
	EventElementsMarkerIncoming
	EventElementsCoreElements
	EventsSetDefinitionsBase
	EventsSetTerminateBase
	definitions.DefinitionsGetElementsBase
	definitions.DefinitionsGetTerminateBase
}

// IntermediateCatchEventRepository ...
type IntermediateCatchEventRepository interface {
	EventElementsBase
	EventElementsCamundaBase
	EventElementsCoreElements
	EventElementsMarker
	EventElementsCoreThrowCatchElements

	SetConditionalEventDefinition()
	GetConditionalEventDefinition() *definitions.ConditionalEventDefinition
	SetTimerEventDefinition()
	GetTimerEventDefinition() *definitions.TimerEventDefinition
}

// IntermediateThrowEventRepository ...
type IntermediateThrowEventRepository interface {
	EventElementsBase
	EventElementsMarker
	EventElementsCoreElements
	EventElementsCoreThrowCatchElements

	SetCompensateEventDefinition()
	GetCompensateEventDefinition() *definitions.CompensateEventDefinition
	SetEscalationEventDefinition()
	GetEscalationEventDefinition() *definitions.EscalationEventDefinition
}

// StartEventRepository ...
type StartEventRepository interface {
	EventElementsBase
	EventElementsCamundaBase
	EventElementsMarkerOutgoing
	EventElementsCoreElements

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
