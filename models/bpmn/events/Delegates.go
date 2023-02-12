package events

import "github.com/deemount/gobpmn/models/bpmn/events/definitions"

// SetTimerEvent ...
func SetTimerEvent(p DelegateParameter) {
	p.ICE.SetID(p.T, p.H[0])
	p.ICE.SetName(p.N)
	p.ICE.SetIncoming(1)
	p.ICE.GetIncoming(0).SetFlow(p.H[1])
	p.ICE.SetOutgoing(1)
	p.ICE.GetOutgoing(0).SetFlow(p.H[2])
	p.ICE.SetTimerEventDefinition()
	definitions.SetTimerEventDefinition(
		definitions.DelegateParameter{
			TED: p.ICE.GetTimerEventDefinition(),
			TD:  p.TD,
			H:   []string{p.TEDH}})
}
