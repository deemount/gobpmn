package canvas

type CanvasRepository interface {
	BoundsRepository
	DiagramRepository
	EdgeRepository
	LabelRepository
	PlaneRepository
	ShapeRepository
	WaypointRepository
}

type Canvas struct {
	CanvasRepository
}

func NewCanvas() CanvasRepository {
	return &Canvas{}
}
