package pool

// e *Participant, typ, typProcessRef, name string, hash ...string
func SetParticipant(p DelegateParameter) {
	p.PPT.SetID(p.T, p.H[0])
	p.PPT.SetName(p.N)
	p.PPT.SetProcessRef(p.PR, p.H[1])
}
