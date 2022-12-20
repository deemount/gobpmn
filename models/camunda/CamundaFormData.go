package camunda

// CamundaFormData ...
type CamundaFormData struct {
	CamundaFormField []CamundaFormField `xml:"camunda:formField,omitempty" json:"formData,omitempty"`
}

/* Elements */

// SetCamundaFormField ...
func (formData *CamundaFormData) SetCamundaFormField(num int) {
	formData.CamundaFormField = make([]CamundaFormField, num)
}
