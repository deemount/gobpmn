package data

import "reflect"

// ReflectDataObjectMethodsToMap ...
func ReflectDataObjectGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(DataObject{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDataObjectMethodsToMap ...
func ReflectDataObjectMethodsToMap() map[int]string {
	var ptr *DataObject
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDataObjectReferenceMethodsToMap ...
func ReflectDataObjectReferenceGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(DataObjectReference{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDataObjectReferenceMethodsToMap ...
func ReflectDataObjectReferenceMethodsToMap() map[int]string {
	var ptr *DataObjectReference
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDataStoreReferenceGetMethodsToMap ...
func ReflectDataStoreReferenceGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(DataStoreReference{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDataStoreReferenceMethodsToMap ...
func ReflectDataStoreReferenceMethodsToMap() map[int]string {
	var ptr *DataStoreReference
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
