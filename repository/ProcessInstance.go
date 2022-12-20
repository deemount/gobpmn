package repository

import (
	"context"
	"log"
	"time"

	"github.com/deemount/gobpmn/models/events"
	"github.com/deemount/gobpmn/spec/process_instance"
)

type ProcessInstance struct {
	engine BpmnEngine
}

func NewProcessInstance(e BpmnEngine) ProcessInstance {
	return ProcessInstance{
		engine: e,
	}
}

// GetProcessInstances returns a list of instance information.
func (processInstance *ProcessInstance) GetProcessInstances() []*ProcessInstanceInfo {
	return processInstance.engine.ProcessInstances
}

// GetProcessInfo ...
func (processInstance *ProcessInstance) GetProcessInfo(ctx context.Context, b BpmnFactory) (*ProcessInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	done := make(chan func() (*ProcessInfo, error))
	go processInstance.engine.requestFile(done, b)
	processInfo, err := (<-done)()
	return processInfo, err
}

// CreateInstance creates a new instance for a process with given processKey
func (processInstance *ProcessInstance) Create(processKey int64, variableContext map[string]interface{}) (*ProcessInstanceInfo, error) {
	if variableContext == nil {
		variableContext = map[string]interface{}{}
	}
	for _, process := range processInstance.engine.Processes {
		if process.ProcessKey == processKey {
			processInstanceInfo := ProcessInstanceInfo{
				processInfo:     &process,
				instanceKey:     processInstance.engine.generateKey(),
				variableContext: variableContext,
				createdAt:       time.Now(),
				state:           process_instance.READY,
			}
			processInstance.engine.ProcessInstances = append(processInstance.engine.ProcessInstances, &processInstanceInfo)
			return &processInstanceInfo, nil
		}
	}
	return nil, nil
}

// Run ...
func (processInstance *ProcessInstance) Run(inst *ProcessInstanceInfo) error {

	type queueElement struct {
		inboundFlowId string
		baseElement   events.TStartEvent
	}

	queue := make([]queueElement, 0)
	processInfo := inst.processInfo

	switch inst.state {
	case process_instance.READY:

		// use start events to start the instance
		for _, proc := range processInfo.Def.Process {
			for _, event := range proc.StartEvent {
				queue = append(queue, queueElement{
					inboundFlowId: "",
					baseElement:   event,
				})
			}
		}

		inst.state = process_instance.ACTIVE

	case process_instance.ACTIVE:

		/*intermediateCatchEvents := state.findIntermediateCatchEventsForContinuation(process, instance)
		for _, ice := range intermediateCatchEvents {
			queue = append(queue, queueElement{
				inboundFlowId: "",
				baseElement:   ice,
			})
		}*/

	case process_instance.COMPLETED:
		return nil

	default:
		panic("Unknown process instance state")
	}

	log.Printf("processinstance -> queue: %+v", queue)

	/*for len(queue) > 0 {

		element := queue[0].baseElement
		inboundFlowId := queue[0].inboundFlowId
		queue = queue[1:]

		continueNextElement := state.handleElement(process, instance, element)

		if continueNextElement {
			if inboundFlowId != "" {
				state.scheduledFlows = remove(state.scheduledFlows, inboundFlowId)
			}
			nextFlows := BPMN20.FindSequenceFlows(&process.definitions.Process.SequenceFlows, element.GetOutgoingAssociation())
			if element.GetType() == BPMN20.ExclusiveGateway {
				nextFlows = exclusivelyFilterByConditionExpression(nextFlows, instance.variableContext)
			}
			for _, flow := range nextFlows {
				// TODO: create test for that
				//if len(flows) < 1 {
				//	panic(fmt.Sprintf("Can't find 'sequenceFlow' element with ID=%s. "+
				//		"This is likely because your BPMN is invalid.", flows[0]))
				//}
				state.scheduledFlows = append(state.scheduledFlows, flow.Id)
				baseElements := BPMN20.FindBaseElementsById(process.definitions, flow.TargetRef)
				// TODO: create test for that
				//if len(baseElements) < 1 {
				//	panic(fmt.Sprintf("Can't find flow element with ID=%s. "+
				//		"This is likely because there are elements in the definition, "+
				//		"which this engine does not support (yet).", flow.Id))
				//}
				targetBaseElement := baseElements[0]
				queue = append(queue, queueElement{
					inboundFlowId: flow.Id,
					baseElement:   targetBaseElement,
				})
			}
		}
	}*/
	return nil
}
