package pool

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/loop"
)

// NewParticipant ...
func NewParticipant() ParticipantRepository {
	return &Participant{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (participant *Participant) SetID(typ string, suffix interface{}) {
	switch typ {
	case "participant":
		participant.ID = fmt.Sprintf("Participant_%v", suffix)
		break
	case "id":
		participant.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (participant *Participant) SetName(name string) {
	participant.Name = name
}

// SetProcessRef ...
func (participant *Participant) SetProcessRef(typ string, suffix string) {
	switch typ {
	case "process":
		participant.ProcessRef = fmt.Sprintf("Process_%s", suffix)
		break
	case "id":
		participant.ProcessRef = fmt.Sprintf("%s", suffix)
		break
	}
}

/*** Make Elements ***/

/** BPMN **/

/** Documentation **/

// SetDocumentation ...
func (participant *Participant) SetDocumentation() {
	participant.Documentation = make([]attributes.Documentation, 1)
}

// SetParticipantMultiplicity ...
func (participant *Participant) SetParticipantMultiplicity() {
	participant.ParticipantMultiplicity = make([]loop.ParticipantMultiplicity, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (participant Participant) GetID() compulsion.STR_PTR {
	return &participant.ID
}

// GetName ...
func (participant Participant) GetName() compulsion.STR_PTR {
	return &participant.Name
}

// GetProcessRef ...
func (participant Participant) GetProcessRef() *string {
	return &participant.ProcessRef
}

/* Elements */

/** BPMN **/

// SetDocumentation ...
func (participant Participant) GetDocumentation() *attributes.Documentation {
	return &participant.Documentation[0]
}

// GetParticipantMultiplicity ...
func (participant Participant) GetParticipantMultiplicity() *loop.ParticipantMultiplicity {
	return &participant.ParticipantMultiplicity[0]
}
