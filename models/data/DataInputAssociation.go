package data

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
)

// DataInputAssociationRepository ...
type DataInputAssociationRepository interface {
	DataBaseID
	DataBaseCoreElements
}

// DataInputAssociation ...
type DataInputAssociation struct {
	ID                string                         `xml:"id,attr,omitempty" json:"id,omitempty"`
	Documentation     []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TDataInputAssociation ...
type TDataInputAssociation struct {
	ID                string                          `xml:"id,attr,omitempty" json:"id,omitempty"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

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
func (dia DataInputAssociation) GetID() STR_PTR {
	return &dia.ID
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (dia DataInputAssociation) GetDocumentation() DOCUMENTATION_PTR {
	return &dia.Documentation[0]
}

// GetExtensionElements ...
func (dia DataInputAssociation) GetExtensionElements() EXTENSION_ELEMENTS_PTR {
	return &dia.ExtensionElements[0]
}
