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
	case "startevent":
		r = fmt.Sprintf("StartEvent_%v", suffix)
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
// e *StartEvent, name string, hash ...string
func SetStartEvent(p DelegateParameter) {
	p.SE.SetID(p.T, p.H[0])
	p.SE.SetName(p.N)
	p.SE.SetOutgoing(1)
	p.SE.GetOutgoing(0).SetFlow(p.H[1])
}

// SetEndEvent ...
// e *EndEvent, name string, hash ...string
func SetEndEvent(p DelegateParameter) {
	p.EE.SetID(p.T, p.H[0])
	p.EE.SetName(p.N)
	p.EE.SetIncoming(1)
	p.EE.GetIncoming(0).SetFlow(p.H[1])
}
