package models

// CamundaMap ...
type CamundaMap struct {
	CamundaEntry []CamundaEntry `xml:"camunda:entry,omitempty" json:"entry,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaEntry
func (mp *CamundaMap) SetCamundaEntry(num int) {
	mp.CamundaEntry = make([]CamundaEntry, num)
}
