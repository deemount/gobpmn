package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/marker"
)

func NewReceiveTask() ReceiveTaskRepository {
	return &ReceiveTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (receiveTask *ReceiveTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		receiveTask.ID = fmt.Sprintf("Activity_%s", suffix)
		break
	case "id":
		receiveTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (receiveTask *ReceiveTask) SetName(name string) {
	receiveTask.Name = name
}

// SetMessageRef ...
func (receiveTask *ReceiveTask) SetMessageRef(suffix string) {
	receiveTask.MessageRef = fmt.Sprintf("Message_%s", suffix)
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (receiveTask *ReceiveTask) SetCamundaAsyncBefore(asyncBefore bool) {
	receiveTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (receiveTask *ReceiveTask) SetCamundaAsyncAfter(asyncAfter bool) {
	receiveTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (receiveTask *ReceiveTask) SetCamundaJobPriority(priority int) {
	receiveTask.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetDocumentation ...
func (receiveTask *ReceiveTask) SetDocumentation() {
	receiveTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (receiveTask *ReceiveTask) SetExtensionElements() {
	receiveTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (receiveTask *ReceiveTask) SetIncoming(num int) {
	receiveTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (receiveTask *ReceiveTask) SetOutgoing(num int) {
	receiveTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (receiveTask ReceiveTask) GetID() compulsion.STR_PTR {
	return &receiveTask.ID
}

// GetName ...
func (receiveTask ReceiveTask) GetName() compulsion.STR_PTR {
	return &receiveTask.Name
}

// GetMessageRef ...
func (receiveTask ReceiveTask) GetMessageRef(suffix string) *string {
	return &receiveTask.MessageRef
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (receiveTask ReceiveTask) GetCamundaAsyncBefore() *bool {
	return &receiveTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (receiveTask ReceiveTask) GetCamundaAsyncAfter() *bool {
	return &receiveTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (receiveTask ReceiveTask) GetCamundaJobPriority() *int {
	return &receiveTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (receiveTask ReceiveTask) GetDocumentation() *attributes.Documentation {
	return &receiveTask.Documentation[0]
}

// GetExtensionElements ...
func (receiveTask ReceiveTask) GetExtensionElements() *attributes.ExtensionElements {
	return &receiveTask.ExtensionElements[0]
}

// GetIncoming ...
func (receiveTask ReceiveTask) GetIncoming(num int) *marker.Incoming {
	return &receiveTask.Incoming[num]
}

// GetOutgoing ...
func (receiveTask ReceiveTask) GetOutgoing(num int) *marker.Outgoing {
	return &receiveTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (receiveTask ReceiveTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", receiveTask.ID, receiveTask.Name)
}
