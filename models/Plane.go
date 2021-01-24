package models

import "strconv"

// Plane ...
type Plane struct {
	ID      string `xml:"id,attr"`
	Element string `xml:"bpmnElement,attr"`
}

// SetID ...
func (p *Plane) SetID(num int64) {
	p.ID = "BPMNPlane_" + strconv.FormatInt(num, 16)
}

// SetElement ...
func (p *Plane) SetElement(suffix string) {
	p.Element = "Process_" + suffix
}
