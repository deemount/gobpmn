package tasks

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/data"
	"github.com/deemount/gobpmn/models/marker"
)

type TasksID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type TasksName interface {
	SetName(name string)
	GetName() STR_PTR
}

type TasksBaseAttributes interface {
	TasksID
	TasksName
}

type TasksMarkers interface {
	SetIncoming(num int)
	GetIncoming(num int) *marker.Incoming
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
}

type TasksCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
}

type TasksBaseCoreElements interface {
	SetDocumentation()
	SetExtensionElements()
	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *attributes.ExtensionElements
}

type TasksBase interface {
	TasksBaseAttributes
	TasksMarkers
	TasksCamundaBase
	TasksBaseCoreElements
}

// TasksElementsRepository ...
type TasksElementsRepository interface {
	SetBusinessRuleTask(num int)
	GetBusinessRuleTask(num int) *BusinessRuleTask
	SetTask(num int)
	GetTask(num int) *Task
	SetUserTask(num int)
	GetUserTask(num int) *UserTask
	SetManualTask(num int)
	GetManualTask(num int) *ManualTask
	SetReceiveTask(num int)
	GetReceiveTask(num int) *ReceiveTask
	SetScriptTask(num int)
	GetScriptTask(num int) *ScriptTask
	SetSendTask(num int)
	GetSendTask(num int) *SendTask
	SetServiceTask(num int)
	GetServiceTask(num int) *ServiceTask
}

// BusinessRuleTaskRepository ...
type BusinessRuleTaskRepository interface {
	TasksBase
	SetCamundaClass(class string)
	GetCamundaClass() *string
	String() string
}

// ManualTaskRepository ...
type ManualTaskRepository interface {
	TasksBase
	String() string
}

// ReceiveTaskRepository ...
type ReceiveTaskRepository interface {
	TasksBase
	SetMessageRef(suffix string)
	GetMessageRef(suffix string) *string
	String() string
}

// ScriptTaskRepository ...
type ScriptTaskRepository interface {
	TasksBase
	String() string
}

// SendTaskRepository ...
type SendTaskRepository interface {
	TasksBase
	String() string
}

// ServiceTaskRepository ...
type ServiceTaskRepository interface {
	TasksBase
	String() string
}

// TaskRepository ...
type TaskRepository interface {
	TasksBase
	SetDataInputAssociation(num int)
	GetDataInputAssociation(num int) *data.DataInputAssociation
	String() string
}

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
