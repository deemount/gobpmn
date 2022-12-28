package subprocesses

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events/elements"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/loop"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/tasks"
)

type SubprocessesID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type SubprocessesName interface {
	SetName(name string)
	GetName() STR_PTR
}

type SubprocessesMarker interface {
	SetIncoming(num int)
	GetIncoming(num int) *marker.Incoming
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
}

type SubprocessesFlow interface {
	SetSequenceFlow(num int)
	GetSequenceFlow(num int) *marker.SequenceFlow
}

type SubprocessesCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
}

type SubprocessesBaseCoreElements interface {
	SetDocumentation()
	GetDocumentation() *attributes.Documentation
	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

type SubprocessesBase interface {
	SubprocessesID
	SubprocessesName
	SubprocessesBaseCoreElements
	SubprocessesMarker
	SubprocessesCamundaBase
}

// SubprocessesElementsRepository ...
type SubprocessesElementsRepository interface {
	SetCallActivity(num int)
	GetCallActivity(num int) *CallActivity
	SetSubProcess(num int)
	GetSubProcess(num int) *SubProcess
	SetTransaction(num int)
	GetTransaction(num int) *Transaction
	SetAdHocSubprocess(num int)
	GetAdHocSubProcess(num int) *AdHocSubProcess
}

// AdHocSubProcessRepository ...
type AdHocSubProcessRepository interface {
	SubprocessesBase
	SubprocessesFlow
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
	SubprocessesFlow
	tasks.TasksElementsRepository
	gateways.GatewaysElementsRepository

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() *bool

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics

	SetStartEvent(num int)
	GetStartEvent(num int) *elements.StartEvent
	SetEndEvent(num int)
	GetEndEvent(num int) *elements.EndEvent

	SetSubProcess(num int)
	SetAdHocSubProcess(num int)
	GetSubProcess(num int) *SubProcess
	GetAdHocSubProcess(num int) *AdHocSubProcess
}

// TransactionRepository ...
type TransactionRepository interface {
	SubprocessesBase
	SubprocessesFlow
}
