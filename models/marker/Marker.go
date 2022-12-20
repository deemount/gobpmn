package marker

type Type_Incoming IncomingRepository
type Type_Outgoing OutgoingRepository
type Type_Sequence_Flow SequenceFlowRepository
type Type_Message_Flow MessageFlowRepository
type Type_Signal SignalRepository

type MarkerRepository[
	I Type_Incoming,
	O Type_Outgoing,
	S Type_Sequence_Flow,
	M Type_Message_Flow,
	SI Type_Signal] interface {
	IncomingRepository() I
	OutgoingRepository() O
	SequenceFlowRepository() S
	MessageFlowRepository() M
	SignalRepository() SI
}

type Marker[
	I Type_Incoming,
	O Type_Outgoing,
	S Type_Sequence_Flow,
	M Type_Message_Flow,
	SI Type_Signal] struct {
	MarkerRepository[I, O, S, M, SI]
}

func NewMarker() MarkerRepository[
	Type_Incoming,
	Type_Outgoing,
	Type_Sequence_Flow,
	Type_Message_Flow,
	Type_Signal] {
	return &Marker[
		Type_Incoming,
		Type_Outgoing,
		Type_Sequence_Flow,
		Type_Message_Flow,
		Type_Signal]{}
}
