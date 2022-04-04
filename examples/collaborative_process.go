package examples

import (
	"fmt"

	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

// CollaborativeProcessRepository ...
type CollaborativeProcessRepository interface {
	// Setter
	Create()
	SetElements()
	SetDefinitionsAttributes()
	SetCollaboration()
	// Customer Support
	SetCustomerSupportProcess()
	SetCustomerSupportStartEvent()
	SetCustomerSupportTaskCheckIncomingClaim()
	SetCustomerSupportTaskDenyWarrantyClaim()
	SetFromCustomerSupportCheckIncomingClaimSequenceFlow()
	// Customer
	SetCustomerProcess()
	SetCustomerStartEvent()
	SetCustomerTaskNoticeOfDefect()
	SetCustomerIntermediateCatchEventWaitingForAnswer()
	SetFromCustomerStartEventSequenceFlow()
	// Diagram
	SetDiagram()
	// Getters
	GetPlane() *models.Plane
	GetCustomerSupportProcess() *models.Process
	GetCustomerProcess() *models.Process
}

// CollaborativeProcess ...
type CollaborativeProcess struct {
	def                         *models.Definitions
	CustomerSupportIsExecutable bool
	CustomerIsExecutable        bool
	CollaborationID             string
	CustomerSupportID           string
	CustomerID                  string
	CustomerSupportProcessID    string
	CustomerProcessID           string
	// Message Flow
	CustomerToCustomerSupportMessageHash string
	// Customer Support
	CustomerSupportStartEventHash             string
	FromCustomerSupportStartEvent             string
	CustomerSupportTaskCheckIncomingClaimHash string
	FromCustomerSupportTaskCheckIncomingClaim string
	CustomerSupportTaskDenyWarrantyClaimHash  string
	// Customer
	CustomerStartEventHash                                    string
	FromCustomerStartEvent                                    string
	CustomerTaskNoticeOfDefectHash                            string
	FromCustomerTaskNoticeOfDefect                            string
	CustomerIntermediateCatchEventWaitingForAnswerHash        string
	CustomerIntermediateCatchEventWaitingForAnswerShapeIDHash string
	CustomerTimerEventDefinitionWaitingForAnswerHash          string
	FromCustomerIntermediateCatchEventWaitingForAnswer        string
	CustomerTaskReceiptWarrantyRefusalHash                    string
	FromCustomerTaskReceiptWarrantyRefusal                    string
}

//
func NewCollaborativeProcess(def *models.Definitions) CollaborativeProcess {
	return CollaborativeProcess{
		def:                                  def,
		CustomerSupportIsExecutable:          true,
		CustomerIsExecutable:                 false,
		CollaborationID:                      "uniqueCollaborationId",
		CustomerSupportID:                    "uniqueCustomerSupportId",
		CustomerID:                           "uniqueCustomerId",
		CustomerSupportProcessID:             "uniqueCustomerSupportProcessId",
		CustomerProcessID:                    "uniqueCustomerProcessId",
		CustomerToCustomerSupportMessageHash: utils.GenerateHash(),
		// Customer Support
		CustomerSupportStartEventHash:             utils.GenerateHash(),
		FromCustomerSupportStartEvent:             utils.GenerateHash(),
		CustomerSupportTaskCheckIncomingClaimHash: utils.GenerateHash(),
		FromCustomerSupportTaskCheckIncomingClaim: utils.GenerateHash(),
		CustomerSupportTaskDenyWarrantyClaimHash:  utils.GenerateHash(),
		// Customer
		CustomerStartEventHash:                                    utils.GenerateHash(),
		FromCustomerStartEvent:                                    utils.GenerateHash(),
		CustomerTaskNoticeOfDefectHash:                            utils.GenerateHash(),
		FromCustomerTaskNoticeOfDefect:                            utils.GenerateHash(),
		CustomerIntermediateCatchEventWaitingForAnswerHash:        utils.GenerateHash(),
		CustomerIntermediateCatchEventWaitingForAnswerShapeIDHash: utils.GenerateHash(),
		CustomerTimerEventDefinitionWaitingForAnswerHash:          utils.GenerateHash(),
		FromCustomerIntermediateCatchEventWaitingForAnswer:        utils.GenerateHash(),
		CustomerTaskReceiptWarrantyRefusalHash:                    utils.GenerateHash(),
		FromCustomerTaskReceiptWarrantyRefusal:                    utils.GenerateHash(),
	}
}

// Create ...
func (collaborativeProcess *CollaborativeProcess) Create() {
	collaborativeProcess.SetElements()
	collaborativeProcess.SetDefinitionsAttributes()
	collaborativeProcess.SetCollaboration()
	collaborativeProcess.SetCustomerSupportProcess()
	collaborativeProcess.SetCustomerProcess()
	// Customer Support
	collaborativeProcess.SetCustomerSupportStartEvent()
	collaborativeProcess.SetCustomerSupportTaskCheckIncomingClaim()
	collaborativeProcess.SetCustomerSupportTaskDenyWarrantyClaim()
	collaborativeProcess.SetFromCustomerSupportStartEventSequenceFlow()
	collaborativeProcess.SetFromCustomerSupportCheckIncomingClaimSequenceFlow()
	// Customer
	collaborativeProcess.SetCustomerStartEvent()
	collaborativeProcess.SetCustomerTaskNoticeOfDefect()
	collaborativeProcess.SetCustomerIntermediateCatchEventWaitingForAnswer()
	collaborativeProcess.SetCustomerTaskReceiptWarrantyRefusal()
	collaborativeProcess.SetFromCustomerStartEventSequenceFlow()
	collaborativeProcess.SetFromCustomerTaskNoticeOfDefectSequenceFlow()
	collaborativeProcess.SetFromCustomerIntermediateCatchEventWaitingForAnswerSequenceFlow()
	// Diagram
	collaborativeProcess.SetDiagram()
}

/**
 *
 * Setter
 *
 **/

// SetElements ...
func (collaborativeProcess *CollaborativeProcess) SetElements() {
	// Collaboration
	collaborativeProcess.def.SetCollaboration()
	collaborativeProcess.def.Collaboration[0].SetParticipant(2)
	collaborativeProcess.def.Collaboration[0].SetMessageFlow(1)
	collaborativeProcess.def.SetProcess(2)
	// Customer Support
	customerSupport := collaborativeProcess.GetCustomerSupportProcess()
	customerSupport.SetStartEvent(1)
	customerSupport.SetTask(2)
	customerSupport.SetSequenceFlow(2)
	// Customer
	customer := collaborativeProcess.GetCustomerProcess()
	customer.SetStartEvent(1)
	customer.SetTask(2)
	customer.SetIntermediateCatchEvent(1)
	customer.SetSequenceFlow(3)
	// Diagram
	collaborativeProcess.def.SetDiagram(1)
	collaborativeProcess.def.Diagram[0].SetPlane()
	plane := collaborativeProcess.GetPlane()
	plane.SetShape(9)
	plane.SetEdge(6)
}

// SetDefinitionsAttributes ...
func (collaborativeProcess *CollaborativeProcess) SetDefinitionsAttributes() {
	collaborativeProcess.def.SetDefaultAttributes()
}

// SetCollaboration ...
func (collaborativeProcess *CollaborativeProcess) SetCollaboration() {
	// assign
	plane := collaborativeProcess.GetPlane()
	customerSupport := collaborativeProcess.GetCustomerSupportParticipant()

	// generics
	collaborativeProcess.def.Collaboration[0].SetID("id", collaborativeProcess.CollaborationID)
	// participant attributes
	// generics
	// customer support
	customerSupport.SetID("id", collaborativeProcess.CustomerSupportID)
	customerSupport.SetName("Customer Support")
	customerSupport.SetProcessRef("id", collaborativeProcess.CustomerSupportProcessID)
	// customer
	collaborativeProcess.def.Collaboration[0].Participant[1].SetID("id", collaborativeProcess.CustomerID)
	collaborativeProcess.def.Collaboration[0].Participant[1].SetName("Customer")
	collaborativeProcess.def.Collaboration[0].Participant[1].SetProcessRef("id", collaborativeProcess.CustomerProcessID)
	// message flow
	collaborativeProcess.def.Collaboration[0].MessageFlow[0].SetID("flow", collaborativeProcess.CustomerToCustomerSupportMessageHash)
	collaborativeProcess.def.Collaboration[0].MessageFlow[0].SetName("Submit warranty claim")
	collaborativeProcess.def.Collaboration[0].MessageFlow[0].SetSourceRef(fmt.Sprintf("Activity_%s", collaborativeProcess.CustomerTaskNoticeOfDefectHash))
	collaborativeProcess.def.Collaboration[0].MessageFlow[0].SetTargetRef(fmt.Sprintf("Activity_%s", collaborativeProcess.CustomerSupportTaskCheckIncomingClaimHash))
	// shape attributes
	// customer support
	plane.Shape[0].SetID("id", collaborativeProcess.CustomerSupportID)
	plane.Shape[0].SetElement("id", collaborativeProcess.CustomerSupportID)
	plane.Shape[0].SetIsHorizontal(true)
	plane.Shape[0].SetBounds()
	plane.Shape[0].Bounds[0].SetCoordinates(150, 80)
	plane.Shape[0].Bounds[0].SetSize(800, 160)
	// customer
	plane.Shape[1].SetID("id", collaborativeProcess.CustomerID)
	plane.Shape[1].SetElement("id", collaborativeProcess.CustomerID)
	plane.Shape[1].SetIsHorizontal(true)
	plane.Shape[1].SetBounds()
	plane.Shape[1].Bounds[0].SetCoordinates(150, 360)
	plane.Shape[1].Bounds[0].SetSize(800, 160)
	// edge attributes
	plane.Edge[5].SetID("flow", collaborativeProcess.CustomerToCustomerSupportMessageHash)
	plane.Edge[5].SetElement("flow", collaborativeProcess.CustomerToCustomerSupportMessageHash)
	plane.Edge[5].SetWaypoint()
	plane.Edge[5].Waypoint[0].SetCoordinates(370, 400)
	plane.Edge[5].Waypoint[1].SetCoordinates(370, 200)
	plane.Edge[5].SetLabel()
	plane.Edge[5].Label[0].SetBounds()
	plane.Edge[5].Label[0].Bounds[0].SetCoordinates(375, 290)
	plane.Edge[5].Label[0].Bounds[0].SetSize(89, 27)
}

/**
 *
 * Customer Support Process
 *
 **/

// SetCustomerSupportProcess ...
func (collaborativeProcess *CollaborativeProcess) SetCustomerSupportProcess() {
	// assign
	process := collaborativeProcess.GetCustomerSupportProcess()
	// generics
	process.SetID("id", collaborativeProcess.CustomerSupportProcessID)
	process.SetIsExecutable(collaborativeProcess.CustomerSupportIsExecutable)
}

// SetCustomerSupportStartEvent ...
func (collaborativeProcess *CollaborativeProcess) SetCustomerSupportStartEvent() {
	// assign
	process := collaborativeProcess.GetCustomerSupportProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.StartEvent[0].SetID("event", collaborativeProcess.CustomerSupportStartEventHash)
	process.StartEvent[0].SetName("Begin of Customer Support Process")
	// outgoing
	process.StartEvent[0].SetOutgoing(1)
	process.StartEvent[0].Outgoing[0].SetFlow(collaborativeProcess.FromCustomerSupportStartEvent)
	// shape attributes
	plane.Shape[2].SetID("event", collaborativeProcess.CustomerSupportStartEventHash)
	plane.Shape[2].SetElement("event", collaborativeProcess.CustomerSupportStartEventHash)
	plane.Shape[2].SetBounds()
	plane.Shape[2].Bounds[0].SetCoordinates(225, 142)
	plane.Shape[2].Bounds[0].SetSize(36, 36)
}

func (collaborativeProcess *CollaborativeProcess) SetCustomerSupportTaskCheckIncomingClaim() {
	// assign
	process := collaborativeProcess.GetCustomerSupportProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.Task[0].SetID(collaborativeProcess.CustomerSupportTaskCheckIncomingClaimHash)
	process.Task[0].SetName("Check incoming claim")
	// incoming
	process.Task[0].SetIncoming(1)
	process.Task[0].Incoming[0].SetFlow(collaborativeProcess.FromCustomerSupportStartEvent)
	// outgoing
	process.Task[0].SetOutgoing(1)
	process.Task[0].Outgoing[0].SetFlow(collaborativeProcess.FromCustomerSupportTaskCheckIncomingClaim)
	// shape attributes
	plane.Shape[3].SetID("activity", collaborativeProcess.CustomerSupportTaskCheckIncomingClaimHash)
	plane.Shape[3].SetElement("activity", collaborativeProcess.CustomerSupportTaskCheckIncomingClaimHash)
	plane.Shape[3].SetBounds()
	plane.Shape[3].Bounds[0].SetCoordinates(320, 120)
	plane.Shape[3].Bounds[0].SetSize(100, 80)
}

//
func (collaborativeProcess *CollaborativeProcess) SetCustomerSupportTaskDenyWarrantyClaim() {
	// assign
	process := collaborativeProcess.GetCustomerSupportProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.Task[1].SetID(collaborativeProcess.CustomerSupportTaskDenyWarrantyClaimHash)
	process.Task[1].SetName("Deny warranty claim")
	// incoming
	process.Task[1].SetIncoming(1)
	process.Task[1].Incoming[0].SetFlow(collaborativeProcess.FromCustomerSupportTaskCheckIncomingClaim)
	// outgoing
	//collaborativeProcess.def.Process[0].Task[0].SetOutgoing(1)
	//collaborativeProcess.def.Process[0].Task[0].Outgoing[0].SetFlow(collaborativeProcess.FromCustomerSupportCheckIncomingClaimTask)
	// shape attributes
	plane.Shape[4].SetID("activity", collaborativeProcess.CustomerSupportTaskDenyWarrantyClaimHash)
	plane.Shape[4].SetElement("activity", collaborativeProcess.CustomerSupportTaskDenyWarrantyClaimHash)
	plane.Shape[4].SetBounds()
	plane.Shape[4].Bounds[0].SetCoordinates(580, 120)
	plane.Shape[4].Bounds[0].SetSize(100, 80)
}

// SetFromCustomerSupportStartEventSequenceFlow ...
func (collaborativeProcess *CollaborativeProcess) SetFromCustomerSupportStartEventSequenceFlow() {
	// assign
	process := collaborativeProcess.GetCustomerSupportProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.SequenceFlow[0].SetID(collaborativeProcess.FromCustomerSupportStartEvent)
	process.SequenceFlow[0].SetSourceRef(fmt.Sprintf("Event_%s", collaborativeProcess.CustomerSupportStartEventHash))
	process.SequenceFlow[0].SetTargetRef(fmt.Sprintf("Activity_%s", collaborativeProcess.CustomerSupportTaskCheckIncomingClaimHash))
	// edge attributes
	plane.Edge[0].SetID("flow", collaborativeProcess.FromCustomerSupportStartEvent)
	plane.Edge[0].SetElement("flow", collaborativeProcess.FromCustomerSupportStartEvent)
	plane.Edge[0].SetWaypoint()
	plane.Edge[0].Waypoint[0].SetCoordinates(261, 160)
	plane.Edge[0].Waypoint[1].SetCoordinates(320, 160)
}

// SetFromCustomerSupportCheckIncomingClaimSequenceFlow ...
func (collaborativeProcess *CollaborativeProcess) SetFromCustomerSupportCheckIncomingClaimSequenceFlow() {
	// assign
	process := collaborativeProcess.GetCustomerSupportProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.SequenceFlow[1].SetID(collaborativeProcess.FromCustomerSupportTaskCheckIncomingClaim)
	process.SequenceFlow[1].SetSourceRef(fmt.Sprintf("Activity_%s", collaborativeProcess.CustomerSupportTaskCheckIncomingClaimHash))
	process.SequenceFlow[1].SetTargetRef(fmt.Sprintf("Activity_%s", collaborativeProcess.CustomerSupportTaskDenyWarrantyClaimHash))
	// edge attributes
	plane.Edge[1].SetID("flow", collaborativeProcess.FromCustomerSupportTaskCheckIncomingClaim)
	plane.Edge[1].SetElement("flow", collaborativeProcess.FromCustomerSupportTaskCheckIncomingClaim)
	plane.Edge[1].SetWaypoint()
	plane.Edge[1].Waypoint[0].SetCoordinates(420, 160)
	plane.Edge[1].Waypoint[1].SetCoordinates(580, 160)
}

/**
 *
 * Customer Process
 *
 **/

// SetCustomerProcess ...
func (collaborativeProcess *CollaborativeProcess) SetCustomerProcess() {
	// assign
	process := collaborativeProcess.GetCustomerProcess()
	// generics
	process.SetID("id", collaborativeProcess.CustomerProcessID)
	process.SetIsExecutable(collaborativeProcess.CustomerIsExecutable)
}

// SetCustomerStartEvent ...
func (collaborativeProcess *CollaborativeProcess) SetCustomerStartEvent() {
	// assign
	process := collaborativeProcess.GetCustomerProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.StartEvent[0].SetID("event", collaborativeProcess.CustomerStartEventHash)
	process.StartEvent[0].SetName("Begin of Customer Process")
	// outgoing
	process.StartEvent[0].SetOutgoing(1)
	process.StartEvent[0].Outgoing[0].SetFlow(collaborativeProcess.FromCustomerStartEvent)
	// shape attributes
	plane.Shape[5].SetID("event", collaborativeProcess.CustomerStartEventHash)
	plane.Shape[5].SetElement("event", collaborativeProcess.CustomerStartEventHash)
	plane.Shape[5].SetBounds()
	plane.Shape[5].Bounds[0].SetCoordinates(225, 422)
	plane.Shape[5].Bounds[0].SetSize(36, 36)
}

// SetCustomerTaskNoticeOfDefect ...
func (collaborativeProcess *CollaborativeProcess) SetCustomerTaskNoticeOfDefect() {
	// assign
	process := collaborativeProcess.GetCustomerProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.Task[0].SetID(collaborativeProcess.CustomerTaskNoticeOfDefectHash)
	process.Task[0].SetName("Notice of defect")
	// incoming
	process.Task[0].SetIncoming(1)
	process.Task[0].Incoming[0].SetFlow(collaborativeProcess.FromCustomerStartEvent)
	// outgoing
	process.Task[0].SetOutgoing(1)
	process.Task[0].Outgoing[0].SetFlow(collaborativeProcess.FromCustomerTaskNoticeOfDefect)
	// shape attributes
	plane.Shape[6].SetID("activity", collaborativeProcess.CustomerTaskNoticeOfDefectHash)
	plane.Shape[6].SetElement("activity", collaborativeProcess.CustomerTaskNoticeOfDefectHash)
	plane.Shape[6].SetBounds()
	plane.Shape[6].Bounds[0].SetCoordinates(320, 400)
	plane.Shape[6].Bounds[0].SetSize(100, 80)
}

// SetCustomerIntermediateCatchEventWaitingForAnswer ...
func (collaborativeProcess *CollaborativeProcess) SetCustomerIntermediateCatchEventWaitingForAnswer() {
	// assign
	process := collaborativeProcess.GetCustomerProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.IntermediateCatchEvent[0].SetID(collaborativeProcess.CustomerIntermediateCatchEventWaitingForAnswerHash)
	process.IntermediateCatchEvent[0].SetName("Waiting for answer")
	// incoming
	process.IntermediateCatchEvent[0].SetIncoming(1)
	process.IntermediateCatchEvent[0].Incoming[0].SetFlow(collaborativeProcess.FromCustomerTaskNoticeOfDefect)
	// outgoing
	process.IntermediateCatchEvent[0].SetOutgoing(1)
	process.IntermediateCatchEvent[0].Outgoing[0].SetFlow(collaborativeProcess.FromCustomerIntermediateCatchEventWaitingForAnswer)
	// timer event definition
	process.IntermediateCatchEvent[0].SetTimerEventDefinition()
	process.IntermediateCatchEvent[0].TimerEventDefinition[0].SetID(collaborativeProcess.CustomerTimerEventDefinitionWaitingForAnswerHash)
	process.IntermediateCatchEvent[0].TimerEventDefinition[0].SetTimeDuration()
	process.IntermediateCatchEvent[0].TimerEventDefinition[0].TimeDuration[0].SetTimerDefinitionType()
	process.IntermediateCatchEvent[0].TimerEventDefinition[0].TimeDuration[0].SetTimerDefinition("PT1")
	// shape attributes
	plane.Shape[7].SetID("event", collaborativeProcess.CustomerIntermediateCatchEventWaitingForAnswerShapeIDHash)
	plane.Shape[7].SetElement("event", collaborativeProcess.CustomerIntermediateCatchEventWaitingForAnswerHash)
	plane.Shape[7].SetBounds()
	plane.Shape[7].Bounds[0].SetCoordinates(482, 422)
	plane.Shape[7].Bounds[0].SetSize(36, 36)
}

// SetCustomerTaskReceiptWarrantyRefusal ...
func (collaborativeProcess *CollaborativeProcess) SetCustomerTaskReceiptWarrantyRefusal() {
	// assign
	process := collaborativeProcess.GetCustomerProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.Task[1].SetID(collaborativeProcess.CustomerTaskReceiptWarrantyRefusalHash)
	process.Task[1].SetName("Receipt warranty refusal")
	// incoming
	process.Task[1].SetIncoming(1)
	process.Task[1].Incoming[0].SetFlow(collaborativeProcess.FromCustomerIntermediateCatchEventWaitingForAnswer)
	// outgoing

	// shape attributes
	plane.Shape[8].SetID("activity", collaborativeProcess.CustomerTaskReceiptWarrantyRefusalHash)
	plane.Shape[8].SetElement("activity", collaborativeProcess.CustomerTaskReceiptWarrantyRefusalHash)
	plane.Shape[8].SetBounds()
	plane.Shape[8].Bounds[0].SetCoordinates(580, 400)
	plane.Shape[8].Bounds[0].SetSize(100, 80)
}

// SetFromStartEventSequenceFlow ...
func (collaborativeProcess *CollaborativeProcess) SetFromCustomerStartEventSequenceFlow() {
	// assign
	process := collaborativeProcess.GetCustomerProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.SequenceFlow[0].SetID(collaborativeProcess.FromCustomerStartEvent)
	process.SequenceFlow[0].SetSourceRef(fmt.Sprintf("Event_%s", collaborativeProcess.CustomerStartEventHash))
	process.SequenceFlow[0].SetTargetRef(fmt.Sprintf("Activity_%s", collaborativeProcess.CustomerTaskNoticeOfDefectHash))
	// edge attributes
	plane.Edge[2].SetID("flow", collaborativeProcess.FromCustomerStartEvent)
	plane.Edge[2].SetElement("flow", collaborativeProcess.FromCustomerStartEvent)
	plane.Edge[2].SetWaypoint()
	plane.Edge[2].Waypoint[0].SetCoordinates(261, 440)
	plane.Edge[2].Waypoint[1].SetCoordinates(320, 440)
}

// SetFromStartEventSequenceFlow ...
func (collaborativeProcess *CollaborativeProcess) SetFromCustomerTaskNoticeOfDefectSequenceFlow() {
	// assign
	process := collaborativeProcess.GetCustomerProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.SequenceFlow[1].SetID(collaborativeProcess.FromCustomerTaskNoticeOfDefect)
	process.SequenceFlow[1].SetSourceRef(fmt.Sprintf("Activity_%s", collaborativeProcess.CustomerTaskNoticeOfDefectHash))
	process.SequenceFlow[1].SetTargetRef(fmt.Sprintf("Event_%s", collaborativeProcess.CustomerIntermediateCatchEventWaitingForAnswerHash))
	// edge attributes
	plane.Edge[3].SetID("flow", collaborativeProcess.FromCustomerTaskNoticeOfDefect)
	plane.Edge[3].SetElement("flow", collaborativeProcess.FromCustomerTaskNoticeOfDefect)
	plane.Edge[3].SetWaypoint()
	plane.Edge[3].Waypoint[0].SetCoordinates(420, 440)
	plane.Edge[3].Waypoint[1].SetCoordinates(482, 440)
}

// SetFromCustomerIntermediateCatchEventWaitingForAnswerSequenceFlow ...
func (collaborativeProcess *CollaborativeProcess) SetFromCustomerIntermediateCatchEventWaitingForAnswerSequenceFlow() {
	// assign
	process := collaborativeProcess.GetCustomerProcess()
	plane := collaborativeProcess.GetPlane()
	// generics
	process.SequenceFlow[2].SetID(collaborativeProcess.FromCustomerTaskReceiptWarrantyRefusal)
	process.SequenceFlow[2].SetSourceRef(fmt.Sprintf("Event_%s", collaborativeProcess.CustomerIntermediateCatchEventWaitingForAnswerHash))
	process.SequenceFlow[2].SetTargetRef(fmt.Sprintf("Activity_%s", collaborativeProcess.CustomerTaskReceiptWarrantyRefusalHash))
	// edge attributes
	plane.Edge[4].SetID("flow", collaborativeProcess.FromCustomerIntermediateCatchEventWaitingForAnswer)
	plane.Edge[4].SetElement("flow", collaborativeProcess.FromCustomerIntermediateCatchEventWaitingForAnswer)
	plane.Edge[4].SetWaypoint()
	plane.Edge[4].Waypoint[0].SetCoordinates(518, 440)
	plane.Edge[4].Waypoint[1].SetCoordinates(580, 440)
}

// SetDiagram ...
func (collaborativeProcess *CollaborativeProcess) SetDiagram() {
	// diagram attributes
	var n int64 = 1
	collaborativeProcess.def.Diagram[0].SetID(n)
	// plane attributes
	p := collaborativeProcess.GetPlane()
	p.SetID(n)
	p.SetElement("id", collaborativeProcess.CollaborationID)
}

/**
 *
 * Getter
 *
 **/

// GetCustomerSupportProces ...
func (collaborativeProcess CollaborativeProcess) GetCustomerSupportProcess() *models.Process {
	return &collaborativeProcess.def.Process[0]
}

// GetCustomerSupportParticipant
func (collaborativeProcess CollaborativeProcess) GetCustomerSupportParticipant() *models.Participant {
	return &collaborativeProcess.def.Collaboration[0].Participant[0]
}

// GetCustomerProcess ...
func (collaborativeProcess CollaborativeProcess) GetCustomerProcess() *models.Process {
	return &collaborativeProcess.def.Process[1]
}

// GetCustomerSupportParticipant
func (collaborativeProcess CollaborativeProcess) GetCustomerParticipant() *models.Participant {
	return &collaborativeProcess.def.Collaboration[0].Participant[1]
}

// GetPlane ...
func (collaborativeProcess CollaborativeProcess) GetPlane() *models.Plane {
	return &collaborativeProcess.def.Diagram[0].Plane[0]
}
