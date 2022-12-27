package subprocesses

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

type STR_PTR *string
type DOCUMENTATION_PTR *attributes.Documentation
type EXTENSION_ELEMENTS_PTR *attributes.ExtensionElements

// subprocesses
const SUB_PROCESS string = "subProcess"
const SUB_PROCESS_AD_HOC string = "adHocSubProcess"
const CALL_ACTIVITY string = "callActivity"
const TRANSACTION string = "transaction"

type SubprocessesID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type SubprocessesName interface {
	SetName(name string)
	GetName() STR_PTR
}

type SubprocessesMarker interface {
	SetIncoming(num int)
	GetIncoming(num int) *marker.Incoming
	SetOutgoing(num int)
	GetOutgoing(num int) *marker.Outgoing
}

type SubprocessesFlow interface {
	SetSequenceFlow(num int)
	GetSequenceFlow(num int) *marker.SequenceFlow
}

type SubprocessesCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() *bool
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() *bool
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() *int
}

type SubprocessesBaseCoreElements interface {
	SetDocumentation()
	GetDocumentation() DOCUMENTATION_PTR
	SetExtensionElements()
	GetExtensionElements() EXTENSION_ELEMENTS_PTR
}

type SubprocessesBase interface {
	SubprocessesID
	SubprocessesName
	SubprocessesBaseCoreElements
	SubprocessesMarker
	SubprocessesCamundaBase
}
