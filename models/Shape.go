package models

import "fmt"

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
func (shape *Shape) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		shape.ID = fmt.Sprintf("Activity_%s_di", suffix)
	case "collaboration":
		shape.ID = fmt.Sprintf("Participant_%s_di", suffix)
	case "event":
		shape.ID = fmt.Sprintf("Event_%s_di", suffix)
	case "startevent":
		shape.ID = fmt.Sprintf("_BPMNShape_StartEvent_%d", suffix)
	}

}

// SetElement ...
func (shape *Shape) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		shape.Element = fmt.Sprintf("Activity_%s", suffix)
	case "collaboration":
		shape.Element = fmt.Sprintf("Participant_%s", suffix)
	case "event":
		shape.Element = fmt.Sprintf("Event_%s", suffix)
	case "startevent":
		shape.Element = fmt.Sprintf("StartEvent_%d", suffix)
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
