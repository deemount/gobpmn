package pool

import (
	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// pool
type DelegateParameter struct {
	T string   // typ
	N string   // name
	H []string // hash
}

/*
 * Elementary
 */

// FlowNodeRef ...
type FlowNodeRef struct {
	impl.CoreInnerID
}

// Lane ...
type Lane struct {
	impl.CoreID
	FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// TLane ...
type TLane struct {
	impl.CoreID
	FlowNodeRef []FlowNodeRef `xml:"flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// LaneSet ...
type LaneSet struct {
	impl.CoreID
	Lane []Lane `xml:"bpmn:lane,omitempty" json:"lane,omitempty"`
}

// TLaneSet ...
type TLaneSet struct {
	impl.CoreID
	Lane []TLane `xml:"lane,omitempty" json:"lane,omitempty"`
}
