package collaboration

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewCollaboration ...
func NewCollaboration() CollaborationRepository {
	return &Collaboration{}
}

/*
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

/*** Make Elements ***/

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
	collaboration.Participant = make(PARTICIPANT_SLC, num)
}

// SetMessageFlow ...
func (collaboration *Collaboration) SetMessageFlow(num int) {
	collaboration.MessageFlow = make([]flow.MessageFlow, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (collaboration Collaboration) GetID() impl.STR_PTR {
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
func (collaboration Collaboration) GetParticipant(num int) PARTICIPANT_PTR {
	return &collaboration.Participant[num]
}

// GetMessageFlow ...
func (collaboration Collaboration) GetMessageFlow(num int) *flow.MessageFlow {
	return &collaboration.MessageFlow[num]
}
