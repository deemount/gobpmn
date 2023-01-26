package subprocesses

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/marker"
)

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

/*** Attributes ***/

// SetDocumentation ...
func (transaction *Transaction) SetDocumentation() {
	transaction.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (transaction *Transaction) SetExtensionElements() {
	transaction.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*** Marker ***/

// SetIncoming ...
func (transaction *Transaction) SetIncoming(num int) {
	transaction.Incoming = make([]marker.Incoming, num)
}

// SetOutgoing ...
func (transaction *Transaction) SetOutgoing(num int) {
	transaction.Outgoing = make([]marker.Outgoing, num)
}

/*** Flow ***/

// SetSequenceFlow ...
func (transaction *Transaction) SetSequenceFlow(num int) {
	transaction.SequenceFlow = make([]flow.SequenceFlow, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (transaction Transaction) GetID() impl.STR_PTR {
	return &transaction.ID
}

// GetName ...
func (transaction Transaction) GetName() impl.STR_PTR {
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

/*** Attributes ***/

// GetDocumentation ...
func (transaction Transaction) GetDocumentation() *attributes.Documentation {
	return &transaction.Documentation[0]
}

// GetExtensionElements ...
func (transaction Transaction) GetExtensionElements() *attributes.ExtensionElements {
	return &transaction.ExtensionElements[0]
}

/*** Marker ***/

// GetIncoming ...
func (transaction Transaction) GetIncoming(num int) *marker.Incoming {
	return &transaction.Incoming[num]
}

// GetOutgoing ...
func (transaction Transaction) GetOutgoing(num int) *marker.Outgoing {
	return &transaction.Outgoing[num]
}

// GetSequenceFlow ...
func (transaction Transaction) GetSequenceFlow(num int) *flow.SequenceFlow {
	return &transaction.SequenceFlow[num]
}
