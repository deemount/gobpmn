package common

import (
	"errors"
	"reflect"
	"testing"
)

// TestProcessMethodStruct tests the ProcessMethod struct initialization
func TestProcessMethodStruct(t *testing.T) {
	method := ProcessMethod{
		Name: "TestMethod",
		Arg:  5,
	}

	if method.Name != "TestMethod" {
		t.Errorf("Expected Name to be 'TestMethod', got '%s'", method.Name)
	}

	if method.Arg != 5 {
		t.Errorf("Expected Arg to be 5, got %d", method.Arg)
	}
}

// TestGetProcessMethods tests the GetProcessMethods function
func TestGetProcessMethods(t *testing.T) {
	// Setup test data
	q := &Quantity[[]reflect.StructField]{
		ProcessElements: make(map[int]map[processElement]int),
	}
	processIdx := 0
	q.ProcessElements[processIdx] = make(map[processElement]int)

	// Set some element counts
	q.ProcessElements[processIdx][startEvent] = 1
	q.ProcessElements[processIdx][endEvent] = 2
	q.ProcessElements[processIdx][userTask] = 3
	q.ProcessElements[processIdx][sequenceFlow] = 4

	// Call the function
	methods := GetProcessMethods(processIdx, q)

	// Check results
	expectedMethods := map[string]int{
		"SetStartEvent":             1,
		"SetEndEvent":               2,
		"SetUserTask":               3,
		"SetSequenceFlow":           4,
		"SetIntermediateCatchEvent": 0,
		"SetIntermediateThrowEvent": 0,
		"SetInclusiveGateway":       0,
		"SetExclusiveGateway":       0,
		"SetParallelGateway":        0,
		"SetGateway":                0,
		"SetServiceTask":            0,
		"SetScriptTask":             0,
		"SetTask":                   0,
	}

	if len(methods) != len(expectedMethods) {
		t.Errorf("Expected %d methods, got %d", len(expectedMethods), len(methods))
	}

	for _, method := range methods {
		expectedArg, exists := expectedMethods[method.Name]
		if !exists {
			t.Errorf("Unexpected method: %s", method.Name)
			continue
		}

		if method.Arg != expectedArg {
			t.Errorf("Expected Arg for %s to be %d, got %d",
				method.Name, expectedArg, method.Arg)
		}
	}
}

// MockStruct with methods for testing callMethod and callMethodValue
type MockStruct struct{}

func (m *MockStruct) TestMethod(arg int) error {
	if arg < 0 {
		return errors.New("negative argument")
	}
	return nil
}

func (m *MockStruct) TestMethodNoError(arg int) {
	// No return value
}

func (m *MockStruct) TestMethodReturnValue(arg int) int {
	return arg * 2
}

func (m *MockStruct) TestMethodReturnValueAndError(arg int) (int, error) {
	if arg < 0 {
		return 0, errors.New("negative argument")
	}
	return arg * 2, nil
}

// TestCallMethod tests the callMethod function
func TestCallMethod(t *testing.T) {
	mock := &MockStruct{}
	target := reflect.ValueOf(mock)

	tests := []struct {
		name        string
		methodName  string
		args        []reflect.Value
		expectError bool
	}{
		{
			name:        "Valid method with no error",
			methodName:  "TestMethod",
			args:        []reflect.Value{reflect.ValueOf(5)},
			expectError: false,
		},
		{
			name:        "Valid method with error",
			methodName:  "TestMethod",
			args:        []reflect.Value{reflect.ValueOf(-1)},
			expectError: true,
		},
		{
			name:        "Method with no return value",
			methodName:  "TestMethodNoError",
			args:        []reflect.Value{reflect.ValueOf(5)},
			expectError: false,
		},
		{
			name:        "Non-existent method",
			methodName:  "NonExistentMethod",
			args:        []reflect.Value{},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := callMethod(target, test.methodName, test.args)

			if test.expectError && err == nil {
				t.Errorf("Expected an error, got nil")
			}

			if !test.expectError && err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
		})
	}
}

// TestCallMethodValue tests the callMethodValue function
func TestCallMethodValue(t *testing.T) {
	mock := &MockStruct{}
	target := reflect.ValueOf(mock)

	tests := []struct {
		name          string
		methodName    string
		args          []reflect.Value
		expectedValue any
		expectError   bool
	}{
		{
			name:          "Method returning a value",
			methodName:    "TestMethodReturnValue",
			args:          []reflect.Value{reflect.ValueOf(5)},
			expectedValue: 10,
			expectError:   false,
		},
		{
			name:          "Method returning a value and no error",
			methodName:    "TestMethodReturnValueAndError",
			args:          []reflect.Value{reflect.ValueOf(5)},
			expectedValue: 10,
			expectError:   false,
		},
		{
			name:          "Method returning a value and error",
			methodName:    "TestMethodReturnValueAndError",
			args:          []reflect.Value{reflect.ValueOf(-1)},
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:        "Method with no return value",
			methodName:  "TestMethodNoError",
			args:        []reflect.Value{reflect.ValueOf(5)},
			expectError: true,
		},
		{
			name:        "Non-existent method",
			methodName:  "NonExistentMethod",
			args:        []reflect.Value{},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := callMethodValue(target, test.methodName, test.args)

			if test.expectError && err == nil {
				t.Errorf("Expected an error, got nil")
			}

			if !test.expectError && err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}

			if !test.expectError && result.Interface() != test.expectedValue {
				t.Errorf("Expected return value %v, got %v",
					test.expectedValue, result.Interface())
			}
		})
	}
}

// TestGetMethodNames tests the getMethodNames function
func TestGetMethodNames(t *testing.T) {
	mock := &MockStruct{}
	target := reflect.ValueOf(mock)

	methods := getMethodNames(target)

	expectedMethods := []string{
		"TestMethod",
		"TestMethodNoError",
		"TestMethodReturnValue",
		"TestMethodReturnValueAndError",
	}

	// Check that all expected methods are present
	for _, expected := range expectedMethods {
		found := false
		for _, actual := range methods {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected method %s not found in results", expected)
		}
	}
}

// TestTargetMethodName tests the targetMethodName function
func TestTargetMethodName(t *testing.T) {
	mock := &MockStruct{}
	target := reflect.ValueOf(mock)

	tests := []struct {
		name        string
		methodName  string
		expectError bool
	}{
		{
			name:        "Valid method",
			methodName:  "TestMethod",
			expectError: false,
		},
		{
			name:        "Non-existent method",
			methodName:  "NonExistentMethod",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			method, err := targetMethodName(target, test.methodName)

			if test.expectError && err == nil {
				t.Errorf("Expected an error, got nil")
			}

			if !test.expectError && err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}

			if !test.expectError && !method.IsValid() {
				t.Errorf("Expected valid method, got invalid")
			}
		})
	}
}
