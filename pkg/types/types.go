package types

import (
	"reflect"
)

// BPMN ...
type BPMN struct {
	Pos, Width, Height, X, Y int
	Type, Hash               string
}

// OrderedField represents a field in a struct or map with its name, value, and metadata.
// Note: Actually not in use and will be implemented in the future.
type OrderedField struct {
	Name  string // field value or map key
	Value any    // field value or type (BPMN/reflect.StructField/etc.)
	Meta  any    // reflect.StructField or Tags
}

// BPMNGeneric is a generic interface that can be used to represent
// either a slice of reflect.StructField or a map[string]any.
type BPMNGeneric interface {
	[]OrderedField | []reflect.StructField | map[string]any
}
