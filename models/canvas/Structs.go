package canvas

// Bounds ...
type Bounds struct {
	X      int `xml:"x,attr,omitempty"`
	Y      int `xml:"y,attr,omitempty"`
	Width  int `xml:"width,attr,omitempty"`
	Height int `xml:"height,attr,omitempty"`
}

// Waypoint ...
type Waypoint struct {
	X int `xml:"x,attr" json:"-"`
	Y int `xml:"y,attr" json:"-"`
}

// Diagram ...
type Diagram struct {
	ID          string  `xml:"id,attr"`
	Description string  `xml:"-" json:"-"`
	Plane       []Plane `xml:"bpmndi:BPMNPlane,omitempty"`
}

// TDiagram ...
type TDiagram struct {
	ID          string   `xml:"id,attr"`
	Description string   `xml:"-" json:"-"`
	Plane       []TPlane `xml:"BPMNPlane,omitempty"`
}

// Edge ...
type Edge struct {
	ID       string     `xml:"id,attr" json:"-"`
	Element  string     `xml:"bpmnElement,attr" json:"-"`
	Waypoint []Waypoint `xml:"di:waypoint" json:"-"`
	Label    []Label    `xml:"bpmndi:BPMNLabel,omitempty" json:"-"`
}

// TEdge ...
type TEdge struct {
	ID       string     `xml:"id,attr" json:"-"`
	Element  string     `xml:"bpmnElement,attr" json:"-"`
	Waypoint []Waypoint `xml:"waypoint" json:"-"`
	Label    []TLabel   `xml:"BPMNLabel,omitempty" json:"-"`
}

// Label ...
type Label struct {
	Bounds []Bounds `xml:"dc:Bounds,omitempty" json:"-"`
}

// TLabel ...
type TLabel struct {
	Bounds []Bounds `xml:"Bounds,omitempty" json:"-"`
}

// Plane ...
type Plane struct {
	ID          string  `xml:"id,attr" json:"-"`
	Element     string  `xml:"bpmnElement,attr" json:"-"`
	Description string  `xml:"-" json:"-"`
	Shape       []Shape `xml:"bpmndi:BPMNShape" json:"-"`
	Edge        []Edge  `xml:"bpmndi:BPMNEdge" json:"-"`
}

// TPlane ...
type TPlane struct {
	ID          string   `xml:"id,attr" json:"-"`
	Element     string   `xml:"bpmnElement,attr" json:"-"`
	Description string   `xml:"-" json:"-"`
	Shape       []TShape `xml:"BPMNShape" json:"-"`
	Edge        []TEdge  `xml:"BPMNEdge" json:"-"`
}

// Shape ...
type Shape struct {
	ID           string   `xml:"id,attr" json:"-"`
	Element      string   `xml:"bpmnElement,attr" json:"-"`
	IsHorizontal bool     `xml:"isHorizontal,attr,omitempty" json:"-"`
	Bounds       []Bounds `xml:"dc:Bounds" json:"-"`
	Label        []Label  `xml:"bpmndi:BPMNLabel" json:"-"`
}

// TShape ...
type TShape struct {
	ID           string   `xml:"id,attr" json:"-"`
	Element      string   `xml:"bpmnElement,attr" json:"-"`
	IsHorizontal bool     `xml:"isHorizontal,attr,omitempty" json:"-"`
	Bounds       []Bounds `xml:"Bounds" json:"-"`
	Label        []TLabel `xml:"BPMNLabel" json:"-"`
}
