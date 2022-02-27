package models

// DecisionRepository ...
type DecisionRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// Decision ...
type Decision struct {
	ID            string        `xml:"id,attr"`
	Name          string        `xml:"name,attr,omitempty"`
	DecisionTable DecisionTable `xml:"decisionTable"`
}

// SetID ...
func (dec *Decision) SetID(suffix string) {
	dec.ID = "Decision_" + suffix
}

// SetName ...
func (dec *Decision) SetName(name string) {
	dec.Name = name
}
