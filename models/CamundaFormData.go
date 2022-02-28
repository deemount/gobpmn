package models

// CamundaFormData ...
type CamundaFormData struct {
	CamundaFormField []CamundaFormField `xml:"camunda:formField,omitempty" json:"formData,omitempty"`
}

/* Elements */

// SetCamundaFormField ...
func (cfd *CamundaFormData) SetCamundaFormField(num int) {
	cfd.CamundaFormField = make([]CamundaFormField, num)
}
