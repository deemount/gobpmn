package models

type Shape struct {
	ID           string `xml:"id,attr"`
	Element      string `xml:"bpmnElement,attr"`
	IsHorizontal string `xml:"isHorizontal,attr,omitempty"`
	Bounds       Bounds `xml:"dc:Bounds"`
}

// SetID ...
func (sh *Shape) SetID(typ string, suffix string) {
	switch typ {
	case "collaboration":
		sh.ID = "Participant_" + suffix + "_di"
	}

}

func (sh *Shape) SetElement(typ string, suffix string) {
	switch typ {
	case "collaboration":
		sh.Element = "Participant_" + suffix
	}
}
