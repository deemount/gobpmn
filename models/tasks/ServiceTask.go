package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

func NewServiceTask() ServiceTaskRepository {
	return &ServiceTask{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (serviceTask *ServiceTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		serviceTask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		serviceTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (serviceTask *ServiceTask) SetName(name string) {
	serviceTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (serviceTask *ServiceTask) SetCamundaAsyncBefore(asyncBefore bool) {
	serviceTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (serviceTask *ServiceTask) SetCamundaAsyncAfter(asyncAfter bool) {
	serviceTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (serviceTask *ServiceTask) SetCamundaJobPriority(priority int) {
	serviceTask.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetDocumentation ...
func (serviceTask *ServiceTask) SetDocumentation() {
	serviceTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (serviceTask *ServiceTask) SetExtensionElements() {
	serviceTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (serviceTask *ServiceTask) SetIncoming(num int) {
	serviceTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (serviceTask *ServiceTask) SetOutgoing(num int) {
	serviceTask.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (serviceTask ServiceTask) GetID() STR_PTR {
	return &serviceTask.ID
}

// GetName ...
func (serviceTask ServiceTask) GetName() STR_PTR {
	return &serviceTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (serviceTask ServiceTask) GetCamundaAsyncBefore() *bool {
	return &serviceTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (serviceTask ServiceTask) GetCamundaAsyncAfter() *bool {
	return &serviceTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (serviceTask ServiceTask) GetCamundaJobPriority() *int {
	return &serviceTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (serviceTask ServiceTask) GetDocumentation() *attributes.Documentation {
	return &serviceTask.Documentation[0]
}

// GetExtensionElements ...
func (serviceTask ServiceTask) GetExtensionElements() *attributes.ExtensionElements {
	return &serviceTask.ExtensionElements[0]
}

// GetIncoming ...
func (serviceTask ServiceTask) GetIncoming(num int) *marker.Incoming {
	return &serviceTask.Incoming[num]
}

// GetOutgoing ...
func (serviceTask ServiceTask) GetOutgoing(num int) *marker.Outgoing {
	return &serviceTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (serviceTask ServiceTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", serviceTask.ID, serviceTask.Name)
}
