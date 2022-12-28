package attributes

type Attributes struct {
	Documentation     DOCUMENTATION_SLC      `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements EXTENSION_ELEMENTS_SLC `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

type TAttributes struct {
	Documentation     DOCUMENTATION_SLC       `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements TEXTENSION_ELEMENTS_SLC `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}
