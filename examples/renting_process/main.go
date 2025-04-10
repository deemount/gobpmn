//go:build go1.23
// +build go1.23

package main

import (
	"context"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"reflect"
	"time"

	"github.com/deemount/gobpmn/internal/parser"
	"github.com/deemount/gobpmn/pkg/core"
	"github.com/deemount/gobpmn/pkg/core/common"
	"github.com/deemount/gobpmn/pkg/core/foundation"
)

var bpmnParser parser.BPMNParser

// RentingProcess is a struct that represents a collaborative process model
// with anonymous fields. It is a composition of two processes.
type (

	// RentingProcess structure
	RentingProcess struct {

		// The first field must be set to a DefinitionsRepository.
		Def foundation.DefinitionsRepository // Refers to the DefinitionsRepository

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
		Collaboration        common.BPMN // BPMN Element
		TenantProcess        common.BPMN // BPMN Element
		TenantParticipant    common.BPMN // BPMN Element
		LandlordProcess      common.BPMN // BPMN Element
		LandlordParticipant  common.BPMN // BPMN Element
	}

	// Tenant
	Tenant struct {
		StartEvent     common.BPMN // BPMN Element
		FromStartEvent common.BPMN // BPMN Element
		Task           common.BPMN // BPMN Element
		FromTask       common.BPMN // BPMN Element
		EndEvent       common.BPMN // BPMN Element
	}

	// Landlord
	Landlord struct {
		StartEvent     common.BPMN // BPMN Element
		FromStartEvent common.BPMN // BPMN Element
		FirstTask      common.BPMN // BPMN Element
		FromFirstTask  common.BPMN // BPMN Element
		SecondTask     common.BPMN // BPMN Element
		FromSecondTask common.BPMN // BPMN Element
		ScriptTask     common.BPMN // BPMN Element
		FromScriptTask common.BPMN // BPMN Element
		EndEvent       common.BPMN // BPMN Element
	}
)

func main() {

	flags := log.Ldate | log.Lshortfile
	log.SetFlags(flags)
	errorLogger := log.New(os.Stdout, "ERROR: ", flags)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	/*
	 * RentalProcess
	 */
	rentalProcess, err := core.NewReflectDI(ctx, RentingProcess{}, []reflect.StructField{})
	if err != nil {
		errorLogger.Fatalf("\033[0;31m:\n%s", err)
		return
	}

	// logging messages for the terminal
	/*
		log.Printf("rentalProcess.Target: %+#v", rentalProcess) // represents the Target
		log.Print("-------------------------")
		log.Printf("rentalProcess.Def: %+#v", rentalProcess.Def) // represents the model to create
	*/

	// create the process model
	bpmn, err := parser.NewBPMNParser(
		parser.WithCounter(),
		parser.WithDefinitions(rentalProcess.Def))
	if err != nil {
		panic(err)
	}
	bpmn.Marshal()

	err = bpmn.Validate()
	if err != nil {
		fmt.Println("Error:", err)
	}

}
