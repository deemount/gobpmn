package tasks

<<<<<<< HEAD
import "github.com/deemount/gobpmn/models/bpmn/canvas"

=======
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
// SetTask ...
// e *Task, name string, hash ...string
func SetTask(p DelegateParameter) {
	p.TA.SetID(p.T, p.H[0])
	p.TA.SetName(p.N)
	p.TA.SetIncoming(1)
	p.TA.GetIncoming(0).SetFlow(p.H[1])
	p.TA.SetOutgoing(1)
	p.TA.GetOutgoing(0).SetFlow(p.H[2])
<<<<<<< HEAD
	if p.SH != nil {
		canvas.SetShape(
			canvas.DelegateParameter{S: p.SH, T: p.T, H: p.H[0], B: p.BS})
	}
=======
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}
