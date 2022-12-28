package loop

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/conditional"
)

// LoopCardinality ...
type LoopCardinality struct{}

// TLoopCardinality ...
type TLoopCardinality struct{}

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

// ParticipantMultiplicity ...
type ParticipantMultiplicity struct{}

// TParticipantMultiplicity ...
type TParticipantMultiplicity struct{}

// StandardLoopCharacteristics ...
type StandardLoopCharacteristics struct{}

// TStandardLoopCharacteristics ...
type TStandardLoopCharacteristics struct{}
