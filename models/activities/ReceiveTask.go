package activities

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// ReceiveTaskRepository ...
type ReceiveTaskRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)
	SetMessageRef(suffix string)

	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaJobPriority(priority int)

	SetExtensionElements()
	SetIncoming(num int)
	SetOutgoing(num int)

	GetID() *string
	GetName() *string
	GetMessageRef(suffix string) *string

	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaJobPriority() *int

	GetExtensionElements() *camunda.ExtensionElements
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing
}

// ReceiveTask ...
type ReceiveTask struct {
	ID                 string                      `xml:"id,attr" json:"id,omitempty"`
	Name               string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	MessageRef         string                      `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
	CamundaAsyncBefore bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,attr,omitempty"`
	CamundaAsyncAfter  bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,attr,omitempty"`
	CamundaJobPriority int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,attr,omitempty"`
	Documentation      []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing           `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TReceiveTask ...
type TReceiveTask struct {
	ID                string                       `xml:"id,attr" json:"id,omitempty"`
	Name              string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	MessageRef        string                       `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
	AsyncBefore       bool                         `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,attr,omitempty"`
	AsyncAfter        bool                         `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,attr,omitempty"`
	JobPriority       int                          `xml:"jobPriority,attr,omitempty" json:"jobPriority,attr,omitempty"`
	Documentation     []attributes.Documentation   `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming            `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing            `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewReceiveTask() ReceiveTaskRepository {
	return &ReceiveTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (receiveTask *ReceiveTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		receiveTask.ID = fmt.Sprintf("Activity_%s", suffix)
		break
	case "id":
		receiveTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (receiveTask *ReceiveTask) SetName(name string) {
	receiveTask.Name = name
}

// SetMessageRef ...
func (receiveTask *ReceiveTask) SetMessageRef(suffix string) {
	receiveTask.MessageRef = fmt.Sprintf("Message_%s", suffix)
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (receiveTask *ReceiveTask) SetCamundaAsyncBefore(asyncBefore bool) {
	receiveTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (receiveTask *ReceiveTask) SetCamundaAsyncAfter(asyncAfter bool) {
	receiveTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (receiveTask *ReceiveTask) SetCamundaJobPriority(priority int) {
	receiveTask.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetDocumentation ...
func (receiveTask *ReceiveTask) SetDocumentation() {
	receiveTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (receiveTask *ReceiveTask) SetExtensionElements() {
	receiveTask.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (receiveTask *ReceiveTask) SetIncoming(num int) {
	receiveTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (receiveTask *ReceiveTask) SetOutgoing(num int) {
	receiveTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (receiveTask ReceiveTask) GetID() *string {
	return &receiveTask.ID
}

// GetName ...
func (receiveTask ReceiveTask) GetName() *string {
	return &receiveTask.Name
}

// GetMessageRef ...
func (receiveTask ReceiveTask) GetMessageRef(suffix string) *string {
	return &receiveTask.MessageRef
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (receiveTask ReceiveTask) GetCamundaAsyncBefore() *bool {
	return &receiveTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (receiveTask ReceiveTask) GetCamundaAsyncAfter() *bool {
	return &receiveTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (receiveTask ReceiveTask) GetCamundaJobPriority() *int {
	return &receiveTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (receiveTask ReceiveTask) GetDocumentation() *attributes.Documentation {
	return &receiveTask.Documentation[0]
}

// GetExtensionElements ...
func (receiveTask ReceiveTask) GetExtensionElements() *camunda.ExtensionElements {
	return &receiveTask.ExtensionElements[0]
}

// GetIncoming ...
func (receiveTask ReceiveTask) GetIncoming(num int) *marker.Incoming {
	return &receiveTask.Incoming[num]
}

// GetOutgoing ...
func (receiveTask ReceiveTask) GetOutgoing(num int) *marker.Outgoing {
	return &receiveTask.Outgoing[num]
}
