package marker

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
)

// CategoryRepository ...
type CategoryRepository interface {
	MarkerID
	MarkerBaseCoreElements

	SetCategoryValue()
	GetCategoryValue() *CategoryValue
}

// Category ...
type Category struct {
	ID                string                         `xml:"id,attr,omitempty" json:"id"`
	CategoryValue     []CategoryValue                `xml:"bpmn:categoryValue,omitempty" json:"categoryValue,omitempty"`
	Documentation     []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// Category ...
type TCategory struct {
	ID                string                          `xml:"id,attr,omitempty" json:"id"`
	CategoryValue     []CategoryValue                 `xml:"categoryValue,omitempty" json:"categoryValue,omitempty"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

func NewCategory() CategoryRepository {
	return &Category{}
}

/* Attributes */

/** BPMN **/

// SetID ...
func (category *Category) SetID(typ string, suffix interface{}) {
	switch typ {
	case "category":
		category.ID = fmt.Sprintf("Category_%s", suffix)
		break
	case "id":
		category.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/*** Make Elements ***/

/** BPMN **/

// SetCategoryValue...
func (category *Category) SetCategoryValue() {
	category.CategoryValue = make([]CategoryValue, 1)
}

// SetDocumentation ...
func (category *Category) SetDocumentation() {
	category.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (category *Category) SetExtensionElements() {
	category.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (category Category) GetID() STR_PTR {
	return &category.ID
}

/* Elements */

/** BPMN **/

// SetCategoryValue...
func (category Category) GetCategoryValue() *CategoryValue {
	return &category.CategoryValue[0]
}

// GetDocumentation ...
func (category Category) GetDocumentation() DOCUMENTATION_PTR {
	return &category.Documentation[0]
}

// GetExtensionElements ...
func (category Category) GetExtensionElements() EXTENSION_ELEMENTS_PTR {
	return &category.ExtensionElements[0]
}
