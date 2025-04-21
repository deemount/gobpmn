package common

import (
	"reflect"

	"github.com/deemount/gobpmn/pkg/config"
)

// BPMN ...
type BPMN struct {
	Pos, Width, Height, X, Y int
	Type, Hash               string
}

/*
 * The following type is used to represent the generic types of BPMN struct fields and maps.
 * The interface is used in the common package to provide common functionality for BPMN elements.
 */

type BPMNStructField []reflect.StructField
type BPMNOrderedField []config.OrderedField
type BPMNMap map[string]any

// BPMNGeneric is a generic interface that can be used to represent
// either a slice of reflect.StructField or a map[string]any.
type BPMNGeneric interface {
	[]config.OrderedField | []reflect.StructField | map[string]any
}
