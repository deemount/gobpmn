package subprocesses

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/camunda"
	"github.com/deemount/gobpmn/models/bpmn/events"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/gateways"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/loop"
	"github.com/deemount/gobpmn/models/bpmn/marker"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
)

type SubprocessesBase interface {
	impl.IFBaseID
	impl.IFBaseName
	attributes.AttributesBaseElements
	marker.MarkerIncomingOutgoing
	camunda.CamundaDefaultAttributes
}

// SubprocessesElementsRepository ...
type SubprocessesElementsRepository interface {
	SetCallActivity(num int)
	GetCallActivity(num int) *CallActivity
	SetSubProcess(num int)
	GetSubProcess(num int) *SubProcess
	SetTransaction(num int)
	GetTransaction(num int) *Transaction
	SetAdHocSubProcess(num int)
	GetAdHocSubProcess(num int) *AdHocSubProcess
}

// AdHocSubProcessRepository ...
type AdHocSubProcessRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
	gateways.GatewaysElementsRepository
	tasks.TasksElementsRepository

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() *bool

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics

	SetStartEvent(num int)
	GetStartEvent(num int) *elements.StartEvent
	SetEndEvent(num int)
	GetEndEvent(num int) *elements.EndEvent

	SetSubProcess(num int)
	GetSubProcess(num int) *SubProcess
	SetAdHocSubProcess(num int)
	GetAdHocSubProcess(num int) *AdHocSubProcess
}

// CallActivityRepository ...
type CallActivityRepository interface {
	SubprocessesBase

	SetCalledElement(element string)
	GetCalledElement() *string

	SetCamundaCalledElementTenantID(tenantID string)
	GetCamundaCalledElementTenantID() *string

	SetCamundaVariableMappingClass(class string)
	GetCamundaVariableMappingClass() *string

	SetStandardLoopCharacteristics()
	GetStandardLoopCharacteristics() *loop.StandardLoopCharacteristics

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics
}

// SubProcessRepository ...
type SubProcessRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository
	events.EventsElementsRepository
	SubprocessesElementsRepository

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() *bool

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics
}

// TransactionRepository ...
type TransactionRepository interface {
	SubprocessesBase
	flow.FlowSequenceFlow
}
