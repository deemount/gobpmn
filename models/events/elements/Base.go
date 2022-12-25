package elements

import (
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/marker"
)

type EventElementsID interface{ eventsbase.EventsID }
type EventElementsName interface{ eventsbase.EventsName }

type EventElementsMarkerIncoming interface {
	SetIncoming(num int)
	GetIncoming(num int) *marker.Incoming
}

type EventElementsMarkerOutgoing interface {
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
}

type EventsElementsDefinitions interface {
	definitions.DefinitionsGetElements
	eventsbase.EventsSetDefinitions
}

type EventElementsCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
}

type EventElementsBase interface {
	EventElementsID
	EventElementsName
}
