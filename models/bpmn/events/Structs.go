package events

import (
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
)

type DelegateParameter struct {
	BE   *elements.BoundaryEvent
	SE   *elements.StartEvent
	EE   *elements.EndEvent
	ICE  *elements.IntermediateCatchEvent
	ITE  *elements.IntermediateThrowEvent
	SH   *canvas.Shape
	BS   canvas.Bounds
	TD   string   // TimerDefinition
	T    string   // Typ
	N    string   // Name
	H    []string // Hash
	TEDH string   // TimerEventDefinition Hash
}

// ProcessEvents ...
type ProcessEvents struct {
	StartEvent             START_EVENT_SLC              `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty" csv:"-"`
	BoundaryEvent          BOUNDARY_EVENT_SLC           `xml:"bpmn:boundaryEvent,omitemnpty" json:"boundaryEvent,omitempty" csv:"-"`
	EndEvent               END_EVENT_SLC                `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent INTERMEDIATE_CATCH_EVENT_SLC `xml:"bpmn:intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty" csv:"-"`
	IntermediateThrowEvent INTERMEDIATE_THROW_EVENT_SLC `xml:"bpmn:intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty" csv:"-"`
}

// TProcessEvents ...
type TProcessEvents struct {
	StartEvent             TSTART_EVENT_SLC              `xml:"startEvent,omitemnpty" json:"startEvent,omitempty" csv:"-"`
	BoundaryEvent          TBOUNDARY_EVENT_SLC           `xml:"boundaryEvent,omitemnpty" json:"boundaryEvent,omitempty" csv:"-"`
	EndEvent               TEND_EVENT_SLC                `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	IntermediateCatchEvent TINTERMEDIATE_CATCH_EVENT_SLC `xml:"intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty" csv:"-"`
	IntermediateThrowEvent TINTERMEDIATE_THROW_EVENT_SLC `xml:"intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty" csv:"-"`
}

// CoreEvents ...
type CoreEvents struct {
	Message MESSAGE_SLC `xml:"bpmn:message,omitemnpty" json:"message,omitempty" csv:"-"`
	Signal  SIGNAL_SLC  `xml:"bpmn:signal,omitemnpty" json:"signal,omitempty" csv:"-"`
}

// TCoreEvents ...
type TCoreEvents struct {
	Message TMESSAGE_SLC `xml:"message,omitemnpty" json:"message,omitempty" csv:"-"`
	Signal  TSIGNAL_SLC  `xml:"signal,omitemnpty" json:"signal,omitempty" csv:"-"`
}
