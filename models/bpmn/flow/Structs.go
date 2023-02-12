package flow

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
<<<<<<< HEAD
	"github.com/deemount/gobpmn/models/bpmn/canvas"
=======
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	"github.com/deemount/gobpmn/models/bpmn/conditional"
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// flow
type DelegateParameter struct {
<<<<<<< HEAD
	SF *SequenceFlow
	MF *MessageFlow
	ED *canvas.Edge
	WP []canvas.Waypoint
	BS canvas.Bounds
	ST string   // source type
	TT string   // target type
	T  string   // typ
	N  string   // name
	H  []string // hash
=======
	SF *SequenceFlow // element pointer
	ST string        // source type
	TT string        // target type
	T  string        // typ
	N  string        // name
	H  []string      // hash
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

type SourceTargetRef struct {
	SourceRef string `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef string `xml:"targetRef,attr" json:"targetRef,omitempty"`
}

// Association ...
type Association struct {
	ID string `xml:"id,attr"`
	SourceTargetRef
	attributes.Attributes
}

// TAssociation ...
type TAssociation struct {
	ID string `xml:"id,attr"`
	SourceTargetRef
	attributes.TAttributes
}

// DataInputAssociation ...
type DataInputAssociation struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	attributes.Attributes
}

// TDataInputAssociation ...
type TDataInputAssociation struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
	attributes.TAttributes
}

// MessageFlow ...
type MessageFlow struct {
	impl.BaseAttributes
	attributes.Attributes
	SourceTargetRef
}

// TMessageFlow ...
type TMessageFlow struct {
	impl.BaseAttributes
	attributes.TAttributes
	SourceTargetRef
}

// SequenceFlow ...
type SequenceFlow struct {
	impl.BaseAttributes
	attributes.Attributes
	SourceTargetRef
	ConditionExpression []conditional.ConditionExpression `xml:"bpmn:conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}

// TSequenceFlow ...
type TSequenceFlow struct {
	impl.BaseAttributes
	attributes.TAttributes
	SourceTargetRef
	ConditionExpression []conditional.ConditionExpression `xml:"conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}
