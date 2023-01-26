package events

import "github.com/deemount/gobpmn/models/bpmn/events/elements"

type EVENTS_SLC []Events

type START_EVENT_SLC []elements.StartEvent
type BOUNDARY_EVENT_SLC []elements.BoundaryEvent
type END_EVENT_SLC []elements.EndEvent
type INTERMEDIATE_CATCH_EVENT_SLC []elements.IntermediateCatchEvent
type INTERMEDIATE_THROW_EVENT_SLC []elements.IntermediateThrowEvent

type TSTART_EVENT_SLC []elements.TStartEvent
type TBOUNDARY_EVENT_SLC []elements.TBoundaryEvent
type TEND_EVENT_SLC []elements.TEndEvent
type TINTERMEDIATE_CATCH_EVENT_SLC []elements.TIntermediateCatchEvent
type TINTERMEDIATE_THROW_EVENT_SLC []elements.TIntermediateThrowEvent
