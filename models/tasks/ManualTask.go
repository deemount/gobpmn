package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

// ManualTaskRepository ...
type ManualTaskRepository interface {
	TasksBase

	SetDocumentation()
	SetExtensionElements()
	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *attributes.ExtensionElements
}

// ManualTask ...
type ManualTask struct {
	ID                 string                         `xml:"id,attr,omitempty" json:"id"`
	Name               string                         `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                           `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                           `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                            `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming              `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing              `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TManualTask ...
type TManualTask struct {
	ID                string                          `xml:"id,attr,omitempty" json:"id"`
	Name              string                          `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore       bool                            `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter        bool                            `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority       int                             `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming               `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing               `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewManualTask() ManualTaskRepository {
	return &ManualTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (manualTask *ManualTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		manualTask.ID = fmt.Sprintf("Activity_%s", suffix)
		break
	case "id":
		manualTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (manualTask *ManualTask) SetName(name string) {
	manualTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (manualTask *ManualTask) SetCamundaAsyncBefore(asyncBefore bool) {
	manualTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (manualTask *ManualTask) SetCamundaAsyncAfter(asyncAfter bool) {
	manualTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (manualTask *ManualTask) SetCamundaJobPriority(priority int) {
	manualTask.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetDocumentation ...
func (manualTask *ManualTask) SetDocumentation() {
	manualTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (manualTask *ManualTask) SetExtensionElements() {
	manualTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (manualTask *ManualTask) SetIncoming(num int) {
	manualTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (manualTask *ManualTask) SetOutgoing(num int) {
	manualTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (manualTask ManualTask) GetID() STR_PTR {
	return &manualTask.ID
}

// GetName ...
func (manualTask ManualTask) GetName() STR_PTR {
	return &manualTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (manualTask ManualTask) GetCamundaAsyncBefore() *bool {
	return &manualTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (manualTask ManualTask) GetCamundaAsyncAfter() *bool {
	return &manualTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (manualTask ManualTask) GetCamundaJobPriority() *int {
	return &manualTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (manualTask ManualTask) GetDocumentation() *attributes.Documentation {
	return &manualTask.Documentation[0]
}

// GetExtensionElements ...
func (manualTask ManualTask) GetExtensionElements() *attributes.ExtensionElements {
	return &manualTask.ExtensionElements[0]
}

// GetIncoming ...
func (manualTask ManualTask) GetIncoming(num int) *marker.Incoming {
	return &manualTask.Incoming[num]
}

// GetOutgoing ...
func (manualTask ManualTask) GetOutgoing(num int) *marker.Outgoing {
	return &manualTask.Outgoing[num]
}
