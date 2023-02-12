package elements

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/camunda"
<<<<<<< HEAD
	"github.com/deemount/gobpmn/models/bpmn/canvas"
=======
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	"github.com/deemount/gobpmn/models/bpmn/events/definitions"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

// event
<<<<<<< HEAD
type DelegateParameter struct {
	BE    *BoundaryEvent
	SE    *StartEvent
	EE    *EndEvent
	ICE   *IntermediateCatchEvent
	ITE   *IntermediateThrowEvent
	MSG   *Message
	SH    *canvas.Shape
	BS    canvas.Bounds
	ISTED bool   // IsTimerEventDefinition
	TD    string // TimerDefinition
	T     string
	N     string
	H     []string
	TEDH  string // TimerEventDefinition Hash
=======
// BE: be *BoundaryEvent, typ string, name string, hash string
// SE: se *StartEvent, typ string, name string, hash string
// EE: ee *EndEvent, typ string, name string, hash string
// ICE: ice *IntermediateCatchEvent, typ string, name string, hash string
// ITE: ite *IntermediateCatchEvent, typ string, name string, hash string
type DelegateParameter struct {
	BE  *BoundaryEvent
	SE  *StartEvent
	EE  *EndEvent
	ICE *IntermediateCatchEvent
	ITE *IntermediateThrowEvent
	T   string
	N   string
	H   []string
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

/*
 * Elementary
 */

// BoundaryEvent ...
type BoundaryEvent struct {
	impl.BaseAttributes
	definitions.Boundaries
	AttachedToRef  string            `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity bool              `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	Outgoing       []marker.Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TBoundaryEvent ...
type TBoundaryEvent struct {
	impl.BaseAttributes
	definitions.TBoundaries
	AttachedToRef  string            `xml:"attachedToRef,attr,omitempty" json:"attachedToRef,omitempty"`
	CancelActivity bool              `xml:"cancelActivity,attr,omitempty" json:"cancelActivity,omitempty"` // maybe @deprecated in new modeler
	Outgoing       []marker.Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

// EndEvent ...
type EndEvent struct {
	impl.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	definitions.EndEvent
	Incoming []marker.Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
}

// TEndEvent ...
type TEndEvent struct {
	impl.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	definitions.TEndEvent
	Incoming []marker.Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	impl.BaseAttributes
	camunda.CoreAttributes
	attributes.Attributes
	marker.TIncomingOutgoing
	definitions.IntermediateCatchEvent
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	impl.BaseAttributes
	camunda.TCoreAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
	definitions.TIntermediateCatchEvent
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	impl.BaseAttributes
	attributes.Attributes
	marker.IncomingOutgoing
	definitions.IntermediateThrowEvent
}

// TIntermediateThrowEvent ...
type TIntermediateThrowEvent struct {
	impl.BaseAttributes
	attributes.TAttributes
	marker.TIncomingOutgoing
	definitions.TIntermediateThrowEvent
}

// StartEvent ...
type StartEvent struct {
	impl.BaseAttributes
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
	impl.BaseAttributes
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

// Message ...
type Message struct {
	impl.BaseAttributes
}

// TMessage ...
type TMessage struct {
	impl.BaseAttributes
}

// Signal ...
type Signal struct {
	impl.BaseAttributes
}

// TSignal ...
type TSignal struct {
	impl.BaseAttributes
}
