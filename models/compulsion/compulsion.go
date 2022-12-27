package compulsion

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

type CompulsionCoreAttributes struct {
	ID   string `xml:"id,attr,omitempty" json:"id" csv:"ID"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" csv:"NAME"`
}

type CompulsionCoreElements struct {
	Documentation     []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

type TCompulsionCoreElements struct {
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

type CompulsionCamundaCoreAttributes struct {
	CamundaAsyncBefore bool `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int  `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
}

type TCompulsionCamundaCoreAttributes struct {
	AsyncBefore bool `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter  bool `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	JobPriority int  `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
}

type CompulsionCoreIncomingOutgoing struct {
	Incoming []marker.Incoming `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing []marker.Outgoing `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
}

type TCompulsionCoreIncomingOutgoing struct {
	Incoming []marker.Incoming `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing []marker.Outgoing `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
}
