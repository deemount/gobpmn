package models

import "fmt"

// Lane ...
type Lane struct {
	ID          string        `xml:"id,attr"`
	FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (lane *Lane) SetID(suffix string) {
	lane.ID = fmt.Sprintf("Lane_%s", suffix)
}

/* Elements */

/** BPMN **/

// SetFlowNodeRef ...
func (lane *Lane) SetFlowNodeRef(num int) {
	lane.FlowNodeRef = make([]FlowNodeRef, num)
}
