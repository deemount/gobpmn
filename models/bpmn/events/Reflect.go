package events

import (
	"reflect"
	"strings"
)

var (
	fieldShort = "vnts"
	fieldLong  = "events"
)

// IsDefinitions ...
func IsEvents(field reflect.StructField) bool {
	return strings.ToLower(field.Name) == fieldShort || strings.ToLower(field.Name) == fieldLong
}
