package marker

import (
	"fmt"
	"log"

	"github.com/deemount/gobpmn/models/attributes"
)

// MessageFlowRepository ...
type MessageFlowRepository interface {
	MarkerBase

	SetSourceRef(typ string, sourceRef interface{})
	SetTargetRef(typ string, targetRef interface{})
	GetSourceRef() *string
	GetTargetRef() *string

	SetDocumentation()
	SetExtensionElements()
	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *attributes.ExtensionElements
}

// MessageFlow ...
type MessageFlow struct {
	ID                string                         `xml:"id,attr" json:"id"`
	Name              string                         `xml:"name,attr,omitempty" json:"name,omitempty"`
	SourceRef         string                         `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string                         `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TMessageFlow ...
type TMessageFlow struct {
	ID                string                          `xml:"id,attr" json:"id"`
	Name              string                          `xml:"name,attr,omitempty" json:"name,omitempty"`
	SourceRef         string                          `xml:"sourceRef,attr" json:"sourceRef,omitempty"`
	TargetRef         string                          `xml:"targetRef,attr" json:"targetRef,omitempty"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

func NewMessageFlow() MessageFlowRepository {
	return &MessageFlow{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (messageFlow *MessageFlow) SetID(typ string, suffix interface{}) {
	switch typ {
	case "flow":
		messageFlow.ID = fmt.Sprintf("Flow_%v", suffix)
		break
	case "id":
		messageFlow.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (messageFlow *MessageFlow) SetName(name string) {
	messageFlow.Name = name
}

// SetSourceRef ...
func (messageFlow *MessageFlow) SetSourceRef(typ string, sourceRef interface{}) {
	switch typ {
	case "activity":
		messageFlow.SourceRef = fmt.Sprintf("Activity_%s", sourceRef)
		break
	case "event":
		messageFlow.SourceRef = fmt.Sprintf("Event_%s", sourceRef)
		break
	case "id":
		messageFlow.SourceRef = fmt.Sprintf("%s", sourceRef)
		break
	case "participant":
		messageFlow.SourceRef = fmt.Sprintf("Participant_%s", sourceRef)
		break
	}
}

// SetTargetRef ...
func (messageFlow *MessageFlow) SetTargetRef(typ string, targetRef interface{}) {
	switch typ {
	case "activity":
		messageFlow.TargetRef = fmt.Sprintf("Activity_%s", targetRef)
		break
	case "event":
		messageFlow.TargetRef = fmt.Sprintf("Event_%s", targetRef)
		break
	case "id":
		messageFlow.TargetRef = fmt.Sprintf("%s", targetRef)
		break
	case "participant":
		messageFlow.TargetRef = fmt.Sprintf("Participant_%s", targetRef)
		break
	default:
		log.Panic("no typ set in target ref for message flow")
	}
}

/*** Make Elements ***/

/** BPMN **/

// SetDocumentation ...
func (messageFlow *MessageFlow) SetDocumentation() {
	messageFlow.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (messageFlow *MessageFlow) SetExtensionElements() {
	messageFlow.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (messageFlow MessageFlow) GetID() STR_PTR {
	return &messageFlow.ID
}

// GetName ...
func (messageFlow MessageFlow) GetName() STR_PTR {
	return &messageFlow.Name
}

// GetSourceRef ...
func (messageFlow MessageFlow) GetSourceRef() *string {
	return &messageFlow.SourceRef
}

// GetTargetRef ...
func (messageFlow MessageFlow) GetTargetRef() *string {
	return &messageFlow.TargetRef
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (messageFlow MessageFlow) GetDocumentation() *attributes.Documentation {
	return &messageFlow.Documentation[0]
}

// GetExtensionElements ...
func (messageFlow MessageFlow) GetExtensionElements() *attributes.ExtensionElements {
	return &messageFlow.ExtensionElements[0]
}
