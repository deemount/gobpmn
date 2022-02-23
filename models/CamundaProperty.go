package models

// CamundaProperty ...
type CamundaProperty struct {
	Name  string `xml:"name,attr,omitempty"`
	Value string `xml:"value,attr,omitempty"`
}

// SetName ...
func (cproperty *CamundaProperty) SetName(name string) {
	cproperty.Name = name
}

// SetValue ...
func (cproperty *CamundaProperty) SetValue(value string) {
	cproperty.Value = value
}
