package definitions

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewMessageEventDefinition ...
func NewMessageEventDefinition() MessageEventDefinitionRepository {
	return &MessageEventDefinition{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (med *MessageEventDefinition) SetID(typ string, suffix interface{}) {
	med.ID = SetID(typ, suffix)
}

// SetMsgRef ...
func (med *MessageEventDefinition) SetMsgRef(suffix string) {
	med.MsgRef = fmt.Sprintf("Message_%s", suffix)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (med MessageEventDefinition) GetID() impl.STR_PTR {
	return &med.ID
}

// GetMsgRef ...
func (med MessageEventDefinition) GetMsgRef() impl.STR_PTR {
	return &med.MsgRef
}
