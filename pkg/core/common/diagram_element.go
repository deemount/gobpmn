package common

type diagramElement string

func (de diagramElement) String() string {
	return string(de)
}

const (
	plane diagramElement = "Plane"
	shape diagramElement = "Shape"
	edge  diagramElement = "Edge"
)
