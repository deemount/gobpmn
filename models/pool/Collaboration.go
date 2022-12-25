package pool

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/marker"
)

// CollaborationRepository ...
type CollaborationRepository interface {
	PoolID

	SetDocumentation()
	SetExtensionElements()
	SetParticipant(num int)
	SetMessageFlow(num int)

	GetDocumentation() *attributes.Documentation
	GetExtensionElements() *attributes.ExtensionElements
	GetParticipant(num int) *Participant
	GetMessageFlow(num int) *marker.MessageFlow
}

// Collaboration ...
type Collaboration struct {
	ID                string                         `xml:"id,attr" json:"id"`
	Documentation     []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
	Participant       []Participant                  `xml:"bpmn:participant" json:"participant,omitempty"`
	MessageFlow       []marker.MessageFlow           `xml:"bpmn:messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// TCollaboration ...
type TCollaboration struct {
	ID                string                          `xml:"id,attr" json:"id"`
	Documentation     []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
	Participant       []TParticipant                  `xml:"participant" json:"participant,omitempty"`
	MessageFlow       []marker.TMessageFlow           `xml:"messageFlow,omitempty" json:"messageFlow,omitempty"`
}

func NewCollaboration() CollaborationRepository {
	return &Collaboration{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (collaboration *Collaboration) SetID(typ string, suffix interface{}) {
	switch typ {
	case "collaboration":
		collaboration.ID = fmt.Sprintf("Collaboration_%v", suffix)
		break
	case "id":
		collaboration.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/* Make Elements */

/** BPMN **/

// SetDocumentation ...
func (collaboration *Collaboration) SetDocumentation() {
	collaboration.Documentation = make([]attributes.Documentation, 1)
}

// SetExtensionElements ...
func (collaboration *Collaboration) SetExtensionElements() {
	collaboration.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

// SetParticipant ...
func (collaboration *Collaboration) SetParticipant(num int) {
	collaboration.Participant = make([]Participant, num)
}

// SetMessageFlow ...
func (collaboration *Collaboration) SetMessageFlow(num int) {
	collaboration.MessageFlow = make([]marker.MessageFlow, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (collaboration Collaboration) GetID() STR_PTR {
	return &collaboration.ID
}

/* Elements */

/** BPMN **/

// GetDocumentation ...
func (collaboration Collaboration) GetDocumentation() *attributes.Documentation {
	return &collaboration.Documentation[0]
}

// GetExtensionElements ...
func (collaboration Collaboration) GetExtensionElements() *attributes.ExtensionElements {
	return &collaboration.ExtensionElements[0]
}

// GetParticipant ...
func (collaboration Collaboration) GetParticipant(num int) *Participant {
	return &collaboration.Participant[num]
}

// GetMessageFlow ...
func (collaboration Collaboration) GetMessageFlow(num int) *marker.MessageFlow {
	return &collaboration.MessageFlow[num]
}
