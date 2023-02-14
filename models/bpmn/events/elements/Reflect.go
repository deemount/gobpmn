package elements

import (
	"reflect"
	"strings"
)

// ReflectBoundaryEventMethodsToMap ...
func ReflectBoundaryEventMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(BoundaryEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectStartEventMethodsToMap ...
func ReflectIntermediateCatchEventMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(IntermediateCatchEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEndEventMethodsToMap ...
func ReflectEndEventMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(EndEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMessageMethodsToMap ...
func ReflectMessageMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Message{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSignalMethodsToMap ...
func ReflectSignalMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Signal{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectStartEventMethodsToMap ...
func ReflectStartEventlMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(StartEvent{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// IsMessage ...
func IsMessage(field string) bool { return strings.ToLower(field) == structMessage }
