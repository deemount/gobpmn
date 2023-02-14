package collaboration

import "reflect"

// ReflectCollaborationMethodsToMap ...
func ReflectBoundaryEventMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Collaboration{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectParticipantMethodsToMap ...
func ReflectParticipantMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Participant{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}
