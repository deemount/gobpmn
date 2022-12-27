package data

import "github.com/deemount/gobpmn/models/attributes"

type STR_PTR *string
type DOCUMENTATION_PTR *attributes.Documentation
type EXTENSION_ELEMENTS_PTR *attributes.ExtensionElements

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
	GetDocumentation() DOCUMENTATION_PTR
}

type DataBaseExtensionElements interface {
	SetExtensionElements()
	GetExtensionElements() EXTENSION_ELEMENTS_PTR
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
