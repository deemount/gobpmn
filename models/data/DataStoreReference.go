package data

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
)

// DataStoreRepository ...
type DataStoreReferenceRepository interface {
	DataBaseAttributes
	DataBaseCoreElements
}

// DataStoreReference ...
type DataStoreReference struct {
	ID                string                         `xml:"id,attr,omitempty" json:"id" csv:"id"`
	Name              string                         `xml:"name,attr,omitempty" json:"name,omitempty" csv:"name"`
	Documentation     []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TDataStoreReference ...
type TDataStoreReference struct {
	ID                string                          `xml:"id,attr,omitempty" json:"id" csv:"id"`
	Name              string                          `xml:"name,attr,omitempty" json:"name,omitempty" csv:"name"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// NewDataStore ...
func NewDataStoreReference() DataStoreReferenceRepository {
	return &DataStoreReference{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (dsr *DataStoreReference) SetID(typ string, suffix interface{}) {
	switch typ {
	case "dsr":
		dsr.ID = fmt.Sprintf("DataStoreReference__%v", suffix)
		break
	case "id":
		dsr.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (dsr *DataStoreReference) SetName(name string) {
	dsr.Name = name
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (dsr *DataStoreReference) SetDocumentation() {
	dsr.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (dsr *DataStoreReference) SetExtensionElements() {
	dsr.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (dsr DataStoreReference) GetID() STR_PTR {
	return &dsr.ID
}

// GetName ...
func (dsr DataStoreReference) GetName() STR_PTR {
	return &dsr.Name
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (dsr DataStoreReference) GetDocumentation() DOCUMENTATION_PTR {
	return &dsr.Documentation[0]
}

// GetExtensionElements ...
func (dsr DataStoreReference) GetExtensionElements() EXTENSION_ELEMENTS_PTR {
	return &dsr.ExtensionElements[0]
}
