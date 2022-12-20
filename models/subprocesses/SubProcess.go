package subprocesses

import (
	"fmt"

	"github.com/deemount/gobpmn/models/activities"
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/events"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/loop"
	"github.com/deemount/gobpmn/models/marker"
)

// SubProcessRepository ...
type SubProcessRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)
	SetTriggeredByEvent(triggered bool)

	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaJobPriority(priority int)

	SetDocumentation()
	SetExtensionElements()
	SetMultiInstanceLoopCharacteristics()

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

	SetSequenceFlow(num int)

	GetID() *string
	GetName() *string
	GetTriggeredByEvent() *bool

	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaJobPriority() *int

	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *camunda.ExtensionElements
	GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics

	GetStartEvent(num int) *events.StartEvent
	GetEndEvent(num int) *events.EndEvent
	GetTask(num int) *activities.Task
	GetUserTask(num int) *activities.UserTask
	GetManualTask(num int) *activities.ManualTask
	GetReceiveTask(num int) *activities.ReceiveTask
	GetScriptTask(num int) *activities.ScriptTask
	GetSendTask(num int) *activities.SendTask
	GetServiceTask(num int) *activities.ServiceTask

	GetSubProcess(num int) *SubProcess
	GetAdHocSubProcess(num int) *AdHocSubProcess

	GetExclusiveGateway(num int) *gateways.ExclusiveGateway
	GetInclusiveGateway(num int) *gateways.InclusiveGateway
	GetParallelGateway(num int) *gateways.ParallelGateway
	GetComplexGateway(num int) *gateways.ComplexGateway
	GetEventBasedGateway(num int) *gateways.EventBasedGateway

	GetSequenceFlow(num int) *marker.SequenceFlow
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
	ExtensionElements                []camunda.ExtensionElements             `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	StartEvent                       []events.StartEvent                     `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty"`
	EndEvent                         []events.EndEvent                       `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	Task                             []activities.Task                       `xml:"bpmn:task,omitempty" json:"task,omitempty"`
	UserTask                         []activities.UserTask                   `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty"`
	ManualTask                       []activities.ManualTask                 `xml:"bpmn:manualTask,omitempty" json:"manualTask,omitempty"`
	ReceiveTask                      []activities.ReceiveTask                `xml:"bpmn:receiveTask,omitempty" json:"receiveTask,omitempty"`
	ScriptTask                       []activities.ScriptTask                 `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty"`
	SendTask                         []activities.SendTask                   `xml:"bpmn:sendTask,omitempty" json:"sendTask,omitempty"`
	ServiceTask                      []activities.ServiceTask                `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty"`
	SubProcess                       []SubProcess                            `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  []AdHocSubProcess                       `xml:"bpmn:adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	ExclusiveGateway                 []gateways.ExclusiveGateway             `xml:"bpmn:exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
	InclusiveGateway                 []gateways.InclusiveGateway             `xml:"bpmn:inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
	ParallelGateway                  []gateways.ParallelGateway              `xml:"bpmn:parallelGateway,omitempty" json:"parallelGateway,omitempty"`
	ComplexGateway                   []gateways.ComplexGateway               `xml:"bpmn:complexGateway,omitempty" json:"complexGateway,omitempty"`
	EventBasedGateway                []gateways.EventBasedGateway            `xml:"bpmn:eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty"`
	SequenceFlow                     []marker.SequenceFlow                   `xml:"bpmn:sequenceFlow,omitempty" json:"equenceFlow,omitempty"`
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
	ExtensionElements                []camunda.TExtensionElements            `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	StartEvent                       []events.TStartEvent                    `xml:"startEvent,omitemnpty" json:"startEvent,omitempty"`
	EndEvent                         []events.TEndEvent                      `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	Task                             []activities.TTask                      `xml:"task,omitempty" json:"task,omitempty"`
	UserTask                         []activities.TUserTask                  `xml:"userTask,omitempty" json:"userTask,omitempty"`
	ManualTask                       []activities.TManualTask                `xml:"manualTask,omitempty" json:"manualTask,omitempty"`
	ReceiveTask                      []activities.ReceiveTask                `xml:"receiveTask,omitempty" json:"receiveTask,omitempty"`
	ScriptTask                       []activities.ScriptTask                 `xml:"scriptTask,omitempty" json:"scriptTask,omitempty"`
	SendTask                         []activities.SendTask                   `xml:"sendTask,omitempty" json:"sendTask,omitempty"`
	ServiceTask                      []activities.ServiceTask                `xml:"serviceTask,omitempty" json:"serviceTask,omitempty"`
	SubProcess                       []TSubProcess                           `xml:"subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  []TAdHocSubProcess                      `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	ExclusiveGateway                 []gateways.ExclusiveGateway             `xml:"exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
	InclusiveGateway                 []gateways.InclusiveGateway             `xml:"inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
	ParallelGateway                  []gateways.ParallelGateway              `xml:"parallelGateway,omitempty" json:"parallelGateway,omitempty"`
	ComplexGateway                   []gateways.ComplexGateway               `xml:"complexGateway,omitempty" json:"complexGateway,omitempty"`
	EventBasedGateway                []gateways.EventBasedGateway            `xml:"eventBasedGateway,omitempty" json:"eventBasedGateway,omitempty"`
	SequenceFlow                     []marker.SequenceFlow                   `xml:"sequenceFlow,omitempty" json:"equenceFlow,omitempty"`
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
	subprocess.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetMultiInstaceLoopCharacteristics ...
func (subprocess *SubProcess) SetMultiInstanceLoopCharacteristics() {
	subprocess.MultiInstanceLoopCharacteristics = make([]loop.MultiInstanceLoopCharacteristics, 1)
}

/*** Event ***/

// SetStartEvent ...
func (subprocess *SubProcess) SetStartEvent(num int) {
	subprocess.StartEvent = make([]events.StartEvent, num)
}

// SetEndEvent ...
func (subprocess *SubProcess) SetEndEvent(num int) {
	subprocess.EndEvent = make([]events.EndEvent, num)
}

/*** Task ***/

// SetTask ...
func (subprocess *SubProcess) SetTask(num int) {
	subprocess.Task = make([]activities.Task, num)
}

// SetUserTask ...
func (subprocess *SubProcess) SetUserTask(num int) {
	subprocess.UserTask = make([]activities.UserTask, num)
}

// SetManualTask ...
func (subprocess *SubProcess) SetManualTask(num int) {
	subprocess.ManualTask = make([]activities.ManualTask, num)
}

// SetReceiveTask ...
func (subprocess *SubProcess) SetReceiveTask(num int) {
	subprocess.ReceiveTask = make([]activities.ReceiveTask, num)
}

// SetScriptTask ...
func (subprocess *SubProcess) SetScriptTask(num int) {
	subprocess.ScriptTask = make([]activities.ScriptTask, num)
}

// SetSendTask ...
func (subprocess *SubProcess) SetSendTask(num int) {
	subprocess.SendTask = make([]activities.SendTask, num)
}

// SetServiceTask ...
func (subprocess *SubProcess) SetServiceTask(num int) {
	subprocess.ServiceTask = make([]activities.ServiceTask, num)
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
func (subprocess *SubProcess) GetID() *string {
	return &subprocess.ID
}

// GetName ...
func (subprocess SubProcess) GetName() *string {
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
func (subprocess SubProcess) GetExtensionElements() *camunda.ExtensionElements {
	return &subprocess.ExtensionElements[0]
}

// GetMultiInstaceLoopCharacteristics ...
func (subprocess SubProcess) GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics {
	return &subprocess.MultiInstanceLoopCharacteristics[0]
}

/*** Event ***/

// GetStartEvent ...
func (subprocess SubProcess) GetStartEvent(num int) *events.StartEvent {
	return &subprocess.StartEvent[num]
}

// GetEndEvent ...
func (subprocess SubProcess) GetEndEvent(num int) *events.EndEvent {
	return &subprocess.EndEvent[num]
}

/*** Task ***/

// GetTask ...
func (subprocess SubProcess) GetTask(num int) *activities.Task {
	return &subprocess.Task[num]
}

// GetUserTask ...
func (subprocess SubProcess) GetUserTask(num int) *activities.UserTask {
	return &subprocess.UserTask[num]
}

// GetManualTask ...
func (subprocess SubProcess) GetManualTask(num int) *activities.ManualTask {
	return &subprocess.ManualTask[num]
}

// GetReceiveTask ...
func (subprocess SubProcess) GetReceiveTask(num int) *activities.ReceiveTask {
	return &subprocess.ReceiveTask[num]
}

// GetScriptTask ...
func (subprocess SubProcess) GetScriptTask(num int) *activities.ScriptTask {
	return &subprocess.ScriptTask[num]
}

// GetSendTask ...
func (subprocess SubProcess) GetSendTask(num int) *activities.SendTask {
	return &subprocess.SendTask[num]
}

// GetServiceTask ...
func (subprocess SubProcess) GetServiceTask(num int) *activities.ServiceTask {
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
