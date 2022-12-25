package marker

import "fmt"

// MessageRepository ...
type MessageRepository interface {
	MarkerBase
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

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (msg *Message) SetID(typ string, suffix interface{}) {
	switch typ {
	case "message":
		msg.ID = fmt.Sprintf("Message_%v", suffix)
		break
	case "id":
		msg.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (msg *Message) SetName(suffix string) {
	msg.Name = fmt.Sprintf("Message_%s", suffix)
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (msg Message) GetID() *string {
	return &msg.ID
}

// GetName ...
func (msg Message) GetName() *string {
	return &msg.Name
}
