package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

// ExclusiveGatewayRepository ...
type ExclusiveGatewayRepository interface {
	GatewayBase

	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

// ExclusiveGateway ...
type ExclusiveGateway struct {
	ID                 string                         `xml:"id,attr" json:"id"`
	Name               string                         `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                           `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore"`
	CamundaAsyncAfter  bool                           `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter"`
	CamundaJobPriority int                            `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming              `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing              `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TExclusiveGateway ...
type TExclusiveGateway struct {
	ID                string                          `xml:"id,attr" json:"id"`
	Name              string                          `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore       bool                            `xml:"asyncBefore,attr,omitempty" json:"asyncBefore"`
	AsyncAfter        bool                            `xml:"asyncAfter,attr,omitempty" json:"asyncAfter"`
	JobPriority       int                             `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming               `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing               `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (exclusiveGateway *ExclusiveGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		exclusiveGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		exclusiveGateway.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (exclusiveGateway *ExclusiveGateway) SetName(name string) {
	exclusiveGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	exclusiveGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	exclusiveGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (exclusiveGateway *ExclusiveGateway) SetCamundaJobPriority(priority int) {
	exclusiveGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetExtensionElements ...
func (exclusiveGateway *ExclusiveGateway) SetExtensionElements() {
	exclusiveGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (exclusiveGateway *ExclusiveGateway) SetIncoming(num int) {
	exclusiveGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (exclusiveGateway *ExclusiveGateway) SetOutgoing(num int) {
	exclusiveGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (exclusiveGateway ExclusiveGateway) GetID() *string {
	return &exclusiveGateway.ID
}

// GetName ...
func (exclusiveGateway ExclusiveGateway) GetName() *string {
	return &exclusiveGateway.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (exclusiveGateway ExclusiveGateway) GetCamundaAsyncBefore() *bool {
	return &exclusiveGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (exclusiveGateway ExclusiveGateway) GetCamundaAsyncAfter() *bool {
	return &exclusiveGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (exclusiveGateway ExclusiveGateway) GetCamundaJobPriority() *int {
	return &exclusiveGateway.CamundaJobPriority
}

/*** Make Elements ***/

// GetExtensionElements ...
func (exclusiveGateway ExclusiveGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &exclusiveGateway.ExtensionElements[0]
}

// GetIncoming ...
func (exclusiveGateway ExclusiveGateway) GetIncoming(num int) *marker.Incoming {
	return &exclusiveGateway.Incoming[num]
}

// GetOutgoing ...
func (exclusiveGateway ExclusiveGateway) GetOutgoing(num int) *marker.Outgoing {
	return &exclusiveGateway.Outgoing[num]
}
