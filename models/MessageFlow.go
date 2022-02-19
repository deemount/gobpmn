package models

type MessageFlow struct {
	ID        string `xml:"id,attr"`
	SourceRef string `xml:"sourceRef,attr"`
	TargetRef string `xml:"targetRef,attr"`
}

// SetID ...
func (mf *MessageFlow) SetID(suffix string) {
	mf.ID = "Flow_" + suffix
}

// SetSourceRef ...
func (mf *MessageFlow) SetSourceRef(sourceRef string) {
	mf.SourceRef = sourceRef
}

// SetTargetRef ...
func (mf *MessageFlow) SetTargetRef(targetRef string) {
	mf.TargetRef = targetRef
}
