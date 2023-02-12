package flow

<<<<<<< HEAD
import "github.com/deemount/gobpmn/models/bpmn/canvas"

=======
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
// SetSequenceFlow ...
func SetSequenceFlow(p DelegateParameter) {
	p.SF.SetID(p.T, p.H[0])
	// Notice: A label must set for name of sequenceflow
	if p.N != "" {
		p.SF.SetName(p.N)
<<<<<<< HEAD
		canvas.SetLabel(
			canvas.DelegateParameter{E: p.ED, B: p.BS})
	}
	p.SF.SetSourceRef(p.ST, p.H[1])
	p.SF.SetTargetRef(p.TT, p.H[2])
	if p.ED != nil {
		canvas.SetEdge(
			canvas.DelegateParameter{E: p.ED, T: p.T, H: p.H[0], W: p.WP})
	}
}

// SetMessageFlow ...
func SetMessageFlow(p DelegateParameter) {
	p.MF.SetID(p.T, p.H[0])
	p.MF.SetName(p.N)
	p.MF.SetSourceRef(p.ST, p.H[1])
	p.MF.SetTargetRef(p.TT, p.H[2])
=======
	}
	p.SF.SetSourceRef(p.ST, p.H[1])
	p.SF.SetTargetRef(p.TT, p.H[2])
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}
