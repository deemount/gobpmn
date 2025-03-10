package common

import (
	"context"
	"fmt"
	"reflect"
)

// ElementType represents the type of BPMN element
type ElementType string

func (et ElementType) String() string {
	return string(et)
}

const (
	TypeStartEvent ElementType = "startevent"
	TypeEvent      ElementType = "event"
	TypeFlow       ElementType = "flow"
	TypeGateway    ElementType = "gateway"
	TypeActivity   ElementType = "activity"
)

// ElementProcessor handles BPMN element processing
type ElementProcessor struct {
	value    *ReflectValue
	quantity *Quantity
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
func NewElementProcessor(v *ReflectValue, q *Quantity) (*ElementProcessor, error) {
	if v == nil || q == nil {
		return nil, fmt.Errorf("invalid parameters: ReflectValue and quantity must not be nil")
	}

	ep := &ElementProcessor{
		value:    v,
		quantity: q,
		handlers: make(map[ElementType]ElementHandler),
	}

	ep.registerHandlers()

	return ep, nil

}

// registerHandlers registers handlers for different types of BPMN elements
func (ep *ElementProcessor) registerHandlers() {
	ep.handlers[TypeStartEvent] = NewStartEventHandler(ep)
	ep.handlers[TypeEvent] = NewEventHandler(ep)
	ep.handlers[TypeActivity] = NewActivityHandler(ep)
	ep.handlers[TypeFlow] = NewFlowHandler(ep)
	ep.handlers[TypeGateway] = NewGatewayHandler(ep)
}

// createProcessingConfig initializes processing configuration
func (ep *ElementProcessor) createProcessingConfig(processIdx int, field reflect.Value) *ProcessingConfig {
	return &ProcessingConfig{
		field:     field,
		numFields: field.NumField(),
		indices:   initializeIndices(),
		//indices: NewElementIndices(),
		counts: ep.quantity.Elements[processIdx],
	}
}

/*
 * @processes
 */

// ProcessElementsWithContext adds context support for processing multiple elements
func (ep *ElementProcessor) ProcessElements(ctx context.Context) error {
	done := make(chan error, 1) // NOTE: buffered channel

	go func() {
		done <- ep.processMultipleElements()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return fmt.Errorf("multiple elements processing cancelled: %w", ctx.Err())
	}

}

// ProcessMultipleElement processes all elements across multiple processes.
// It handles error slice for better resource management.
// It ranges over the process names and processes each process.
// Each field is taken from v.Instance.
func (ep *ElementProcessor) processMultipleElements() error {
	var errs []error

	for processIdx, processName := range ep.value.ProcessName {
		field := targetFieldName(ep.value, processName) // NOTE: v.Instance is called once here in the whole element processor.

		if !field.IsValid() {
			errs = append(errs, fmt.Errorf("invalid field for process: %s", processName))
			continue // silently continuing
		}

		if err := ep.processProcess(processIdx, field); err != nil {
			errs = append(errs, fmt.Errorf("process %s: %w", processName, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("multiple errors occurred: %v", errs)
	}

	return nil

}

// processProcess handles processing for a single process in a multiple process model, like collaboration.
func (ep *ElementProcessor) processProcess(processIdx int, field reflect.Value) error {
	config := ep.createProcessingConfig(processIdx, field)

	for fieldIdx := 0; fieldIdx < config.numFields; fieldIdx++ {
		if err := ep.processField(processIdx, fieldIdx, config); err != nil {
			return fmt.Errorf("failed to process field at index %d: %w", fieldIdx, err)
		}
	}

	return nil

}

// processField handles processing of individual fields.
func (ep *ElementProcessor) processField(processIdx, fieldIdx int, config *ProcessingConfig) error {

	info := extractFieldInfo(config.field, fieldIdx)
	if !isValidField(info) {
		return nil
	}

	elementType := ElementType(info.typ)

	handler, exists := ep.handlers[elementType]
	if !exists {
		return fmt.Errorf("no handler found for element type: %s", elementType)
	}

	return handler.Handle(processIdx, info, config)

}

/*
 * @single process
 */

// ProcessStandalone adds context support for processing a single element.
func (ep *ElementProcessor) ProcessStandalone(ctx context.Context) error {
	done := make(chan error, 1)
	go func() {
		done <- ep.processSingleElement()
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return fmt.Errorf("single element processing cancelled: %w", ctx.Err())
	}
}

// ProcessStandalone processes elements for a single process
func (ep *ElementProcessor) processSingleElement() error {

	config := &ProcessingConfig{
		indices: initializeIndices(),
		counts:  ep.quantity.Elements[0],
	}

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

		// NOTE: somehow the first field is DEF, which is a call on a interface value. This is not handled.
		if err := ep.processSingleStructField(field, fieldType, fieldIdx, config); err != nil {
			return fmt.Errorf("failed to process field %s: %w", fieldType.Name, err)
		}

	}

	return nil

}

// processSingleStructField processes a single struct field.
func (ep *ElementProcessor) processSingleStructField(field reflect.Value, fieldType reflect.StructField, fieldIdx int, config *ProcessingConfig) error {

	typ := field.FieldByName("Type").String()

	elementType := ElementType(typ)
	handler, exists := ep.handlers[elementType]
	if !exists {
		return fmt.Errorf("no handler found for element type: %s", elementType)
	}

	info := FieldInfo{
		name:       fieldType.Name,
		typ:        typ,
		hash:       field.FieldByName("Hash").String(),
		hashBefore: ep.value.Instance.Field(fieldIdx - 1).FieldByName("Hash").String(),
		nextHash:   ep.nextHash(fieldIdx), // The index is NOT always 0
		element:    matchElementType(fieldType.Name),
	}

	return handler.Handle(0, info, config)

}

// nextHash returns the next hash value
// Note:
func (ep *ElementProcessor) nextHash(fieldIdx int) string {
	if fieldIdx+1 < ep.value.InstanceNumField {
		return ep.value.Instance.Field(fieldIdx + 1).FieldByName("Hash").String()
	}
	return ""
}

// cleanup method
func (ep *ElementProcessor) Cleanup() {
	// Clean up handlers
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
