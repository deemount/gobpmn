package marker

import "fmt"

// MessageRepository ...
type MessageRepository interface {
	SetID(suffix string)
	SetName(suffix string)
	GetID() string
	GetName() string
}

// Message ...
type Message struct {
	ID   string `xml:"id,attr,omitempty" json:"id"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

// TMessage ...
type TMessage struct {
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

// SetID ...
func (msg *Message) SetID(suffix string) {
	msg.ID = fmt.Sprintf("Message_%s", suffix)
}

// SetName ...
func (msg *Message) SetName(suffix string) {
	msg.Name = fmt.Sprintf("Message_%s", suffix)
}

// GetID ...
func (msg Message) GetID() *string {
	return &msg.ID
}

// GetName ...
func (msg Message) GetName() *string {
	return &msg.Name
}
