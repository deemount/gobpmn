package conditional

import "reflect"

// ReflectCompletionConditionGetMethodsToMap ...
func ReflectCompletionConditionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(CompletionCondition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCompletionConditionMethodsToMap ...
func ReflectCompletionConditionMethodsToMap() map[int]string {
	var ptr *CompletionCondition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCompletionConditionMethodsToSlice ...
func ReflectCompletionConditionMethodsToSlice() []string {
	var ptr *CompletionCondition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectConditionGetMethodsToMap ...
func ReflectConditionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Condition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectConditionMethodsToMap ...
func ReflectConditionMethodsToMap() map[int]string {
	var ptr *Condition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectConditionMethodsToSlice ...
func ReflectConditionMethodsToSlice() []string {
	var ptr *Condition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectConditionExpressionGetMethodsToMap ...
func ReflectConditionExpressionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ConditionExpression{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectConditionExpressionMethodsToMap ...
func ReflectConditionExpressionMethodsToMap() map[int]string {
	var ptr *ConditionExpression
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectConditionExpressionMethodsToSlice ...
func ReflectConditionExpressionMethodsToSlice() []string {
	var ptr *ConditionExpression
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
