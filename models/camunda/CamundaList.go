package camunda

// CamundaListRepository ...
type CamundaListRepository interface {
	SetCamundaValue(num int)
	GetCamundaValue(num int) *CamundaValue
}

// CamundaList ...
type CamundaList struct {
	CamundaValue []CamundaValue `xml:"camunda:value,omitempty" json:"value,omitempty"`
}

// TCamundaList ...
type TCamundaList struct {
	Value []CamundaValue `xml:"value,omitempty" json:"value,omitempty"`
}

// NewCamundaList ...
func NewCamundaList() CamundaListRepository {
	return &CamundaList{}
}

/*
 * Default Setters
 */

/*** Make Elements ***/

/** Camunda **/

// SetCamundaValue ...
func (list *CamundaList) SetCamundaValue(num int) {
	list.CamundaValue = make([]CamundaValue, num)
}

/*
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaValue ...
func (list CamundaList) GetCamundaValue(num int) *CamundaValue {
	return &list.CamundaValue[num]
}
