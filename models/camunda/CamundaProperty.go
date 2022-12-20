package camunda

// CamundaProperty ...
type CamundaProperty struct {
	Name  string `xml:"name,attr,omitempty" json:"name"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

/* Attributes */

// SetName ...
func (property *CamundaProperty) SetName(name string) {
	property.Name = name
}

// SetValue ...
func (property *CamundaProperty) SetValue(value string) {
	property.Value = value
}
