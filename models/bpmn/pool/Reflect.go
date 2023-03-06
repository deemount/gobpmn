package pool

import "reflect"

// ReflectFlowNodeRefMethodsToMap ...
func ReflectFlowNodeRefGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(FlowNodeRef{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectFlowNodeRefMethodsToMap ...
func ReflectFlowNodeRefMethodsToMap() map[int]string {
	var ptr *FlowNodeRef
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectFlowNodeRefMethodsToSlice ...
func ReflectFlowNodeRefMethodsToSlice() []string {
	var ptr *FlowNodeRef
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectLaneMethodsToMap ...
func ReflectLaneGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Lane{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLaneMethodsToMap ...
func ReflectLaneMethodsToMap() map[int]string {
	var ptr *Lane
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLaneMethodsToSlice ...
func ReflectLaneMethodsToSlice() []string {
	var ptr *Lane
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectLaneSetGetMethodsToMap ...
func ReflectLaneSetGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(LaneSet{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLaneSetMethodsToMap ...
func ReflectLaneSetMethodsToMap() map[int]string {
	var ptr *LaneSet
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLaneSetMethodsToSlice ...
func ReflectLaneSetMethodsToSlice() []string {
	var ptr *LaneSet
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
