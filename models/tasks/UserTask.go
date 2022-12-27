package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/marker"
)

// UserTaskRepository ...
type UserTaskRepository interface {
	TasksBase

	SetCamundaFormKey(formKey string)
	GetCamundaFormKey() *string
	SetCamundaAssignee(assignee string)
	GetCamundaAssignee() *string
	SetCamundaCandidateUsers(users string)
	GetCamundaCandidateUsers() *string
	SetCamundaCandidateGroups(groups string)
	GetCamundaCandidateGroups() *string
	SetCamundaDueDate(due string)
	GetCamundaDueDate() *string
	SetCamundaFollowUpDate(followUp string)
	GetCamundaFollowUpDate() *string
	SetCamundaPriority(priority int)
	GetCamundaPriority() *int

	String() string
}

// UserTask ...
type UserTask struct {
	compulsion.CompulsionCoreAttributes
	compulsion.CompulsionCoreElements
	compulsion.CompulsionCamundaCoreAttributes
	compulsion.CompulsionCoreIncomingOutgoing
	CamundaFormKey         string `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaAssignee        string `xml:"camunda:assignee,attr,omitempty" json:"assignee,omitempty"`
	CamundaCandidateUsers  string `xml:"camunda:candidateUsers,attr,omitempty" json:"candidateUsers,omitempty"`
	CamundaCandidateGroups string `xml:"camunda:candidateGroups,attr,omitempty" json:"candidateGroups,omitempty"`
	CamundaDueDate         string `xml:"camunda:dueDate,attr,omitempty" json:"dueDate,omitempty"`
	CamundaFollowUpDate    string `xml:"camunda:followUpDate,attr,omitempty" json:"followUpDate,omitempty"`
	CamundaPriority        int    `xml:"camunda:priority,attr,omitempty" json:"priority,omitempty"`
}

// TUserTask ...
type TUserTask struct {
	compulsion.CompulsionCoreAttributes
	compulsion.TCompulsionCoreElements
	compulsion.TCompulsionCamundaCoreAttributes
	compulsion.CompulsionCoreIncomingOutgoing
	FormKey         string `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	Assignee        string `xml:"assignee,attr,omitempty" json:"assignee,omitempty"`
	CandidateUsers  string `xml:"candidateUsers,attr,omitempty" json:"candidateUsers,omitempty"`
	CandidateGroups string `xml:"candidateGroups,attr,omitempty" json:"candidateGroups,omitempty"`
	DueDate         string `xml:"dueDate,attr,omitempty" json:"dueDate,omitempty"`
	FollowUpDate    string `xml:"followUpDate,attr,omitempty" json:"followUpDate,omitempty"`
	Priority        int    `xml:"priority,attr,omitempty" json:"priority,omitempty"`
}

func NewUserTask() UserTaskRepository {
	return &UserTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (utask *UserTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		utask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		utask.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (utask *UserTask) SetName(name string) {
	utask.Name = name
}

/** Camunda **/

// SetCamundaFormKey ...
func (utask *UserTask) SetCamundaFormKey(formKey string) {
	utask.CamundaFormKey = formKey
}

// SetCamundaAsyncBefore ...
func (utask *UserTask) SetCamundaAsyncBefore(asyncBefore bool) {
	utask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (utask *UserTask) SetCamundaAsyncAfter(asyncAfter bool) {
	utask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (utask *UserTask) SetCamundaJobPriority(priority int) {
	utask.CamundaJobPriority = priority
}

// SetCamundaAssignee ...
func (utask *UserTask) SetCamundaAssignee(assignee string) {
	utask.CamundaAssignee = assignee
}

// SetCamundaCandidUsers ...
func (ut *UserTask) SetCamundaCandidateUsers(users string) {
	ut.CamundaCandidateUsers = users
}

// SetCamundaCandidGroups ...
func (ut *UserTask) SetCamundaCandidateGroups(groups string) {
	ut.CamundaCandidateGroups = groups
}

// SetCamundaDueDate ...
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask *UserTask) SetCamundaDueDate(due string) {
	utask.CamundaDueDate = due
}

// SetCamundaFollowUpDate
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask *UserTask) SetCamundaFollowUpDate(followUp string) {
	utask.CamundaFollowUpDate = followUp
}

// SetCamundaPriority ...
func (utask *UserTask) SetCamundaPriority(priority int) {
	utask.CamundaPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (utask *UserTask) SetDocumentation() {
	utask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (utask *UserTask) SetExtensionElements() {
	utask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (utask *UserTask) SetIncoming(num int) {
	utask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (utask *UserTask) SetOutgoing(num int) {
	utask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (utask UserTask) GetID() STR_PTR {
	return &utask.ID
}

// GetName ...
func (utask UserTask) GetName() STR_PTR {
	return &utask.Name
}

/** Camunda **/

// GetCamundaFormKey ...
func (utask UserTask) GetCamundaFormKey() *string {
	return &utask.CamundaFormKey
}

// GetCamundaAsyncBefore ...
func (utask UserTask) GetCamundaAsyncBefore() *bool {
	return &utask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (utask UserTask) GetCamundaAsyncAfter() *bool {
	return &utask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (utask UserTask) GetCamundaJobPriority() *int {
	return &utask.CamundaJobPriority
}

// GetCamundaAssignee ...
func (utask UserTask) GetCamundaAssignee() *string {
	return &utask.CamundaAssignee
}

// GetCamundaCandidUsers ...
func (ut UserTask) GetCamundaCandidateUsers() *string {
	return &ut.CamundaCandidateUsers
}

// GetCamundaCandidGroups ...
func (ut UserTask) GetCamundaCandidateGroups() *string {
	return &ut.CamundaCandidateGroups
}

// GetCamundaDueDate ...
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask UserTask) GetCamundaDueDate() *string {
	return &utask.CamundaDueDate
}

// GetCamundaFollowUpDate
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask UserTask) GetCamundaFollowUpDate() *string {
	return &utask.CamundaFollowUpDate
}

// GetCamundaPriority ...
func (utask UserTask) GetCamundaPriority() *int {
	return &utask.CamundaPriority
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (utask UserTask) GetDocumentation() DOCUMENTATION_PTR {
	return &utask.Documentation[0]
}

// GetExtensionElements ...
func (utask UserTask) GetExtensionElements() EXTENSION_ELEMENTS_PTR {
	return &utask.ExtensionElements[0]
}

// GetIncoming ...
func (utask UserTask) GetIncoming(num int) *marker.Incoming {
	return &utask.Incoming[num]
}

// GetOutgoing ...
func (utask UserTask) GetOutgoing(num int) *marker.Outgoing {
	return &utask.Outgoing[num]
}

/*
 * Default String
 */

// String ...
func (utask UserTask) String() string {
	return fmt.Sprintf("id=%v, name=%v", utask.ID, utask.Name)
}
