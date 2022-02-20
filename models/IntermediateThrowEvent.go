package models

import "fmt"

// SetIntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	ID                string              `xml:"id,attr,omitempty"`
	Name              string              `xml:"name,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (ite *IntermediateThrowEvent) SetID(suffix string) {
	ite.ID = fmt.Sprintf("Event_%s", suffix)
}

// SetName ...
func (ite *IntermediateThrowEvent) SetName(name string) {
	ite.Name = name
}

/* Elements */

// SetExtensionElements ...
func (ite *IntermediateThrowEvent) SetExtensionElements() {
	ite.ExtensionElements = make([]ExtensionElements, 1)
}
