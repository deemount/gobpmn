package models

import "strconv"

// Plane ...
type Plane struct {
	ID      string  `xml:"id,attr"`
	Element string  `xml:"bpmnElement,attr"`
	Shape   []Shape `xml:"bpmndi:BPMNShape"`
	Edge    []Edge  `xml:"bpmndi:BPMNEdge"`
}

// SetID ...
func (p *Plane) SetID(num int64) {
	p.ID = "BPMNPlane_" + strconv.FormatInt(num, 16)
}

// SetElement ...
func (p *Plane) SetElement(typ string, suffix string) {
	switch typ {
	case "process":
		p.Element = "Process_" + suffix
	case "collaboration":
		p.Element = "Collaboration_" + suffix
	}
}

// SetCollaborationElement ...
func (p *Plane) SetCollaborationElement(suffix string) {
	p.Element = "Collaboration_" + suffix
}
