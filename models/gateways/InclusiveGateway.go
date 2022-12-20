package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// InclusiveGatewayRepository ...
type InclusiveGatewayRepository interface {
	SetID(typ string, suffix interface{})
	SetName(name string)

	SetCamundaAsyncBefore(asyncBefore bool)
	SetCamundaAsyncAfter(asyncAfter bool)
	SetCamundaJobPriority(priority int)

	SetExtensionElements()
	SetIncoming(num int)
	SetOutgoing(num int)

	GetID() *string
	GetName() *string

	GetCamundaAsyncBefore() *bool
	GetCamundaAsyncAfter() *bool
	GetCamundaJobPriority() *int

	GetExtensionElements() *camunda.ExtensionElements
	GetIncoming(num int) *marker.Incoming
	GetOutgoing(num int) *marker.Outgoing
}

// InclusiveGateway ...
type InclusiveGateway struct {
	ID                 string                      `xml:"id,attr" json:"id"`
	Name               string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore"`
	CamundaAsyncAfter  bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter"`
	CamundaJobPriority int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing           `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TInclusiveGateway ...
type TInclusiveGateway struct {
	ID                string                       `xml:"id,attr" json:"id"`
	Name              string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore       bool                         `xml:"asyncBefore,attr,omitempty" json:"asyncBefore"`
	AsyncAfter        bool                         `xml:"asyncAfter,attr,omitempty" json:"asyncAfter"`
	JobPriority       int                          `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation     []attributes.Documentation   `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming            `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing            `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewInclusiveGateway() InclusiveGatewayRepository {
	return &InclusiveGateway{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusiveGateway *InclusiveGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		inclusiveGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		inclusiveGateway.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (inclusiveGateway *InclusiveGateway) SetName(name string) {
	inclusiveGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (inclusiveGateway *InclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	inclusiveGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (inclusiveGateway *InclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	inclusiveGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (inclusiveGateway *InclusiveGateway) SetCamundaJobPriority(priority int) {
	inclusiveGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetExtensionElements ...
func (inclusiveGateway *InclusiveGateway) SetExtensionElements() {
	inclusiveGateway.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (inclusiveGateway *InclusiveGateway) SetIncoming(num int) {
	inclusiveGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (inclusiveGateway *InclusiveGateway) SetOutgoing(num int) {
	inclusiveGateway.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (inclusiveGateway InclusiveGateway) GetID() *string {
	return &inclusiveGateway.ID

}

// SetName ...
func (inclusiveGateway InclusiveGateway) GetName() *string {
	return &inclusiveGateway.Name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (inclusiveGateway InclusiveGateway) GetCamundaAsyncBefore() *bool {
	return &inclusiveGateway.CamundaAsyncBefore
}

// SetCamundaAsyncAfter ...
func (inclusiveGateway InclusiveGateway) GetCamundaAsyncAfter() *bool {
	return &inclusiveGateway.CamundaAsyncAfter
}

// SetCamundaJobPriority ...
func (inclusiveGateway InclusiveGateway) GetCamundaJobPriority() *int {
	return &inclusiveGateway.CamundaJobPriority
}

/* Elements */

// GetExtensionElements ...
func (inclusiveGateway InclusiveGateway) GetExtensionElements() *camunda.ExtensionElements {
	return &inclusiveGateway.ExtensionElements[0]
}

// GetIncoming ...
func (inclusiveGateway InclusiveGateway) GetIncoming(num int) *marker.Incoming {
	return &inclusiveGateway.Incoming[num]
}

// GetOutgoing ...
func (inclusiveGateway InclusiveGateway) GetOutgoing(num int) *marker.Outgoing {
	return &inclusiveGateway.Outgoing[num]
}
