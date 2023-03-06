package flow

import "reflect"

// ReflectAssociationyGetMethodsToMap ...
func ReflectAssociationGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Association{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectAssociationyMethodsToMap ...
func ReflectAssociationMethodsToMap() map[int]string {
	var ptr *Association
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectAssociationyMethodsToSlice ...
func ReflectAssociationMethodsToSlice() []string {
	var ptr *Association
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectDataInputAssociationyGetMethodsToMap ...
func ReflectDataInputAssociationGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(DataInputAssociation{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDataInputAssociationyMethodsToMap ...
func ReflectDataInputAssociationMethodsToMap() map[int]string {
	var ptr *DataInputAssociation
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDataInputAssociationyMethodsToSlice ...
func ReflectDataInputAssociationMethodsToSlice() []string {
	var ptr *DataInputAssociation
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectMessageFlowMethodsToMap ...
func ReflectMessageFlowGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(MessageFlow{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMessageFlowMethodsToMap ...
func ReflectMessageFlowMethodsToMap() map[int]string {
	var ptr *MessageFlow
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMessageFlowMethodsToSlice ...
func ReflectMessageFlowMethodsToSlice() []string {
	var ptr *MessageFlow
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectSequenceFlowGetMethodsToMap ...
func ReflectSequenceFlowGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(SequenceFlow{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSequenceFlowMethodsToMap ...
func ReflectSequenceFlowMethodsToMap() map[int]string {
	var ptr *SequenceFlow
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSequenceFlowMethodsToSlice ...
func ReflectSequenceFlowMethodsToSlice() []string {
	var ptr *SequenceFlow
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
