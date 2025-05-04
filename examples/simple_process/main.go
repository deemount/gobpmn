//go:build go1.24
// +build go1.24

package main

import (
	"os"
	"time"

	"github.com/deemount/gobpmn"
)

// SimpleProcess is a struct that represents a process model
// with BPMN elements, with named fields.
type SimpleProcess struct {

	// Def MUST be set to a DefinitionsRepository,
	// otherwise the model will not be valid. It MUST be set at the first place.
	Def gobpmn.Repository // Refers to the DefinitionsRepository

	// If IsExecutable is set, the process is executable and set to true.
	// Otherwise, the process is not executable and set to false.
	// If more than one process in a model is given, only one process can be executable.
	// It will be then the first process given in the model, which will be executable.
	IsExecutable bool // Process Configuration

	// All elements of the BPMN model
	Process             gobpmn.BPMN // BPMN Element
	StartEvent          gobpmn.BPMN // BPMN Element
	FromStartEvent      gobpmn.BPMN // BPMN Element
	Task                gobpmn.BPMN // BPMN Element
	FromTask            gobpmn.BPMN // BPMN Element
	SecondTask          gobpmn.BPMN // BPMN Element
	FromSecondTask      gobpmn.BPMN // BPMN Element
	ScriptTask          gobpmn.BPMN // BPMN Element
	FromScriptTask      gobpmn.BPMN // BPMN Element
	OtherScriptTask     gobpmn.BPMN // BPMN Element
	FromOtherScriptTask gobpmn.BPMN // BPMN Element
	UserTask            gobpmn.BPMN // BPMN Element
	FromUserTask        gobpmn.BPMN // BPMN Element
	SpecialUserTask     gobpmn.BPMN // BPMN Element
	FromSpecialUserTask gobpmn.BPMN // BPMN Element
	EndEvent            gobpmn.BPMN // BPMN Element
}

func (rp SimpleProcess) GetDefinitions() gobpmn.Repository {
	return rp.Def
}

func main() {

	log := gobpmn.NewLogger("simple", gobpmn.DebugLevel, os.Stdout)
	if err := log.SetOutputFile("files/log/simple_process.log"); err != nil {
		log.Errorf("could not open log file: %v", err)
	}

	opts := gobpmn.Options{
		Timeout:  60 * time.Second,
		Validate: true,
		Logger:   log,
	}

	model, err := gobpmn.FromStruct(SimpleProcess{}, opts)
	if err != nil {
		log.Errorf("bpmn generation failed: %v", err)
		os.Exit(1)
	}

	log.Infof("bpmn model generated successfully: %+v", model)

}
