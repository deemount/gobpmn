package models

import (
	"fmt"
	"log"
)

// Association ...
type Association struct {
	ID                string              `xml:"id,attr"`
	SourceRef         string              `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string              `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []Documentation     `xml:"bpmn:documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TAssociation ...
type TAssociation struct {
	ID                string              `xml:"id,attr"`
	SourceRef         string              `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string              `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []Documentation     `xml:"bpmn:documentation,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (association *Association) SetID(suffix string) {
	association.ID = fmt.Sprintf("Association_%s", suffix)
}

// SetSourceRef ...
func (association *Association) SetSourceRef(typ string, sourceRef string) {
	switch typ {
	case "activity":
		association.SourceRef = fmt.Sprintf("Activity_%s", sourceRef)
		break
	case "event":
		association.SourceRef = fmt.Sprintf("Event_%s", sourceRef)
		break
	case "id":
		association.SourceRef = fmt.Sprintf("%s", sourceRef)
		break
	case "participant":
		association.SourceRef = fmt.Sprintf("Participant_%s", sourceRef)
		break
	case "textannotation":
		association.SourceRef = fmt.Sprintf("TextAnnotation_%s", sourceRef)
		break
	}
}

// SetTargetRef ...
func (association *Association) SetTargetRef(typ string, targetRef string) {
	switch typ {
	case "activity":
		association.TargetRef = fmt.Sprintf("Activity_%s", targetRef)
		break
	case "event":
		association.TargetRef = fmt.Sprintf("Event_%s", targetRef)
		break
	case "id":
		association.TargetRef = fmt.Sprintf("%s", targetRef)
		break
	case "participant":
		association.TargetRef = fmt.Sprintf("Participant_%s", targetRef)
		break
	case "textannotation":
		association.TargetRef = fmt.Sprintf("TextAnnotation_%s", targetRef)
		break
	default:
		log.Panic("no typ set in target ref for message flow")
	}
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (association *Association) SetDocumentation() {
	association.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (association *Association) SetExtensionElements() {
	association.ExtensionElements = make([]ExtensionElements, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (association Association) GetID() string {
	return association.ID
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (association Association) GetDocumentation() *Documentation {
	return &association.Documentation[0]
}
