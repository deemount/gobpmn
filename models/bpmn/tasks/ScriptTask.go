package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

func NewScriptTask() ScriptTaskRepository {
	return &ScriptTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (scriptTask *ScriptTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		scriptTask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		scriptTask.ID = fmt.Sprintf("Activity_%s", suffix)
		break
	}
}

// SetName ...
func (scriptTask *ScriptTask) SetName(name string) {
	scriptTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (scriptTask *ScriptTask) SetCamundaAsyncBefore(asyncBefore bool) {
	scriptTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (scriptTask *ScriptTask) SetCamundaAsyncAfter(asyncAfter bool) {
	scriptTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (scriptTask *ScriptTask) SetCamundaJobPriority(priority int) {
	scriptTask.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetDocumentation ...
func (scriptTask *ScriptTask) SetDocumentation() {
	scriptTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (scriptTask *ScriptTask) SetExtensionElements() {
	scriptTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (scriptTask *ScriptTask) SetIncoming(num int) {
	scriptTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (scriptTask *ScriptTask) SetOutgoing(num int) {
	scriptTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (scriptTask ScriptTask) GetID() impl.STR_PTR {
	return &scriptTask.ID
}

// GetName ...
func (scriptTask ScriptTask) GetName() impl.STR_PTR {
	return &scriptTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (scriptTask ScriptTask) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &scriptTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (scriptTask ScriptTask) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &scriptTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (scriptTask ScriptTask) GetCamundaJobPriority() impl.INT_PTR {
	return &scriptTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (scriptTask ScriptTask) GetDocumentation() *attributes.Documentation {
	return &scriptTask.Documentation[0]
}

// GetExtensionElements ...
func (scriptTask ScriptTask) GetExtensionElements() *attributes.ExtensionElements {
	return &scriptTask.ExtensionElements[0]
}

// GetIncoming ...
func (scriptTask ScriptTask) GetIncoming(num int) *marker.Incoming {
	return &scriptTask.Incoming[num]
}

// GetOutgoing ...
func (scriptTask ScriptTask) GetOutgoing(num int) *marker.Outgoing {
	return &scriptTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (scriptTask ScriptTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", scriptTask.ID, scriptTask.Name)
}
