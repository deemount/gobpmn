package models

import "fmt"

// CallActivity ...
type CallActivity struct {
	ID                               string                             `xml:"id,attr,omitempty"`
	Name                             string                             `xml:"name,attr,omitempty"`
	CalledElement                    string                             `xml:"calledElement,attr,omitempty"`
	CamundaAsyncBefore               bool                               `xml:"camunda:asyncBefore,attr,omitempty"`
	CamundaAsyncAfter                bool                               `xml:"camunda:asyncAfter,attr,omitempty"`
	CamundaJobPriority               int                                `xml:"camunda:jobPriority,attr,omitempty"`
	CamundaCalledElementTenantID     string                             `xml:"camunda:calledElementTenantId,attr,omitempty"`
	CamundaVariableMappingClass      string                             `xml:"camunda:variableMappingClass,attr,omitempty"`
	Documentation                    []Documentation                    `xml:"bpmn:documentation,omitempty"`
	ExtensionElements                []ExtensionElements                `xml:"bpmn:extensionElements,omitempty"`
	Incoming                         []Incoming                         `xml:"bpmn:incoming,omitempty"`
	Outgoing                         []Outgoing                         `xml:"bpmn:outgoing,omitempty"`
	StandardLoopCharacteristics      []StandardLoopCharacteristics      `xml:"bpmn:standardLoopCharacteristics,omitempty"`
	MultiInstanceLoopCharacteristics []MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (ca *CallActivity) SetID(suffix string) {
	ca.ID = fmt.Sprintf("Activity_%s", suffix)
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

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (ca *CallActivity) SetDocumentation() {
	ca.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (ca *CallActivity) SetExtensionElements() {
	ca.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (ca *CallActivity) SetIncoming(num int) {
	ca.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (ca *CallActivity) SetOutgoing(num int) {
	ca.Outgoing = make([]Outgoing, num)
}

// SetStandardLoopCharacteristics ...
func (ca *CallActivity) SetStandardLoopCharacteristics() {
	ca.StandardLoopCharacteristics = make([]StandardLoopCharacteristics, 1)
}

// SetMultiInstanceLoopCharacteristics ...
func (ca *CallActivity) SetMultiInstanceLoopCharacteristics() {
	ca.MultiInstanceLoopCharacteristics = make([]MultiInstanceLoopCharacteristics, 1)
}
