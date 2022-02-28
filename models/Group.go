package models

import "fmt"

// Group ...
type Group struct {
	ID                string              `xml:"id,attr,omitempty" json:"id"`
	CategoryValueRef  string              `xml:"categoryValueRef,attr,omitempty" json:"categoryValueRef,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (group *Group) SetID(suffix string) {
	group.ID = "_" + suffix
}

// SetCategoryValRef ...
func (group *Group) SetCategoryValueRef(suffix string) {
	group.CategoryValueRef = fmt.Sprintf("_%s", suffix)
}

/* Elements */

/** BPMN **/

// SetExtensionElements ...
