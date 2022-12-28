package data

import "github.com/deemount/gobpmn/models/attributes"

// DataBaseID ...
type DataBaseID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

// DataBaseName ...
type DataBaseName interface {
	SetName(name string)
	GetName() STR_PTR
}

type DataBaseDocumentation interface {
	SetDocumentation()
	GetDocumentation() *attributes.Documentation
}

type DataBaseExtensionElements interface {
	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

type DataBaseCoreElements interface {
	DataBaseDocumentation
	DataBaseExtensionElements
}

// DataBaseAttributes ...
type DataBaseAttributes interface {
	DataBaseID
	DataBaseName
}

// DataInputAssociationRepository ...
type DataInputAssociationRepository interface {
	DataBaseID
	DataBaseCoreElements
}

// DataObjectRepository ...
type DataObjectRepository interface {
	DataBaseID
}

// DataObjectReferenceRepository ...
type DataObjectReferenceRepository interface{}

// DataStoreRepository ...
type DataStoreReferenceRepository interface {
	DataBaseAttributes
	DataBaseCoreElements
}
