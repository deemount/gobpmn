package collaborative_process

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
	// configuration
	CustomerSupportIsExecutable bool // inject: 1, next: -
	CustomerIsExecutable        bool // inject: 2, next: - (no injection; see injection rules for configuration: boolean)
	// pool related
	Collaboration          factory.Builder // inject: 3, index: 0, next: 3
	CustomerSupportID      factory.Builder // inject: 4, index: 0, next:
	CustomerSupportProcess factory.Builder // inject: 5, index: 0, next: 2
	CustomerID             factory.Builder // inject: 6, index: 0, next: 3
	CustomerProcess        factory.Builder // inject: 7, index: 0, next: 4
}

// Message ...
type Message struct {
	CustomerToCustomerSupportMessage factory.Builder //
	CustomerSupportToCustomerMessage factory.Builder // next: 5
}

// CustomerSupportProcess ...
type CustomerSupport struct {
	CustomerSupportStartEvent     factory.Builder //
	FromCustomerSupportStartEvent factory.Builder //next: 6
	CheckIncomingClaimTask        factory.Builder //next: 7
	FromCheckIncomingClaimTask    factory.Builder //next: 8
	DenyWarrantyClaimTask         factory.Builder //next: 9
	FromDenyWarrantyClaimTask     factory.Builder //next: 10
	CustomerSupportEndEvent       factory.Builder //next: 11
}

// Customer ...
type Customer struct {
	CustomerStartEvent                   factory.Builder //
	FromCustomerStartEvent               factory.Builder //next: 12
	NoticeOfDefectTask                   factory.Builder //next: 13
	FromNoticeOfDefectTask               factory.Builder //next: 14
	WaitingForAnswerTask                 factory.Builder //next: 15
	TimerEventDefinitionWaitingForAnswer factory.Builder //next: 16
	FromWaitingForAnswerTask             factory.Builder //next: 17
	ReceiptWarrantyRefusalTask           factory.Builder //next: 18
	FromReceiptWarrantyRefusalTask       factory.Builder //next: 19
	CustomerEndEvent                     factory.Builder //next: 20
}

// NewCollaborativeProcess refers to the definitions struct to start building the model
// e.g. def: new(models.Definitions) use factory function, if zero value is insufficent (allocates memory)
func New() Proxy {
	p := Builder.Inject(Process{}).(Process)
	p.Def = core.NewDefinitions()
	Builder.Build(p.Def)
	return &p
}

// Build ...
func (p Process) Build() Process {
	// Elements
	p.elements()
	p.attributes()
	// Collaboration
	p.setCollaboration()
	// Customer Support
	p.setCustomerSupportProcess()
	p.setCustomerSupportStartEvent()
	p.setCheckIncomingClaimTask()
	p.setDenyWarrantyClaimTask()
	p.setCustomerSupportEndEvent()
	// Customer
	p.setCustomerProcess()
	p.setCustomerStartEvent()
	p.setNoticeOfDefectTask()
	p.setWaitingForAnswerTask()
	p.setReceiptWarrantyRefusalTask()
	p.setCustomerEndEvent()
	// Canvas
	p.setDiagram()
	return p
}

// Def ...
func (p Process) Call() core.DefinitionsRepository { return p.Def }

// setInnerElements ...
func (p *Process) elements() {
	// Collaboration
	collaboration := p.collaboration()
	collaboration.SetID("collaboration", p.Collaboration.Suffix)
	collaboration.SetParticipant(Builder.NumPart)
	collaboration.SetMessageFlow(Builder.NumMsg)
	// Processes
	p.customerSupportProcess().SetStartEvent(Builder.NumStartEvent / Builder.NumPart)
	p.customerSupportProcess().SetEndEvent(Builder.NumEndEvent / Builder.NumPart)
	p.customerSupportProcess().SetTask(2)
	p.customerSupportProcess().SetSequenceFlow(3)
	p.customerProcess().SetStartEvent(Builder.NumStartEvent / Builder.NumPart)
	p.customerProcess().SetEndEvent(Builder.NumEndEvent / Builder.NumPart)
	p.customerProcess().SetTask(2)
	p.customerProcess().SetIntermediateCatchEvent(1)
	p.customerProcess().SetSequenceFlow(4)
	// Canvas
	diagram := p.diagram()
	diagram.SetPlane()
	plane := p.plane()
	plane.SetShape(Builder.NumShape)
	plane.SetEdge(9)
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
			ED: p.plane().GetEdge(7),
			N:  "claim",
			ST: "activity",
			TT: "activity",
			H:  []string{p.CustomerToCustomerSupportMessage.Suffix, p.NoticeOfDefectTask.Suffix, p.CheckIncomingClaimTask.Suffix},
			BS: canvas.Bounds{X: 387, Y: 290, Width: 26, Height: 14},
			WP: []canvas.Waypoint{{X: 370, Y: 400}, {X: 370, Y: 200}}})
	flow.SetMessageFlow(
		flow.DelegateParameter{
			MF: p.collaboration().GetMessageFlow(1),
			ED: p.plane().GetEdge(8),
			N:  "refusal",
			ST: "activity",
			TT: "activity",
			H:  []string{p.CustomerSupportToCustomerMessage.Suffix, p.DenyWarrantyClaimTask.Suffix, p.ReceiptWarrantyRefusalTask.Suffix},
			BS: canvas.Bounds{X: 643, Y: 290, Width: 34, Height: 14},
			WP: []canvas.Waypoint{{X: 630, Y: 200}, {X: 630, Y: 400}}})
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
			H: p.CustomerSupportID.Suffix})
}

// setCustomerSupportProcess ...
func (p *Process) setCustomerSupportProcess() {
	p.customerSupportProcess().SetID("process", p.CustomerSupportProcess.Suffix)
	p.customerSupportProcess().SetIsExecutable(p.CustomerSupportIsExecutable)
}

// setCustomerSupportStartEvent ...
func (p *Process) setCustomerSupportStartEvent() {
	elements.SetStartEvent(
		elements.DelegateParameter{
			SE:    p.customerSupportProcess().GetStartEvent(0),
			SH:    p.plane().GetShape(1),
			BSPTR: p.plane().GetShape(0).GetBounds(),
			H:     []string{p.CustomerSupportStartEvent.Suffix, p.FromCustomerSupportStartEvent.Suffix}})
	p.fromCustomerSupportStartEvent()
}

// fromCustomerSupportStartEvent ...
func (p *Process) fromCustomerSupportStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF:    p.customerSupportProcess().GetSequenceFlow(0),
			ED:    p.plane().GetEdge(0),
			BSPTR: p.plane().GetShape(1).GetBounds(),
			ST:    "event",
			TT:    "activity",
			H:     []string{p.FromCustomerSupportStartEvent.Suffix, p.CustomerSupportStartEvent.Suffix, p.CheckIncomingClaimTask.Suffix}})
}

// setCheckIncomingClaim ...
func (p *Process) setCheckIncomingClaimTask() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA:     p.customerSupportProcess().GetTask(0),
			SH:     p.plane().GetShape(2),
			WPPREV: p.plane().GetEdge(0).GetWaypoint(1),
			N:      "Check incoming claim",
			H:      []string{p.CheckIncomingClaimTask.Suffix, p.FromCustomerSupportStartEvent.Suffix, p.FromCheckIncomingClaimTask.Suffix}})
	p.fromCheckIncomingClaimTask()
}

// fromCheckIncomingClaim ...
func (p *Process) fromCheckIncomingClaimTask() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF:    p.customerSupportProcess().GetSequenceFlow(1),
			ED:    p.plane().GetEdge(1),
			BSPTR: p.plane().GetShape(2).GetBounds(),
			N:     "decide",
			ST:    "activity",
			TT:    "activity",
			H:     []string{p.FromCheckIncomingClaimTask.Suffix, p.CheckIncomingClaimTask.Suffix, p.DenyWarrantyClaimTask.Suffix},
			BS:    canvas.Bounds{X: 485, Y: 142, Width: 33, Height: 14}}) // TODO: label bounds needs to fit to the flow
}

// setDenyWarrantyClaim ...
func (p *Process) setDenyWarrantyClaimTask() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA:     p.customerSupportProcess().GetTask(1),
			SH:     p.plane().GetShape(3),
			WPPREV: p.plane().GetEdge(1).GetWaypoint(1),
			N:      "Deny warranty claim",
			H:      []string{p.DenyWarrantyClaimTask.Suffix, p.FromCheckIncomingClaimTask.Suffix, p.FromDenyWarrantyClaimTask.Suffix}})
	p.fromDenyWarrantyClaimTask()
}

// fromDenyWarrantyClaim ...
func (p *Process) fromDenyWarrantyClaimTask() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF:    p.customerSupportProcess().GetSequenceFlow(2),
			ED:    p.plane().GetEdge(2),
			BSPTR: p.plane().GetShape(3).GetBounds(),
			ST:    "activity",
			TT:    "event",
			H:     []string{p.FromDenyWarrantyClaimTask.Suffix, p.DenyWarrantyClaimTask.Suffix, p.CustomerSupportEndEvent.Suffix}})
}

// setCustomerSupportEndEvent ...
func (p *Process) setCustomerSupportEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE:     p.customerSupportProcess().GetEndEvent(0),
			SH:     p.plane().GetShape(4),
			WPPREV: p.plane().GetEdge(2).GetWaypoint(1),
			H:      []string{p.CustomerSupportEndEvent.Suffix, p.FromDenyWarrantyClaimTask.Suffix}})
}

/**** Customer Process ****/

// setCustomerPool
func (p *Process) setCustomerPool() {
	canvas.SetPool(
		canvas.DelegateParameter{
			S:     p.plane().GetShape(5),
			BSPTR: p.plane().GetShape(0).GetBounds(),
			T:     "participant",
			I:     true,
			H:     p.CustomerID.Suffix})
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
			SE:    p.customerProcess().GetStartEvent(0),
			SH:    p.plane().GetShape(6),
			BSPTR: p.plane().GetShape(5).GetBounds(),
			H:     []string{p.CustomerStartEvent.Suffix, p.FromCustomerStartEvent.Suffix}})
	p.fromCustomerStartEvent()
}

// fromCustomerStartEvent ...
func (p *Process) fromCustomerStartEvent() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF:    p.customerProcess().GetSequenceFlow(0),
			ED:    p.plane().GetEdge(3),
			BSPTR: p.plane().GetShape(6).GetBounds(),
			ST:    "event",
			TT:    "activity",
			H:     []string{p.FromCustomerStartEvent.Suffix, p.CustomerStartEvent.Suffix, p.NoticeOfDefectTask.Suffix}})
}

// setNoticeOfDefect ...
func (p *Process) setNoticeOfDefectTask() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA:     p.customerProcess().GetTask(0),
			SH:     p.plane().GetShape(7),
			WPPREV: p.plane().GetEdge(3).GetWaypoint(1),
			N:      "Notice of defect",
			H:      []string{p.NoticeOfDefectTask.Suffix, p.FromCustomerStartEvent.Suffix, p.FromNoticeOfDefectTask.Suffix}})
	p.fromNoticeOfDefect()
}

// fromNoticeOfDefect ...
func (p *Process) fromNoticeOfDefect() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF:    p.customerProcess().GetSequenceFlow(1),
			ED:    p.plane().GetEdge(4),
			BSPTR: p.plane().GetShape(7).GetBounds(),
			ST:    "activity",
			TT:    "event",
			H:     []string{p.FromNoticeOfDefectTask.Suffix, p.NoticeOfDefectTask.Suffix, p.WaitingForAnswerTask.Suffix}})
}

// setWaitingForAnswer ...
func (p *Process) setWaitingForAnswerTask() {
	elements.SetIntermediateCatchEvent(
		elements.DelegateParameter{
			ISTED:  true,
			ICE:    p.customerProcess().GetIntermediateCatchEvent(0),
			SH:     p.plane().GetShape(8),
			WPPREV: p.plane().GetEdge(4).GetWaypoint(1),
			T:      "event",
			N:      "Waiting for answer",
			TD:     "PT1M",
			H:      []string{p.WaitingForAnswerTask.Suffix, p.FromNoticeOfDefectTask.Suffix, p.FromWaitingForAnswerTask.Suffix},
			TEDH:   p.TimerEventDefinitionWaitingForAnswer.Suffix})
	p.fromWaitingForAnswer()
}

// fromWaitingForAnswer ...
func (p *Process) fromWaitingForAnswer() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF:    p.customerProcess().GetSequenceFlow(2),
			ED:    p.plane().GetEdge(5),
			BSPTR: p.plane().GetShape(8).GetBounds(),
			ST:    "event",
			TT:    "activity",
			H:     []string{p.FromWaitingForAnswerTask.Suffix, p.WaitingForAnswerTask.Suffix, p.ReceiptWarrantyRefusalTask.Suffix}})
}

// setReceiptWarrantyRefusal ...
func (p *Process) setReceiptWarrantyRefusalTask() {
	tasks.SetTask(
		tasks.DelegateParameter{
			TA:     p.customerProcess().GetTask(1),
			SH:     p.plane().GetShape(9),
			WPPREV: p.plane().GetEdge(5).GetWaypoint(1),
			N:      "Receipt warranty refusal",
			H:      []string{p.ReceiptWarrantyRefusalTask.Suffix, p.FromWaitingForAnswerTask.Suffix, p.FromReceiptWarrantyRefusalTask.Suffix}})
	p.fromReceiptWarrantyRefusal()
}

// fromReceiptWarrantyRefusal ...
func (p *Process) fromReceiptWarrantyRefusal() {
	flow.SetSequenceFlow(
		flow.DelegateParameter{
			SF:    p.customerProcess().GetSequenceFlow(3),
			ED:    p.plane().GetEdge(6),
			BSPTR: p.plane().GetShape(9).GetBounds(),
			ST:    "activity",
			TT:    "event",
			H:     []string{p.FromReceiptWarrantyRefusalTask.Suffix, p.ReceiptWarrantyRefusalTask.Suffix, p.CustomerEndEvent.Suffix}})
}

// setCustomerEndEvent ...
func (p *Process) setCustomerEndEvent() {
	elements.SetEndEvent(
		elements.DelegateParameter{
			EE:     p.customerProcess().GetEndEvent(0),
			SH:     p.plane().GetShape(10),
			WPPREV: p.plane().GetEdge(6).GetWaypoint(1),
			H:      []string{p.CustomerEndEvent.Suffix, p.FromReceiptWarrantyRefusalTask.Suffix}})
}

// setDiagram ...
func (p *Process) setDiagram() {
	// diagram attributes
	var n int64 = 1
	p.diagram().SetID("diagram", n)
	// plane attributes
	p.plane().SetID("plane", n)
	p.plane().SetElement("collaboration", p.Collaboration.Suffix)
}

/**** Default Setter/Getter ****/

func (p *Process) attributes()                                { p.Def.SetDefaultAttributes() }
func (p Process) collaboration() *collaboration.Collaboration { return p.Def.GetCollaboration() }
func (p Process) customerSupportProcess() *process.Process    { return p.Def.GetProcess(0) }
func (p Process) customerProcess() *process.Process           { return p.Def.GetProcess(1) }
func (p Process) diagram() *canvas.Diagram                    { return p.Def.GetDiagram(0) }
func (p Process) plane() *canvas.Plane                        { return p.diagram().GetPlane() }
