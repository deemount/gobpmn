package models

import "strconv"

// SequenceFlow ...
type SequenceFlow struct {
	ID        string `xml:"id,attr"`
	SourceRef string `xml:"sourceRef,attr"`
	TargetRef string `xml:"targetRef,attr"`
}

// SetID ...
func (sf *SequenceFlow) SetID(suffix string) {
	sf.ID = "Flow_" + suffix
}

// SetStartEvent ...
func (sf *SequenceFlow) SetStartEvent(num int64) {
	sf.SourceRef = "StartEvent_" + strconv.FormatInt(num, 16)
}

// SetSourceRef ...
func (sf *SequenceFlow) SetSourceRef(name string, num int64) {
	sf.SourceRef = name + "_" + strconv.FormatInt(num, 16)
}

// SetEventRef ...
func (sf *SequenceFlow) SetEventRef(suffix string) {
	sf.TargetRef = "Event_" + suffix
}

// SetTargetRef ...
func (sf *SequenceFlow) SetTargetRef(name string, suffix string) {
	sf.TargetRef = name + "_" + suffix
}
