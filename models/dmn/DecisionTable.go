package dmn

type DecisionTable struct {
	ID string `xml:"id,attr"`
}

// SetID ...
func (decTbl *DecisionTable) SetID(suffix string) {
	decTbl.ID = "DecisionTable_" + suffix
}
