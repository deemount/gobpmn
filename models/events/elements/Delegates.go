package elements

import (
	"fmt"
)

// SetID ...
func SetID(typ string, suffix interface{}) string {
	var r string
	switch typ {
	case "event":
		r = fmt.Sprintf("Event_%v", suffix)
		break
	case "counter":
		r = fmt.Sprintf("StartEvent_%d", suffix)
		break
	case "id":
		r = fmt.Sprintf("%s", suffix)
		break
	}
	return r
}

// SetIntermediateCatchEvent ...
func SetIntermediateCatchEvent(e *IntermediateCatchEvent, name string, hash ...string) {
	e.SetID("event", hash[0])
	e.SetName(name)
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(hash[1])
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(hash[2])
}

// SetStartEvent ...
func SetStartEvent(e *StartEvent, name string, hash ...string) {
	e.SetID("event", hash[0])
	e.SetName(name)
	e.SetOutgoing(1)
	e.GetOutgoing(0).SetFlow(hash[1])
}

// SetEndEvent ...
func SetEndEvent(e *EndEvent, name string, hash ...string) {
	e.SetID("event", hash[0])
	e.SetName(name)
	e.SetIncoming(1)
	e.GetIncoming(0).SetFlow(hash[1])
}
