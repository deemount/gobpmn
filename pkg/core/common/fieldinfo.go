package common

import (
	"maps"
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

	info := FieldInfo{
		name:    field.Type().Field(idx).Name,
		extName: extractLastTwoWords(field.Type().Field(idx).Name),
		element: matchElementType(field.Type().Field(idx).Name),
	}

	// Safely get the type if the field has one
	typeField := field.Field(idx).FieldByName("Type")
	if typeField.IsValid() {
		info.typ = typeField.String()
	}

	// Safely get the hash if the field has one
	hashField := field.Field(idx).FieldByName("Hash")
	if hashField.IsValid() {
		info.hash = hashField.String()
	}

	// Safely get the hashBefore only if idx > 0
	if idx > 0 {
		prevField := field.Field(idx - 1).FieldByName("Hash")
		if prevField.IsValid() {
			info.hashBefore = prevField.String()
		}
	}

	// Get the next hash
	info.nextHash = getNextHash(field, idx, field.NumField())

	return info

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

// collectFromFieldsWithNeighbors collects all fields with their neighbors
func collectFromFieldsWithNeighbors(data any) map[string]map[string]any {
	results := make(map[string]map[string]any)
	val := reflect.ValueOf(data)

	// pointer dereferencing
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Falls val kein Struct ist, abbrechen
	if val.Kind() != reflect.Struct {
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

		// Fall 1: Feld beginnt mit "From" => Verbindung speichern
		if strings.HasPrefix(fieldName, "From") {
			fieldInfo := make(map[string]any)
			fieldInfo["Value"] = fieldValue.Interface()

			// SourceRef bestimmen (vorheriges Feld)
			if i > 0 {
				sourceRefFieldName = val.Type().Field(i - 1).Name
				sourceRefFieldValue = val.Field(i - 1).Interface()
				fieldInfo["SourceRef"] = map[string]any{
					"FieldName": sourceRefFieldName,
					"Value":     sourceRefFieldValue,
				}
			}

			// TargetRef bestimmen (nächstes Feld, aber mit Sicherheit prüfen)
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

		// Fall 2: Rekursion auf Nested Structs
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
