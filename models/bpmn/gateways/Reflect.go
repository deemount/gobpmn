package gateways

import "reflect"

// ReflectComplexGatewayGetMethodsToMap ...
func ReflectComplexGatewayGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ComplexGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectComplexGatewayMethodsToMap ...
func ReflectComplexGatewayMethodsToMap() map[int]string {
	var ptr *ComplexGateway
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectComplexGatewayMethodsToSlice ...
func ReflectComplexGatewayMethodsToSlice() []string {
	var ptr *ComplexGateway
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectEventBasedGatewayGetMethodsToMap ...
func ReflectEventBasedGatewayGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(EventBasedGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEventBasedGatewayMethodsToMap ...
func ReflectEventBasedGatewayMethodsToMap() map[int]string {
	var ptr *EventBasedGateway
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEventBasedGatewayMethodsToSlice ...
func ReflectEventBasedGatewayMethodsToSlice() []string {
	var ptr *EventBasedGateway
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectExclusiveGatewayGetMethodsToMap ...
func ReflectExclusiveGatewayGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(EventBasedGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectExclusiveGatewayMethodsToMap ...
func ReflectExclusiveGatewayMethodsToMap() map[int]string {
	var ptr *EventBasedGateway
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectExclusiveGatewayMethodsToSlice ...
func ReflectExclusiveGatewayMethodsToSlice() []string {
	var ptr *EventBasedGateway
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectInclusiveGatewayMethodsToMap ...
func ReflectInclusiveGatewayGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(InclusiveGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectInclusiveGatewayMethodsToMap ...
func ReflectInclusiveGatewayMethodsToMap() map[int]string {
	var ptr *InclusiveGateway
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectInclusiveGatewayMethodsToSlice ...
func ReflectInclusiveGatewayMethodsToSlice() []string {
	var ptr *InclusiveGateway
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectParallelGatewayGetMethodsToMap ...
func ReflectParallelGatewayGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ParallelGateway{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectParallelGatewayMethodsToMap ...
func ReflectParallelGatewayMethodsToMap() map[int]string {
	var ptr *ParallelGateway
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectParallelGatewayMethodsToSlice ...
func ReflectParallelGatewayMethodsToSlice() []string {
	var ptr *ParallelGateway
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
