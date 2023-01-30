package canvas

import "github.com/deemount/gobpmn/models/bpmn/impl"

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
	impl.IFBaseCoordinates
	SetSize(width, height int)
	GetSize() (impl.INT_PTR, impl.INT_PTR)
	SetWidth(width int)
	GetWidth() impl.INT_PTR
	SetHeight(height int)
	GetHeight() impl.INT_PTR
}

// WaypointRepository ...
type WaypointRepository interface {
	impl.IFBaseCoordinates
}

// Diagram ...
type DiagramRepository interface {
	impl.IFBaseID

	SetPlane()
	GetPlane() *Plane

	GetDescription() string
}

// Edge ...
type EdgeRepository interface {
	impl.IFBaseID
	impl.IFBaseElement
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
	impl.IFBaseID
	impl.IFBaseElement

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
	impl.IFBaseID
	impl.IFBaseElement
	CanvasBoundsElements
	CanvasLabelElements

	SetIsHorizontal(isHorizontal bool)
	GetIsHorizontal() *bool
}
