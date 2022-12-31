package data

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
)

// NewDataInputAssociation ...
func NewDataInputAssociation() DataInputAssociationRepository {
	return &DataInputAssociation{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (dia *DataInputAssociation) SetID(typ string, suffix interface{}) {
	switch typ {
	case "dia":
		dia.ID = fmt.Sprintf("DataInputAssociation_%v", suffix)
		break
	case "id":
		dia.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (dia *DataInputAssociation) SetDocumentation() {
	dia.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (dia *DataInputAssociation) SetExtensionElements() {
	dia.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (dia DataInputAssociation) GetID() compulsion.STR_PTR {
	return &dia.ID
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (dia DataInputAssociation) GetDocumentation() *attributes.Documentation {
	return &dia.Documentation[0]
}

// GetExtensionElements ...
func (dia DataInputAssociation) GetExtensionElements() *attributes.ExtensionElements {
	return &dia.ExtensionElements[0]
}
