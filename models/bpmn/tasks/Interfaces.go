package tasks

import (
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

type TasksBaseAttributes interface {
	impl.IFBaseID
	impl.IFBaseName
}

type TasksMarkers interface {
	SetIncoming(num int)
	GetIncoming(num int) *marker.Incoming
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
}

type TasksCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() impl.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() impl.BOOL_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() impl.INT_PTR
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

/*
 * @Repositories
 */

// TasksElementsRepository ...
type TasksElementsRepository interface {
	SetBusinessRuleTask(num int)
	GetBusinessRuleTask(num int) BUSINESS_RULE_TASK_PTR
	SetTask(num int)
	GetTask(num int) TASK_PTR
	SetUserTask(num int)
	GetUserTask(num int) USER_TASK_PTR
	SetManualTask(num int)
	GetManualTask(num int) MANUAL_TASK_PTR
	SetReceiveTask(num int)
	GetReceiveTask(num int) RECEIVE_TASK_PTR
	SetScriptTask(num int)
	GetScriptTask(num int) SCRIPT_TASK_PTR
	SetSendTask(num int)
	GetSendTask(num int) SEND_TASK_PTR
	SetServiceTask(num int)
	GetServiceTask(num int) SERVICE_TASK_PTR
}

// BusinessRuleTaskRepository ...
type BusinessRuleTaskRepository interface {
	TasksBase
	SetCamundaClass(class string)
	GetCamundaClass() impl.STR_PTR
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
	GetMessageRef(suffix string) impl.STR_PTR
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
	GetDataInputAssociation(num int) *flow.DataInputAssociation
	String() string
}

// UserTaskRepository ...
type UserTaskRepository interface {
	TasksBase
	SetCamundaFormKey(formKey string)
	GetCamundaFormKey() impl.STR_PTR
	SetCamundaAssignee(assignee string)
	GetCamundaAssignee() impl.STR_PTR
	SetCamundaCandidateUsers(users string)
	GetCamundaCandidateUsers() impl.STR_PTR
	SetCamundaCandidateGroups(groups string)
	GetCamundaCandidateGroups() impl.STR_PTR
	SetCamundaDueDate(due string)
	GetCamundaDueDate() impl.STR_PTR
	SetCamundaFollowUpDate(followUp string)
	GetCamundaFollowUpDate() impl.STR_PTR
	SetCamundaPriority(priority int)
	GetCamundaPriority() impl.INT_PTR
	String() string
}
