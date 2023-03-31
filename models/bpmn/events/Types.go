package events

import "github.com/deemount/gobpmn/models/bpmn/events/elements"

type BOUNDARY_EVENT_PTR *elements.BoundaryEvent
type END_EVENT_PTR *elements.EndEvent
type INTERMEDIATE_CATCH_EVENT_PTR *elements.IntermediateCatchEvent
type INTERMEDIATE_THROW_EVENT_PTR *elements.IntermediateThrowEvent
type MESSAGE_PTR *elements.Message
type SIGNAL_PTR *elements.Signal
type START_EVENT_PTR *elements.StartEvent

type PROCESS_EVENTS_SLC []ProcessEvents
type CORE_EVENTS_SLC []CoreEvents

type BOUNDARY_EVENT_SLC []elements.BoundaryEvent
type END_EVENT_SLC []elements.EndEvent
type INTERMEDIATE_CATCH_EVENT_SLC []elements.IntermediateCatchEvent
type INTERMEDIATE_THROW_EVENT_SLC []elements.IntermediateThrowEvent
type MESSAGE_SLC []elements.Message
type SIGNAL_SLC []elements.Signal
type START_EVENT_SLC []elements.StartEvent

// T

type TBOUNDARY_EVENT_SLC []elements.TBoundaryEvent
type TEND_EVENT_SLC []elements.TEndEvent
type TINTERMEDIATE_CATCH_EVENT_SLC []elements.TIntermediateCatchEvent
type TINTERMEDIATE_THROW_EVENT_SLC []elements.TIntermediateThrowEvent
type TMESSAGE_SLC []elements.TMessage
type TSIGNAL_SLC []elements.TSignal
type TSTART_EVENT_SLC []elements.TStartEvent
