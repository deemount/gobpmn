package models

import "fmt"

// CamundaFormField ...
type CamundaFormField struct {
	ID                string              `xml:"id,attr,omitempty"`
	Label             string              `xml:"label,attr,omitempty"`
	Typ               string              `xml:"type,attr,omitempty"`
	DefaultValue      string              `xml:"defaultValue,attr,omitempty"`
	CamundaProperties []CamundaProperties `xml:"camunda:properties,omitempty"`
	CamundaValidation []CamundaValidation `xml:"camunda:validation,omitempty"`
}

/* Attributes */

// SetID ...
func (cformfield *CamundaFormField) SetID(suffix string) {
	cformfield.ID = fmt.Sprintf("FormField_%s", suffix)
}

// SetLabel ...
func (cformfield *CamundaFormField) SetLabel(label string) {
	cformfield.Label = label
}

// SetType ...
// can be: string, long, boolean, date, enum, custom type
func (cformfield *CamundaFormField) SetType(typ string) {
	cformfield.Typ = typ
}

// SetDefaultValue ...
func (cformfield *CamundaFormField) SetDefaultValue(defaultValue string) {
	cformfield.DefaultValue = defaultValue
}

/* Elements */

/** Camunda **/

// SetCamundaProperties ...
func (cformfield *CamundaFormField) SetCamundaProperties() {
	cformfield.CamundaProperties = make([]CamundaProperties, 1)
}

// SetCamundaValidation ...
func (cformfield *CamundaFormField) SetCamundaValidation() {
	cformfield.CamundaValidation = make([]CamundaValidation, 1)
}
