package common

// go:build 1.23

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

const DefaultTimeout = 1 * time.Second

// NewReflectDI ...
// If a error occurs, the function returns a new instance of the model and the error.
// The new instance is then a zero value of T by the model.
func NewReflectDI[T any](ctx context.Context, model T) (result T, err error) {

	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()
	}

	reflectVal, err := NewReflectValue(model)
	if err != nil {
		return *new(T), NewError(fmt.Errorf("failed to create reflect value"))
	}
	defer reflectVal.Cleanup()

	if err := reflectVal.initialize(); err != nil {
		return *new(T), NewError(fmt.Errorf("failed to initialize reflect value: %w", err))
	}

	// create a new mapping
	mapping := new(Mapping) // NOTE: Mapping and Quantity is a helper structure and can be put together in a single structure?
	defer mapping.Cleanup()

	// assign the fields of the ReflectValue to the corresponding maps.
	mapping.Assign(reflectVal)

	// create a new quantity
	quantity := new(Quantity)

	// handle model type
	if err := reflectVal.handleModelType(quantity, mapping); err != nil {
		return *new(T), NewError(fmt.Errorf("failed to handle model type: %v", err))
	}

	// reflect the processes in the BPMN model by given quantity.
	// v.Process[] is a slice of reflect.Value, which represents the processes in the BPMN model.
	if err := reflectVal.reflectProcess(quantity); err != nil {
		return *new(T), NewError(fmt.Errorf("failed to reflect processes: %v", err))
	}

	// get the instance of the ReflectValue.
	instance := reflectVal.instance(mapping)
	typed, ok := instance.(T)
	if !ok {
		return *new(T), NewError(fmt.Errorf("invalid type conversion"))
	}

	// finalize the model construction
	if err := reflectVal.finalizeModel(ctx, quantity); err != nil {
		return *new(T), NewError(fmt.Errorf("failed to finalize model:\n%v", err))
	}

	return typed, nil
}

// instanceFieldName returns the field value by the name in v.Instance.
func instanceFieldName(v *ReflectValue, name string) reflect.Value {
	value := v.Instance.FieldByName(name)
	if !value.IsValid() {
		fmt.Printf("targetFieldName: field %s not found in instance. Available fields: %v\n", name, getFieldNames(v.Instance))
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
