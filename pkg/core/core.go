package core

import (
	"context"
	"fmt"

	"github.com/deemount/gobpmn/pkg/core/common"
	"github.com/deemount/gobpmn/pkg/types"
)

// Core is a generic function, acting as valve, that creates a new BPMN model from a
// static struct definition and a generic or typed map.
// It takes a context, a model of type T, and a generic type M.
// If a error occurs, the function returns a new instance of the model and the error.
// The new instance is then a zero value of T by the model.
func Core[T any, M types.BPMNGeneric](ctx context.Context, model T, genericType M) (result T, err error) {

	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), common.DefaultTimeout)
		defer cancel()
	}

	// create a new reflect value
	value, err := common.NewReflectValue(model, genericType)
	if err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to create reflect value:\n%w", err))
	}
	defer value.Cleanup()

	if err := value.Setup(); err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to setup reflect value:\n%w", err))
	}

	// create a new mapping
	mapping := new(common.Mapping[M]) // NOTE: Mapping and Quantity is a helper structure and can be put together in a single structure?
	defer mapping.Cleanup()

	// assign the fields of the value to the corresponding maps.
	mapping.Assign(value)

	// create a new quantity
	quantity := new(common.Quantity[M])

	// handle model type
	if err := value.HandleModelType(quantity, mapping); err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to handle model type:\n%w", err))
	}

	// reflect the processes in the bpmnmodel by given quantity.
	// v.Process[] is a slice of reflect.Value, which represents the processes in the bpmn model.
	if err := value.ReflectProcess(quantity); err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to reflect processes:\n%w", err))
	}

	// get the current value of the reflected value.
	instance := value.CurrentValue(mapping)
	// convert the instance to the type T
	// NOTE: this is a type assertion, which checks if the instance is of type T.
	// S1020: type assertion error
	// https://staticcheck.dev/docs/checks#S1020
	convertedInstance, ok := instance.(T) // NOTE: instance is a reflect.Value, which represents the current value of the reflected value.
	// check if the instance is of type T
	// if not, return a new instance of T and an error.
	// if ok is false, it means that the type conversion failed.
	// if ok is true, it means that the type conversion succeeded.
	if !ok {
		return *new(T), common.NewError(fmt.Errorf("type assertion error"))
	}

	// finalize the model construction
	if err := value.FinalizeModel(ctx, quantity); err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to finalize model:\n%w", err))
	}

	return convertedInstance, nil
}
