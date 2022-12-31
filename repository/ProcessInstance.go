package repository

import (
	"context"
	"log"
	"time"

	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/events/elements"
	"github.com/deemount/gobpmn/spec/process_instance"
	"github.com/deemount/gobpmn/workflows"
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

	queue := make([]workflows.EventQueue, 0)
	processInfo := inst.processInfo

	switch inst.state {
	case process_instance.READY:

		// use start events to start the instance
		for _, proc := range processInfo.Def.Process {
			for _, event := range proc.StartEvent {
				queue = append(queue, workflows.EventQueue{
					InboundFlowId: "",
					Element:       event,
				})
			}
		}

		inst.state = process_instance.ACTIVE

	case process_instance.ACTIVE:

	case process_instance.COMPLETED:
		return nil

	default:
		panic("Unknown process instance state")
	}

	// get the id of startevent: queue[0].baseElement.BaseAttributes.ID
	log.Printf("processinstance -> queue: %+v", queue)

	dispatcher := workflows.NewDispatcher()
	dispatcher.Register(queue[0].Element)

	log.Printf("processinstance -> dispatcher: %+v ", dispatcher)

	go func() {
		err := dispatcher.Dispatch(context.Background(), elements.TStartEvent{
			BaseAttributes: compulsion.BaseAttributes{
				ID: "111",
			},
		})
		if err != nil {
			log.Println(err)
		}
	}()

	select {}

	/*for len(queue) > 0 {
		element := queue[0].element
		inboundFlowId := queue[0].inboundFlowId
		queue = queue[1:]
	}*/

	return nil
}
