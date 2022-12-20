package models

import (
	"github.com/deemount/gobpmn/models/activities"
	"github.com/deemount/gobpmn/models/events"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/subprocesses"
)

type Type_Definitions DefinitionsRepository

type Type_Events events.EventsRepository[
	events.Type_Cancel_Event_Definition,
	events.Type_Compensate_Event_Definition,
	events.Type_Conditional_Event_Definition,
	events.Type_Error_Event_Definition,
	events.Type_Escalation_Event_Definition,
	events.Type_Link_Event_Definition,
	events.Type_Start_Event,
	events.Type_End_Event,
	events.Type_Boundary_Event,
	events.Type_Signal_Event_Definition,
	events.Type_Timer_Event_Definition,
	events.Type_Terminate_Event_Definition,
	events.Type_Intermediate_Catch_Event,
	events.Type_Intermediate_Throw_Event]

type Type_Activities activities.ActivitiesRepository[
	activities.Type_Task,
	activities.Type_User_Task,
	activities.Type_Receive_Task,
	activities.Type_Manual_Task,
	activities.Type_Script_Task,
	activities.Type_Send_Task,
	activities.Type_Business_Rule_Task]

type Type_Subprocesses subprocesses.SubprocessesRepository[subprocesses.Type_Call_Activity]

type Type_Marker marker.MarkerRepository[
	marker.Type_Incoming,
	marker.Type_Outgoing,
	marker.Type_Sequence_Flow,
	marker.Type_Message_Flow,
	marker.Type_Signal]

type IDefinition[T Type_Definitions] interface {
	DefinitionsRepository() T
}

type IEvent[E Type_Events] interface {
	EventsRepository() E
}

type IActivities[A Type_Activities] interface {
	ActivitiesRepository() A
}

type ISubprocesses[S Type_Subprocesses] interface {
	SubprocessesRepository() S
}

type IMarker[M Type_Marker] interface {
	MarkerRepository() M
}

type IDiagram interface {
	PlaneRepository
	ShapeRepository
	WaypointRepository
}

type Repository[
	Definitions Type_Definitions,
	Events Type_Events,
	Activities Type_Activities,
	Subprocesses Type_Subprocesses,
	Marker Type_Marker] interface {
	IDefinition() Definitions
	IEvent() Events
	IActivities() Activities
	ISubprocesses() Subprocesses
	IMarker() Marker
	/*
		CollaborationRepository
		ParticipantRepository
		LaneRepository
		ProcessRepository
		IMarker
		IDiagram
	*/
}

type Models[
	Definitions Type_Definitions,
	Events Type_Events,
	Activities Type_Activities,
	Subprocesses Type_Subprocesses,
	Marker Type_Marker] struct {
	Repository[
		Definitions,
		Events,
		Activities,
		Subprocesses,
		Marker]
}

func NewModels() Repository[Type_Definitions, Type_Events, Type_Activities, Type_Subprocesses, Type_Marker] {
	return &Models[Type_Definitions, Type_Events, Type_Activities, Type_Subprocesses, Type_Marker]{}
}

func (m *Models[Type_Definitions, Type_Events, Type_Activities, Type_Subprocesses, Type_Marker]) Writer() {

}
