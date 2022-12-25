package canvas

import "fmt"

// Edge ...
type EdgeRepository interface {
	CanvasBase

	SetWaypoint()
	GetWaypoint() (*Waypoint, *Waypoint)

	SetLabel()
	GetLabel() *Label
}

// Edge ...
type Edge struct {
	ID       string     `xml:"id,attr" json:"-"`
	Element  string     `xml:"bpmnElement,attr" json:"-"`
	Waypoint []Waypoint `xml:"di:waypoint" json:"-"`
	Label    []Label    `xml:"bpmndi:BPMNLabel,omitempty" json:"-"`
}

func NewEdge() EdgeRepository {
	return &Edge{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (edge *Edge) SetID(typ string, suffix interface{}) {
	switch typ {
	case "association":
		edge.ID = fmt.Sprintf("Association_%s_di", suffix)
	case "flow":
		edge.ID = fmt.Sprintf("Flow_%s_di", suffix)
	}
}

// SetElement ...
func (edge *Edge) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "association":
		edge.Element = fmt.Sprintf("Association_%s", suffix)
	case "flow":
		edge.Element = fmt.Sprintf("Flow_%s", suffix)
	}
}

/*** Make Elements ***/

/** BPMNDI **/

// SetWaypoint ...
func (edge *Edge) SetWaypoint() {
	edge.Waypoint = make([]Waypoint, 2)
}

// SetLabel ...
func (edge *Edge) SetLabel() {
	edge.Label = make([]Label, 1)
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (edge Edge) GetID() STR_PTR {
	return &edge.ID
}

// GetElement ...
func (edge Edge) GetElement() STR_PTR {
	return &edge.Element
}

/* Elements */

/** BPMNDI **/

// GetWaypoint ...
func (edge Edge) GetWaypoint() (*Waypoint, *Waypoint) {
	return &edge.Waypoint[0], &edge.Waypoint[1]
}

// GetLabel ...
func (edge Edge) GetLabel() *Label {
	return &edge.Label[0]
}
