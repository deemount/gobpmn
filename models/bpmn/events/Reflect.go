package events

import (
	"reflect"
	"strings"
)

var (
	structEvents = "events"
)

// IsEvents ...
func IsEvents(field reflect.StructField) bool {
	return strings.ToLower(field.Name) == structEvents
}
