package models

// Edge ...
type Edge struct {
	ID       string     `xml:"id,attr"`
	Element  string     `xml:"bpmnElement,attr"`
	Label    Label      `xml:"bpmndi:BPMNLabel,omitempty"`
	Waypoint []Waypoint `xml:"di:waypoint"`
}

// SetID ...
func (e *Edge) SetID(suffix string) {
	e.ID = "Flow_" + suffix + "_di"
}

// SetElement ...
func (e *Edge) SetElement(suffix string) {
	e.Element = "Flow_" + suffix
}
