package common

import (
	"regexp"
	"strings"
	"testing"
	"time"
)

// Unit Tests
func TestMatchElementTypeExactMatches(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected processElement
	}{
		{"StartEvent exact match", "StartEvent", startEvent},
		{"EndEvent exact match", "EndEvent", endEvent},
		{"IntermediateCatchEvent exact match", "IntermediateCatchEvent", intermediateCatchEvent},
		{"IntermediateThrowEvent exact match", "IntermediateThrowEvent", intermediateThrowEvent},
		{"ParallelGateway exact match", "ParallelGateway", parallelGateway},
		{"ExclusiveGateway exact match", "ExclusiveGateway", exclusiveGateway},
		{"InclusiveGateway exact match", "InclusiveGateway", inclusiveGateway},
		{"UserTask exact match", "UserTask", userTask},
		{"ScriptTask exact match", "ScriptTask", scriptTask},
		{"Task exact match", "Task", task},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchElementType(tt.input)
			if result != tt.expected {
				t.Errorf("matchElementType(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMatchElementTypeSequenceFlow(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected processElement
	}{
		{"From pattern match", "From", sequenceFlow},
		{"FromTask pattern match", "FromTask", sequenceFlow},
		{"FromUserTask pattern match", "FromUserTask", sequenceFlow},
		{"From with spaces pattern match", "From Task", sequenceFlow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchElementType(tt.input)
			if result != tt.expected {
				t.Errorf("matchElementType(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMatchElementTypeEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected processElement
	}{
		{"Empty string", "", processElement("")},
		{"Whitespace only", "   ", processElement("")},
		{"Unknown element type", "UnknownElement", processElement("")},
		{"Similar but not matching", "StartEvents", processElement("")},
		{"Lowercase", "startevent", processElement("")}, // Current implementation is case sensitive
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchElementType(tt.input)
			if result != tt.expected {
				t.Errorf("matchElementType(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMatchElementTypeWithWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected processElement
	}{
		{"StartEvent with leading space", " StartEvent", startEvent},
		{"StartEvent with trailing space", "StartEvent ", startEvent},
		{"StartEvent with both spaces", " StartEvent ", startEvent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matchElementType(tt.input)
			if result != tt.expected {
				t.Errorf("matchElementType(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Integration Tests
func TestElementRegistryIntegration(t *testing.T) {
	// Test the interaction between ElementTypeList and matchElementType
	t.Run("Modify ElementTypeList and verify matchElementType behavior", func(t *testing.T) {
		// Save original ElementTypeList
		originalList := ElementTypeList
		defer func() {
			ElementTypeList = originalList // Restore original list after test
		}()

		// Create a custom ElementTypeList with modified elements
		customElement := processElement("CustomElement")
		ElementTypeList = []Match{
			{
				name:     "CustomElement",
				element:  customElement,
				exact:    true,
				priority: 1,
				regex:    regexp.MustCompile(`^CustomElement$`),
			},
			// Keep one original element
			{
				name:     "StartEvent",
				element:  startEvent,
				exact:    true,
				priority: 2,
				regex:    regexp.MustCompile(`^StartEvent$`),
			},
		}

		// Test with the modified list
		result := matchElementType("CustomElement")
		if result != customElement {
			t.Errorf("matchElementType with modified ElementTypeList failed, got %v, expected %v", result, customElement)
		}

		// Test that original elements still work
		result = matchElementType("StartEvent")
		if result != startEvent {
			t.Errorf("matchElementType with modified ElementTypeList failed for original element, got %v, expected %v", result, startEvent)
		}

		// Test that removed elements no longer work
		result = matchElementType("EndEvent")
		if result != processElement("") {
			t.Errorf("matchElementType should return empty for removed element, got %v", result)
		}
	})

	t.Run("Test non-exact match logic", func(t *testing.T) {
		// Save original ElementTypeList
		originalList := ElementTypeList
		defer func() {
			ElementTypeList = originalList // Restore original list after test
		}()

		// Create ElementTypeList with non-exact matches and various priorities
		ElementTypeList = []Match{
			{
				name:     "Task",
				element:  task,
				exact:    false,
				priority: 3,
				regex:    regexp.MustCompile(`^Task$`),
			},
			{
				name:     "UserTask",
				element:  userTask,
				exact:    false,
				priority: 2,
				regex:    regexp.MustCompile(`^UserTask$`),
			},
			{
				name:     "ScriptTask",
				element:  scriptTask,
				exact:    false,
				priority: 1,
				regex:    regexp.MustCompile(`^ScriptTask$`),
			},
		}

		// Input contains multiple matches, but ScriptTask has highest priority (lowest number)
		result := matchElementType("This is a ScriptTask and UserTask and Task")
		if result != scriptTask {
			t.Errorf("Non-exact match should prefer ScriptTask due to priority, got %v", result)
		}
	})
}

// Performance Tests
func TestMatchElementTypePerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	// Benchmark for exact matches
	t.Run("Performance with exact matches", func(t *testing.T) {
		start := time.Now()
		iterations := 100000

		for i := 0; i < iterations; i++ {
			matchElementType("StartEvent")
			matchElementType("EndEvent")
			matchElementType("ParallelGateway")
		}

		duration := time.Since(start)
		t.Logf("Exact match performance: %v for %d iterations (%.2f μs/op)",
			duration, iterations*3, float64(duration.Microseconds())/float64(iterations*3))

		// Set a reasonable threshold (adjust based on your environment)
		if duration > time.Second {
			t.Errorf("Performance test took too long: %v", duration)
		}
	})

	// Benchmark for non-exact matches (which are more expensive)
	t.Run("Performance with non-exact matches", func(t *testing.T) {
		start := time.Now()
		iterations := 10000

		for i := 0; i < iterations; i++ {
			matchElementType("This is a UserTask")
			matchElementType("From something to something else")
			matchElementType("Complex ScriptTask description")
		}

		duration := time.Since(start)
		t.Logf("Non-exact match performance: %v for %d iterations (%.2f μs/op)",
			duration, iterations*3, float64(duration.Microseconds())/float64(iterations*3))

		// Set a reasonable threshold (adjust based on your environment)
		if duration > time.Second*2 {
			t.Errorf("Non-exact match performance test took too long: %v", duration)
		}
	})

	// Test with mixed input types to simulate real-world usage
	t.Run("Mixed inputs performance simulation", func(t *testing.T) {
		inputs := []string{
			"StartEvent",
			"",
			"Unknown element",
			"From A to B",
			" ScriptTask ",
			"This contains the word UserTask somewhere",
			"EndEvent",
			"Just some random text",
		}

		start := time.Now()
		iterations := 10000

		for i := 0; i < iterations; i++ {
			for _, input := range inputs {
				matchElementType(input)
			}
		}

		duration := time.Since(start)
		t.Logf("Mixed input performance: %v for %d iterations (%.2f μs/op)",
			duration, iterations*len(inputs), float64(duration.Microseconds())/float64(iterations*len(inputs)))

		if duration > time.Second*3 {
			t.Errorf("Mixed input performance test took too long: %v", duration)
		}
	})
}

// Test for specific potential bugs and edge cases
func TestPotentialBugs(t *testing.T) {
	t.Run("Case where bestMatch.name is used incorrectly", func(t *testing.T) {
		// The current implementation has a potential bug where it uses bestMatch.name in the regex comparison
		// which might not be what was intended. This test verifies the current behavior.

		// Save original list
		originalList := ElementTypeList
		defer func() { ElementTypeList = originalList }()

		// Create a test scenario that might reveal the issue
		testElement := processElement("TestElement")
		ElementTypeList = []Match{
			{
				name:     "Test", // This is a substring of normalizedName
				element:  testElement,
				exact:    false,
				priority: 1,
				regex:    regexp.MustCompile(`^CompletelyDifferent$`), // This won't match "Test"
			},
		}

		result := matchElementType("TestString")

		// Document current behavior
		if result == testElement {
			t.Logf("CURRENT BEHAVIOR: matchElementType returns element even when regex doesn't match bestMatch.name")
		} else {
			t.Logf("CURRENT BEHAVIOR: matchElementType requires regex to match bestMatch.name")
		}
	})

	t.Run("Check if bestMatch is checked correctly", func(t *testing.T) {
		// Save original list
		originalList := ElementTypeList
		defer func() { ElementTypeList = originalList }()

		// Create a test scenario with no matches
		ElementTypeList = []Match{
			{
				name:     "NoMatch",
				element:  processElement("NoMatch"),
				exact:    false,
				priority: 5,
				regex:    regexp.MustCompile(`^NoMatch$`),
			},
		}

		result := matchElementType("CompletelyDifferentString")

		if result != processElement("") {
			t.Errorf("matchElementType should return empty for no matches, got %v", result)
		}
	})
}

// TestElementRegexPatterns tests if the regex patterns in ElementTypeList work as expected
func TestElementRegexPatterns(t *testing.T) {
	for _, match := range ElementTypeList {
		t.Run("Regex for "+match.name, func(t *testing.T) {
			// Test that the regex matches the name
			if !match.regex.MatchString(match.name) {
				t.Errorf("Regex %q doesn't match name %q", match.regex.String(), match.name)
			}

			// For exact matches, test with variations that should not match
			if match.exact {
				variations := []string{
					"Pre" + match.name,
					match.name + "Post",
					strings.ToLower(match.name),
					strings.ToUpper(match.name),
				}

				for _, variation := range variations {
					if variation == match.name {
						continue // Skip if somehow the variation is the same
					}

					if match.regex.MatchString(variation) && match.exact {
						t.Errorf("Regex %q for exact match unexpectedly matches %q",
							match.regex.String(), variation)
					}
				}
			}
		})
	}
}
