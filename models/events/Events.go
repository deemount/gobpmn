package events

type Type_Cancel_Event_Definition CancelEventDefinitionRepository
type Type_Compensate_Event_Definition CompensateEventDefinitionRepository
type Type_Conditional_Event_Definition ConditionalEventDefinitionRepository
type Type_Error_Event_Definition ErrorEventDefinitionRepository
type Type_Escalation_Event_Definition EscalationEventDefinitionRepository
type Type_Link_Event_Definition LinkEventDefinitionRepository
type Type_Start_Event StartEventRepository
type Type_End_Event EndEventRepository
type Type_Boundary_Event BoundaryEventRepository
type Type_Signal_Event_Definition SignalEventDefinitionRepository
type Type_Timer_Event_Definition TimerEventDefinitionRepository
type Type_Terminate_Event_Definition TerminateEventDefinitionRepository
type Type_Intermediate_Catch_Event IntermediateCatchEventRepository
type Type_Intermediate_Throw_Event IntermediateThrowEventRepository

type EventsRepository[
	CED Type_Cancel_Event_Definition,
	CME Type_Compensate_Event_Definition,
	CNE Type_Conditional_Event_Definition,
	EED Type_Error_Event_Definition,
	ESCED Type_Escalation_Event_Definition,
	LED Type_Link_Event_Definition,
	S Type_Start_Event,
	E Type_End_Event,
	B Type_Boundary_Event,
	SED Type_Signal_Event_Definition,
	T Type_Timer_Event_Definition,
	TERED Type_Terminate_Event_Definition,
	IC Type_Intermediate_Catch_Event,
	IT Type_Intermediate_Throw_Event] interface {
	CancelEventDefinitionRepository() CED
	CompensateEventDefinitionRepository() CME
	ConditionalEventDefinitionRepository() CNE
	ErrorEventDefinitionRepository() EED
	EscalationEventDefinitionRepository() ESCED
	LinkEventDefinitionRepository() LED
	StartEventRepository() S
	EndEventRepository() E
	BoundaryEventRepository() B
	SignalEventDefinitionRepository() SED
	TimerEventDefinitionRepository() T
	TerminateEventDefinitionRepository() TERED
	IntermediateCatchEventRepository() IC
	IntermediateThrowEventRepository() IT
}

type Events[
	CED Type_Cancel_Event_Definition,
	CME Type_Compensate_Event_Definition,
	CNE Type_Conditional_Event_Definition,
	EED Type_Error_Event_Definition,
	ESCED Type_Escalation_Event_Definition,
	LED Type_Link_Event_Definition,
	S Type_Start_Event,
	E Type_End_Event,
	B Type_Boundary_Event,
	SED Type_Signal_Event_Definition,
	T Type_Timer_Event_Definition,
	TERED Type_Terminate_Event_Definition,
	IC Type_Intermediate_Catch_Event,
	IT Type_Intermediate_Throw_Event] struct {
	EventsRepository[CED, CME, CNE, EED, ESCED, LED, S, E, B, SED, T, TERED, IC, IT]
}

func NewEvents() EventsRepository[
	Type_Cancel_Event_Definition,
	Type_Compensate_Event_Definition,
	Type_Conditional_Event_Definition,
	Type_Error_Event_Definition,
	Type_Escalation_Event_Definition,
	Type_Link_Event_Definition,
	Type_Start_Event,
	Type_End_Event,
	Type_Boundary_Event,
	Type_Signal_Event_Definition,
	Type_Timer_Event_Definition,
	Type_Terminate_Event_Definition,
	Type_Intermediate_Catch_Event,
	Type_Intermediate_Throw_Event] {
	return &Events[
		Type_Cancel_Event_Definition,
		Type_Compensate_Event_Definition,
		Type_Conditional_Event_Definition,
		Type_Error_Event_Definition,
		Type_Escalation_Event_Definition,
		Type_Link_Event_Definition,
		Type_Start_Event,
		Type_End_Event,
		Type_Boundary_Event,
		Type_Signal_Event_Definition,
		Type_Timer_Event_Definition,
		Type_Terminate_Event_Definition,
		Type_Intermediate_Catch_Event,
		Type_Intermediate_Throw_Event]{}
}
