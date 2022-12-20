package models

import (
	"fmt"

	"github.com/deemount/gobpmn/models/camunda"
)

type GroupRepository interface {
	SetID(typ string, suffix string)
	SetCategoryValueRef(suffix string)

	SetDocumentation()
	SetExtensionElements()
}

// Group ...
type Group struct {
	ID                string                      `xml:"id,attr,omitempty" json:"id"`
	CategoryValueRef  string                      `xml:"categoryValueRef,attr,omitempty" json:"categoryValueRef,omitempty"`
	ExtensionElements []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (group *Group) SetID(typ string, suffix string) {
	group.ID = "_" + suffix
}

// SetCategoryValRef ...
func (group *Group) SetCategoryValueRef(suffix string) {
	group.CategoryValueRef = fmt.Sprintf("_%s", suffix)
}

/* Elements */

/** BPMN **/

// SetDocumentation ...

// SetExtensionElements ...
