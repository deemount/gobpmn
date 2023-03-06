package core

import (
	"reflect"
	"strings"
)

var (
	repository = "DefinitionsRepository"
	fieldLong  = "definitions"
	fieldShort = "def"
)

// ReflectDefinitions ...
func ReflectDefinitions() reflect.Type {
	var def *DefinitionsRepository
	return reflect.TypeOf(def)
}

func ReflectDefinitionsGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Definitions{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

func ReflectDefinitionsMethodsToMap() map[int]string {
	var ptr *Definitions
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

func ReflectDefinitionsMethodsToSlice() []string {
	var ptr *Definitions
	t := reflect.TypeOf(ptr)
	m := []string{}
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectDefinitionsIsFirst ...
func ReflectDefinitionsIsFirst(p any) bool {
	if reflect.TypeOf(p).Field(0).Type.Name() == repository {
		return true
	}
	return false
}

// IsDefinitions ...
func IsDefinitions(field reflect.StructField) bool {
	return strings.ToLower(field.Name) == fieldShort || strings.ToLower(field.Name) == fieldLong
}

// IsNotDefinitions ...
func IsNotDefinitions(field reflect.StructField) bool {
	return strings.ToLower(field.Name) != fieldShort && strings.ToLower(field.Name) != fieldLong
}
