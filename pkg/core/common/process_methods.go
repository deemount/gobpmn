package common

import (
	"fmt"
	"reflect"
)

// ProcessMethod represents a BPMN process method with its argument
type ProcessMethod struct {
	Name string
	Arg  int
}

// GetProcessMethods returns the standard BPMN process methods with their quantities.
// Each method represents a BPMN element type and its required quantity.
// Note: q Quantity for a single process should be used in the Arg field.
func GetProcessMethods[M BPMNGeneric](processIdx int, q *Quantity[M]) []ProcessMethod {
	process := q.Elements[processIdx]
	return []ProcessMethod{
		{Name: "SetStartEvent", Arg: process[startEvent]},
		{Name: "SetEndEvent", Arg: process[endEvent]},
		{Name: "SetIntermediateCatchEvent", Arg: process[intermediateCatchEvent]},
		{Name: "SetIntermediateThrowEvent", Arg: process[intermediateThrowEvent]},
		{Name: "SetInclusiveGateway", Arg: process[inclusiveGateway]},
		{Name: "SetExclusiveGateway", Arg: process[exclusiveGateway]},
		{Name: "SetParallelGateway", Arg: process[parallelGateway]},
		{Name: "SetGateway", Arg: process[gateway]},
		{Name: "SetUserTask", Arg: process[userTask]},
		{Name: "SetServiceTask", Arg: process[serviceTask]},
		{Name: "SetScriptTask", Arg: process[scriptTask]},
		{Name: "SetTask", Arg: process[task]},
		{Name: "SetSequenceFlow", Arg: process[sequenceFlow]},
	}
}

// callMethod is a helper function to safely call reflect methods.
// It returns an error if the method is not found or if the method call fails,
// because there are more than one results and the first one is not nil.
func callMethod(target reflect.Value, methodName string, args []reflect.Value) error {

	method := target.MethodByName(methodName)
	if !method.IsValid() {
		return NewError(fmt.Errorf("method %s not found\n", methodName))
	}

	results := method.Call(args)
	if len(results) > 0 && !results[0].IsNil() {
		return NewError(fmt.Errorf("method %s failed: %v", methodName, results[0].Interface()))
	}

	return nil

}

// callMethodValue is a helper function to safely call reflect methods and return the result
// of the method call.
// It MUST return exactly one value.
// If it's less or more than one, it returns an error.
// It is mostly used to call a getter method from the elements.go file.
func callMethodValue(v reflect.Value, methodName string, args []reflect.Value) (reflect.Value, error) {
	var result reflect.Value

	method := v.MethodByName(methodName)
	if !method.IsValid() {
		return reflect.Value{}, NewError(fmt.Errorf("method %s not found", methodName))
	}

	results := method.Call(args)

	if len(results) == 0 {
		return reflect.Value{}, NewError(fmt.Errorf("method %s returned no value", methodName))
	}

	if len(results) > 1 && !results[1].IsNil() {
		return reflect.Value{}, NewError(fmt.Errorf("method %s failed", methodName))
	}

	result = results[0]

	return result, nil
}

// getMethodNames returns the method names of a target.
// NOTE: actually not in use.
func getMethodNames(target reflect.Value) []string {
	var methods []string
	for i := range target.NumMethod() {
		methods = append(methods, target.Type().Method(i).Name)
	}
	return methods
}

// targetMethodName returns the target method by the name.
// NOTE: actually not in use.
func targetMethodName(target reflect.Value, methodName string) (reflect.Value, error) {
	method := target.MethodByName(methodName)
	if !method.IsValid() {
		return reflect.Value{}, NewError(fmt.Errorf("method %s not found", methodName))
	}
	return method, nil
}
