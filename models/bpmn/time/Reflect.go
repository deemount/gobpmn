package time

import "reflect"

// ReflectTimeCycleGetMethodsToMap ...
func ReflectTimeCycleGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(TimeCycle{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTimeCycleMethodsToMap ...
func ReflectTimeCycleMethodsToMap() map[int]string {
	var ptr *TimeCycle
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTimeDateGetMethodsToMap ...
func ReflectTimeDateGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(TimeDate{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTimeDateMethodsToMap ...
func ReflectTimeDateMethodsToMap() map[int]string {
	var ptr *TimeDate
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTimeDurationGetMethodsToMap ...
func ReflectTimeDurationGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(TimeDuration{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectTimeDurationMethodsToMap ...
func ReflectTimeDurationMethodsToMap() map[int]string {
	var ptr *TimeDuration
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
