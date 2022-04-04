package models

import (
	"fmt"
	"strconv"
)

// Plane ...
type Plane struct {
	ID          string  `xml:"id,attr" json:"-"`
	Element     string  `xml:"bpmnElement,attr" json:"-"`
	Description string  `xml:"-" json:"-"`
	Shape       []Shape `xml:"bpmndi:BPMNShape" json:"-"`
	Edge        []Edge  `xml:"bpmndi:BPMNEdge" json:"-"`
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
		break
	case "collaboration":
		plane.Element = "Collaboration_" + suffix
		break
	case "id":
		plane.Element = fmt.Sprintf("%s", suffix)
		break
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

/* Specification */

// Description ...
func (plane Plane) GetDescription() string {
	plane.Description = `
	A BPMNPlane specializes DI::Plane and redefines its model element
	reference to be of the (BPMN) BaseElement. A BPMNPlane can only reference
	a BaseElement of the types: Process, SubProcess, AdHocSubProcess, Transaction,
	Collaboration, Choreography or SubChoreography.

	BPMNPlane element is always owned by a BPMNDiagram and represents the root
	diagram element of that diagram. The plane represents a 2 dimensional surface
	with an origin at (0,0) along the x and y axes with increasing coordinates to
	the right and bottom. Only positive coordinates are allowed for diagram elements
	that are nested in a BPMNPlane. This means that the union of all the nested
	elements' bounds is deemed to be located at the plane's origin point.
	`
	return fmt.Sprintf("%s", plane.Description)
}
