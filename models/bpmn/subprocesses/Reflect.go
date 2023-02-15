package subprocesses

import "reflect"

// ReflectAdHocSubProcessGetMethodsToMap ...
func ReflectAdHocProcessGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(AdHocSubProcess{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectAdHocSubProcessMethodsToMap ...
func ReflectAdHocProcessMethodsToMap() map[int]string {
	var ptr *AdHocSubProcess
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCallActivityGetMethodsToMap ...
func ReflectCallActivityGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(CallActivity{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCallActivityMethodsToMap ...
func ReflectCallActivityMethodsToMap() map[int]string {
	var ptr *CallActivity
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSubProcessMethodsToMap ...
func ReflectSubProcessGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(SubProcess{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSubProcessMethodsToMap ...
func ReflectSubProcessMethodsToMap() map[int]string {
	var ptr *SubProcess
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTransactionGetMethodsToMap ...
func ReflectTransactionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Transaction{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTransactionMethodsToMap ...
func ReflectTransactionMethodsToMap() map[int]string {
	var ptr *Transaction
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
