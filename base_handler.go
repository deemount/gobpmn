package main

import "reflect"

type ElementHandler interface {
	Handle(processIdx int, info FieldInfo, config *ProcessingConfig) error
}

type PropertySetter interface {
	SetProperties(element reflect.Value, info FieldInfo) error
}
