package activities

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// TaskRepository ...
type TaskRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)

	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaJobPriority(priority int)

	SetDocumentation()
	SetExtensionElements()
	SetIncoming(num int)
	SetOutgoing(num int)

	String() string

	GetID() *string
	GetName() *string

	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaJobPriority() *int

	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *camunda.ExtensionElements
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing
}

// Task ...
type Task struct {
	ID                 string                      `xml:"id,attr,omitempty" json:"id" csv:"id"`
	Name               string                      `xml:"name,attr,omitempty" json:"name,omitempty" csv:"name"`
	CamundaAsyncBefore bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing           `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TTask ...
type TTask struct {
	ID                string                       `xml:"id,attr,omitempty" json:"id"`
	Name              string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore       bool                         `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter        bool                         `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority       int                          `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation     []attributes.Documentation   `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming            `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing            `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewTask() TaskRepository {
	return &Task{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (task *Task) SetID(typ string, suffix interface{}) {
	task.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (task *Task) SetName(name string) {
	task.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (task *Task) SetCamundaAsyncBefore(asyncBefore bool) {
	task.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (task *Task) SetCamundaAsyncAfter(asyncAfter bool) {
	task.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (task *Task) SetCamundaJobPriority(priority int) {
	task.CamundaJobPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (task *Task) SetDocumentation() {
	task.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (task *Task) SetExtensionElements() {
	task.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (task *Task) SetIncoming(num int) {
	task.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (task *Task) SetOutgoing(num int) {
	task.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default String
 */

// String ...
func (task Task) String() string {
	return fmt.Sprintf("id=%v, name=%v", task.ID, task.Name)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (task Task) GetID() *string {
	return &task.ID
}

// GetName ...
func (task Task) GetName() *string {
	return &task.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (task Task) GetCamundaAsyncBefore() *bool {
	return &task.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (task Task) GetCamundaAsyncAfter() *bool {
	return &task.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (task Task) GetCamundaJobPriority() *int {
	return &task.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (task Task) GetDocumentation() *attributes.Documentation {
	return &task.Documentation[0]
}

// GetExtensionElements ...
func (task Task) GetExtensionElements() *camunda.ExtensionElements {
	return &task.ExtensionElements[0]
}

// GetIncoming ...
func (task Task) GetIncoming(num int) *marker.Incoming {
	return &task.Incoming[num]
}

// GetOutgoing ...
func (task Task) GetOutgoing(num int) *marker.Outgoing {
	return &task.Outgoing[num]
}
