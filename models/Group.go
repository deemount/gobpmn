package models

// Group ...
type Group struct {
	ID             string            `xml:"id,attr,omitempty"`
	CategoryValRef string            `xml:"categoryValueRef,attr,omitempty"`
	ExtensionEl    ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}
