package main

import (
	"reflect"
	"strings"
)

// FieldInfo holds information about a field being processed
type FieldInfo struct {
	name string

	// extName is the last two words of the field name in a CamelCase string
	extName string

	// typ is the type of the BPMN element
	typ     string
	element processElement

	// hash is the current hash value of the BPMN element
	hash       string
	hashBefore string
	nextHash   string
}

// extractFieldInfo gathers all necessary field information.
func extractFieldInfo(field reflect.Value, idx int) FieldInfo {
	return FieldInfo{
		name:       field.Type().Field(idx).Name,
		extName:    extractLastTwoWords(field.Type().Field(idx).Name),
		typ:        field.Field(idx).FieldByName("Type").String(),
		element:    matchElementType(field.Type().Field(idx).Name),
		hash:       field.Field(idx).FieldByName("Hash").String(),
		hashBefore: field.Field(idx - 1).FieldByName("Hash").String(),
		nextHash:   getNextHash(field, idx, field.NumField()),
	}
}

// isValidField checks if a field is valid
func isValidField(info FieldInfo) bool {
	if info.name == "" {
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

// collectFromFieldsWithNeighbors collects information from fields with neighbors
// by recursively traversing the data structure.
func collectFromFieldsWithNeighbors(data interface{}) map[string]map[string]interface{} {

	results := make(map[string]map[string]interface{})
	val := reflect.ValueOf(data)

	// pointer dereferencing
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:

		var sourceRefFieldName string
		var sourceRefFieldValue interface{}

		var targetRefFieldName string
		var targetRefFieldValue interface{}

		// iterate over all fields
		for i := 0; i < val.NumField(); i++ {

			field := val.Type().Field(i)
			fieldValue := val.Field(i)
			fieldName := field.Name

			// find "From" fields
			if strings.HasPrefix(fieldName, "From") {

				fieldInfo := make(map[string]interface{})
				fieldInfo["Value"] = fieldValue.Interface()

				// get previous field (if any)
				if i > 0 {

					sourceRefFieldName = val.Type().Field(i - 1).Name
					sourceRefFieldValue = val.Field(i - 1).Interface()
					fieldInfo["SourceRef"] = map[string]interface{}{
						"FieldName": sourceRefFieldName,
						"Value":     sourceRefFieldValue,
					}

					targetRefFieldName = val.Type().Field(i + 1).Name
					targetRefFieldValue = val.Field(i + 1).Interface()
					fieldInfo["TargetRef"] = map[string]interface{}{
						"FieldName": targetRefFieldName,
						"Value":     targetRefFieldValue,
					}

				}

				// save field information
				results[fieldName] = fieldInfo

			}

		}
	}

	return results
}
