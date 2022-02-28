package models

import "fmt"

// BusinessRuleTaskRepository ...
type UserTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// UserTask ...
type UserTask struct {
	ID                  string              `xml:"id,attr" json:"id"`
	Name                string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaFormKey      string              `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaAsyncBef     bool                `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAft     bool                `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority  int                 `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	CamundaAssignee     string              `xml:"camunda:assignee,attr,omitempty" json:"assignee,omitempty"`
	CamundaCandidUsers  string              `xml:"camunda:candidateUsers,attr,omitempty" json:"candidateUsers,omitempty"`
	CamundaCandidGroups string              `xml:"camunda:candidateGroups,attr,omitempty" json:"candidateGroups,omitempty"`
	CamundaDueDate      string              `xml:"camunda:dueDate,attr,omitempty" json:"dueDate,omitempty"`
	CamundaFollowUpDate string              `xml:"camunda:followUpDate,attr,omitempty" json:"followUpDate,omitempty"`
	CamundaPriority     int                 `xml:"camunda:priority,attr,omitempty" json:"priority,omitempty"`
	Documentation       []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements   []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming            []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing            []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (utask *UserTask) SetID(suffix string) {
	utask.ID = fmt.Sprintf("Activity_%s", suffix)
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

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (utask *UserTask) SetDocumentation() {
	utask.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (utask *UserTask) SetExtensionElements() {
	utask.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (utask *UserTask) SetIncoming(num int) {
	utask.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (utask *UserTask) SetOutgoing(num int) {
	utask.Outgoing = make([]Outgoing, num)
}
