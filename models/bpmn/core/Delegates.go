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

// SetMainElements ...
func SetMainElements(d DefinitionsRepository, num int) {
	d.SetMainElements(num)
}
