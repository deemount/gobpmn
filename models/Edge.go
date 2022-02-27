package models

import "fmt"

// Edge ...
type Edge struct {
	ID       string     `xml:"id,attr"`
	Element  string     `xml:"bpmnElement,attr"`
	Waypoint []Waypoint `xml:"di:waypoint"`
}

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

/* Elements */

/** BPMNDI **/

// SetWaypoint ...
func (edge *Edge) SetWaypoint() {
	edge.Waypoint = make([]Waypoint, 2)
}
