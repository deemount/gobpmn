package process

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/camunda"
	"github.com/deemount/gobpmn/models/bpmn/data"
	"github.com/deemount/gobpmn/models/bpmn/events"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/gateways"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/pool"
	"github.com/deemount/gobpmn/models/bpmn/subprocesses"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
)

// ProcessRepository ...
type ProcessRepository interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
	events.EventsElementsRepository
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository
	subprocesses.SubprocessesElementsRepository
	camunda.CamundaProcessAttributes
	flow.FlowSequenceFlow

	SetIsExecutable(isExec bool)
	GetIsExecutable() *bool

	SetLaneSet()
	GetLaneSet() *pool.LaneSet

	SetDataObject(num int)
	GetDataObject(num int) *data.DataObject
}
