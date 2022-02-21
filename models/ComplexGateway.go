package models

import "fmt"

// ComplexGateway ...
type ComplexGateway struct {
	ID                string              `xml:"id,attr"`
	Name              string              `xml:"name,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (complexgate *ComplexGateway) SetID(suffix string) {
	complexgate.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (complexgate *ComplexGateway) SetName(name string) {
	complexgate.Name = name
}

/* Elements */

// SetExtensionElements ...
func (complexgate *ComplexGateway) SetExtensionElements() {
	complexgate.ExtensionElements = make([]ExtensionElements, 1)
}
