package models

import "strconv"

// Diagram ...
type Diagram struct {
	ID    string `xml:"id,attr"`
	Plane Plane  `xml:"bpmndi:BPMNPlane,omitempty"`
}

// SetID ...
func (dia *Diagram) SetID(num int64) {
	dia.ID = "BPMNDiagram_" + strconv.FormatInt(num, 16)
}
