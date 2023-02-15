package marker

import "reflect"

// ReflectCategoryGetMethodsToMap ...
func ReflectCategoryGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Category{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCategoryMethodsToMap ...
func ReflectCategoryMethodsToMap() map[int]string {
	var ptr *Category
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCategoryValueMethodsToMap ...
func ReflectCategoryValueGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(CategoryValue{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCategoryValueMethodsToMap ...
func ReflectCategoryValueMethodsToMap() map[int]string {
	var ptr *CategoryValue
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectGroupMethodsToMap ...
func ReflectGroupGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Group{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectGroupMethodsToMap ...
func ReflectGroupMethodsToMap() map[int]string {
	var ptr *Group
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectIncomingGetMethodsToMap ...
func ReflectIncomingGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Incoming{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectIncomingMethodsToMap ...
func ReflectIncomingMethodsToMap() map[int]string {
	var ptr *Incoming
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectOutgoingMethodsToMap ...
func ReflectOutgoingGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Outgoing{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectOutgoingMethodsToMap ...
func ReflectOutgoingMethodsToMap() map[int]string {
	var ptr *Outgoing
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
