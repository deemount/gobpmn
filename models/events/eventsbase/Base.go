package eventsbase

type STR_PTR *string

type EventsID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type EventsName interface {
	SetName(name string)
	GetName() STR_PTR
}

type EventsMarker interface {
}

type EventsSetTerminateBase interface {
	SetTerminateEventDefinition()
}

type EventsSetDefinitionsBase interface {
	SetMessageEventDefinition()
	SetEscalationEventDefinition()
	SetErrorEventDefinition()
	SetSignalEventDefinition()
	SetCompensateEventDefinition()
}

type EventsSetDefinitions interface {
	EventsSetDefinitionsBase

	SetTimerEventDefinition()
	SetCancelEventDefinition()
	SetConditionalEventDefinition()
}

type EventsCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
}

type EventsBase interface {
	EventsID
	EventsName
}
