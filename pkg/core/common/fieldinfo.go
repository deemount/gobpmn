package common

import (
	"log"
	"maps"
	"reflect"
	"strings"
)

// FieldInfo holds information about a field being processed
type FieldInfo struct {
	Name string

	// extName is the last two words of the field name in a CamelCase string
	extName string

	// typ is the type of the BPMN element
	typ     string
	element processElement

	// hash is the current hash value of the BPMN element
	hash       string
	hashBefore string
	nextHash   string

	// used for converter
	Type  reflect.Type
	Value any
	Pos   int
}

// extractFieldInfo gathers all necessary field information.
func extractFieldInfo(field reflect.Value, idx int) FieldInfo {

	info := FieldInfo{
		Name:    field.Type().Field(idx).Name,
		extName: extractLastTwoWords(field.Type().Field(idx).Name),
		element: matchElementType(field.Type().Field(idx).Name),
	}

	// safely get the type if the field has one
	typeField := field.Field(idx).FieldByName("Type")
	if typeField.IsValid() {
		info.typ = typeField.String()
	}

	// safely get the hash if the field has one
	hashField := field.Field(idx).FieldByName("Hash")
	if hashField.IsValid() {
		if info.element == "StartEvent" && typeField.String() == "startevent" {
			info.hash = "1"
		} else {
			info.hash = hashField.String()
		}
	}

	// safely get the hashBefore only if idx > 0
	if idx > 0 {
		prevField := field.Field(idx - 1).FieldByName("Hash")
		if prevField.IsValid() {
			info.hashBefore = prevField.String()
		}
	}

	// get the next hash
	info.nextHash = getNextHash(field, idx, field.NumField())

	return info

}

// isValidField checks if a field is valid
func isValidField(info FieldInfo) bool {
	if info.Name == "" {
		return false
	}
	if info.typ == "" {
		return false
	}
	if info.element == "" {
		return false
	}
	return true
}

// collectFromFieldsWithNeighbors collects all fields with their neighbors
func collectFromFieldsWithNeighbors(data any) map[string]map[string]any {
	results := make(map[string]map[string]any)
	val := reflect.ValueOf(data)

	// pointer dereferencing
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// cancel, if val is not a struct
	if val.Kind() != reflect.Struct {
		log.Printf("Data is not a struct: %v", val)
		return results
	}

	var sourceRefFieldName string
	var sourceRefFieldValue any

	var targetRefFieldName string
	var targetRefFieldValue any

	// iterate over all fields
	for i := range val.NumField() {
		field := val.Type().Field(i)
		fieldValue := val.Field(i)
		fieldName := field.Name

		// case 1: field starts with "From" => save the connection
		if strings.HasPrefix(fieldName, "From") {
			fieldInfo := make(map[string]any)
			fieldInfo["Value"] = fieldValue.Interface()

			// get SourceRef bestimmen (previous field)
			if i > 0 {
				sourceRefFieldName = val.Type().Field(i - 1).Name
				sourceRefFieldValue = val.Field(i - 1).Interface()
				fieldInfo["SourceRef"] = map[string]any{
					"FieldName": sourceRefFieldName,
					"Value":     sourceRefFieldValue,
				}
			}

			// get TargetRef (next field, if available)
			if i+1 < val.NumField() {
				targetRefFieldName = val.Type().Field(i + 1).Name
				targetRefFieldValue = val.Field(i + 1).Interface()
				fieldInfo["TargetRef"] = map[string]any{
					"FieldName": targetRefFieldName,
					"Value":     targetRefFieldValue,
				}
			}

			results[fieldName] = fieldInfo
		}

		// case 2: recoursion on nested Structs
		if fieldValue.Kind() == reflect.Struct {
			subResults := collectFromFieldsWithNeighbors(fieldValue.Interface())
			maps.Copy(results, subResults)
		} else if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() && fieldValue.Elem().Kind() == reflect.Struct {
			subResults := collectFromFieldsWithNeighbors(fieldValue.Elem().Interface())
			maps.Copy(results, subResults)
		}
	}

	return results
}
