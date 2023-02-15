package elements

import (
	"reflect"
	"strings"
)

// ReflectBoundaryEventGetMethodsToMap ...
func ReflectBoundaryEventGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(BoundaryEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectBoundaryEventMethodsToMap ...
func ReflectBoundaryEventMethodsToMap() map[int]string {
	var ptr *BoundaryEvent
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectIntermediateCatchEventGetMethodsToMap ...
func ReflectIntermediateCatchEventGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(IntermediateCatchEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectIntermediateCatchEventMethodsToMap ...
func ReflectIntermediateCatchEventMethodsToMap() map[int]string {
	var ptr *IntermediateCatchEvent
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectIntermediateCatchEventGetMethodsToMap ...
func ReflectIntermediateThrowEventGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(IntermediateThrowEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectIntermediateCatchEventMethodsToMap ...
func ReflectIntermediateThrowEventMethodsToMap() map[int]string {
	var ptr *IntermediateThrowEvent
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEndEventGetMethodsToMap ...
func ReflectEndEventGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(EndEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEndEventMethodsToMap ...
func ReflectEndEventMethodsToMap() map[int]string {
	var ptr *EndEvent
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMessageMethodsToMap ...
func ReflectMessageGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Message{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMessageMethodsToMap ...
func ReflectMessageMethodsToMap() map[int]string {
	var ptr *Message
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSignalGetMethodsToMap ...
func ReflectSignalGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Signal{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSignalMethodsToMap ...
func ReflectSignalMethodsToMap() map[int]string {
	var ptr *Signal
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectStartEventGetMethodsToMap ...
func ReflectStartEventlGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(StartEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectStartEventMethodsToMap ...
func ReflectStartEventlMethodsToMap() map[int]string {
	var ptr *StartEvent
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// IsMessage ...
func IsMessage(field string) bool { return strings.ToLower(field) == structMessage }
