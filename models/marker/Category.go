package marker

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
)

// CategoryRepository ...
type CategoryRepository interface {
	SetID(typ string, suffix interface{})

	SetDocumentation()
	SetExtensionElements()

	SetCategoryValue()

	GetID() *string
}

// Category ...
type Category struct {
	ID                string                      `xml:"id,attr,omitempty" json:"id"`
	CategoryValue     []CategoryValue             `xml:"bpmn:categoryValue,omitempty" json:"categoryValue,omitempty"`
	Documentation     []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// Category ...
type TCategory struct {
	ID                string                       `xml:"id,attr,omitempty" json:"id"`
	CategoryValue     []CategoryValue              `xml:"categoryValue,omitempty" json:"categoryValue,omitempty"`
	Documentation     []attributes.Documentation   `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.TExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
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
	category.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (category Category) GetID() *string {
	return &category.ID
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (category Category) GetDocumentation() *attributes.Documentation {
	return &category.Documentation[0]
}

// GetExtensionElements ...
func (category Category) GetExtensionElements() *camunda.ExtensionElements {
	return &category.ExtensionElements[0]
}
