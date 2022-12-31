package canvas

import "github.com/deemount/gobpmn/models/compulsion"

type CanvasBoundsElements interface {
	SetBounds()
	GetBounds() *Bounds
}

type CanvasLabelElements interface {
	SetLabel()
	GetLabel() *Label
}

// BoundsRepository ...
type BoundsRepository interface {
	compulsion.IFBaseCoordinates

	SetSize(width, height int)
	GetSize() (*int, *int)
	SetWidth(width int)
	GetWidth() *int
	SetHeight(height int)
	GetHeight() *int
}

// WaypointRepository ...
type WaypointRepository interface {
	compulsion.IFBaseCoordinates
}

// Diagram ...
type DiagramRepository interface {
	compulsion.IFBaseID

	SetPlane()
	GetPlane() *Plane

	GetDescription() string
}

// Edge ...
type EdgeRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseElement
	CanvasLabelElements

	SetWaypoint()
	GetWaypoint(num int) *Waypoint
}

// Label ...
type LabelRepository interface {
	CanvasBoundsElements
}

// PlaneRepository ...
type PlaneRepository interface {
	compulsion.IFBaseID
	compulsion.IFBaseElement

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
	compulsion.IFBaseID
	compulsion.IFBaseElement
	CanvasBoundsElements
	CanvasLabelElements

	SetIsHorizontal(isHorizontal bool)
	GetIsHorizontal() *bool
}
