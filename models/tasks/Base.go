package tasks

import "github.com/deemount/gobpmn/models/marker"

type STR_PTR *string

const TASK = "task"
const TASK_SCRIPT = "scriptTask"
const TASK_SERVICE = "serviceTask"
const TASK_BUSINESS_RULE = "businessRuleTask"
const TASK_MANUAL_TASK = "manualTask"
const TASK_USER_TASK = "userTask"
const TASK_SEND_TASK = "sendTask"
const TASK_RECEIVE_TASK = "receiveTask"

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

type TasksBase interface {
	TasksBaseAttributes
	TasksMarkers
	TasksCamundaBase
}
