package factory

import (
	"log"
	"reflect"

	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/utils"
)

type Reflect struct {
	Analyze
	IF        interface{}
	Temporary reflect.Value
	Element   reflect.Value
}

// NewReflect initialize interface {} p
func NewReflect(p interface{}) *Reflect {
	return &Reflect{IF: p}
}

// Interface where Element is the interface {}
func (ref *Reflect) Interface() *Reflect {
	ref.Element = reflect.ValueOf(&ref.IF).Elem()
	return ref
}

// New allocates a temporary variable with type of the struct.
// ref.Element.Elem() is the value contained in the interface
func (ref *Reflect) New() *Reflect {
	ref.Temporary = reflect.New(ref.Element.Elem().Type()).Elem()
	ref.Temporary.Set(ref.Element.Elem())
	return ref
}

// Make initializes maps to analyze later
// anonym: all anonymous fields
// config: all boolean fields
// builder: all builder fields
// words: all collected words splitted
func (ref *Reflect) Maps() {
	ref.anonym = make(map[int]string)
	ref.config = make(map[int]string)
	ref.builder = make(map[int]string)
	ref.words = make(map[int][]string)
}

// Set temporary variable values to interface {}
func (ref *Reflect) Set() any {
	ref.Element.Set(ref.Temporary)
	return ref.Element.Interface()
}

// reflectAppAnonymousFields takes all compounds which are anonymous
// The implementation is O(N) for insertions
func (ref *Reflect) reflectAnonymousFields() {

	fields := reflect.VisibleFields(reflect.TypeOf(ref.IF))
	i := 0
	for _, field := range fields {
		if ref.isAnonymous(field) {
			ref.anonym[i] = field.Name
			ref.words[i] = utils.Split(ref.anonym[i])
			i++
		}
	}
	if i == 0 {
		log.Println("factory.reflect: main struct contains map of zero anonymous fields")
	}

}

// reflectBuilderType with three static filter options
func (ref Reflect) reflectBuilderType() {
	fields := reflect.VisibleFields(reflect.TypeOf(ref.IF))
	count := 0
	index := len(ref.words)
	for _, field := range fields {
		if ref.isNotAnonymous(field) && ref.isNotDefinitions(field) && ref.isBpmnBuilder(field) {
			ref.builder[count] = field.Name
			ref.words[index] = utils.Split(ref.builder[count])
			count++
			index++
		}
	}
}

// reflectBoolType with one static filter option, which must be kind of reflect.Bool
// The bool type of a field in a struct describes mostly configuration settings
// The field must contain a sibling in title case, e.g. IsExecutable
func (ref Reflect) reflectBoolType() {
	fields := reflect.VisibleFields(reflect.TypeOf(ref.IF))
	count := 0
	index := len(ref.words)
	for _, field := range fields {
		if field.Type.Kind() == reflect.Bool {
			ref.config[count] = field.Name
			ref.words[index] = utils.Split(ref.config[count])
			count++
			index++
		}
	}
}

// countWords ...
func (ref Reflect) countWords() {
	l := 0
	length := len(ref.words)
	for i := 0; i < length; i++ {
		l += len(ref.words[i])
	}
	log.Printf("factory.reflect: count app words %v", l)
}

// isNotDefinitions ...
func (ref *Reflect) isDefinitions(field reflect.StructField) bool { return core.IsDefinitions(field) }

// isNotDefinitions ...
func (ref *Reflect) isNotDefinitions(field reflect.StructField) bool {
	return core.IsNotDefinitions(field)
}

// hasAnonymous ...
func (ref *Reflect) hasAnonymous() bool { return len(ref.anonym) > 0 }

// hasNotAnonymous ...
func (ref *Reflect) hasNotAnonymous() bool { return len(ref.anonym) == 0 }

// isAnonymous ...
func (ref *Reflect) isAnonymous(field reflect.StructField) bool { return field.Anonymous }

// isNotAnonymous ...
func (ref *Reflect) isNotAnonymous(field reflect.StructField) bool { return !field.Anonymous }

// isBpmnBuilder ...
func (ref *Reflect) isBpmnBuilder(field reflect.StructField) bool {
	return field.Type.Name() == typeBuilder
}
