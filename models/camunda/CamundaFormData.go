package camunda

// CamundaFormDataRepository ...
type CamundaFormDataRepository interface {
	SetCamundaFormField(num int)
	GetCamundaFormField(num int) *CamundaFormField
}

// CamundaFormData ...
type CamundaFormData struct {
	CamundaFormField []CamundaFormField `xml:"camunda:formField,omitempty" json:"formData,omitempty"`
}

// TCamundaFormData ...
type TCamundaFormData struct {
	FormField []CamundaFormField `xml:"formField,omitempty" json:"formData,omitempty"`
}

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
