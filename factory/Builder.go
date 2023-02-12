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

	c.Write([]byte(s))
	defer c.Reset()

	r := Builder{Suffix: fmt.Sprintf("%x", string(c.Sum(nil)))}

	return r
}

// inject ...
func (h *Builder) inject(p interface{}) interface{} {

	log.Println("factory.builder: start injecting fields")

	ref := NewReflect(p)
	ref.Interface().New().Maps()

	// get anonymous field from interface {}
	ref.reflectAnonymousFields()

	// after reflected anonymous fields
	ref.reflectBuilderType()
	ref.reflectBoolType()

	// anonymous field are reflected
	if ref.hasAnonymous() {

		// walk through the map with names of anonymous fields
		for _, field := range ref.anonym {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(field)

			// walk through the values of fields by the interface {}
			for i := 0; i < n.NumField(); i++ {

				// set by kind of reflected value above
				switch n.Field(i).Kind() {

				// kind is a struct
				case reflect.Struct:

					builderField := n.Type().Field(i).Name

					if h.isPool(field) {
						if strings.Contains(builderField, fieldProcess) {
							h.NumProc++
						}
						if strings.Contains(builderField, fieldID) {
							h.NumPart++
						}
					}

					if h.isMessage(field) {
						if strings.Contains(builderField, fieldMessage) {
							h.NumMsg++
						}
					}

					if strings.Contains(builderField, field) {

					}

					/*
						if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == fieldStartEvent {
							h.NumStartEvent++
						}

						if strings.Contains(builderField, fieldEndEvent) && utils.After(builderField, "From") == fieldEndEvent {
							h.NumEndEvent++
						}
					*/

					// generate hash value
					hash := h.hash()
					n.Field(i).Set(reflect.ValueOf(hash))

					log.Printf("factory.builder: inject struct field %s", builderField)

					break

				// kind is a bool
				case reflect.Bool:

					val := n.Type().Field(i).Name

					// only the first field, which IsExecutable is set to true,
					// means, only one process in a collaboration can be executed at runtime
					// this can be changed in the future, if the engine fits for more execution
					// options
					if strings.Contains(val, boolIsExecutable) && i == 0 {
						n.Field(0).SetBool(true)
						log.Printf("factory.builder: inject bool field %s", val)
					}

					break
				}
			}
		}

		p = ref.Set()

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

			// generate hash value
			hash := h.hash()
			n.Set(reflect.ValueOf(hash))

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

		p = ref.Set()

	}

	ref.countWords()

	return p

}

// build ...
func (h *Builder) build(p interface{}) {
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

// isZero ...
func (h *Builder) isZero() bool { return h.Suffix == "" }

// isPool ...
func (h *Builder) isPool(field string) bool { return pool.IsPool(field) }

// isMessage ...
func (h *Builder) isMessage(field string) bool { return elements.IsMessage(field) }

// isUnknown ...
func (h *Builder) isUnknown(field string) bool { return true }
