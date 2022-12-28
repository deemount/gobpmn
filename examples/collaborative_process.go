package examples

/**************************************************************************************/
/**
 * @Import
 *
 * Import models and utils package to build the model
 **/

import (
	"log"

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
	cp.setMainElements()
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
 * @setMainElements
 * @@models.Definitions: SetCollaboration
 * @@models.Definitions: SetProcess
 * @@models.Definitions: SetDiagram
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

// setMainElements
func (cp *collaborativeProcess) setMainElements() {
	cp.def.SetCollaboration()
	cp.def.SetProcess(2)
	cp.def.SetDiagram(1)
}

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
	log.Printf("%+v", diagram)
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
	customerSupportParticipant.SetID("id", cp.CustomerSupportID)
	customerSupportParticipant.SetName("Customer Support")
	customerSupportParticipant.SetProcessRef("id", cp.CustomerSupportProcessID)
	// customer
	customerParticipant.SetID("id", cp.CustomerID)
	customerParticipant.SetName("Customer")
	customerParticipant.SetProcessRef("id", cp.CustomerProcessID)

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
	e.SetID("id", cp.CustomerSupportID)
	e.SetElement("id", cp.CustomerSupportID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.GetBounds().SetCoordinates(150, 80)
	e.GetBounds().SetSize(800, 160)
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
	e.SetID("event", cp.CustomerSupportStartEventHash)
	e.SetName("Begin of Customer Support Process")
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(cp.FromCustomerSupportStartEventHash)
	d.SetID("event", cp.CustomerSupportStartEventHash)
	d.SetElement("event", cp.CustomerSupportStartEventHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(225, 142)
	d.GetBounds().SetSize(36, 36)
}

// setCustomerSupportEndEvent ...
func (cp *collaborativeProcess) setCustomerSupportEndEvent() {
	e, d := cp.GetCustomerSupportEndEvent()
	e.SetID("event", cp.CustomerSupportEndEventHash)
	e.SetName("End of Customer Support Process")
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(cp.FromDenyWarrantyClaimHash)
	d.SetID("event", cp.CustomerSupportEndEventHash)
	d.SetElement("event", cp.CustomerSupportEndEventHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(822, 142)
	d.GetBounds().SetSize(36, 36)
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
	e.SetID("activity", cp.CheckIncomingClaimHash)
	e.SetName("Check incoming claim")
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(cp.FromCustomerSupportStartEventHash)
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(cp.FromCheckIncomingClaimHash)
	d.SetID("activity", cp.CheckIncomingClaimHash)
	d.SetElement("activity", cp.CheckIncomingClaimHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(320, 120)
	d.GetBounds().SetSize(100, 80)
}

// setDenyWarrantyClaim ...
func (cp *collaborativeProcess) setDenyWarrantyClaim() {
	e, d := cp.GetDenyWarrantyClaim()
	e.SetID("activity", cp.DenyWarrantyClaimHash)
	e.SetName("Deny warranty claim")
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(cp.FromCheckIncomingClaimHash)
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(cp.FromDenyWarrantyClaimHash)
	d.SetID("activity", cp.DenyWarrantyClaimHash)
	d.SetElement("activity", cp.DenyWarrantyClaimHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(580, 120)
	d.GetBounds().SetSize(100, 80)
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
	e.SetID("flow", cp.FromCustomerSupportStartEventHash)
	e.SetSourceRef("event", cp.CustomerSupportStartEventHash)
	e.SetTargetRef("activity", cp.CheckIncomingClaimHash)
	d.SetID("flow", cp.FromCustomerSupportStartEventHash)
	d.SetElement("flow", cp.FromCustomerSupportStartEventHash)
	d.SetWaypoint()
	d.GetWaypoint(0).SetCoordinates(261, 160)
	d.GetWaypoint(1).SetCoordinates(320, 160)
}

// fromCheckIncomingClaim ...
func (cp *collaborativeProcess) fromCheckIncomingClaim() {
	e, d := cp.GetFromCheckIncomingClaim()
	e.SetID("flow", cp.FromCheckIncomingClaimHash)
	e.SetName("decide")
	e.SetSourceRef("activity", cp.CheckIncomingClaimHash)
	e.SetTargetRef("activity", cp.DenyWarrantyClaimHash)
	d.SetID("flow", cp.FromCheckIncomingClaimHash)
	d.SetElement("flow", cp.FromCheckIncomingClaimHash)
	d.SetWaypoint()
	d.GetWaypoint(0).SetCoordinates(420, 160)
	d.GetWaypoint(1).SetCoordinates(580, 160)
	d.SetLabel()
	d.GetLabel().SetBounds()
	d.GetLabel().GetBounds().SetCoordinates(485, 142)
	d.GetLabel().GetBounds().SetSize(33, 14)
}

// fromDenyWarrantyClaim ...
func (cp *collaborativeProcess) fromDenyWarrantyClaim() {
	e, d := cp.GetFromDenyWarrantyClaim()
	e.SetID("flow", cp.FromDenyWarrantyClaimHash)
	e.SetSourceRef("activity", cp.DenyWarrantyClaimHash)
	e.SetTargetRef("event", cp.CustomerSupportEndEventHash)
	d.SetID("flow", cp.FromDenyWarrantyClaimHash)
	d.SetElement("flow", cp.FromDenyWarrantyClaimHash)
	d.SetWaypoint()
	d.GetWaypoint(0).SetCoordinates(680, 160)
	d.GetWaypoint(1).SetCoordinates(822, 160)
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
	e.SetID("flow", cp.CustomerToCustomerSupportMessageHash)
	e.SetElement("flow", cp.CustomerToCustomerSupportMessageHash)
	e.SetWaypoint()
	e.GetWaypoint(0).SetCoordinates(370, 400)
	e.GetWaypoint(1).SetCoordinates(370, 200)
	e.SetLabel()
	e.GetLabel().SetBounds()
	e.GetLabel().GetBounds().SetCoordinates(387, 290)
	e.GetLabel().GetBounds().SetSize(26, 14)
}

// setEdgeMessageRefusal ...
func (cp *collaborativeProcess) setEdgeMessageRefusal() {
	e := cp.GetEdgeMessageRefusal(cp.GetPlane())
	e.SetID("flow", cp.CustomerSupportToCustomerMessageHash)
	e.SetElement("flow", cp.CustomerSupportToCustomerMessageHash)
	e.SetWaypoint()
	e.GetWaypoint(0).SetCoordinates(630, 200)
	e.GetWaypoint(1).SetCoordinates(630, 400)
	e.SetLabel()
	e.GetLabel().SetBounds()
	e.GetLabel().GetBounds().SetCoordinates(643, 290)
	e.GetLabel().GetBounds().SetSize(34, 14)
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
	e.SetID("id", cp.CustomerID)
	e.SetElement("id", cp.CustomerID)
	e.SetIsHorizontal(true)
	e.SetBounds()
	e.GetBounds().SetCoordinates(150, 360)
	e.GetBounds().SetSize(800, 160)
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
	e.SetID("event", cp.CustomerStartEventHash)
	e.SetName("Begin of Customer Process")
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(cp.FromCustomerStartEventHash)
	d.SetID("event", cp.CustomerStartEventHash)
	d.SetElement("event", cp.CustomerStartEventHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(225, 422)
	d.GetBounds().SetSize(36, 36)
}

// setCustomerEndEvent ...
func (cp *collaborativeProcess) setCustomerEndEvent() {
	e, d := cp.GetCustomerEndEvent()
	e.SetID("event", cp.CustomerEndEventHash)
	e.SetName("End of Customer Process")
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(cp.FromReceiptWarrantyRefusalHash)
	d.SetID("event", cp.CustomerEndEventHash)
	d.SetElement("event", cp.CustomerEndEventHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(822, 422)
	d.GetBounds().SetSize(36, 36)
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
	e.SetID("activity", cp.NoticeOfDefectHash)
	e.SetName("Notice of defect")
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(cp.FromCustomerStartEventHash)
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(cp.FromNoticeOfDefectHash)
	d.SetID("activity", cp.NoticeOfDefectHash)
	d.SetElement("activity", cp.NoticeOfDefectHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(320, 400)
	d.GetBounds().SetSize(100, 80)
}

// setWaitingForAnswer ...
func (cp *collaborativeProcess) setWaitingForAnswer() {
	e, d := cp.GetWaitingForAnswer()
	e.SetID("event", cp.WaitingForAnswerHash)
	e.SetName("Waiting for answer")
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(cp.FromNoticeOfDefectHash)
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(cp.FromWaitingForAnswerHash)
	t := cp.GetWaitingForAnswerTimerEventDefinition(e)
	t.SetID("ted", cp.TimerEventDefinitionWaitingForAnswerHash)
	t.SetTimeDuration()
	t.GetTimeDuration().SetTimerDefinitionType()
	t.GetTimeDuration().SetTimerDefinition("PT1M")
	d.SetID("event", cp.WaitingForAnswerShapeIDHash)
	d.SetElement("event", cp.WaitingForAnswerHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(482, 422)
	d.GetBounds().SetSize(36, 36)
}

// setReceiptWarrantyRefusal ...
func (cp *collaborativeProcess) setReceiptWarrantyRefusal() {
	e, d := cp.GetReceiptWarrantyRefusal()
	e.SetID("activity", cp.ReceiptWarrantyRefusalHash)
	e.SetName("Receipt warranty refusal")
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(cp.FromWaitingForAnswerHash)
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(cp.FromReceiptWarrantyRefusalHash)
	d.SetID("activity", cp.ReceiptWarrantyRefusalHash)
	d.SetElement("activity", cp.ReceiptWarrantyRefusalHash)
	d.SetBounds()
	d.GetBounds().SetCoordinates(580, 400)
	d.GetBounds().SetSize(100, 80)
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
	e.SetID("flow", cp.FromCustomerStartEventHash)
	e.SetSourceRef("event", cp.CustomerStartEventHash)
	e.SetTargetRef("activity", cp.NoticeOfDefectHash)
	d.SetID("flow", cp.FromCustomerStartEventHash)
	d.SetElement("flow", cp.FromCustomerStartEventHash)
	d.SetWaypoint()
	d.GetWaypoint(0).SetCoordinates(261, 440)
	d.GetWaypoint(1).SetCoordinates(320, 440)
}

// FromNoticeOfDefect ...
func (cp *collaborativeProcess) fromNoticeOfDefect() {
	e, d := cp.GetFromNoticeOfDefect()
	e.SetID("flow", cp.FromNoticeOfDefectHash)
	e.SetSourceRef("activity", cp.NoticeOfDefectHash)
	e.SetTargetRef("event", cp.WaitingForAnswerHash)
	d.SetID("flow", cp.FromNoticeOfDefectHash)
	d.SetElement("flow", cp.FromNoticeOfDefectHash)
	d.SetWaypoint()
	d.GetWaypoint(0).SetCoordinates(420, 440)
	d.GetWaypoint(1).SetCoordinates(482, 440)
}

// FromWaitingForAnswer ...
func (cp *collaborativeProcess) fromWaitingForAnswer() {
	e, d := cp.GetFromWaitingForAnswer()
	e.SetID("flow", cp.FromWaitingForAnswerHash)
	e.SetSourceRef("event", cp.WaitingForAnswerHash)
	e.SetTargetRef("activity", cp.ReceiptWarrantyRefusalHash)
	d.SetID("flow", cp.FromWaitingForAnswerHash)
	d.SetElement("flow", cp.FromWaitingForAnswerHash)
	d.SetWaypoint()
	d.GetWaypoint(0).SetCoordinates(518, 440)
	d.GetWaypoint(1).SetCoordinates(580, 440)
}

// FromReceiptWarrantyRefusal ...
func (cp *collaborativeProcess) fromReceiptWarrantyRefusal() {
	e, d := cp.GetFromReceiptWarrantyRefusal()
	e.SetID("flow", cp.FromReceiptWarrantyRefusalHash)
	e.SetSourceRef("activity", cp.ReceiptWarrantyRefusalHash)
	e.SetTargetRef("event", cp.CustomerEndEventHash)
	d.SetID("flow", cp.FromReceiptWarrantyRefusalHash)
	d.SetElement("flow", cp.FromReceiptWarrantyRefusalHash)
	d.SetWaypoint()
	d.GetWaypoint(0).SetCoordinates(680, 440)
	d.GetWaypoint(1).SetCoordinates(822, 440)
}

// setDiagram ...
func (cp *collaborativeProcess) setDiagram() {
	// diagram attributes
	var n int64 = 1
	diagram := cp.GetDiagram()
	log.Printf("%+v", diagram)
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
