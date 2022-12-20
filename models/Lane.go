package models

import "fmt"

// LaneRepository ...
type LaneRepository interface {
	SetID(typ string, suffix interface{})

	SetFlowNodeRef(num int)

	GetID() *string

	GetFlowNodeRef(num int) *FlowNodeRef
}

// Lane ...
type Lane struct {
	ID          string        `xml:"id,attr" json:"id"`
	FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

func NewLane() LaneRepository {
	return &Lane{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (lane *Lane) SetID(typ string, suffix interface{}) {
	switch typ {
	case "lane":
		lane.ID = fmt.Sprintf("Lane_%v", suffix)
		break
	case "id":
		lane.ID = fmt.Sprintf("%s", suffix)
	}
}

/*** Make Elements ***/

/** BPMN **/

// SetFlowNodeRef ...
func (lane *Lane) SetFlowNodeRef(num int) {
	lane.FlowNodeRef = make([]FlowNodeRef, num)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (lane Lane) GetID() *string {
	return &lane.ID
}

/* Elements */

/** BPMN **/

// GetFlowNodeRef ...
func (lane Lane) GetFlowNodeRef(num int) *FlowNodeRef {
	return &lane.FlowNodeRef[num]
}
