package infrastructure

import (
	"fmt"
	"strings"
)

// Collaboration
type (

	// Collaboration ...
	Collaboration struct {
		ID          string        `xml:"id,attr,omitempty" json:"id"`
		Participant []Participant `xml:"bpmn:participant" json:"participant,omitempty"`
	}

	// TCollaboration ...
	TCollaboration struct {
		ID          string         `xml:"id,attr,omitempty" json:"id"`
		Participant []TParticipant `xml:"participant" json:"participant,omitempty"`
	}
)

// SetID ...
func (c *Collaboration) SetID(typ string, suffix interface{}) {
	switch typ {
	case "collaboration":
		c.ID = fmt.Sprintf("Collaboration_%v", suffix)
	case "id":
		c.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (c Collaboration) GetID() *string {
	return &c.ID
}

// SetParticipant ...
func (c *Collaboration) SetParticipant(num int) {
	c.Participant = make([]Participant, num)
}

// GetParticipant ...
func (c Collaboration) GetParticipant(num int) *Participant {
	return &c.Participant[num]
}

// Participant
type (

	// Participant ...
	Participant struct {
		ID         string `xml:"id,attr,omitempty" json:"id"`
		Name       string `xml:"name,attr,omitempty" json:"name,omitempty"`
		ProcessRef string `xml:"processRef,attr,omitempty" json:"processRef,omitempty"`
	}

	// TParticipant ...
	TParticipant struct {
		ID         string `xml:"id,attr,omitempty" json:"id"`
		Name       string `xml:"name,attr,omitempty" json:"name,omitempty"`
		ProcessRef string `xml:"processRef,attr" json:"processRef,omitempty"`
	}
)

// SetID ...
func (p *Participant) SetID(typ string, suffix interface{}) {
	switch typ {
	case "participant":
		p.ID = fmt.Sprintf("Participant_%v", suffix)
	case "id":
		p.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (p Participant) GetID() *string {
	return &p.ID
}

// SetName ...
func (p *Participant) SetName(name string) {
	p.Name = name
}

// GetName ...
func (p Participant) GetName() *string {
	return &p.Name
}

// SetProcessRef ...
func (participant *Participant) SetProcessRef(typ string, suffix interface{}) {
	switch typ {
	case "process":
		participant.ProcessRef = fmt.Sprintf("Process_%s", suffix)
	case "id":
		participant.ProcessRef = fmt.Sprintf("%s", suffix)
	}
}

// GetProcessRef ...
func (p Participant) GetProcessRef() *string {
	return &p.ProcessRef
}

// Process
type (

	// Process ...
	Process struct {
		ID                     string                   `xml:"id,attr,omitempty" json:"id"`
		Name                   string                   `xml:"name,attr,omitempty" json:"name,omitempty"`
		IsExecutable           bool                     `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
		LaneSet                []LaneSet                `xml:"bpmn:laneSet,omitempty" json:"laneSet,omitempty"` // Note: Pool
		StartEvent             []StartEvent             `xml:"bpmn:startEvent,omitempty" json:"startEvent,omitempty"`
		EndEvent               []EndEvent               `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
		IntermediateCatchEvent []IntermediateCatchEvent `xml:"bpmn:intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty"`
		IntermediateThrowEvent []IntermediateThrowEvent `xml:"bpmn:intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty"`
		InclusiveGateway       []InclusiveGateway       `xml:"bpmn:inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
		ExclusiveGateway       []ExclusiveGateway       `xml:"bpmn:exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
		ParallelGateway        []ParallelGateway        `xml:"bpmn:parallelGateway,omitempty" json:"parallelGateway,omitempty"`
		Gateway                []Gateway                `xml:"bpmn:gateway,omitempty" json:"gateway,omitempty"`
		UserTask               []UserTask               `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty"`
		ServiceTask            []ServiceTask            `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty"`
		ScriptTask             []ScriptTask             `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty"`
		Task                   []Task                   `xml:"bpmn:task,omitempty" json:"task,omitempty"`
		SequenceFlow           []SequenceFlow           `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
	}

	// TProcess ...
	TProcess struct {
		ID                     string                    `xml:"id,attr,omitempty" json:"id"`
		Name                   string                    `xml:"name,attr,omitempty" json:"name,omitempty"`
		IsExecutable           bool                      `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
		LaneSet                []TLaneSet                `xml:"laneSet,omitempty" json:"laneSet,omitempty"` // Note: Pool
		StartEvent             []TStartEvent             `xml:"startEvent,omitempty" json:"startEvent,omitempty"`
		EndEvent               []TEndEvent               `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
		IntermediateCatchEvent []TIntermediateCatchEvent `xml:"intermediateCatchEvent,omitempty" json:"intermediateCatchEvent,omitempty"`
		IntermediateThrowEvent []TIntermediateThrowEvent `xml:"intermediateThrowEvent,omitempty" json:"intermediateThrowEvent,omitempty"`
		InclusiveGateway       []TInclusiveGateway       `xml:"inclusiveGateway,omitempty" json:"inclusiveGateway,omitempty"`
		ExclusiveGateway       []TExclusiveGateway       `xml:"exclusiveGateway,omitempty" json:"exclusiveGateway,omitempty"`
		ParallelGateway        []TParallelGateway        `xml:"parallelGateway,omitempty" json:"parallelGateway,omitempty"`
		Gateway                []TGateway                `xml:"gateway,omitempty" json:"gateway,omitempty"`
		UserTask               []TUserTask               `xml:"userTask,omitempty" json:"userTask,omitempty"`
		ServiceTask            []TServiceTask            `xml:"serviceTask,omitempty" json:"serviceTask,omitempty"`
		ScriptTask             []TScriptTask             `xml:"scriptTask,omitempty" json:"scriptTask,omitempty"`
		Task                   []TTask                   `xml:"task,omitempty" json:"task,omitempty"`
		SequenceFlow           []TSequenceFlow           `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
	}
)

// SetID ...
func (ps *Process) SetID(typ string, suffix interface{}) {
	switch typ {
	case "process":
		ps.ID = fmt.Sprintf("Process_%v", suffix)
	case "id":
		ps.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (ps Process) GetID() *string {
	return &ps.ID
}

// SetName ...
func (ps *Process) SetName(name string) {
	ps.Name = name
}

// GetName ...
func (ps Process) GetName() *string {
	return &ps.Name
}

// SetIsExecutable ...
func (ps *Process) SetIsExecutable(isExec bool) {
	ps.IsExecutable = isExec
}

// GetIsExecutable ...
func (ps Process) GetIsExecutable() *bool {
	return &ps.IsExecutable
}

// SetStartEvent ...
func (ps *Process) SetStartEvent(num int) {
	ps.StartEvent = make([]StartEvent, num)
}

// GetStartEvent ...
func (ps Process) GetStartEvent(num int) *StartEvent {
	return &ps.StartEvent[num]
}

// SetEndEvent ...
func (ps *Process) SetEndEvent(num int) {
	ps.EndEvent = make([]EndEvent, num)
}

// GetEndEvent ...
func (ps Process) GetEndEvent(num int) *EndEvent {
	return &ps.EndEvent[num]
}

// SetIntermediateCatchEvent ...
func (ps *Process) SetIntermediateCatchEvent(num int) {
	ps.IntermediateCatchEvent = make([]IntermediateCatchEvent, num)
}

// GetIntermediateCatchEvent ...
func (ps Process) GetIntermediateCatchEvent(num int) *IntermediateCatchEvent {
	return &ps.IntermediateCatchEvent[num]
}

// SetIntermediateThrowEvent ...
func (ps *Process) SetIntermediateThrowEvent(num int) {
	ps.IntermediateThrowEvent = make([]IntermediateThrowEvent, num)
}

// GetIntermediateThrowEvent ...
func (ps Process) GetIntermediateThrowEvent(num int) *IntermediateThrowEvent {
	return &ps.IntermediateThrowEvent[num]
}

// SetInclusiveGateway ...
func (ps *Process) SetInclusiveGateway(num int) {
	ps.InclusiveGateway = make([]InclusiveGateway, num)
}

// GetInclusiveGateway ...
func (ps Process) GetInclusiveGateway(num int) *InclusiveGateway {
	return &ps.InclusiveGateway[num]
}

// SetExclusiveGateway ...
func (ps *Process) SetExclusiveGateway(num int) {
	ps.ExclusiveGateway = make([]ExclusiveGateway, num)
}

// GetExclusiveGateway ...
func (ps Process) GetExclusiveGateway(num int) *ExclusiveGateway {
	return &ps.ExclusiveGateway[num]
}

// SetParallelGateway ...
func (ps *Process) SetParallelGateway(num int) {
	ps.ParallelGateway = make([]ParallelGateway, num)
}

// GetParallelGateway ...
func (ps Process) GetParallelGateway(num int) *ParallelGateway {
	return &ps.ParallelGateway[num]
}

// SetGateway ...
func (ps *Process) SetGateway(num int) {
	ps.Gateway = make([]Gateway, num)
}

// GetGateway ...
func (ps Process) GetGateway(num int) *Gateway {
	return &ps.Gateway[num]
}

// SetUserTask ...
func (ps *Process) SetUserTask(num int) {
	ps.UserTask = make([]UserTask, num)
}

// GetUserTask ...
func (ps Process) GetUserTask(num int) *UserTask {
	return &ps.UserTask[num]
}

// SetServiceTask ...
func (ps *Process) SetServiceTask(num int) {
	ps.ServiceTask = make([]ServiceTask, num)
}

// GetServiceTask ...
func (ps Process) GetServiceTask(num int) *ServiceTask {
	return &ps.ServiceTask[num]
}

// SetScriptTask ...
func (ps *Process) SetScriptTask(num int) {
	ps.ScriptTask = make([]ScriptTask, num)
}

// GetScriptTask ...
func (ps Process) GetScriptTask(num int) *ScriptTask {
	return &ps.ScriptTask[num]
}

// SetTask ...
func (ps *Process) SetTask(num int) {
	ps.Task = make([]Task, num)
}

// GetTask ...
func (ps Process) GetTask(num int) *Task {
	return &ps.Task[num]
}

// SetSequenceFlow ...
func (ps *Process) SetSequenceFlow(num int) {
	ps.SequenceFlow = make([]SequenceFlow, num)
}

// GetSequenceFlow ...
func (ps Process) GetSequenceFlow(num int) *SequenceFlow {
	return &ps.SequenceFlow[num]
}

// LaneSet ...
type (

	// LaneSet ...
	LaneSet struct {
		ID   string `xml:"id,attr,omitempty" json:"id"`
		Lane []Lane `xml:"bpmn:lane,omitempty" json:"lane,omitempty"`
	}

	// TLaneSet ...
	TLaneSet struct {
		ID   string  `xml:"id,attr,omitempty" json:"id"`
		Lane []TLane `xml:"lane,omitempty" json:"lane,omitempty"`
	}
)

// SetID ...
func (ls *LaneSet) SetID(typ string, suffix interface{}) {
	switch typ {
	case "laneset":
		ls.ID = fmt.Sprintf("LaneSet_%v", suffix)
	case "id":
		ls.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (ls LaneSet) GetID() *string {
	return &ls.ID
}

// SetLane ...
func (ls *LaneSet) SetLane(num int) {
	ls.Lane = make([]Lane, num)
}

// GetLane ...
func (ls LaneSet) GetLane(num int) *Lane {
	return &ls.Lane[num]
}

// Lane ...
type (

	// Lane ...
	Lane struct {
		ID          string        `xml:"id,attr,omitempty" json:"id"`
		FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
	}

	// TLane ...
	TLane struct {
		ID          string        `xml:"id,attr,omitempty" json:"id"`
		FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
	}
)

// SetID ...
func (lane *Lane) SetID(typ string, suffix interface{}) {
	switch typ {
	case "lane":
		lane.ID = fmt.Sprintf("Lane_%v", suffix)
	case "id":
		lane.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (lane Lane) GetID() *string {
	return &lane.ID
}

// SetFlowNodeRef ...
func (lane *Lane) SetFlowNodeRef(num int) {
	lane.FlowNodeRef = make([]FlowNodeRef, num)
}

// GetFlowNodeRef ...
func (lane Lane) GetFlowNodeRef(num int) *FlowNodeRef {
	return &lane.FlowNodeRef[num]
}

// FlowNodeRef ...
type FlowNodeRef struct {
	ID string `xml:",innerxml" json:"id"`
}

// SetID ...
func (fnr *FlowNodeRef) SetID(typ string, suffix interface{}) {
	switch typ {
	case "id":
		fnr.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (fnr FlowNodeRef) GetID() *string {
	return &fnr.ID
}

// StartEvent ...
type (

	// StartEvent ...
	StartEvent struct {
		ID             string     `xml:"id,attr,omitempty" json:"id"`
		Name           string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		IsInterrupting bool       `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
		Outgoing       []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TStartEvent ...
	TStartEvent struct {
		ID             string     `xml:"id,attr,omitempty" json:"id"`
		Name           string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		IsInterrupting bool       `xml:"isInterrupting,attr,omitempty" json:"isInterrupting,omitempty"`
		Outgoing       []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (se *StartEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	// startevent case MUST/SHOULD be used, when only one start event is used
	// or for the first start event, when more than one is set
	case "startevent":
		se.ID = fmt.Sprint("StartEvent_1")
	// event case MUST/SHOULD be used, when more than one start event is used
	case "event":
		se.ID = fmt.Sprintf("Event_%v", suffix)
	case "id":
		se.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (se StartEvent) GetID() *string {
	return &se.ID
}

// SetName ...
func (se *StartEvent) SetName(name string) {
	se.Name = name
}

// GetName ...
func (se StartEvent) GetName() *string {
	return &se.Name
}

// SetIsInterrupting ...
func (se *StartEvent) SetIsInterrupting(isInterrupting bool) {
	se.IsInterrupting = isInterrupting
}

// GetIsInterrupting ...
func (se StartEvent) GetIsInterrupting() *bool {
	return &se.IsInterrupting
}

// SetOutgoing ...
func (se *StartEvent) SetOutgoing(num int) {
	se.Outgoing = make([]Outgoing, num)
}

// GetOutgoing ...
func (se StartEvent) GetOutgoing(num int) *Outgoing {
	return &se.Outgoing[num]
}

// EndEvent ...
type (

	// EndEvent ...
	EndEvent struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	}

	// TEndEvent ...
	TEndEvent struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
	}
)

// SetID ...
func (ee *EndEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	case "event":
		ee.ID = fmt.Sprintf("Event_%v", suffix)
	case "id":
		ee.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (ee EndEvent) GetID() *string {
	return &ee.ID
}

// SetName ...
func (ee *EndEvent) SetName(name string) {
	ee.Name = name
}

// GetName ...
func (ee EndEvent) GetName() *string {
	return &ee.Name
}

// SetIncoming ...
func (ee *EndEvent) SetIncoming(num int) {
	ee.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (ee EndEvent) GetIncoming(num int) *Incoming {
	return &ee.Incoming[num]
}

// IntermediateCatchEvent ...
type (

	// IntermediateCatchEvent ...
	IntermediateCatchEvent struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TIntermediateCatchEvent ...
	TIntermediateCatchEvent struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (ice *IntermediateCatchEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	case "event":
		ice.ID = fmt.Sprintf("Event_%v", suffix)
	case "id":
		ice.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (ice IntermediateCatchEvent) GetID() *string {
	return &ice.ID
}

// SetName ...
func (ice *IntermediateCatchEvent) SetName(name string) {
	ice.Name = name
}

// GetName ...
func (ice IntermediateCatchEvent) GetName() *string {
	return &ice.Name
}

// SetIncoming ...
func (ice *IntermediateCatchEvent) SetIncoming(num int) {
	ice.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (ice IntermediateCatchEvent) GetIncoming(num int) *Incoming {
	return &ice.Incoming[num]
}

// SetOutgoing ...
func (ice *IntermediateCatchEvent) SetOutgoing(num int) {
	ice.Outgoing = make([]Outgoing, num)
}

// GetOutgoing ...
func (ice IntermediateCatchEvent) GetOutgoing(num int) *Outgoing {
	return &ice.Outgoing[num]
}

// IntermediateThrowEvent ...
type (

	// IntermediateThrowEvent ...
	IntermediateThrowEvent struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TIntermediateThrowEvent ...
	TIntermediateThrowEvent struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (ite *IntermediateThrowEvent) SetID(typ string, suffix interface{}) {
	switch typ {
	case "event":
		ite.ID = fmt.Sprintf("Event_%v", suffix)
	case "id":
		ite.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (ite IntermediateThrowEvent) GetID() *string {
	return &ite.ID
}

// SetName ...
func (ite *IntermediateThrowEvent) SetName(name string) {
	ite.Name = name
}

// GetName ...
func (ite IntermediateThrowEvent) GetName() *string {
	return &ite.Name
}

// SetIncoming ...
func (ite *IntermediateThrowEvent) SetIncoming(num int) {
	ite.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (ite IntermediateThrowEvent) GetIncoming(num int) *Incoming {
	return &ite.Incoming[num]
}

// SetOutgoing ...
func (ite *IntermediateThrowEvent) SetOutgoing(num int) {
	ite.Outgoing = make([]Outgoing, num)
}

// GetOutgoing ...
func (ite IntermediateThrowEvent) GetOutgoing(num int) *Outgoing {
	return &ite.Outgoing[num]
}

// InclusiveGateway ...
type (

	// InclusiveGateway ...
	InclusiveGateway struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TInclusiveGateway ...
	TInclusiveGateway struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (ig *InclusiveGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		ig.ID = fmt.Sprintf("Gateway_%v", suffix)
	case "id":
		ig.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (ig InclusiveGateway) GetID() *string {
	return &ig.ID
}

// SetName ...
func (ig *InclusiveGateway) SetName(name string) {
	ig.Name = name
}

// GetName ...
func (ig InclusiveGateway) GetName() *string {
	return &ig.Name
}

// SetIncoming ...
func (ig *InclusiveGateway) SetIncoming(num int) {
	ig.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (ig InclusiveGateway) GetIncoming(num int) *Incoming {
	return &ig.Incoming[num]
}

// SetOutgoing ...
func (ig *InclusiveGateway) SetOutgoing(num int) {
	ig.Outgoing = make([]Outgoing, num)
}

// GetOutgoing ...
func (ig InclusiveGateway) GetOutgoing(num int) *Outgoing {
	return &ig.Outgoing[num]
}

// ExclusiveGateway ...
type (

	// ExclusiveGateway ...
	ExclusiveGateway struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TExclusiveGateway ...
	TExclusiveGateway struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (eg *ExclusiveGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		eg.ID = fmt.Sprintf("Gateway_%v", suffix)
	case "id":
		eg.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (eg ExclusiveGateway) GetID() *string {
	return &eg.ID
}

// SetName ...
func (eg *ExclusiveGateway) SetName(name string) {
	eg.Name = name
}

// GetName ...
func (eg ExclusiveGateway) GetName() *string {
	return &eg.Name
}

// SetIncoming ...
func (eg *ExclusiveGateway) SetIncoming(num int) {
	eg.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (eg ExclusiveGateway) GetIncoming(num int) *Incoming {
	return &eg.Incoming[num]
}

// SetOutgoing ...
func (eg *ExclusiveGateway) SetOutgoing(num int) {
	eg.Outgoing = make([]Outgoing, num)
}

// GetOutgoing ...
func (eg ExclusiveGateway) GetOutgoing(num int) *Outgoing {
	return &eg.Outgoing[num]
}

// ParallelGateway ...
type (

	// ParallelGateway ...
	ParallelGateway struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TParallelGateway ...
	TParallelGateway struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (pg *ParallelGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		pg.ID = fmt.Sprintf("Gateway_%v", suffix)
	case "id":
		pg.ID = fmt.Sprintf("%s", suffix)

	}
}

// GetID ...
func (pg ParallelGateway) GetID() *string {
	return &pg.ID
}

// SetName ...
func (pg *ParallelGateway) SetName(name string) {
	pg.Name = name
}

// GetName ...
func (pg ParallelGateway) GetName() *string {
	return &pg.Name
}

// SetIncoming ...
func (eg *ParallelGateway) SetIncoming(num int) {
	eg.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (pg ParallelGateway) GetIncoming(num int) *Incoming {
	return &pg.Incoming[num]
}

// SetOutgoing ...
func (pg *ParallelGateway) SetOutgoing(num int) {
	pg.Outgoing = make([]Outgoing, num)
}

// GetOutgoing ...
func (pg ParallelGateway) GetOutgoing(num int) *Outgoing {
	return &pg.Outgoing[num]
}

// Gateway ...
type (

	// Gateway ...
	Gateway struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TGateway ...
	TGateway struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (gt *Gateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		gt.ID = fmt.Sprintf("Gateway_%v", suffix)
	case "id":
		gt.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (gt Gateway) GetID() *string {
	return &gt.ID
}

// SetName ...
func (gt *Gateway) SetName(name string) {
	gt.Name = name
}

// GetName ...
func (gt Gateway) GetName() *string {
	return &gt.Name
}

// SetIncoming ...
func (gt *Gateway) SetIncoming(num int) {
	gt.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (gt Gateway) GetIncoming(num int) *Incoming {
	return &gt.Incoming[num]
}

// SetOutgoing ...
func (gt *Gateway) SetOutgoing(num int) {
	gt.Outgoing = make([]Outgoing, num)
}

// GetOutgoing ...
func (gt Gateway) GetOutgoing(num int) *Outgoing {
	return &gt.Outgoing[num]
}

// Task ...
type (

	// Task ..
	Task struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TTask ..
	TTask struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (tk *Task) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		tk.ID = fmt.Sprintf("Activity_%v", suffix)
	case "id":
		tk.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (tk Task) GetID() *string {
	return &tk.ID
}

// SetName ...
func (tk *Task) SetName(name string) {
	tk.Name = name
}

// GetName ...
func (tk Task) GetName() *string {
	return &tk.Name
}

// SetIncoming ...
func (tk *Task) SetIncoming(num int) {
	tk.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (tk Task) GetIncoming(num int) *Incoming {
	return &tk.Incoming[num]
}

// SetOutgoing ...
func (tk *Task) SetOutgoing(num int) {
	tk.Outgoing = make([]Outgoing, num)
}

// GetOutgoing ...
func (tk Task) GetOutgoing(num int) *Outgoing {
	return &tk.Outgoing[num]
}

// ScriptTask ...
type (

	// ScriptTask ..
	ScriptTask struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TScriptTask ..
	TScriptTask struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (st *ScriptTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		st.ID = fmt.Sprintf("Activity_%v", suffix)
	case "id":
		st.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (st ScriptTask) GetID() *string {
	return &st.ID
}

// SetName ...
func (st *ScriptTask) SetName(name string) {
	st.Name = name
}

// GetName ...
func (st ScriptTask) GetName() *string {
	return &st.Name
}

// SetIncoming ...
func (st *ScriptTask) SetIncoming(num int) {
	st.Incoming = make([]Incoming, num)
}

func (st ScriptTask) GetIncoming(num int) *Incoming {
	return &st.Incoming[num]
}

// SetOutgoing ...
func (st *ScriptTask) SetOutgoing(num int) {
	st.Outgoing = make([]Outgoing, num)
}

func (st ScriptTask) GetOutgoing(num int) *Outgoing {
	return &st.Outgoing[num]
}

// UserTask ...
type (

	// UserTask ...
	UserTask struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TUserTask ...
	TUserTask struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (ut *UserTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		ut.ID = fmt.Sprintf("Activity_%v", suffix)
	case "id":
		ut.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (ut UserTask) GetID() *string {
	return &ut.ID
}

// SetName ...
func (ut *UserTask) SetName(name string) {
	ut.Name = name
}

// GetName ...
func (ut UserTask) GetName() *string {
	return &ut.Name
}

// SetIncoming ...
func (ut *UserTask) SetIncoming(num int) {
	ut.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (ut UserTask) GetIncoming(num int) *Incoming {
	return &ut.Incoming[num]
}

// SetOutgoing ...
func (ut *UserTask) SetOutgoing(num int) {
	ut.Outgoing = make([]Outgoing, num)
}

func (ut UserTask) GetOutgoing(num int) *Outgoing {
	return &ut.Outgoing[num]
}

// ServiceTask ...
type (

	// ServiceTask ...
	ServiceTask struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	}

	// TServiceTask ...
	TServiceTask struct {
		ID       string     `xml:"id,attr,omitempty" json:"id"`
		Name     string     `xml:"name,attr,omitempty" json:"name,omitempty"`
		Incoming []Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
		Outgoing []Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	}
)

// SetID ...
func (st *ServiceTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		st.ID = fmt.Sprintf("Activity_%v", suffix)
	case "id":
		st.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (st ServiceTask) GetID() *string {
	return &st.ID
}

// SetName ...
func (st *ServiceTask) SetName(name string) {
	st.Name = name
}

// GetName ...
func (st ServiceTask) GetName() *string {
	return &st.Name
}

// SetIncoming ...
func (st *ServiceTask) SetIncoming(num int) {
	st.Incoming = make([]Incoming, num)
}

// GetIncoming ...
func (st ServiceTask) GetIncoming(num int) *Incoming {
	return &st.Incoming[num]
}

// SetOutgoing ...
func (st *ServiceTask) SetOutgoing(num int) {
	st.Outgoing = make([]Outgoing, num)
}

func (st ServiceTask) GetOutgoing(num int) *Outgoing {
	return &st.Outgoing[num]
}

// SequenceFlow ...
type (

	// SequenceFlow ...
	SequenceFlow struct {
		ID        string `xml:"id,attr,omitempty" json:"id"`
		Name      string `xml:"name,attr,omitempty" json:"name,omitempty"`
		SourceRef string `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
		TargetRef string `xml:"targetRef,attr" json:"targetRef,omitempty"`
	}

	// TSequenceFlow ...
	TSequenceFlow struct {
		ID        string `xml:"id,attr,omitempty" json:"id"`
		Name      string `xml:"name,attr,omitempty" json:"name,omitempty"`
		SourceRef string `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
		TargetRef string `xml:"targetRef,attr" json:"targetRef,omitempty"`
	}
)

// SetID ...
func (sf *SequenceFlow) SetID(typ string, suffix interface{}) {
	switch typ {
	case "flow":
		sf.ID = fmt.Sprintf("Flow_%v", suffix)
	case "id":
		sf.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (sf SequenceFlow) GetID() *string {
	return &sf.ID
}

// SetName ...
func (sf *SequenceFlow) SetName(name string) {
	sf.Name = name
}

// GetName ...
func (sf SequenceFlow) GetName() *string {
	return &sf.Name
}

// SetSourceRef ...
func (sf *SequenceFlow) SetSourceRef(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		sf.SourceRef = fmt.Sprintf("Activity_%s", suffix)
	case "event":
		sf.SourceRef = fmt.Sprintf("Event_%s", suffix)
	case "id":
		sf.SourceRef = fmt.Sprintf("%s", suffix)
	case "startevent":
		sf.SourceRef = fmt.Sprintf("StartEvent_%s", suffix)
	}
}

// GetSourceRef ...
func (sf SequenceFlow) GetSourceRef() *string {
	return &sf.SourceRef
}

// SetTargetRef ...
func (sf *SequenceFlow) SetTargetRef(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		sf.TargetRef = fmt.Sprintf("Activity_%s", suffix)
	case "event":
		sf.TargetRef = fmt.Sprintf("Event_%s", suffix)
	case "id":
		sf.TargetRef = fmt.Sprintf("%s", suffix)
	case "startevent":
		sf.TargetRef = fmt.Sprintf("StartEvent_%s", suffix)
	}
}

// GetTargetRef ...
func (sf SequenceFlow) GetTargetRef() *string {
	return &sf.TargetRef
}

// Incoming ...
type Incoming struct {
	Flow string `xml:",innerxml" json:"flow"`
}

// SetFlow ...
func (in *Incoming) SetFlow(suffix interface{}) {
	stringArg := getLastPart(suffix.(string)) // Note: using type assertion.
	in.Flow = fmt.Sprintf("Flow_%s", stringArg)
}

// GetFlow ...
func (in Incoming) GetFlow() *string {
	return &in.Flow
}

// Outgoing
type Outgoing struct {
	Flow string `xml:",innerxml" json:"flow"`
}

// SetFlow ...
func (out *Outgoing) SetFlow(suffix interface{}) {
	stringArg := getLastPart(suffix.(string)) // Note: using type assertion.
	out.Flow = fmt.Sprintf("Flow_%s", stringArg)
}

// GetFlow ...
func (out Outgoing) GetFlow() *string {
	return &out.Flow
}

/*
 * @helper
 */

// getLastPart returns the last part of a string, which is separated by an underscore.
// If no Snake-Case-String, return the original string.
func getLastPart(input string) string {
	if strings.Contains(input, "_") {
		parts := strings.Split(input, "_")
		// return the last part of the string
		return parts[len(parts)-1]
	}
	// no Snake-Case-String. Return the original string.
	return input
}
