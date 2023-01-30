package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

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
func (manualTask ManualTask) GetID() impl.STR_PTR {
	return &manualTask.ID
}

// GetName ...
func (manualTask ManualTask) GetName() impl.STR_PTR {
	return &manualTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (manualTask ManualTask) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &manualTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (manualTask ManualTask) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &manualTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (manualTask ManualTask) GetCamundaJobPriority() impl.INT_PTR {
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

/*
 * Default String
 */

// String ...
func (manualTask ManualTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", manualTask.ID, manualTask.Name)
}
