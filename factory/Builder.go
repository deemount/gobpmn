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

// Builder ...
type Builder struct {
	Reflect
	Settings
	Suffix string
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

// inject is one of hte main ideas behind the project
// The method itself reflects a given struct and inject
// signed fields with hash values. I chose this type of algorithm,
// because it does the logic behind the building of a bpmn.
// Kinda magic ...
func (h *Builder) inject(p interface{}) interface{} {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("factory.builder: inject recovered", r)
		}
	}()

	log.Println("factory.builder: start injecting fields")

	ref := NewReflect(p)
	ref.Interface().New().Maps().Reflection()

	// TODO: Describe how the injection rules for anonymous fields are
	// anonymous field are reflected
	if ref.hasAnonymous() {

		// walk through the map with names of anonymous fields
		for _, field := range ref.anonym {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(field)

			// walk through the values of fields by the interface {}
			for i := 0; i < n.NumField(); i++ {

				name := n.Type().Field(i).Name
				if strings.Contains(name, field) {
					log.Printf("factory.builder: detected field %s contains anonymous field %s", name, field)
				}

				// set by kind of reflected value above
				switch n.Field(i).Kind() {

				// kind is a struct
				case reflect.Struct:

					h.inPool(field, name)
					h.inMessage(field, name)

					hash := h.hash()                      // generate hash value
					n.Field(i).Set(reflect.ValueOf(hash)) // inject the field

					log.Printf("factory.builder: inject struct field %s", name)

					break

				// kind is a bool
				case reflect.Bool:

					// only the first field, which IsExecutable is set to true,
					// means, only one process in a collaboration can be executed at runtime
					// this can be changed in the future, if the engine fits for more execution
					// options
					if strings.Contains(name, boolIsExecutable) && i == 0 {
						n.Field(0).SetBool(true)
						log.Printf("factory.builder: inject first bool field %s once", name)
					}

					break
				}
			}
		}

	}

	// anonymous field aren't reflected
	if ref.hasNotAnonymous() {

		// walk through the map with names of builder fields
		for _, builderField := range ref.builder {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(builderField)

			if strings.Contains(builderField, fieldProcess) {
				h.NumProc++
			}

			if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == fieldStartEvent {
				h.NumStartEvent++
			}

			if strings.Contains(builderField, fieldTask) && utils.After(builderField, "From") == fieldTask {
				h.NumTask++
			}

			if strings.Contains(builderField, fieldEndEvent) {
				h.NumEndEvent++
			}

			hash := h.hash()             // generate hash value
			n.Set(reflect.ValueOf(hash)) // inject the field

			log.Printf("factory.builder: inject struct field %v", builderField)

		}

		// walk through the map with names of boolean fields
		for _, configField := range ref.config {

			// get the reflected value of field
			n2 := ref.Temporary.FieldByName(configField)

			// only the first field, which IsExecutable is set to true
			n2.SetBool(true)

			log.Printf("factory.builder: inject config field %v", configField)

		}

	}

	p = ref.Set()
	ref.countWords()

	log.Printf("factory.builder: eof inject number of reflected methods %v", reflect.TypeOf(p).NumMethod())

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

// inPool ...
func (h *Builder) inPool(field, builderField string) {
	if h.isPool(field) {
		if strings.Contains(builderField, fieldProcess) {
			h.NumProc++
		}
		if strings.Contains(builderField, fieldID) {
			h.NumPart++
		}
	}
}

// inMessage ...
func (h *Builder) inMessage(field, builderField string) {
	if h.isMessage(field) {
		if strings.Contains(builderField, fieldMessage) {
			h.NumMsg++
		}
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
