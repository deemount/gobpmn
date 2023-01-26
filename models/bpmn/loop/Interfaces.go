package loop

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/conditional"
)

type LoopBaseID interface{}
type LoopBaseName interface{}
type LoopCamundaBase interface{}
type LoopBaseCoreElements interface{}
type LoopBase interface{}

// LoopCardinalityRepository ...
type LoopCardinalityRepository interface{}

// MultiInstanceLoopCharacteristicsRepository ...
type MultiInstanceLoopCharacteristicsRepository interface {
	SetIsSequential(is bool)
	GetIsSequential() *bool

	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool

	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool

	SetCamundaCollection(collection string)
	GetCamundaCollection() *string

	SetCamundaElementVariable(element string)
	GetCamundaElementVariable() *string

	SetLoopCardinality()
	GetLoopCardinality() *LoopCardinality

	SetCompletionCondition()
	GetCompletionCondition() *conditional.CompletionCondition

	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

// ParticipantMultiplicity ...
type ParticipantMultiplicityRepository interface{}

// StandardLoopCharacteristicsRepository ...
type StandardLoopCharacteristicsRepository interface{}
