package elements

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/marker"
)

// BoundaryEvent ...
type BoundaryEvent struct {
	compulsion.BaseAttributes
	definitions.Boundaries
	AttachedToRef  string            `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity bool              `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	Outgoing       []marker.Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TBoundaryEvent ...
type TBoundaryEvent struct {
	compulsion.BaseAttributes
	definitions.TBoundaries
	AttachedToRef  string            `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity bool              `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	Outgoing       []marker.Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

// EndEvent ...
type EndEvent struct {
	compulsion.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	definitions.EndEvent
	Incoming []marker.Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
}

// TEndEvent ...
type TEndEvent struct {
	compulsion.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	definitions.TEndEvent
	Incoming []marker.Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	compulsion.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	marker.TIncomingOutgoing
	definitions.IntermediateCatchEvent
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	compulsion.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
	definitions.TIntermediateCatchEvent
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	compulsion.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
	definitions.IntermediateThrowEvent
}

// TIntermediateThrowEvent ...
type TIntermediateThrowEvent struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
	definitions.TIntermediateThrowEvent
}

// StartEvent ...
type StartEvent struct {
	compulsion.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	definitions.StartEvent
	IsInterrupting        bool              `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	CamundaFormKey        string            `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaFormRef        string            `xml:"camunda:formRef,attr,omitempty" json:"formRef,omitempty"`
	CamundaFormRefBind    string            `xml:"camunda:formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	CamundaFormRefVersion string            `xml:"camunda:formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	CamundaInitiator      string            `xml:"camunda:initiator,attr,omitempty" json:"init,omitempty"`
	Outgoing              []marker.Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	compulsion.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	definitions.TStartEvent
	IsInterrupting bool              `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
	FormKey        string            `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	FormRef        string            `xml:"formRef,attr,omitempty" json:"formRef,omitempty"`
	FormRefBind    string            `xml:"formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	FormRefVersion string            `xml:"formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	Initiator      string            `xml:"initiator,attr,omitempty" json:"init,omitempty"`
	Outgoing       []marker.Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}
