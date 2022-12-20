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

// AdHocSubProcessRepository ...
type AdHocSubProcessRepository interface {
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

// AdHocSubProcess ...
type AdHocSubProcess struct {
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
type TAdHocSubProcess struct {
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
	SubProcess                       []SubProcess                            `xml:"subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
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
	adhoc.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetMultiInstaceLoopCharacteristics ...
func (adhoc *AdHocSubProcess) SetMultiInstanceLoopCharacteristics() {
	adhoc.MultiInstanceLoopCharacteristics = make([]loop.MultiInstanceLoopCharacteristics, 1)
}

/*** Event ***/

// SetStartEvent ...
func (adhoc *AdHocSubProcess) SetStartEvent(num int) {
	adhoc.StartEvent = make([]events.StartEvent, num)
}

// SetEndEvent ...
func (adhoc *AdHocSubProcess) SetEndEvent(num int) {
	adhoc.EndEvent = make([]events.EndEvent, num)
}

/*** Task ***/

// SetTask ...
func (adhoc *AdHocSubProcess) SetTask(num int) {
	adhoc.Task = make([]activities.Task, num)
}

// SetUserTask ...
func (adhoc *AdHocSubProcess) SetUserTask(num int) {
	adhoc.UserTask = make([]activities.UserTask, num)
}

// SetManualTask ...
func (adhoc *AdHocSubProcess) SetManualTask(num int) {
	adhoc.ManualTask = make([]activities.ManualTask, num)
}

// SetReceiveTask ...
func (adhoc *AdHocSubProcess) SetReceiveTask(num int) {
	adhoc.ReceiveTask = make([]activities.ReceiveTask, num)
}

// SetScriptTask ...
func (adhoc *AdHocSubProcess) SetScriptTask(num int) {
	adhoc.ScriptTask = make([]activities.ScriptTask, num)
}

// SetSendTask ...
func (adhoc *AdHocSubProcess) SetSendTask(num int) {
	adhoc.SendTask = make([]activities.SendTask, num)
}

// SetServiceTask ...
func (adhoc *AdHocSubProcess) SetServiceTask(num int) {
	adhoc.ServiceTask = make([]activities.ServiceTask, num)
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
func (adhoc *AdHocSubProcess) GetID() *string {
	return &adhoc.ID
}

// GetName ...
func (adhoc AdHocSubProcess) GetName() *string {
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

// GetDocumentation ...
func (adhoc AdHocSubProcess) GetDocumentation() *attributes.Documentation {
	return &adhoc.Documentation[0]
}

// GetExtensionElements ...
func (adhoc AdHocSubProcess) GetExtensionElements() *camunda.ExtensionElements {
	return &adhoc.ExtensionElements[0]
}

// GetMultiInstaceLoopCharacteristics ...
func (adhoc AdHocSubProcess) GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics {
	return &adhoc.MultiInstanceLoopCharacteristics[0]
}

/*** Event ***/

// GetStartEvent ...
func (adhoc AdHocSubProcess) GetStartEvent(num int) *events.StartEvent {
	return &adhoc.StartEvent[num]
}

// GetEndEvent ...
func (adhoc AdHocSubProcess) GetEndEvent(num int) *events.EndEvent {
	return &adhoc.EndEvent[num]
}

/*** Task ***/

// GetTask ...
func (adhoc AdHocSubProcess) GetTask(num int) *activities.Task {
	return &adhoc.Task[num]
}

// GetUserTask ...
func (adhoc AdHocSubProcess) GetUserTask(num int) *activities.UserTask {
	return &adhoc.UserTask[num]
}

// GetManualTask ...
func (adhoc AdHocSubProcess) GetManualTask(num int) *activities.ManualTask {
	return &adhoc.ManualTask[num]
}

// GetReceiveTask ...
func (adhoc AdHocSubProcess) GetReceiveTask(num int) *activities.ReceiveTask {
	return &adhoc.ReceiveTask[num]
}

// GetScriptTask ...
func (adhoc AdHocSubProcess) GetScriptTask(num int) *activities.ScriptTask {
	return &adhoc.ScriptTask[num]
}

// GetSendTask ...
func (adhoc AdHocSubProcess) GetSendTask(num int) *activities.SendTask {
	return &adhoc.SendTask[num]
}

// GetServiceTask ...
func (adhoc AdHocSubProcess) GetServiceTask(num int) *activities.ServiceTask {
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
