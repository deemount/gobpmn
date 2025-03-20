package common

// go:build 1.23

import (
	"fmt"
	"reflect"
)

// StartEventHandler handles start events
type StartEventHandler struct {
	*BaseElement
}

// StartEventIndices holds indices for start events
type StartEventIndices struct {
	startEvent int
}

// NewStartEventHandler creates a new StartEventHandler instance
func NewStartEventHandler(ep *ElementProcessor) *StartEventHandler {
	return &StartEventHandler{
		BaseElement: &BaseElement{
			processor: ep,
		},
	}
}

// Handle processes start events.
func (h *StartEventHandler) Handle(processIdx int, info FieldInfo, config *ProcessingConfig) error {
	indices := &StartEventIndices{
		startEvent: config.indices[startEvent],
	}

	if err := h.handleStartEvent(processIdx, info, config, &indices.startEvent); err != nil {
		return NewError(fmt.Errorf("failed to handle start event: %w", err))
	}

	config.indices[startEvent] = indices.startEvent

	return nil
}

// handleStartEvent handles start events.
func (h *StartEventHandler) handleStartEvent(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	// if it's NOT the first start event in the process,
	// then a start event get the "event"-prefix for the id
	if processIdx > 0 && int(*idx) == 0 {
		info.typ = "event"
	}

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		"GetStartEvent",
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get start event: %w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set start event properties: %w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure start event flow: %w", err))
	}

	(*idx)++

	return nil

}

// SetProperties sets basic properties and additional specific properties for start events.
func (h *StartEventHandler) SetProperties(element reflect.Value, info FieldInfo) error {

	// first call base implementation
	if err := h.BaseElement.SetProperties(element, info); err != nil {
		return err
	}

	// start event specific properties
	return h.setStartEventSpecificProperties(element)

}

// setStartEventSpecificProperties sets start event specific properties.
func (h *StartEventHandler) setStartEventSpecificProperties(el reflect.Value) error {
	properties := map[string]interface{}{
		// TODO: IsInterrupting needs a more generic handling. It's not always true.
		"IsInterrupting": true,
	}

	for propName, propValue := range properties {
		if err := callMethod(el, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s: %w", propName, err))
		}
	}

	return nil

}

// configureFlow configures the flow for a start event
// Note: A start event always has exactly one outgoing flow and no incoming flow
func (h *StartEventHandler) configureFlow(el reflect.Value, info FieldInfo) error {
	if err := h.setOutgoingFlowCount(el); err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	if err := h.configureOutgoingFlow(el, info.nextHash); err != nil {
		return NewError(fmt.Errorf("failed to configure outgoing flow:\n%w", err))
	}

	return nil

}

// setOutgoingFlowCount sets the number of outgoing flows for a start event
// Note: An outgoing can only be set, if there's a following known element
// Actually, there is no check if the next element is a known element.
func (h *StartEventHandler) setOutgoingFlowCount(el reflect.Value) error {
	const outgoingFlowCount = 1 // start events always have exactly one outgoing flow

	err := callMethod(el, "SetOutgoing", []reflect.Value{
		reflect.ValueOf(outgoingFlowCount),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	return nil

}

// configureOutgoingFlow configures the outgoing flow of a start event
func (h *StartEventHandler) configureOutgoingFlow(el reflect.Value, nextHash string) error {
	const firstFlowIndex = 0

	// Get the outgoing flow object
	outFlow, err := callMethodValue(el, "GetOutgoing", []reflect.Value{
		reflect.ValueOf(firstFlowIndex),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to get outgoing flow:\n%w", err))
	}

	// Set the flow hash
	err = callMethod(outFlow, "SetFlow", []reflect.Value{
		reflect.ValueOf(nextHash),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set flow hash:\n%w", err))
	}

	return nil

}

// ---------------------------------------------------------- //

// EventIndices holds indices for different types of events
type EventIndices struct {
	catch int
	throw int
	end   int
}

// EventHandler handles different types of events
type EventHandler struct {
	*BaseElement
}

// NewEventHandler creates a new EventHandler instance
func NewEventHandler(ep *ElementProcessor) *EventHandler {
	return &EventHandler{
		BaseElement: &BaseElement{
			processor: ep,
		},
	}
}

// Handle processes different types of events
func (h *EventHandler) Handle(processIdx int, info FieldInfo, config *ProcessingConfig) error {
	indices := EventIndices{
		catch: config.indices[intermediateCatchEvent],
		throw: config.indices[intermediateThrowEvent],
		end:   config.indices[endEvent],
	}

	var err error

	switch info.element {
	case intermediateCatchEvent:
		err = h.handleIntermediateCatchEvent(processIdx, info, config, &indices.catch)
	case intermediateThrowEvent:
		err = h.handleIntermediateThrowEvent(processIdx, info, config, &indices.throw)
	case endEvent:
		err = h.handleEndEvent(processIdx, info, config, &indices.end)
	default:
		return NewError(fmt.Errorf("unsupported event type: %s", info.element))
	}

	if err != nil {
		return NewError(fmt.Errorf("failed to handle event %w", err))
	}

	config.indices[intermediateCatchEvent] = indices.catch
	config.indices[intermediateThrowEvent] = indices.throw
	config.indices[endEvent] = indices.end

	return nil

}

// handleIntermediateCatchEvent ...
func (h *EventHandler) handleIntermediateCatchEvent(processIndex int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIndex],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)})
	if err != nil {
		return NewError(fmt.Errorf("failed to get intermediate catch event %v", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set intermediate catch event properties: %w", err))
	}

	(*idx)++

	return nil

}

// handleIntermediateThrowEvent processes intermediate throw events.
// It calls the GetIntermediateThrowEvent method on the process and sets
// the properties for the intermediate throw event.
func (h *EventHandler) handleIntermediateThrowEvent(processIndex int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIndex],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)})
	if err != nil {
		return NewError(fmt.Errorf("failed to get intermediate throw event: %v", err))
	}

	// Set the element properties
	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set intermediate throw event properties: %w", err))
	}

	(*idx)++

	return nil

}

// handleEndEvent processes end events.
// It calls the GetEndEvent method on the process and sets the properties for the end event.
func (h *EventHandler) handleEndEvent(processIndex int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIndex],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)})
	if err != nil {
		return NewError(fmt.Errorf("failed to get end event: %v\n", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set end event properties:\n%w", err))
	}

	(*idx)++

	return nil

}

// SetProperties sets properties for events.
func (h *EventHandler) SetProperties(element reflect.Value, info FieldInfo) error {

	// first call base implementation
	if err := h.BaseElement.SetProperties(element, info); err != nil {
		return err
	}

	// event specific properties based on type
	switch info.element {
	case intermediateCatchEvent:
		return h.setIntermediateCatchEventProperties(element)
	case intermediateThrowEvent:
		return h.setIntermediateThrowEventProperties(element)
	case endEvent:
		return h.setEndEventProperties(element)
	default:
		return NewError(fmt.Errorf("unsupported event type: %s", info.element))
	}

}

// setIntermediateCatchEventProperties sets start event specific properties.
func (h *EventHandler) setIntermediateCatchEventProperties(el reflect.Value) error {
	properties := map[string]interface{}{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(el, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// setIntermediateThrowEventProperties sets start event specific properties.
func (h *EventHandler) setIntermediateThrowEventProperties(element reflect.Value) error {
	properties := map[string]interface{}{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// setEndEventProperties sets start event specific properties.
func (h *EventHandler) setEndEventProperties(element reflect.Value) error {
	properties := map[string]interface{}{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// configureFlow configures the flow for a event
func (h *EventHandler) configureFlow(el reflect.Value, info FieldInfo) error {
	if err := h.setIncomingFlowCount(el); err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	if err := h.configureIncomingFlow(el, info.hashBefore); err != nil {
		return NewError(fmt.Errorf("failed to configure outgoing flow:\n%w", err))
	}

	if err := h.setOutgoingFlowCount(el); err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	if err := h.configureOutgoingFlow(el, info.nextHash); err != nil {
		return NewError(fmt.Errorf("failed to configure outgoing flow:\n%w", err))
	}

	return nil

}

// setIncomingFlowCount sets the number of incoming flows for a event type
func (h *EventHandler) setIncomingFlowCount(element reflect.Value) error {
	const incomingFlowCount = 1 // start events always have exactly one outgoing flow

	err := callMethod(element, "SetIncoming", []reflect.Value{
		reflect.ValueOf(incomingFlowCount),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set incoming flow count:\n%w", err))
	}

	return nil

}

// configureOutgoingFlow configures the outgoing flow of a start event
func (h *EventHandler) configureIncomingFlow(element reflect.Value, hashBefore string) error {
	const firstFlowIndex = 0

	// Get the outgoing flow object
	inFlow, err := callMethodValue(element, "GetIncoming", []reflect.Value{
		reflect.ValueOf(firstFlowIndex),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to get outgoing flow:\n%w", err))
	}

	// Set the flow hash
	err = callMethod(inFlow, "SetFlow", []reflect.Value{
		reflect.ValueOf(hashBefore),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set flow hash:\n%w", err))
	}

	return nil

}

// setOutgoingFlowCount sets the number of outgoing flows for a event
// Note: An outgoing can only be set, if there's a following known element
// Actually, there is no check if the next element is a known element.
func (h *EventHandler) setOutgoingFlowCount(el reflect.Value) error {
	const outgoingFlowCount = 1 // start events always have exactly one outgoing flow

	err := callMethod(el, "SetOutgoing", []reflect.Value{
		reflect.ValueOf(outgoingFlowCount),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	return nil

}

// configureOutgoingFlow configures the outgoing flow of a start event
func (h *EventHandler) configureOutgoingFlow(el reflect.Value, nextHash string) error {
	flowIndex := 0

	// Get the outgoing flow object
	outFlow, err := callMethodValue(el, "GetOutgoing", []reflect.Value{
		reflect.ValueOf(flowIndex),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to get outgoing flow:\n%w", err))
	}

	// Set the flow hash
	err = callMethod(outFlow, "SetFlow", []reflect.Value{
		reflect.ValueOf(nextHash),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set flow hash:\n%w", err))
	}

	return nil

}

// ---------------------------------------------------------- //

// GatewayIndices holds indices for different types of gateways.
type GatewayIndices struct {
	inclusive int
	exclusive int
	parallel  int
	gateway   int
}

// GatewayHandler handles different types of gateways
type GatewayHandler struct {
	*BaseElement
}

// NewGatewayHandler creates a new GatewayHandler instance
func NewGatewayHandler(ep *ElementProcessor) *GatewayHandler {
	return &GatewayHandler{
		BaseElement: &BaseElement{
			processor: ep,
		},
	}
}

// handleGateway ...
func (h *GatewayHandler) Handle(processIndex int, info FieldInfo, config *ProcessingConfig) error {
	indices := GatewayIndices{
		inclusive: config.indices[inclusiveGateway],
		exclusive: config.indices[exclusiveGateway],
		parallel:  config.indices[parallelGateway],
		gateway:   config.indices[gateway],
	}

	var err error

	switch info.element {
	case inclusiveGateway:
		err = h.handleInclusiveGateway(processIndex, info, config, &indices.inclusive)
	case exclusiveGateway:
		err = h.handleExclusiveGateway(processIndex, info, config, &indices.exclusive)
	case parallelGateway:
		err = h.handleParallelGateway(processIndex, info, config, &indices.parallel)
	case gateway:
		err = h.handleGateway(processIndex, info, config, &indices.gateway)
	default:
		return NewError(fmt.Errorf("unsupported gateway type: %s\n", info.element))
	}

	if err != nil {
		return NewError(fmt.Errorf("failed to handle gateway:\n%w", err))
	}

	config.indices[inclusiveGateway] = indices.inclusive
	config.indices[exclusiveGateway] = indices.exclusive
	config.indices[parallelGateway] = indices.parallel
	config.indices[gateway] = indices.gateway

	return nil

}

// handleInclusiveGateway processes inclusive gateway elements
func (h *GatewayHandler) handleInclusiveGateway(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get inclusive gateway:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set inclusive gateway properties:\n%w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure inclusive gateway flow:\n%w", err))
	}

	(*idx)++

	return nil

}

// handleExclusiveGateway processes exclusive gateway elements
func (h *GatewayHandler) handleExclusiveGateway(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get exclusive gateway:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set exclusive gateway properties:\n%w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure exclusive event flow:\n%w", err))
	}

	(*idx)++

	return nil

}

// handleParallelGateway processes parallel gateway elements.
// It calls the GetParallelGateway method on the process and sets
// the properties for the gateway.
func (h *GatewayHandler) handleParallelGateway(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get parallel gateway:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set parallel gateway properties:\n%w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure parallel gateway flow:\n%w", err))
	}

	(*idx)++

	return nil

}

// handleGateway processes gateway elements.
// It calls the GetGateway method on the process and sets the properties for the gateway.
func (h *GatewayHandler) handleGateway(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get gateway:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set gateway properties:\n%w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure gateway flow:\n%w", err))
	}

	(*idx)++

	return nil

}

// SetProperties sets properties for gateways.
func (h *GatewayHandler) SetProperties(element reflect.Value, info FieldInfo) error {

	// first call base implementation
	if err := h.BaseElement.SetProperties(element, info); err != nil {
		return err
	}

	// gateway specific properties based on type
	switch info.element {
	case inclusiveGateway:
		return h.setInclusiveGatewayProperties(element)
	case exclusiveGateway:
		return h.setExclusiveGatewayProperties(element)
	case parallelGateway:
		return h.setParallelGatewayProperties(element)
	case gateway:
		return h.setGatewayProperties(element)
	default:
		return NewError(fmt.Errorf("unsupported gateway type: %s\n", info.element))
	}

}

// setInclusiveGatewayProperties sets inclusive gateway specific properties.
func (h *GatewayHandler) setInclusiveGatewayProperties(element reflect.Value) error {
	properties := map[string]any{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// setExclusiveGatewayProperties sets exclusive gateway specific properties.
func (h *GatewayHandler) setExclusiveGatewayProperties(element reflect.Value) error {
	properties := map[string]any{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// setParallelGatewayProperties sets parallel gateway specific properties.
func (h *GatewayHandler) setParallelGatewayProperties(element reflect.Value) error {
	properties := map[string]any{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// setGatewayProperties sets gateway specific properties.
func (h *GatewayHandler) setGatewayProperties(element reflect.Value) error {
	properties := map[string]any{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// configureFlow configures the flow for a event
func (h *GatewayHandler) configureFlow(el reflect.Value, info FieldInfo) error {
	if err := h.setIncomingFlowCount(el); err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	if err := h.configureIncomingFlow(el, info.hashBefore); err != nil {
		return NewError(fmt.Errorf("failed to configure outgoing flow:\n%w", err))
	}

	if err := h.setOutgoingFlowCount(el); err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	if err := h.configureOutgoingFlow(el, info.nextHash); err != nil {
		return NewError(fmt.Errorf("failed to configure outgoing flow:\n%w", err))
	}

	return nil

}

// setIncomingFlowCount sets the number of incoming flows for a gateway
func (h *GatewayHandler) setIncomingFlowCount(el reflect.Value) error {
	incomingFlowCount := 1 // TODO: the incoming flow count is yet not to be determined

	err := callMethod(el, "SetIncoming", []reflect.Value{
		reflect.ValueOf(incomingFlowCount),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set incoming flow count:\n%w", err))
	}

	return nil

}

// configureOutgoingFlow configures the outgoing flow of a gateway
func (h *GatewayHandler) configureIncomingFlow(el reflect.Value, hashBefore string) error {
	flowIndex := 0 // TODO: the flow index is yet not to be determined

	// Get the outgoing flow object
	inFlow, err := callMethodValue(el, "GetIncoming", []reflect.Value{
		reflect.ValueOf(flowIndex),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to get incoming flow:\n%w", err))
	}

	// Set the flow hash
	err = callMethod(inFlow, "SetFlow", []reflect.Value{
		reflect.ValueOf(hashBefore),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set flow hash:\n%w", err))
	}

	return nil

}

// setOutgoingFlowCount sets the number of outgoing flows for a gateway
// Note: An outgoing can only be set, if there's a following known element
// Actually, there is no check if the next element is a known element.
func (h *GatewayHandler) setOutgoingFlowCount(el reflect.Value) error {
	const outgoingFlowCount = 1 // TODO: the outgoing flow count is yet not to be determined

	err := callMethod(el, "SetOutgoing", []reflect.Value{
		reflect.ValueOf(outgoingFlowCount),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	return nil

}

// configureOutgoingFlow configures the outgoing flow of a start event
func (h *GatewayHandler) configureOutgoingFlow(el reflect.Value, nextHash string) error {
	flowIndex := 0 // TODO: the flow index is yet not to be determined

	// Get the outgoing flow object
	outFlow, err := callMethodValue(el, "GetOutgoing", []reflect.Value{
		reflect.ValueOf(flowIndex),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to get outgoing flow:\n%w", err))
	}

	// Set the flow hash
	err = callMethod(outFlow, "SetFlow", []reflect.Value{
		reflect.ValueOf(nextHash),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set flow hash:\n%w", err))
	}

	return nil

}

// ---------------------------------------------------------- //

// ActivityIndices holds indices for different types of activities
type ActivityIndices struct {
	userTask    int
	serviceTask int
	scriptTask  int
	task        int
}

// ActivityHandler handles different types of activities
type ActivityHandler struct {
	*BaseElement
}

// NewActivityHandler creates a new ActivityHandler instance
func NewActivityHandler(ep *ElementProcessor) *ActivityHandler {
	return &ActivityHandler{
		BaseElement: &BaseElement{
			processor: ep,
		},
	}
}

// Handle processes different types of activities.
// It uses indices to keep track of the number of activities processed.
func (h *ActivityHandler) Handle(processIdx int, info FieldInfo, config *ProcessingConfig) error {

	// Initialize indices
	indices := ActivityIndices{
		userTask:    config.indices[userTask],
		serviceTask: config.indices[serviceTask],
		scriptTask:  config.indices[scriptTask],
		task:        config.indices[task],
	}

	var err error

	switch info.element {
	case userTask:
		err = h.handleUserTask(processIdx, info, config, &indices.userTask)
	case serviceTask:
		err = h.handleServiceTask(processIdx, info, config, &indices.serviceTask)
	case scriptTask:
		err = h.handleScriptTask(processIdx, info, config, &indices.scriptTask)
	case task:
		err = h.handleTask(processIdx, info, config, &indices.task)
	default:
		return NewError(fmt.Errorf("unsupported activity type: %s\n", info.element))
	}

	if err != nil {
		return NewError(fmt.Errorf("failed to handle activity:\n%w", err))
	}

	config.indices[task] = indices.task
	config.indices[userTask] = indices.userTask
	config.indices[serviceTask] = indices.serviceTask
	config.indices[scriptTask] = indices.scriptTask

	return nil

}

// handleTask processes task elements.
// It calls the GetTask method on the process and sets the properties for the task.
func (h *ActivityHandler) handleTask(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get task:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set task properties:\n%w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure task flow:\n%w", err))
	}

	(*idx)++

	return nil

}

// handleUserTask processes user task elements.
// It calls the GetUserTask method on the process and sets the properties for the user task
func (h *ActivityHandler) handleUserTask(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get user task:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set user task properties:\n%w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure user task flow:\n%w", err))
	}

	(*idx)++

	return nil

}

// handleScriptTask calls the GetScriptTask method on the process
// and sets the properties for the script task.
func (h *ActivityHandler) handleScriptTask(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get script task:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set script task properties:\n%w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure script task flow:\n%w", err))
	}

	(*idx)++

	return nil

}

// handleServiceTask processes service task elements.
// It calls the GetServiceTask method on the process and sets the properties for the service task.
func (h *ActivityHandler) handleServiceTask(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get service task:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set service task properties:\n%w", err))
	}

	if err := h.configureFlow(el, info); err != nil {
		return NewError(fmt.Errorf("failed to configure service task flow:\n%w", err))
	}

	(*idx)++

	return nil

}

// SetProperties sets properties for activities.
// It calls the SetProperties method in the BaseElement to set common properties, like ID and Name.
func (h *ActivityHandler) SetProperties(element reflect.Value, info FieldInfo) error {

	// first call base implementation
	if err := h.BaseElement.SetProperties(element, info); err != nil {
		return err
	}

	// activity specific properties based on type
	switch info.element {
	case userTask:
		return h.setUserTaskProperties(element)
	case serviceTask:
		return h.setServiceTaskProperties(element)
	case scriptTask:
		return h.setScriptTaskProperties(element)
	case task:
		return h.setTaskProperties(element)
	default:
		return NewError(fmt.Errorf("unsupported activity type: %s\n", info.element))
	}

}

// setUserTaskProperties sets specific properties.
func (h *ActivityHandler) setUserTaskProperties(element reflect.Value) error {
	properties := map[string]any{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// setServiceTaskProperties sets specific properties.
func (h *ActivityHandler) setServiceTaskProperties(element reflect.Value) error {
	properties := map[string]any{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// setScriptTaskProperties sets specific properties.
func (h *ActivityHandler) setScriptTaskProperties(element reflect.Value) error {
	properties := map[string]any{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// setTaskProperties sets specific properties.
func (h *ActivityHandler) setTaskProperties(element reflect.Value) error {
	properties := map[string]any{
		// Add more properties as needed
	}

	for propName, propValue := range properties {
		if err := callMethod(element, "Set"+propName, []reflect.Value{
			reflect.ValueOf(propValue),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set %s:\n%w", propName, err))
		}
	}

	return nil

}

// configureFlow configures the flow for a activity
func (h *ActivityHandler) configureFlow(el reflect.Value, info FieldInfo) error {
	if err := h.setIncomingFlowCount(el); err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	if err := h.configureIncomingFlow(el, info.hashBefore); err != nil {
		return NewError(fmt.Errorf("failed to configure outgoing flow:\n%w", err))
	}

	if err := h.setOutgoingFlowCount(el); err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	if err := h.configureOutgoingFlow(el, info.nextHash); err != nil {
		return NewError(fmt.Errorf("failed to configure outgoing flow:\n%w", err))
	}

	return nil

}

// setIncomingFlowCount sets the number of incoming flows for a event type
func (h *ActivityHandler) setIncomingFlowCount(el reflect.Value) error {
	incomingFlowCount := 1 // TODO: the incoming flow count is yet not to be determined

	err := callMethod(el, "SetIncoming", []reflect.Value{
		reflect.ValueOf(incomingFlowCount),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set incoming flow count:\n%w", err))
	}

	return nil

}

// configureOutgoingFlow configures the outgoing flow of a activity
func (h *ActivityHandler) configureIncomingFlow(el reflect.Value, hashBefore string) error {
	flowIndex := 0 // TODO: the flow index is yet not to be determined

	// Get the outgoing flow object
	inFlow, err := callMethodValue(el, "GetIncoming", []reflect.Value{
		reflect.ValueOf(flowIndex),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to get incoming flow:\n%w", err))
	}

	// Set the flow hash
	err = callMethod(inFlow, "SetFlow", []reflect.Value{
		reflect.ValueOf(hashBefore),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set flow hash:\n%w", err))
	}

	return nil

}

// setOutgoingFlowCount sets the number of outgoing flows for a activity
// Note: An outgoing can only be set, if there's a following known element
// Actually, there is no check if the next element is a known element.
func (h *ActivityHandler) setOutgoingFlowCount(el reflect.Value) error {
	const outgoingFlowCount = 1 // TODO: the outgoing flow count is yet not to be determined

	err := callMethod(el, "SetOutgoing", []reflect.Value{
		reflect.ValueOf(outgoingFlowCount),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set outgoing flow count:\n%w", err))
	}

	return nil

}

// configureOutgoingFlow configures the outgoing flow of a activity
func (h *ActivityHandler) configureOutgoingFlow(el reflect.Value, nextHash string) error {
	flowIndex := 0 // TODO: the flow index is yet not to be determined

	// Get the outgoing flow object
	outFlow, err := callMethodValue(el, "GetOutgoing", []reflect.Value{
		reflect.ValueOf(flowIndex),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to get outgoing flow:\n%w", err))
	}

	// Set the flow hash
	err = callMethod(outFlow, "SetFlow", []reflect.Value{
		reflect.ValueOf(nextHash),
	})
	if err != nil {
		return NewError(fmt.Errorf("failed to set flow hash:\n%w", err))
	}

	return nil

}

// ---------------------------------------------------------- //

// FlowIndices holds indices for flow elements.
type FlowIndices struct {
	flow int
}

// FlowHandler handles flow elements
type FlowHandler struct {
	*BaseElement
}

// NewFlowHandler creates a new FlowHandler instance
func NewFlowHandler(ep *ElementProcessor) *FlowHandler {
	handler := &FlowHandler{
		BaseElement: &BaseElement{
			processor: ep,
		},
	}
	return handler
}

// Handle processes a flow element
func (h *FlowHandler) Handle(processIdx int, info FieldInfo, config *ProcessingConfig) error {

	indices := FlowIndices{
		flow: config.indices[sequenceFlow],
	}

	var err error

	switch info.element {
	case sequenceFlow:
		err = h.handleFlow(processIdx, info, config, &indices.flow)
	default:
		return NewError(fmt.Errorf("unsupported flow type: %s\n", info.element))
	}

	if err != nil {
		return NewError(fmt.Errorf("failed to handle flow:\n%w", err))
	}

	config.indices[sequenceFlow] = indices.flow

	return nil

}

// handleFlow processes flow elements.
func (h *FlowHandler) handleFlow(processIdx int, info FieldInfo, config *ProcessingConfig, idx *int) error {
	if *idx >= config.counts[info.element] {
		return nil
	}

	methodName := fmt.Sprintf("Get%s", info.element)

	el, err := callMethodValue(
		h.processor.value.Process[processIdx],
		methodName,
		[]reflect.Value{reflect.ValueOf(*idx)},
	)
	if err != nil {
		return NewError(fmt.Errorf("failed to get flow:\n%w", err))
	}

	if err := h.SetProperties(el, info); err != nil {
		return NewError(fmt.Errorf("failed to set flow properties:\n%w", err))
	}

	(*idx)++

	return nil

}

// SetProperties sets properties for flow elements
func (h *FlowHandler) SetProperties(flow reflect.Value, info FieldInfo) error {

	// call base implementation
	if err := h.BaseElement.SetProperties(flow, info); err != nil {
		return err
	}

	// Set flow-specific properties
	if err := h.setFlowSpecificProperties(flow, info); err != nil {
		return err
	}
	return nil
}

// setFlowSpecificProperties sets properties specific to sequence flows
func (h *FlowHandler) setFlowSpecificProperties(flow reflect.Value, info FieldInfo) error {

	// Set source and target references if available
	sourceRef := h.processor.value.FlowNeighbors[info.name]["SourceRef"].(map[string]any)
	targetRef := h.processor.value.FlowNeighbors[info.name]["TargetRef"].(map[string]any)

	if info.hash != "" {
		if err := callMethod(flow, "SetSourceRef", []reflect.Value{
			reflect.ValueOf(sourceRef["Value"].(BPMN).Type),
			reflect.ValueOf(sourceRef["Value"].(BPMN).Hash),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set source reference:\n %w", err))
		}
	}

	if info.nextHash != "" {
		if err := callMethod(flow, "SetTargetRef", []reflect.Value{
			reflect.ValueOf(targetRef["Value"].(BPMN).Type),
			reflect.ValueOf(targetRef["Value"].(BPMN).Hash),
		}); err != nil {
			return NewError(fmt.Errorf("failed to set target reference:\n %w", err))
		}
	}

	return nil

}
