package process

import (
	"reflect"
)

// ReflectProcessMethodsToMap ...
func ReflectProcessGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Process{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectProcessMethodsToMap ...
func ReflectProcessMethodsToMap() map[int]string {
	var ptr *Process
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectProcessMethodsToSlice ...
func ReflectProcessMethodsToSlice() []string {
	var ptr *Process
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
