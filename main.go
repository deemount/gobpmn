//go:build go1.23
// +build go1.23

package main

import (
	"context"
	"log"
	_ "net/http/pprof"
	"time"
)

var bpmnModeller BPMNModeller

// ExampleProcess is a struct that represents a process model
// with BPMN elements, with named fields.
type ExampleProcess struct {

	// Def MUST be set to a DefinitionsRepository,
	// otherwise the model will not be valid. It MUST be set at the first place.
	Def DefinitionsRepository // Refers to the DefinitionsRepository

	// If IsExecutable is set, the process is executable and set to true.
	// Otherwise, the process is not executable and set to false.
	// If more than one process in a model is given, only one process can be executable.
	// It will be then the first process given in the model, which will be executable.
	IsExecutable bool // Process Configuration

	// All elements of the BPMN model
	Process             BPMN // BPMN Element
	StartEvent          BPMN // BPMN Element
	FromStartEvent      BPMN // BPMN Element
	Task                BPMN // BPMN Element
	FromTask            BPMN // BPMN Element
	SecondTask          BPMN // BPMN Element
	FromSecondTask      BPMN // BPMN Element
	ScriptTask          BPMN // BPMN Element
	FromScriptTask      BPMN // BPMN Element
	OtherScriptTask     BPMN // BPMN Element
	FromOtherScriptTask BPMN // BPMN Element
	UserTask            BPMN // BPMN Element
	FromUserTask        BPMN // BPMN Element
	SpecialUserTask     BPMN // BPMN Element
	FromSpecialUserTask BPMN // BPMN Element
	EndEvent            BPMN // BPMN Element
}

// RentingProcess is a struct that represents a collaborative process model
// with anonymous fields. It is a composition of two processes.
type (

	// RentingProcess structure
	RentingProcess struct {

		// The first field must be set to a DefinitionsRepository.
		Def DefinitionsRepository // Refers to the DefinitionsRepository

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
		TenantIsExecutable   bool // Process Configuration
		LandlordIsExecutable bool // Process Configuration
		Collaboration        BPMN // BPMN Element
		TenantProcess        BPMN // BPMN Element
		TenantParticipant    BPMN // BPMN Element
		LandlordProcess      BPMN // BPMN Element
		LandlordParticipant  BPMN // BPMN Element
	}

	// Tenant
	Tenant struct {
		StartEvent     BPMN // BPMN Element
		FromStartEvent BPMN // BPMN Element
		Task           BPMN // BPMN Element
		FromTask       BPMN // BPMN Element
		EndEvent       BPMN // BPMN Element
	}

	// Landlord
	Landlord struct {
		StartEvent     BPMN // BPMN Element
		FromStartEvent BPMN // BPMN Element
		FirstTask      BPMN // BPMN Element
		FromFirstTask  BPMN // BPMN Element
		SecondTask     BPMN // BPMN Element
		FromSecondTask BPMN // BPMN Element
		ScriptTask     BPMN // BPMN Element
		FromScriptTask BPMN // BPMN Element
		EndEvent       BPMN // BPMN Element
	}
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	/*
	 * ExampleProcess
	 */

	exampleProcess, err := NewReflectDI[ExampleProcess](ctx, ExampleProcess{})
	if err != nil {
		log.Fatal(err)
	}

	// some logging messages for the terminal
	/*
		log.Printf("exampleProcess.Target: %+#v", exampleProcess) // represents the Target
		log.Print("-------------------------")
		log.Printf("exampleProcess.Def: %+#v", exampleProcess.Def) // represents the model to create
	*/

	// create the process model
	example, err := NewBPMNModeller(
		WithCounter(),
		WithDefinitions(exampleProcess.Def))
	if err != nil {
		panic(err)
	}
	example.Marshal()

	//log.Print("--NEW-------------------------")

	/*
	 * RentalProcess
	 */
	/*
		rentalProcess, err := NewReflectDI[RentingProcess](ctx, RentingProcess{})
		if err != nil {
			log.Printf("Error: %s", err)
			return
		}
	*/
	// logging messages for the terminal
	/*
		log.Printf("rentalProcess.Target: %+#v", rentalProcess) // represents the Target
		log.Print("-------------------------")
		log.Printf("rentalProcess.Def: %+#v", rentalProcess.Def) // represents the model to create

		rental, err := NewBPMNModeller(
			WithCounter(),
			WithDefinitions(rentalProcess.Def))
		if err != nil {
			panic(err)
		}
		rental.Marshal()
	*/

}
