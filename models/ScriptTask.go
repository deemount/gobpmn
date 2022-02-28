package models

import "fmt"

// ScriptTaskRepository ...
type ScriptTaskRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// ScriptTask ...
type ScriptTask struct {
	ID                string              `xml:"id,attr,omitempty"`
	Name              string              `xml:"name,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
	Incoming          []Incoming          `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []Outgoing          `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
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

/* Elements */

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
