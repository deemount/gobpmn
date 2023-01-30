package data

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewDataStoreReference ...
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
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (dsr DataStoreReference) GetID() impl.STR_PTR {
	return &dsr.ID
}

// GetName ...
func (dsr DataStoreReference) GetName() impl.STR_PTR {
	return &dsr.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (dsr DataStoreReference) GetDocumentation() *attributes.Documentation {
	return &dsr.Documentation[0]
}

// GetExtensionElements ...
func (dsr DataStoreReference) GetExtensionElements() *attributes.ExtensionElements {
	return &dsr.ExtensionElements[0]
}
