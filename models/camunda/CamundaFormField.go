package camunda

import "fmt"

// CamundaFormField ...
type CamundaFormFieldRepository interface {
	CamundaBaseID
	CamundaBaseLabel
	CamundaBaseType
	SetDefaultValue(defaultValue string)
	GetDefaultValue() *string
	SetCamundaProperties()
	GetCamundaProperties() *CamundaProperties
	SetCamundaValidation()
	GetCamundaValidation() *CamundaValidation
}

// CamundaFormField ...
type CamundaFormField struct {
	ID                string              `xml:"id,attr,omitempty" json:"id"`
	Label             string              `xml:"label,attr,omitempty" json:"label,omitempty"`
	Typ               string              `xml:"type,attr,omitempty" json:"type,omitempty"`
	DefaultValue      string              `xml:"defaultValue,attr,omitempty" json:"defaultValue,omitempty"`
	CamundaProperties []CamundaProperties `xml:"camunda:properties,omitempty" json:"properties,omitempty"`
	CamundaValidation []CamundaValidation `xml:"camunda:validation,omitempty" json:"validation,omitempty"`
}

// TCamundaFormField ...
type TCamundaFormField struct {
	ID           string              `xml:"id,attr,omitempty" json:"id"`
	Label        string              `xml:"label,attr,omitempty" json:"label,omitempty"`
	Typ          string              `xml:"type,attr,omitempty" json:"type,omitempty"`
	DefaultValue string              `xml:"defaultValue,attr,omitempty" json:"defaultValue,omitempty"`
	Properties   []CamundaProperties `xml:"properties,omitempty" json:"properties,omitempty"`
	Validation   []CamundaValidation `xml:"validation,omitempty" json:"validation,omitempty"`
}

// NewCamundaFormField ...
func NewCamundaFormField() CamundaFormFieldRepository {
	return &CamundaFormField{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetID ...
func (formfield *CamundaFormField) SetID(typ string, suffix string) {
	switch typ {
	case "formfield":
		formfield.ID = fmt.Sprintf("FormField_%s", suffix)
		break
	case "id":
		formfield.ID = fmt.Sprintf("%s", suffix)
		break
	}
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

/*** Make Elements ***/

/** Camunda **/

// SetCamundaProperties ...
func (formfield *CamundaFormField) SetCamundaProperties() {
	formfield.CamundaProperties = make([]CamundaProperties, 1)
}

// SetCamundaValidation ...
func (formfield *CamundaFormField) SetCamundaValidation() {
	formfield.CamundaValidation = make([]CamundaValidation, 1)
}

/*
 * Default Getters
 */

/* Attributes */

// GetID ...
func (formfield CamundaFormField) GetID() STR_PTR {
	return &formfield.ID
}

// GetLabel ...
func (formfield CamundaFormField) GetLabel() *string {
	return &formfield.Label
}

// GetType ...
// can be: string, long, boolean, date, enum, custom type
func (formfield CamundaFormField) GetType() *string {
	return &formfield.Typ
}

// GetDefaultValue ...
func (formfield CamundaFormField) GetDefaultValue() *string {
	return &formfield.DefaultValue
}

/*** Make Elements ***/

/** Camunda **/

// GetCamundaProperties ...
func (formfield CamundaFormField) GetCamundaProperties() *CamundaProperties {
	return &formfield.CamundaProperties[0]
}

// GetCamundaValidation ...
func (formfield CamundaFormField) GetCamundaValidation() *CamundaValidation {
	return &formfield.CamundaValidation[0]
}
