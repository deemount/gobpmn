package models

// Participant ...
type Participant struct {
	ID         string `xml:"id,attr"`
	Name       string `xml:"name,attr,omitempty"`
	ProcessRef string `xml:"processRef,attr"`
}

// SetID ...
func (pp *Participant) SetID(suffix string) {
	pp.ID = "Participant_" + suffix
}

// SetName ...
func (pp *Participant) SetName(name string) {
	pp.Name = name
}

// SetProcessRef ...
func (pp *Participant) SetProcessRef(suffix string) {
	pp.ProcessRef = "Process_" + suffix
}
