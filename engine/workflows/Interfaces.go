package workflows

import "context"

// All custom event types must satisfy this interface.
type Event interface {
	Handle(ctx context.Context) error
}

// All custom event listeners must satisfy this interface.
type Listener interface {
	Listen(event interface{})
}
