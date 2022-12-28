package marker

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/conditional"
)

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
	GetDocumentation() *attributes.Documentation
	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

type MarkerBase interface {
	MarkerID
	MarkerName
}

// AssociationRepository ...
type AssociationRepository interface {
	MarkerID
	MarkerBaseCoreElements

	MarkerBaseReferences
}

// CategoryRepository ...
type CategoryRepository interface {
	MarkerID
	MarkerBaseCoreElements

	SetCategoryValue()
	GetCategoryValue() *CategoryValue
}

// CategoryValueRepository ...
type CategoryValueRepository interface {
	MarkerID

	SetValue(value string)
	GetValue() *string
}

// IncomingRepository ...
type IncomingRepository interface {
	SetFlow(suffix string)
	GetFlow() *string
}

// OutgoingRepository ...
type OutgoingRepository interface {
	SetFlow(suffix string)
	GetFlow() *string
}

// GroupRepository ...
type GroupRepository interface {
	MarkerID

	SetCategoryValueRef(suffix string)
	GetCategoryValueRef() *string

	SetDocumentation()
	GetDocumentation() *attributes.Documentation
	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

// MessageRepository ...
type MessageRepository interface {
	MarkerBase
}

// MessageFlowRepository ...
type MessageFlowRepository interface {
	MarkerBase
	MarkerBaseCoreElements
	MarkerBaseReferences
}

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	MarkerBase
	MarkerBaseReferences
	MarkerBaseCoreElements

	SetConditionExpression()
	GetConditionExpression() *conditional.ConditionExpression
}

// SignalRepository ...
type SignalRepository interface {
	MarkerBase
}
