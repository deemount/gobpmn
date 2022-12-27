package camunda

// CamundaPropertyRepository ...
type CamundaPropertyRepository interface {
	CamundaBaseName
	CamundaBaseValue
}

// CamundaProperty ...
type CamundaProperty struct {
	Name  string `xml:"name,attr,omitempty" json:"name"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

// NewCamundaProperty ...
func NewCamundaProperty() CamundaPropertyRepository {
	return &CamundaProperty{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetName ...
func (property *CamundaProperty) SetName(name string) {
	property.Name = name
}

// SetValue ...
func (property *CamundaProperty) SetValue(value string) {
	property.Value = value
}

/*
 * Default Getters
 */

/* Attributes */

// GetName ...
func (property CamundaProperty) GetName() STR_PTR {
	return &property.Name
}

// GetValue ...
func (property CamundaProperty) GetValue() STR_PTR {
	return &property.Value
}
