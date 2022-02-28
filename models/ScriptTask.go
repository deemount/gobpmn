package models

import "fmt"

// ScriptTaskRepository ...
type ScriptTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ScriptTask ...
type ScriptTask struct {
	ID                 string              `xml:"id,attr,omitempty" json:"id"`
	Name               string              `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                 `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (scriptTask *ScriptTask) SetID(suffix string) {
	scriptTask.ID = fmt.Sprintf("Activity_%s", suffix)
}

// SetName ...
func (scriptTask *ScriptTask) SetName(name string) {
	scriptTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (scriptTask *ScriptTask) SetCamundaAsyncBefore(asyncBefore bool) {
	scriptTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (scriptTask *ScriptTask) SetCamundaAsyncAfter(asyncAfter bool) {
	scriptTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (scriptTask *ScriptTask) SetCamundaJobPriority(priority int) {
	scriptTask.CamundaJobPriority = priority
}

/* Elements */

// SetDocumentation ...
func (scriptTask *ScriptTask) SetDocumentation() {
	scriptTask.Documentation = make([]Documentation, 1)
}

// SetExtensionElements ...
func (scriptTask *ScriptTask) SetExtensionElements() {
	scriptTask.ExtensionElements = make([]ExtensionElements, 1)
}

// SetIncoming ...
func (scriptTask *ScriptTask) SetIncoming(num int) {
	scriptTask.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (scriptTask *ScriptTask) SetOutgoing(num int) {
	scriptTask.Outgoing = make([]Outgoing, num)
}
