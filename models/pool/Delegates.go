package pool

func SetParticipant(e *Participant, typID, typProcessRef, name string, hash ...string) {
	e.SetID(typID, hash[0])
	e.SetName(name)
	e.SetProcessRef(typProcessRef, hash[1])
}
