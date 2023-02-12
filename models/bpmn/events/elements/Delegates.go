package elements

import (
	"fmt"
<<<<<<< HEAD
	"strings"

	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/events/definitions"
)

var (
	structMessage = "message"
=======
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
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
<<<<<<< HEAD
func SetIntermediateCatchEvent(p DelegateParameter) {
	p.ICE.SetID(p.T, p.H[0])
	p.ICE.SetName(p.N)
	p.ICE.SetIncoming(1)
	p.ICE.GetIncoming(0).SetFlow(p.H[1])
	p.ICE.SetOutgoing(1)
	p.ICE.GetOutgoing(0).SetFlow(p.H[2])

	if p.ISTED {
		p.ICE.SetTimerEventDefinition()
		definitions.SetTimerEventDefinition(
			definitions.DelegateParameter{
				TED: p.ICE.GetTimerEventDefinition(),
				TD:  p.TD,
				H:   []string{p.TEDH}})
	}

	if p.SH != nil {
		canvas.SetShape(
			canvas.DelegateParameter{S: p.SH, T: p.T, H: p.H[0], B: p.BS})
	}

}

// SetStartEvent ...
=======
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
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
func SetStartEvent(p DelegateParameter) {
	p.SE.SetID(p.T, p.H[0])
	p.SE.SetName(p.N)
	p.SE.SetOutgoing(1)
	p.SE.GetOutgoing(0).SetFlow(p.H[1])
<<<<<<< HEAD
	if p.SH != nil {
		canvas.SetShape(
			canvas.DelegateParameter{S: p.SH, T: p.T, H: p.H[0], B: p.BS})
	}
}

// SetEndEvent ...
=======
}

// SetEndEvent ...
// e *EndEvent, name string, hash ...string
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
func SetEndEvent(p DelegateParameter) {
	p.EE.SetID(p.T, p.H[0])
	p.EE.SetName(p.N)
	p.EE.SetIncoming(1)
	p.EE.GetIncoming(0).SetFlow(p.H[1])
<<<<<<< HEAD
	if p.SH != nil {
		canvas.SetShape(
			canvas.DelegateParameter{S: p.SH, T: p.T, H: p.H[0], B: p.BS})
	}
}

// SetMessage ...
func SetMessage(p DelegateParameter) {}

// IsMessage ...
func IsMessage(field string) bool { return strings.ToLower(field) == structMessage }
=======
}
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
