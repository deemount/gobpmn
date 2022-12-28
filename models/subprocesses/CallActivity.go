package subprocesses

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/loop"
	"github.com/deemount/gobpmn/models/marker"
)

func NewCallActivity() CallActivityRepository {
	return &CallActivity{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ca *CallActivity) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		ca.ID = fmt.Sprintf("Activity_%s", suffix)
		break
	case "id":
		ca.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (ca *CallActivity) SetName(name string) {
	ca.Name = name
}

// SetCalledElement ...
func (ca *CallActivity) SetCalledElement(element string) {
	ca.CalledElement = element
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (ca *CallActivity) SetCamundaAsyncBefore(asyncBefore bool) {
	ca.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (ca *CallActivity) SetCamundaAsyncAfter(asyncAfter bool) {
	ca.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (ca *CallActivity) SetCamundaJobPriority(priority int) {
	ca.CamundaJobPriority = priority
}

// SetCamundaCalledElementTenantID ...
func (ca *CallActivity) SetCamundaCalledElementTenantID(tenantID string) {
	ca.CamundaCalledElementTenantID = tenantID
}

// SetCamundaVariableMappingClass ...
func (ca *CallActivity) SetCamundaVariableMappingClass(class string) {
	ca.CamundaVariableMappingClass = class
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (ca *CallActivity) SetDocumentation() {
	ca.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (ca *CallActivity) SetExtensionElements() {
	ca.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (ca *CallActivity) SetIncoming(num int) {
	ca.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (ca *CallActivity) SetOutgoing(num int) {
	ca.Outgoing = make([]marker.Outgoing, num)
}

/*** Loop ***/

// SetStandardLoopCharacteristics ...
func (ca *CallActivity) SetStandardLoopCharacteristics() {
	ca.StandardLoopCharacteristics = make([]loop.StandardLoopCharacteristics, 1)
}

// SetMultiInstanceLoopCharacteristics ...
func (ca *CallActivity) SetMultiInstanceLoopCharacteristics() {
	ca.MultiInstanceLoopCharacteristics = make([]loop.MultiInstanceLoopCharacteristics, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ca CallActivity) GetID() STR_PTR {
	return &ca.ID
}

// GetName ...
func (ca CallActivity) GetName() STR_PTR {
	return &ca.Name
}

// GetCalledElement ...
func (ca CallActivity) GetCalledElement() *string {
	return &ca.CalledElement
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (ca CallActivity) GetCamundaAsyncBefore() *bool {
	return &ca.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (ca CallActivity) GetCamundaAsyncAfter() *bool {
	return &ca.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (ca CallActivity) GetCamundaJobPriority() *int {
	return &ca.CamundaJobPriority
}

// GetCamundaCalledElementTenantID ...
func (ca CallActivity) GetCamundaCalledElementTenantID() *string {
	return &ca.CamundaCalledElementTenantID
}

// GetCamundaVariableMappingClass ...
func (ca CallActivity) GetCamundaVariableMappingClass() *string {
	return &ca.CamundaVariableMappingClass
}

/*** Make Elements ***/

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (ca CallActivity) GetDocumentation() *attributes.Documentation {
	return &ca.Documentation[0]
}

// GetExtensionElements ...
func (ca CallActivity) GetExtensionElements() *attributes.ExtensionElements {
	return &ca.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (ca CallActivity) GetIncoming(num int) *marker.Incoming {
	return &ca.Incoming[num]
}

// GetOutgoing ...
func (ca CallActivity) GetOutgoing(num int) *marker.Outgoing {
	return &ca.Outgoing[num]
}

/*** Loop ***/

// GetStandardLoopCharacteristics ...
func (ca CallActivity) GetStandardLoopCharacteristics() *loop.StandardLoopCharacteristics {
	return &ca.StandardLoopCharacteristics[0]
}

// GetMultiInstanceLoopCharacteristics ...
func (ca CallActivity) GetMultiInstanceLoopCharacteristics() *loop.MultiInstanceLoopCharacteristics {
	return &ca.MultiInstanceLoopCharacteristics[0]
}
