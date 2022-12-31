package marker

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/conditional"
)

/*
 * Base
 */

// MarkerBaseReference ...
type MarkerBaseReferences interface {
	SetSourceRef(typ string, sourceRef interface{})
	GetSourceRef() *string
	SetTargetRef(typ string, targetRef interface{})
	GetTargetRef() *string
}

// MarkerFlow ...
type MarkerFlow interface {
	SetFlow(suffix string)
	GetFlow() *string
}

// MarkerSequenceFlow ...
type MarkerSequenceFlow interface {
	SetSequenceFlow(num int)
	GetSequenceFlow(num int) *SequenceFlow
}

// MarkerIncoming ...
type MarkerIncoming interface {
	SetIncoming(num int)
	GetIncoming(num int) *Incoming
}

// MarkerOutgoing ...
type MarkerOutgoing interface {
	SetOutgoing(num int)
	GetOutgoing(num int) *Outgoing
}

// MarkerIncomingOutgoing
type MarkerIncomingOutgoing interface {
	MarkerIncoming
	MarkerOutgoing
}

// MarkerBase ...
type MarkerBase interface {
	compulsion.IFBaseID
	compulsion.IFBaseName
}

/*
 * Repositories
 */

// AssociationRepository ...
type AssociationRepository interface {
	compulsion.IFBaseID
	attributes.AttributesBaseElements

	MarkerBaseReferences
}

// CategoryRepository ...
type CategoryRepository interface {
	compulsion.IFBaseID
	attributes.AttributesBaseElements

	SetCategoryValue()
	GetCategoryValue() *CategoryValue
}

// CategoryValueRepository ...
type CategoryValueRepository interface {
	compulsion.IFBaseID

	SetValue(value string)
	GetValue() *string
}

// IncomingRepository ...
type IncomingRepository interface{ MarkerFlow }

// OutgoingRepository ...
type OutgoingRepository interface{ MarkerFlow }

// GroupRepository ...
type GroupRepository interface {
	compulsion.IFBaseID

	SetCategoryValueRef(suffix string)
	GetCategoryValueRef() *string

	attributes.AttributesBaseElements
}

// MessageRepository ...
type MessageRepository interface {
	MarkerBase
}

// MessageFlowRepository ...
type MessageFlowRepository interface {
	MarkerBase
	attributes.AttributesBaseElements
	MarkerBaseReferences
}

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	MarkerBase
	MarkerBaseReferences
	attributes.AttributesBaseElements

	SetConditionExpression()
	GetConditionExpression() *conditional.ConditionExpression
}

// SignalRepository ...
type SignalRepository interface {
	MarkerBase
}
