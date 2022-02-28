package models

import "fmt"

// LaneSet ...
type LaneSet struct {
	ID   string `xml:"id,attr" json:"id"`
	Lane []Lane `xml:"bpmn:lane,omitempty" json:"lane,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (ls *LaneSet) SetID(suffix string) {
	ls.ID = fmt.Sprintf("LaneSet_%s", suffix)
}

/* Elements */

/** BPMN **/

// SetLane ...
func (ls *LaneSet) SetLane(num int) {
	ls.Lane = make([]Lane, num)
}
