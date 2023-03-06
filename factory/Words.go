package factory

// NOTICE:
// It's a first build of a library for usage in natural language processing
// They may be used in some example processes, but as a whole it could be,
// that the whole library is getting deleted or put in an other structure.

var (

	// type
	typeBuilder = "Builder"

	// anonymous classified (plural)
	fieldEventDefinitions = "EventDefinitions"
	fieldEventElements    = "EventElements"
	fieldEvents           = "Events"
	fieldGateways         = "Gateways"
	fieldMessages         = "Messages"
	fieldPool             = "Pool"
	fieldProcesses        = "Processes"
	fieldTasks            = "Tasks"

	// field
	fieldCollaboration = "Collaboration"
	fieldParticipant   = "Participant"
	fieldID            = "ID"
	fieldProcess       = "Process"

	// field event elements
	fieldBoundaryEvent          = "BoundaryEvent"
	fieldEndEvent               = "EndEvent"
	fieldIntermediateCatchEvent = "IntermediateCatchEvent"
	fieldIntermediateThrowEvent = "IntermediateThrowEvent"
	fieldMessage                = "Message"
	fieldSignal                 = "Signal"
	fieldStartEvent             = "StartEvent"

	// field event definitions
	fieldCancelEventDefinition     = "CancelEventDefinition"
	fieldCompensateEventDefinition = "CompensateEventDefinition"
	fieldErrorEventDefinition      = "ErrorEventDefinition"
	fieldEscalationEventDefinition = "EscalationEventDefinition"
	fieldLinkEventDefinition       = "LinkEventDefinition"
	fieldMessageEventDefinition    = "MessageEventDefinition"
	fieldSignalEventDefinition     = "SignalEventDefinition"
	fieldTerminateEventDefinition  = "TerminateEventDefinition"
	fieldTimerEventDefinition      = "TimerEventDefinition"

	// field gateways
	fieldComplexGateway    = "ComplexGateway"
	fieldEventBasedGateway = "EventBasedGateway"
	fieldExclusiveGateway  = "ExclusiveGateway"
	fieldInclusiveGateway  = "InclusiveGateway"
	fieldParallelGateway   = "ParallelGateway"

	// tasks
	fieldBusinessRuleTask = "BusinessRuleTask"
	fieldManualTask       = "ManualTask"
	fieldReceiveTask      = "ReceiveTask"
	fieldScriptTask       = "ScriptTask"
	fieldServiceTask      = "ServiceTask"
	fieldTask             = "Task"
	fieldUserTask         = "UserTask"

	// unknown
	fieldUnknown = "Unknown"

	// bool
	boolIsExecutable = "IsExecutable"

	// methods (actually used in Builder.build only)
	MethodSetCollaboration = "SetCollaboration"
	MethodSetProcess       = "SetProcess"
	MethodSetDiagram       = "SetDiagram"

	// map
	words = map[string]map[int]string{
		"type": {
			0: typeBuilder,
		},
		"anonymous": {
			0: fieldEventDefinitions,
			1: fieldEventElements,
			2: fieldEvents,
			3: fieldMessages,
			4: fieldPool,
			5: fieldProcesses,
			6: fieldTasks,
		},
		"pool": {
			0: fieldCollaboration,
			1: fieldParticipant,
			2: fieldID,
			3: fieldProcess,
		},
		"event_elements": {
			0: fieldBoundaryEvent,
			1: fieldEndEvent,
			2: fieldIntermediateCatchEvent,
			3: fieldIntermediateThrowEvent,
			4: fieldMessage,
			5: fieldSignal,
			6: fieldStartEvent,
		},
		"event_definitions": {
			0: fieldCancelEventDefinition,
			1: fieldCompensateEventDefinition,
			2: fieldErrorEventDefinition,
			3: fieldEscalationEventDefinition,
			4: fieldLinkEventDefinition,
			5: fieldMessageEventDefinition,
			6: fieldSignalEventDefinition,
			7: fieldTerminateEventDefinition,
			8: fieldTimerEventDefinition,
		},
		"tasks": {
			0: fieldBusinessRuleTask,
			1: fieldManualTask,
			2: fieldReceiveTask,
			3: fieldScriptTask,
			4: fieldServiceTask,
			5: fieldTask,
			6: fieldUserTask,
		},
		"unknown": {
			0: fieldUnknown,
		},
		"configuration": {
			0: boolIsExecutable,
		},
	}
)
