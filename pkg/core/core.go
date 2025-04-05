package core

// go:build 1.23

import (
	"context"
	"fmt"
	"time"

	"github.com/deemount/gobpmn/pkg/core/common"
)

const DefaultTimeout = 1 * time.Second

// NewReflectDI ...
// If a error occurs, the function returns a new instance of the model and the error.
// The new instance is then a zero value of T by the model.
func NewReflectDI[T any, M common.BPMNGeneric](ctx context.Context, model T, genericType M) (result T, err error) {

	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()
	}

	reflectVal, err := common.NewReflectValue(model, genericType)
	if err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to create reflect value:\n%w", err))
	}
	defer reflectVal.Cleanup()

	if err := reflectVal.Setup(); err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to setup reflect value:\n%w", err))
	}

	// create a new mapping
	mapping := new(common.Mapping[M]) // NOTE: Mapping and Quantity is a helper structure and can be put together in a single structure?
	defer mapping.Cleanup()

	// assign the fields of the ReflectValue to the corresponding maps.
	mapping.Assign(reflectVal)

	// create a new quantity
	quantity := new(common.Quantity[M])

	// handle model type
	if err := reflectVal.HandleModelType(quantity, mapping); err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to handle model type:\n%w", err))
	}

	// reflect the processes in the BPMN model by given quantity.
	// v.Process[] is a slice of reflect.Value, which represents the processes in the BPMN model.
	if err := reflectVal.ReflectProcess(quantity); err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to reflect processes:\n%w", err))
	}

	// get the instance of the ReflectValue.
	instance := reflectVal.Instnc(mapping)
	typed, ok := instance.(T)
	if !ok {
		return *new(T), common.NewError(fmt.Errorf("invalid type conversion"))
	}

	// finalize the model construction
	if err := reflectVal.FinalizeModel(ctx, quantity); err != nil {
		return *new(T), common.NewError(fmt.Errorf("failed to finalize model:\n%w", err))
	}

	return typed, nil
}
