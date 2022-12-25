package pool

import "fmt"

type LaneSetRepository interface {
	PoolID
	SetLane(num int)
	GetLane() *Lane
}

// LaneSet ...
type LaneSet struct {
	ID   string `xml:"id,attr" json:"id"`
	Lane []Lane `xml:"bpmn:lane,omitempty" json:"lane,omitempty"`
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ls *LaneSet) SetID(typ string, suffix interface{}) {
	switch typ {
	case "laneset":
		ls.ID = fmt.Sprintf("LaneSet_%v", suffix)
		break
	case "id":
		ls.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/* Elements */

/** BPMN **/

// SetLane ...
func (ls *LaneSet) SetLane(num int) {
	ls.Lane = make([]Lane, num)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ls LaneSet) GetID() STR_PTR {
	return &ls.ID
}

/* Elements */

/** BPMN **/

// GetLane ...
func (ls LaneSet) GetLane(num int) *Lane {
	return &ls.Lane[num]
}
