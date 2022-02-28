package models

import (
	"fmt"
	"strconv"
)

// Plane ...
type Plane struct {
	ID      string  `xml:"id,attr" json:"-"`
	Element string  `xml:"bpmnElement,attr" json:"-"`
	Shape   []Shape `xml:"bpmndi:BPMNShape" json:"-"`
	Edge    []Edge  `xml:"bpmndi:BPMNEdge" json:"-"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (plane *Plane) SetID(num int64) {
	plane.ID = "BPMNPlane_" + strconv.FormatInt(num, 16)
}

// SetElement ...
func (plane *Plane) SetElement(typ string, suffix string) {
	switch typ {
	case "process":
		plane.Element = "Process_" + suffix
	case "collaboration":
		plane.Element = "Collaboration_" + suffix
	}
}

// SetAttrProcessElement ...
func (plane *Plane) SetAttrProcessElement(suffix string) {
	plane.Element = fmt.Sprintf("Process_%s", suffix)
}

// SetAttrCollaborationElement ...
func (plane *Plane) SetAttrCollaborationElement(suffix string) {
	plane.Element = fmt.Sprintf("Collaboration_%s", suffix)
}

/* Elements */

/** BPMNDI **/

// SetShape ...
func (plane *Plane) SetShape(num int) {
	plane.Shape = make([]Shape, num)
}

// SetEdge ...
func (plane *Plane) SetEdge(num int) {
	plane.Edge = make([]Edge, num)
}
