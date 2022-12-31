package data

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
)

// DataInputAssociationRepository ...
type DataInputAssociationRepository interface {
	compulsion.IFBaseID
	attributes.AttributesBaseElements
}

// DataObjectRepository ...
type DataObjectRepository interface {
	compulsion.IFBaseID
}

// DataObjectReferenceRepository ...
type DataObjectReferenceRepository interface{}

// DataStoreRepository ...
type DataStoreReferenceRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
	attributes.AttributesBaseElements
}
