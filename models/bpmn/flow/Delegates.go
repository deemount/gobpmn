package flow

// SetSequenceFlow ...
func SetSequenceFlow(p DelegateParameter) {
	p.SF.SetID(p.T, p.H[0])
	// Notice: A label must set for name of sequenceflow
	if p.N != "" {
		p.SF.SetName(p.N)
	}
	p.SF.SetSourceRef(p.ST, p.H[1])
	p.SF.SetTargetRef(p.TT, p.H[2])
}
