package common

import (
	"reflect"
	"sort"
)

// Mapping ...
type Mapping[M BPMNGeneric] struct {
	Anonym,
	Config,
	BPMNType map[int]string
}

// Assign maps the fields of a ReflectValue to the corresponding maps, maintaining the order based on the Pos field.
func (m *Mapping[M]) Assign(v *ReflectValue[M]) {
	// Initialize the maps.
	m.Anonym = make(map[int]string)
	m.Config = make(map[int]string)
	m.BPMNType = make(map[int]string)

	// Define a mapping of field types to their corresponding maps.
	fieldMap := map[reflect.Kind]map[int]string{
		reflect.Bool:   m.Config,
		reflect.Struct: m.BPMNType,
	}

	fields := v.Fields

	switch typedFields := any(fields).(type) {
	case []reflect.StructField:
		// Handle struct fields.
		anonymIndex := 0
		configIndex := 0
		bpmnTypeIndex := 0

		for _, field := range typedFields {
			if field.Anonymous {
				m.Anonym[anonymIndex] = field.Name
				anonymIndex++
			} else {
				// get the map for the field type.
				fieldMapForType, ok := fieldMap[field.Type.Kind()]
				if !ok {
					// if the field type is not recognized, check if it's a BPMN struct.
					if field.Type.Name() == "BPMN" {
						m.BPMNType[bpmnTypeIndex] = field.Name
						bpmnTypeIndex++
					}
				} else {
					// assign the field to the corresponding map.
					fieldMapForType[configIndex] = field.Name
					configIndex++
				}
			}
		}

	case map[string]any:

		// sort by Pos
		type sortedEntry struct {
			Key string
			Pos int
		}

		var sortedEntries []sortedEntry
		configIndex := 0 // index for m.Config

		for key, value := range typedFields {
			valueType := reflect.TypeOf(value)

			if valueType == nil {
				continue // skip, if the value is nil
			}

			// extract the Pos-value, if the value is BPMN
			if valueType.Name() == "BPMN" {
				bpmnValue, ok := any(value).(BPMN)
				if ok {
					sortedEntries = append(sortedEntries, sortedEntry{Key: key, Pos: bpmnValue.Pos})
				}
				continue
			}

			// save all other values in m.Config
			// NOTE: this is temporary, until I have a better solution
			switch valueType.Kind() {
			case reflect.Bool, reflect.String, reflect.Int:
				m.Config[configIndex] = key
				configIndex++
			}

		}

		// sort by Pos
		sort.Slice(sortedEntries, func(i, j int) bool {
			return sortedEntries[i].Pos < sortedEntries[j].Pos
		})

		// save in m.BPMNType after sorting
		for _, entry := range sortedEntries {
			m.BPMNType[entry.Pos] = entry.Key
		}

	}
}

func (m *Mapping[M]) Cleanup() {
	// Clear maps and slices
	m.BPMNType = nil
	m.Config = nil
	m.Anonym = nil
}
