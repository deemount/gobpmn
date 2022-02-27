package models

import "strconv"

// Diagram ...
type Diagram struct {
	ID    string  `xml:"id,attr"`
	Plane []Plane `xml:"bpmndi:BPMNPlane,omitempty"`
}

/* Attributes */

// SetID ...
func (diagram *Diagram) SetID(num int64) {
	diagram.ID = "BPMNDiagram_" + strconv.FormatInt(num, 16)
}

/* Elements */

// SetPlane ...
func (diagram *Diagram) SetPlane() {
	diagram.Plane = make([]Plane, 1)
}
