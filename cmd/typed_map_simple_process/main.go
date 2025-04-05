package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/deemount/gobpmn/internal/parser"
	"github.com/deemount/gobpmn/pkg/core"
	"github.com/deemount/gobpmn/pkg/core/common"
	"github.com/deemount/gobpmn/pkg/core/foundation"
)

type TypedMapSimpleProcess map[string]any

/*func (s TypedMapSimpleProcess) Has(key string) bool {
	_, ok := s[key]
	return ok
}*/

// Testing some logging
var (
	debugLogger,
	errorLogger,
	infoLogger,
	warningLogger *log.Logger
)

func main() {

	flags := log.Ldate | log.Ltime | log.Llongfile
	log.SetFlags(flags)

	debugLogger = log.New(os.Stdout, "DEBUG: ", flags)
	errorLogger = log.New(os.Stdout, "ERROR: ", flags)
	infoLogger = log.New(os.Stdout, "INFO: ", flags)
	warningLogger = log.New(os.Stdout, "WARNING: ", flags)

	// TypedMapSimpleProcess is a typed map that represents a process model
	// with BPMN elements, with named fields.
	type TypedMapSimpleProcess map[string]any
	typedMapSimpleProcess := TypedMapSimpleProcess{
		"Def":                 foundation.NewDefinitions(),
		"IsExecutable":        true,
		"Process":             common.BPMN{Pos: 1},
		"StartEvent":          common.BPMN{Pos: 2},
		"FromStartEvent":      common.BPMN{Pos: 3},
		"Task":                common.BPMN{Pos: 4},
		"FromTask":            common.BPMN{Pos: 5},
		"SecondTask":          common.BPMN{Pos: 6},
		"FromSecondTask":      common.BPMN{Pos: 7},
		"ScriptTask":          common.BPMN{Pos: 8},
		"FromScriptTask":      common.BPMN{Pos: 9},
		"OtherScriptTask":     common.BPMN{Pos: 10},
		"FromOtherScriptTask": common.BPMN{Pos: 11},
		"UserTask":            common.BPMN{Pos: 12},
		"FromUserTask":        common.BPMN{Pos: 13},
		"SpecialUserTask":     common.BPMN{Pos: 14},
		"FromSpecialUserTask": common.BPMN{Pos: 15},
		"EndEvent":            common.BPMN{Pos: 16},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	type T any
	process, err := core.NewReflectDI[T](ctx, typedMapSimpleProcess, map[string]any{})
	if err != nil {
		errorLogger.Fatalf("\033[0;31m\n%s", err)
		return
	}

	// assuming process contains dynamic struct
	value := reflect.ValueOf(process)

	// if process is an interface{}, get the underlying value
	if value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	// get the Def field
	defField := value.FieldByName("Def")
	if !defField.IsValid() {
		fmt.Println("Def field not found")
		return
	}

	// convert to *foundation.Definitions
	def, ok := defField.Interface().(*foundation.Definitions)
	if !ok {
		fmt.Println("Def field is not of type *foundation.Definitions")
		return
	}

	// create the process model
	bpmn, err := parser.NewBPMNParser(
		parser.WithCounter(),
		parser.WithDefinitions(def))
	if err != nil {
		panic(err)
	}
	bpmn.Marshal()

	err = bpmn.Validate()
	if err != nil {
		fmt.Println("Error:", err)
	}

}
