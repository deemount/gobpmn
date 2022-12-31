package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/data"
	"github.com/deemount/gobpmn/models/marker"
)

// NewTask ...
func NewTask() TaskRepository {
	return &Task{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (task *Task) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		task.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		task.ID = fmt.Sprintf("%s", suffix)
		break
	}
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

// SetProperty ...
func (task *Task) SetProperty() {
	task.Property = make([]attributes.Property, 1)
}

// SetDocumentation ...
func (task *Task) SetDocumentation() {
	task.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (task *Task) SetExtensionElements() {
	task.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetDataInputAssociation ...
func (task *Task) SetDataInputAssociation(num int) {
	task.DataInputAssociation = make([]data.DataInputAssociation, num)
}

// SetIncoming ...
func (task *Task) SetIncoming(num int) {
	task.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (task *Task) SetOutgoing(num int) {
	task.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (task Task) GetID() compulsion.STR_PTR {
	return &task.ID
}

// GetName ...
func (task Task) GetName() compulsion.STR_PTR {
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

// GetProperty ...
func (task Task) GetProperty() *attributes.Property {
	return &task.Property[0]
}

// GetDocumentation ...
func (task Task) GetDocumentation() *attributes.Documentation {
	return &task.Documentation[0]
}

// GetExtensionElements ...
func (task Task) GetExtensionElements() *attributes.ExtensionElements {
	return &task.ExtensionElements[0]
}

// GetDataInputAssociation ...
func (task Task) GetDataInputAssociation(num int) *data.DataInputAssociation {
	return &task.DataInputAssociation[num]
}

// GetIncoming ...
func (task Task) GetIncoming(num int) *marker.Incoming {
	return &task.Incoming[num]
}

// GetOutgoing ...
func (task Task) GetOutgoing(num int) *marker.Outgoing {
	return &task.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (task Task) String() string {
	return fmt.Sprintf("id=%v, name=%v", task.ID, task.Name)
}
