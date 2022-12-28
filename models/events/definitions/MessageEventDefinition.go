package definitions

import (
	"fmt"
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
	switch typ {
	case "med":
		med.ID = fmt.Sprintf("MessageEventDefinition_%v", suffix)
		break
	case "id":
		med.ID = fmt.Sprintf("%s", suffix)
		break
	}
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
func (med MessageEventDefinition) GetID() STR_PTR {
	return &med.ID
}

// GetMsgRef ...
func (med MessageEventDefinition) GetMsgRef() *string {
	return &med.MsgRef
}
