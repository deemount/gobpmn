package common

import (
	"fmt"
	"reflect"
)

// BaseElement provides common functionality for element handlers
type BaseElement struct {
	processor *ElementProcessor
}

// Implement PropertySetter interface at the base level
func (b *BaseElement) SetProperties(element reflect.Value, info FieldInfo) error {
	// Common property setting logic
	if err := b.setCommonProperties(element, info); err != nil {
		return fmt.Errorf("failed to set common properties: %w", err)
	}
	return nil
}

// setElementProperties sets common properties for BPMN elements.
func (h *BaseElement) setCommonProperties(el reflect.Value, info FieldInfo) error {

	if err := callMethod(el, "SetID", []reflect.Value{
		reflect.ValueOf(info.typ),
		reflect.ValueOf(info.hash),
	}); err != nil {
		return fmt.Errorf("failed to set ID: %w", err)
	}

	if err := callMethod(el, "SetName", []reflect.Value{
		reflect.ValueOf(info.name),
	}); err != nil {
		return fmt.Errorf("failed to set name: %w", err)
	}

	return nil

}

func (h *BaseElement) Cleanup() {
	// Cleanup resources
}
