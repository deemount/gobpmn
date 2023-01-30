package loop

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/conditional"
)

// NewMultiInstanceLoopCharacteristics ...
func NewMultiInstanceLoopCharacteristics() MultiInstanceLoopCharacteristicsRepository {
	return &MultiInstanceLoopCharacteristics{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetIsSequential ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetIsSequential(is bool) {
	multiInstanceLoopCharacteristics.IsSequential = is
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCamundaAsyncBefore(asyncBefore bool) {
	multiInstanceLoopCharacteristics.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCamundaAsyncAfter(asyncAfter bool) {
	multiInstanceLoopCharacteristics.CamundaAsyncAfter = asyncAfter
}

// SetCamundaCollection ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCamundaCollection(collection string) {
	multiInstanceLoopCharacteristics.CamundaCollection = collection
}

// SetCamundaElementVariable ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCamundaElementVariable(element string) {
	multiInstanceLoopCharacteristics.CamundaElementVariable = element
}

/*** Make Elements ***/

/** BPMN **/

// SetExtensionElements ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetExtensionElements() {
	multiInstanceLoopCharacteristics.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetLoopCardinality ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetLoopCardinality() {
	multiInstanceLoopCharacteristics.LoopCardinality = make([]LoopCardinality, 1)
}

// SetCompletionCondition ...
func (multiInstanceLoopCharacteristics *MultiInstanceLoopCharacteristics) SetCompletionCondition() {
	multiInstanceLoopCharacteristics.CompletionCondition = make([]conditional.CompletionCondition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetIsSequential ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetIsSequential() *bool {
	return &multiInstanceLoopCharacteristics.IsSequential
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaAsyncBefore() *bool {
	return &multiInstanceLoopCharacteristics.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaAsyncAfter() *bool {
	return &multiInstanceLoopCharacteristics.CamundaAsyncAfter
}

// GetCamundaCollection ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaCollection() *string {
	return &multiInstanceLoopCharacteristics.CamundaCollection
}

// GetCamundaElementVariable ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCamundaElementVariable() *string {
	return &multiInstanceLoopCharacteristics.CamundaElementVariable
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetExtensionElements() *attributes.ExtensionElements {
	return &multiInstanceLoopCharacteristics.ExtensionElements[0]
}

// GetLoopCardinality ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetLoopCardinality() *LoopCardinality {
	return &multiInstanceLoopCharacteristics.LoopCardinality[0]
}

// GetCompletionCondition ...
func (multiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics) GetCompletionCondition() *conditional.CompletionCondition {
	return &multiInstanceLoopCharacteristics.CompletionCondition[0]
}
