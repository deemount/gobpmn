package common

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/deemount/gobpmn/pkg/core/foundation"
)

// Converter is a struct
type Converter struct{}

// DynamicStruct holds the generated type and value
type DynamicStruct struct {
	Type   reflect.Type
	Value  reflect.Value
	Fields []FieldInfo
}

// ToStruct converts a map[string]any to a struct
func (c *Converter) ToStruct(input any) (*DynamicStruct, error) {
	// convert input to map[string]interface{}
	inputMap, err := c.normalizeToMap(input)
	if err != nil {
		return nil, err
	}
	// extract and order fields
	fields, err := c.getOrderedFields(inputMap)
	if err != nil {
		return nil, err
	}
	// generate struct type dynamically
	structType := c.createStructType(fields)
	// create new instance and populate values
	structValue := reflect.New(structType).Elem()
	if err := c.populateStruct(structValue, fields); err != nil {
		return nil, err
	}
	return &DynamicStruct{
		Type:   structType,
		Value:  structValue,
		Fields: fields,
	}, nil
}

// normalizeToMap converts the input to a map[string]any
func (c *Converter) normalizeToMap(input any) (map[string]any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, fmt.Errorf("input must be a map")
	}
	if val.Type().Key().Kind() != reflect.String {
		return nil, fmt.Errorf("map keys must be strings")
	}
	result := make(map[string]any)
	for _, key := range val.MapKeys() {
		result[key.String()] = val.MapIndex(key).Interface()
	}
	return result, nil
}

// getOrderedFields extracts and orders fields from a map
func (c *Converter) getOrderedFields(m map[string]any) ([]FieldInfo, error) {
	var fields []FieldInfo
	for name, value := range m {
		pos := 0
		if bpmn, ok := value.(BPMN); ok {
			pos = bpmn.Pos
		}
		fields = append(fields, FieldInfo{
			Name:  name,
			Pos:   pos,
			Type:  reflect.TypeOf(value),
			Value: value,
		})
	}
	// sort by pos
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Pos < fields[j].Pos
	})
	return fields, nil
}

// createStructType generates a struct type from a list of fields
func (c *Converter) createStructType(fields []FieldInfo) reflect.Type {
	var structFields []reflect.StructField
	for _, field := range fields {
		structFields = append(structFields, reflect.StructField{
			Name: field.Name,
			Type: field.Type,
			Tag:  reflect.StructTag(fmt.Sprintf(`map:"%s"`, field.Name)),
		})
	}
	return reflect.StructOf(structFields)
}

// populateStruct populates a struct with values from a list of fields
func (c *Converter) populateStruct(structValue reflect.Value, fields []FieldInfo) error {
	// First ensure fields are in correct order (Def, IsExecutable, then BPMN fields by Pos)
	sort.Slice(fields, func(i, j int) bool {
		// Def always comes first
		if fields[i].Name == "Def" {
			return true
		}
		if fields[j].Name == "Def" {
			return false
		}
		// IsExecutable comes second
		if fields[i].Name == "IsExecutable" {
			return true
		}
		if fields[j].Name == "IsExecutable" {
			return false
		}
		// Then sort BPMN fields by Pos
		_, iIsBPMN := fields[i].Value.(BPMN)
		_, jIsBPMN := fields[j].Value.(BPMN)
		if iIsBPMN && jIsBPMN {
			return fields[i].Pos < fields[j].Pos
		}
		// non-BPMN fields come before BPMN fields
		return !iIsBPMN && jIsBPMN
	})

	// now populate the struct fields in sorted order
	for i, field := range fields {
		fieldValue := structValue.Field(i)
		// handle BPMN fields with position preservation
		if bpmn, ok := field.Value.(BPMN); ok {
			fieldValue.Set(reflect.ValueOf(BPMN{Pos: bpmn.Pos}))
			continue
		}
		// handle special cases with proper type checking
		switch field.Name {
		case "Def":
			if defPtr, ok := field.Value.(*foundation.Definitions); ok {
				converted := convertDefinitions(defPtr)
				// check if converted is nil and assign a new Definitions
				if converted != nil && reflect.TypeOf(converted).AssignableTo(fieldValue.Type()) {
					fieldValue.Set(reflect.ValueOf(converted))
					continue
				}
			}
			return NewError(fmt.Errorf("Def must be *foundation.Definitions or convertible to DefinitionsRepository"))
		case "IsExecutable":
			if b, ok := field.Value.(bool); ok {
				fieldValue.SetBool(b)
				continue
			}
			return NewError(fmt.Errorf("IsExecutable must be bool, got %T", field.Value))
		}
		// normal field assignment with type conversion
		value := reflect.ValueOf(field.Value)
		if value.Type().ConvertibleTo(fieldValue.Type()) {
			fieldValue.Set(value.Convert(fieldValue.Type()))
		} else {
			return NewError(fmt.Errorf("cannot assign %v to %v", value.Type(), fieldValue.Type()))
		}
	}
	return nil
}

// convertDefinitions is a helper function
func convertDefinitions(def *foundation.Definitions) foundation.DefinitionsRepository {
	if def == nil {
		return foundation.NewDefinitions()
	}
	return def
}

const DefaultTimeout = 20 * time.Millisecond

// ConvertDynamicStructToDefinitions ...
func ConvertDynamicStructToDefinitions(ctx context.Context, process any) (*foundation.Definitions, error) {

	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, DefaultTimeout)
		defer cancel()
		return ConvertDynamicStructToDefinitions(ctx, process)
	}

	// assuming process contains dynamic struct
	value := reflect.ValueOf(process)
	// if process is an interface{}, get the underlying value
	if value.Kind() == reflect.Interface {
		value = value.Elem()
	}
	// get the Def field
	defField := value.FieldByName("Def")
	if !defField.IsValid() {
		fmt.Println("Def field not found")
		return nil, fmt.Errorf("Def field not found")
	}
	// convert to *foundation.Definitions
	def, ok := defField.Interface().(*foundation.Definitions)
	if !ok {
		return nil, fmt.Errorf("Def field is not of type *foundation.Definitions")
	}
	return def, nil
}
