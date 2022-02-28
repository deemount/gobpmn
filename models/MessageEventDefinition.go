package models

import "fmt"

// MessageEventDefinition ...
type MessageEventDefinition struct {
	ID     string `xml:"id,attr,omitempty" json:"id"`
	MsgRef string `xml:"messageRef,attr,omitempty" json:"messageRef,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (med *MessageEventDefinition) SetID(suffix string) {
	med.ID = fmt.Sprintf("MessageEventDefinition_%s", suffix)
}

// SetMsgRef ...
func (med *MessageEventDefinition) SetMsgRef(suffix string) {
	med.MsgRef = fmt.Sprintf("Message_%s", suffix)
}
