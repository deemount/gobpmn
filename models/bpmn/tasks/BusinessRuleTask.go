package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

func NewBusinessRuleTask() BusinessRuleTaskRepository {
	return &BusinessRuleTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (businessRuleTask *BusinessRuleTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		businessRuleTask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		businessRuleTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (businessRuleTask *BusinessRuleTask) SetName(name string) {
	businessRuleTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (businessRuleTask *BusinessRuleTask) SetCamundaAsyncBefore(asyncBefore bool) {
	businessRuleTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (businessRuleTask *BusinessRuleTask) SetCamundaAsyncAfter(asyncAfter bool) {
	businessRuleTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (businessRuleTask *BusinessRuleTask) SetCamundaJobPriority(priority int) {
	businessRuleTask.CamundaJobPriority = priority
}

// SetCamundaClass ...
func (businessRuleTask *BusinessRuleTask) SetCamundaClass(class string) {
	businessRuleTask.CamundaClass = class
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (businessRuleTask *BusinessRuleTask) SetDocumentation() {
	businessRuleTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (businessRuleTask *BusinessRuleTask) SetExtensionElements() {
	businessRuleTask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (businessRuleTask *BusinessRuleTask) SetIncoming(num int) {
	businessRuleTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (businessRuleTask *BusinessRuleTask) SetOutgoing(num int) {
	businessRuleTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (businessRuleTask BusinessRuleTask) GetID() impl.STR_PTR {
	return &businessRuleTask.ID
}

// GetName ...
func (businessRuleTask BusinessRuleTask) GetName() impl.STR_PTR {
	return &businessRuleTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (businessRuleTask BusinessRuleTask) GetCamundaAsyncBefore() *bool {
	return &businessRuleTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (businessRuleTask BusinessRuleTask) GetCamundaAsyncAfter() *bool {
	return &businessRuleTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (businessRuleTask BusinessRuleTask) GetCamundaJobPriority() *int {
	return &businessRuleTask.CamundaJobPriority
}

// GetCamundaClass ...
func (businessRuleTask BusinessRuleTask) GetCamundaClass() *string {
	return &businessRuleTask.CamundaClass
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (businessRuleTask BusinessRuleTask) GetDocumentation() *attributes.Documentation {
	return &businessRuleTask.Documentation[0]
}

// SetExtensionElements ...
func (businessRuleTask BusinessRuleTask) GetExtensionElements() *attributes.ExtensionElements {
	return &businessRuleTask.ExtensionElements[0]
}

// SetIncoming ...
func (businessRuleTask BusinessRuleTask) GetIncoming(num int) *marker.Incoming {
	return &businessRuleTask.Incoming[num]
}

// SetOutgoing ...
func (businessRuleTask BusinessRuleTask) GetOutgoing(num int) *marker.Outgoing {
	return &businessRuleTask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (businessRuleTask BusinessRuleTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", businessRuleTask.ID, businessRuleTask.Name)
}
