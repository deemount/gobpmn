package events

import "fmt"

// MessageEventDefinitionRepository ...
type MessageEventDefinitionRepository interface {
	SetID(typ string, suffix interface{})
	SetMsgRef(suffix string)

	GetID() *string
	GetMsgRef() *string
}

// MessageEventDefinition ...
type MessageEventDefinition struct {
	ID     string `xml:"id,attr,omitempty" json:"id"`
	MsgRef string `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
}

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
func (med MessageEventDefinition) GetID() *string {
	return &med.ID
}

// GetMsgRef ...
func (med MessageEventDefinition) GetMsgRef() *string {
	return &med.MsgRef
}
