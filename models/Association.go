package models

import (
	"fmt"
	"log"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
)

// AssociationRepository ...
type AssociationRepository interface {
	SetID(typ string, suffix interface{})
	SetSourceRef(typ string, sourceRef string)
	SetTargetRef(typ string, targetRef string)

	SetDocumentation()
	SetExtensionElements()

	GetID() *string
	GetSourceRef() *string
	GetTargetRef() *string

	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *camunda.ExtensionElements
}

// Association ...
type Association struct {
	ID                string                      `xml:"id,attr"`
	SourceRef         string                      `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string                      `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []attributes.Documentation  `xml:"bpmn:documentation,omitempty"`
	ExtensionElements []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TAssociation ...
type TAssociation struct {
	ID                string                       `xml:"id,attr"`
	SourceRef         string                       `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string                       `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []attributes.Documentation   `xml:"documentation,omitempty"`
	ExtensionElements []camunda.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

func NewAssociation() AssociationRepository {
	return &Association{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (association *Association) SetID(typ string, suffix interface{}) {
	switch typ {
	case "association":
		association.ID = fmt.Sprintf("Association_%v", suffix)
		break
	case "id":
		association.ID = fmt.Sprintf("%s", suffix)
		break
	}
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

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (association *Association) SetDocumentation() {
	association.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (association *Association) SetExtensionElements() {
	association.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (association Association) GetID() *string {
	return &association.ID
}

// GetSourceRef ...
func (association *Association) GetSourceRef() *string {
	return &association.SourceRef
}

// GetTargetRef ...
func (association Association) GetTargetRef() *string {
	return &association.TargetRef
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (association Association) GetDocumentation() *attributes.Documentation {
	return &association.Documentation[0]
}

// GetExtensionElements ...
func (association *Association) GetExtensionElements() *camunda.ExtensionElements {
	return &association.ExtensionElements[0]
}
