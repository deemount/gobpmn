package gateways

import "reflect"

// ReflectComplexGatewayMethodsToMap ...
func ReflectComplexGatewayMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ComplexGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEventBasedGatewayMethodsToMap ...
func ReflectEventBasedGatewayMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(EventBasedGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectExclusiveGatewayMethodsToMap ...
func ReflectExclusiveGatewayMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(EventBasedGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectInclusiveGatewayMethodsToMap ...
func ReflectInclusiveGatewayMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(InclusiveGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectParallelGatewayMethodsToMap ...
func ReflectParallelGatewayMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ParallelGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
