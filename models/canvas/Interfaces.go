package canvas

type CanvasID interface {
	SetID(typ string, suffix interface{})
	GetID() STR_PTR
}

type CanvasElement interface {
	SetElement(typ string, suffix interface{})
	GetElement() STR_PTR
}

type CanvasBase interface {
	CanvasID
	CanvasElement
}

// BoundsRepository ...
type BoundsRepository interface {
	SetX(x int)
	GetX() *int
	SetY(y int)
	GetY() *int
	SetCoordinates(x, y int)
	GetCoordinates() (*int, *int)

	SetSize(width, height int)
	GetSize() (*int, *int)
	SetWidth(width int)
	GetWidth() *int
	SetHeight(height int)
	GetHeight() *int
}

// WaypointRepository ...
type WaypointRepository interface {
	SetX(x int)
	GetX() *int
	SetY(y int)
	GetY() *int
	SetCoordinates(x, y int)
	GetCoordinates() (*int, *int)
}

// Diagram ...
type DiagramRepository interface {
	CanvasID

	SetPlane()
	GetPlane() *Plane

	GetDescription() string
}

// Edge ...
type EdgeRepository interface {
	CanvasBase

	SetWaypoint()
	GetWaypoint(num int) *Waypoint

	SetLabel()
	GetLabel() *Label
}

// Label ...
type LabelRepository interface {
	SetBounds()
	GetBounds() *Bounds
}

// PlaneRepository ...
type PlaneRepository interface {
	CanvasBase

	SetAttrProcessElement(suffix string)
	SetAttrCollaborationElement(suffix string)

	SetShape(num int)
	GetShape(num int) *Shape
	SetEdge(num int)
	GetEdge(num int) *Edge

	GetDescription() string
}

// ShapeRepository ...
type ShapeRepository interface {
	CanvasBase

	SetIsHorizontal(isHorizontal bool)
	GetIsHorizontal() *bool

	SetBounds()
	GetBounds() *Bounds

	SetLabel()
	GetLabel() *Label
}
