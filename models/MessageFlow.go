package models

type MessageFlow struct {
	ID                string              `xml:"id,attr"`
	SourceRef         string              `xml:"sourceRef,attr"`
	TargetRef         string              `xml:"targetRef,attr"`
	Documentation     []Documentation     `xml:"bpmn:documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (mf *MessageFlow) SetID(suffix string) {
	mf.ID = "Flow_" + suffix
}

// SetSourceRef ...
func (mf *MessageFlow) SetSourceRef(sourceRef string) {
	mf.SourceRef = sourceRef
}

// SetTargetRef ...
func (mf *MessageFlow) SetTargetRef(targetRef string) {
	mf.TargetRef = targetRef
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (mf *MessageFlow) SetDocumentation() {
	mf.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (mf *MessageFlow) SetExtensionElements() {
	mf.ExtensionElements = make([]ExtensionElements, 1)
}
