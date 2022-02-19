package models

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

// SetSourceRef ...
func (sf *SequenceFlow) SetSourceRef(sourceRef string) {
	sf.SourceRef = sourceRef
}

// SetTargetRef ...
func (sf *SequenceFlow) SetTargetRef(targetRef string) {
	sf.TargetRef = targetRef
}
