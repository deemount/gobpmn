package models

// Plane ...
type Plane struct {
	ID      string `xml:"id,attr"`
	Element string `xml:"bpmnElement,attr"`
}

// SetID ...
func (p *Plane) SetID() {
	p.ID = "BPMNPlane_"
}

// SetElement ...
func (p *Plane) SetElement() {
	p.Element = "Process_"
}
