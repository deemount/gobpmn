package models

import (
	"fmt"
	"strconv"
)

// Diagram ...
type Diagram struct {
	ID          string  `xml:"id,attr"`
	Description string  `xml:"-" json:"-"`
	Plane       []Plane `xml:"bpmndi:BPMNPlane,omitempty"`
}

/* Attributes */

// SetID ...
func (diagram *Diagram) SetID(num int64) {
	diagram.ID = "BPMNDiagram_" + strconv.FormatInt(num, 16)
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
	to beof type BPMNPlane. A BPMN diagram can also own a collection of BPMNStyle elements
	that are referenced by BPMNLabel elements in the diagram. These style elements represent
	the unique appearance styles used in the diagram.
	`
	return fmt.Sprintf("%s", diagram.Description)
}
