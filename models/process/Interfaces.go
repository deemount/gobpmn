package process

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/data"
	"github.com/deemount/gobpmn/models/events"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/models/subprocesses"
	"github.com/deemount/gobpmn/models/tasks"
)

type ProcessBaseID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type ProcessBaseName interface {
	SetName(name string)
	GetName() STR_PTR
}

type ProcessBaseDocumentation interface {
	SetDocumentation()
	GetDocumentation() *attributes.Documentation
}

type ProcessBaseExtensionElements interface {
	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

type ProcessBaseCoreElements interface {
	ProcessBaseDocumentation
	ProcessBaseExtensionElements
}

type ProcessBase interface {
	ProcessBaseID
	ProcessBaseName
	ProcessBaseCoreElements
}

// ProcessRepository ...
type ProcessRepository interface {
	ProcessBase
	events.EventsElementsRepository
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository
	subprocesses.SubprocessesElementsRepository

	SetIsExecutable(isExec bool)
	GetIsExecutable() *bool

	SetSequenceFlow(num int)
	GetSequenceFlow(num int) *marker.SequenceFlow

	SetCamundaVersionTag(tag string)
	GetCamundaVersionTag() *string
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
	SetCamundaTaskPriority(priority int)
	GetCamundaTaskPriority() *int
	SetCamundaCandidateStarterGroups(groups string)
	GetCamundaCandidateStarterGroups() *string
	SetCamundaCandiddateStarterUsers(users string)
	GetCamundaCandiddateStarterUsers() *string
	SetCamundaHistoryTimeToLive(tolive string)
	GetCamundaHistoryTimeToLive() *string

	SetLaneSet()
	GetLaneSet() *pool.LaneSet

	SetDataObject(num int)
	GetDataObject(num int) *data.DataObject
}
