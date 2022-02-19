package models

import "fmt"

// MessageEventDefinition ...
type MessageEventDefinition struct {
	ID     string `xml:"id,attr,omitempty"`
	MsgRef string `xml:"messageRef,attr,omitempty"`
}

// SetID ...
func (med *MessageEventDefinition) SetID(suffix string) {
	med.ID = fmt.Sprintf("MessageEventDefinition_%s", suffix)
}

// SetMsgRef ...
func (med *MessageEventDefinition) SetMsgRef(suffix string) {
	med.MsgRef = fmt.Sprintf("Message_%s", suffix)
}
