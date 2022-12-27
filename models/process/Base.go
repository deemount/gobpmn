package process

import "github.com/deemount/gobpmn/models/attributes"

type STR_PTR *string
type DOCUMENTATION_PTR *attributes.Documentation
type EXTENSION_ELEMENTS_PTR *attributes.ExtensionElements

type ProcessBaseID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type ProcessBaseName interface {
	SetName(name string)
	GetName() STR_PTR
}

type ProcessBaseDocumentation interface {
	SetDocumentation()
	GetDocumentation() DOCUMENTATION_PTR
}

type ProcessBaseExtensionElements interface {
	SetExtensionElements()
	GetExtensionElements() EXTENSION_ELEMENTS_PTR
}

type ProcessBaseCoreElements interface {
	ProcessBaseDocumentation
	ProcessBaseExtensionElements
}

type ProcessBase interface {
	ProcessBaseID
	ProcessBaseName
	ProcessBaseCoreElements
}
