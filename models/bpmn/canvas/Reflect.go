package canvas

import "reflect"

// ReflectBoundsMethodsToMap ...
func ReflectBoundsMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Bounds{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDiagramMethodsToMap ...
func ReflectDiagramMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Diagram{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEdgeMethodsToMap ...
func ReflectEdgeMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Edge{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLabelMethodsToMap ...
func ReflectLabelMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Label{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectPlaneMethodsToMap ...
func ReflectPlaneMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Plane{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectShapeMethodsToMap ...
func ReflectShapeMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Shape{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectWaypointMethodsToMap ...
func ReflectWaypointMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Waypoint{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
