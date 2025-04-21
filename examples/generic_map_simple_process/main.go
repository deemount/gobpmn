package main

import (
	"log"

	"github.com/deemount/gobpmn"
)

// Testing some logging
var errorLogger *log.Logger

// Process is a generic map that represents a process model
// with BPMN elements, with named fields.
var Process = map[string]any{
	"Def":                 gobpmn.Definitions(),
	"IsExecutable":        true,
	"Process":             gobpmn.BPMN{Pos: 1},
	"StartEvent":          gobpmn.BPMN{Pos: 2},
	"FromStartEvent":      gobpmn.BPMN{Pos: 3},
	"Task":                gobpmn.BPMN{Pos: 4},
	"FromTask":            gobpmn.BPMN{Pos: 5},
	"SecondTask":          gobpmn.BPMN{Pos: 6},
	"FromSecondTask":      gobpmn.BPMN{Pos: 7},
	"ScriptTask":          gobpmn.BPMN{Pos: 8},
	"FromScriptTask":      gobpmn.BPMN{Pos: 9},
	"OtherScriptTask":     gobpmn.BPMN{Pos: 10},
	"FromOtherScriptTask": gobpmn.BPMN{Pos: 11},
	"UserTask":            gobpmn.BPMN{Pos: 12},
	"FromUserTask":        gobpmn.BPMN{Pos: 13},
	"SpecialUserTask":     gobpmn.BPMN{Pos: 14},
	"FromSpecialUserTask": gobpmn.BPMN{Pos: 15},
	"EndEvent":            gobpmn.BPMN{Pos: 16},
}

func main() {

	flags := log.Ldate | log.Ltime | log.Llongfile
	log.SetFlags(flags)

	type T any
	_, err := gobpmn.FromMap[T](Process)
	if err != nil {
		errorLogger.Fatalf("\033[0;31m:\n%s", err)
		return
	}

}
