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

// SubProcessRepository ...
type SubProcessRepository interface {
	SubprocessesBase
	SubprocessesFlow

	SetTriggeredByEvent(triggered bool)
	GetTriggeredByEvent() *bool

	SetMultiInstanceLoopCharacteristics()
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics

	SetStartEvent(num int)
	SetEndEvent(num int)

	SetTask(num int)
	SetUserTask(num int)
	SetManualTask(num int)
	SetReceiveTask(num int)
	SetScriptTask(num int)
	SetSendTask(num int)
	SetServiceTask(num int)

	SetSubProcess(num int)
	SetAdHocSubProcess(num int)

	SetExclusiveGateway(num int)
	SetInclusiveGateway(num int)
	SetParallelGateway(num int)
	SetComplexGateway(num int)
	SetEventBasedGateway(num int)

	GetStartEvent(num int) *elements.StartEvent
	GetEndEvent(num int) *elements.EndEvent
	GetTask(num int) *tasks.Task
	GetUserTask(num int) *tasks.UserTask
	GetManualTask(num int) *tasks.ManualTask
	GetReceiveTask(num int) *tasks.ReceiveTask
	GetScriptTask(num int) *tasks.ScriptTask
	GetSendTask(num int) *tasks.SendTask
	GetServiceTask(num int) *tasks.ServiceTask

	GetSubProcess(num int) *SubProcess
	GetAdHocSubProcess(num int) *AdHocSubProcess

	GetExclusiveGateway(num int) *gateways.ExclusiveGateway
	GetInclusiveGateway(num int) *gateways.InclusiveGateway
	GetParallelGateway(num int) *gateways.ParallelGateway
	GetComplexGateway(num int) *gateways.ComplexGateway
	GetEventBasedGateway(num int) *gateways.EventBasedGateway
}

// SubProcess ...
type SubProcess struct {
	ID                               string                                  `xml:"id,attr" json:"id"`
	Name                             string                                  `xml:"name,attr,omitempty" json:"name,omitempty"`
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	CamundaAsyncBefore               bool                                    `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter                bool                                    `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority               int                                     `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation                    []attributes.Documentation              `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements                []attributes.ExtensionElements          `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                         []marker.Incoming                       `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                         []marker.Outgoing                       `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	StartEvent                       []elements.StartEvent                   `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty"`
	EndEvent                         []elements.EndEvent                     `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	Task                             []tasks.Task                            `xml:"bpmn:task,omitempty" json:"task,omitempty"`
	UserTask                         []tasks.UserTask                        `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty"`
	ManualTask                       []tasks.ManualTask                      `xml:"bpmn:manualTask,omitempty" json:"manualTask,omitempty"`
	ReceiveTask                      []tasks.ReceiveTask                     `xml:"bpmn:receiveTask,omitempty" json:"receiveTask,omitempty"`
	ScriptTask                       []tasks.ScriptTask                      `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty"`
	SendTask                         []tasks.SendTask                        `xml:"bpmn:sendTask,omitempty" json:"sendTask,omitempty"`
	ServiceTask                      []tasks.ServiceTask                     `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty"`
	SubProcess                       []SubProcess                            `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  []AdHocSubProcess                       `xml:"bpmn:adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	ExclusiveGateway                 []gateways.ExclusiveGateway             `xml:"bpmn:exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
	InclusiveGateway                 []gateways.InclusiveGateway             `xml:"bpmn:inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
	ParallelGateway                  []gateways.ParallelGateway              `xml:"bpmn:parallelGateway,omitempty" json:"parallelGateway,omitempty"`
	ComplexGateway                   []gateways.ComplexGateway               `xml:"bpmn:complexGateway,omitempty" json:"complexGateway,omitempty"`
	EventBasedGateway                []gateways.EventBasedGateway            `xml:"bpmn:eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty"`
	SequenceFlow                     []marker.SequenceFlow                   `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// TSubProcess ...
type TSubProcess struct {
	ID                               string                                  `xml:"id,attr" json:"id"`
	Name                             string                                  `xml:"name,attr,omitempty" json:"name,omitempty"`
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	CamundaAsyncBefore               bool                                    `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter                bool                                    `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority               int                                     `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation                    []attributes.Documentation              `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements                []attributes.TExtensionElements         `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming                         []marker.Incoming                       `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing                         []marker.Outgoing                       `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	StartEvent                       []elements.TStartEvent                  `xml:"startEvent,omitemnpty" json:"startEvent,omitempty"`
	EndEvent                         []elements.TEndEvent                    `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	Task                             []tasks.TTask                           `xml:"task,omitempty" json:"task,omitempty"`
	UserTask                         []tasks.TUserTask                       `xml:"userTask,omitempty" json:"userTask,omitempty"`
	ManualTask                       []tasks.TManualTask                     `xml:"manualTask,omitempty" json:"manualTask,omitempty"`
	ReceiveTask                      []tasks.ReceiveTask                     `xml:"receiveTask,omitempty" json:"receiveTask,omitempty"`
	ScriptTask                       []tasks.ScriptTask                      `xml:"scriptTask,omitempty" json:"scriptTask,omitempty"`
	SendTask                         []tasks.SendTask                        `xml:"sendTask,omitempty" json:"sendTask,omitempty"`
	ServiceTask                      []tasks.ServiceTask                     `xml:"serviceTask,omitempty" json:"serviceTask,omitempty"`
	SubProcess                       []TSubProcess                           `xml:"subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  []TAdHocSubProcess                      `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	ExclusiveGateway                 []gateways.ExclusiveGateway             `xml:"exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
	InclusiveGateway                 []gateways.InclusiveGateway             `xml:"inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
	ParallelGateway                  []gateways.ParallelGateway              `xml:"parallelGateway,omitempty" json:"parallelGateway,omitempty"`
	ComplexGateway                   []gateways.ComplexGateway               `xml:"complexGateway,omitempty" json:"complexGateway,omitempty"`
	EventBasedGateway                []gateways.EventBasedGateway            `xml:"eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty"`
	SequenceFlow                     []marker.SequenceFlow                   `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (subprocess *SubProcess) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		subprocess.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		subprocess.ID = fmt.Sprintf("%v", suffix)
		break
	}
}

// SetName ...
func (subprocess *SubProcess) SetName(name string) {
	subprocess.Name = name
}

// SetTriggeredByEvent ...
func (subprocess *SubProcess) SetTriggeredByEvent(triggered bool) {
	subprocess.TriggeredByEvent = triggered
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (subprocess *SubProcess) SetCamundaAsyncBefore(asyncBefore bool) {
	subprocess.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (subprocess *SubProcess) SetCamundaAsyncAfter(asyncAfter bool) {
	subprocess.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (subprocess *SubProcess) SetCamundaJobPriority(priority int) {
	subprocess.CamundaJobPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (subprocess *SubProcess) SetDocumentation() {
	subprocess.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (subprocess *SubProcess) SetExtensionElements() {
	subprocess.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (subprocess *SubProcess) SetIncoming(num int) {
	subprocess.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (subprocess *SubProcess) SetOutgoing(num int) {
	subprocess.Outgoing = make([]marker.Outgoing, num)
}

// SetMultiInstaceLoopCharacteristics ...
func (subprocess *SubProcess) SetMultiInstanceLoopCharacteristics() {
	subprocess.MultiInstanceLoopCharacteristics = make([]loop.MultiInstanceLoopCharacteristics, 1)
}

/*** Event ***/

// SetStartEvent ...
func (subprocess *SubProcess) SetStartEvent(num int) {
	subprocess.StartEvent = make([]elements.StartEvent, num)
}

// SetEndEvent ...
func (subprocess *SubProcess) SetEndEvent(num int) {
	subprocess.EndEvent = make([]elements.EndEvent, num)
}

/*** Task ***/

// SetTask ...
func (subprocess *SubProcess) SetTask(num int) {
	subprocess.Task = make([]tasks.Task, num)
}

// SetUserTask ...
func (subprocess *SubProcess) SetUserTask(num int) {
	subprocess.UserTask = make([]tasks.UserTask, num)
}

// SetManualTask ...
func (subprocess *SubProcess) SetManualTask(num int) {
	subprocess.ManualTask = make([]tasks.ManualTask, num)
}

// SetReceiveTask ...
func (subprocess *SubProcess) SetReceiveTask(num int) {
	subprocess.ReceiveTask = make([]tasks.ReceiveTask, num)
}

// SetScriptTask ...
func (subprocess *SubProcess) SetScriptTask(num int) {
	subprocess.ScriptTask = make([]tasks.ScriptTask, num)
}

// SetSendTask ...
func (subprocess *SubProcess) SetSendTask(num int) {
	subprocess.SendTask = make([]tasks.SendTask, num)
}

// SetServiceTask ...
func (subprocess *SubProcess) SetServiceTask(num int) {
	subprocess.ServiceTask = make([]tasks.ServiceTask, num)
}

/*** Subprocess ***/

// SetSubProcess ...
func (subprocess *SubProcess) SetSubProcess(num int) {
	subprocess.SubProcess = make([]SubProcess, num)
}

// SetAdHocSubProcess ...
func (subprocess *SubProcess) SetAdHocSubProcess(num int) {
	subprocess.AdHocSubProcess = make([]AdHocSubProcess, num)
}

/*** Gateway ***/

// SetExclusiveGateway
func (subprocess *SubProcess) SetExclusiveGateway(num int) {
	subprocess.ExclusiveGateway = make([]gateways.ExclusiveGateway, num)
}

// SetInclsuiveGateway
func (subprocess *SubProcess) SetInclusiveGateway(num int) {
	subprocess.InclusiveGateway = make([]gateways.InclusiveGateway, num)
}

// SetParallelGateway
func (subprocess *SubProcess) SetParallelGateway(num int) {
	subprocess.ParallelGateway = make([]gateways.ParallelGateway, num)
}

// SetComplexGateway
func (subprocess *SubProcess) SetComplexGateway(num int) {
	subprocess.ComplexGateway = make([]gateways.ComplexGateway, num)
}

// SetEventBasedGateway
func (subprocess *SubProcess) SetEventBasedGateway(num int) {
	subprocess.EventBasedGateway = make([]gateways.EventBasedGateway, num)
}

/*** Marker ***/

// SetSequenceFlow ...
func (subprocess *SubProcess) SetSequenceFlow(num int) {
	subprocess.SequenceFlow = make([]marker.SequenceFlow, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (subprocess SubProcess) GetID() STR_PTR {
	return &subprocess.ID
}

// GetName ...
func (subprocess SubProcess) GetName() STR_PTR {
	return &subprocess.Name
}

// GetTriggeredByEvent ...
func (subprocess SubProcess) GetTriggeredByEvent() *bool {
	return &subprocess.TriggeredByEvent
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (subprocess SubProcess) GetCamundaAsyncBefore() *bool {
	return &subprocess.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (subprocess SubProcess) GetCamundaAsyncAfter() *bool {
	return &subprocess.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (subprocess SubProcess) GetCamundaJobPriority() *int {
	return &subprocess.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (subprocess SubProcess) GetDocumentation() *attributes.Documentation {
	return &subprocess.Documentation[0]
}

// GetExtensionElements ...
func (subprocess SubProcess) GetExtensionElements() *attributes.ExtensionElements {
	return &subprocess.ExtensionElements[0]
}

// GetIncoming ...
func (subprocess SubProcess) GetIncoming(num int) *marker.Incoming {
	return &subprocess.Incoming[num]
}

// GetOutgoing ...
func (subprocess SubProcess) GetOutgoing(num int) *marker.Outgoing {
	return &subprocess.Outgoing[num]
}

// GetMultiInstaceLoopCharacteristics ...
func (subprocess SubProcess) GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics {
	return &subprocess.MultiInstanceLoopCharacteristics[0]
}

/*** Event ***/

// GetStartEvent ...
func (subprocess SubProcess) GetStartEvent(num int) *elements.StartEvent {
	return &subprocess.StartEvent[num]
}

// GetEndEvent ...
func (subprocess SubProcess) GetEndEvent(num int) *elements.EndEvent {
	return &subprocess.EndEvent[num]
}

/*** Task ***/

// GetTask ...
func (subprocess SubProcess) GetTask(num int) *tasks.Task {
	return &subprocess.Task[num]
}

// GetUserTask ...
func (subprocess SubProcess) GetUserTask(num int) *tasks.UserTask {
	return &subprocess.UserTask[num]
}

// GetManualTask ...
func (subprocess SubProcess) GetManualTask(num int) *tasks.ManualTask {
	return &subprocess.ManualTask[num]
}

// GetReceiveTask ...
func (subprocess SubProcess) GetReceiveTask(num int) *tasks.ReceiveTask {
	return &subprocess.ReceiveTask[num]
}

// GetScriptTask ...
func (subprocess SubProcess) GetScriptTask(num int) *tasks.ScriptTask {
	return &subprocess.ScriptTask[num]
}

// GetSendTask ...
func (subprocess SubProcess) GetSendTask(num int) *tasks.SendTask {
	return &subprocess.SendTask[num]
}

// GetServiceTask ...
func (subprocess SubProcess) GetServiceTask(num int) *tasks.ServiceTask {
	return &subprocess.ServiceTask[num]
}

/*** Subprocess ***/

// GetSubProcess ...
func (subprocess SubProcess) GetSubProcess(num int) *SubProcess {
	return &subprocess.SubProcess[num]
}

// GetAdHocSubProcess ...
func (subprocess SubProcess) GetAdHocSubProcess(num int) *AdHocSubProcess {
	return &subprocess.AdHocSubProcess[num]
}

/*** Gateway ***/

// GetExclusiveGateway
func (subprocess SubProcess) GetExclusiveGateway(num int) *gateways.ExclusiveGateway {
	return &subprocess.ExclusiveGateway[num]
}

// GetInclsuiveGateway
func (subprocess SubProcess) GetInclusiveGateway(num int) *gateways.InclusiveGateway {
	return &subprocess.InclusiveGateway[num]
}

// GetParallelGateway
func (subprocess SubProcess) GetParallelGateway(num int) *gateways.ParallelGateway {
	return &subprocess.ParallelGateway[num]
}

// GetComplexGateway
func (subprocess SubProcess) GetComplexGateway(num int) *gateways.ComplexGateway {
	return &subprocess.ComplexGateway[num]
}

// GetEventBasedGateway
func (subprocess SubProcess) GetEventBasedGateway(num int) *gateways.EventBasedGateway {
	return &subprocess.EventBasedGateway[num]
}

/*** Marker ***/

// GetSequenceFlow ...
func (subprocess SubProcess) GetSequenceFlow(num int) *marker.SequenceFlow {
	return &subprocess.SequenceFlow[num]
}
