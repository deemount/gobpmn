package examples

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
)

/**************************************************************************************/
/**
 * @BASE CLASS
 *
 *
 **/

// CollaborativeProcess ...
type CollaborativeProcess interface {
	Create() collaborativeProcess
}

/**************************************************************************************/
/**
 * @OBJECTS
 * Global
 * @private collaborativeProcess
 * @private collaborativeProcessPool
 * @private collaborativeProcessMessage
 * Customer Support
 * @private collaborativeProcessCustomerSupportEvent
 * @private collaborativeProcessCustomerSupportSequence
 * @private collaborativeProcessCustomerSupportTask
 * Customer
 * @private collaborativeProcessCustomerEvent
 * @private collaborativeProcessCustomerSequence
 * @private collaborativeProcessCustomerTask
 *
 **/

// collaborativeProcess refers to models definition repository, globals and hashes
type collaborativeProcess struct {
	def core.DefinitionsRepository
	/**** Global ****/
	collaborativeProcessPool
	collaborativeProcessMessage
	/**** Customer Support Hashes ****/
	collaborativeProcessCustomerSupportEvent
	collaborativeProcessCustomerSupportSequence
	collaborativeProcessCustomerSupportTask
	/**** Customer Hashes ****/
	collaborativeProcessCustomerEvent
	collaborativeProcessCustomerSequence
	collaborativeProcessCustomerTask
}

// collaborativeProcessPool ...
type collaborativeProcessPool struct {
	CustomerSupportIsExecutable bool
	CustomerIsExecutable        bool
	CollaborationID             string
	CustomerSupportID           string
	CustomerID                  string
	CustomerSupportProcessID    string
	CustomerProcessID           string
}

// collaborativeProcessMessage ...
type collaborativeProcessMessage struct {
	CustomerToCustomerSupportMessageHash string
	CustomerSupportToCustomerMessageHash string
}

/**** Customer Support Hashes ****/

// collaborativeProcessCustomerSupportEvent ...
type collaborativeProcessCustomerSupportEvent struct {
	// @StartEvent
	CustomerSupportStartEventHash string
	// @EndEvent
	CustomerSupportEndEventHash string
}

// collaborativeProcessCustomerSupportSequence ...
type collaborativeProcessCustomerSupportSequence struct {
	// @StartEvent
	FromCustomerSupportStartEventHash string
	// @EndEvent
	FromCustomerSupportEndEventHash string
	// @CheckIncomingClaim
	FromCheckIncomingClaimHash string
	// @DenyWarrantyClaim
	FromDenyWarrantyClaimHash string
}

type collaborativeProcessCustomerSupportTask struct {
	// @CheckIncomingClaim
	CheckIncomingClaimHash string
	// @DenyWarrantyClaim
	DenyWarrantyClaimHash string
}

/**** Customer Hashes ****/

// collaborativeProcessCustomerEvent ...
type collaborativeProcessCustomerEvent struct {
	// @StartEvent
	CustomerStartEventHash string
	// @EndEvent
	CustomerEndEventHash string
	// @WaitingForAnswer
	TimerEventDefinitionWaitingForAnswerHash string
}

// collaborativeProcessCustomerSequence ...
type collaborativeProcessCustomerSequence struct {
	// @StartEvent
	FromCustomerStartEventHash string
	// @EndEvent
	FromCustomerEndEventHash string
	// @NoticeOfDefect
	FromNoticeOfDefectHash string
	// @WaitingForAnswer
	FromWaitingForAnswerHash string
	// @ReceiptWarrantyRefusal
	FromReceiptWarrantyRefusalHash string
}

// collaborativeProcessCustomerTask ...
type collaborativeProcessCustomerTask struct {
	// @NoticeOfDefect
	NoticeOfDefectHash string
	// @WaitingForAnswer
	WaitingForAnswerHash        string
	WaitingForAnswerShapeIDHash string
	// @ReceiptWarrantyRefusal
	ReceiptWarrantyRefusalHash string
}

/**************************************************************************************/
/**
 * @BUILDER
 *
 *
 **/

// NewCollaborativeProcess ...
func NewCollaborativeProcess() CollaborativeProcess {
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

// Create ...
func (cp collaborativeProcess) Create() collaborativeProcess {
	// Elements
	core.SetMainElements(cp.def, 2)
	cp.setInnerElements()
	cp.setDefinitionsAttributes()
	cp.setCollaboration()
	// Process
	cp.setProcessCustomerSupport()
	cp.setProcessCustomer()
	// Customer Support Event
	cp.setCustomerSupportStartEvent()
	cp.setCustomerSupportEndEvent()
	// Customer Support Task
	cp.setCheckIncomingClaim()
	cp.setDenyWarrantyClaim()
	// Customer Support Sequence
	cp.fromCustomerSupportStartEvent()
	cp.fromCheckIncomingClaim()
	cp.fromDenyWarrantyClaim()
	// Customer Event
	cp.setCustomerStartEvent()
	cp.setCustomerEndEvent()
	// Customer Task
	cp.setNoticeOfDefect()
	cp.setWaitingForAnswer()
	cp.setReceiptWarrantyRefusal()
	// Customer Sequence
	cp.fromCustomerStartEvent()
	cp.fromNoticeOfDefect()
	cp.fromWaitingForAnswer()
	cp.fromReceiptWarrantyRefusal()
	// Canvas
	cp.setDiagram()
	cp.setPlane()
	return cp
}

// Def ...
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

	// Canvas
	diagram := cp.diagram()
	diagram.SetPlane()
	plane := cp.plane()
	plane.SetShape(11)
	plane.SetEdge(9)
}

// setDefinitionsAttributes ...
func (cp *collaborativeProcess) setDefinitionsAttributes() {
	cp.def.SetDefaultAttributes()
}

// setCollaboration ...
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
			T: "flow",
			H: cp.CustomerSupportToCustomerMessageHash,
			W: []canvas.Waypoint{{X: 630, Y: 200}, {X: 630, Y: 400}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
			E: cp.edgeMessageRefusal(cp.plane()),
			B: canvas.Bounds{X: 643, Y: 290, Width: 34, Height: 14}})
}

/***************************************************************************************************************/

/**** Customer Process ****/

/**
 * @Customer Process I
 * @Process
 * @private setPoolCustomer
 * @private setProcessCustomer
 **/

// setPoolCustomer
func (cp *collaborativeProcess) setPoolCustomer() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: cp.shapePoolCustomer(cp.plane()),
			T: "id",
			I: true,
			H: cp.CustomerID,
			B: canvas.Bounds{X: 150, Y: 360, Width: 800, Height: 160}})
}

// setProcessCustomer ...
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
}

/**
 * @Customer Process II
 * @Event
 * @SetCustomerStartEvent
 * @SetCustomerEndEvent
 **/

// setCustomerStartEvent ...
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
}

/**
 * @Customer Process III
 * @Task
 * @private setNoticeOfDefect
 * @private setWaitingForAnswer
 * @private setReceiptWarrantyRefusal
 **/

// setNoticeOfDefect ...
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

/**** Getter ****/

// collaboration ...
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
}
