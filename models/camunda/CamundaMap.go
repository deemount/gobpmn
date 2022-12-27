package camunda

// CamundaMapRepository ...
type CamundaMapRepository interface {
	SetCamundaEntry(num int)
	GetCamundaEntry(num int) *CamundaEntry
}

// CamundaMap ...
type CamundaMap struct {
	CamundaEntry []CamundaEntry `xml:"camunda:entry,omitempty" json:"entry,omitempty"`
}

// TCamundaMap ...
type TCamundaMap struct {
	Entry []CamundaEntry `xml:"entry,omitempty" json:"entry,omitempty"`
}

// NewCamundaMap ...
func NewCamundaMap() CamundaMapRepository {
	return &CamundaMap{}
}

/*
 * Default Setters
 */

/*** Make Elements ***/

/** Camunda **/

// SetCamundaEntry
func (mp *CamundaMap) SetCamundaEntry(num int) {
	mp.CamundaEntry = make([]CamundaEntry, num)
}

/*
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaEntry
func (mp CamundaMap) GetCamundaEntry(num int) *CamundaEntry {
	return &mp.CamundaEntry[num]
}
