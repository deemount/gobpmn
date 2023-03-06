package factory

import (
	"sync"
)

// Event ...
type Event struct {
	Name string
	Data interface{}
}

// EventBus ...
type EventBus struct {
	subscribers map[string][]chan Event
	mutex       sync.RWMutex
}

// NewEventBus ...
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan Event),
	}
}

// Subscribe ...
func (eb *EventBus) Subscribe(eventName string) chan Event {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	ch := make(chan Event, 1)
	eb.subscribers[eventName] = append(eb.subscribers[eventName], ch)

	return ch
}

// Unsubscribe
func (eb *EventBus) Unsubscribe(eventName string, ch chan Event) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	if subscribers, ok := eb.subscribers[eventName]; ok {
		for i, subscriber := range subscribers {
			if subscriber == ch {
				eb.subscribers[eventName] = append(subscribers[:i], subscribers[i+1:]...)
				close(subscriber)
				return
			}
		}
	}
}

// Publish
func (eb *EventBus) Publish(event Event) {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()

	if subscribers, ok := eb.subscribers[event.Name]; ok {
		for _, subscriber := range subscribers {
			go func(subscriber chan Event) {
				subscriber <- event
			}(subscriber)
		}
	}
}
