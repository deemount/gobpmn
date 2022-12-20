package camunda

import "fmt"

// CamundaFormField ...
type CamundaFormField struct {
	ID                string              `xml:"id,attr,omitempty" json:"id"`
	Label             string              `xml:"label,attr,omitempty" json:"label,omitempty"`
	Typ               string              `xml:"type,attr,omitempty" json:"type,omitempty"`
	DefaultValue      string              `xml:"defaultValue,attr,omitempty" json:"defaultValue,omitempty"`
	CamundaProperties []CamundaProperties `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
	CamundaValidation []CamundaValidation `xml:"camunda:validation,omitempty" json:"validation,omitempty"`
}

/* Attributes */

// SetID ...
func (formfield *CamundaFormField) SetID(suffix string) {
	formfield.ID = fmt.Sprintf("FormField_%s", suffix)
}

// SetLabel ...
func (formfield *CamundaFormField) SetLabel(label string) {
	formfield.Label = label
}

// SetType ...
// can be: string, long, boolean, date, enum, custom type
func (formfield *CamundaFormField) SetType(typ string) {
	formfield.Typ = typ
}

// SetDefaultValue ...
func (formfield *CamundaFormField) SetDefaultValue(defaultValue string) {
	formfield.DefaultValue = defaultValue
}

/* Elements */

/** Camunda **/

// SetCamundaProperties ...
func (formfield *CamundaFormField) SetCamundaProperties() {
	formfield.CamundaProperties = make([]CamundaProperties, 1)
}

// SetCamundaValidation ...
func (formfield *CamundaFormField) SetCamundaValidation() {
	formfield.CamundaValidation = make([]CamundaValidation, 1)
}
