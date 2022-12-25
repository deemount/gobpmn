package loop

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/conditional"
)

// MultiInstanceLoopCharacteristicsRepository ...
type MultiInstanceLoopCharacteristicsRepository interface {
	SetIsSequential(is bool)

	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaCollection(collection string)
	SetCamundaElementVariable(element string)

	SetExtensionElements()
	SetLoopCardinality()
	SetCompletionCondition()

	GetIsSequential() *bool

	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaCollection() *string
	GetCamundaElementVariable() *string

	GetExtensionElements() *attributes.ExtensionElements
	GetLoopCardinality() *LoopCardinality
	GetCompletionCondition() *conditional.CompletionCondition
}

// MultiInstanceLoopCharacteristics ...
type MultiInstanceLoopCharacteristics struct {
	IsSequential           bool                              `xml:"isSequential,attr,omitempty" json:"isSequential,omitempty"`
	CamundaAsyncBefore     bool                              `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter      bool                              `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaCollection      string                            `xml:"camunda:collection,attr,omitempty" json:"collection,omitempty"`
	CamundaElementVariable string                            `xml:"camunda:elementVariable,attr,omitempty" json:"elementVariable,omitempty"`
	ExtensionElements      []attributes.ExtensionElements    `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	LoopCardinality        []LoopCardinality                 `xml:"bpmn:loopCardinality,omitempty" json:"loopCardinality,omitempty"`
	CompletionCondition    []conditional.CompletionCondition `xml:"bpmn:completionCondition,omitempty" json:"completionCondition,omitempty"`
}

// TMultiInstanceLoopCharacteristics ...
type TMultiInstanceLoopCharacteristics struct {
	IsSequential        bool                              `xml:"isSequential,attr,omitempty" json:"isSequential,omitempty"`
	AsyncBefore         bool                              `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter          bool                              `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	Collection          string                            `xml:"collection,attr,omitempty" json:"collection,omitempty"`
	ElementVariable     string                            `xml:"elementVariable,attr,omitempty" json:"elementVariable,omitempty"`
	ExtensionElements   []attributes.TExtensionElements   `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	LoopCardinality     []LoopCardinality                 `xml:"loopCardinality,omitempty" json:"loopCardinality,omitempty"`
	CompletionCondition []conditional.CompletionCondition `xml:"completionCondition,omitempty" json:"completionCondition,omitempty"`
}

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
