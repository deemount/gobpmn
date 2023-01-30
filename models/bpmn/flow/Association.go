package flow

import (
	"fmt"
	"log"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewAssociation ...
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
func (association *Association) SetSourceRef(typ string, sourceRef interface{}) {
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
func (association *Association) SetTargetRef(typ string, targetRef interface{}) {
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
	association.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (association Association) GetID() impl.STR_PTR {
	return &association.ID
}

// GetSourceRef ...
func (association Association) GetSourceRef() impl.STR_PTR {
	return &association.SourceRef
}

// GetTargetRef ...
func (association Association) GetTargetRef() impl.STR_PTR {
	return &association.TargetRef
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (association Association) GetDocumentation() *attributes.Documentation {
	return &association.Documentation[0]
}

// GetExtensionElements ...
func (association Association) GetExtensionElements() *attributes.ExtensionElements {
	return &association.ExtensionElements[0]
}
