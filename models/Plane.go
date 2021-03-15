package models

import "strconv"

// Plane ...
type Plane struct {
	ID      string  `xml:"id,attr"`
	Element string  `xml:"bpmnElement,attr"`
	Edge    Edge    `xml:"bpmndi:BPMNEdge"`
	Shape   []Shape `xml:"bpmndi:BPMNShape"`
}

// SetID ...
func (p *Plane) SetID(num int64) {
	p.ID = "BPMNPlane_" + strconv.FormatInt(num, 16)
}

// SetElement ...
func (p *Plane) SetElement(name, suffix string) {
	p.Element = name + "_" + suffix
}

// SetProcessElement ...
func (p *Plane) SetProcessElement(suffix string) {
	p.Element = "Process_" + suffix
}

// SetCollaborationElement ...
func (p *Plane) SetCollaborationElement(suffix string) {
	p.Element = "Collaboration_" + suffix
}
