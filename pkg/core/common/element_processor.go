package common

import (
	"context"
	"fmt"
	"reflect"

	"github.com/deemount/gobpmn/pkg/types"
)

type ElementHandler interface {
	Handle(processIdx int, info FieldInfo, config *ProcessingConfig) error
}

// ElementProcessor handles BPMN element processing
type ElementProcessor[M types.BPMNGeneric] struct {
	value    *ReflectValue[M]
	quantity *Quantity[M]
	handlers map[ElementType]ElementHandler
}

// ProcessingConfig holds the configuration for element processing
type ProcessingConfig struct {
	field     reflect.Value
	numFields int
	indices   map[processElement]int
	counts    map[processElement]int
}

// NewElementProcessor creates a new ElementProcessor instance
func NewElementProcessor[M types.BPMNGeneric](v *ReflectValue[M], q *Quantity[M]) (*ElementProcessor[M], error) {
	if v == nil || q == nil {
		return nil, NewError(fmt.Errorf("invalid parameters: ReflectValue and quantity must not be nil"))
	}

	ep := &ElementProcessor[M]{
		value:    v,
		quantity: q,
		handlers: make(map[ElementType]ElementHandler),
	}

	ep.registerHandlers()

	return ep, nil

}

// registerHandlers registers handlers for different types of BPMN elements
func (ep *ElementProcessor[M]) registerHandlers() {
	ep.handlers[TypeStartEvent] = NewStartEventHandler(ep)
	ep.handlers[TypeEvent] = NewEventHandler(ep)
	ep.handlers[TypeActivity] = NewActivityHandler(ep)
	ep.handlers[TypeFlow] = NewFlowHandler(ep)
	ep.handlers[TypeGateway] = NewGatewayHandler(ep)
}

// createProcessingConfig initializes processing configuration
func (ep *ElementProcessor[M]) createProcessingConfig(processIdx int, field reflect.Value) *ProcessingConfig {
	return &ProcessingConfig{
		field:     field,
		numFields: field.NumField(),
		indices:   initializeIndices(),
		counts:    ep.quantity.Elements[processIdx],
	}
}

/*
 * @processes
 */

// ProcessElementsWithContext adds context support for processing multiple elements
func (ep *ElementProcessor[M]) ProcessElementsWithContext(ctx context.Context) error {
	done := make(chan error, 1) // NOTE: use a buffered channel

	go func() {
		done <- ep.processMultipleElements()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return NewError(fmt.Errorf("multiple elements processing cancelled:\n%w", ctx.Err()))
	}

}

// processMultipleElements processes all elements across multiple processes.
// It handles error slice for better resource management.
// It ranges over the process names and processes each process.
// Each field is taken from v.Instance.
func (ep *ElementProcessor[M]) processMultipleElements() error {
	var errs []error

	for processIdx, processName := range ep.value.ProcessName {
		field := instanceFieldName(ep.value, processName) // NOTE: v.Instance is called once here in the whole element processor.

		if !field.IsValid() {
			errs = append(errs, NewError(fmt.Errorf("invalid field for process:\n%s", processName)))
			continue // silently continuing
		}

		// collect errors
		if err := ep.processProcess(processIdx, field); err != nil {
			errs = append(errs, NewError(fmt.Errorf("process %s:\n%w", processName, err)))
		}
	}

	if len(errs) > 0 {
		return NewError(fmt.Errorf("multiple errors occurred:\n%v", errs))
	}

	return nil

}

// processProcess handles processing for a single process in a multiple process model, like collaboration.
func (ep *ElementProcessor[M]) processProcess(processIdx int, field reflect.Value) error {
	config := ep.createProcessingConfig(processIdx, field)

	for fieldIdx := range config.numFields {
		if err := ep.processField(processIdx, fieldIdx, config); err != nil {
			return NewError(fmt.Errorf("failed to process field at index %d:\n%w", fieldIdx, err))
		}
	}

	return nil

}

// processField handles processing of individual fields.
func (ep *ElementProcessor[M]) processField(processIdx, fieldIdx int, config *ProcessingConfig) error {

	info := extractFieldInfo(config.field, fieldIdx)
	if !isValidField(info) {
		return nil
	}

	elementType := ElementType(info.typ)

	handler, exists := ep.handlers[elementType]
	if !exists {
		return NewError(fmt.Errorf("no handler found for element type:\n%s", elementType))
	}

	return handler.Handle(processIdx, info, config)

}

/*
 * @standalone process
 */

// ProcessStandalone adds context support for processing a single element.
func (ep *ElementProcessor[M]) ProcessStandalone(ctx context.Context) error {
	done := make(chan error, 1)
	go func() {
		done <- ep.processElement()
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return NewError(fmt.Errorf("single element processing cancelled:\n%w", ctx.Err()))
	}
}

// processElement processes elements for a standalone process
// The Pos field is used to sort the elements and process them in order.
func (ep *ElementProcessor[M]) processElement() error {
	config := &ProcessingConfig{
		indices: initializeIndices(),
		counts:  ep.quantity.Elements[0],
	}

	/*
		switch fields := any(ep.value.Fields).(type) {
		case []reflect.StructField:
	*/
	// NOTE: the first field is DEF, which is a call on a interface value.
	//       The second field is mostly a bool value, then the Fields with the BPMN types
	//       are following. The first field of the bpmn types is a process itself.
	//       These fields described above shouldn't processed here.
	//	     Write an condition to skip these fields, if given.
	for fieldIdx := 3; fieldIdx < ep.value.InstanceNumField; fieldIdx++ {
		field := ep.value.Instance.Field(fieldIdx)
		fieldType := ep.value.Instance.Type().Field(fieldIdx)

		if field.Kind() == reflect.Bool {
			continue // NOTE: placeholder for future implementation of configuration handling
		}

		if err := ep.processEntry(field, fieldType, fieldIdx, config); err != nil {
			return NewError(fmt.Errorf("failed to process field %s:\n%w", fieldType.Name, err))
		}
	}

	return nil
}

// processEntry processes a single struct field or map entry
// BUG: processEntry is not working as expected with map[string]any
func (ep *ElementProcessor[M]) processEntry(field reflect.Value, fieldType reflect.StructField, fieldIdx int, config *ProcessingConfig) error {
	var typ, hash, hashBefore, nextHash string

	typ = field.FieldByName("Type").String()
	hash = field.FieldByName("Hash").String()
	hashBefore = ep.value.Instance.Field(fieldIdx - 1).FieldByName("Hash").String()
	nextHash = ep.nextHash(fieldIdx)

	// get the element type
	elementType := ElementType(typ)
	handler, exists := ep.handlers[elementType]
	if !exists {
		return NewError(fmt.Errorf("no handler found for element type:\n%s", elementType))
	}

	// create FieldInfo
	info := FieldInfo{
		Name:       fieldType.Name,
		typ:        typ,
		hash:       hash,
		hashBefore: hashBefore,
		nextHash:   nextHash,
		element:    matchElementType(fieldType.Name),
	}

	return handler.Handle(0, info, config)
}

// nextHash returns the next hash value
// Note:
func (ep *ElementProcessor[M]) nextHash(fieldIdx int) string {
	if fieldIdx+1 < ep.value.InstanceNumField {
		return ep.value.Instance.Field(fieldIdx + 1).FieldByName("Hash").String()
	}
	return ""
}

// Cleanup ...
func (ep *ElementProcessor[M]) Cleanup() {
	// clean up handlers
	for _, handler := range ep.handlers {
		if cleaner, ok := handler.(Cleanup); ok {
			cleaner.Cleanup()
		}
	}

	// clear maps and slices
	ep.handlers = nil
	ep.value = nil
	ep.quantity = nil

}
