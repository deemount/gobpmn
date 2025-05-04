package main

import (
	"os"
	"time"

	"github.com/deemount/gobpmn"
)

// RentingProcess is a struct that represents a collaborative process model
// with anonymous fields. It is a composition of two processes without a laneset.
type (

	// RentingProcess structure
	RentingProcess struct {

		// The first field must be set to a DefinitionsRepository.
		Def gobpmn.Repository // Refers to the DefinitionsRepository

		// The second field must be set to a pool.
		// A pool must have the name Pool in itself to become identified as such and should/can
		// have the name as the process. Only the first process in the model is executable.
		RentingPool // RentingPool represents a collaboration

		// All anonymous fields, which are set after a pool, are considered as processes.
		Tenant   // Tenant represents a process
		Landlord // Landlord represents a process
	}

	// RentingPool ...
	RentingPool struct {
		TenantIsExecutable   bool        // Process Configuration
		LandlordIsExecutable bool        // Process Configuration
		Collaboration        gobpmn.BPMN // BPMN Element
		TenantProcess        gobpmn.BPMN // BPMN Element
		TenantParticipant    gobpmn.BPMN // BPMN Element
		LandlordProcess      gobpmn.BPMN // BPMN Element
		LandlordParticipant  gobpmn.BPMN // BPMN Element
	}

	// Tenant
	Tenant struct {
		StartEvent     gobpmn.BPMN // BPMN Element
		FromStartEvent gobpmn.BPMN // Flow Element
		Task           gobpmn.BPMN // BPMN Element
		FromTask       gobpmn.BPMN // Flow Element
		EndEvent       gobpmn.BPMN // BPMN Element
	}

	// Landlord
	Landlord struct {
		StartEvent     gobpmn.BPMN // BPMN Element
		FromStartEvent gobpmn.BPMN // Flow Element
		FirstTask      gobpmn.BPMN // BPMN Element
		FromFirstTask  gobpmn.BPMN // Flow Element
		SecondTask     gobpmn.BPMN // BPMN Element
		FromSecondTask gobpmn.BPMN // Flow Element
		ScriptTask     gobpmn.BPMN // BPMN Element
		FromScriptTask gobpmn.BPMN // Flow Element
		EndEvent       gobpmn.BPMN // BPMN Element
	}
)

func (rp RentingProcess) GetDefinitions() gobpmn.Repository {
	return rp.Def
}

func main() {

	log := gobpmn.NewLogger("renting", gobpmn.DebugLevel, os.Stdout)
	if err := log.SetOutputFile("files/log/renting_process.log"); err != nil {
		log.Errorf("could not open log file: %v", err)
	}

	opts := gobpmn.Options{
		Timeout:  60 * time.Second,
		Validate: true,
		Logger:   log,
	}

	model, err := gobpmn.FromStruct(RentingProcess{}, opts)
	if err != nil {
		log.Errorf("bpmn generation failed: %v", err)
		os.Exit(1)
	}

	log.Infof("bpmn model generated successfully: %+v", model)

}
