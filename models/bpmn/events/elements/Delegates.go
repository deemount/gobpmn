package elements

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/canvas"
	"github.com/deemount/gobpmn/models/bpmn/events/definitions"
)

var (
	structMessage = "message"
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
			canvas.DelegateParameter{S: p.SH, T: p.T, H: p.H[0], B: p.BS, WPPREV: p.WPPREV})
	}

}

// SetStartEvent ...
func SetStartEvent(p DelegateParameter) {

	if p.T == "" {
		p.T = "event"
	}

	p.SE.SetID(p.T, p.H[0])
	p.SE.SetName(p.N)
	p.SE.SetOutgoing(1)                 // startevent has by default one outgoing
	p.SE.GetOutgoing(0).SetFlow(p.H[1]) // set second hash value to flow
	if p.SH != nil {
		canvas.SetShape(
			canvas.DelegateParameter{S: p.SH, T: p.T, H: p.H[0], B: p.BS, BSPTR: p.BSPTR})
	}
}

// SetEndEvent ...
func SetEndEvent(p DelegateParameter) {

	if p.T == "" {
		p.T = "event"
	}

	p.EE.SetID(p.T, p.H[0])
	p.EE.SetName(p.N)
	p.EE.SetIncoming(1)
	p.EE.GetIncoming(0).SetFlow(p.H[1])
	if p.SH != nil {
		canvas.SetShape(
			canvas.DelegateParameter{S: p.SH, T: p.T, H: p.H[0], B: p.BS, WPPREV: p.WPPREV})
	}
}

// SetMessage ...
func SetMessage(p DelegateParameter) {}
