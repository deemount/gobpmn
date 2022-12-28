package subprocesses

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/events/elements"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/loop"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/tasks"
)

func NewAdHocSubProcess() AdHocSubProcessRepository {
	return &AdHocSubProcess{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (adhoc *AdHocSubProcess) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		adhoc.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		adhoc.ID = fmt.Sprintf("%v", suffix)
		break
	}
}

// SetName ...
func (adhoc *AdHocSubProcess) SetName(name string) {
	adhoc.Name = name
}

// SetTriggeredByEvent ...
func (adhoc *AdHocSubProcess) SetTriggeredByEvent(triggered bool) {
	adhoc.TriggeredByEvent = triggered
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (adhoc *AdHocSubProcess) SetCamundaAsyncBefore(asyncBefore bool) {
	adhoc.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (adhoc *AdHocSubProcess) SetCamundaAsyncAfter(asyncAfter bool) {
	adhoc.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (adhoc *AdHocSubProcess) SetCamundaJobPriority(priority int) {
	adhoc.CamundaJobPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (adhoc *AdHocSubProcess) SetDocumentation() {
	adhoc.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (adhoc *AdHocSubProcess) SetExtensionElements() {
	adhoc.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (adhoc *AdHocSubProcess) SetIncoming(num int) {
	adhoc.Incoming = make([]marker.Incoming, num)
}

/*** Marker ***/

// SetOutgoing ...
func (adhoc *AdHocSubProcess) SetOutgoing(num int) {
	adhoc.Outgoing = make([]marker.Outgoing, num)
}

/*** Loop ***/

// SetMultiInstaceLoopCharacteristics ...
func (adhoc *AdHocSubProcess) SetMultiInstanceLoopCharacteristics() {
	adhoc.MultiInstanceLoopCharacteristics = make([]loop.MultiInstanceLoopCharacteristics, 1)
}

/*** Events ***/

// SetStartEvent ...
func (adhoc *AdHocSubProcess) SetStartEvent(num int) {
	adhoc.StartEvent = make([]elements.StartEvent, num)
}

// SetEndEvent ...
func (adhoc *AdHocSubProcess) SetEndEvent(num int) {
	adhoc.EndEvent = make([]elements.EndEvent, num)
}

/*** Tasks ***/

// SetBusinessRuleTask ...
func (adhoc *AdHocSubProcess) SetBusinessRuleTask(num int) {
	adhoc.BusinessRuleTask = make([]tasks.BusinessRuleTask, num)
}

// SetTask ...
func (adhoc *AdHocSubProcess) SetTask(num int) {
	adhoc.Task = make([]tasks.Task, num)
}

// SetUserTask ...
func (adhoc *AdHocSubProcess) SetUserTask(num int) {
	adhoc.UserTask = make([]tasks.UserTask, num)
}

// SetManualTask ...
func (adhoc *AdHocSubProcess) SetManualTask(num int) {
	adhoc.ManualTask = make([]tasks.ManualTask, num)
}

// SetReceiveTask ...
func (adhoc *AdHocSubProcess) SetReceiveTask(num int) {
	adhoc.ReceiveTask = make([]tasks.ReceiveTask, num)
}

// SetScriptTask ...
func (adhoc *AdHocSubProcess) SetScriptTask(num int) {
	adhoc.ScriptTask = make([]tasks.ScriptTask, num)
}

// SetSendTask ...
func (adhoc *AdHocSubProcess) SetSendTask(num int) {
	adhoc.SendTask = make([]tasks.SendTask, num)
}

// SetServiceTask ...
func (adhoc *AdHocSubProcess) SetServiceTask(num int) {
	adhoc.ServiceTask = make([]tasks.ServiceTask, num)
}

/*** Subprocess ***/

// SetSubProcess ...
func (adhoc *AdHocSubProcess) SetSubProcess(num int) {
	adhoc.SubProcess = make([]SubProcess, num)
}

// SetAdHocSubProcess ...
func (adhoc *AdHocSubProcess) SetAdHocSubProcess(num int) {
	adhoc.AdHocSubProcess = make([]AdHocSubProcess, num)
}

/*** Gateway ***/

// SetExclusiveGateway
func (adhoc *AdHocSubProcess) SetExclusiveGateway(num int) {
	adhoc.ExclusiveGateway = make([]gateways.ExclusiveGateway, num)
}

// SetInclsuiveGateway
func (adhoc *AdHocSubProcess) SetInclusiveGateway(num int) {
	adhoc.InclusiveGateway = make([]gateways.InclusiveGateway, num)
}

// SetParallelGateway
func (adhoc *AdHocSubProcess) SetParallelGateway(num int) {
	adhoc.ParallelGateway = make([]gateways.ParallelGateway, num)
}

// SetComplexGateway
func (adhoc *AdHocSubProcess) SetComplexGateway(num int) {
	adhoc.ComplexGateway = make([]gateways.ComplexGateway, num)
}

// SetEventBasedGateway
func (adhoc *AdHocSubProcess) SetEventBasedGateway(num int) {
	adhoc.EventBasedGateway = make([]gateways.EventBasedGateway, num)
}

/*** Marker ***/

// SetSequenceFlow ...
func (adhoc *AdHocSubProcess) SetSequenceFlow(num int) {
	adhoc.SequenceFlow = make([]marker.SequenceFlow, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (adhoc *AdHocSubProcess) GetID() STR_PTR {
	return &adhoc.ID
}

// GetName ...
func (adhoc AdHocSubProcess) GetName() STR_PTR {
	return &adhoc.Name
}

// GetTriggeredByEvent ...
func (adhoc AdHocSubProcess) GetTriggeredByEvent() *bool {
	return &adhoc.TriggeredByEvent
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (adhoc AdHocSubProcess) GetCamundaAsyncBefore() *bool {
	return &adhoc.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (adhoc AdHocSubProcess) GetCamundaAsyncAfter() *bool {
	return &adhoc.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (adhoc AdHocSubProcess) GetCamundaJobPriority() *int {
	return &adhoc.CamundaJobPriority
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (adhoc AdHocSubProcess) GetDocumentation() *attributes.Documentation {
	return &adhoc.Documentation[0]
}

// GetExtensionElements ...
func (adhoc AdHocSubProcess) GetExtensionElements() *attributes.ExtensionElements {
	return &adhoc.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (adhoc AdHocSubProcess) GetIncoming(num int) *marker.Incoming {
	return &adhoc.Incoming[num]
}

// GetOutgoing ...
func (adhoc AdHocSubProcess) GetOutgoing(num int) *marker.Outgoing {
	return &adhoc.Outgoing[num]
}

/*** Loop ***/

// GetMultiInstaceLoopCharacteristics ...
func (adhoc AdHocSubProcess) GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics {
	return &adhoc.MultiInstanceLoopCharacteristics[0]
}

/*** Events ***/

// GetStartEvent ...
func (adhoc AdHocSubProcess) GetStartEvent(num int) *elements.StartEvent {
	return &adhoc.StartEvent[num]
}

// GetEndEvent ...
func (adhoc AdHocSubProcess) GetEndEvent(num int) *elements.EndEvent {
	return &adhoc.EndEvent[num]
}

/*** Tasks ***/

// GetBusinessRuleTask ...
func (adhoc AdHocSubProcess) GetBusinessRuleTask(num int) *tasks.BusinessRuleTask {
	return &adhoc.BusinessRuleTask[num]
}

// GetTask ...
func (adhoc AdHocSubProcess) GetTask(num int) *tasks.Task {
	return &adhoc.Task[num]
}

// GetUserTask ...
func (adhoc AdHocSubProcess) GetUserTask(num int) *tasks.UserTask {
	return &adhoc.UserTask[num]
}

// GetManualTask ...
func (adhoc AdHocSubProcess) GetManualTask(num int) *tasks.ManualTask {
	return &adhoc.ManualTask[num]
}

// GetReceiveTask ...
func (adhoc AdHocSubProcess) GetReceiveTask(num int) *tasks.ReceiveTask {
	return &adhoc.ReceiveTask[num]
}

// GetScriptTask ...
func (adhoc AdHocSubProcess) GetScriptTask(num int) *tasks.ScriptTask {
	return &adhoc.ScriptTask[num]
}

// GetSendTask ...
func (adhoc AdHocSubProcess) GetSendTask(num int) *tasks.SendTask {
	return &adhoc.SendTask[num]
}

// GetServiceTask ...
func (adhoc AdHocSubProcess) GetServiceTask(num int) *tasks.ServiceTask {
	return &adhoc.ServiceTask[num]
}

/*** Subprocess ***/

// GetSubProcess ...
func (adhoc AdHocSubProcess) GetSubProcess(num int) *SubProcess {
	return &adhoc.SubProcess[num]
}

// GetAdHocSubProcess ...
func (adhoc AdHocSubProcess) GetAdHocSubProcess(num int) *AdHocSubProcess {
	return &adhoc.AdHocSubProcess[num]
}

/*** Gateway ***/

// GetExclusiveGateway
func (adhoc AdHocSubProcess) GetExclusiveGateway(num int) *gateways.ExclusiveGateway {
	return &adhoc.ExclusiveGateway[num]
}

// GetInclsuiveGateway
func (adhoc AdHocSubProcess) GetInclusiveGateway(num int) *gateways.InclusiveGateway {
	return &adhoc.InclusiveGateway[num]
}

// GetParallelGateway
func (adhoc AdHocSubProcess) GetParallelGateway(num int) *gateways.ParallelGateway {
	return &adhoc.ParallelGateway[num]
}

// GetComplexGateway
func (adhoc AdHocSubProcess) GetComplexGateway(num int) *gateways.ComplexGateway {
	return &adhoc.ComplexGateway[num]
}

// GetEventBasedGateway
func (adhoc AdHocSubProcess) GetEventBasedGateway(num int) *gateways.EventBasedGateway {
	return &adhoc.EventBasedGateway[num]
}

/*** Marker ***/

// GetSequenceFlow ...
func (adhoc AdHocSubProcess) GetSequenceFlow(num int) *marker.SequenceFlow {
	return &adhoc.SequenceFlow[num]
}
