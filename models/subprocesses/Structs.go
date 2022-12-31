package subprocesses

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/camunda"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/events"
	"github.com/deemount/gobpmn/models/events/elements"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/loop"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/tasks"
)

// Subprocesses ...
type Subprocesses struct {
	CallActivity    CALL_ACTIVITY_SLC    `xml:"bpmn:callActivity,omitempty" json:"callActivity,omitempty" csv:"-"`
	SubProcess      SUBPROCESS_SLC       `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty" csv:"-"`
	Transaction     TRANSACTION_SLC      `xml:"bpmn:transaction,omitempty" json:"transaction,omitempty" csv:"-"`
	AdHocSubProcess ADHOC_SUBPROCESS_SLC `xml:"bpmn:adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty" csv:"-"`
}

// TSubprocesses ...
type TSubprocesses struct {
	CallActivity    TCALL_ACTIVITY_SLC    `xml:"callActivity,omitempty" json:"callActivity,omitempty"`
	SubProcess      TSUBPROCESS_SLC       `xml:"subProcess,omitempty" json:"subProcess,omitempty"`
	Transaction     TTRANSACTION_SLC      `xml:"transaction,omitempty" json:"transaction,omitempty"`
	AdHocSubProcess TADHOC_SUBPROCESS_SLC `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty" csv:"-"`
}

// AdHocSubProcess ...
type AdHocSubProcess struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	gateways.Gateways
	tasks.Tasks
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	StartEvent                       []elements.StartEvent                   `xml:"bpmn:startEvent,omitemnpty" json:"startEvent,omitempty"`
	EndEvent                         []elements.EndEvent                     `xml:"bpmn:endEvent,omitempty" json:"endEvent,omitempty"`
	SubProcess                       SUBPROCESS_SLC                          `xml:"bpmn:subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  ADHOC_SUBPROCESS_SLC                    `xml:"bpmn:adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	SequenceFlow                     []marker.SequenceFlow                   `xml:"bpmn:sequenceFlow,omitempty" json:"equenceFlow,omitempty"`
}

// TSubProcess ...
type TAdHocSubProcess struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	gateways.TGateways
	tasks.TTasks
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	StartEvent                       []elements.TStartEvent                  `xml:"startEvent,omitemnpty" json:"startEvent,omitempty"`
	EndEvent                         []elements.TEndEvent                    `xml:"endEvent,omitempty" json:"endEvent,omitempty"`
	SubProcess                       TSUBPROCESS_SLC                         `xml:"subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  TADHOC_SUBPROCESS_SLC                   `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	SequenceFlow                     []marker.SequenceFlow                   `xml:"sequenceFlow,omitempty" json:"equenceFlow,omitempty"`
}

// CallActivity ...
type CallActivity struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	CalledElement                    string                                  `xml:"calledElement,attr,omitempty" json:"calledElement,omitempty"`
	CamundaCalledElementTenantID     string                                  `xml:"camunda:calledElementTenantId,attr,omitempty" json:"calledElementTenantId,omitempty"`
	CamundaVariableMappingClass      string                                  `xml:"camunda:variableMappingClass,attr,omitempty" json:"variableMappingClass,omitempty"`
	StandardLoopCharacteristics      []loop.StandardLoopCharacteristics      `xml:"bpmn:standardLoopCharacteristics,omitempty" json:"standardLoopCharacteristics,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics,omitempty"`
}

// TCallActivity ...
type TCallActivity struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	CalledElement                    string                                  `xml:"calledElement,attr,omitempty" json:"calledElement,omitempty"`
	CamundaCalledElementTenantID     string                                  `xml:"calledElementTenantId,attr,omitempty" json:"calledElementTenantId,omitempty"`
	CamundaVariableMappingClass      string                                  `xml:"variableMappingClass,attr,omitempty" json:"variableMappingClass,omitempty"`
	StandardLoopCharacteristics      []loop.StandardLoopCharacteristics      `xml:"standardLoopCharacteristics,omitempty" json:"standardLoopCharacteristics,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics,omitempty"`
}

// SubProcess ...
type SubProcess struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	events.Events
	tasks.Tasks
	gateways.Gateways
	Subprocesses
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"bpmn:multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	SequenceFlow                     []marker.SequenceFlow                   `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// TSubProcess ...
type TSubProcess struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.TCoreAttributes
	marker.IncomingOutgoing
	events.TEvents
	tasks.TTasks
	gateways.TGateways
	TriggeredByEvent                 bool                                    `xml:"triggeredByEvent,attr,omitempty" json:"triggeredByEvent,omitempty"`
	MultiInstanceLoopCharacteristics []loop.MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics,omitempty" json:"multiInstanceLoopCharacteristics"`
	SubProcess                       TSUBPROCESS_SLC                         `xml:"subProcess,omitempty" json:"subProcess,omitempty"`           // is that possible ?
	AdHocSubProcess                  TADHOC_SUBPROCESS_SLC                   `xml:"adhocSubprocess,omitempty" json:"adhocSubprocess,omitempty"` // is that possible ?
	SequenceFlow                     []marker.SequenceFlow                   `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// Transaction ...
type Transaction struct {
	compulsion.BaseAttributes
	attributes.Attributes
	camunda.CoreAttributes
	marker.IncomingOutgoing
	SequenceFlow []marker.SequenceFlow `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// TTransaction ...
type TTransaction struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	camunda.TCoreAttributes
	marker.TIncomingOutgoing
	SequenceFlow []marker.SequenceFlow `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}
