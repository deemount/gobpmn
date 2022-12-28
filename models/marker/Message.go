package marker

import "fmt"

// NewMessage ...
func NewMessage() MessageRepository {
	return &Message{}
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
func (msg Message) GetID() STR_PTR {
	return &msg.ID
}

// GetName ...
func (msg Message) GetName() STR_PTR {
	return &msg.Name
}
