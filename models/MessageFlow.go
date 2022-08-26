package models

import (
	"fmt"
	"log"
)

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

/**
 * Default Setters
 */

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
func (messageFlow *MessageFlow) SetSourceRef(typ string, sourceRef string) {
	switch typ {
	case "activity":
		messageFlow.SourceRef = fmt.Sprintf("Activity_%s", sourceRef)
		break
	case "event":
		messageFlow.SourceRef = fmt.Sprintf("Event_%s", sourceRef)
		break
	case "id":
		messageFlow.SourceRef = fmt.Sprintf("%s", sourceRef)
		break
	case "participant":
		messageFlow.SourceRef = fmt.Sprintf("Participant_%s", sourceRef)
		break
	}
}

// SetTargetRef ...
func (messageFlow *MessageFlow) SetTargetRef(typ string, targetRef string) {
	switch typ {
	case "activity":
		messageFlow.TargetRef = fmt.Sprintf("Activity_%s", targetRef)
		break
	case "event":
		messageFlow.TargetRef = fmt.Sprintf("Event_%s", targetRef)
		break
	case "id":
		messageFlow.TargetRef = fmt.Sprintf("%s", targetRef)
		break
	case "participant":
		messageFlow.TargetRef = fmt.Sprintf("Participant_%s", targetRef)
		break
	default:
		log.Panic("no typ set in target ref for message flow")
	}
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

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (messageFlow MessageFlow) GetID() string {
	return messageFlow.ID
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (messageFlow MessageFlow) GetDocumentation() *Documentation {
	return &messageFlow.Documentation[0]
}
