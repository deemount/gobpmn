package marker

import "github.com/deemount/gobpmn/models/attributes"

type STR_PTR *string
type DOCUMENTATION_PTR *attributes.Documentation
type EXTENSION_ELEMENTS_PTR *attributes.ExtensionElements

type MarkerID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type MarkerName interface {
	SetName(suffix string)
	GetName() STR_PTR
}

type MarkerBaseReferences interface {
	SetSourceRef(typ string, sourceRef interface{})
	GetSourceRef() *string
	SetTargetRef(typ string, targetRef interface{})
	GetTargetRef() *string
}

type MarkerBaseCoreElements interface {
	SetDocumentation()
	GetDocumentation() DOCUMENTATION_PTR
	SetExtensionElements()
	GetExtensionElements() EXTENSION_ELEMENTS_PTR
}

type MarkerBase interface {
	MarkerID
	MarkerName
}
