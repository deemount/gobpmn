package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/deemount/gobpmn/pkg/core"
	"github.com/deemount/gobpmn/pkg/core/common"
	"github.com/deemount/gobpmn/pkg/core/foundation"
)

type MappedSimpleProcess map[string]any

func (s MappedSimpleProcess) Has(key string) bool {
	_, ok := s[key]
	return ok
}

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

	mappedSimpleProcess := MappedSimpleProcess{
		"Def":                 foundation.NewDefinitions(),
		"IsExecutable":        true,
		"Process":             common.BPMN{Pos: 0},
		"StartEvent":          common.BPMN{Pos: 1},
		"FromStartEvent":      common.BPMN{Pos: 2},
		"Task":                common.BPMN{Pos: 3},
		"FromTask":            common.BPMN{Pos: 4},
		"SecondTask":          common.BPMN{Pos: 5},
		"FromSecondTask":      common.BPMN{Pos: 6},
		"ScriptTask":          common.BPMN{Pos: 7},
		"FromScriptTask":      common.BPMN{Pos: 8},
		"OtherScriptTask":     common.BPMN{Pos: 9},
		"FromOtherScriptTask": common.BPMN{Pos: 10},
		"UserTask":            common.BPMN{Pos: 11},
		"FromUserTask":        common.BPMN{Pos: 12},
		"SpecialUserTask":     common.BPMN{Pos: 13},
		"FromSpecialUserTask": common.BPMN{Pos: 14},
		"EndEvent":            common.BPMN{Pos: 15},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	process, err := core.NewReflectDI[MappedSimpleProcess, map[string]any](ctx, mappedSimpleProcess)
	if err != nil {
		errorLogger.Fatalf("\033[0;31m\n%s", err)
		return
	}

	debugLogger.Printf("SimpleProcess: %+v", process)

}
