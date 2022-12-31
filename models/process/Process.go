package process

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/data"
	"github.com/deemount/gobpmn/models/events/elements"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/models/subprocesses"
	"github.com/deemount/gobpmn/models/tasks"
)

// NewProcess ...
func NewProcess() ProcessRepository {
	return &Process{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (process *Process) SetID(typ string, suffix interface{}) {
	switch typ {
	case "hash":
		process.ID = fmt.Sprintf("Process_%v", suffix)
		break
	case "id":
		process.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (process *Process) SetName(name string) {
	process.Name = name
}

// SetIsExecutable ...
func (process *Process) SetIsExecutable(isExec bool) {
	process.IsExecutable = isExec
}

/** Camunda **/

// SetCamundaVersionTag ...
func (process *Process) SetCamundaVersionTag(tag string) {
	process.CamundaVersionTag = tag
}

// SetCamundaJobpriority ...
func (process *Process) SetCamundaJobPriority(priority int) {
	process.CamundaJobPriority = priority
}

// SetCamundaTaskPriority ...
func (process *Process) SetCamundaTaskPriority(priority int) {
	process.CamundaTaskPriority = priority
}

// SetCamundaCandidateStarterGroups ...
func (process *Process) SetCamundaCandidateStarterGroups(groups string) {
	process.CamundaCandidateStarterGroups = groups
}

// SetCamundaCandidateStarterUsers
func (process *Process) SetCamundaCandiddateStarterUsers(users string) {
	process.CamundaCandidateStarterUsers = users
}

// SetCamundaHistoryTimeToLive ...
func (process *Process) SetCamundaHistoryTimeToLive(tolive string) {
	process.CamundaHistoryTimeToLive = tolive
}

/*** Make Elements ***/

/** BPMN **/

/** Documentation **/

// SetDocumentation ...
func (process *Process) SetDocumentation() {
	process.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (process *Process) SetExtensionElements() {
	process.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/** Pool **/

// SetLaneSet ...
func (process *Process) SetLaneSet() {
	process.LaneSet = make([]pool.LaneSet, 1)
}

/*** Events ***/

// SetStartEvent ...
func (process *Process) SetStartEvent(num int) {
	process.StartEvent = make([]elements.StartEvent, num)
}

// SetBoundEvent ...
func (process *Process) SetBoundaryEvent(num int) {
	process.BoundaryEvent = make([]elements.BoundaryEvent, num)
}

// SetEndEvent ...
func (process *Process) SetEndEvent(num int) {
	process.EndEvent = make([]elements.EndEvent, num)
}

// SetIntermedCatchEvent ...
func (process *Process) SetIntermediateCatchEvent(num int) {
	process.IntermediateCatchEvent = make([]elements.IntermediateCatchEvent, num)
}

// SetIntermedThrowEvent ...
func (process *Process) SetIntermediateThrowEvent(num int) {
	process.IntermediateThrowEvent = make([]elements.IntermediateThrowEvent, num)
}

/*** Tasks ***/

// SetBusinessRuleTask ...
func (process *Process) SetBusinessRuleTask(num int) {
	process.BusinessRuleTask = make([]tasks.BusinessRuleTask, num)
}

// SetTask ...
func (process *Process) SetTask(num int) {
	process.Task = make([]tasks.Task, num)
}

// SetUserTask ...
func (process *Process) SetUserTask(num int) {
	process.UserTask = make([]tasks.UserTask, num)
}

// SetManualTask ...
func (process *Process) SetManualTask(num int) {
	process.ManualTask = make([]tasks.ManualTask, num)
}

// SetReceiveTask ...
func (process *Process) SetReceiveTask(num int) {
	process.ReceiveTask = make([]tasks.ReceiveTask, num)
}

// SetScriptTask ...
func (process *Process) SetScriptTask(num int) {
	process.ScriptTask = make([]tasks.ScriptTask, num)
}

// SetSendTask ...
func (process *Process) SetSendTask(num int) {
	process.SendTask = make([]tasks.SendTask, num)
}

// SetServiceTask ...
func (process *Process) SetServiceTask(num int) {
	process.ServiceTask = make([]tasks.ServiceTask, num)
}

/*** Subprocesses ***/

// SetCallActivity ...
func (process *Process) SetCallActivity(num int) {
	process.CallActivity = make([]subprocesses.CallActivity, num)
}

// SetSubProcess ...
func (process *Process) SetSubProcess(num int) {
	process.SubProcess = make([]subprocesses.SubProcess, num)
}

// SetTransaction ...
func (process *Process) SetTransaction(num int) {
	process.Transaction = make([]subprocesses.Transaction, num)
}

// SetAdHocSubProcess ...
func (process *Process) SetAdHocSubProcess(num int) {
	process.AdHocSubProcess = make([]subprocesses.AdHocSubProcess, num)
}

/*** Gateways ***/

// SetExclusiveGateway
func (process *Process) SetExclusiveGateway(num int) {
	process.ExclusiveGateway = make([]gateways.ExclusiveGateway, num)
}

// SetInclsuiveGateway
func (process *Process) SetInclusiveGateway(num int) {
	process.InclusiveGateway = make([]gateways.InclusiveGateway, num)
}

// SetParallelGateway
func (process *Process) SetParallelGateway(num int) {
	process.ParallelGateway = make([]gateways.ParallelGateway, num)
}

// SetComplexGateway
func (process *Process) SetComplexGateway(num int) {
	process.ComplexGateway = make([]gateways.ComplexGateway, num)
}

// SetEventBasedGateway
func (process *Process) SetEventBasedGateway(num int) {
	process.EventBasedGateway = make([]gateways.EventBasedGateway, num)
}

/*** Marker ***/

// SetAssociation ...
func (process *Process) SetAssociation(num int) {
	process.Association = make([]marker.Association, num)
}

// SetSequenceFlow ...
func (process *Process) SetSequenceFlow(num int) {
	process.SequenceFlow = make([]marker.SequenceFlow, num)
}

/*** Data ***/

// SetDataObject ...
func (process *Process) SetDataObject(num int) {
	process.DataObject = make([]data.DataObject, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

/*** Attributes ***/

// GetID ...
func (process Process) GetID() compulsion.STR_PTR {
	return &process.ID
}

// GetName ...
func (process Process) GetName() compulsion.STR_PTR {
	return &process.Name
}

// GetIsExecutable ...
func (process Process) GetIsExecutable() *bool {
	return &process.IsExecutable
}

/** Camunda **/

// GetCamundaVersionTag ...
func (process *Process) GetCamundaVersionTag() *string {
	return &process.CamundaVersionTag
}

// GetCamundaJobpriority ...
func (process *Process) GetCamundaJobPriority() *int {
	return &process.CamundaJobPriority
}

// GetCamundaTaskPriority ...
func (process *Process) GetCamundaTaskPriority() *int {
	return &process.CamundaTaskPriority
}

// GetCamundaCandidateStarterGroups ...
func (process *Process) GetCamundaCandidateStarterGroups() *string {
	return &process.CamundaCandidateStarterGroups
}

// GetCamundaCandidateStarterUsers
func (process *Process) GetCamundaCandiddateStarterUsers() *string {
	return &process.CamundaCandidateStarterUsers
}

// GetCamundaHistoryTimeToLive ...
func (process *Process) GetCamundaHistoryTimeToLive() *string {
	return &process.CamundaHistoryTimeToLive
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (process Process) GetDocumentation() *attributes.Documentation {
	return &process.Documentation[0]
}

// GetExtensionElements ...
func (process Process) GetExtensionElements() *attributes.ExtensionElements {
	return &process.ExtensionElements[0]
}

/*** Pool ***/

// GetLaneSet ...
func (process Process) GetLaneSet() *pool.LaneSet {
	return &process.LaneSet[0]
}

/*** Events ***/

// GetStartEvent ...
func (process Process) GetStartEvent(num int) *elements.StartEvent {
	return &process.StartEvent[num]
}

// GetBoundaryEvent ...
func (process Process) GetBoundaryEvent(num int) *elements.BoundaryEvent {
	return &process.BoundaryEvent[num]
}

// GetEndEvent ...
func (process Process) GetEndEvent(num int) *elements.EndEvent {
	return &process.EndEvent[num]
}

// GetIntermedCatchEvent ...
func (process Process) GetIntermediateCatchEvent(num int) *elements.IntermediateCatchEvent {
	return &process.IntermediateCatchEvent[num]
}

// GetIntermedThrowEvent ...
func (process Process) GetIntermediateThrowEvent(num int) *elements.IntermediateThrowEvent {
	return &process.IntermediateThrowEvent[num]
}

/*** Tasks ***/

// GetBusinessRuleTask ...
func (process Process) GetBusinessRuleTask(num int) *tasks.BusinessRuleTask {
	return &process.BusinessRuleTask[num]
}

// GetTask ...
func (process Process) GetTask(num int) *tasks.Task {
	return &process.Task[num]
}

// GetUserTask ...
func (process Process) GetUserTask(num int) *tasks.UserTask {
	return &process.UserTask[num]
}

// GetManualTask ...
func (process Process) GetManualTask(num int) *tasks.ManualTask {
	return &process.ManualTask[num]
}

// GetReceiveTask ...
func (process Process) GetReceiveTask(num int) *tasks.ReceiveTask {
	return &process.ReceiveTask[num]
}

// GetScriptTask ...
func (process Process) GetScriptTask(num int) *tasks.ScriptTask {
	return &process.ScriptTask[num]
}

// GetSendTask ...
func (process Process) GetSendTask(num int) *tasks.SendTask {
	return &process.SendTask[num]
}

// GetServiceTask ...
func (process Process) GetServiceTask(num int) *tasks.ServiceTask {
	return &process.ServiceTask[num]
}

// GetCallActivity ...
func (process Process) GetCallActivity(num int) *subprocesses.CallActivity {
	return &process.CallActivity[num]
}

// GetSubProcess ...
func (process Process) GetSubProcess(num int) *subprocesses.SubProcess {
	return &process.SubProcess[num]
}

// GetTransaction ...
func (process Process) GetTransaction(num int) *subprocesses.Transaction {
	return &process.Transaction[num]
}

// GetAdHocSubProcess ...
func (process Process) GetAdHocSubProcess(num int) *subprocesses.AdHocSubProcess {
	return &process.AdHocSubProcess[num]
}

/*** Gateways ***/

// GetExclusiveGateway
func (process Process) GetExclusiveGateway(num int) *gateways.ExclusiveGateway {
	return &process.ExclusiveGateway[num]
}

// GetInclsuiveGateway
func (process Process) GetInclusiveGateway(num int) *gateways.InclusiveGateway {
	return &process.InclusiveGateway[num]
}

// GetParallelGateway
func (process Process) GetParallelGateway(num int) *gateways.ParallelGateway {
	return &process.ParallelGateway[num]
}

// GetComplexGateway
func (process Process) GetComplexGateway(num int) *gateways.ComplexGateway {
	return &process.ComplexGateway[num]
}

// GetEventBasedGateway
func (process Process) GetEventBasedGateway(num int) *gateways.EventBasedGateway {
	return &process.EventBasedGateway[num]
}

/*** Marker ***/

// GetAssociation ...
func (process Process) GetAssociation(num int) *marker.Association {
	return &process.Association[num]
}

// GetSequenceFlow ...
func (process Process) GetSequenceFlow(num int) *marker.SequenceFlow {
	return &process.SequenceFlow[num]
}

/*** Data ***/

// GetDataObject ...
func (process Process) GetDataObject(num int) *data.DataObject {
	return &process.DataObject[num]
}
