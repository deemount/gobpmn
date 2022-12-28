package tasks

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/data"
	"github.com/deemount/gobpmn/models/marker"
)

type Tasks struct {
	BusinessRuleTask BUSINESS_RULE_TASK_SLC `xml:"bpmn:businessRuleTask,omitempty" json:"businessRuleTask,omitempty" csv:"-"`
	Task             TASK_SLC               `xml:"bpmn:task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask         USER_TASK_SLC          `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask       MANUAL_TASK_SLC        `xml:"bpmn:manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask      RECEIVE_TASK_SLC       `xml:"bpmn:receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask       SCRIPT_TASK_SLC        `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask         SEND_TASK_SLC          `xml:"bpmn:sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask      SERVICE_TASK_SLC       `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
}

type TTasks struct {
	BusinessRuleTask BUSINESS_RULE_TASK_SLC `xml:"businessRuleTask,omitempty" json:"businessRuleTask,omitempty" csv:"-"`
	Task             TASK_SLC               `xml:"task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask         USER_TASK_SLC          `xml:"userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask       MANUAL_TASK_SLC        `xml:"manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask      RECEIVE_TASK_SLC       `xml:"receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask       SCRIPT_TASK_SLC        `xml:"scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask         SEND_TASK_SLC          `xml:"sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask      SERVICE_TASK_SLC       `xml:"serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
}

// BusinessRuleTask ...
type BusinessRuleTask struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	CamundaClass string `xml:"camunda:class,attr,omitempty" json:"class,omitempty"`
}

// TBusinessRuleTask ...
type TBusinessRuleTask struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`
}

// ManualTask ...
type ManualTask struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
}

// TManualTask ...
type TManualTask struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
}

// ReceiveTask ...
type ReceiveTask struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	MessageRef string `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
}

// TReceiveTask ...
type TReceiveTask struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	MessageRef string `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
}

// ScriptTask ...
type ScriptTask struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
}

// TScriptTask ...
type TScriptTask struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
}

// SendTask ...
type SendTask struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
}

// TSendTask ...
type TSendTask struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
}

// ServiceTask ...
type ServiceTask struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
}

// TServiceTask ...
type TServiceTask struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
}

// Task ...
type Task struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	Property             []attributes.Property       `xml:"bpmn:property,omitempty" json:"property,omitempty"`
	DataInputAssociation []data.DataInputAssociation `xml:"bpmn:dataInputAssociation,omitempty" json:"dataInputAssociation,omitempty"`
}

// TTask ...
type TTask struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	Property             []attributes.Property       `xml:"property,omitempty" json:"property,omitempty"`
	DataInputAssociation []data.DataInputAssociation `xml:"dataInputAssociation,omitempty" json:"dataInputAssociation,omitempty"`
}

// UserTask ...
type UserTask struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
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
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	FormKey         string `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	Assignee        string `xml:"assignee,attr,omitempty" json:"assignee,omitempty"`
	CandidateUsers  string `xml:"candidateUsers,attr,omitempty" json:"candidateUsers,omitempty"`
	CandidateGroups string `xml:"candidateGroups,attr,omitempty" json:"candidateGroups,omitempty"`
	DueDate         string `xml:"dueDate,attr,omitempty" json:"dueDate,omitempty"`
	FollowUpDate    string `xml:"followUpDate,attr,omitempty" json:"followUpDate,omitempty"`
	Priority        int    `xml:"priority,attr,omitempty" json:"priority,omitempty"`
}
