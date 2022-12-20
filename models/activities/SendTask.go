package activities

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// SendTaskRepository ...
type SendTaskRepository interface {
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

// SendTask ...
type SendTask struct {
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

// TSendTask ...
type TSendTask struct {
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

func NewSendTask() SendTaskRepository {
	return &SendTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (sendTask *SendTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		sendTask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		sendTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (sendTask *SendTask) SetName(name string) {
	sendTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (sendTask *SendTask) SetCamundaAsyncBefore(asyncBefore bool) {
	sendTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (sendTask *SendTask) SetCamundaAsyncAfter(asyncAfter bool) {
	sendTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (sendTask *SendTask) SetCamundaJobPriority(priority int) {
	sendTask.CamundaJobPriority = priority
}

/* Elements */

// SetDocumentation ...
func (sendTask *SendTask) SetDocumentation() {
	sendTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (sendTask *SendTask) SetExtensionElements() {
	sendTask.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (sendTask *SendTask) SetIncoming(num int) {
	sendTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (sendTask *SendTask) SetOutgoing(num int) {
	sendTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (sendTask SendTask) GetID() *string {
	return &sendTask.ID
}

// GetName ...
func (sendTask SendTask) GetName() *string {
	return &sendTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (sendTask SendTask) GetCamundaAsyncBefore() *bool {
	return &sendTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (sendTask SendTask) GetCamundaAsyncAfter() *bool {
	return &sendTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (sendTask SendTask) GetCamundaJobPriority() *int {
	return &sendTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (sendTask SendTask) GetDocumentation() *attributes.Documentation {
	return &sendTask.Documentation[0]
}

// GetExtensionElements ...
func (sendTask SendTask) GetExtensionElements() *camunda.ExtensionElements {
	return &sendTask.ExtensionElements[0]
}

// GetIncoming ...
func (sendTask SendTask) GetIncoming(num int) *marker.Incoming {
	return &sendTask.Incoming[num]
}

// GetOutgoing ...
func (sendTask SendTask) GetOutgoing(num int) *marker.Outgoing {
	return &sendTask.Outgoing[num]
}
