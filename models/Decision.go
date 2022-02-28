package models

// DecisionRepository ...
type DecisionRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// Decision ...
type Decision struct {
	ID            string          `xml:"id,attr" json:"id"`
	Name          string          `xml:"name,attr,omitempty" json:"name,omitempty"`
	DecisionTable []DecisionTable `xml:"decisionTable" json:"decisionTable,omitempty"`
}

// SetID ...
func (dec *Decision) SetID(suffix string) {
	dec.ID = "Decision_" + suffix
}

// SetName ...
func (dec *Decision) SetName(name string) {
	dec.Name = name
}
