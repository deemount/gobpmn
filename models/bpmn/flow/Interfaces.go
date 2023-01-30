package flow

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/conditional"
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

/*
 * Base
 */

// FlowBaseReference ...
type FlowBaseReferences interface {
	SetSourceRef(typ string, sourceRef interface{})
	GetSourceRef() impl.STR_PTR
	SetTargetRef(typ string, targetRef interface{})
	GetTargetRef() impl.STR_PTR
}

// FlowSequenceFlow ...
type FlowSequenceFlow interface {
	SetSequenceFlow(num int)
	GetSequenceFlow(num int) *SequenceFlow
}

// FlowBase ...
type FlowBase interface {
	impl.IFBaseID
	impl.IFBaseName
}

/*
 * Repositories
 */

// AssociationRepository ...
type AssociationRepository interface {
	impl.IFBaseID
	attributes.AttributesBaseElements

	FlowBaseReferences
}

// DataInputAssociationRepository ...
type DataInputAssociationRepository interface {
	impl.IFBaseID
	attributes.AttributesBaseElements
}

// MessageFlowRepository ...
type MessageFlowRepository interface {
	FlowBase
	attributes.AttributesBaseElements
	FlowBaseReferences
}

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	FlowBase
	FlowBaseReferences
	attributes.AttributesBaseElements

	SetConditionExpression()
	GetConditionExpression() *conditional.ConditionExpression
}
