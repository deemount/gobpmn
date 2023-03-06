package attributes

import "reflect"

// ReflectDocumentationGetMethodsToMap ...
func ReflectDocumentationGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Documentation{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDocumentationMethodsToMap ...
func ReflectDocumentationMethodsToMap() map[int]string {
	var ref *Documentation
	m := make(map[int]string)
	t := reflect.TypeOf(ref)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDocumentationMethodsToSlice ...
func ReflectDocumentationMethodsToSlice() []string {
	var ref *Documentation
	m := []string{}
	t := reflect.TypeOf(ref)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectExtensionElementsGetMethodsToMap ...
func ReflectExtensionElementsGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ExtensionElements{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectExtensionElementsMethodsToMap ...
func ReflectExtensionElementsMethodsToMap() map[int]string {
	var ref ExtensionElements
	m := make(map[int]string)
	t := reflect.TypeOf(ref)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectExtensionElementsMethodsToSlice ...
func ReflectExtensionElementsMethodsToSlice() []string {
	var ref ExtensionElements
	m := []string{}
	t := reflect.TypeOf(ref)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectPropertyGetMethodsToMap ...
func ReflectPropertyGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Property{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectPropertyMethodsToMap ...
func ReflectPropertyMethodsToMap() map[int]string {
	var ref *Property
	m := make(map[int]string)
	t := reflect.TypeOf(ref)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectPropertyMethodsToSlice ...
func ReflectPropertyMethodsToSlice() []string {
	var ref *Property
	m := []string{}
	t := reflect.TypeOf(ref)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
