package workflows

import "github.com/deemount/gobpmn/models/events/elements"

// Job represents events. When a new event is dispatched, it
// gets tuned into a job and put into `Dispatcher.jobs` channel.
type Job struct {
	eventName EVENT_NAME
	eventType interface{}
}

// Dispatcher represents ...
type Dispatcher struct {
	events []string
}

type EventQueue struct {
	InboundFlowId string
	Element       elements.TStartEvent
}
