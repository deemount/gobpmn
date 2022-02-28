package models

import "fmt"

type Message struct {
	ID   string `xml:"id,attr,omitempty" json:"id"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

// SetID ...
func (msg *Message) SetID(suffix string) {
	msg.ID = fmt.Sprintf("Message_%s", suffix)
}

// SetName ...
func (msg *Message) SetName(suffix string) {
	msg.Name = fmt.Sprintf("Message_%s", suffix)
}
