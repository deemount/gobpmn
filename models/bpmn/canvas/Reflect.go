package canvas

import "reflect"

// ReflectBoundsGetMethodsToMap ...
func ReflectBoundsGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Bounds{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectBoundsMethodsToMap ...
func ReflectBoundsMethodsToMap() map[int]string {
	var ptr *Bounds
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDiagramGetMethodsToMap ...
func ReflectDiagramGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Diagram{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectDiagramMethodsToMap ...
func ReflectDiagramMethodsToMap() map[int]string {
	var ptr *Diagram
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEdgeGetMethodsToMap ...
func ReflectEdgeGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Edge{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectEdgeMethodsToMap ...
func ReflectEdgeMethodsToMap() map[int]string {
	var ptr *Edge
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLabelGetMethodsToMap ...
func ReflectLabelGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Label{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectLabelMethodsToMap ...
func ReflectLabelMethodsToMap() map[int]string {
	var ptr *Label
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectPlaneGetMethodsToMap ...
func ReflectPlaneGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Plane{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectPlaneMethodsToMap ...
func ReflectPlaneMethodsToMap() map[int]string {
	var ptr *Plane
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectShapeGetMethodsToMap ...
func ReflectShapeGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Shape{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectShapeMethodsToMap ...
func ReflectShapeMethodsToMap() map[int]string {
	var ptr *Shape
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectWaypointGetMethodsToMap ...
func ReflectWaypointGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Waypoint{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectWaypointMethodsToMap ...
func ReflectWaypointMethodsToMap() map[int]string {
	var ptr *Waypoint
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
