package gateways

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/marker"
)

// ComplexGatewayRepository ...
type ComplexGatewayRepository interface {
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

// ComplexGateway ...
type ComplexGateway struct {
	ID                 string                      `xml:"id,attr" json:"id"`
	Name               string                      `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                        `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                        `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                         `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation  `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []camunda.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming           `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing           `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

// TComplexGateway ...
type TComplexGateway struct {
	ID                string                       `xml:"id,attr" json:"id"`
	Name              string                       `xml:"name,attr,omitempty" json:"name,omitempty"`
	AsyncBefore       bool                         `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter        bool                         `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority       int                          `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation     []attributes.Documentation   `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []camunda.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming          []marker.Incoming            `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing          []marker.Outgoing            `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}

func NewComplexGateway() ComplexGatewayRepository {
	return &ComplexGateway{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (complexgate *ComplexGateway) SetID(typ string, suffix interface{}) {
	switch typ {
	case "gateway":
		complexgate.ID = fmt.Sprintf("Gateway_%v", suffix)
		break
	case "id":
		complexgate.ID = fmt.Sprintf("%s", suffix)
	}
}

// SetName ...
func (complexgate *ComplexGateway) SetName(name string) {
	complexgate.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexgate *ComplexGateway) SetCamundaAsyncBefore(asyncBefore bool) {
	complexgate.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (complexgate *ComplexGateway) SetCamundaAsyncAfter(asyncAfter bool) {
	complexgate.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (complexgate *ComplexGateway) SetCamundaJobPriority(priority int) {
	complexgate.CamundaJobPriority = priority
}

/*** Make Elements ***/

// SetExtensionElements ...
func (complexgate *ComplexGateway) SetExtensionElements() {
	complexgate.ExtensionElements = make([]camunda.ExtensionElements, 1)
}

// SetIncoming ...
func (complexgate *ComplexGateway) SetIncoming(num int) {
	complexgate.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (complexgate *ComplexGateway) SetOutgoing(num int) {
	complexgate.Outgoing = make([]marker.Outgoing, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (complexgate ComplexGateway) GetID() *string {
	return &complexgate.ID
}

// SetName ...
func (complexgate ComplexGateway) GetName() *string {
	return &complexgate.Name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (complexgate ComplexGateway) GetCamundaAsyncBefore() *bool {
	return &complexgate.CamundaAsyncBefore
}

// SetCamundaAsyncAfter ...
func (complexgate ComplexGateway) GetCamundaAsyncAfter() *bool {
	return &complexgate.CamundaAsyncAfter
}

// SetCamundaJobPriority ...
func (complexgate ComplexGateway) GetCamundaJobPriority() *int {
	return &complexgate.CamundaJobPriority
}

/*** Make Elements ***/

// SetExtensionElements ...
func (complexgate ComplexGateway) GetExtensionElements() *camunda.ExtensionElements {
	return &complexgate.ExtensionElements[0]
}

// SetIncoming ...
func (complexgate ComplexGateway) GetIncoming(num int) *marker.Incoming {
	return &complexgate.Incoming[num]
}

// SetOutgoing ...
func (complexgate ComplexGateway) GetOutgoing(num int) *marker.Outgoing {
	return &complexgate.Outgoing[num]
}
