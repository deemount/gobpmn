package common

import "reflect"

// BPMN ...
type BPMN struct {
	Pos, Width, Height, X, Y int
	Type, Hash               string
}

/*
 * The following types are used to represent the generic types of BPMN struct fields and maps.
 * They are used in the common package to provide common functionality for BPMN elements.
 * NOTE: Actually, the types aren't used.
 */

// bstruct type represents the generic type of BPMN struct fields
type bstruct []reflect.StructField

// bmap type represents the generic type of BPMN map fields
type bmap map[string]any
