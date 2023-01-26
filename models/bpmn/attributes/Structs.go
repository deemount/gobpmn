package attributes

import (
	"github.com/deemount/gobpmn/models/bpmn/camunda"
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

type Attributes struct {
	Documentation     DOCUMENTATION_SLC      `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements EXTENSION_ELEMENTS_SLC `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

type TAttributes struct {
	Documentation     DOCUMENTATION_SLC       `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements TEXTENSION_ELEMENTS_SLC `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// Documentation ...
type Documentation struct {
	Data string `xml:",innerxml,omitempty" json:"documentation,omitempty"`
}

// ExtensionElements ...
type ExtensionElements struct {
	camunda.ExtensionElements
}

// TExtensionElements ...
type TExtensionElements struct {
	camunda.ExtensionElements
}

// Property ...
type Property struct {
	impl.BaseAttributes
}
