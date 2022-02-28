package models

// MultiInstanceLoopCharacteristics ...
type MultiInstanceLoopCharacteristics struct {
	IsSequential           bool                  `xml:"isSequential,attr,omitempty" json:"isSequential,omitempty"`
	CamundaAsyncBefore     bool                  `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter      bool                  `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaCollection      string                `xml:"camunda:collection,attr,omitempty" json:"collection,omitempty"`
	CamundaElementVariable string                `xml:"camunda:elementVariable,attr,omitempty" json:"elementVariable,omitempty"`
	ExtensionElements      []ExtensionElements   `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	LoopCardinality        []LoopCardinality     `xml:"bpmn:loopCardinality,omitempty" json:"loopCardinality,omitempty"`
	CompletionCondition    []CompletionCondition `xml:"bpmn:completionCondition,omitempty" json:"completionCondition,omitempty"`
}

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

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (milc *MultiInstanceLoopCharacteristics) SetExtensionElements() {
	milc.ExtensionElements = make([]ExtensionElements, 1)
}

// SetLoopCardinality ...
func (milc *MultiInstanceLoopCharacteristics) SetLoopCardinality() {
	milc.LoopCardinality = make([]LoopCardinality, 1)
}

// SetCompletionCondition ...
func (milc *MultiInstanceLoopCharacteristics) SetCompletionCondition() {
	milc.CompletionCondition = make([]CompletionCondition, 1)
}
