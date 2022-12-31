package canvas

import (
	"fmt"

	"github.com/deemount/gobpmn/models/compulsion"
)

// NewEdge ...
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
func (edge Edge) GetID() compulsion.STR_PTR {
	return &edge.ID
}

// GetElement ...
func (edge Edge) GetElement() compulsion.STR_PTR {
	return &edge.Element
}

/* Elements */

/** BPMNDI **/

// GetWaypoint ...
func (edge Edge) GetWaypoint(num int) *Waypoint {
	return &edge.Waypoint[num]
}

// GetLabel ...
func (edge Edge) GetLabel() *Label {
	return &edge.Label[0]
}
