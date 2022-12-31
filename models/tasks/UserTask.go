package tasks

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/marker"
)

func NewUserTask() UserTaskRepository {
	return &UserTask{}
}

/*
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

/*** Attributes ***/

// SetDocumentation ...
func (utask *UserTask) SetDocumentation() {
	utask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (utask *UserTask) SetExtensionElements() {
	utask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (utask *UserTask) SetIncoming(num int) {
	utask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (utask *UserTask) SetOutgoing(num int) {
	utask.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (utask UserTask) GetID() compulsion.STR_PTR {
	return &utask.ID
}

// GetName ...
func (utask UserTask) GetName() compulsion.STR_PTR {
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

/*** Attributes ***/

// GetDocumentation ...
func (utask UserTask) GetDocumentation() *attributes.Documentation {
	return &utask.Documentation[0]
}

// GetExtensionElements ...
func (utask UserTask) GetExtensionElements() *attributes.ExtensionElements {
	return &utask.ExtensionElements[0]
}

/*** Marker ***/

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
