package canvas

// canvas
// Shape: d *Shape, typ string, hash string, b Bounds
// Edge: d *Edge, typ string, hash string, w []Waypoint
// Label: d *Edge, b Bounds
// Pool: e *Shape, typ string, isHorizontal bool, hash string, b Bounds
type DelegateParameter struct {
	S *Shape
	E *Edge
	T string
	H string
	I bool
	W []Waypoint
	B Bounds
}

/*
 * Elementary
 */

// Bounds ...
type Bounds struct {
	X      int `xml:"x,attr,omitempty" json:"x,omitempty"`
	Y      int `xml:"y,attr,omitempty" json:"y,omitempty"`
	Width  int `xml:"width,attr,omitempty"`
	Height int `xml:"height,attr,omitempty"`
}

// Waypoint ...
type Waypoint struct {
	X int `xml:"x,attr" json:"x,omitempty"`
	Y int `xml:"y,attr" json:"y,omitempty"`
}

// Diagram ...
type Diagram struct {
	ID          string  `xml:"id,attr" json:"id,omitempty"`
	Description string  `xml:"-" json:"-"`
	Plane       []Plane `xml:"bpmndi:BPMNPlane,omitempty" json:"plane,omitempty"`
}

// TDiagram ...
type TDiagram struct {
	ID          string   `xml:"id,attr" json:"id,omitempty"`
	Description string   `xml:"-" json:"-"`
	Plane       []TPlane `xml:"BPMNPlane,omitempty" json:"plane,omitempty"`
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

// Label ...
type Label struct {
	Bounds []Bounds `xml:"dc:Bounds,omitempty" json:"label,omitempty"`
}

// TLabel ...
type TLabel struct {
	Bounds []Bounds `xml:"Bounds,omitempty" json:"bounds,omitempty"`
}

// Plane ...
type Plane struct {
	ID          string  `xml:"id,attr" json:"id,omitempty"`
	Element     string  `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Description string  `xml:"-" json:"-"`
	Shape       []Shape `xml:"bpmndi:BPMNShape" json:"shape,omitempty"`
	Edge        []Edge  `xml:"bpmndi:BPMNEdge" json:"edge,omitempty"`
}

// TPlane ...
type TPlane struct {
	ID          string   `xml:"id,attr" json:"id,omitempty"`
	Element     string   `xml:"bpmnElement,attr" json:"bpmnElement,omitempty"`
	Description string   `xml:"-" json:"-"`
	Shape       []TShape `xml:"BPMNShape" json:"shape,omitempty"`
	Edge        []TEdge  `xml:"BPMNEdge" json:"edge,omitempty"`
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
