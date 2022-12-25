package subprocesses

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

// TransactionRepository ...
type TransactionRepository interface {
	SubprocessesBase
	SubprocessesFlow
}

// Transaction ...
type Transaction struct {
	ID                 string                         `xml:"id,attr" json:"id"`
	Name               string                         `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                           `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                           `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                            `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming              `xml:"bpmn:incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing              `xml:"bpmn:outgoing,omitempty" json:"outgoing,omitempty"`
	SequenceFlow       []marker.SequenceFlow          `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

// TTransaction ...
type TTransaction struct {
	ID                 string                          `xml:"id,attr" json:"id"`
	Name               string                          `xml:"name,attr,omitempty" json:"name,omitempty"`
	CamundaAsyncBefore bool                            `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter  bool                            `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaJobPriority int                             `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	Documentation      []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements  []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Incoming           []marker.Incoming               `xml:"incoming,omitempty" json:"incoming,omitempty"`
	Outgoing           []marker.Outgoing               `xml:"outgoing,omitempty" json:"outgoing,omitempty"`
	SequenceFlow       []marker.SequenceFlow           `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
}

func NewTransaction() TransactionRepository {
	return &Transaction{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (transaction *Transaction) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		transaction.ID = fmt.Sprintf("Activity_%v", suffix)
		break
	case "id":
		transaction.ID = fmt.Sprintf("%v", suffix)
		break
	}
}

// SetName ...
func (transaction *Transaction) SetName(name string) {
	transaction.Name = name
}

/** Camunda **/

// SetCamundaAsyncBefore ...
func (transaction *Transaction) SetCamundaAsyncBefore(asyncBefore bool) {
	transaction.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (transaction *Transaction) SetCamundaAsyncAfter(asyncAfter bool) {
	transaction.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (transaction *Transaction) SetCamundaJobPriority(priority int) {
	transaction.CamundaJobPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (transaction *Transaction) SetDocumentation() {
	transaction.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (transaction *Transaction) SetExtensionElements() {
	transaction.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetIncoming ...
func (transaction *Transaction) SetIncoming(num int) {
	transaction.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (transaction *Transaction) SetOutgoing(num int) {
	transaction.Outgoing = make([]marker.Outgoing, num)
}

/*** Marker ***/

// SetSequenceFlow ...
func (transaction *Transaction) SetSequenceFlow(num int) {
	transaction.SequenceFlow = make([]marker.SequenceFlow, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (transaction Transaction) GetID() STR_PTR {
	return &transaction.ID
}

// GetName ...
func (transaction Transaction) GetName() STR_PTR {
	return &transaction.Name
}

/** Camunda **/

// GetCamundaAsyncBefore ...
func (transaction Transaction) GetCamundaAsyncBefore() *bool {
	return &transaction.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (transaction Transaction) GetCamundaAsyncAfter() *bool {
	return &transaction.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (transaction Transaction) GetCamundaJobPriority() *int {
	return &transaction.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (transaction Transaction) GetDocumentation() *attributes.Documentation {
	return &transaction.Documentation[0]
}

// GetExtensionElements ...
func (transaction Transaction) GetExtensionElements() *attributes.ExtensionElements {
	return &transaction.ExtensionElements[0]
}

// GetIncoming ...
func (transaction Transaction) GetIncoming(num int) *marker.Incoming {
	return &transaction.Incoming[num]
}

// GetOutgoing ...
func (transaction Transaction) GetOutgoing(num int) *marker.Outgoing {
	return &transaction.Outgoing[num]
}

/*** Marker ***/

// GetSequenceFlow ...
func (transaction Transaction) GetSequenceFlow(num int) *marker.SequenceFlow {
	return &transaction.SequenceFlow[num]
}
