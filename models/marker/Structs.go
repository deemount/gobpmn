package marker

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/conditional"
)

// Notice: can be implemented in seperate package
type SourceTargetRef struct {
	SourceRef string `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef string `xml:"targetRef,attr" json:"targetRef,omitempty"`
}

type IncomingOutgoing struct {
	Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

type TIncomingOutgoing struct {
	Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
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

// Category ...
type Category struct {
	ID            string          `xml:"id,attr,omitempty" json:"id"`
	CategoryValue []CategoryValue `xml:"bpmn:categoryValue,omitempty" json:"categoryValue,omitempty"`
	attributes.Attributes
}

// TCategory ...
type TCategory struct {
	ID            string          `xml:"id,attr,omitempty" json:"id"`
	CategoryValue []CategoryValue `xml:"categoryValue,omitempty" json:"categoryValue,omitempty"`
	attributes.TAttributes
}

// CategoryValue ...
type CategoryValue struct {
	ID    string `xml:"id,attr,omitempty" json:"id"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

// Group ...
type Group struct {
	ID               string `xml:"id,attr,omitempty" json:"id"`
	CategoryValueRef string `xml:"categoryValueRef,attr,omitempty" json:"categoryValueRef,omitempty"`
	attributes.Attributes
}

// TGroup ...
type TGroup struct {
	ID               string `xml:"id,attr,omitempty" json:"id"`
	CategoryValueRef string `xml:"categoryValueRef,attr,omitempty" json:"categoryValueRef,omitempty"`
	attributes.TAttributes
}

// Incoming ...
type Incoming struct {
	Flow string `xml:",innerxml" json:"flow"`
}

// Outgoing ...
type Outgoing struct {
	Flow string `xml:",innerxml" json:"flow"`
}

// Message ...
type Message struct {
	compulsion.BaseAttributes
}

// TMessage ...
type TMessage struct {
	compulsion.BaseAttributes
}

// MessageFlow ...
type MessageFlow struct {
	compulsion.BaseAttributes
	attributes.Attributes
	SourceTargetRef
}

// TMessageFlow ...
type TMessageFlow struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	SourceTargetRef
}

// SequenceFlow ...
type SequenceFlow struct {
	compulsion.BaseAttributes
	attributes.Attributes
	SourceTargetRef
	ConditionExpression []conditional.ConditionExpression `xml:"bpmn:conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}

// TSequenceFlow ...
type TSequenceFlow struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	SourceTargetRef
	ConditionExpression []conditional.ConditionExpression `xml:"conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}

// Signal ...
type Signal struct {
	compulsion.BaseAttributes
}

// TSignal ...
type TSignal struct {
	compulsion.BaseAttributes
}
