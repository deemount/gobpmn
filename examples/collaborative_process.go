package examples

<<<<<<< HEAD
import (
	"github.com/deemount/gobpmn/factory"
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/collaboration"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
=======
//PoolCustomerSupport: canvas.Bounds{X: 150, Y: 80, Width: 800, Height: 160}
//PoolCustomer: canvas.Bounds{X: 150, Y: 360, Width: 800, Height: 160}
//CustomerSupportStartEvent: canvas.Bounds{X: 225, Y: 142, Width: 36, Height: 36}
//CustomerStartEvent: canvas.Bounds{X: 225, Y: 422, Width: 36, Height: 36}
//CustomerSupportEndEvent: canvas.Bounds{X: 822, Y: 142, Width: 36, Height: 36}
//CustomerEndEvent: canvas.Bounds{X: 822, Y: 422, Width: 36, Height: 36}

/**************************************************************************************/
/**
 * @Import
 *
 * Import models and utils package to build the model
 **/

import (
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/definitions"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/pool"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
	"github.com/deemount/gobpmn/utils"
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
)

// CollaborativeProcess ...
type CollaborativeProcess interface {
	Create() CProcess
}

// CProcess ...
type CProcess struct {
	Def core.DefinitionsRepository
	Pool
	Message
	CustomerSupport
	Customer
}

// Pool ...
type Pool struct {
	CustomerSupportIsExecutable bool
	CustomerIsExecutable        bool
	Collaboration               factory.BpmnBuilder
	CustomerSupportID           factory.BpmnBuilder
	CustomerSupportProcess      factory.BpmnBuilder
	CustomerID                  factory.BpmnBuilder
	CustomerProcess             factory.BpmnBuilder
}

// Message ...
type Message struct {
	CustomerToCustomerSupportMessage factory.BpmnBuilder
	CustomerSupportToCustomerMessage factory.BpmnBuilder
}

// CustomerSupportProcess ...
type CustomerSupport struct {
	CustomerSupportStartEvent     factory.BpmnBuilder
	FromCustomerSupportStartEvent factory.BpmnBuilder
	CheckIncomingClaim            factory.BpmnBuilder
	FromCheckIncomingClaim        factory.BpmnBuilder
	DenyWarrantyClaim             factory.BpmnBuilder
	FromDenyWarrantyClaim         factory.BpmnBuilder
	CustomerSupportEndEvent       factory.BpmnBuilder
}

// Customer ...
type Customer struct {
	CustomerStartEvent                   factory.BpmnBuilder
	FromCustomerStartEvent               factory.BpmnBuilder
	NoticeOfDefect                       factory.BpmnBuilder
	FromNoticeOfDefect                   factory.BpmnBuilder
	WaitingForAnswer                     factory.BpmnBuilder
	TimerEventDefinitionWaitingForAnswer factory.BpmnBuilder
	FromWaitingForAnswer                 factory.BpmnBuilder
	ReceiptWarrantyRefusal               factory.BpmnBuilder
	FromReceiptWarrantyRefusal           factory.BpmnBuilder
	CustomerEndEvent                     factory.BpmnBuilder
}

// NewCollaborativeProcess refers to the definitions struct to start building the model
// e.g. def: new(models.Definitions) use factory function, if zero value is insufficent (allocates memory)
func NewCollaborativeProcess() CollaborativeProcess {
<<<<<<< HEAD
	cp := bb.Inject(CProcess{}).(CProcess)
	cp.Def = core.NewDefinitions(cp)
	return &cp
}

=======
	return &collaborativeProcess{
		// you need to refer to the definitions struct to start building the model
		// Notice:
		// Also possible to use (look below), but allocates memory
		// def: new(models.Definitions), // use factory function, if zero value is insufficent
		def: core.NewDefinitions(),
		// Pool
		collaborativeProcessPool: collaborativeProcessPool{
			CustomerSupportIsExecutable: true,
			CustomerIsExecutable:        false,
			CollaborationID:             "uniqueCollaborationId",
			CustomerSupportID:           "uniqueCustomerSupportId",
			CustomerID:                  "uniqueCustomerId",
			CustomerSupportProcessID:    "uniqueCustomerSupportProcessId",
			CustomerProcessID:           "uniqueCustomerProcessId",
		},
		// Message
		collaborativeProcessMessage: collaborativeProcessMessage{
			CustomerToCustomerSupportMessageHash: utils.GenerateHash(),
			CustomerSupportToCustomerMessageHash: utils.GenerateHash(),
		},
		// Customer Support
		// @Event
		collaborativeProcessCustomerSupportEvent: collaborativeProcessCustomerSupportEvent{
			CustomerSupportStartEventHash: utils.GenerateHash(),
			CustomerSupportEndEventHash:   utils.GenerateHash(),
		},
		// @Sequence
		collaborativeProcessCustomerSupportSequence: collaborativeProcessCustomerSupportSequence{
			FromCustomerSupportStartEventHash: utils.GenerateHash(),
			FromCustomerSupportEndEventHash:   utils.GenerateHash(),
			FromCheckIncomingClaimHash:        utils.GenerateHash(),
			FromDenyWarrantyClaimHash:         utils.GenerateHash(),
		},
		// @Task
		collaborativeProcessCustomerSupportTask: collaborativeProcessCustomerSupportTask{
			CheckIncomingClaimHash: utils.GenerateHash(),
			DenyWarrantyClaimHash:  utils.GenerateHash(),
		},
		// Customer
		// @Event
		collaborativeProcessCustomerEvent: collaborativeProcessCustomerEvent{
			CustomerStartEventHash:                   utils.GenerateHash(),
			CustomerEndEventHash:                     utils.GenerateHash(),
			TimerEventDefinitionWaitingForAnswerHash: utils.GenerateHash(),
		},
		// @Sequence
		collaborativeProcessCustomerSequence: collaborativeProcessCustomerSequence{
			FromCustomerStartEventHash:     utils.GenerateHash(),
			FromCustomerEndEventHash:       utils.GenerateHash(),
			FromNoticeOfDefectHash:         utils.GenerateHash(),
			FromWaitingForAnswerHash:       utils.GenerateHash(),
			FromReceiptWarrantyRefusalHash: utils.GenerateHash(),
		},
		// @Task
		collaborativeProcessCustomerTask: collaborativeProcessCustomerTask{
			NoticeOfDefectHash:          utils.GenerateHash(),
			WaitingForAnswerHash:        utils.GenerateHash(),
			WaitingForAnswerShapeIDHash: utils.GenerateHash(),
			ReceiptWarrantyRefusalHash:  utils.GenerateHash(),
		},
	}
}

/**************************************************************************************/
/**
 * @Create
 *
 * Die Create-Methode ruft ... Methoden auf, die ...
 *
 * @@setMainElements
 * @@setInnerElements
 * @@setDefinitionsAttributes
 * @@setCollaboration
 * @@setProcessCustomerSupport
 * @@setProcessCustomer
 * @@setCustomerSupportStartEvent
 * @@setCustomerSupportEndEvent
 * @@setCheckIncomingClaim
 * @@setDenyWarrantyClaim
 * @@fromCustomerSupportStartEvent
 * @@fromCheckIncomingClaim
 * @@fromDenyWarrantyClaim
 * @@setCustomerStartEvent
 * @@setCustomerEndEvent
 * @@setNoticeOfDefect
 * @@setWaitingForAnswer
 * @@setReceiptWarrantyRefusal
 * @@fromCustomerStartEvent
 * @@fromNoticeOfDefect
 * @@fromWaitingForAnswer
 * @@fromReceiptWarrantyRefusal
 * @@setDiagram
 * @@setPlane
 * @Def
 **/

>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
// Create ...
func (cp CProcess) Create() CProcess {
	// Elements
	bb.Build(cp.Def)
	cp.setInnerElements()
	cp.setDefinitionsAttributes()
	cp.setCollaboration()
	// Customer Support
	cp.setCustomerSupportProcess()
	cp.setCustomerSupportStartEvent()
	cp.setCheckIncomingClaim()
	cp.setDenyWarrantyClaim()
	cp.setCustomerSupportEndEvent()
	// Customer
	cp.setCustomerProcess()
	cp.setCustomerStartEvent()
	cp.setNoticeOfDefect()
	cp.setWaitingForAnswer()
	cp.setReceiptWarrantyRefusal()
<<<<<<< HEAD
	cp.setCustomerEndEvent()
=======
	// Customer Sequence
	cp.fromCustomerStartEvent()
	cp.fromNoticeOfDefect()
	cp.fromWaitingForAnswer()
	cp.fromReceiptWarrantyRefusal()
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	// Canvas
	cp.setDiagram()
	cp.setPlane()
	return cp
}

// Def ...
<<<<<<< HEAD
func (cp CProcess) Build() core.DefinitionsRepository { return cp.Def }

// setInnerElements ...
func (cp *CProcess) setInnerElements() {
	// Collaboration
	collaboration := cp.collaboration()
	collaboration.SetID("collaboration", cp.Collaboration.Suffix)
	collaboration.SetParticipant(bb.NumPart)
	collaboration.SetMessageFlow(bb.NumMsg)
	// Processes
	cp.customerSupportProcess().SetStartEvent(1)
	cp.customerSupportProcess().SetEndEvent(1)
	cp.customerSupportProcess().SetTask(2)
	cp.customerSupportProcess().SetSequenceFlow(3)
	cp.customerProcess().SetStartEvent(1)
	cp.customerProcess().SetEndEvent(1)
	cp.customerProcess().SetTask(2)
	cp.customerProcess().SetIntermediateCatchEvent(1)
	cp.customerProcess().SetSequenceFlow(4)
=======
func (cp collaborativeProcess) Def() core.DefinitionsRepository {
	return cp.def
}

/**************************************************************************************/

// setInnerElements ...
//
// All elements are created in the number according to the integer. Outgoing
// are the previously created parent elements by means of the method SetMainElements,
// which in the context of its execution calls the methods SetCollaboration, SetProcess
// (with number X of processes) and SetDiagram. This makes it possible to define all further
// elements as a map.
func (cp *collaborativeProcess) setInnerElements() {
	// Collaboration
	collaboration := cp.collaboration()
	collaboration.SetID("id", cp.CollaborationID)
	collaboration.SetParticipant(2)
	collaboration.SetMessageFlow(2)

	// Processes
	cp.setInnerElementsCustomerSupport(cp.customerSupportProcess())
	cp.setInnerElementsCustomer(cp.customerProcess())

>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
	// Canvas
	diagram := cp.diagram()
	diagram.SetPlane()
	plane := cp.plane()
	plane.SetShape(11)
	plane.SetEdge(9)
}

// setDefinitionsAttributes ...
func (cp *CProcess) setDefinitionsAttributes() {
	cp.Def.SetDefaultAttributes()
}

// setCollaboration ...
<<<<<<< HEAD
func (cp *CProcess) setCollaboration() {
	cp.setParticipants()
	cp.setMessageFlows()
	cp.setPools()
}

// setParticipants ...
func (cp *CProcess) setParticipants() {
	// customer support
	collaboration.SetParticipant(
		collaboration.DelegateParameter{
			PPT: cp.collaboration().GetParticipant(0),
			T:   "participant",
			PR:  "process",
			N:   "Customer Support",
			H:   []string{cp.CustomerSupportID.Suffix, cp.CustomerSupportProcess.Suffix}})
	// customer
	collaboration.SetParticipant(
		collaboration.DelegateParameter{
			PPT: cp.collaboration().GetParticipant(1),
			T:   "participant",
			PR:  "process",
			N:   "Customer",
			H:   []string{cp.CustomerID.Suffix, cp.CustomerProcess.Suffix}})
}

// setMessageFlows ...
func (cp *CProcess) setMessageFlows() {
	flow.SetMessageFlow(
		flow.DelegateParameter{
			MF: cp.collaboration().GetMessageFlow(0),
			T:  "flow",
			N:  "claim",
			ST: "activity",
			TT: "activity",
			H:  []string{cp.CustomerToCustomerSupportMessage.Suffix, cp.NoticeOfDefect.Suffix, cp.CheckIncomingClaim.Suffix}})
	flow.SetMessageFlow(
		flow.DelegateParameter{
			MF: cp.collaboration().GetMessageFlow(1),
			T:  "flow",
			N:  "refusal",
			ST: "activity",
			TT: "activity",
			H:  []string{cp.CustomerSupportToCustomerMessage.Suffix, cp.DenyWarrantyClaim.Suffix, cp.ReceiptWarrantyRefusal.Suffix}})
	cp.setMessageRefusal()
}

// setEdgeMessageRefusal ...
func (cp *CProcess) setMessageRefusal() {
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: cp.plane().GetEdge(8),
=======
func (cp *collaborativeProcess) setCollaboration() {
	cp.setParticipants()
	cp.setMessageFlows()
	cp.setPools()
	// edge
	// message:edge:claim
	cp.setEdgeMessageClaim()
	// message:edge:refusal
	cp.setEdgeMessageRefusal()

}

// setParticipants ...
func (cp *collaborativeProcess) setParticipants() {
	// customer support
	/*
		test1 := cp.collaborativeProcessPool.CustomerSupportID
		test2 := cp.CustomerSupportID
		log.Printf("test1 %v, test2 %v", test1, test2)
	*/
	pool.SetParticipant(
		pool.DelegateParameter{
			PPT: cp.customerSupportParticipant(cp.collaboration()),
			T:   "id",
			PR:  "id",
			N:   "Customer Support",
			H:   []string{cp.CustomerSupportID, cp.CustomerSupportProcessID}})
	// customer
	pool.SetParticipant(
		pool.DelegateParameter{
			PPT: cp.customerParticipant(cp.collaboration()),
			T:   "id",
			PR:  "id",
			N:   "Customer",
			H:   []string{cp.CustomerID, cp.CustomerProcessID}})
}

// setMessageFlows ...
func (cp *collaborativeProcess) setMessageFlows() {
	// claim
	claim := cp.messageClaim()
	claim.SetID("flow", cp.CustomerToCustomerSupportMessageHash)
	claim.SetName("claim")
	claim.SetSourceRef("activity", cp.NoticeOfDefectHash)
	claim.SetTargetRef("activity", cp.CheckIncomingClaimHash)
	// refusal
	refusal := cp.messageRefusal()
	refusal.SetID("flow", cp.CustomerSupportToCustomerMessageHash)
	refusal.SetName("refusal")
	refusal.SetSourceRef("activity", cp.DenyWarrantyClaimHash)
	refusal.SetTargetRef("activity", cp.ReceiptWarrantyRefusalHash)
}

// setPools ...
func (cp *collaborativeProcess) setPools() {
	cp.setPoolCustomerSupport() // plane:shape:customersupport
	cp.setPoolCustomer()        // plane:shape:customer
}

/***************************************************************************************************************/
/* Customer Support Process */

// setPoolCustomerSupport ...
func (cp *collaborativeProcess) setPoolCustomerSupport() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: cp.shapePoolCustomerSupport(cp.plane()),
			T: "id",
			I: true,
			H: cp.CustomerSupportID,
			B: canvas.Bounds{X: 150, Y: 80, Width: 800, Height: 160}})
}

// setProcessCustomerSupport ...
func (cp *collaborativeProcess) setProcessCustomerSupport() {
	cp.customerSupportProcess().SetID("id", cp.CustomerSupportProcessID)
	cp.customerSupportProcess().SetIsExecutable(cp.CustomerSupportIsExecutable)
}

// setInnerElementsCustomerSupport ...
func (cp *collaborativeProcess) setInnerElementsCustomerSupport(e *process.Process) {
	e.SetStartEvent(1)
	e.SetEndEvent(1)
	e.SetTask(2)
	e.SetSequenceFlow(3)
}

// setCustomerSupportStartEvent ...
func (cp *collaborativeProcess) setCustomerSupportStartEvent() {
	e, d := cp.customerSupportStartEvent()
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: e,
			T:  "event",
			N:  "Begin of Customer Support Process",
			H:  []string{cp.CustomerSupportStartEventHash, cp.FromCustomerSupportStartEventHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.CustomerSupportStartEventHash,
			B: canvas.Bounds{X: 225, Y: 142}})
}

// setCustomerSupportEndEvent ...
func (cp *collaborativeProcess) setCustomerSupportEndEvent() {
	e, d := cp.customerSupportEndEvent()
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: e,
			T:  "event",
			N:  "End of Customer Support Process",
			H:  []string{cp.CustomerSupportEndEventHash, cp.FromDenyWarrantyClaimHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.CustomerSupportEndEventHash,
			B: canvas.Bounds{X: 822, Y: 142}})
}

// setCheckIncomingClaim ...
func (cp *collaborativeProcess) setCheckIncomingClaim() {
	e, d := cp.checkIncomingClaim()
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: e,
			T:  "activity",
			N:  "Check incoming claim",
			H:  []string{cp.CheckIncomingClaimHash, cp.FromCustomerSupportStartEventHash, cp.FromCheckIncomingClaimHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: cp.CheckIncomingClaimHash,
			B: canvas.Bounds{X: 320, Y: 120}})
}

// setDenyWarrantyClaim ...
func (cp *collaborativeProcess) setDenyWarrantyClaim() {
	e, d := cp.denyWarrantyClaim()
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: e,
			T:  "activity",
			N:  "Deny warranty claim",
			H:  []string{cp.DenyWarrantyClaimHash, cp.FromCheckIncomingClaimHash, cp.FromDenyWarrantyClaimHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: cp.DenyWarrantyClaimHash,
			B: canvas.Bounds{X: 580, Y: 120}})
}

// fromCustomerSupportStartEvent ...
func (cp *collaborativeProcess) fromCustomerSupportStartEvent() {
	e, d := cp.flowCustomerSupportStartEvent()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{cp.FromCustomerSupportStartEventHash, cp.CustomerSupportStartEventHash, cp.CheckIncomingClaimHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromCustomerSupportStartEventHash,
			W: []canvas.Waypoint{{X: 261, Y: 160}, {X: 320, Y: 160}}})
}

// fromCheckIncomingClaim ...
func (cp *collaborativeProcess) fromCheckIncomingClaim() {
	e, d := cp.flowCheckIncomingClaim()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			N:  "decide",
			ST: "activity",
			TT: "activity",
			H:  []string{cp.FromCheckIncomingClaimHash, cp.CheckIncomingClaimHash, cp.DenyWarrantyClaimHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromCheckIncomingClaimHash,
			W: []canvas.Waypoint{{X: 420, Y: 160}, {X: 580, Y: 160}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
			E: d,
			B: canvas.Bounds{X: 485, Y: 142, Width: 33, Height: 14}})
}

// fromDenyWarrantyClaim ...
func (cp *collaborativeProcess) fromDenyWarrantyClaim() {
	e, d := cp.flowDenyWarrantyClaim()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{cp.FromDenyWarrantyClaimHash, cp.DenyWarrantyClaimHash, cp.CustomerSupportEndEventHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromDenyWarrantyClaimHash,
			W: []canvas.Waypoint{{X: 680, Y: 160}, {X: 822, Y: 160}}})
}

/**
 * @Process
 * @Message Edge
 * @private
 * @private
 **/

// setEdgeMessageClaim ...
func (cp *collaborativeProcess) setEdgeMessageClaim() {
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: cp.edgeMessageClaim(cp.plane()),
			T: "flow",
			H: cp.CustomerToCustomerSupportMessageHash,
			W: []canvas.Waypoint{{X: 370, Y: 400}, {X: 370, Y: 200}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
			E: cp.edgeMessageClaim(cp.plane()),
			B: canvas.Bounds{X: 387, Y: 290, Width: 26, Height: 14}})
}

// setedgeMessageRefusal ...
func (cp *collaborativeProcess) setEdgeMessageRefusal() {
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: cp.edgeMessageRefusal(cp.plane()),
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
			T: "flow",
			H: cp.CustomerSupportToCustomerMessage.Suffix,
			W: []canvas.Waypoint{{X: 630, Y: 200}, {X: 630, Y: 400}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
<<<<<<< HEAD
			E: cp.plane().GetEdge(8),
			B: canvas.Bounds{X: 643, Y: 290, Width: 34, Height: 14}})
}

// setPools ...
func (cp *CProcess) setPools() {
	cp.setCustomerSupportPool()
	cp.setCustomerPool()
}

/* Customer Support Process */

// setPoolCustomerSupport ...
func (cp *CProcess) setCustomerSupportPool() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: cp.plane().GetShape(0),
			T: "participant",
			I: true,
			H: cp.CustomerSupportID.Suffix,
			B: canvas.Bounds{X: 150, Y: 80, Width: 800, Height: 160}})
}

// setProcessCustomerSupport ...
func (cp *CProcess) setCustomerSupportProcess() {
	cp.customerSupportProcess().SetID("process", cp.CustomerSupportProcess.Suffix)
	cp.customerSupportProcess().SetIsExecutable(cp.CustomerSupportIsExecutable)
}

// setCustomerSupportStartEvent ...
func (cp *CProcess) setCustomerSupportStartEvent() {
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: cp.customerSupportProcess().GetStartEvent(0),
			SH: cp.plane().GetShape(2),
			T:  "event",
			N:  "Begin of Customer Support Process",
			H:  []string{cp.CustomerSupportStartEvent.Suffix, cp.FromCustomerSupportStartEvent.Suffix},
			BS: canvas.Bounds{X: 225, Y: 142}})
	cp.fromCustomerSupportStartEvent()
}

// fromCustomerSupportStartEvent ...
func (cp *CProcess) fromCustomerSupportStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: cp.customerSupportProcess().GetSequenceFlow(0),
			ED: cp.plane().GetEdge(0),
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{cp.FromCustomerSupportStartEvent.Suffix, cp.CustomerSupportStartEvent.Suffix, cp.CheckIncomingClaim.Suffix},
			WP: []canvas.Waypoint{{X: 261, Y: 160}, {X: 320, Y: 160}}})
}

// setCheckIncomingClaim ...
func (cp *CProcess) setCheckIncomingClaim() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: cp.customerSupportProcess().GetTask(0),
			SH: cp.plane().GetShape(4),
			T:  "activity",
			N:  "Check incoming claim",
			H:  []string{cp.CheckIncomingClaim.Suffix, cp.FromCustomerSupportStartEvent.Suffix, cp.FromCheckIncomingClaim.Suffix},
			BS: canvas.Bounds{X: 320, Y: 120}})
	cp.fromCheckIncomingClaim()
}

// fromCheckIncomingClaim ...
func (cp *CProcess) fromCheckIncomingClaim() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: cp.customerSupportProcess().GetSequenceFlow(1),
			ED: cp.plane().GetEdge(1),
			T:  "flow",
			N:  "decide",
			ST: "activity",
			TT: "activity",
			H:  []string{cp.FromCheckIncomingClaim.Suffix, cp.CheckIncomingClaim.Suffix, cp.DenyWarrantyClaim.Suffix},
			WP: []canvas.Waypoint{{X: 420, Y: 160}, {X: 580, Y: 160}},
			BS: canvas.Bounds{X: 485, Y: 142, Width: 33, Height: 14}})
}

// setDenyWarrantyClaim ...
func (cp *CProcess) setDenyWarrantyClaim() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: cp.customerSupportProcess().GetTask(1),
			T:  "activity",
			N:  "Deny warranty claim",
			H:  []string{cp.DenyWarrantyClaim.Suffix, cp.FromCheckIncomingClaim.Suffix, cp.FromDenyWarrantyClaim.Suffix},
			SH: cp.plane().GetShape(5),
			BS: canvas.Bounds{X: 580, Y: 120}})
	cp.fromDenyWarrantyClaim()
}

// fromDenyWarrantyClaim ...
func (cp *CProcess) fromDenyWarrantyClaim() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: cp.customerSupportProcess().GetSequenceFlow(2),
			ED: cp.plane().GetEdge(2),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{cp.FromDenyWarrantyClaim.Suffix, cp.DenyWarrantyClaim.Suffix, cp.CustomerSupportEndEvent.Suffix},
			WP: []canvas.Waypoint{{X: 680, Y: 160}, {X: 822, Y: 160}}})
}

// setCustomerSupportEndEvent ...
func (cp *CProcess) setCustomerSupportEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: cp.customerSupportProcess().GetEndEvent(0),
			SH: cp.plane().GetShape(3),
			T:  "event",
			N:  "End of Customer Support Process",
			H:  []string{cp.CustomerSupportEndEvent.Suffix, cp.FromDenyWarrantyClaim.Suffix},
			BS: canvas.Bounds{X: 822, Y: 142}})
}
=======
			E: cp.edgeMessageRefusal(cp.plane()),
			B: canvas.Bounds{X: 643, Y: 290, Width: 34, Height: 14}})
}

/***************************************************************************************************************/
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4

/**** Customer Process ****/

// setPoolCustomer
<<<<<<< HEAD
func (cp *CProcess) setCustomerPool() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: cp.plane().GetShape(1),
			T: "participant",
			I: true,
			H: cp.CustomerID.Suffix,
=======
func (cp *collaborativeProcess) setPoolCustomer() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: cp.shapePoolCustomer(cp.plane()),
			T: "id",
			I: true,
			H: cp.CustomerID,
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
			B: canvas.Bounds{X: 150, Y: 360, Width: 800, Height: 160}})
}

// setProcessCustomer ...
<<<<<<< HEAD
func (cp *CProcess) setCustomerProcess() {
	cp.customerProcess().SetID("process", cp.CustomerProcess.Suffix)
	cp.customerProcess().SetIsExecutable(cp.CustomerIsExecutable)
=======
func (cp *collaborativeProcess) setProcessCustomer() {
	cp.customerProcess().SetID("id", cp.CustomerProcessID)
	cp.customerProcess().SetIsExecutable(cp.CustomerIsExecutable)
}

// setInnerElementsCustomer ...
func (cp *collaborativeProcess) setInnerElementsCustomer(p *process.Process) {
	// Event
	p.SetStartEvent(1)
	p.SetEndEvent(1)
	// Task
	p.SetTask(2)
	// Event
	p.SetIntermediateCatchEvent(1)
	// Sequence
	p.SetSequenceFlow(4)
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

// setCustomerStartEvent ...
<<<<<<< HEAD
func (cp *CProcess) setCustomerStartEvent() {
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: cp.customerProcess().GetStartEvent(0),
			SH: cp.plane().GetShape(6),
			T:  "event",
			N:  "Begin of Customer Process",
			H:  []string{cp.CustomerStartEvent.Suffix, cp.FromCustomerStartEvent.Suffix},
			BS: canvas.Bounds{X: 225, Y: 422}})
	cp.fromCustomerStartEvent()
}

// fromCustomerStartEvent ...
func (cp *CProcess) fromCustomerStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: cp.customerProcess().GetSequenceFlow(0),
			ED: cp.plane().GetEdge(3),
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{cp.FromCustomerStartEvent.Suffix, cp.CustomerStartEvent.Suffix, cp.NoticeOfDefect.Suffix},
			WP: []canvas.Waypoint{{X: 261, Y: 440}, {X: 320, Y: 440}}})
=======
func (cp *collaborativeProcess) setCustomerStartEvent() {
	e, d := cp.customerStartEvent()
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: e,
			T:  "event",
			N:  "Begin of Customer Process",
			H:  []string{cp.CustomerStartEventHash, cp.FromCustomerStartEventHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.CustomerStartEventHash,
			B: canvas.Bounds{X: 225, Y: 422}})
}

// setCustomerEndEvent ...
func (cp *collaborativeProcess) setCustomerEndEvent() {
	e, d := cp.customerEndEvent()
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: e,
			T:  "event",
			N:  "End of Customer Process",
			H:  []string{cp.CustomerEndEventHash, cp.FromReceiptWarrantyRefusalHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.CustomerEndEventHash,
			B: canvas.Bounds{X: 822, Y: 422}})
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

// setNoticeOfDefect ...
<<<<<<< HEAD
func (cp *CProcess) setNoticeOfDefect() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: cp.customerProcess().GetTask(0),
			SH: cp.plane().GetShape(8),
			T:  "activity",
			N:  "Notice of defect",
			H:  []string{cp.NoticeOfDefect.Suffix, cp.FromCustomerStartEvent.Suffix, cp.FromNoticeOfDefect.Suffix},
			BS: canvas.Bounds{X: 320, Y: 400}})
	cp.fromNoticeOfDefect()
}

// setMessageClaim ...
func (cp *CProcess) setMessageClaim() {
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: cp.plane().GetEdge(7),
			T: "flow",
			H: cp.CustomerToCustomerSupportMessage.Suffix,
			W: []canvas.Waypoint{{X: 370, Y: 400}, {X: 370, Y: 200}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
			E: cp.plane().GetEdge(7),
			B: canvas.Bounds{X: 387, Y: 290, Width: 26, Height: 14}})
}

// fromNoticeOfDefect ...
func (cp *CProcess) fromNoticeOfDefect() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: cp.customerProcess().GetSequenceFlow(1),
			ED: cp.plane().GetEdge(4),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{cp.FromNoticeOfDefect.Suffix, cp.NoticeOfDefect.Suffix, cp.WaitingForAnswer.Suffix},
			WP: []canvas.Waypoint{{X: 420, Y: 440}, {X: 482, Y: 440}}})
}

// setWaitingForAnswer ...
func (cp *CProcess) setWaitingForAnswer() {
	elements.SetIntermediateCatchEvent(
		elements.DelegateParameter{
			ISTED: true,
			ICE:   cp.customerProcess().GetIntermediateCatchEvent(0),
			SH:    cp.plane().GetShape(9),
			T:     "event",
			N:     "Waiting for answer",
			TD:    "PT1M",
			H:     []string{cp.WaitingForAnswer.Suffix, cp.FromNoticeOfDefect.Suffix, cp.FromWaitingForAnswer.Suffix},
			TEDH:  cp.TimerEventDefinitionWaitingForAnswer.Suffix,
			BS:    canvas.Bounds{X: 482, Y: 422}})
	cp.fromWaitingForAnswer()
}

// fromWaitingForAnswer ...
func (cp *CProcess) fromWaitingForAnswer() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: cp.customerProcess().GetSequenceFlow(2),
			ED: cp.plane().GetEdge(5),
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{cp.FromWaitingForAnswer.Suffix, cp.WaitingForAnswer.Suffix, cp.ReceiptWarrantyRefusal.Suffix},
			WP: []canvas.Waypoint{{X: 518, Y: 440}, {X: 580, Y: 440}}})
}

// setReceiptWarrantyRefusal ...
func (cp *CProcess) setReceiptWarrantyRefusal() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: cp.customerProcess().GetTask(1),
			SH: cp.plane().GetShape(10),
			T:  "activity",
			N:  "Receipt warranty refusal",
			H:  []string{cp.ReceiptWarrantyRefusal.Suffix, cp.FromWaitingForAnswer.Suffix, cp.FromReceiptWarrantyRefusal.Suffix},
			BS: canvas.Bounds{X: 580, Y: 400}})
	cp.fromReceiptWarrantyRefusal()
}

// fromReceiptWarrantyRefusal ...
func (cp *CProcess) fromReceiptWarrantyRefusal() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: cp.customerProcess().GetSequenceFlow(3),
			ED: cp.plane().GetEdge(6),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{cp.FromReceiptWarrantyRefusal.Suffix, cp.ReceiptWarrantyRefusal.Suffix, cp.CustomerEndEvent.Suffix},
			WP: []canvas.Waypoint{{X: 680, Y: 440}, {X: 822, Y: 440}}})
}

// setCustomerEndEvent ...
func (cp *CProcess) setCustomerEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: cp.customerProcess().GetEndEvent(0),
			SH: cp.plane().GetShape(7),
			T:  "event",
			N:  "End of Customer Process",
			H:  []string{cp.CustomerEndEvent.Suffix, cp.FromReceiptWarrantyRefusal.Suffix},
			BS: canvas.Bounds{X: 822, Y: 422}})
}

// setDiagram ...
func (cp *CProcess) setDiagram() {
	// diagram attributes
	var n int64 = 1
	cp.diagram().SetID("diagram", n)
	// plane attributes
	p := cp.plane()
	p.SetID("plane", n)
	p.SetElement("collaboration", cp.Collaboration.Suffix)
}
=======
func (cp *collaborativeProcess) setNoticeOfDefect() {
	e, d := cp.noticeOfDefect()
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: e,
			T:  "activity",
			N:  "Notice of defect",
			H:  []string{cp.NoticeOfDefectHash, cp.FromCustomerStartEventHash, cp.FromNoticeOfDefectHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: cp.NoticeOfDefectHash,
			B: canvas.Bounds{X: 320, Y: 400}})
}

// setWaitingForAnswer ...
func (cp *collaborativeProcess) setWaitingForAnswer() {
	e, d := cp.waitingForAnswer()
	elements.SetIntermediateCatchEvent(e, "Waiting for answer", cp.WaitingForAnswerHash, cp.FromNoticeOfDefectHash, cp.FromWaitingForAnswerHash)
	t := cp.waitingForAnswerTimerEventDefinition(e)
	definitions.SetTimerEventDefinition(t, "PT1M", cp.TimerEventDefinitionWaitingForAnswerHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.WaitingForAnswerHash,
			B: canvas.Bounds{X: 482, Y: 422}})
}

// setReceiptWarrantyRefusal ...
func (cp *collaborativeProcess) setReceiptWarrantyRefusal() {
	e, d := cp.receiptWarrantyRefusal()
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: e,
			T:  "activity",
			N:  "Receipt warranty refusal",
			H:  []string{cp.ReceiptWarrantyRefusalHash, cp.FromWaitingForAnswerHash, cp.FromReceiptWarrantyRefusalHash}})
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: cp.ReceiptWarrantyRefusalHash,
			B: canvas.Bounds{X: 580, Y: 400}})
}

/**
 * @Customer Process IV
 * @Sequence
 * @private fromCustomerStartEvent
 * @private fromNoticeOfDefect
 * @private fromWaitingForAnswer
 * @private fromReceiptWarrantyRefusal
 **/

// fromCustomerStartEvent ...
func (cp *collaborativeProcess) fromCustomerStartEvent() {
	e, d := cp.flowCustomerStartEvent()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{cp.FromCustomerStartEventHash, cp.CustomerStartEventHash, cp.NoticeOfDefectHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromCustomerStartEventHash,
			W: []canvas.Waypoint{{X: 261, Y: 440}, {X: 320, Y: 440}}})
}

// fromNoticeOfDefect ...
func (cp *collaborativeProcess) fromNoticeOfDefect() {
	e, d := cp.flowNoticeOfDefect()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{cp.FromNoticeOfDefectHash, cp.NoticeOfDefectHash, cp.WaitingForAnswerHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromNoticeOfDefectHash,
			W: []canvas.Waypoint{{X: 420, Y: 440}, {X: 482, Y: 440}}})
}

// fromWaitingForAnswer ...
func (cp *collaborativeProcess) fromWaitingForAnswer() {
	e, d := cp.flowWaitingForAnswer()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{cp.FromWaitingForAnswerHash, cp.WaitingForAnswerHash, cp.ReceiptWarrantyRefusalHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromWaitingForAnswerHash,
			W: []canvas.Waypoint{{X: 518, Y: 440}, {X: 580, Y: 440}}})
}

// fromReceiptWarrantyRefusal ...
func (cp *collaborativeProcess) fromReceiptWarrantyRefusal() {
	e, d := cp.flowReceiptWarrantyRefusal()
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: e,
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{cp.FromReceiptWarrantyRefusalHash, cp.ReceiptWarrantyRefusalHash, cp.CustomerEndEventHash}})
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromReceiptWarrantyRefusalHash,
			W: []canvas.Waypoint{{X: 680, Y: 440}, {X: 822, Y: 440}}})
}

/***************************************************************************************************************/
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4

/**** Getter ****/

// collaboration ...
<<<<<<< HEAD
func (cp CProcess) collaboration() *collaboration.Collaboration {
	return cp.Def.GetCollaboration()
}

// customerSupportProces ...
func (cp CProcess) customerSupportProcess() *process.Process {
	return cp.Def.GetProcess(0)
}

// customerProcess ...
func (cp CProcess) customerProcess() *process.Process {
	return cp.Def.GetProcess(1)
}

// diagram ...
func (cp CProcess) diagram() *canvas.Diagram {
	return cp.Def.GetDiagram(0)
}

// plane ...
func (cp CProcess) plane() *canvas.Plane {
	return cp.diagram().GetPlane()
=======
func (cp collaborativeProcess) collaboration() *pool.Collaboration {
	return cp.def.GetCollaboration()
}

// customerSupportParticipant ...
func (cp collaborativeProcess) customerSupportParticipant(e *pool.Collaboration) *pool.Participant {
	return e.GetParticipant(0)
}

// customerSupportProces ...
func (cp collaborativeProcess) customerSupportProcess() *process.Process {
	return cp.def.GetProcess(0)
}

// customerSupportStartEvent ...
func (cp collaborativeProcess) customerSupportStartEvent() (*elements.StartEvent, *canvas.Shape) {
	start := cp.customerSupportProcess().GetStartEvent(0)
	shape := cp.shapeCustomerSupportStartEvent(cp.plane())
	return start, shape
}

// customerSupportEndEvent ...
func (cp collaborativeProcess) customerSupportEndEvent() (*elements.EndEvent, *canvas.Shape) {
	end := cp.customerSupportProcess().GetEndEvent(0)
	shape := cp.shapeCustomerSupportEndEvent(cp.plane())
	return end, shape
}

// checkIncomingClaim ...
func (cp collaborativeProcess) checkIncomingClaim() (tasks.TASK_PTR, *canvas.Shape) {
	task := cp.customerSupportProcess().GetTask(0)
	shape := cp.shapeCheckIncomingClaim(cp.plane())
	return task, shape
}

// denyWarrantyClaim ...
func (cp collaborativeProcess) denyWarrantyClaim() (*tasks.Task, *canvas.Shape) {
	task := cp.customerSupportProcess().GetTask(1)
	shape := cp.shapeDenyWarrantyClaim(cp.plane())
	return task, shape
}

// flowCustomerSupportStartEvent ...
func (cp collaborativeProcess) flowCustomerSupportStartEvent() (*flow.SequenceFlow, *canvas.Edge) {
	flow := cp.customerSupportProcess().GetSequenceFlow(0)
	edge := cp.edgeFromCustomerSupportStartEvent(cp.plane())
	return flow, edge
}

// flowCheckIncomingClaim ...
func (cp collaborativeProcess) flowCheckIncomingClaim() (*flow.SequenceFlow, *canvas.Edge) {
	flow := cp.customerSupportProcess().GetSequenceFlow(1)
	edge := cp.edgeFromCheckIncomingClaim(cp.plane())
	return flow, edge
}

// flowDenyWarrantyClaim ...
func (cp collaborativeProcess) flowDenyWarrantyClaim() (*flow.SequenceFlow, *canvas.Edge) {
	flow := cp.customerSupportProcess().GetSequenceFlow(2)
	edge := cp.edgeFromDenyWarrantyClaim(cp.plane())
	return flow, edge
}

/*** CUSTOMER ***/

// CustomerProcess ...
func (cp collaborativeProcess) customerProcess() *process.Process {
	return cp.def.GetProcess(1)
}

// customerParticipant ...
func (cp collaborativeProcess) customerParticipant(e *pool.Collaboration) *pool.Participant {
	return e.GetParticipant(1)
}

// customerStartEvent ...
func (cp collaborativeProcess) customerStartEvent() (*elements.StartEvent, *canvas.Shape) {
	start := cp.customerProcess().GetStartEvent(0)
	shape := cp.shapeCustomerStartEvent(cp.plane())
	return start, shape
}

// customerEndEvent ...
func (cp collaborativeProcess) customerEndEvent() (*elements.EndEvent, *canvas.Shape) {
	end := cp.customerProcess().GetEndEvent(0)
	shape := cp.shapeCustomerEndEvent(cp.plane())
	return end, shape
}

// noticeOfDefect ...
func (cp collaborativeProcess) noticeOfDefect() (*tasks.Task, *canvas.Shape) {
	task := cp.customerProcess().GetTask(0)
	shape := cp.shapeNoticeOfDefect(cp.plane())
	return task, shape
}

// waitingForAnswer ...
func (cp collaborativeProcess) waitingForAnswer() (*elements.IntermediateCatchEvent, *canvas.Shape) {
	timer := cp.customerProcess().GetIntermediateCatchEvent(0)
	shape := cp.shapeWaitingForAnswer(cp.plane())
	return timer, shape
}

// waitingForAnswerTimerEventDefinition ...
func (cp collaborativeProcess) waitingForAnswerTimerEventDefinition(e *elements.IntermediateCatchEvent) *definitions.TimerEventDefinition {
	e.SetTimerEventDefinition()
	return e.GetTimerEventDefinition()
}

// receiptWarrantyRefusal ...
func (cp collaborativeProcess) receiptWarrantyRefusal() (*tasks.Task, *canvas.Shape) {
	task := cp.customerProcess().GetTask(1)
	shape := cp.shapeReceiptWarrantyRefusal(cp.plane())
	return task, shape
}

// flowCustomerStartEvent ...
func (cp collaborativeProcess) flowCustomerStartEvent() (*flow.SequenceFlow, *canvas.Edge) {
	flow := cp.customerProcess().GetSequenceFlow(0)
	edge := cp.edgeFromCustomerStartEvent(cp.plane())
	return flow, edge
}

// flowNoticeOfDefect ...
func (cp collaborativeProcess) flowNoticeOfDefect() (*flow.SequenceFlow, *canvas.Edge) {
	flow := cp.customerProcess().GetSequenceFlow(1)
	edge := cp.edgeFromNoticeOfDefect(cp.plane())
	return flow, edge
}

// flowWaitingForAnswer ...
func (cp collaborativeProcess) flowWaitingForAnswer() (*flow.SequenceFlow, *canvas.Edge) {
	flow := cp.customerProcess().GetSequenceFlow(2)
	edge := cp.edgeFromWaitingForAnswer(cp.plane())
	return flow, edge
}

// flowReceiptWarrantyRefusal ...
func (cp collaborativeProcess) flowReceiptWarrantyRefusal() (*flow.SequenceFlow, *canvas.Edge) {
	flow := cp.customerProcess().GetSequenceFlow(3)
	edge := cp.edgeFromReceiptWarrantyRefusal(cp.plane())
	return flow, edge
}

/**** Messages ****/

// messageClaim ...
func (cp collaborativeProcess) messageClaim() *flow.MessageFlow {
	return cp.collaboration().GetMessageFlow(0)
}

// messageRefusal ...
func (cp collaborativeProcess) messageRefusal() *flow.MessageFlow {
	return cp.collaboration().GetMessageFlow(1)
}

/***************************************************************************************************************/

/**** Diagram & Plane ****/

// setDiagram ...
func (cp *collaborativeProcess) setDiagram() {
	// diagram attributes
	var n int64 = 1
	diagram := cp.diagram()
	diagram.SetID("diagram", n)
	/*
		// plane attributes
		p := cp.plane()
		p.SetID("plane", n)
		p.SetElement("id", cp.CollaborationID)
	*/
}

// diagram ...
func (cp collaborativeProcess) diagram() *canvas.Diagram {
	return cp.def.GetDiagram(0)
}

// setPlane ...
func (cp *collaborativeProcess) setPlane() {
	var n int64 = 1
	p := cp.plane()
	p.SetID("plane", n)
	p.SetElement("id", cp.CollaborationID)
}

// plane ...
func (cp collaborativeProcess) plane() *canvas.Plane {
	return cp.diagram().GetPlane()
}

/**** Shapes ****/

// shapePoolCustomerSupport ...
func (cp collaborativeProcess) shapePoolCustomerSupport(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(0)
}

// shapePoolCustomer ...
func (cp collaborativeProcess) shapePoolCustomer(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(1)
}

// shapeCustomerSupportStartEvent ...
func (cp collaborativeProcess) shapeCustomerSupportStartEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(2)
}

// shapeCustomerSupportEndEvent ...
func (cp collaborativeProcess) shapeCustomerSupportEndEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(3)
}

// shapeCheckIncomingClaim ...
func (cp collaborativeProcess) shapeCheckIncomingClaim(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(4)
}

// shapeDenyWarrantyClaim ...
func (cp collaborativeProcess) shapeDenyWarrantyClaim(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(5)
}

// shapeCustomerStartEvent ...
func (cp collaborativeProcess) shapeCustomerStartEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(6)
}

// shapeCustomerEndEvent ...
func (cp collaborativeProcess) shapeCustomerEndEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(7)
}

// shapeNoticeOfDefect ...
func (cp collaborativeProcess) shapeNoticeOfDefect(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(8)
}

// shapeWaitingForAnswer ...
func (cp collaborativeProcess) shapeWaitingForAnswer(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(9)
}

// shapeReceiptWarrantyRefusal ...
func (cp collaborativeProcess) shapeReceiptWarrantyRefusal(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(10)
}

/**** Edges ****/

// edgeFromCustomerSupportStartEvent ...
func (cp collaborativeProcess) edgeFromCustomerSupportStartEvent(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(0)
}

// edgeFromCheckIncomingClaim ...
func (cp collaborativeProcess) edgeFromCheckIncomingClaim(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(1)
}

// edgeFromDenyWarrantyClaim ...
func (cp collaborativeProcess) edgeFromDenyWarrantyClaim(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(2)
}

// edgeFromCustomerStartEvent ...
func (cp collaborativeProcess) edgeFromCustomerStartEvent(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(3)
}

// edgeFromNoticeOfDefect ...
func (cp collaborativeProcess) edgeFromNoticeOfDefect(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(4)
}

// edgeFromWaitingForAnswer ...
func (cp collaborativeProcess) edgeFromWaitingForAnswer(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(5)
}

// edgeFromReceiptWarrantyRefusal ...
func (cp collaborativeProcess) edgeFromReceiptWarrantyRefusal(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(6)
}

// edgeMessageClaim ...
func (cp collaborativeProcess) edgeMessageClaim(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(7)
}

// edgeMessageRefusal ...
func (cp collaborativeProcess) edgeMessageRefusal(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(8)
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}
