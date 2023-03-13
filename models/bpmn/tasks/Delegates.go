package tasks

import "github.com/deemount/gobpmn/models/bpmn/canvas"

// SetTask ...
// e *Task, name string, hash ...string
func SetTask(p DelegateParameter) {

	if p.T == "" {
		p.T = "activity"
	}

	p.TA.SetID(p.T, p.H[0])
	p.TA.SetName(p.N)
	p.TA.SetIncoming(1)
	p.TA.GetIncoming(0).SetFlow(p.H[1])
	p.TA.SetOutgoing(1)
	p.TA.GetOutgoing(0).SetFlow(p.H[2])
	if p.SH != nil {
		canvas.SetShape(
			canvas.DelegateParameter{S: p.SH, T: p.T, H: p.H[0], WPPREV: p.WPPREV, B: p.BS})
	}
}
