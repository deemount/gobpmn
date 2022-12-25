package canvas

import (
	"fmt"
)

// Diagram ...
type DiagramRepository interface {
	CanvasID

	SetPlane()
	GetPlane() *Plane

	GetDescription() string
}

// Diagram ...
type Diagram struct {
	ID          string  `xml:"id,attr"`
	Description string  `xml:"-" json:"-"`
	Plane       []Plane `xml:"bpmndi:BPMNPlane,omitempty"`
}

func NewDiagram() DiagramRepository {
	return &Diagram{}
}

/* Attributes */

// SetID ...
// SetID(typ string, num int64)
func (diagram *Diagram) SetID(typ string, suffix interface{}) {
	switch typ {
	case "diagram":
		//diagram.ID = "BPMNDiagram_" + strconv.FormatInt(num, 16)
		diagram.ID = fmt.Sprintf("BPMNDiagram_%v", suffix)
		break
	case "id":
		diagram.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/* Elements */

/** BPMNDI **/

// SetPlane ...
func (diagram *Diagram) SetPlane() {
	diagram.Plane = make([]Plane, 1)
}

/**
 * Default Getters
 */

/* Attributes */

// GetID ...
func (diagram Diagram) GetID() STR_PTR {
	return &diagram.ID
}

/* Elements */

/** BPMN **/

// GetPlane ...
func (diagram Diagram) GetPlane() *Plane {
	return &diagram.Plane[0]
}

/* Specification */

// Description ...
func (diagram Diagram) GetDescription() string {
	diagram.Description = `
	BPMNDiagram represents a depiction of all or part of a BPMN model. It
	specializes DI::Diagram and redefines the root element (the top most diagram element)
	to be of type BPMNPlane. A BPMN diagram can also own a collection of BPMNStyle elements
	that are referenced by BPMNLabel elements in the diagram. These style elements represent
	the unique appearance styles used in the diagram.
	`
	return fmt.Sprintf("%s", diagram.Description)
}
