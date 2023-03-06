package definitions

import "reflect"

// ReflectCancelEventDefinitionGetMethodsToMap ...
func ReflectCancelEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(CancelEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCancelEventDefinitionMethodsToMap ...
func ReflectCancelEventDefinitionMethodsToMap() map[int]string {
	var ptr *CancelEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCancelEventDefinitionMethodsToSlice ...
func ReflectCancelEventDefinitionMethodsToSlice() []string {
	var ptr *CancelEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectCompensateEventDefinitionGetMethodsToMap ...
func ReflectCompensateEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(CompensateEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCompensateEventDefinitionMethodsToMap ...
func ReflectCompensateEventDefinitionMethodsToMap() map[int]string {
	var ptr *CompensateEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCompensateEventDefinitionMethodsToSlice ...
func ReflectCompensateEventDefinitionMethodsToSlice() []string {
	var ptr *CompensateEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectConditionalEventDefinitionGetMethodsToMap ...
func ReflectConditionalEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ConditionalEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectConditionalEventDefinitionMethodsToMap ...
func ReflectConditionalEventDefinitionMethodsToMap() map[int]string {
	var ptr *ConditionalEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectConditionalEventDefinitionMethodsToSlice ...
func ReflectConditionalEventDefinitionMethodsToSlice() []string {
	var ptr *ConditionalEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectErrorEventDefinitionGetMethodsToMap ...
func ReflectErrorEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ErrorEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectErrorEventDefinitionMethodsToMap ...
func ReflectErrorEventDefinitionMethodsToMap() map[int]string {
	var ptr *ErrorEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectErrorEventDefinitionMethodsToSlice ...
func ReflectErrorEventDefinitionMethodsToSlice() []string {
	var ptr *ErrorEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectEscalationEventDefinitionMethodsToMap ...
func ReflectEscalationEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(EscalationEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEscalationEventDefinitionMethodsToMap ...
func ReflectEscalationEventDefinitionMethodsToMap() map[int]string {
	var ptr *EscalationEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEscalationEventDefinitionMethodsToSlice ...
func ReflectEscalationEventDefinitionMethodsToSlice() []string {
	var ptr *EscalationEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectLinkEventDefinitionGetMethodsToMap ...
func ReflectLinkEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(LinkEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLinkEventDefinitionMethodsToMap ...
func ReflectLinkEventDefinitionMethodsToMap() map[int]string {
	var ptr *LinkEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLinkEventDefinitionMethodsToSlice ...
func ReflectLinkEventDefinitionMethodsToSlice() []string {
	var ptr *LinkEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectMessageEventDefinitionGetMethodsToMap ...
func ReflectMessageEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(MessageEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMessageEventDefinitionMethodsToMap ...
func ReflectMessageEventDefinitionMethodsToMap() map[int]string {
	var ptr *MessageEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMessageEventDefinitionMethodsToSlice ...
func ReflectMessageEventDefinitionMethodsToSlice() []string {
	var ptr *MessageEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectSignalEventDefinitionGetMethodsToMap ...
func ReflectSignalEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(SignalEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSignalEventDefinitionMethodsToMap ...
func ReflectSignalEventDefinitionMethodsToMap() map[int]string {
	var ptr *SignalEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSignalEventDefinitionMethodsToSlice ...
func ReflectSignalEventDefinitionMethodsToSlice() []string {
	var ptr *SignalEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectTerminateEventDefinitionMethodsToMap ...
func ReflectTerminateEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(TerminateEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTerminateEventDefinitionMethodsToMap ...
func ReflectTerminateEventDefinitionMethodsToMap() map[int]string {
	var ptr *TerminateEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTerminateEventDefinitionMethodsToSlice ...
func ReflectTerminateEventDefinitionMethodsToSlice() []string {
	var ptr *TerminateEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectTimerEventDefinitionGetMethodsToMap ...
func ReflectTimerEventDefinitionGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(TimerEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTimerEventDefinitionMethodsToMap ...
func ReflectTimerEventDefinitionMethodsToMap() map[int]string {
	var ptr *TimerEventDefinition
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTimerEventDefinitionMethodsToSlice ...
func ReflectTimerEventDefinitionMethodsToSlice() []string {
	var ptr *TimerEventDefinition
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
