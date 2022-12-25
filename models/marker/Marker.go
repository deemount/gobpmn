package marker

type MarkerRepository interface {
	IncomingRepository
	OutgoingRepository
	SequenceFlowRepository
	MessageFlowRepository
	SignalRepository
}

type Marker struct {
	MarkerRepository
}

func NewMarker() MarkerRepository {
	return &Marker{}
}
