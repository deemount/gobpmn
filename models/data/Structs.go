package data

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
)

// DataInputAssociation ...
type DataInputAssociation struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	attributes.Attributes
}

// TDataInputAssociation ...
type TDataInputAssociation struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	attributes.TAttributes
}

// DataObject ...
type DataObject struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// DataObjectReference ...
type DataObjectReference struct {
	compulsion.BaseAttributes
	attributes.Attributes
	DataObjectRef string `xml:"dataObjectRef,attr,omitempty" json:"dataObjectRef"`
}

// TDataObjectReference ...
type TDataObjectReference struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	DataObjectRef string `xml:"dataObjectRef,attr,omitempty" json:"dataObjectRef"`
}

// DataStoreReference ...
type DataStoreReference struct {
	compulsion.BaseAttributes
	attributes.Attributes
}

// TDataStoreReference ...
type TDataStoreReference struct {
	compulsion.BaseAttributes
	attributes.TAttributes
}
