package camunda

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
