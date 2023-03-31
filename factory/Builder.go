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

type Def core.DefinitionsRepository

// Builder ...
type Builder struct {
	Reflect
	Settings
	Suffix    string
	HashTable []string
}

// Hash ...
func (h *Builder) Hash() string {
	if h.isZero() {
		r := h.hash()
		h.Suffix = r.Suffix
	}
	return h.Suffix
}

// Inject all anonymous fields with a hash value by fields with type Builder
// It also setup reflected fields with boolean type and checks out the configuration by wording
// usage:
// e.g.
// var p CProcess
// p = h.inject(CProcess{}).(CProcess)
func (h *Builder) Inject(p interface{}) interface{} { return h.inject(p) }

// Create receives the definitions repository by the app in p argument
// and calls the main elements to set the maps, including process parameters
// n of process.
func (h *Builder) Build(p interface{}) { h.build(p) }

/*
 * @private
 */

// hash ...
func (h Builder) hash() Builder {

	n := 8
	b := make([]byte, n)
	c := fnv.New32a()

	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)

	if _, err := c.Write([]byte(s)); err != nil {
		panic(err)
	}
	defer c.Reset()

	r := Builder{Suffix: fmt.Sprintf("%x", string(c.Sum(nil)))}

	return r
}

// hashTable ...
// Note:
// @self is a pseudo and means "this hash is set to this element"
// Ruleset for a straight process without any decision:
// A Collaboration has one hash value: self
// A Participant has two hash values: id, process
// A Pool has one hash value: id
// A StartEvent has two hash values: self, next (flow: from)
// A Flow has three hash values: self, previous, next (task ...)
// A Task, Event has three hash values: self, previous (flow: from), next (flow. from)
// A EndEvent has two hash values: self, previous (flow: from)
func (h Builder) hashTable(index int, val reflect.Value) Builder {

	n := 8
	b := make([]byte, n)
	c := fnv.New32a()

	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)

	if _, err := c.Write([]byte(s)); err != nil {
		panic(err)
	}
	defer c.Reset()

	r := Builder{
		HashTable: []string{
			fmt.Sprintf("%x", string(c.Sum(nil))),
		},
	}

	strHash := fmt.Sprintf("%s", val.Field(index).FieldByName("Suffix"))
	if strHash == "" {
		val.Field(index).Set(reflect.ValueOf(r))
	}

	if index+1 < val.NumField() {
		nexthash := Builder{HashTable: []string{}}          // generate hash value
		val.Field(index + 1).Set(reflect.ValueOf(nexthash)) // inject the field
	}

	return r
}

// inject itself reflects a given struct and inject
// signed fields with hash values.
// There are two conditions to assign fields of a strcut:
// a) The struct has anonymous fields
// b) The struct has no anymous fields
// It also counts the element in their specification to know
// how much elements of each package needs to be mapped later then.
func (h *Builder) inject(p interface{}) interface{} {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("factory.builder: inject recovered", r)
		}
	}()

	log.Println("factory.builder: start injecting fields")

	ref := NewReflect(p)
	ref.Interface().New().Maps().Reflection()

	// TODO: ref.Analyze fields can be used here
	//log.Printf("ref.Analyze.Anonym %+v", ref.Analyze.Anonym)
	//log.Printf("ref.Analyze.Builder %+v", ref.Analyze.Builder)
	//log.Printf("ref.Analyze.Config %+v", ref.Analyze.Config)
	//log.Printf("ref.Analyze.Words %+v", ref.Analyze.Words)

	l := 0

	// TODO: Describe how the injection rules for anonymous fields are
	// anonymous field are reflected
	if ref.hasAnonymous() {

		// walk through the map with names of anonymous fields
		for index, field := range ref.Anonym {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(field)

			// walk through the values of fields assigned to the interface {}
			for i := 0; i < n.NumField(); i++ {

				name := n.Type().Field(i).Name
				if strings.Contains(name, field) {
					// TODO: Thanks for the info, but it seems that it's useless
					//log.Printf("factory.builder: detected field %s contains anonymous field %s", name, field)
				}

				// set by kind of reflected value above
				switch n.Field(i).Kind() {

				// kind is a bool
				case reflect.Bool:

					// only the first field, which IsExecutable is set to true,
					// means, only one process in a collaboration can be executed at runtime
					// this can be changed in the future, if the engine fits for more execution
					// options
					if strings.Contains(name, boolIsExecutable) && i == 0 {
						n.Field(0).SetBool(true)
						log.Printf("factory.builder: inject first bool field %s once", name)
					} else {
						log.Printf("factory.builder: second bool field %s stays false", name)
					}

					log.Printf("factory.builder: >> field injection %d of %d done", i+1, n.NumField())

					break

				// kind is a struct
				case reflect.Struct:

					// TODO: The counting list could become bigger, when different elements needs to be counted
					h.countPool(field, name)    // counts processes, participants and their shapes
					h.countMessage(field, name) // counts message flows and their edges
					h.countStartEvent(name)     // counts startevents
					h.countTask(name)           // counts tasks
					h.countEndEvent(name)       // counts endevent

					strHash := fmt.Sprintf("%s", n.Field(i).FieldByName("Suffix"))
					if strHash == "" {
						hash := h.hash()                      // generate hash value
						n.Field(i).Set(reflect.ValueOf(hash)) // inject the field
						//h.HashTable[0] = hash.Suffix
						log.Printf("factory.builder: inject first at process index [%d] => %v on current fieldname %s (%v) with hash value ", index, ref.Analyze.Words[index], n.Type().Field(i).Name, n.Field(i).FieldByName("Suffix"))
					} else {
						log.Printf("factory.builder: current fieldname %s (%v) has got hash value before", n.Type().Field(i).Name, n.Field(i).FieldByName("Suffix"))
					}

					// TODO: Check only previous element; maybe a solution trying to erate the hash slice
					if i > 0 {
						if n.Field(i-1).Kind() != reflect.Bool {
							log.Printf("factory.builder: previous fieldname was %s (%v)", n.Type().Field(i-1).Name, n.Field(i-1).FieldByName("Suffix"))
						}
					} else {
						log.Printf("factory.builder: current fieldname %s has no previous field", n.Type().Field(i).Name)
					}

					// TODO: Check next element and set hash value; maybe a solution trying to erate the hash slice
					if i+1 < n.NumField() {
						nexthash := h.hash()                          // generate hash value
						n.Field(i + 1).Set(reflect.ValueOf(nexthash)) // inject the field
						//h.HashTable[1] = nexthash.Suffix
						log.Printf("factory.builder: inject next fieldname at %v: %s (%v)", n.Type().Name(), n.Type().Field(i+1).Name, n.Field(i+1).FieldByName("Suffix"))
					} else {
						log.Printf("factory.builder: current fieldname %s has no fieldname to the next", n.Type().Field(i).Name)
					}

					log.Printf("factory.builder: >> field injection %d of %d done", i+1, n.NumField())

					break

				}

			}
			l++
		}
		log.Printf("factory.builder: range %dx anonym ", l)
	}

	log.Printf("hashtable %v", h.HashTable)

	// anonymous field aren't reflected
	if ref.hasNotAnonymous() {

		// walk through the map with names of builder fields
		for _, builderField := range ref.Builder {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(builderField)

			if strings.Contains(builderField, fieldProcess) {
				h.NumProc++
			}

			if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == fieldStartEvent {
				h.NumStartEvent++
				h.NumShape++
			}

			if strings.Contains(builderField, fieldTask) && utils.After(builderField, "From") == fieldTask {
				h.NumTask++
				h.NumShape++
			}

			if strings.Contains(builderField, fieldEndEvent) {
				h.NumEndEvent++
				h.NumShape++
			}

			hash := h.hash()             // generate hash value
			n.Set(reflect.ValueOf(hash)) // inject the field

			log.Printf("factory.builder: inject struct field %v", builderField)

		}

		// walk through the map with names of boolean fields
		for _, configField := range ref.Config {

			// get the reflected value of field
			n2 := ref.Temporary.FieldByName(configField)

			// only the first field, which IsExecutable is set to true
			n2.SetBool(true)

			log.Printf("factory.builder: inject config field %v", configField)

		}

	}

	p = ref.Set()
	ref.countWords()

	log.Printf("factory.builder: eof inject; detect %d of reflected public methods", reflect.TypeOf(p).NumMethod())

	return p

}

// build contains the reflected process definition
// and can call SetMainElements from the core package or
// calls it by the reflected method name.
// This method hides some of the setters by building the BPMN
// with reflection
func (h *Builder) build(p interface{}) {

	// el is the interface {}
	el := reflect.ValueOf(&p).Elem()

	// Allocate a temporary variable with type of the struct.
	// el.Elem() is the value contained in the interface
	definitions := reflect.New(el.Elem().Type()).Elem() // *core.Definitions
	definitions.Set(el.Elem())                          // reflected process definitions el will be assigned to the core definitions

	// set collaboration, process and diagram
	collaboration := definitions.MethodByName(MethodSetCollaboration)
	collaboration.Call([]reflect.Value{})

	process := definitions.MethodByName(MethodSetProcess)
	process.Call([]reflect.Value{reflect.ValueOf(h.NumProc)}) // h.numProc represents number of processes

	diagram := definitions.MethodByName(MethodSetDiagram)
	diagram.Call([]reflect.Value{reflect.ValueOf(1)})

}

/*** Semantic, Analyze and Statistic ***/

// inPool ...
func (h *Builder) countPool(field, builderField string) {
	if h.isPool(field) {
		if strings.Contains(builderField, fieldProcess) {
			h.NumProc++
		}
		if strings.Contains(builderField, fieldID) {
			h.NumPart++
			h.NumShape++
		}
	}
}

// inMessage ...
func (h *Builder) countMessage(field, builderField string) {
	if h.isMessage(field) {
		if strings.Contains(builderField, fieldMessage) {
			h.NumMsg++
			h.NumEdge++
		}
	}
}

// isStartEvent ...
func (h *Builder) countStartEvent(builderField string) {
	if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == "" {
		h.NumStartEvent++
		h.NumShape++
	}
}

// isTask ...
func (h *Builder) countTask(builderField string) {
	if strings.Contains(builderField, fieldTask) && utils.After(builderField, "From") == "" {
		h.NumTask++
		h.NumShape++
	}
}

// isEndEvent ...
func (h *Builder) countEndEvent(builderField string) {
	if strings.Contains(builderField, fieldEndEvent) {
		h.NumEndEvent++
		h.NumShape++
	}
}

// isZero ...
func (h *Builder) isZero() bool { return h.Suffix == "" }

// isPool ...
func (h *Builder) isPool(field string) bool { return pool.IsPool(field) }

// isMessage ...
func (h *Builder) isMessage(field string) bool { return elements.IsMessage(field) }

// isUnknown ...
func (h *Builder) isUnknown(field string) bool { return true }
