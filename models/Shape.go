package models

// Shape ...
type Shape struct {
	ID           string   `xml:"id,attr"`
	Element      string   `xml:"bpmnElement,attr"`
	IsHorizontal bool     `xml:"isHorizontal,attr,omitempty"`
	Bounds       []Bounds `xml:"dc:Bounds"`
	Label        []Label  `xml:"bpmndi:BPMNLabel"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (shape *Shape) SetID(typ string, suffix string) {
	switch typ {
	case "collaboration":
		shape.ID = "Participant_" + suffix + "_di"
	}

}

// SetElement ...
func (shape *Shape) SetElement(typ string, suffix string) {
	switch typ {
	case "collaboration":
		shape.Element = "Participant_" + suffix
	}
}

// SetIsHorizontal ...
func (shape *Shape) SetIsHorizontal(isHorizontal bool) {
	shape.IsHorizontal = isHorizontal
}

/* Elements */

/** DC **/

// SetBounds ...
func (shape *Shape) SetBounds() {
	shape.Bounds = make([]Bounds, 1)
}

/** BPMNDI **/

// SetLabel ...
func (shape *Shape) SetLabel() {
	shape.Label = make([]Label, 1)
}
