package flow

import "github.com/deemount/gobpmn/models/bpmn/canvas"

// SetSequenceFlow ...
func SetSequenceFlow(p DelegateParameter) {
	p.SF.SetID(p.T, p.H[0]) // set id with first hash
	// Notice: A label must set for name of sequenceflow
	if p.N != "" {
		p.SF.SetName(p.N)
		canvas.SetLabel(
			canvas.DelegateParameter{E: p.ED, B: p.BS})
	}
	p.SF.SetSourceRef(p.ST, p.H[1]) // set source ref with second hash
	p.SF.SetTargetRef(p.TT, p.H[2]) // set target ref with third hash
	if p.ED != nil {
		canvas.SetEdge(
			canvas.DelegateParameter{
				E:     p.ED,
				T:     p.T,
				BSPTR: p.BSPTR, // Bounds Pointer (getting the element size before >>this<< edge)
				ST:    p.ST,
				TT:    p.TT,
				H:     p.H[0],
				W:     p.WP})
	}
}

// SetMessageFlow ...
func SetMessageFlow(p DelegateParameter) {
	p.MF.SetID(p.T, p.H[0])
	p.MF.SetName(p.N)
	p.MF.SetSourceRef(p.ST, p.H[1])
	p.MF.SetTargetRef(p.TT, p.H[2])
}
