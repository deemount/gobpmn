package events

import "reflect"

// EditorGetID returns a string pointer
// The usage in an other package to create a bpmn could be:
// e.g.: events.Events.GetID(events.Events{}, e)
// e.g.: events.Events.GetID(events.Events{}, *elements.StartEvent)
// where argument e is mostly a pointer to
// a) definitions or
// b) elements
// This method is tested, but not used actually
func (events TEvents) EditorGetID(e interface{}) reflect.Value {
	return reflect.ValueOf(e)
}
