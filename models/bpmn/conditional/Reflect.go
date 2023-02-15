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
