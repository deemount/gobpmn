package models

import "fmt"

type MessageFlow struct {
	ID                string              `xml:"id,attr" json:"id"`
	SourceRef         string              `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string              `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (messageFlow *MessageFlow) SetID(suffix string) {
	messageFlow.ID = fmt.Sprintf("Flow_%s", suffix)
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
