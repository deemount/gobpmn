package factory

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"log"
	"reflect"
	"strings"

	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/pool"
	"github.com/deemount/gobpmn/utils"
)

type def core.DefinitionsRepository

var (
	configID           = "ID"
	configProcess      = "Process"
	configIsExecutable = "IsExecutable"
	configMessage      = "Message"
)

// BpmnBuilder ...
type BpmnBuilder struct {
	Analyze
	Settings
	Suffix string
}

// BpmnBuilder ...
func (h *BpmnBuilder) Hash() string {
	if h.isZero() {
		r := h.hash()
		h.Suffix = r.Suffix
	}
	return h.Suffix
}

// Inject all anonymous fields with a hash value by fields with type BpmnBuilder
// It also setup reflected fields with boolean type and checks out the configuration by wording
// usage:
// e.g.
// var p CProcess
// p = h.inject(CProcess{}).(CProcess)
func (h *BpmnBuilder) Inject(p interface{}) interface{} { return h.inject(p) }

// Create receives the definitions repository by the app in p argument
// and calls the main elements to set the maps, including process parameters
// n of process.
func (h *BpmnBuilder) Build(p interface{}) { h.build(p) }

/*
 * @private
 */

// hash ...
func (h BpmnBuilder) hash() BpmnBuilder {
	n := 8
	b := make([]byte, n)
	c := fnv.New32a()

	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)

	c.Write([]byte(s))
	defer c.Reset()

	r := BpmnBuilder{Suffix: fmt.Sprintf("%x", string(c.Sum(nil)))}

	return r
}

// inject ...
func (h *BpmnBuilder) inject(p interface{}) interface{} {

	log.Println("builder: start injecting fields")

	// el is the interface {}
	el := reflect.ValueOf(&p).Elem()

	// Allocate a temporary variable with type of the struct.
	// el.Elem() is the value contained in the interface
	tmp := reflect.New(el.Elem().Type()).Elem()
	tmp.Set(el.Elem())

	// initialize maps to analyze
	// anonym: all anonymous fields
	// config: all boolean fields
	// builder: all builder fields
	// words: all collected words splitted
	h.anonym = make(map[int]string)
	h.config = make(map[int]string)
	h.builder = make(map[int]string)
	h.words = make(map[int][]string)

	// get anonymous field from interface {}
	h.reflectAnonymousFields(p)

	// after reflected anonymous fields
	h.reflectBuilderType(p)
	h.reflectBoolType(p)

	// anonymous field are reflected
	if h.hasAnonymous() {

		// walk through the map with names of anonymous fields
		for _, field := range h.anonym {

			// get the reflected value of field
			n := tmp.FieldByName(field)

			// walk through the values of fields by the interface {}
			for i := 0; i < n.NumField(); i++ {

				// set by kind of reflected value above
				switch n.Field(i).Kind() {

				// kind is a struct
				case reflect.Struct:

					val := n.Type().Field(i).Name

					if h.isPool(field) {
						if strings.Contains(val, configProcess) {
							h.NumProc++
						}
						if strings.Contains(val, configID) {
							h.NumPart++
						}
					}

					if h.isMessage(field) {
						if strings.Contains(val, configMessage) {
							h.NumMsg++
						}
					}

					// generate hash value
					hash := h.hash()
					n.Field(i).Set(reflect.ValueOf(hash))

					log.Printf("builder: inject struct field %s", val)

					break

				// kind is a bool
				case reflect.Bool:

					val := n.Type().Field(i).Name

					// only the first field, which IsExecutable is set to true,
					// means, only one process in a collaboration can be executed at runtime
					// this can be changed in the future, if the engine fits for more execution
					// options
					if strings.Contains(val, configIsExecutable) && i == 0 {
						n.Field(0).SetBool(true)
						log.Printf("builder: inject bool field %s", val)
					}

					break
				}
			}
		}

		el.Set(tmp)
		p = el.Interface()

	}

	// anonymous field aren't reflected
	if h.hasNotAnonymous() {

		// walk through the map with names of anonymous fields
		for _, field := range h.builder {

			// get the reflected value of field
			n := tmp.FieldByName(field)

			// generate hash value
			hash := h.hash()
			n.Set(reflect.ValueOf(hash))

		}

		// walk through the map with names of anonymous fields
		for _, field2 := range h.config {

			// get the reflected value of field
			n2 := tmp.FieldByName(field2)

			//
			n2.SetBool(true)

		}

		el.Set(tmp)
		p = el.Interface()

	}

	h.countWords()

	return p

}

// build ...
func (h *BpmnBuilder) build(p interface{}) {
	// el is the interface {}
	el := reflect.ValueOf(&p).Elem()

	// Allocate a temporary variable with type of the struct.
	// el.Elem() is the value contained in the interface
	tmp := reflect.New(el.Elem().Type()).Elem()
	tmp.Set(el.Elem())

	// set main elements of the core package
	// h.numProc represents number of processes
	core.SetMainElements(tmp.Interface().(def), h.NumProc)
}

// reflectAppAnonymousFields takes all compounds which are anonymous
// The implementation is O(N) for insertions
func (h *BpmnBuilder) reflectAnonymousFields(p interface{}) {

	fields := reflect.VisibleFields(reflect.TypeOf(p))
	i := 0
	for _, field := range fields {
		if h.isAnonymous(field) {
			h.anonym[i] = field.Name
			h.words[i] = utils.Split(h.anonym[i])
			i++
		}
	}
	if i == 0 {
		log.Println("builder: main struct contains map of zero anonymous fields")
	}

}

// reflectBuilderType with three static filter options
func (h BpmnBuilder) reflectBuilderType(p interface{}) {
	fields := reflect.VisibleFields(reflect.TypeOf(p))
	count := 0
	index := len(h.words)
	for _, field := range fields {
		if h.isNotAnonymous(field) && h.isNotDefinitions(field) && h.isBpmnBuilder(field) {
			h.builder[count] = field.Name
			h.words[index] = utils.Split(h.builder[count])
			count++
			index++
		}
	}
}

// reflectBoolType with one static filter option, which must be kind of reflect.Bool
// The bool type of a field in a struct describes mostly configuration settings
// The field must contain a sibling in title case, e.g. IsExecutable
func (h BpmnBuilder) reflectBoolType(p interface{}) {
	fields := reflect.VisibleFields(reflect.TypeOf(p))
	count := 0
	index := len(h.words)
	for _, field := range fields {
		if field.Type.Kind() == reflect.Bool {
			h.config[count] = field.Name
			h.words[index] = utils.Split(h.config[count])
			count++
			index++
		}
	}
}

// countWords ...
func (h BpmnBuilder) countWords() {
	l := 0
	length := len(h.words)
	for i := 0; i < length; i++ {
		l += len(h.words[i])
	}
	log.Printf("builder: count app words %v", l)
}

// isZero ...
func (h *BpmnBuilder) isZero() bool { return h.Suffix == "" }

// isNotDefinitions ...
func (h *BpmnBuilder) isDefinitions(field reflect.StructField) bool { return core.IsDefinitions(field) }

// isNotDefinitions ...
func (h *BpmnBuilder) isNotDefinitions(field reflect.StructField) bool {
	return core.IsNotDefinitions(field)
}

// isPool ...
func (h *BpmnBuilder) isPool(field string) bool { return pool.IsPool(field) }

// isMessage ...
func (h *BpmnBuilder) isMessage(field string) bool { return elements.IsMessage(field) }

// isUnknown ...
func (h *BpmnBuilder) isUnknown(field string) bool { return true }

// isAnonymous ...
func (h *BpmnBuilder) isAnonymous(field reflect.StructField) bool { return field.Anonymous }

// isNotAnonymous ...
func (h *BpmnBuilder) isNotAnonymous(field reflect.StructField) bool { return !field.Anonymous }

// hasAnonymous ...
func (h *BpmnBuilder) hasAnonymous() bool { return len(h.anonym) > 0 }

// hasNotAnonymous ...
func (h *BpmnBuilder) hasNotAnonymous() bool { return len(h.anonym) == 0 }

// isBpmnBuilder ...
func (h *BpmnBuilder) isBpmnBuilder(field reflect.StructField) bool {
	return field.Type.Name() == "BpmnBuilder"
}
