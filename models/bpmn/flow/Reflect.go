package flow

import "reflect"

// ReflectAssociationyMethodsToMap ...
func ReflectAssociationMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Association{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDataInputAssociationyMethodsToMap ...
func ReflectDataInputAssociationMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(DataInputAssociation{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMessageFlowMethodsToMap ...
func ReflectMessageFlowMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(MessageFlow{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSequenceFlowMethodsToMap ...
func ReflectSequenceFlowMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(SequenceFlow{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
