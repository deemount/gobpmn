package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

func NewSendTask() SendTaskRepository {
	return &SendTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (sendTask *SendTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		sendTask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		sendTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (sendTask *SendTask) SetName(name string) {
	sendTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (sendTask *SendTask) SetCamundaAsyncBefore(asyncBefore bool) {
	sendTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (sendTask *SendTask) SetCamundaAsyncAfter(asyncAfter bool) {
	sendTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (sendTask *SendTask) SetCamundaJobPriority(priority int) {
	sendTask.CamundaJobPriority = priority
}

/* Elements */

// SetDocumentation ...
func (sendTask *SendTask) SetDocumentation() {
	sendTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (sendTask *SendTask) SetExtensionElements() {
	sendTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (sendTask *SendTask) SetIncoming(num int) {
	sendTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (sendTask *SendTask) SetOutgoing(num int) {
	sendTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (sendTask SendTask) GetID() STR_PTR {
	return &sendTask.ID
}

// GetName ...
func (sendTask SendTask) GetName() STR_PTR {
	return &sendTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (sendTask SendTask) GetCamundaAsyncBefore() *bool {
	return &sendTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (sendTask SendTask) GetCamundaAsyncAfter() *bool {
	return &sendTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (sendTask SendTask) GetCamundaJobPriority() *int {
	return &sendTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (sendTask SendTask) GetDocumentation() *attributes.Documentation {
	return &sendTask.Documentation[0]
}

// GetExtensionElements ...
func (sendTask SendTask) GetExtensionElements() *attributes.ExtensionElements {
	return &sendTask.ExtensionElements[0]
}

// GetIncoming ...
func (sendTask SendTask) GetIncoming(num int) *marker.Incoming {
	return &sendTask.Incoming[num]
}

// GetOutgoing ...
func (sendTask SendTask) GetOutgoing(num int) *marker.Outgoing {
	return &sendTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (sendTask SendTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", sendTask.ID, sendTask.Name)
}
