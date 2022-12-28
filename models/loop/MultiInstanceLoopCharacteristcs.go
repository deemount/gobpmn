package loop

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/conditional"
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
func (milc *MultiInstanceLoopCharacteristics) SetIsSequential(is bool) {
	milc.IsSequential = is
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (milc *MultiInstanceLoopCharacteristics) SetCamundaAsyncBefore(asyncBefore bool) {
	milc.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (milc *MultiInstanceLoopCharacteristics) SetCamundaAsyncAfter(asyncAfter bool) {
	milc.CamundaAsyncAfter = asyncAfter
}

// SetCamundaCollection ...
func (milc *MultiInstanceLoopCharacteristics) SetCamundaCollection(collection string) {
	milc.CamundaCollection = collection
}

// SetCamundaElementVariable ...
func (milc *MultiInstanceLoopCharacteristics) SetCamundaElementVariable(element string) {
	milc.CamundaElementVariable = element
}

/*** Make Elements ***/

/** BPMN **/

// SetExtensionElements ...
func (milc *MultiInstanceLoopCharacteristics) SetExtensionElements() {
	milc.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetLoopCardinality ...
func (milc *MultiInstanceLoopCharacteristics) SetLoopCardinality() {
	milc.LoopCardinality = make([]LoopCardinality, 1)
}

// SetCompletionCondition ...
func (milc *MultiInstanceLoopCharacteristics) SetCompletionCondition() {
	milc.CompletionCondition = make([]conditional.CompletionCondition, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetIsSequential ...
func (milc MultiInstanceLoopCharacteristics) GetIsSequential() *bool {
	return &milc.IsSequential
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (milc MultiInstanceLoopCharacteristics) GetCamundaAsyncBefore() *bool {
	return &milc.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (milc MultiInstanceLoopCharacteristics) GetCamundaAsyncAfter() *bool {
	return &milc.CamundaAsyncAfter
}

// GetCamundaCollection ...
func (milc MultiInstanceLoopCharacteristics) GetCamundaCollection() *string {
	return &milc.CamundaCollection
}

// GetCamundaElementVariable ...
func (milc MultiInstanceLoopCharacteristics) GetCamundaElementVariable() *string {
	return &milc.CamundaElementVariable
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (milc MultiInstanceLoopCharacteristics) GetExtensionElements() *attributes.ExtensionElements {
	return &milc.ExtensionElements[0]
}

// GetLoopCardinality ...
func (milc MultiInstanceLoopCharacteristics) GetLoopCardinality() *LoopCardinality {
	return &milc.LoopCardinality[0]
}

// GetCompletionCondition ...
func (milc MultiInstanceLoopCharacteristics) GetCompletionCondition() *conditional.CompletionCondition {
	return &milc.CompletionCondition[0]
}
