package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

// EventBasedGatewayRepository ...
type EventBasedGatewayRepository interface {
	GatewayBase

	SetExtensionElements()
	GetExtensionElements() *attributes.ExtensionElements
}

// EventBasedGateway ...
type EventBasedGateway struct {
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

// TEventBasedGateway ...
type TEventBasedGateway struct {
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
func (eventBasedGateway *EventBasedGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		eventBasedGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		eventBasedGateway.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (eventBasedGateway *EventBasedGateway) SetName(name string) {
	eventBasedGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (eventBasedGateway *EventBasedGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	eventBasedGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (eventBasedGateway *EventBasedGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	eventBasedGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (eventBasedGateway *EventBasedGateway) SetCamundaJobPriority(priority int) {
	eventBasedGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetExtensionElements ...
func (eventBasedGateway *EventBasedGateway) SetExtensionElements() {
	eventBasedGateway.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (eventBasedGateway *EventBasedGateway) SetIncoming(num int) {
	eventBasedGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (eventBasedGateway *EventBasedGateway) SetOutgoing(num int) {
	eventBasedGateway.Outgoing = make([]marker.Outgoing, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (eventBasedGateway EventBasedGateway) GetID() STR_PTR {
	return &eventBasedGateway.ID
}

// GetName ...
func (eventBasedGateway EventBasedGateway) GetName() STR_PTR {
	return &eventBasedGateway.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (eventBasedGateway EventBasedGateway) GetCamundaAsyncBefore() *bool {
	return &eventBasedGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (eventBasedGateway EventBasedGateway) GetCamundaAsyncAfter() *bool {
	return &eventBasedGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (eventBasedGateway EventBasedGateway) GetCamundaJobPriority() *int {
	return &eventBasedGateway.CamundaJobPriority
}

/* Elements */

// GetExtensionElements ...
func (eventBasedGateway EventBasedGateway) GetExtensionElements() *attributes.ExtensionElements {
	return &eventBasedGateway.ExtensionElements[0]
}

// GetIncoming ...
func (eventBasedGateway EventBasedGateway) GetIncoming(num int) *marker.Incoming {
	return &eventBasedGateway.Incoming[num]
}

// GetOutgoing ...
func (eventBasedGateway EventBasedGateway) GetOutgoing(num int) *marker.Outgoing {
	return &eventBasedGateway.Outgoing[num]
}
