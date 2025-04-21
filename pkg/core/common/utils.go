package common

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

// typ ...
func typ(n string) string {
	// define a mapping of BPMN element types to their corresponding names.
	bpmnTypes := map[string]string{
		"IsExecutable":           "config",
		"Process":                "process",
		"StartEvent":             "startevent",
		"EndEvent":               "event",
		"IntermediateCatchEvent": "event",
		"IntermediateThrowEvent": "event",
		"InclusiveGateway":       "gateway",
		"ExclusiveGateway":       "gateway",
		"ParallelGateway":        "gateway",
		"Gateway":                "gateway",
		"ScriptTask":             "activity",
		"UserTask":               "activity",
		"ServiceTask":            "activity",
		"Task":                   "activity",
	}

	// check if the name contains "From" to determine if it's a flow.
	if strings.Contains(n, "From") {
		return "flow"
	}
	// iterate over the BPMN types and return the type if the name contains the type.
	for t, bpmnType := range bpmnTypes {
		if strings.Contains(n, t) {
			return bpmnType
		}
	}
	// Return an empty string if no matching type is found.
	return ""
}

// inject returns a hash value of a given string.
// It uses the FNV-1a algorithm to generate a hash value.
// The method returns a struct, called BPMN, with the hash value.
// The argument typ is the type of the BPMN element.
func inject(typ string, idx int) (BPMN, error) {
	n := 8
	b := make([]byte, n)
	c := fnv.New32a()
	if _, err := rand.Read(b); err != nil {
		return BPMN{}, err
	}
	s := fmt.Sprintf("%x", b)
	if _, err := c.Write([]byte(s)); err != nil {
		return BPMN{}, err
	}
	defer c.Reset()

	var h string

	/*
	 * Note: the following code block is not clear. The first if statement is not clear.
	 * The if statement checks if the index is 0 and the type is "startevent".
	 * If the condition is true, the hash value is set to "1".
	 * The else block sets the hash value to the string representation of the checksum.
	 *
	 * This specific handling is a mock up from the Camunda Modeler 5.2.0.
	 */
	if idx == 0 && typ == "startevent" {
		h = fmt.Sprintf("%x", 1)
	} else {
		h = fmt.Sprintf("%x", string(c.Sum(nil)))
	}

	result := BPMN{
		Type: typ, // Note: needs to be reconsidered. Is this the right method, to put the typ value into the Type field?
		Hash: h,
	}
	return result, nil
}

// getNextHash ...
// Note: maybe redudant? look at ReflectValue method "nextHash". Refactor?
func getNextHash(field reflect.Value, fieldIdx, numFields int) string {
	if fieldIdx+1 < numFields {
		return field.Field(fieldIdx + 1).FieldByName("Hash").String()
	}
	return ""
}

// extractPrefixBeforeProcess extracts the prefix before the word "Process" in a string.
// The method returns the prefix as a string, which is used as the name of the process.
// Note: the mthod is used two times: first in newReflectValue (reflect_di.go) and second countFieldsInPool (quantities.go)
func extractPrefixBeforeProcess(input string) string {
	re := regexp.MustCompile(`([A-Za-z]+)Process$`)
	match := re.FindStringSubmatch(input)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// extractLastTwoWords extracts the last two words from a string
func extractLastTwoWords(input string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	words := re.FindAllString(input, -1)
	if len(words) < 2 {
		return strings.Join(words, "")
	}
	lastTwo := words[len(words)-2:]
	return strings.Join(lastTwo, "")
}

// sortedKeys returns the sorted keys of a map[string]any.
// It sorts the keys by the BPMN element position (Pos).
// The BPMN elements are stored in a temporary struct.
// Note: the method is not used actually. There is an another solution in the
// mapping.go file, which is used to sort the BPMN elements by their position.
func sortedKeys(fields map[string]any, valueOf reflect.Value) []string {
	var sortedKeys []string
	// temporary struct to store BPMN elements
	bpmnElements := make([]struct {
		Key string
		Pos int
	}, 0)
	// iterate over the map keys
	for _, key := range valueOf.MapKeys() {
		strKey := key.String()
		value := valueOf.MapIndex(key).Interface()
		switch v := value.(type) {
		case BPMN:
			bpmnElements = append(bpmnElements, struct {
				Key string
				Pos int
			}{strKey, v.Pos})
		}
	}
	// sort BPMN elements by Pos
	sort.SliceStable(bpmnElements, func(i, j int) bool {
		return bpmnElements[i].Pos < bpmnElements[j].Pos
	})
	// create sorted keys
	for _, elem := range bpmnElements {
		sortedKeys = append(sortedKeys, elem.Key)
	}
	return sortedKeys
}
