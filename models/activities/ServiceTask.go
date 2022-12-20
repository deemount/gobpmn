package activities

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// ServiceTaskRepository ...
type ServiceTaskRepository interface {
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

// ServiceTask ...
type ServiceTask struct {
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

// TServiceTask ...
type TServiceTask struct {
	ID                 string                      `xml:"id,attr,omitempty" json:"id"`
	Name               string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                        `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                        `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                         `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation  `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []camunda.ExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming           `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing           `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewServiceTask() ServiceTaskRepository {
	return &ServiceTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (serviceTask *ServiceTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		serviceTask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		serviceTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (serviceTask *ServiceTask) SetName(name string) {
	serviceTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (serviceTask *ServiceTask) SetCamundaAsyncBefore(asyncBefore bool) {
	serviceTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (serviceTask *ServiceTask) SetCamundaAsyncAfter(asyncAfter bool) {
	serviceTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (serviceTask *ServiceTask) SetCamundaJobPriority(priority int) {
	serviceTask.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetDocumentation ...
func (serviceTask *ServiceTask) SetDocumentation() {
	serviceTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (serviceTask *ServiceTask) SetExtensionElements() {
	serviceTask.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (serviceTask *ServiceTask) SetIncoming(num int) {
	serviceTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (serviceTask *ServiceTask) SetOutgoing(num int) {
	serviceTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (serviceTask ServiceTask) GetID() *string {
	return &serviceTask.ID
}

// GetName ...
func (serviceTask ServiceTask) GetName() *string {
	return &serviceTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (serviceTask ServiceTask) GetCamundaAsyncBefore() *bool {
	return &serviceTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (serviceTask ServiceTask) GetCamundaAsyncAfter() *bool {
	return &serviceTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (serviceTask ServiceTask) GetCamundaJobPriority() *int {
	return &serviceTask.CamundaJobPriority
}

/* Elements */

// GetDocumentation ...
func (serviceTask ServiceTask) GetDocumentation() *attributes.Documentation {
	return &serviceTask.Documentation[0]
}

// GetExtensionElements ...
func (serviceTask ServiceTask) GetExtensionElements() *camunda.ExtensionElements {
	return &serviceTask.ExtensionElements[0]
}

// GetIncoming ...
func (serviceTask ServiceTask) GetIncoming(num int) *marker.Incoming {
	return &serviceTask.Incoming[num]
}

// GetOutgoing ...
func (serviceTask ServiceTask) GetOutgoing(num int) *marker.Outgoing {
	return &serviceTask.Outgoing[num]
}
