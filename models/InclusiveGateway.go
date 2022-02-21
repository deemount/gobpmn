package models

import "fmt"

// InclusiveGateway ...
type InclusiveGateway struct {
	ID                string              `xml:"id,attr"`
	Name              string              `xml:"name,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusivegate *InclusiveGateway) SetID(suffix string) {
	inclusivegate.ID = fmt.Sprintf("_%s", suffix)
}

// SetName ...
func (inclusivegate *InclusiveGateway) SetName(name string) {
	inclusivegate.Name = name
}

/* Elements */

// SetExtensionElements ...
func (inclusivegate *InclusiveGateway) SetExtensionElements() {
	inclusivegate.ExtensionElements = make([]ExtensionElements, 1)
}
