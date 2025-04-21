package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/deemount/gobpmn/pkg/core/foundation"
	"github.com/deemount/gobpmn/pkg/core/infrastructure"
)

// Validate checks, if all incoming, outgoing, sourceRef and targetRef values are valid.
func (parser BPMNParser) Validate() error {

	jsonFilename := parser.GetFilename()
	jsonFilename = jsonFilename + ".json"
	jsonFilepath := parser.FilePathJSON + "/" + jsonFilename

	data, err := os.ReadFile(jsonFilepath)
	if err != nil {
		return fmt.Errorf("failed reading json file %v", err)
	}

	// parse json in definitions struct
	var definitions foundation.Definitions
	err = json.Unmarshal(data, &definitions)
	if err != nil {
		return fmt.Errorf("failed parsing json file %v", err)
	}

	// save all element IDs and flow IDs
	elementIDs := make(map[string]bool)
	flowIDs := make(map[string]infrastructure.SequenceFlow)

	// run through the processes
	for _, process := range definitions.Process {
		// collect the id's bpmn elements
		for _, el := range process.StartEvent {
			elementIDs[el.ID] = true
		}
		for _, el := range process.EndEvent {
			elementIDs[el.ID] = true
		}
		for _, el := range process.IntermediateCatchEvent {
			elementIDs[el.ID] = true
		}
		for _, el := range process.IntermediateThrowEvent {
			elementIDs[el.ID] = true
		}
		for _, el := range process.ExclusiveGateway {
			elementIDs[el.ID] = true
		}
		for _, el := range process.ParallelGateway {
			elementIDs[el.ID] = true
		}
		for _, el := range process.InclusiveGateway {
			elementIDs[el.ID] = true
		}
		for _, el := range process.UserTask {
			elementIDs[el.ID] = true
		}
		for _, el := range process.ServiceTask {
			elementIDs[el.ID] = true
		}
		for _, el := range process.ScriptTask {
			elementIDs[el.ID] = true
		}
		for _, el := range process.Task {
			elementIDs[el.ID] = true
		}

		// save flows
		for _, flow := range process.SequenceFlow {
			flowIDs[flow.ID] = flow
		}
	}

	// validation of incoming & outgoing flows
	for _, process := range definitions.Process {
		for _, el := range process.IntermediateCatchEvent {
			validateFlow("IntermediateCatchEvent", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
		for _, el := range process.IntermediateThrowEvent {
			validateFlow("IntermediateThrowEvent", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
		for _, el := range process.ExclusiveGateway {
			validateFlow("ExclusiveGateway", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
		for _, el := range process.ParallelGateway {
			validateFlow("ParallelGateway", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
		for _, el := range process.InclusiveGateway {
			validateFlow("InclusiveGateway", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
		for _, el := range process.UserTask {
			validateFlow("UserTask", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
		for _, el := range process.ScriptTask {
			validateFlow("ScriptTask", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
		for _, el := range process.ServiceTask {
			validateFlow("ScriptTask", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
		for _, el := range process.Task {
			validateFlow("Task", el.ID, el.Incoming, el.Outgoing, flowIDs, elementIDs)
		}
	}

	// validate the sequence flows
	for _, process := range definitions.Process {
		for _, flow := range process.SequenceFlow {
			if _, exists := elementIDs[flow.SourceRef]; !exists {
				fmt.Printf("invalid SourceRef: %s in Flow %s\n", flow.SourceRef, flow.ID)
			}
			if _, exists := elementIDs[flow.TargetRef]; !exists {
				fmt.Printf("invalid TargetRef: %s in Flow %s\n", flow.TargetRef, flow.ID)
			}
		}
	}

	fmt.Println("finished validation")
	return nil
}

// validateFlow checks incoming and outgoing connections for bpmn elements.
func validateFlow(elementType, elementID string, incoming []infrastructure.Incoming, outgoing []infrastructure.Outgoing, flowIDs map[string]infrastructure.SequenceFlow, elementIDs map[string]bool) {
	for _, inc := range incoming {
		if _, exists := flowIDs[inc.Flow]; !exists {
			fmt.Printf("missing incoming flow: %s at %s %s\n", inc.Flow, elementType, elementID)
		}
	}
	for _, out := range outgoing {
		if _, exists := flowIDs[out.Flow]; !exists {
			fmt.Printf("missing outgoing flow: %s at %s %s\n", out.Flow, elementType, elementID)
		}
	}
}
