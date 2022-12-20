package activities

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// ScriptTaskRepository ...
type ScriptTaskRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)

	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaJobPriority(priority int)

	SetDocumentation()
	SetExtensionElements()
	SetIncoming(num int)
	SetOutgoing(num int)

	GetID() *string
	GetName() *string

	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaJobPriority() *int

	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *camunda.ExtensionElements
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing
}

// ScriptTask ...
type ScriptTask struct {
	ID                 string                      `xml:"id,attr,omitempty" json:"id"`
	Name               string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing           `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TScriptTask ...
type TScriptTask struct {
	ID                string                       `xml:"id,attr,omitempty" json:"id"`
	Name              string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore       bool                         `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter        bool                         `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority       int                          `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation     []attributes.Documentation   `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming            `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing            `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewScriptTask() ScriptTaskRepository {
	return &ScriptTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (scriptTask *ScriptTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		scriptTask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		scriptTask.ID = fmt.Sprintf("Activity_%s", suffix)
		break
	}
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

/*** Make Elements ***/

// SetDocumentation ...
func (scriptTask *ScriptTask) SetDocumentation() {
	scriptTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (scriptTask *ScriptTask) SetExtensionElements() {
	scriptTask.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (scriptTask *ScriptTask) SetIncoming(num int) {
	scriptTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (scriptTask *ScriptTask) SetOutgoing(num int) {
	scriptTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (scriptTask ScriptTask) GetID() *string {
	return &scriptTask.ID
}

// GetName ...
func (scriptTask ScriptTask) GetName() *string {
	return &scriptTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (scriptTask ScriptTask) GetCamundaAsyncBefore() *bool {
	return &scriptTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (scriptTask ScriptTask) GetCamundaAsyncAfter() *bool {
	return &scriptTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (scriptTask ScriptTask) GetCamundaJobPriority() *int {
	return &scriptTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (scriptTask ScriptTask) GetDocumentation() *attributes.Documentation {
	return &scriptTask.Documentation[0]
}

// GetExtensionElements ...
func (scriptTask ScriptTask) GetExtensionElements() *camunda.ExtensionElements {
	return &scriptTask.ExtensionElements[0]
}

// GetIncoming ...
func (scriptTask ScriptTask) GetIncoming(num int) *marker.Incoming {
	return &scriptTask.Incoming[num]
}

// GetOutgoing ...
func (scriptTask ScriptTask) GetOutgoing(num int) *marker.Outgoing {
	return &scriptTask.Outgoing[num]
}
