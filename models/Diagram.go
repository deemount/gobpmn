package models

// Diagram ...
type Diagram struct {
	ID    string `xml:"id,attr"`
	Plane Plane  `xml:"bpmndi:BPMNPlane"`
}

// SetID ...
func (dia *Diagram) SetID() {
	dia.ID = "BPMNDiagram_"
}
