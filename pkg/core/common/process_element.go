package common

// processElement represents a BPMN process element type
type processElement string

func (pe processElement) String() string {
	return string(pe)
}

const (
	startEvent             processElement = "StartEvent"
	endEvent               processElement = "EndEvent"
	intermediateCatchEvent processElement = "IntermediateCatchEvent"
	intermediateThrowEvent processElement = "IntermediateThrowEvent"
	inclusiveGateway       processElement = "InclusiveGateway"
	exclusiveGateway       processElement = "ExclusiveGateway"
	parallelGateway        processElement = "ParallelGateway"
	gateway                processElement = "Gateway"
	userTask               processElement = "UserTask"
	serviceTask            processElement = "ServiceTask"
	scriptTask             processElement = "ScriptTask"
	task                   processElement = "Task"
	sequenceFlow           processElement = "SequenceFlow"
)
