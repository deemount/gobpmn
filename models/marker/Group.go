package marker

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
)

// NewGroup ...
func NewGroup() GroupRepository {
	return &Group{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (group *Group) SetID(typ string, suffix interface{}) {
	switch typ {
	case "group":
		group.ID = fmt.Sprintf("Group_%v", suffix)
		break
	case "id":
		group.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetCategoryValRef ...
func (group *Group) SetCategoryValueRef(suffix string) {
	group.CategoryValueRef = fmt.Sprintf("_%s", suffix)
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (group *Group) SetDocumentation() {
	group.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (group *Group) SetExtensionElements() {
	group.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (group Group) GetID() STR_PTR {
	return &group.ID

}

// GetCategoryValRef ...
func (group Group) GetCategoryValueRef() *string {
	return &group.CategoryValueRef
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (group Group) GetDocumentation() *attributes.Documentation {
	return &group.Documentation[0]
}

// GetExtensionElements ...
func (group Group) GetExtensionElements() *attributes.ExtensionElements {
	return &group.ExtensionElements[0]
}
