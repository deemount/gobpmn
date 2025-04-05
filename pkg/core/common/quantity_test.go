package common

import (
	"reflect"
	"testing"
)

// TestQuantityInitialization verifies that a new Quantity struct is properly initialized
func TestQuantityInitialization(t *testing.T) {
	q := Quantity[StructF]{}

	if q.Pool != 0 {
		t.Errorf("Expected Pool to be 0, got %d", q.Pool)
	}

	if q.Process != 0 {
		t.Errorf("Expected Process to be 0, got %d", q.Process)
	}

	if q.Participant != 0 {
		t.Errorf("Expected Participant to be 0, got %d", q.Participant)
	}

	if q.Elements != nil {
		t.Errorf("Expected Elements to be nil, got %v", q.Elements)
	}
}

// TestCountPool tests the countPool method
func TestCountPool(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected int
	}{
		{"Pool field", "TestPool", 1},
		{"Pool substring", "This_is_a_Pool_field", 1},
		{"No Pool in name", "TestField", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := Quantity[StructF]{}
			q.countPool(test.field)

			if q.Pool != test.expected {
				t.Errorf("Expected Pool count to be %d after counting %s, got %d",
					test.expected, test.field, q.Pool)
			}
		})
	}
}

// Mock struct for testing countFieldsInPool
type MockPool struct {
	TestProcess     string
	AnotherProcess  int
	SomeParticipant bool
	NormalField     float64
}

// TestCountFieldsInPool tests the countFieldsInPool method
func TestCountFieldsInPool(t *testing.T) {
	q := Quantity[StructF]{}
	mockPool := MockPool{}
	v := &ReflectValue[StructF]{}

	v.Pool = reflect.ValueOf(mockPool)
	v.ParticipantName = []string{}
	v.ProcessName = []string{}

	q.countFieldsInPool(v)

	if q.Process != 2 {
		t.Errorf("Expected Process count to be 2, got %d", q.Process)
	}

	if q.Participant != 1 {
		t.Errorf("Expected Participant count to be 1, got %d", q.Participant)
	}

	expectedProcessNames := []string{"Test", "Another"}
	if !reflect.DeepEqual(v.ProcessName, expectedProcessNames) {
		t.Errorf("Expected ProcessName to be %v, got %v", expectedProcessNames, v.ProcessName)
	}

	expectedParticipantNames := []string{"SomeParticipant"}
	if !reflect.DeepEqual(v.ParticipantName, expectedParticipantNames) {
		t.Errorf("Expected ParticipantName to be %v, got %v", expectedParticipantNames, v.ParticipantName)
	}
}

// TestCountFieldsInPoolPanics tests that countFieldsInPool panics when given a non-struct
func TestCountFieldsInPoolPanics(t *testing.T) {
	q := Quantity[StructF]{}
	mockInt := 5
	v := &ReflectValue[StructF]{}

	v.Pool = reflect.ValueOf(mockInt)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected countFieldsInPool to panic with non-struct input")
		}
	}()

	q.countFieldsInPool(v)
}

// Mock structs and functions for testing countFieldsInInstance
type MockInstance struct {
	FromStartEvent string
	UserTask       string
	Gateway        string
	ServiceTask    string
}

// Helper function to create a mocked ReflectValue for single process testing
func createSingleProcessReflectValue() *ReflectValue[StructF] {
	instance := MockInstance{}

	// Create Fields slice similar to what would be in a real ReflectValue
	instanceType := reflect.TypeOf(instance)
	fields := make([]reflect.StructField, instanceType.NumField())
	for i := range instanceType.NumField() {
		fields[i] = instanceType.Field(i)
	}

	v := &ReflectValue[StructF]{}
	v.Process = []reflect.Value{reflect.ValueOf(instance)}
	v.ProcessName = []string{"TestProcess"}
	v.Fields = fields
	v.Instance = reflect.ValueOf(instance)

	return v
}

// Define mock ElementTypeList for tests
type elementMatcher struct {
	name    string
	element processElement
	exact   bool
}

// TestCountFieldsInInstanceSingleProcess tests counting elements in a single process
func TestCountFieldsInInstanceSingleProcess(t *testing.T) {
	q := Quantity[StructF]{}
	v := createSingleProcessReflectValue()

	err := q.countFieldsInInstance(v)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check that elements map is created with the right process index
	if q.Elements[0] == nil {
		t.Fatalf("Expected Elements[0] to be initialized")
	}

	// Check sequence flow count
	if q.Elements[0]["SequenceFlow"] != 1 {
		t.Errorf("Expected 1 SequenceFlow, got %d", q.Elements[0]["SequenceFlow"])
	}

	// Check other element counts
	expected := map[processElement]int{
		"SequenceFlow": 1,
		"UserTask":     1,
		"Gateway":      1,
		"ServiceTask":  1,
	}

	for element, count := range expected {
		if q.Elements[0][element] != count {
			t.Errorf("Expected %d %s, got %d", count, element, q.Elements[0][element])
		}
	}
}

// TestHasParticipant tests the hasParticipant method
func TestHasParticipant(t *testing.T) {
	tests := []struct {
		name        string
		participant int
		process     int
		expected    bool
	}{
		{"Equal process and participant counts", 2, 2, true},
		{"No participants", 0, 2, false},
		{"More processes than participants", 1, 2, false},
		{"More participants than processes", 2, 1, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := Quantity[StructF]{
				Participant: test.participant,
				Process:     test.process,
			}

			if q.hasParticipant() != test.expected {
				t.Errorf("Expected hasParticipant() to return %v with Participant=%d and Process=%d, got %v",
					test.expected, test.participant, test.process, q.hasParticipant())
			}
		})
	}
}

// TestMatchAndCountElement tests the matchAndCountElement method
func TestMatchAndCountElement(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		expected  map[processElement]int
	}{
		{
			"Exact match UserTask",
			"UserTask",
			map[processElement]int{"UserTask": 1},
		},
		{
			"Exact match Gateway",
			"Gateway",
			map[processElement]int{"Gateway": 1},
		},
		{
			"Partial match with Task",
			"CustomTask",
			map[processElement]int{"Task": 1},
		},
		{
			"No match",
			"NoMatch",
			map[processElement]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := Quantity[StructF]{
				Elements: make(map[int]map[processElement]int),
			}
			q.Elements[0] = make(map[processElement]int)

			q.matchAndCountElement(0, test.fieldName)

			for element, expectedCount := range test.expected {
				if q.Elements[0][element] != expectedCount {
					t.Errorf("Expected %d %s, got %d",
						expectedCount, element, q.Elements[0][element])
				}
			}

			// Also check that we don't have unexpected elements
			for element, count := range q.Elements[0] {
				if count > 0 && test.expected[element] != count {
					t.Errorf("Unexpected element count: %s = %d", element, count)
				}
			}
		})
	}
}
