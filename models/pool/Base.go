package pool

import "github.com/deemount/gobpmn/models/attributes"

type STR_PTR *string
type DOCUMENTATION_PTR *attributes.Documentation
type EXTENSION_ELEMENTS_PTR *attributes.ExtensionElements

type PoolID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type PoolName interface {
	SetName(name string)
	GetName() STR_PTR
}

type PoolBaseDocumentation interface {
	SetDocumentation()
	GetDocumentation() DOCUMENTATION_PTR
}

type PoolBaseExtensionElements interface {
	SetExtensionElements()
	GetExtensionElements() EXTENSION_ELEMENTS_PTR
}

type PoolBaseCoreElements interface {
	PoolBaseDocumentation
	PoolBaseExtensionElements
}

type PoolBase interface {
	PoolID
	PoolName
}
