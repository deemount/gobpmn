package examples

/**************************************************************************************/
/**
 * @Import
 *
 * Import models and utils package to build the model
 **/

import (
	"github.com/deemount/gobpmn/models/canvas"
	"github.com/deemount/gobpmn/models/core"
	"github.com/deemount/gobpmn/models/events/definitions"
	"github.com/deemount/gobpmn/models/events/elements"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/models/process"
	"github.com/deemount/gobpmn/models/tasks"
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
		// def: core.NewCore(), // Notice: creates <Core></Core> element as root, which actually the unmarshaler isn't possible to parse
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
	// Diagram
	cp.setDiagram()
	return cp
}

// Def ...
func (cp collaborativeProcess) Def() core.DefinitionsRepository {
	return cp.def
}

/**************************************************************************************/

/**
 * @Setters I
 *
 * @setInnerElements
 * @@Model: GetCollaboration
 * @@Model: GetDiagram
 * @@Model: GetCustomerSupportProcess
 * @@Model: GetCustomerProcess
 *
 * @setDefinitionsAttributes
 * @@models.Definitions: SetDefaultAttributes
 *
 * @setCollaboration
 * @@Model: GetParticipant
 * @@Model: GetCollaboration
 **/

// setInnerElements ...
func (cp *collaborativeProcess) setInnerElements() {
	// Collaboration
	collaboration := cp.GetCollaboration()
	collaboration.SetParticipant(2)
	collaboration.SetMessageFlow(2)

	// Process
	// Customer Support
	customerSupport := cp.GetCustomerSupportProcess()
	// Event
	customerSupport.SetStartEvent(1)
	customerSupport.SetEndEvent(1)
	// Task
	customerSupport.SetTask(2)
	// Sequence
	customerSupport.SetSequenceFlow(3)

	// Process
	// Customer
	customer := cp.GetCustomerProcess()
	// Event
	customer.SetStartEvent(1)
	customer.SetEndEvent(1)
	// Task
	customer.SetTask(2)
	// Event
	customer.SetIntermediateCatchEvent(1)
	// Sequence
	customer.SetSequenceFlow(4)

	// Diagram
	diagram := cp.GetDiagram()
	diagram.SetPlane()
	plane := cp.GetPlane()
	plane.SetShape(11)
	plane.SetEdge(9)
}

// setDefinitionsAttributes ...
func (cp *collaborativeProcess) setDefinitionsAttributes() {
	cp.def.SetDefaultAttributes()
}

// setCollaboration ...
func (cp *collaborativeProcess) setCollaboration() {
	// generics
	customerSupportParticipant := cp.GetCustomerSupportParticipant(cp.GetCollaboration())
	customerParticipant := cp.GetCustomerParticipant(cp.GetCollaboration())
	claim := cp.GetMessageClaim()
	refusal := cp.GetMessageRefusal()
	// generics
	cp.GetCollaboration().SetID("id", cp.CollaborationID)
	// participant attributes
	// generics
	// customer support
	pool.SetParticipant(customerSupportParticipant, "id", "id", "Customer Support", cp.CustomerSupportID, cp.CustomerSupportProcessID)
	// customer
	pool.SetParticipant(customerParticipant, "id", "id", "Customer", cp.CustomerID, cp.CustomerProcessID)

	// message flow
	// claim
	claim.SetID("flow", cp.CustomerToCustomerSupportMessageHash)
	claim.SetName("claim")
	claim.SetSourceRef("activity", cp.NoticeOfDefectHash)
	claim.SetTargetRef("activity", cp.CheckIncomingClaimHash)
	// refusal
	refusal.SetID("flow", cp.CustomerSupportToCustomerMessageHash)
	refusal.SetName("refusal")
	refusal.SetSourceRef("activity", cp.DenyWarrantyClaimHash)
	refusal.SetTargetRef("activity", cp.ReceiptWarrantyRefusalHash)

	// shape
	// plane:shape:customersupport
	cp.setPoolCustomerSupport()
	// plane:shape:customer
	cp.setPoolCustomer()

	// edge
	// message:edge:claim
	cp.setEdgeMessageClaim()
	// message:edge:refusal
	cp.setEdgeMessageRefusal()

}

/**** Customer Support Process ****/

/**
 * @Customer Support Process
 * @Process
 * @private setPoolCustomerSupport
 * @private setProcessCustomerSupport
 **/

// setCustomerSupportPool ...
func (cp *collaborativeProcess) setPoolCustomerSupport() {
	e := cp.GetShapePoolCustomerSupport(cp.GetPlane())
	canvas.SetPool(e, "id", true, cp.CustomerSupportID,
		canvas.Bounds{X: 150, Y: 80, Width: 800, Height: 160})
}

// setCustomerSupportProcess ...
func (cp *collaborativeProcess) setProcessCustomerSupport() {
	cp.GetCustomerSupportProcess().SetID("id", cp.CustomerSupportProcessID)
	cp.GetCustomerSupportProcess().SetIsExecutable(cp.CustomerSupportIsExecutable)
}

/**
 * @Customer Support Process
 * @Event
 * @private setCustomerSupportStartEvent
 * @private setCustomerSupportEndEvent
 **/

// setCustomerSupportStartEvent ...
func (cp *collaborativeProcess) setCustomerSupportStartEvent() {
	e, d := cp.GetCustomerSupportStartEvent()
	elements.SetStartEvent(e, "Begin of Customer Support Process", cp.CustomerSupportStartEventHash, cp.FromCustomerSupportStartEventHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.CustomerSupportStartEventHash,
			B: canvas.Bounds{X: 225, Y: 142, Width: 36, Height: 36}})
}

// setCustomerSupportEndEvent ...
func (cp *collaborativeProcess) setCustomerSupportEndEvent() {
	e, d := cp.GetCustomerSupportEndEvent()
	elements.SetEndEvent(e, "End of Customer Support Process", cp.CustomerSupportEndEventHash, cp.FromDenyWarrantyClaimHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.CustomerSupportEndEventHash,
			B: canvas.Bounds{X: 822, Y: 142, Width: 36, Height: 36}})
}

/**
 * @Customer Support Process
 * @Task
 * @private setCheckIncomingClaim
 * @private setDenyWarrantyClaim
 **/

// setCheckIncomingClaim ...
func (cp *collaborativeProcess) setCheckIncomingClaim() {
	e, d := cp.GetCheckIncomingClaim()
	tasks.SetTask(e, "Check incoming claim", cp.CheckIncomingClaimHash, cp.FromCustomerSupportStartEventHash, cp.FromCheckIncomingClaimHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: cp.CheckIncomingClaimHash,
			B: canvas.Bounds{X: 320, Y: 120, Width: 100, Height: 80}})
}

// setDenyWarrantyClaim ...
func (cp *collaborativeProcess) setDenyWarrantyClaim() {
	e, d := cp.GetDenyWarrantyClaim()
	tasks.SetTask(e, "Deny warranty claim", cp.DenyWarrantyClaimHash, cp.FromCheckIncomingClaimHash, cp.FromDenyWarrantyClaimHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: cp.DenyWarrantyClaimHash,
			B: canvas.Bounds{X: 580, Y: 120, Width: 100, Height: 80}})
}

/**
 * @Customer Support Process
 * @Sequence
 * @private fromCustomerSupportStartEvent
 * @private fromCheckIncomingClaim
 * @private fromDenyWarrantyClaim
 **/

// fromCustomerSupportStartEvent ...
func (cp *collaborativeProcess) fromCustomerSupportStartEvent() {
	// assign
	e, d := cp.GetFromCustomerSupportStartEvent()
	marker.SetSequenceFlow(e, "flow", "event", "activity", cp.FromCustomerSupportStartEventHash, cp.CustomerSupportStartEventHash, cp.CheckIncomingClaimHash)
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromCustomerSupportStartEventHash,
			W: []canvas.Waypoint{{X: 261, Y: 160}, {X: 320, Y: 160}}})
}

// fromCheckIncomingClaim ...
func (cp *collaborativeProcess) fromCheckIncomingClaim() {
	e, d := cp.GetFromCheckIncomingClaim()
	e.SetID("flow", cp.FromCheckIncomingClaimHash)
	e.SetName("decide")
	e.SetSourceRef("activity", cp.CheckIncomingClaimHash)
	e.SetTargetRef("activity", cp.DenyWarrantyClaimHash)
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
	e, d := cp.GetFromDenyWarrantyClaim()
	marker.SetSequenceFlow(e, "flow", "activity", "event", cp.FromDenyWarrantyClaimHash, cp.DenyWarrantyClaimHash, cp.CustomerSupportEndEventHash)
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
	e := cp.GetEdgeMessageClaim(cp.GetPlane())
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: e,
			T: "flow",
			H: cp.CustomerToCustomerSupportMessageHash,
			W: []canvas.Waypoint{{X: 370, Y: 400}, {X: 370, Y: 200}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
			E: e,
			B: canvas.Bounds{X: 387, Y: 290, Width: 26, Height: 14}})
}

// setEdgeMessageRefusal ...
func (cp *collaborativeProcess) setEdgeMessageRefusal() {
	e := cp.GetEdgeMessageRefusal(cp.GetPlane())
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: e,
			T: "flow",
			H: cp.CustomerSupportToCustomerMessageHash,
			W: []canvas.Waypoint{{X: 630, Y: 200}, {X: 630, Y: 400}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
			E: e,
			B: canvas.Bounds{X: 643, Y: 290, Width: 34, Height: 14}})
}

/**** Customer Process ****/

/**
 * @Customer Process I
 * @Process
 * @private setPoolCustomer
 * @private setProcessCustomer
 **/

// setPoolCustomer
func (cp *collaborativeProcess) setPoolCustomer() {
	e := cp.GetShapePoolCustomer(cp.GetPlane())
	canvas.SetPool(e, "id", true, cp.CustomerID,
		canvas.Bounds{X: 150, Y: 360, Width: 800, Height: 160})
}

// setProcessCustomer ...
func (cp *collaborativeProcess) setProcessCustomer() {
	cp.GetCustomerProcess().SetID("id", cp.CustomerProcessID)
	cp.GetCustomerProcess().SetIsExecutable(cp.CustomerIsExecutable)
}

/**
 * @Customer Process II
 * @Event
 * @SetCustomerStartEvent
 * @SetCustomerEndEvent
 **/

// setCustomerStartEvent ...
func (cp *collaborativeProcess) setCustomerStartEvent() {
	e, d := cp.GetCustomerStartEvent()
	elements.SetStartEvent(e, "Begin of Customer Process", cp.CustomerStartEventHash, cp.FromCustomerStartEventHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.CustomerStartEventHash,
			B: canvas.Bounds{X: 225, Y: 422, Width: 36, Height: 36}})
}

// setCustomerEndEvent ...
func (cp *collaborativeProcess) setCustomerEndEvent() {
	e, d := cp.GetCustomerEndEvent()
	elements.SetEndEvent(e, "End of Customer Process", cp.CustomerEndEventHash, cp.FromReceiptWarrantyRefusalHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.CustomerEndEventHash,
			B: canvas.Bounds{X: 822, Y: 422, Width: 36, Height: 36}})
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
	e, d := cp.GetNoticeOfDefect()
	tasks.SetTask(e, "Notice of defect", cp.NoticeOfDefectHash, cp.FromCustomerStartEventHash, cp.FromNoticeOfDefectHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: cp.NoticeOfDefectHash,
			B: canvas.Bounds{X: 320, Y: 400, Width: 100, Height: 80}})
}

// setWaitingForAnswer ...
func (cp *collaborativeProcess) setWaitingForAnswer() {
	e, d := cp.GetWaitingForAnswer()
	elements.SetIntermediateCatchEvent(e, "Waiting for answer", cp.WaitingForAnswerHash, cp.FromNoticeOfDefectHash, cp.FromWaitingForAnswerHash)
	t := cp.GetWaitingForAnswerTimerEventDefinition(e)
	definitions.SetTimerEventDefinition(t, "PT1M", cp.TimerEventDefinitionWaitingForAnswerHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "event",
			H: cp.WaitingForAnswerHash,
			B: canvas.Bounds{X: 482, Y: 422, Width: 36, Height: 36}})
}

// setReceiptWarrantyRefusal ...
func (cp *collaborativeProcess) setReceiptWarrantyRefusal() {
	e, d := cp.GetReceiptWarrantyRefusal()
	tasks.SetTask(e, "Receipt warranty refusal", cp.ReceiptWarrantyRefusalHash, cp.FromWaitingForAnswerHash, cp.FromReceiptWarrantyRefusalHash)
	canvas.SetShape(
		canvas.DelegateParameter{
			S: d,
			T: "activity",
			H: cp.ReceiptWarrantyRefusalHash,
			B: canvas.Bounds{X: 580, Y: 400, Width: 100, Height: 80}})
}

/**
 * @Customer Process IV
 * @Sequence
 * @private fromCustomerStartEvent
 * @private fromNoticeOfDefect
 * @private fromWaitingForAnswer
 * @private fromReceiptWarrantyRefusal
 **/

// FromCustomerStartEvent ...
func (cp *collaborativeProcess) fromCustomerStartEvent() {
	e, d := cp.GetFromCustomerStartEvent()
	marker.SetSequenceFlow(e, "flow", "event", "activity", cp.FromCustomerStartEventHash, cp.CustomerStartEventHash, cp.NoticeOfDefectHash)
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromCustomerStartEventHash,
			W: []canvas.Waypoint{{X: 261, Y: 440}, {X: 320, Y: 440}}})
}

// FromNoticeOfDefect ...
func (cp *collaborativeProcess) fromNoticeOfDefect() {
	e, d := cp.GetFromNoticeOfDefect()
	marker.SetSequenceFlow(e, "flow", "activity", "event", cp.FromNoticeOfDefectHash, cp.NoticeOfDefectHash, cp.WaitingForAnswerHash)
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromNoticeOfDefectHash,
			W: []canvas.Waypoint{{X: 420, Y: 440}, {X: 482, Y: 440}}})
}

// FromWaitingForAnswer ...
func (cp *collaborativeProcess) fromWaitingForAnswer() {
	e, d := cp.GetFromWaitingForAnswer()
	marker.SetSequenceFlow(e, "flow", "event", "activity", cp.FromWaitingForAnswerHash, cp.WaitingForAnswerHash, cp.ReceiptWarrantyRefusalHash)
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromWaitingForAnswerHash,
			W: []canvas.Waypoint{{X: 518, Y: 440}, {X: 580, Y: 440}}})
}

// FromReceiptWarrantyRefusal ...
func (cp *collaborativeProcess) fromReceiptWarrantyRefusal() {
	e, d := cp.GetFromReceiptWarrantyRefusal()
	marker.SetSequenceFlow(e, "flow", "activity", "event", cp.FromReceiptWarrantyRefusalHash, cp.ReceiptWarrantyRefusalHash, cp.CustomerEndEventHash)
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: d,
			T: "flow",
			H: cp.FromReceiptWarrantyRefusalHash,
			W: []canvas.Waypoint{{X: 680, Y: 440}, {X: 822, Y: 440}}})
}

// setDiagram ...
func (cp *collaborativeProcess) setDiagram() {
	// diagram attributes
	var n int64 = 1
	diagram := cp.GetDiagram()
	diagram.SetID("diagram", n)
	// plane attributes
	p := cp.GetPlane()
	p.SetID("plane", n)
	p.SetElement("id", cp.CollaborationID)
}

/**** Getter ****/

// GetCollaboration ...
func (cp collaborativeProcess) GetCollaboration() *pool.Collaboration {
	return cp.def.GetCollaboration()
}

// GetCustomerSupportParticipant ...
func (cp collaborativeProcess) GetCustomerSupportParticipant(e *pool.Collaboration) *pool.Participant {
	return e.GetParticipant(0)
}

// GetCustomerSupportProces ...
func (cp collaborativeProcess) GetCustomerSupportProcess() *process.Process {
	return cp.def.GetProcess(0)
}

// GetCustomerSupportStartEvent ...
func (cp collaborativeProcess) GetCustomerSupportStartEvent() (*elements.StartEvent, *canvas.Shape) {
	start := cp.GetCustomerSupportProcess().GetStartEvent(0)
	plane := cp.GetPlane()
	shape := cp.GetShapeCustomerSupportStartEvent(plane)
	return start, shape
}

// GetCustomerSupportEndEvent ...
func (cp collaborativeProcess) GetCustomerSupportEndEvent() (*elements.EndEvent, *canvas.Shape) {
	end := cp.GetCustomerSupportProcess().GetEndEvent(0)
	plane := cp.GetPlane()
	shape := cp.GetShapeCustomerSupportEndEvent(plane)
	return end, shape
}

// GetCheckIncomingClaim ...
func (cp collaborativeProcess) GetCheckIncomingClaim() (*tasks.Task, *canvas.Shape) {
	task := cp.GetCustomerSupportProcess().GetTask(0)
	plane := cp.GetPlane()
	shape := cp.GetShapeCheckIncomingClaim(plane)
	return task, shape
}

// GetDenyWarrantyClaim ...
func (cp collaborativeProcess) GetDenyWarrantyClaim() (*tasks.Task, *canvas.Shape) {
	task := cp.GetCustomerSupportProcess().GetTask(1)
	plane := cp.GetPlane()
	shape := cp.GetShapeDenyWarrantyClaim(plane)
	return task, shape
}

// GetFromCustomerSupportStartEvent ...
func (cp collaborativeProcess) GetFromCustomerSupportStartEvent() (*marker.SequenceFlow, *canvas.Edge) {
	flow := cp.GetCustomerSupportProcess().GetSequenceFlow(0)
	plane := cp.GetPlane()
	edge := cp.GetEdgeFromCustomerSupportStartEvent(plane)
	return flow, edge
}

// GetFromCheckIncomingClaim ...
func (cp collaborativeProcess) GetFromCheckIncomingClaim() (*marker.SequenceFlow, *canvas.Edge) {
	flow := cp.GetCustomerSupportProcess().GetSequenceFlow(1)
	plane := cp.GetPlane()
	edge := cp.GetEdgeFromCheckIncomingClaim(plane)
	return flow, edge
}

// GetFromDenyWarrantyClaim ...
func (cp collaborativeProcess) GetFromDenyWarrantyClaim() (*marker.SequenceFlow, *canvas.Edge) {
	flow := cp.GetCustomerSupportProcess().GetSequenceFlow(2)
	plane := cp.GetPlane()
	edge := cp.GetEdgeFromDenyWarrantyClaim(plane)
	return flow, edge
}

/*** CUSTOMER ***/

// GetCustomerProcess ...
func (cp collaborativeProcess) GetCustomerProcess() *process.Process {
	return cp.def.GetProcess(1)
}

// GetCustomerParticipant ...
func (cp collaborativeProcess) GetCustomerParticipant(e *pool.Collaboration) *pool.Participant {
	return e.GetParticipant(1)
}

// GetCustomerStartEvent ...
func (cp collaborativeProcess) GetCustomerStartEvent() (*elements.StartEvent, *canvas.Shape) {
	start := cp.GetCustomerProcess().GetStartEvent(0)
	plane := cp.GetPlane()
	shape := cp.GetShapeCustomerStartEvent(plane)
	return start, shape
}

// GetCustomerEndEvent ...
func (cp collaborativeProcess) GetCustomerEndEvent() (*elements.EndEvent, *canvas.Shape) {
	end := cp.GetCustomerProcess().GetEndEvent(0)
	plane := cp.GetPlane()
	shape := cp.GetShapeCustomerEndEvent(plane)
	return end, shape
}

// GetNoticeOfDefect ...
func (cp collaborativeProcess) GetNoticeOfDefect() (*tasks.Task, *canvas.Shape) {
	task := cp.GetCustomerProcess().GetTask(0)
	plane := cp.GetPlane()
	shape := cp.GetShapeNoticeOfDefect(plane)
	return task, shape
}

// GetWaitingForAnswer ...
func (cp collaborativeProcess) GetWaitingForAnswer() (*elements.IntermediateCatchEvent, *canvas.Shape) {
	timer := cp.GetCustomerProcess().GetIntermediateCatchEvent(0)
	plane := cp.GetPlane()
	shape := cp.GetShapeWaitingForAnswer(plane)
	return timer, shape
}

// GetWaitingForAnswerTimerEventDefinition ...
func (cp collaborativeProcess) GetWaitingForAnswerTimerEventDefinition(e *elements.IntermediateCatchEvent) *definitions.TimerEventDefinition {
	e.SetTimerEventDefinition()
	return e.GetTimerEventDefinition()
}

// GetReceiptWarrantyRefusal ...
func (cp collaborativeProcess) GetReceiptWarrantyRefusal() (*tasks.Task, *canvas.Shape) {
	task := cp.GetCustomerProcess().GetTask(1)
	plane := cp.GetPlane()
	shape := cp.GetShapeReceiptWarrantyRefusal(plane)
	return task, shape
}

// GetFromCustomerStartEvent ...
func (cp collaborativeProcess) GetFromCustomerStartEvent() (*marker.SequenceFlow, *canvas.Edge) {
	flow := cp.GetCustomerProcess().GetSequenceFlow(0)
	plane := cp.GetPlane()
	edge := cp.GetEdgeFromCustomerStartEvent(plane)
	return flow, edge
}

// GetFromNoticeOfDefect ...
func (cp collaborativeProcess) GetFromNoticeOfDefect() (*marker.SequenceFlow, *canvas.Edge) {
	flow := cp.GetCustomerProcess().GetSequenceFlow(1)
	plane := cp.GetPlane()
	edge := cp.GetEdgeFromNoticeOfDefect(plane)
	return flow, edge
}

// GetFromWaitingForAnswer ...
func (cp collaborativeProcess) GetFromWaitingForAnswer() (*marker.SequenceFlow, *canvas.Edge) {
	flow := cp.GetCustomerProcess().GetSequenceFlow(2)
	plane := cp.GetPlane()
	edge := cp.GetEdgeFromWaitingForAnswer(plane)
	return flow, edge
}

// GetFromReceiptWarrantyRefusal ...
func (cp collaborativeProcess) GetFromReceiptWarrantyRefusal() (*marker.SequenceFlow, *canvas.Edge) {
	flow := cp.GetCustomerProcess().GetSequenceFlow(3)
	plane := cp.GetPlane()
	edge := cp.GetEdgeFromReceiptWarrantyRefusal(plane)
	return flow, edge
}

/**** Messages ****/

// GetMessageClaim ...
func (cp collaborativeProcess) GetMessageClaim() *marker.MessageFlow {
	return cp.GetCollaboration().GetMessageFlow(0)
}

// GetMessageRefusal ...
func (cp collaborativeProcess) GetMessageRefusal() *marker.MessageFlow {
	return cp.GetCollaboration().GetMessageFlow(1)
}

/**** Diagram ****/

// GetDiagram ...
func (cp collaborativeProcess) GetDiagram() *canvas.Diagram {
	return cp.def.GetDiagram(0)
}

/**** Plane ****/

// GetPlane ...
func (cp collaborativeProcess) GetPlane() *canvas.Plane {
	return cp.GetDiagram().GetPlane()
}

/**** Shapes ****/

// GetShapePoolCustomerSupport ...
func (cp collaborativeProcess) GetShapePoolCustomerSupport(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(0)
}

// GetShapePoolCustomer ...
func (cp collaborativeProcess) GetShapePoolCustomer(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(1)
}

// GetShapeCustomerSupportStartEvent ...
func (cp collaborativeProcess) GetShapeCustomerSupportStartEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(2)
}

// GetShapeCustomerSupportEndEvent ...
func (cp collaborativeProcess) GetShapeCustomerSupportEndEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(3)
}

// GetShapeCheckIncomingClaim ...
func (cp collaborativeProcess) GetShapeCheckIncomingClaim(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(4)
}

// GetShapeDenyWarrantyClaim ...
func (cp collaborativeProcess) GetShapeDenyWarrantyClaim(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(5)
}

// GetShapeCustomerStartEvent ...
func (cp collaborativeProcess) GetShapeCustomerStartEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(6)
}

// GetShapeCustomerEndEvent ...
func (cp collaborativeProcess) GetShapeCustomerEndEvent(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(7)
}

// GetShapeNoticeOfDefect ...
func (cp collaborativeProcess) GetShapeNoticeOfDefect(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(8)
}

// GetShapeWaitingForAnswer ...
func (cp collaborativeProcess) GetShapeWaitingForAnswer(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(9)
}

// GetShapeReceiptWarrantyRefusal ...
func (cp collaborativeProcess) GetShapeReceiptWarrantyRefusal(e *canvas.Plane) *canvas.Shape {
	return e.GetShape(10)
}

/**** Edges ****/

// GetEdgeFromCustomerSupportStartEvent ...
func (cp collaborativeProcess) GetEdgeFromCustomerSupportStartEvent(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(0)
}

// GetEdgeFromCheckIncomingClaim ...
func (cp collaborativeProcess) GetEdgeFromCheckIncomingClaim(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(1)
}

// GetEdgeFromDenyWarrantyClaim ...
func (cp collaborativeProcess) GetEdgeFromDenyWarrantyClaim(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(2)
}

// GetEdgeFromCustomerStartEvent ...
func (cp collaborativeProcess) GetEdgeFromCustomerStartEvent(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(3)
}

// GetEdgeFromNoticeOfDefect ...
func (cp collaborativeProcess) GetEdgeFromNoticeOfDefect(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(4)
}

// GetEdgeFromWaitingForAnswer ...
func (cp collaborativeProcess) GetEdgeFromWaitingForAnswer(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(5)
}

// GetEdgeFromReceiptWarrantyRefusal ...
func (cp collaborativeProcess) GetEdgeFromReceiptWarrantyRefusal(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(6)
}

// GetEdgeMessageClaim ...
func (cp collaborativeProcess) GetEdgeMessageClaim(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(7)
}

// GetEdgeMessageRefusal ...
func (cp collaborativeProcess) GetEdgeMessageRefusal(e *canvas.Plane) *canvas.Edge {
	return e.GetEdge(8)
}
