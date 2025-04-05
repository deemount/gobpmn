package common

// go:build 1.23

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"reflect"
	"sort"
	"strings"
	"sync"
)

// Quantity holds all the quantities of the BPMN elements
// in the BPMN model. It is used to count the number of elements
type Quantity[M BPMNGeneric] struct {
	sync.RWMutex
	Elements    map[int]map[processElement]int
	Pool        int
	Process     int
	Participant int
}

// countPool counts a pool in the process model.
func (q *Quantity[M]) countPool(field string) {
	if strings.Contains(field, "Pool") {
		q.Pool++
	}
}

// countFieldsInPool counts the number of participants in the BPMN model, which are defined in a Pool.
// A pool is structured and have configurations, a collaboration, processes, ID's and messages.
// v.ProcessName string slice contains the names of the processes in the BPMN model.
// v.ParticipantName string slice contains the names of the participants in the BPMN model.
// Ruleset:
//   - All fields of the pool interface MUST be type of struct
//   - If the reflection field contains the word "Process", then count a process.
//   - If the reflection field contains the word "Participant", then count a participant.
func (q *Quantity[M]) countFieldsInPool(v *ReflectValue[M]) {
	if v.Pool.Kind() != reflect.Struct {
		panic("Input data must be a struct") // Note: panic or error handling?
	}
	for i := range v.Pool.NumField() {
		field := v.Pool.Type().Field(i)
		if strings.Contains(field.Name, "Process") {
			v.ProcessName = append(v.ProcessName, extractPrefixBeforeProcess(field.Name))
			q.Process++
		}
		if strings.Contains(field.Name, "Participant") {
			v.ParticipantName = append(v.ParticipantName, field.Name)
			q.Participant++
		}
	}
}

// countFieldsInInstance counts the fields in a process and stores them in a map.
// NOTE: Works correctly for standalone and multiple processes designed by struct,
// as well as a standalone mapped processes
func (q *Quantity[M]) countFieldsInInstance(v *ReflectValue[M]) error {
	if q.Elements == nil {
		q.Elements = make(map[int]map[processElement]int)
	}
	for processIdx, processName := range v.ProcessName {
		if q.Elements[processIdx] == nil {
			q.Elements[processIdx] = make(map[processElement]int)
		}
		if q.Process > 1 {
			if err := q.countMultipleProcessElements(v, processIdx, processName); err != nil {
				log.Printf("Error processing multiple processes: %v", err)
				continue // skip the process and continue with the next one; silent error handling
			}
		} else {
			q.countStandaloneProcessElements(v, processIdx)
		}
	}
	return nil
}

// countMultipleProcessElements handles counting for elements in multiple processes.
// - It counts the elements in v.Instance.
func (q *Quantity[M]) countMultipleProcessElements(v *ReflectValue[M], processIdx int, processName string) error {
	field := instanceFieldName(v, processName)
	if !field.IsValid() {
		return NewError(fmt.Errorf("invalid field for process %s", processName))
	}
	return q.countFieldElements(field, processIdx)
}

// countFieldElements counts elements in a reflected struct field.
func (q *Quantity[M]) countFieldElements(field reflect.Value, processIdx int) error {
	for i := range field.NumField() {
		fieldName := field.Type().Field(i).Name
		// handle sequence flows first
		if strings.HasPrefix(fieldName, "From") {
			q.Elements[processIdx]["SequenceFlow"]++
			continue
		}
		q.matchAndCountElement(processIdx, fieldName)
	}
	return nil
}

// countStandaloneProcessElements handles counting for a single process
func (q *Quantity[M]) countStandaloneProcessElements(v *ReflectValue[M], processIdx int) {
	switch fields := any(v.Fields).(type) {
	case []reflect.StructField:
		// run directly through the struct fields
		for _, field := range fields {
			fieldName := field.Name
			// count sequence flows first
			if strings.HasPrefix(fieldName, "From") {
				q.Elements[processIdx]["SequenceFlow"]++
				continue
			}
			q.matchAndCountElement(processIdx, fieldName)
		}

	case map[string]any:
		// collect map fields and sort them by Pos
		type sortedEntry struct {
			Key string
			Pos int
		}

		var sortedEntries []sortedEntry

		for key, value := range fields {
			if bpmnValue, ok := value.(BPMN); ok {
				// save BPMN elements by Pos
				sortedEntries = append(sortedEntries, sortedEntry{Key: key, Pos: bpmnValue.Pos})
			} else {
				// save directly, if it's not a BPMN element
				q.matchAndCountElement(processIdx, key)
			}
		}

		// sort by pos
		sort.Slice(sortedEntries, func(i, j int) bool {
			return sortedEntries[i].Pos < sortedEntries[j].Pos
		})

		// count elements after sorting
		for _, entry := range sortedEntries {
			// count sequence flows first
			if strings.HasPrefix(entry.Key, "From") {
				q.Elements[processIdx]["SequenceFlow"]++
				continue
			}
			q.matchAndCountElement(processIdx, entry.Key)
		}
	}
}

// matchAndCountElement matches and counts a single element using the two-pass method.
// NOTE: seperated exact and partial matching
func (q *Quantity[M]) matchAndCountElement(processIdx int, fieldName string) {
	for _, matcher := range ElementTypeList {
		if matcher.exact && fieldName == matcher.name {
			q.Elements[processIdx][matcher.element]++
			return // found an exact match, no need to continue; early return
		}
	}
	for _, matcher := range ElementTypeList {
		if !matcher.exact && strings.Contains(fieldName, matcher.name) {
			q.Elements[processIdx][matcher.element]++
			return // found a partial match, no need to continue; early return
		}
	}
}

// hasParticipant checks if a process has a participant.
func (q *Quantity[M]) hasParticipant() bool {
	return q.Participant > 0 && (q.Process == q.Participant)
}
