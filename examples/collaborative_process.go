package examples

import (
	"github.com/deemount/gobpmn/factory"
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/collaboration"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
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
	cp := bb.Inject(CProcess{}).(CProcess)
	cp.Def = core.NewDefinitions(cp)
	return &cp
}

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
	cp.setCustomerEndEvent()
	// Canvas
	cp.setDiagram()
	return cp
}

// Def ...
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
			T: "flow",
			H: cp.CustomerSupportToCustomerMessage.Suffix,
			W: []canvas.Waypoint{{X: 630, Y: 200}, {X: 630, Y: 400}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
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

/**** Customer Process ****/

// setPoolCustomer
func (cp *CProcess) setCustomerPool() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: cp.plane().GetShape(1),
			T: "participant",
			I: true,
			H: cp.CustomerID.Suffix,
			B: canvas.Bounds{X: 150, Y: 360, Width: 800, Height: 160}})
}

// setProcessCustomer ...
func (cp *CProcess) setCustomerProcess() {
	cp.customerProcess().SetID("process", cp.CustomerProcess.Suffix)
	cp.customerProcess().SetIsExecutable(cp.CustomerIsExecutable)
}

// setCustomerStartEvent ...
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
}

// setNoticeOfDefect ...
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

/**** Getter ****/

// collaboration ...
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
}
