package models

import "fmt"

// Shape ...
type Shape struct {
	ID           string   `xml:"id,attr" json:"-"`
	Element      string   `xml:"bpmnElement,attr" json:"-"`
	IsHorizontal bool     `xml:"isHorizontal,attr,omitempty" json:"-"`
	Bounds       []Bounds `xml:"dc:Bounds" json:"-"`
	Label        []Label  `xml:"bpmndi:BPMNLabel" json:"-"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (shape *Shape) SetID(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		shape.ID = fmt.Sprintf("Activity_%s_di", suffix)
		break
	case "collaboration":
		shape.ID = fmt.Sprintf("Participant_%s_di", suffix)
		break
	case "event":
		shape.ID = fmt.Sprintf("Event_%s_di", suffix)
		break
	case "participant":
		shape.ID = fmt.Sprintf("Participant_%s_di", suffix)
		break
	case "startevent":
		shape.ID = fmt.Sprintf("_BPMNShape_StartEvent_%d", suffix)
		break
	}

}

// SetElement ...
func (shape *Shape) SetElement(typ string, suffix interface{}) {
	switch typ {
	case "activity":
		shape.Element = fmt.Sprintf("Activity_%s", suffix)
		break
	case "collaboration":
		shape.Element = fmt.Sprintf("Participant_%s", suffix)
		break
	case "event":
		shape.Element = fmt.Sprintf("Event_%s", suffix)
		break
	case "participant":
		shape.Element = fmt.Sprintf("Participant_%s", suffix)
		break
	case "startevent":
		shape.Element = fmt.Sprintf("StartEvent_%d", suffix)
		break
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
