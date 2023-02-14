package collaborative_process

import (
	"log"
	"reflect"

	"github.com/deemount/gobpmn/factory"
	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/collaboration"
	"github.com/deemount/gobpmn/models/bpmn/core"
	"github.com/deemount/gobpmn/models/bpmn/events/elements"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/process"
	"github.com/deemount/gobpmn/models/bpmn/tasks"
)

// Process ...
type Process struct {
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
	Collaboration               factory.Builder
	CustomerSupportID           factory.Builder
	CustomerSupportProcess      factory.Builder
	CustomerID                  factory.Builder
	CustomerProcess             factory.Builder
}

// Message ...
type Message struct {
	CustomerToCustomerSupportMessage factory.Builder
	CustomerSupportToCustomerMessage factory.Builder
}

// CustomerSupportProcess ...
type CustomerSupport struct {
	CustomerSupportStartEvent     factory.Builder
	FromCustomerSupportStartEvent factory.Builder
	CheckIncomingClaim            factory.Builder
	FromCheckIncomingClaim        factory.Builder
	DenyWarrantyClaim             factory.Builder
	FromDenyWarrantyClaim         factory.Builder
	CustomerSupportEndEvent       factory.Builder
}

// Customer ...
type Customer struct {
	CustomerStartEvent                   factory.Builder
	FromCustomerStartEvent               factory.Builder
	NoticeOfDefect                       factory.Builder
	FromNoticeOfDefect                   factory.Builder
	WaitingForAnswer                     factory.Builder
	TimerEventDefinitionWaitingForAnswer factory.Builder
	FromWaitingForAnswer                 factory.Builder
	ReceiptWarrantyRefusal               factory.Builder
	FromReceiptWarrantyRefusal           factory.Builder
	CustomerEndEvent                     factory.Builder
}

// NewCollaborativeProcess refers to the definitions struct to start building the model
// e.g. def: new(models.Definitions) use factory function, if zero value is insufficent (allocates memory)
func New() Proxy {
	p := Builder.Inject(Process{}).(Process)
	p.Def = core.NewDefinitions()
	miType := reflect.TypeOf(p.Def)
	for i := 0; i < miType.NumMethod(); i++ {
		method := miType.Method(i)
		log.Printf("%v", method.Name)
	}
	return &p
}

// Build ...
func (p Process) Build() Process {
	// Elements
	Builder.Build(p.Def)
	p.setInnerElements()
	p.setDefinitionsAttributes()
	p.setCollaboration()
	// Customer Support
	p.setCustomerSupportProcess()
	p.setCustomerSupportStartEvent()
	p.setCheckIncomingClaim()
	p.setDenyWarrantyClaim()
	p.setCustomerSupportEndEvent()
	// Customer
	p.setCustomerProcess()
	p.setCustomerStartEvent()
	p.setNoticeOfDefect()
	p.setWaitingForAnswer()
	p.setReceiptWarrantyRefusal()
	p.setCustomerEndEvent()
	// Canvas
	p.setDiagram()
	return p
}

// Def ...
func (p Process) Call() core.DefinitionsRepository { return p.Def }

// setInnerElements ...
func (p *Process) setInnerElements() {
	// Collaboration
	collaboration := p.collaboration()
	collaboration.SetID("collaboration", p.Collaboration.Suffix)
	collaboration.SetParticipant(Builder.NumPart)
	collaboration.SetMessageFlow(Builder.NumMsg)
	// Processes
	p.customerSupportProcess().SetStartEvent(1)
	p.customerSupportProcess().SetEndEvent(1)
	p.customerSupportProcess().SetTask(2)
	p.customerSupportProcess().SetSequenceFlow(3)
	p.customerProcess().SetStartEvent(1)
	p.customerProcess().SetEndEvent(1)
	p.customerProcess().SetTask(2)
	p.customerProcess().SetIntermediateCatchEvent(1)
	p.customerProcess().SetSequenceFlow(4)
	// Canvas
	diagram := p.diagram()
	diagram.SetPlane()
	plane := p.plane()
	plane.SetShape(11)
	plane.SetEdge(9)
}

// setDefinitionsAttributes ...
func (p *Process) setDefinitionsAttributes() {
	p.Def.SetDefaultAttributes()
}

// setCollaboration ...
func (p *Process) setCollaboration() {
	p.setParticipants()
	p.setMessageFlows()
	p.setPools()
}

// setParticipants ...
func (p *Process) setParticipants() {
	// customer support
	collaboration.SetParticipant(
		collaboration.DelegateParameter{
			PPT: p.collaboration().GetParticipant(0),
			T:   "participant",
			PR:  "process",
			N:   "Customer Support",
			H:   []string{p.CustomerSupportID.Suffix, p.CustomerSupportProcess.Suffix}})
	// customer
	collaboration.SetParticipant(
		collaboration.DelegateParameter{
			PPT: p.collaboration().GetParticipant(1),
			T:   "participant",
			PR:  "process",
			N:   "Customer",
			H:   []string{p.CustomerID.Suffix, p.CustomerProcess.Suffix}})
}

// setMessageFlows ...
func (p *Process) setMessageFlows() {
	flow.SetMessageFlow(
		flow.DelegateParameter{
			MF: p.collaboration().GetMessageFlow(0),
			T:  "flow",
			N:  "claim",
			ST: "activity",
			TT: "activity",
			H:  []string{p.CustomerToCustomerSupportMessage.Suffix, p.NoticeOfDefect.Suffix, p.CheckIncomingClaim.Suffix}})
	p.setMessageClaim()
	flow.SetMessageFlow(
		flow.DelegateParameter{
			MF: p.collaboration().GetMessageFlow(1),
			T:  "flow",
			N:  "refusal",
			ST: "activity",
			TT: "activity",
			H:  []string{p.CustomerSupportToCustomerMessage.Suffix, p.DenyWarrantyClaim.Suffix, p.ReceiptWarrantyRefusal.Suffix}})
	p.setMessageRefusal()
}

// setMessageClaim ...
func (p *Process) setMessageClaim() {
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: p.plane().GetEdge(7),
			T: "flow",
			H: p.CustomerToCustomerSupportMessage.Suffix,
			W: []canvas.Waypoint{{X: 370, Y: 400}, {X: 370, Y: 200}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
			E: p.plane().GetEdge(7),
			B: canvas.Bounds{X: 387, Y: 290, Width: 26, Height: 14}})
}

// setMessageRefusal ...
func (p *Process) setMessageRefusal() {
	canvas.SetEdge(
		canvas.DelegateParameter{
			E: p.plane().GetEdge(8),
			T: "flow",
			H: p.CustomerSupportToCustomerMessage.Suffix,
			W: []canvas.Waypoint{{X: 630, Y: 200}, {X: 630, Y: 400}}})
	canvas.SetLabel(
		canvas.DelegateParameter{
			E: p.plane().GetEdge(8),
			B: canvas.Bounds{X: 643, Y: 290, Width: 34, Height: 14}})
}

// setPools ...
func (p *Process) setPools() {
	p.setCustomerSupportPool()
	p.setCustomerPool()
}

/* Customer Support Process */

// setPoolCustomerSupport ...
func (p *Process) setCustomerSupportPool() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: p.plane().GetShape(0),
			T: "participant",
			I: true,
			H: p.CustomerSupportID.Suffix,
			B: canvas.Bounds{X: 150, Y: 80, Width: 800, Height: 160}})
}

// setProcessCustomerSupport ...
func (p *Process) setCustomerSupportProcess() {
	p.customerSupportProcess().SetID("process", p.CustomerSupportProcess.Suffix)
	p.customerSupportProcess().SetIsExecutable(p.CustomerSupportIsExecutable)
}

// setCustomerSupportStartEvent ...
func (p *Process) setCustomerSupportStartEvent() {
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: p.customerSupportProcess().GetStartEvent(0),
			SH: p.plane().GetShape(2),
			T:  "event",
			N:  "Begin of Customer Support Process",
			H:  []string{p.CustomerSupportStartEvent.Suffix, p.FromCustomerSupportStartEvent.Suffix},
			BS: canvas.Bounds{X: 225, Y: 142}})
	p.fromCustomerSupportStartEvent()
}

// fromCustomerSupportStartEvent ...
func (p *Process) fromCustomerSupportStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.customerSupportProcess().GetSequenceFlow(0),
			ED: p.plane().GetEdge(0),
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{p.FromCustomerSupportStartEvent.Suffix, p.CustomerSupportStartEvent.Suffix, p.CheckIncomingClaim.Suffix},
			WP: []canvas.Waypoint{{X: 261, Y: 160}, {X: 320, Y: 160}}})
}

// setCheckIncomingClaim ...
func (p *Process) setCheckIncomingClaim() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: p.customerSupportProcess().GetTask(0),
			SH: p.plane().GetShape(4),
			T:  "activity",
			N:  "Check incoming claim",
			H:  []string{p.CheckIncomingClaim.Suffix, p.FromCustomerSupportStartEvent.Suffix, p.FromCheckIncomingClaim.Suffix},
			BS: canvas.Bounds{X: 320, Y: 120}})
	p.fromCheckIncomingClaim()
}

// fromCheckIncomingClaim ...
func (p *Process) fromCheckIncomingClaim() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.customerSupportProcess().GetSequenceFlow(1),
			ED: p.plane().GetEdge(1),
			T:  "flow",
			N:  "decide",
			ST: "activity",
			TT: "activity",
			H:  []string{p.FromCheckIncomingClaim.Suffix, p.CheckIncomingClaim.Suffix, p.DenyWarrantyClaim.Suffix},
			WP: []canvas.Waypoint{{X: 420, Y: 160}, {X: 580, Y: 160}},
			BS: canvas.Bounds{X: 485, Y: 142, Width: 33, Height: 14}})
}

// setDenyWarrantyClaim ...
func (p *Process) setDenyWarrantyClaim() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: p.customerSupportProcess().GetTask(1),
			T:  "activity",
			N:  "Deny warranty claim",
			H:  []string{p.DenyWarrantyClaim.Suffix, p.FromCheckIncomingClaim.Suffix, p.FromDenyWarrantyClaim.Suffix},
			SH: p.plane().GetShape(5),
			BS: canvas.Bounds{X: 580, Y: 120}})
	p.fromDenyWarrantyClaim()
}

// fromDenyWarrantyClaim ...
func (p *Process) fromDenyWarrantyClaim() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.customerSupportProcess().GetSequenceFlow(2),
			ED: p.plane().GetEdge(2),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{p.FromDenyWarrantyClaim.Suffix, p.DenyWarrantyClaim.Suffix, p.CustomerSupportEndEvent.Suffix},
			WP: []canvas.Waypoint{{X: 680, Y: 160}, {X: 822, Y: 160}}})
}

// setCustomerSupportEndEvent ...
func (p *Process) setCustomerSupportEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: p.customerSupportProcess().GetEndEvent(0),
			SH: p.plane().GetShape(3),
			T:  "event",
			N:  "End of Customer Support Process",
			H:  []string{p.CustomerSupportEndEvent.Suffix, p.FromDenyWarrantyClaim.Suffix},
			BS: canvas.Bounds{X: 822, Y: 142}})
}

/**** Customer Process ****/

// setPoolCustomer
func (p *Process) setCustomerPool() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S: p.plane().GetShape(1),
			T: "participant",
			I: true,
			H: p.CustomerID.Suffix,
			B: canvas.Bounds{X: 150, Y: 360, Width: 800, Height: 160}})
}

// setProcessCustomer ...
func (p *Process) setCustomerProcess() {
	p.customerProcess().SetID("process", p.CustomerProcess.Suffix)
	p.customerProcess().SetIsExecutable(p.CustomerIsExecutable)
}

// setCustomerStartEvent ...
func (p *Process) setCustomerStartEvent() {
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE: p.customerProcess().GetStartEvent(0),
			SH: p.plane().GetShape(6),
			T:  "event",
			N:  "Begin of Customer Process",
			H:  []string{p.CustomerStartEvent.Suffix, p.FromCustomerStartEvent.Suffix},
			BS: canvas.Bounds{X: 225, Y: 422}})
	p.fromCustomerStartEvent()
}

// fromCustomerStartEvent ...
func (p *Process) fromCustomerStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.customerProcess().GetSequenceFlow(0),
			ED: p.plane().GetEdge(3),
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{p.FromCustomerStartEvent.Suffix, p.CustomerStartEvent.Suffix, p.NoticeOfDefect.Suffix},
			WP: []canvas.Waypoint{{X: 261, Y: 440}, {X: 320, Y: 440}}})
}

// setNoticeOfDefect ...
func (p *Process) setNoticeOfDefect() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: p.customerProcess().GetTask(0),
			SH: p.plane().GetShape(8),
			T:  "activity",
			N:  "Notice of defect",
			H:  []string{p.NoticeOfDefect.Suffix, p.FromCustomerStartEvent.Suffix, p.FromNoticeOfDefect.Suffix},
			BS: canvas.Bounds{X: 320, Y: 400}})
	p.fromNoticeOfDefect()
}

// fromNoticeOfDefect ...
func (p *Process) fromNoticeOfDefect() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.customerProcess().GetSequenceFlow(1),
			ED: p.plane().GetEdge(4),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{p.FromNoticeOfDefect.Suffix, p.NoticeOfDefect.Suffix, p.WaitingForAnswer.Suffix},
			WP: []canvas.Waypoint{{X: 420, Y: 440}, {X: 482, Y: 440}}})
}

// setWaitingForAnswer ...
func (p *Process) setWaitingForAnswer() {
	elements.SetIntermediateCatchEvent(
		elements.DelegateParameter{
			ISTED: true,
			ICE:   p.customerProcess().GetIntermediateCatchEvent(0),
			SH:    p.plane().GetShape(9),
			T:     "event",
			N:     "Waiting for answer",
			TD:    "PT1M",
			H:     []string{p.WaitingForAnswer.Suffix, p.FromNoticeOfDefect.Suffix, p.FromWaitingForAnswer.Suffix},
			TEDH:  p.TimerEventDefinitionWaitingForAnswer.Suffix,
			BS:    canvas.Bounds{X: 482, Y: 422}})
	p.fromWaitingForAnswer()
}

// fromWaitingForAnswer ...
func (p *Process) fromWaitingForAnswer() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.customerProcess().GetSequenceFlow(2),
			ED: p.plane().GetEdge(5),
			T:  "flow",
			ST: "event",
			TT: "activity",
			H:  []string{p.FromWaitingForAnswer.Suffix, p.WaitingForAnswer.Suffix, p.ReceiptWarrantyRefusal.Suffix},
			WP: []canvas.Waypoint{{X: 518, Y: 440}, {X: 580, Y: 440}}})
}

// setReceiptWarrantyRefusal ...
func (p *Process) setReceiptWarrantyRefusal() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA: p.customerProcess().GetTask(1),
			SH: p.plane().GetShape(10),
			T:  "activity",
			N:  "Receipt warranty refusal",
			H:  []string{p.ReceiptWarrantyRefusal.Suffix, p.FromWaitingForAnswer.Suffix, p.FromReceiptWarrantyRefusal.Suffix},
			BS: canvas.Bounds{X: 580, Y: 400}})
	p.fromReceiptWarrantyRefusal()
}

// fromReceiptWarrantyRefusal ...
func (p *Process) fromReceiptWarrantyRefusal() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF: p.customerProcess().GetSequenceFlow(3),
			ED: p.plane().GetEdge(6),
			T:  "flow",
			ST: "activity",
			TT: "event",
			H:  []string{p.FromReceiptWarrantyRefusal.Suffix, p.ReceiptWarrantyRefusal.Suffix, p.CustomerEndEvent.Suffix},
			WP: []canvas.Waypoint{{X: 680, Y: 440}, {X: 822, Y: 440}}})
}

// setCustomerEndEvent ...
func (p *Process) setCustomerEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE: p.customerProcess().GetEndEvent(0),
			SH: p.plane().GetShape(7),
			T:  "event",
			N:  "End of Customer Process",
			H:  []string{p.CustomerEndEvent.Suffix, p.FromReceiptWarrantyRefusal.Suffix},
			BS: canvas.Bounds{X: 822, Y: 422}})
}

// setDiagram ...
func (p *Process) setDiagram() {
	// diagram attributes
	var n int64 = 1
	p.diagram().SetID("diagram", n)
	// plane attributes
	plane := p.plane()
	plane.SetID("plane", n)
	plane.SetElement("collaboration", p.Collaboration.Suffix)
}

/**** Getter ****/

// collaboration ...
func (p Process) collaboration() *collaboration.Collaboration { return p.Def.GetCollaboration() }

// customerSupportProces ...
func (p Process) customerSupportProcess() *process.Process { return p.Def.GetProcess(0) }

// customerProcess ...
func (p Process) customerProcess() *process.Process { return p.Def.GetProcess(1) }

// diagram ...
func (p Process) diagram() *canvas.Diagram { return p.Def.GetDiagram(0) }

// plane ...
func (p Process) plane() *canvas.Plane { return p.diagram().GetPlane() }
