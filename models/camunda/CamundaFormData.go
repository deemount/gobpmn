package camunda

// NewCamundaFormData ...
func NewCamundaFormData() CamundaFormDataRepository {
	return &CamundaFormData{}
}

/*
 * Default Setters
 */

/*** Make Elements ***/

/** Camunda **/

// SetCamundaFormField ...
func (formData *CamundaFormData) SetCamundaFormField(num int) {
	formData.CamundaFormField = make([]CamundaFormField, num)
}

/*
 * Default Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaFormField ...
func (formData CamundaFormData) GetCamundaFormField(num int) *CamundaFormField {
	return &formData.CamundaFormField[num]
}
