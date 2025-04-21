package common

import (
	"fmt"
	"log"
	"maps"
	"reflect"
	"sort"
	"strings"

	"github.com/deemount/gobpmn/pkg/config"
)

// extractOrderedFields extracts ordered fields from a struct or map.
func extractOrderedFields[T any](model T) ([]config.OrderedField, error) {

	typeOf := reflect.TypeOf(model)
	valueOf := reflect.ValueOf(model)

	switch typeOf.Kind() {
	case reflect.Struct:
		fields := reflect.VisibleFields(typeOf)
		ordered := make([]config.OrderedField, 0, len(fields))
		for _, field := range fields {
			val := valueOf.FieldByIndex(field.Index)
			ordered = append(ordered, config.OrderedField{
				Name:  field.Name,
				Value: val.Interface(),
				Meta:  field,
			})
		}
		return ordered, nil

	case reflect.Map:
		fields := make(map[string]any)
		var sortedKeys []string
		bpmnElements := make([]struct {
			Key string
			Pos int
		}, 0)
		for _, key := range valueOf.MapKeys() {
			strKey := key.String()
			value := valueOf.MapIndex(key).Interface()
			switch v := value.(type) {
			case BPMN:
				bpmnElements = append(bpmnElements, struct {
					Key string
					Pos int
				}{strKey, v.Pos})
				fields[strKey] = v
			case bool, any:
				fields[strKey] = v
			}
		}
		sort.SliceStable(bpmnElements, func(i, j int) bool {
			return bpmnElements[i].Pos < bpmnElements[j].Pos
		})
		for _, elem := range bpmnElements {
			sortedKeys = append(sortedKeys, elem.Key)
		}
		ordered := make([]config.OrderedField, 0, len(sortedKeys))
		for _, key := range sortedKeys {
			ordered = append(ordered, config.OrderedField{
				Name:  key,
				Value: fields[key],
				Meta:  nil,
			})
		}
		return ordered, nil
	default:
		return nil, fmt.Errorf("unsupported kind: %s", typeOf.Kind())
	}
}

// FieldInfo holds information about a field being processed
type FieldInfo struct {
	Name string

	// used for converter
	Pos   int // position of the field in the struct (maybe redudant)
	Type  reflect.Type
	Value any

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

// instanceFieldName returns the field value by the name in v.Instance.
func instanceFieldName[M BPMNGeneric](v *ReflectValue[M], name string) reflect.Value {
	value := v.Instance.FieldByName(name)
	if !value.IsValid() {
		fmt.Printf("instanceFieldName: field %s not found in struct instance. Available fields: %v\n", name, getFieldNames(v.Instance))
		return reflect.Value{}
	}
	return value
}

// getFieldNames returns the names of all fields in a reflect.Value instance
func getFieldNames(instance reflect.Value) []string {
	var fieldNames []string
	for i := range instance.NumField() {
		fieldNames = append(fieldNames, instance.Type().Field(i).Name)
	}
	return fieldNames
}
