package process

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/data"
	"github.com/deemount/gobpmn/models/events"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/models/subprocesses"
	"github.com/deemount/gobpmn/models/tasks"
)

// ProcessRepository ...
type ProcessRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
	attributes.AttributesBaseElements
	events.EventsElementsRepository
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository
	subprocesses.SubprocessesElementsRepository
	camunda.CamundaProcessAttributes
	marker.MarkerSequenceFlow

	SetIsExecutable(isExec bool)
	GetIsExecutable() *bool

	SetLaneSet()
	GetLaneSet() *pool.LaneSet

	SetDataObject(num int)
	GetDataObject(num int) *data.DataObject
}
