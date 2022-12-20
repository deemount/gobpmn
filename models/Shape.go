package models

import "fmt"

// ShapeRepository ...
type ShapeRepository interface {
	SetID(typ string, suffix interface{})
	SetElement(typ string, suffix interface{})
	SetIsHorizontal(isHorizontal bool)
	SetBounds()
	SetLabel()

	GetID() *string
	GetElement() *string
	GetIsHorizontal() *bool

	GetBounds() *Bounds
	GetLabel() *Label
}

// Shape ...
type Shape struct {
	ID           string   `xml:"id,attr" json:"-"`
	Element      string   `xml:"bpmnElement,attr" json:"-"`
	IsHorizontal bool     `xml:"isHorizontal,attr,omitempty" json:"-"`
	Bounds       []Bounds `xml:"dc:Bounds" json:"-"`
	Label        []Label  `xml:"bpmndi:BPMNLabel" json:"-"`
}

// TShape ...
type TShape struct {
	ID           string   `xml:"id,attr" json:"-"`
	Element      string   `xml:"bpmnElement,attr" json:"-"`
	IsHorizontal bool     `xml:"isHorizontal,attr,omitempty" json:"-"`
	Bounds       []Bounds `xml:"Bounds" json:"-"`
	Label        []Label  `xml:"BPMNLabel" json:"-"`
}

func NewShape() ShapeRepository {
	return &Shape{}
}

/**
 * Default Setters
 */

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
	case "id":
		shape.ID = fmt.Sprintf("%s_di", suffix)
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
	case "id":
		shape.Element = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetIsHorizontal ...
func (shape *Shape) SetIsHorizontal(isHorizontal bool) {
	shape.IsHorizontal = isHorizontal
}

/*** Make Elements ***/

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

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (shape Shape) GetID() *string {
	return &shape.ID
}

// GetElement ...
func (shape Shape) GetElement() *string {
	return &shape.Element
}

// GetIsHorizontal ...
func (shape Shape) GetIsHorizontal() *bool {
	return &shape.IsHorizontal
}

/* Elements */

/** DC **/

// SetBounds ...
func (shape Shape) GetBounds() *Bounds {
	return &shape.Bounds[0]
}

/** BPMNDI **/

// SetLabel ...
func (shape Shape) GetLabel() *Label {
	return &shape.Label[0]
}
