package models

import "fmt"

// MessageFlow ...
type MessageFlow struct {
	ID                string              `xml:"id,attr" json:"id"`
	Name              string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	SourceRef         string              `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string              `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TMessageFlow ...
type TMessageFlow struct {
	ID                string              `xml:"id,attr" json:"id"`
	Name              string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	SourceRef         string              `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string              `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []Documentation     `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (messageFlow *MessageFlow) SetID(typ string, suffix string) {
	switch typ {
	case "flow":
		messageFlow.ID = fmt.Sprintf("Flow_%s", suffix)
		break
	case "id":
		messageFlow.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (messageFlow *MessageFlow) SetName(name string) {
	messageFlow.Name = name
}

// SetSourceRef ...
func (messageFlow *MessageFlow) SetSourceRef(sourceRef string) {
	messageFlow.SourceRef = sourceRef
}

// SetTargetRef ...
func (messageFlow *MessageFlow) SetTargetRef(targetRef string) {
	messageFlow.TargetRef = targetRef
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (messageFlow *MessageFlow) SetDocumentation() {
	messageFlow.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (messageFlow *MessageFlow) SetExtensionElements() {
	messageFlow.ExtensionElements = make([]ExtensionElements, 1)
}
