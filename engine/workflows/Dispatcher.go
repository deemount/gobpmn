package workflows

import (
	"context"
	"fmt"
	"reflect"
)

// http://www.inanzzz.com/index.php/post/2qdl/event-listener-and-dispatcher-example-with-golang

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Register(events ...Event) {
	for _, v := range events {
		d.events = append(d.events, reflect.TypeOf(v).String())
	}
}

func (d *Dispatcher) Dispatch(ctx context.Context, event Event) error {
	name := reflect.TypeOf(event).String()
	for _, v := range d.events {
		if v == name {
			return event.Handle(ctx)
		}
	}
	return fmt.Errorf("%s is not a registered event", name)
}

func (d *Dispatcher) Consume() {}
