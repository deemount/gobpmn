package artifacts

import "reflect"

func ReflectTextGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Text{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

func ReflectTextMethodsToMap() map[int]string {
	var ptr *Text
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

func ReflectTextAnnotationGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(TextAnnotation{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

func ReflectTextAnnotationMethodsToMap() map[int]string {
	var ptr *TextAnnotation
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
