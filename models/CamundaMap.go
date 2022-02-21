package models

// CamundaMap ...
type CamundaMap struct {
	CamundaEntry []CamundaEntry `xml:"camunda:entry,omitempty"`
}

/* Elements */

/** Camunda **/

// SetCamundaEntry
func (cmap *CamundaMap) SetCamundaEntry(num int) {
	cmap.CamundaEntry = make([]CamundaEntry, num)
}
