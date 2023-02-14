package definitions

import "reflect"

// ReflectCancelEventDefinitionMethodsToMap ...
func ReflectCancelEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(CancelEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCompensateEventDefinitionMethodsToMap ...
func ReflectCompensateEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(CompensateEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectConditionalEventDefinitionMethodsToMap ...
func ReflectConditionalEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ConditionalEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectErrorEventDefinitionMethodsToMap ...
func ReflectErrorEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ErrorEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEscalationEventDefinitionMethodsToMap ...
func ReflectEscalationEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(EscalationEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLinkEventDefinitionMethodsToMap ...
func ReflectLinkEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(LinkEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLinkEventDefinitionMethodsToMap ...
func ReflectMessageEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(MessageEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectSignalEventDefinitionMethodsToMap ...
func ReflectSignalEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(SignalEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTerminateEventDefinitionMethodsToMap ...
func ReflectTerminateEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(TerminateEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTimerEventDefinitionMethodsToMap ...
func ReflectTimerEventDefinitionMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(TimerEventDefinition{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
