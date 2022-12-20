package activities

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// BusinessRuleTaskRepository ...
type BusinessRuleTaskRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)

	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaJobPriority(priority int)
	SetCamundaClass(class string)

	SetDocumentation()
	SetExtensionElements()
	SetIncoming(num int)
	SetOutgoing(num int)

	GetID() *string
	GetName() *string

	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaJobPriority() *int
	GetCamundaClass() *string

	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *camunda.ExtensionElements
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing
}

// BusinessRuleTask ...
type BusinessRuleTask struct {
	ID                 string                      `xml:"id,attr" json:"id"`
	Name               string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	CamundaClass       string                      `xml:"camunda:class,attr,omitempty" json:"class,omitempty"`
	Documentation      []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing           `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TBusinessRuleTask ...
type TBusinessRuleTask struct {
	ID                string                      `xml:"id,attr" json:"id"`
	Name              string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore       bool                        `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter        bool                        `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority       int                         `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Class             string                      `xml:"class,attr,omitempty" json:"class,omitempty"`
	Documentation     []attributes.Documentation  `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.ExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming           `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing           `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewBusinessRuleTask() BusinessRuleTaskRepository {
	return &BusinessRuleTask{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (businessRuleTask *BusinessRuleTask) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		businessRuleTask.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		businessRuleTask.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (businessRuleTask *BusinessRuleTask) SetName(name string) {
	businessRuleTask.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (businessRuleTask *BusinessRuleTask) SetCamundaAsyncBefore(asyncBefore bool) {
	businessRuleTask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (businessRuleTask *BusinessRuleTask) SetCamundaAsyncAfter(asyncAfter bool) {
	businessRuleTask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (businessRuleTask *BusinessRuleTask) SetCamundaJobPriority(priority int) {
	businessRuleTask.CamundaJobPriority = priority
}

// SetCamundaClass ...
func (businessRuleTask *BusinessRuleTask) SetCamundaClass(class string) {
	businessRuleTask.CamundaClass = class
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (businessRuleTask *BusinessRuleTask) SetDocumentation() {
	businessRuleTask.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (businessRuleTask *BusinessRuleTask) SetExtensionElements() {
	businessRuleTask.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (businessRuleTask *BusinessRuleTask) SetIncoming(num int) {
	businessRuleTask.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (businessRuleTask *BusinessRuleTask) SetOutgoing(num int) {
	businessRuleTask.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (businessRuleTask BusinessRuleTask) GetID() *string {
	return &businessRuleTask.ID
}

// GetName ...
func (businessRuleTask BusinessRuleTask) GetName() *string {
	return &businessRuleTask.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (businessRuleTask BusinessRuleTask) GetCamundaAsyncBefore() *bool {
	return &businessRuleTask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (businessRuleTask BusinessRuleTask) GetCamundaAsyncAfter() *bool {
	return &businessRuleTask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (businessRuleTask BusinessRuleTask) GetCamundaJobPriority() *int {
	return &businessRuleTask.CamundaJobPriority
}

// GetCamundaClass ...
func (businessRuleTask BusinessRuleTask) GetCamundaClass() *string {
	return &businessRuleTask.CamundaClass
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (businessRuleTask BusinessRuleTask) GetDocumentation() *attributes.Documentation {
	return &businessRuleTask.Documentation[0]
}

// SetExtensionElements ...
func (businessRuleTask BusinessRuleTask) GetExtensionElements() *camunda.ExtensionElements {
	return &businessRuleTask.ExtensionElements[0]
}

// SetIncoming ...
func (businessRuleTask BusinessRuleTask) GetIncoming(num int) *marker.Incoming {
	return &businessRuleTask.Incoming[num]
}

// SetOutgoing ...
func (businessRuleTask BusinessRuleTask) GetOutgoing(num int) *marker.Outgoing {
	return &businessRuleTask.Outgoing[num]
}
