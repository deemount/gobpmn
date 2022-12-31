package data

import "github.com/deemount/gobpmn/models/attributes"

// DataInputAssociation ...
type DataInputAssociation struct {
	ID                string                         `xml:"id,attr,omitempty" json:"id,omitempty"`
	Documentation     []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TDataInputAssociation ...
type TDataInputAssociation struct {
	ID                string                          `xml:"id,attr,omitempty" json:"id,omitempty"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// DataObject ...
type DataObject struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// DataObjectReference ...
type DataObjectReference struct{}

// DataStoreReference ...
type DataStoreReference struct {
	ID                string                         `xml:"id,attr,omitempty" json:"id" csv:"id"`
	Name              string                         `xml:"name,attr,omitempty" json:"name,omitempty" csv:"name"`
	Documentation     []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TDataStoreReference ...
type TDataStoreReference struct {
	ID                string                          `xml:"id,attr,omitempty" json:"id" csv:"id"`
	Name              string                          `xml:"name,attr,omitempty" json:"name,omitempty" csv:"name"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}
