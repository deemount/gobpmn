package loop

import "reflect"

// ReflectLoopCardinalityGetMethodsToMap ...
func ReflectLoopCardinalityGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(LoopCardinality{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLoopCardinalityMethodsToMap ...
func ReflectLoopCardinalityMethodsToMap() map[int]string {
	var ptr *LoopCardinality
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLoopCardinalityMethodsToSlice ...
func ReflectLoopCardinalityMethodsToSlice() []string {
	var ptr *LoopCardinality
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectMultiInstanceLoopCharacteristicsGetMethodsToMap ...
func ReflectMultiInstanceLoopCharacteristicsGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(MultiInstanceLoopCharacteristics{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMultiInstanceLoopCharacteristicsMethodsToMap ...
func ReflectMultiInstanceLoopCharacteristicsMethodsToMap() map[int]string {
	var ptr *MultiInstanceLoopCharacteristics
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectMultiInstanceLoopCharacteristicsMethodsToSlice ...
func ReflectMultiInstanceLoopCharacteristicsMethodsToSlice() []string {
	var ptr *MultiInstanceLoopCharacteristics
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectParticipantMultiplicityMethodsToMap ...
func ReflectParticipantMultiplicityGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(ParticipantMultiplicity{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectParticipantMultiplicityMethodsToMap ...
func ReflectParticipantMultiplicityMethodsToMap() map[int]string {
	var ptr *ParticipantMultiplicity
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectParticipantMultiplicityMethodsToSlice ...
func ReflectParticipantMultiplicityMethodsToSlice() []string {
	var ptr *ParticipantMultiplicity
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectStandardLoopCharacteristicsGetMethodsToMap ...
func ReflectStandardLoopCharacteristicsGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(StandardLoopCharacteristics{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectStandardLoopCharacteristicsMethodsToMap ...
func ReflectStandardLoopCharacteristicsMethodsToMap() map[int]string {
	var ptr *StandardLoopCharacteristics
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectStandardLoopCharacteristicsMethodsToSlice ...
func ReflectStandardLoopCharacteristicsMethodsToSlice() []string {
	var ptr *StandardLoopCharacteristics
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
