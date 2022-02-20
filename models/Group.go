package models

import "fmt"

// Group ...
type Group struct {
	ID             string            `xml:"id,attr,omitempty"`
	CategoryValRef string            `xml:"categoryValueRef,attr,omitempty"`
	ExtensionEl    ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}

// SetID ...
func (group *Group) SetID(suffix string) {
	group.ID = "_" + suffix
}

// SetCategoryValRef ...
func (group *Group) SetCategoryValueRef(suffix string) {
	group.CategoryValRef = fmt.Sprintf("_%s", suffix)
}
