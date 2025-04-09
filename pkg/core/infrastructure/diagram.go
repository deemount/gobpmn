package infrastructure

import "fmt"

// Diagram ...
type Diagram struct {
	ID    string  `xml:"id,attr" json:"id,omitempty"`
	Plane []Plane `xml:"bpmndi:BPMNPlane,omitempty" json:"plane,omitempty"`
}

// TDiagram ...
type TDiagram struct {
	ID    string   `xml:"id,attr" json:"id,omitempty"`
	Plane []TPlane `xml:"BPMNPlane,omitempty" json:"plane,omitempty"`
}

// SetID ...
func (diagram *Diagram) SetID(typ string, suffix any) {
	switch typ {
	case "diagram":
		diagram.ID = fmt.Sprintf("BPMNDiagram_%v", suffix)
		break
	case "id":
		diagram.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetPlane ...
func (diagram *Diagram) SetPlane() {
	diagram.Plane = make([]Plane, 1)
}

// GetID ...
func (diagram Diagram) GetID() *string {
	return &diagram.ID
}

// GetPlane ...
func (diagram Diagram) GetPlane() *Plane {
	return &diagram.Plane[0]
}

// Plane ...
type Plane struct {
	ID      string  `xml:"id,attr" json:"id,omitempty"`
	Element string  `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Shape   []Shape `xml:"bpmndi:BPMNShape" json:"shape,omitempty"`
	Edge    []Edge  `xml:"bpmndi:BPMNEdge" json:"edge,omitempty"`
}

// TPlane ...
type TPlane struct {
	ID      string   `xml:"id,attr" json:"id,omitempty"`
	Element string   `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Shape   []TShape `xml:"BPMNShape" json:"shape,omitempty"`
	Edge    []TEdge  `xml:"BPMNEdge" json:"edge,omitempty"`
}

// SetID ...
func (plane *Plane) SetID(typ string, suffix any) {
	switch typ {
	case "plane":
		plane.ID = fmt.Sprintf("BPMNPlane_%d", suffix)
		break
	case "id":
		plane.ID = fmt.Sprintf("%s", suffix)
	}
}

// GetID ...
func (plane Plane) GetID() *string {
	return &plane.ID
}

// SetElement ...
func (plane *Plane) SetElement(typ string, suffix any) {
	switch typ {
	case "process":
		plane.Element = fmt.Sprintf("Process_%s", suffix)
		break
	case "collaboration":
		plane.Element = fmt.Sprintf("Collaboration_%s", suffix)
		break
	case "id":
		plane.Element = fmt.Sprintf("%s", suffix)
		break
	}
}

// GetElement ...
func (plane Plane) GetElement() *string {
	return &plane.Element
}

// SetAttrProcessElement ...
func (plane *Plane) SetAttrProcessElement(suffix string) {
	plane.Element = fmt.Sprintf("Process_%s", suffix)
}

// SetAttrCollaborationElement ...
func (plane *Plane) SetAttrCollaborationElement(suffix string) {
	plane.Element = fmt.Sprintf("Collaboration_%s", suffix)
}

// SetShape ...
func (plane *Plane) SetShape(num int) {
	plane.Shape = make([]Shape, num)
}

// GetShape ...
func (plane Plane) GetShape(num int) *Shape {
	return &plane.Shape[num]
}

// SetEdge ...
func (plane *Plane) SetEdge(num int) {
	plane.Edge = make([]Edge, num)
}

// GetEdge ...
func (plane Plane) GetEdge(num int) *Edge {
	return &plane.Edge[num]
}

// Shape ...
type Shape struct {
	ID              string   `xml:"id,attr" json:"id,omitempty"`
	Element         string   `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	IsHorizontal    bool     `xml:"isHorizontal,attr,omitempty" json:"isHorizontal,omitempty"`
	IsMarkerVisible bool     `xml:"isMarkerVisible,attr,omitempty" json:"isMarkerVisible,omitempty"`
	Bounds          []Bounds `xml:"dc:Bounds" json:"bounds,omitempty"`
	Label           []Label  `xml:"bpmndi:BPMNLabel" json:"label,omitempty"`
}

// TShape ...
type TShape struct {
	ID              string   `xml:"id,attr" json:"id,omitempty"`
	Element         string   `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	IsHorizontal    bool     `xml:"isHorizontal,attr,omitempty" json:"isHorizontal,omitemptyy"`
	IsMarkerVisible bool     `xml:"isMarkerVisible,attr,omitempty" json:"isMarkerVisible,omitempty"`
	Bounds          []Bounds `xml:"Bounds" json:"bounds,omitempty"`
	Label           []TLabel `xml:"BPMNLabel" json:"label,omitempty"`
}

// SetID ...
func (shape *Shape) SetID(typ string, suffix any) {
	switch typ {
	case "activity":
		shape.ID = fmt.Sprintf("Activity_%s_di", suffix)
		break
	case "collaboration":
		shape.ID = fmt.Sprintf("Participant_%s_di", suffix)
		break
	case "event":
		shape.ID = fmt.Sprintf("Event_%s_di", suffix)
		break
	case "participant":
		shape.ID = fmt.Sprintf("Participant_%s_di", suffix)
		break
	case "startevent":
		shape.ID = fmt.Sprintf("_BPMNShape_StartEvent_%v", suffix)
		break
	case "id":
		shape.ID = fmt.Sprintf("%s_di", suffix)
		break
	}
}

// GetID ...
func (shape Shape) GetID() *string {
	return &shape.ID
}

// SetElement ...
func (shape *Shape) SetElement(typ string, suffix any) {
	switch typ {
	case "activity":
		shape.Element = fmt.Sprintf("Activity_%s", suffix)
		break
	case "collaboration":
		shape.Element = fmt.Sprintf("Participant_%s", suffix)
		break
	case "event":
		shape.Element = fmt.Sprintf("Event_%s", suffix)
		break
	case "participant":
		shape.Element = fmt.Sprintf("Participant_%s", suffix)
		break
	case "startevent":
		shape.Element = fmt.Sprintf("StartEvent_%v", suffix)
		break
	case "id":
		shape.Element = fmt.Sprintf("%s", suffix)
		break
	}
}

// GetElement ...
func (shape Shape) GetElement() *string {
	return &shape.Element
}

// SetIsHorizontal ...
func (shape *Shape) SetIsHorizontal(isHorizontal bool) {
	shape.IsHorizontal = isHorizontal
}

// GetIsHorizontal ...
func (shape Shape) GetIsHorizontal() *bool {
	return &shape.IsHorizontal
}

// SetIsMarkerVisible ...
func (shape *Shape) SetIsMarkerVisible(isMarkerVisible bool) {
	shape.IsMarkerVisible = isMarkerVisible
}

// GetIsMarkerVisible ...
func (shape Shape) GetIsMarkerVisible() *bool {
	return &shape.IsMarkerVisible
}

// SetBounds ...
func (shape *Shape) SetBounds() {
	shape.Bounds = make([]Bounds, 1)
}

// GetBounds ...
func (shape Shape) GetBounds() *Bounds {
	return &shape.Bounds[0]
}

// SetLabel ...
func (shape *Shape) SetLabel() {
	shape.Label = make([]Label, 1)
}

// SetLabel ...
func (shape Shape) GetLabel() *Label {
	return &shape.Label[0]
}

// Edge ...
type Edge struct {
	ID       string     `xml:"id,attr" json:"-"`
	Element  string     `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Waypoint []Waypoint `xml:"di:waypoint" json:"waypoint,omitempty"`
	Label    []Label    `xml:"bpmndi:BPMNLabel,omitempty" json:"label,omitempty"`
}

// TEdge ...
type TEdge struct {
	ID       string     `xml:"id,attr" json:"-"`
	Element  string     `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Waypoint []Waypoint `xml:"waypoint" json:"waypoint,omitempty"`
	Label    []TLabel   `xml:"BPMNLabel,omitempty" json:"label,omitempty"`
}

// SetID ...
func (edge *Edge) SetID(typ string, suffix any) {
	switch typ {
	case "association":
		edge.ID = fmt.Sprintf("Association_%s_di", suffix)
	case "flow":
		edge.ID = fmt.Sprintf("Flow_%s_di", suffix)
	}
}

// GetID ...
func (edge Edge) GetID() *string {
	return &edge.ID
}

// SetElement ...
func (edge *Edge) SetElement(typ string, suffix any) {
	switch typ {
	case "association":
		edge.Element = fmt.Sprintf("Association_%s", suffix)
	case "flow":
		edge.Element = fmt.Sprintf("Flow_%s", suffix)
	}
}

// GetElement ...
func (edge Edge) GetElement() *string {
	return &edge.Element
}

// SetWaypoint ...
func (edge *Edge) SetWaypoint() {
	edge.Waypoint = make([]Waypoint, 2)
}

// SetWaypoints ...
func (edge *Edge) SetWaypoints(num int) {
	edge.Waypoint = make([]Waypoint, num)
}

// GetWaypoint ...
func (edge Edge) GetWaypoint(num int) *Waypoint {
	return &edge.Waypoint[num]
}

// SetLabel ...
func (edge *Edge) SetLabel() {
	edge.Label = make([]Label, 1)
}

// GetLabel ...
func (edge Edge) GetLabel() *Label {
	return &edge.Label[0]
}

// Label ...
type Label struct {
	Bounds []Bounds `xml:"dc:Bounds,omitempty" json:"label,omitempty"`
}

// TLabel ...
type TLabel struct {
	Bounds []Bounds `xml:"Bounds,omitempty" json:"bounds,omitempty"`
}

// SetBounds ...
func (label *Label) SetBounds() {
	label.Bounds = make([]Bounds, 1)
}

// GetBounds ...
func (label Label) GetBounds() *Bounds {
	return &label.Bounds[0]
}

// Bounds ...
type Bounds struct {
	X      int `xml:"x,attr,omitempty" json:"x,omitempty"`
	Y      int `xml:"y,attr,omitempty" json:"y,omitempty"`
	Width  int `xml:"width,attr,omitempty"`
	Height int `xml:"height,attr,omitempty"`
}

// SetCoordinates ...
func (bounds *Bounds) SetCoordinates(x, y int) {
	bounds.X = x
	bounds.Y = y
}

// GetCoordinates ...
func (bounds Bounds) GetCoordinates() (*int, *int) {
	return &bounds.X, &bounds.Y
}

// SetX ...
func (bounds *Bounds) SetX(x int) {
	bounds.X = x
}

// GetX ...
func (bounds Bounds) GetX() *int {
	return &bounds.X
}

// SetY ...
func (bounds *Bounds) SetY(y int) {
	bounds.Y = y
}

// GetY ...
func (bounds Bounds) GetY() *int {
	return &bounds.Y
}

// SetSize ...
func (bounds *Bounds) SetSize(width, height int) {
	bounds.Width = width
	bounds.Height = height
}

// GetSize ...
func (bounds Bounds) GetSize() (*int, *int) {
	return &bounds.Width, &bounds.Height
}

// SetWidth ...
func (bounds *Bounds) SetWidth(width int) {
	bounds.Width = width
}

// GetWidth ...
func (bounds Bounds) GetWidth() *int {
	return &bounds.Width
}

// SetHeight ...
func (bounds *Bounds) SetHeight(height int) {
	bounds.Height = height
}

// GetHeight ...
func (bounds Bounds) GetHeight() *int {
	return &bounds.Height
}

// Waypoint ...
type Waypoint struct {
	X int `xml:"x,attr" json:"x,omitempty"`
	Y int `xml:"y,attr" json:"y,omitempty"`
}

// SetX ...
func (waypoint *Waypoint) SetX(x int) {
	waypoint.X = x
}

// GetX ...
func (waypoint Waypoint) GetX() *int {
	return &waypoint.X
}

// SetY ...
func (waypoint *Waypoint) SetY(y int) {
	waypoint.Y = y
}

// GetY ...
func (waypoint Waypoint) GetY() *int {
	return &waypoint.Y
}

// SetCoordinates ...
func (waypoint *Waypoint) SetCoordinates(x, y int) {
	waypoint.X = x
	waypoint.Y = y
}

// GetCoordinates ...
func (waypoint Waypoint) GetCoordinates() (*int, *int) {
	return &waypoint.X, &waypoint.Y
}
