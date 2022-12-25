package definitions

import (
	"fmt"

	"github.com/deemount/gobpmn/models/events/eventsbase"
)

// SignalEventDefinitionRepository ...
type SignalEventDefinitionRepository interface {
	DefinitionsID

	SetSignalRef(suffix string)
	GetSignalRef() *string
}

// SignalEventDefinition ...
type SignalEventDefinition struct {
	ID        string `xml:"id,attr,omitempty" json:"id"`
	SignalRef string `xml:"signalRef,attr,omitempty" json:"signalRef,omitempty"`
}

func NewSignalEventDefinition() SignalEventDefinitionRepository {
	return &SignalEventDefinition{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (sed *SignalEventDefinition) SetID(typ string, suffix interface{}) {
	switch typ {
	case "sed":
		sed.ID = fmt.Sprintf("SignalEventDefinition_%v", suffix)
		break
	case "id":
		sed.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetSignalRef ...
func (sed *SignalEventDefinition) SetSignalRef(suffix string) {
	sed.SignalRef = fmt.Sprintf("Signal_%s", suffix)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (sed SignalEventDefinition) GetID() eventsbase.STR_PTR {
	return &sed.ID
}

// GetSignalRef ...
func (sed SignalEventDefinition) GetSignalRef() *string {
	return &sed.SignalRef
}
