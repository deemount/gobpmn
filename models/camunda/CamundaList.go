package camunda

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
