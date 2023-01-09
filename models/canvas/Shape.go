package canvas

import (
	"fmt"

	"github.com/deemount/gobpmn/models/compulsion"
)

// NewShape ...
func NewShape() ShapeRepository {
	return &Shape{}
}

/*
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

// SetIsMarkerVisible ...
func (shape *Shape) SetIsMarkerVisible(isMarkerVisible bool) {
	shape.IsMarkerVisible = isMarkerVisible
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

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (shape Shape) GetID() compulsion.STR_PTR {
	return &shape.ID
}

// GetElement ...
func (shape Shape) GetElement() compulsion.STR_PTR {
	return &shape.Element
}

// GetIsHorizontal ...
func (shape Shape) GetIsHorizontal() *bool {
	return &shape.IsHorizontal
}

// GetIsMarkerVisible ...
func (shape Shape) GetIsMarkerVisible() *bool {
	return &shape.IsMarkerVisible
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
