package common

import (
	"fmt"
	"reflect"

	"github.com/deemount/gobpmn/pkg/types"
)

// BaseElement provides common functionality for element handlers
type BaseElement[M types.BPMNGeneric] struct {
	processor *ElementProcessor[M]
}

// SetProperties sets common properties for BPMN elements, like ID and Name.
func (b *BaseElement[M]) SetProperties(element reflect.Value, info FieldInfo) error {
	// Common property setting logic
	if err := b.setCommonProperties(element, info); err != nil {
		return fmt.Errorf("failed to set common properties: %w", err)
	}
	return nil
}

// setCommonProperties sets common properties for BPMN elements.
func (h *BaseElement[M]) setCommonProperties(el reflect.Value, info FieldInfo) error {

	if err := callMethod(el, "SetID", []reflect.Value{
		reflect.ValueOf(info.typ),
		reflect.ValueOf(info.hash),
	}); err != nil {
		return fmt.Errorf("failed to set ID: %w", err)
	}

	if err := callMethod(el, "SetName", []reflect.Value{
		reflect.ValueOf(info.Name),
	}); err != nil {
		return fmt.Errorf("failed to set name: %w", err)
	}

	return nil

}
