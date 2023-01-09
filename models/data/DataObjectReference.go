package data

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
)

// NewDataObjectReference ...
func NewDataObjectReference() DataObjectReferenceRepository {
	return &DataObjectReference{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (dor *DataObjectReference) SetID(typ string, suffix interface{}) {
	switch typ {
	case "dataobjref":
		dor.ID = fmt.Sprintf("DataObjecteReference_%v", suffix)
		break
	case "id":
		dor.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (dor *DataObjectReference) SetName(name string) {
	dor.Name = name
}

// SetDataObjectRef ...
func (dor *DataObjectReference) SetDataObjectRef(suffix interface{}) {
	dor.Name = fmt.Sprintf("DataObject_%v", suffix)
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (dor *DataObjectReference) SetDocumentation() {
	dor.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (dor *DataObjectReference) SetExtensionElements() {
	dor.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (dor DataObjectReference) GetID() compulsion.STR_PTR {
	return &dor.ID
}

// SetName ...
func (dor DataObjectReference) GetName() compulsion.STR_PTR {
	return &dor.Name
}

// SetDataObjectRef ...
func (dor DataObjectReference) GetDataObjectRef() *string {
	return &dor.Name
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetDocumentation ...
func (dor DataObjectReference) GetDocumentation() *attributes.Documentation {
	return &dor.Documentation[0]
}

// GetExtensionElements ...
func (dor DataObjectReference) GetExtensionElements() *attributes.ExtensionElements {
	return &dor.ExtensionElements[0]
}
