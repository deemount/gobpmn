package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// ParallelGatewayRepository ...
type ParallelGatewayRepository interface {
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

// ParallelGateway ...
type ParallelGateway struct {
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

// TParallelGateway ...
type TParallelGateway struct {
	ID                string                      `xml:"id,attr" json:"id"`
	Name              string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore       bool                        `xml:"asyncBefore,attr,omitempty" json:"asyncBefore"`
	AsyncAfter        bool                        `xml:"asyncAfter,attr,omitempty" json:"asyncAfter"`
	JobPriority       int                         `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation     []attributes.Documentation  `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.ExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming           `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing           `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewParallelGateway() ParallelGatewayRepository {
	return &ParallelGateway{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (parallelGateway *ParallelGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		parallelGateway.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		parallelGateway.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (parallelGateway *ParallelGateway) SetName(name string) {
	parallelGateway.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (parallelGateway *ParallelGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	parallelGateway.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (parallelGateway *ParallelGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	parallelGateway.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (parallelGateway *ParallelGateway) SetCamundaJobPriority(priority int) {
	parallelGateway.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetExtensionElements ...
func (parallelGateway *ParallelGateway) SetExtensionElements() {
	parallelGateway.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (parallelGateway *ParallelGateway) SetIncoming(num int) {
	parallelGateway.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (parallelGateway *ParallelGateway) SetOutgoing(num int) {
	parallelGateway.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (parallelGateway ParallelGateway) GetID() *string {
	return &parallelGateway.ID
}

// GetName ...
func (parallelGateway ParallelGateway) GetName() *string {
	return &parallelGateway.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (parallelGateway ParallelGateway) GetCamundaAsyncBefore() *bool {
	return &parallelGateway.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (parallelGateway ParallelGateway) GetCamundaAsyncAfter() *bool {
	return &parallelGateway.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (parallelGateway ParallelGateway) GetCamundaJobPriority() *int {
	return &parallelGateway.CamundaJobPriority
}

/* Elements */

// GetExtensionElements ...
func (parallelGateway ParallelGateway) GetExtensionElements() *camunda.ExtensionElements {
	return &parallelGateway.ExtensionElements[0]
}

// GetIncoming ...
func (parallelGateway ParallelGateway) GetIncoming(num int) *marker.Incoming {
	return &parallelGateway.Incoming[num]
}

// GetOutgoing ...
func (parallelGateway ParallelGateway) GetOutgoing(num int) *marker.Outgoing {
	return &parallelGateway.Outgoing[num]
}
