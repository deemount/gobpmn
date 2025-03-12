//go:build go1.23
// +build go1.23

package main

import (
	"context"
	"log"
	_ "net/http/pprof"
	"time"

	"github.com/deemount/gobpmn/internal/parser"
	"github.com/deemount/gobpmn/pkg/core/common"
	"github.com/deemount/gobpmn/pkg/core/foundation"
)

var bpmnParser parser.BPMNParser

// SimpleProcess is a struct that represents a process model
// with BPMN elements, with named fields.
type SimpleProcess struct {

	// Def MUST be set to a DefinitionsRepository,
	// otherwise the model will not be valid. It MUST be set at the first place.
	Def foundation.DefinitionsRepository // Refers to the DefinitionsRepository

	// If IsExecutable is set, the process is executable and set to true.
	// Otherwise, the process is not executable and set to false.
	// If more than one process in a model is given, only one process can be executable.
	// It will be then the first process given in the model, which will be executable.
	IsExecutable bool // Process Configuration

	// All elements of the BPMN model
	Process             common.BPMN // BPMN Element
	StartEvent          common.BPMN // BPMN Element
	FromStartEvent      common.BPMN // BPMN Element
	Task                common.BPMN // BPMN Element
	FromTask            common.BPMN // BPMN Element
	SecondTask          common.BPMN // BPMN Element
	FromSecondTask      common.BPMN // BPMN Element
	ScriptTask          common.BPMN // BPMN Element
	FromScriptTask      common.BPMN // BPMN Element
	OtherScriptTask     common.BPMN // BPMN Element
	FromOtherScriptTask common.BPMN // BPMN Element
	UserTask            common.BPMN // BPMN Element
	FromUserTask        common.BPMN // BPMN Element
	SpecialUserTask     common.BPMN // BPMN Element
	FromSpecialUserTask common.BPMN // BPMN Element
	EndEvent            common.BPMN // BPMN Element
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	/*
	 * SimpleProcess
	 */
	simpleProcess, err := common.NewReflectDI[SimpleProcess](ctx, SimpleProcess{})
	if err != nil {
		log.Fatal(err)
	}

	// logging messages for the terminal
	/*
		log.Printf("simpleProcess.Target: %+#v", simpleProcess) // represents the Target
		log.Print("-------------------------")
		log.Printf("simpleProcess.Def: %+#v", simpleProcess.Def) // represents the model to create
	*/

	// create the process model
	bpmn, err := parser.NewBPMNParser(
		parser.WithCounter(),
		parser.WithDefinitions(simpleProcess.Def))
	if err != nil {
		panic(err)
	}
	bpmn.Marshal()

}
