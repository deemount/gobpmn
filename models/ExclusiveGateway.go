package models

import "fmt"

// ExclusiveGateway ...
type ExclusiveGateway struct {
	ID                string              `xml:"id,attr"`
	Name              string              `xml:"name,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (exclusivegate *ExclusiveGateway) SetID(suffix string) {
	exclusivegate.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (exclusivegate *ExclusiveGateway) SetName(name string) {
	exclusivegate.Name = name
}

/* Elements */

// SetExtensionElements ...
func (exclusivegate *ExclusiveGateway) SetExtensionElements() {
	exclusivegate.ExtensionElements = make([]ExtensionElements, 1)
}
