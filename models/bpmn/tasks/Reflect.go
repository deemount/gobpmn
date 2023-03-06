package tasks

import "reflect"

// ReflectBusinessRuleTaskGetMethodsToMap ...
func ReflectBusinessRuleTaskGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(BusinessRuleTask{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectBusinessRuleTaskMethodsToMap ...
func ReflectBusinessRuleTaskMethodsToMap() map[int]string {
	var ptr *BusinessRuleTask
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectBusinessRuleTaskMethodsToSlice ...
func ReflectBusinessRuleTaskMethodsToSlice() []string {
	var ptr *BusinessRuleTask
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectManualTaskGetMethodsToMap ...
func ReflectManualTaskGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ManualTask{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectManualTaskMethodsToMap ...
func ReflectManualTaskMethodsToMap() map[int]string {
	var ptr *ManualTask
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectManualTaskMethodsToSlice ...
func ReflectManualTaskMethodsToSlice() []string {
	var ptr *ManualTask
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectReceiveTaskGetMethodsToMap ...
func ReflectReceiveTaskGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ReceiveTask{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectReceiveTaskMethodsToMap ...
func ReflectReceiveTaskMethodsToMap() map[int]string {
	var ptr *ReceiveTask
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectReceiveTaskMethodsToSlice ...
func ReflectReceiveTaskMethodsToSlice() []string {
	var ptr *ReceiveTask
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectScriptTaskGetMethodsToMap ...
func ReflectScriptTaskGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ScriptTask{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectScriptTaskMethodsToMap ...
func ReflectScriptTaskMethodsToMap() map[int]string {
	var ptr *ScriptTask
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectScriptTaskMethodsToSlice...
func ReflectScriptTaskMethodsToSlice() []string {
	var ptr *ScriptTask
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectSendTaskGetMethodsToMap ...
func ReflectSendTaskGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(SendTask{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSendTaskMethodsToMap ...
func ReflectSendTaskMethodsToMap() map[int]string {
	var ptr *SendTask
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSendTaskMethodsToSlice ...
func ReflectSendTaskMethodsToSlice() []string {
	var ptr *SendTask
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectServiceTaskGetMethodsToMap ...
func ReflectServiceTaskGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ServiceTask{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectServiceTaskMethodsToMap ...
func ReflectServiceTaskMethodsToMap() map[int]string {
	var ptr *ServiceTask
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectServiceTaskMethodsToSlice ...
func ReflectServiceTaskMethodsToSlice() []string {
	var ptr *ServiceTask
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectTaskMethodsToMap ...
func ReflectTaskGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Task{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTaskMethodsToMap ...
func ReflectTaskMethodsToMap() map[int]string {
	var ptr *Task
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTaskMethodsToSlice ...
func ReflectTaskMethodsToSlice() []string {
	var ptr *Task
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectUserTaskMethodsToMap ...
func ReflectUserTaskGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(UserTask{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectUserTaskMethodsToMap ...
func ReflectUserTaskMethodsToMap() map[int]string {
	var ptr *UserTask
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectUserTaskMethodsToSlice ...
func ReflectUserTaskMethodsToSlice() []string {
	var ptr *UserTask
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
