package activities

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// UserTaskRepository ...
type UserTaskRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)

	SetCamundaFormKey(formKey string)
	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaJobPriority(priority int)
	SetCamundaAssignee(assignee string)
	SetCamundaCandidUsers(users string)
	SetCamundaCandidGroups(groups string)
	SetCamundaDueDate(due string)
	SetCamundaFollowUpDate(followUp string)
	SetCamundaPriority(priority int)

	SetDocumentation()
	SetExtensionElements()
	SetIncoming(num int)
	SetOutgoing(num int)

	GetID() *string
	GetName() *string

	GetCamundaFormKey() *string
	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaJobPriority() *int
	GetCamundaAssignee() *string
	GetCamundaCandidUsers() *string
	GetCamundaCandidGroups() *string
	GetCamundaDueDate() *string
	GetCamundaFollowUpDate() *string
	GetCamundaPriority() *int
}

// UserTask ...
type UserTask struct {
	ID                  string                      `xml:"id,attr" json:"id"`
	Name                string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaFormKey      string                      `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaAsyncBef     bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAft     bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority  int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	CamundaAssignee     string                      `xml:"camunda:assignee,attr,omitempty" json:"assignee,omitempty"`
	CamundaCandidUsers  string                      `xml:"camunda:candidateUsers,attr,omitempty" json:"candidateUsers,omitempty"`
	CamundaCandidGroups string                      `xml:"camunda:candidateGroups,attr,omitempty" json:"candidateGroups,omitempty"`
	CamundaDueDate      string                      `xml:"camunda:dueDate,attr,omitempty" json:"dueDate,omitempty"`
	CamundaFollowUpDate string                      `xml:"camunda:followUpDate,attr,omitempty" json:"followUpDate,omitempty"`
	CamundaPriority     int                         `xml:"camunda:priority,attr,omitempty" json:"priority,omitempty"`
	Documentation       []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements   []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming            []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing            []marker.Outgoing           `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TUserTask ...
type TUserTask struct {
	ID                string                      `xml:"id,attr" json:"id"`
	Name              string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	FormKey           string                      `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	AsyncBef          bool                        `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAft          bool                        `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority       int                         `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Assignee          string                      `xml:"assignee,attr,omitempty" json:"assignee,omitempty"`
	CandidUsers       string                      `xml:"candidateUsers,attr,omitempty" json:"candidateUsers,omitempty"`
	CandidGroups      string                      `xml:"candidateGroups,attr,omitempty" json:"candidateGroups,omitempty"`
	DueDate           string                      `xml:"dueDate,attr,omitempty" json:"dueDate,omitempty"`
	FollowUpDate      string                      `xml:"followUpDate,attr,omitempty" json:"followUpDate,omitempty"`
	Priority          int                         `xml:"priority,attr,omitempty" json:"priority,omitempty"`
	Documentation     []attributes.Documentation  `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.ExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming           `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing           `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
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
	utask.CamundaAsyncBef = asyncBefore
}

// SetCamundaAsyncAfter ...
func (utask *UserTask) SetCamundaAsyncAfter(asyncAfter bool) {
	utask.CamundaAsyncBef = asyncAfter
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
func (ut *UserTask) SetCamundaCandidUsers(users string) {
	ut.CamundaCandidUsers = users
}

// SetCamundaCandidGroups ...
func (ut *UserTask) SetCamundaCandidGroups(groups string) {
	ut.CamundaCandidGroups = groups
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
	utask.ExtensionElements = make([]camunda.ExtensionElements, 1)
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
func (utask UserTask) GetID() *string {
	return &utask.ID
}

// GetName ...
func (utask UserTask) GetName() *string {
	return &utask.Name
}

/** Camunda **/

// GetCamundaFormKey ...
func (utask UserTask) GetCamundaFormKey() *string {
	return &utask.CamundaFormKey
}

// GetCamundaAsyncBefore ...
func (utask UserTask) GetCamundaAsyncBefore() *bool {
	return &utask.CamundaAsyncBef
}

// GetCamundaAsyncAfter ...
func (utask UserTask) GetCamundaAsyncAfter() *bool {
	return &utask.CamundaAsyncBef
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
func (ut UserTask) GetCamundaCandidUsers() *string {
	return &ut.CamundaCandidUsers
}

// GetCamundaCandidGroups ...
func (ut UserTask) GetCamundaCandidGroups() *string {
	return &ut.CamundaCandidGroups
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
